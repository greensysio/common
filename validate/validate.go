package validate

import (
	"fmt"
	"regexp"

	"github.com/greensysio/common/log"
	"github.com/asaskevich/govalidator"
)

func init() {
	log.InitLogger(false)
}

// InitCustomValidator init custom validator
func InitCustomValidator() {
	govalidator.TagMap["emoji"] = govalidator.Validator(func(str string) bool {
		// Regex uses from https://gist.github.com/hnq90/316f08047a3bf348b823
		var rg = regexp.MustCompile(`[\x{1F600}-\x{1F64F}]|[\x{1F300}-\x{1F5FF}]|[\x{1F680}-\x{1F6FF}]|[\x{2600}-\x{26FF}]|[\x{2700}-\x{27BF}]`)
		return !rg.MatchString(str)
	})

	govalidator.CustomTypeTagMap.Set("checkLongitude", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		return govalidator.IsLongitude(fmt.Sprintf("%f", i.(float64)))
	}))

	govalidator.CustomTypeTagMap.Set("checkLatitude", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		return govalidator.IsLatitude(fmt.Sprintf("%f", i.(float64)))
	}))
	log.Print("Init custom validator successfully!")
}
