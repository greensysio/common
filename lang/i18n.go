package lang

import (
	"github.com/greensysio/common/log"
	"os"

	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/yaml"
)

// I18n global transaltion
var tr *i18n.I18n

// InitI18N : Init i18n which was used as global
func InitI18N(path string) {
	i18n.Default = "vn"
	if path == "" {
		path = "./config/i18n/"
	}

	if tr == nil {
		tr = i18n.New(
			yaml.New(path), // load translations from the YAML files in directory `lang/locales`
		)
		if string(tr.T("VN", "jwt.missing")) == "" {
			log.Error("Init I18N failed!")
			os.Exit(1)
		}
	}
	log.Info("Init I18N successfully!")
}

// I18n : use for translation
func I18n() *i18n.I18n {
	return tr
}
