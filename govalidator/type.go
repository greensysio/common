package govalidator

import (
	"reflect"
	"regexp"
)

var (
	fieldsRequiredByDefault bool
	paramsRegexp            = regexp.MustCompile("\\(.*\\)$")
)

// Basic regular expressions for validating strings
const (
	tagName string = "valid"
)

type (
	tagOptionsMap map[string]string
	stringValues  []reflect.Value
)

func (sv stringValues) Len() int           { return len(sv) }
func (sv stringValues) Swap(i, j int)      { sv[i], sv[j] = sv[j], sv[i] }
func (sv stringValues) Less(i, j int) bool { return sv.get(i) < sv.get(j) }
func (sv stringValues) get(i int) string   { return sv[i].String() }
