package services

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var LOCALIZER *i18n.Localizer

func InitLocalizer() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("./resources/lang/pt_BR.toml")
	LOCALIZER = i18n.NewLocalizer(bundle, "pt_BR")
}

func Translate(key string, templateData map[string]string) string {
	if LOCALIZER == nil {
		InitLocalizer()
	}
	translateString, err := LOCALIZER.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: templateData,
	})
	if err != nil {
		return key
	} else {
		return translateString
	}
}
