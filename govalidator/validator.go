package govalidator

import (
	"bitbucket.org/greensys-tech/common/lang"
	"fmt"
	"github.com/asaskevich/govalidator"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

// ValidateStruct use tags for fields.
// result will be equal to `false` if there are any errors.
func ValidateStruct(s interface{}, locale string) (bool, error) {
	if s == nil {
		return true, nil
	}
	result := true
	var err error
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	// we only accept structs
	if val.Kind() != reflect.Struct {
		return false, fmt.Errorf("function only accepts structs; got %s", val.Kind())
	}
	var errs govalidator.Errors
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if typeField.PkgPath != "" {
			continue // Private field
		}
		structResult := true
		if (valueField.Kind() == reflect.Struct ||
			(valueField.Kind() == reflect.Ptr && valueField.Elem().Kind() == reflect.Struct)) &&
			typeField.Tag.Get(tagName) != "-" {
			var err error
			structResult, err = ValidateStruct(valueField.Interface(), locale)
			if err != nil {
				errs = append(errs, err)
			}
		}
		resultField, err2 := typeCheck(valueField, typeField, val, nil, locale)
		if err2 != nil {

			// Replace structure name with JSON name if there is a tag on the variable
			jsonTag := toJSONName(typeField.Tag.Get("json"))
			if jsonTag != "" {
				switch jsonError := err2.(type) {
				case govalidator.Error:
					jsonError.Name = getAttributeName(locale, jsonTag)
					err2 = jsonError
				case govalidator.Errors:
					for i2, err3 := range jsonError {
						switch customErr := err3.(type) {
						case govalidator.Error:
							customErr.Name = getAttributeName(locale, jsonTag)
							jsonError[i2] = customErr
						}
					}

					err2 = jsonError
				}
			}

			errs = append(errs, err2)
		}
		result = result && resultField && structResult
	}
	if len(errs) > 0 {
		err = errs
	}
	return result, err
}

func toJSONName(tag string) string {
	if tag == "" {
		return ""
	}

	// JSON name always comes first. If there's no options then split[0] is
	// JSON name, if JSON name is not set, then split[0] is an empty string.
	split := strings.SplitN(tag, ",", 2)

	name := split[0]

	// However it is possible that the field is skipped when
	// (de-)serializing from/to JSON, in which case assume that there is no
	// tag name to use
	if name == "-" {
		return ""
	}
	return name
}

func typeCheck(v reflect.Value, t reflect.StructField, o reflect.Value, options tagOptionsMap, locale string) (isValid bool, resultErr error) {
	if !v.IsValid() {
		return false, nil
	}

	tag := t.Tag.Get(tagName)

	// Check if the field should be ignored
	switch tag {
	case "":
		if !fieldsRequiredByDefault {
			return true, nil
		}
		return false, govalidator.Error{
			Name: t.Name,
			Err:  fmt.Errorf(string(lang.I18n().T(locale, "valid.msg.at_least_one_validation_defined"))),
			CustomErrorMessageExists: false,
			Validator:                "required",
		}
	case "-":
		return true, nil
	}

	isRootType := false
	if options == nil {
		isRootType = true
		options = parseTagIntoMap(tag)
	}

	if isEmptyValue(v) {
		// an empty value is not validated, check only required
		return checkRequired(v, t, options, locale)
	}

	var customTypeErrors govalidator.Errors
	for validatorName, customErrorMessage := range options {
		if validatefunc, ok := govalidator.CustomTypeTagMap.Get(validatorName); ok {
			delete(options, validatorName)

			if result := validatefunc(v.Interface(), o.Interface()); !result {
				if len(customErrorMessage) > 0 {
					customTypeErrors = append(customTypeErrors, govalidator.Error{
						Name: t.Name,
						Err:  fmt.Errorf(customErrorMessage),
						CustomErrorMessageExists: true,
						Validator:                stripParams(validatorName),
					})
					continue
				}
				customTypeErrors = append(customTypeErrors, govalidator.Error{
					Name: getAttributeName(locale, t.Name),
					Err:  formatError(locale, v, validatorName, false),
					CustomErrorMessageExists: false,
					Validator:                stripParams(validatorName),
				})
			}
		}
	}

	if len(customTypeErrors.Errors()) > 0 {
		return false, customTypeErrors
	}

	if isRootType {
		// Ensure that we've checked the value by all specified validators before report that the value is valid
		defer func() {
			delete(options, "optional")
			delete(options, "required")

			if isValid && resultErr == nil && len(options) != 0 {
				for validator := range options {
					isValid = false
					resultErr = govalidator.Error{
						Name: getAttributeName(locale, t.Name),
						Err: fmt.Errorf(string(lang.I18n().T(locale, "valid.msg.can_not_apply_validator", map[string]string{
							"field": validator,
						}))),
						CustomErrorMessageExists: false,
						Validator:                stripParams(validator),
					}
					return
				}
			}
		}()
	}

	switch v.Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.String:
		// for each tag option check the map of validator functions
		for validatorSpec, customErrorMessage := range options {
			var negate bool
			validator := validatorSpec
			customMsgExists := len(customErrorMessage) > 0

			// Check whether the tag looks like '!something' or 'something'
			if validator[0] == '!' {
				validator = validator[1:]
				negate = true
			}

			// Check for param validators
			for key, value := range govalidator.ParamTagRegexMap {
				ps := value.FindStringSubmatch(validator)
				if len(ps) == 0 {
					continue
				}

				validatefunc, ok := govalidator.ParamTagMap[key]
				if !ok {
					continue
				}

				delete(options, validatorSpec)

				switch v.Kind() {
				case reflect.String,
					reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
					reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
					reflect.Float32, reflect.Float64:

					field := fmt.Sprint(v) // make value into string, then validate with regex
					if result := validatefunc(field, ps[1:]...); (!result && !negate) || (result && negate) {
						if customMsgExists {
							return false, govalidator.Error{
								Name: t.Name,
								Err:  fmt.Errorf(customErrorMessage),
								CustomErrorMessageExists: customMsgExists,
								Validator:                stripParams(validatorSpec),
							}
						}
						if negate {
							return false, govalidator.Error{
								Name: getAttributeName(locale, t.Name),
								Err:  formatError(locale, v, validator, true),
								CustomErrorMessageExists: customMsgExists,
								Validator:                stripParams(validatorSpec),
							}
						}
						return false, govalidator.Error{
							Name: getAttributeName(locale, t.Name),
							Err:  formatError(locale, v, validator, false),
							CustomErrorMessageExists: customMsgExists,
							Validator:                stripParams(validatorSpec),
						}
					}
				default:
					// type not yet supported, fail
					return false, govalidator.Error{
						Name: getAttributeName(locale, t.Name),
						Err: fmt.Errorf(string(lang.I18n().T(locale, "valid.msg.can_not_support", map[string]string{
							"validator": validator,
							"kind":      v.Kind().String(),
						}))),
						CustomErrorMessageExists: false,
						Validator:                stripParams(validatorSpec),
					}
				}
			}

			if validatefunc, ok := govalidator.TagMap[validator]; ok {
				delete(options, validatorSpec)

				switch v.Kind() {
				case reflect.String:
					field := fmt.Sprint(v) // make value into string, then validate with regex
					if result := validatefunc(field); !result && !negate || result && negate {
						if customMsgExists {
							return false, govalidator.Error{
								Name: t.Name,
								Err:  fmt.Errorf(customErrorMessage),
								CustomErrorMessageExists: customMsgExists,
								Validator:                stripParams(validatorSpec),
							}
						}
						if negate {
							return false, govalidator.Error{
								Name: getAttributeName(locale, t.Name),
								Err:  formatError(locale, v, validator, true),
								CustomErrorMessageExists: customMsgExists,
								Validator:                stripParams(validatorSpec),
							}
						}
						customTypeErrors = append(customTypeErrors, govalidator.Error{
							Name: getAttributeName(locale, t.Name),
							Err:  formatError(locale, v, validator, false),
							CustomErrorMessageExists: customMsgExists,
							Validator:                stripParams(validatorSpec),
						})
					}
				default:
					// Not Yet Supported Types (Fail here!)
					return false, govalidator.Error{
						Name: getAttributeName(locale, t.Name),
						Err: fmt.Errorf(string(lang.I18n().T(locale, "valid.msg.can_not_support", map[string]string{
							"validator": validator,
							"kind":      v.Kind().String(),
						}))),
						CustomErrorMessageExists: false,
						Validator:                stripParams(validatorSpec),
					}
				}
			}
		}
		return true, nil
	case reflect.Map:
		if v.Type().Key().Kind() != reflect.String {
			return false, &govalidator.UnsupportedTypeError{Type: v.Type()}
		}
		var sv stringValues
		sv = v.MapKeys()
		sort.Sort(sv)
		result := true
		for _, k := range sv {
			var resultItem bool
			var err error
			if v.MapIndex(k).Kind() != reflect.Struct {
				resultItem, err = typeCheck(v.MapIndex(k), t, o, options, locale)
				if err != nil {
					return false, err
				}
			} else {
				resultItem, err = ValidateStruct(v.MapIndex(k).Interface(), locale)
				if err != nil {
					return false, err
				}
			}
			result = result && resultItem
		}
		return result, nil
	case reflect.Slice, reflect.Array:
		result := true
		for i := 0; i < v.Len(); i++ {
			var resultItem bool
			var err error
			if v.Index(i).Kind() != reflect.Struct {
				resultItem, err = typeCheck(v.Index(i), t, o, options, locale)
				if err != nil {
					return false, err
				}
			} else {
				resultItem, err = ValidateStruct(v.Index(i).Interface(), locale)
				if err != nil {
					return false, err
				}
			}
			result = result && resultItem
		}
		return result, nil
	case reflect.Interface:
		// If the value is an interface then encode its element
		if v.IsNil() {
			return true, nil
		}
		return ValidateStruct(v.Interface(), locale)
	case reflect.Ptr:
		// If the value is a pointer then check its element
		if v.IsNil() {
			return true, nil
		}
		return typeCheck(v.Elem(), t, o, options, locale)
	case reflect.Struct:
		return ValidateStruct(v.Interface(), locale)
	default:
		return false, &govalidator.UnsupportedTypeError{Type: v.Type()}
	}
}

func formatError(locale string, value reflect.Value, validator string, negate bool) error {
	if strings.Contains(validator, "range(") {
		re := regexp.MustCompile("[0-9]+")
		ranges := re.FindAllString(validator, -1)
		path := "valid.msg.not_in_range"
		if negate {
			path = "valid.msg.negate_not_in_range"
		}
		return fmt.Errorf(string(lang.I18n().T(locale, path, map[string]string{
			"value": fmt.Sprintf("%.f", value),
			"min":   ranges[0],
			"max":   ranges[1],
		})))
	}
	if strings.Contains(validator, "length(") {
		re := regexp.MustCompile("[0-9]+")
		ranges := re.FindAllString(validator, -1)
		path := "valid.msg.not_in_length"
		if negate {
			path = "valid.msg.negate_not_in_length"
		}
		return fmt.Errorf(string(lang.I18n().T(locale, path, map[string]string{
			"value": fmt.Sprint(value),
			"min":   ranges[0],
			"max":   ranges[1],
		})))
	}
	if strings.Contains(validator, "runelength(") {
		re := regexp.MustCompile("[0-9]+")
		ranges := re.FindAllString(validator, -1)
		path := "valid.msg.not_in_runelength"
		if negate {
			path = "valid.msg.negate_not_in_runelength"
		}
		return fmt.Errorf(string(lang.I18n().T(locale, path, map[string]string{
			"value": fmt.Sprint(value),
			"min":   ranges[0],
			"max":   ranges[1],
		})))
	}
	if strings.Contains(validator, "in(") {
		t := strings.Replace(validator, "in(", "", -1)
		t = strings.Replace(t, ")", "", -1)
		path := "valid.msg.not_in_array"
		if negate {
			path = "valid.msg.negate_not_in_array"
		}
		return fmt.Errorf(string(lang.I18n().T(locale, path, map[string]string{
			"value":  fmt.Sprint(value),
			"ranges": t,
		})))
	}
	return fmt.Errorf(string(lang.I18n().T(locale, "valid.msg.can_not_format")))
}

func isValidTag(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		switch {
		case strings.ContainsRune("\\'\"!#$%&()*+-./:<=>?@[]^_{|}~ ", c):
			// Backslash and quote chars are reserved, but
			// otherwise any punctuation chars are allowed
			// in a tag name.
		default:
			if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
				return false
			}
		}
	}
	return true
}

func stripParams(validatorString string) string {
	return paramsRegexp.ReplaceAllString(validatorString, "")
}
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

func checkRequired(v reflect.Value, t reflect.StructField, options tagOptionsMap, locale string) (bool, error) {
	if requiredOption, isRequired := options["required"]; isRequired {
		if len(requiredOption) > 0 {
			return false, govalidator.Error{
				Name: t.Name,
				Err:  fmt.Errorf(requiredOption),
				CustomErrorMessageExists: true,
				Validator:                "required",
			}
		}
		return false, govalidator.Error{
			Name: getAttributeName(locale, t.Name),
			Err:  fmt.Errorf(string(lang.I18n().T(locale, "valid.msg.non_zero_value_required"))),
			CustomErrorMessageExists: false,
			Validator:                "required",
		}
	} else if _, isOptional := options["optional"]; fieldsRequiredByDefault && !isOptional {
		return false, govalidator.Error{
			Name: getAttributeName(locale, t.Name),
			Err:  fmt.Errorf(string(lang.I18n().T(locale, "valid.msg.missing_required_field"))),
			CustomErrorMessageExists: false,
			Validator:                "required",
		}
	}
	// not required and empty is valid
	return true, nil
}

// parseTagIntoMap parses a struct tag `valid:required~Some error message,length(2|3)` into map[string]string{"required": "Some error message", "length(2|3)": ""}
func parseTagIntoMap(tag string) tagOptionsMap {
	optionsMap := make(tagOptionsMap)
	options := strings.Split(tag, ",")

	for _, option := range options {
		option = strings.TrimSpace(option)

		validationOptions := strings.Split(option, "~")
		if !isValidTag(validationOptions[0]) {
			continue
		}
		if len(validationOptions) == 2 {
			optionsMap[validationOptions[0]] = validationOptions[1]
		} else {
			optionsMap[validationOptions[0]] = ""
		}
	}
	return optionsMap
}

func getAttributeName(locale string, attName string) string {
	name := string(lang.I18n().T(locale, fmt.Sprintf("valid.field.%s", strings.ToLower(attName))))
	if name == "" {
		name = attName
	}
	return name
}
