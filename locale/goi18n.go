package locale

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
)

var bundle *i18n.Bundle
var localizer *i18n.Localizer

func Init() {
	setLangPreferences()
	setLanguageBundle()
	fmt.Println(Localize("helloWorld"))
}

func setLanguageBundle() *i18n.Localizer {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("resources/en.json")
	bundle.LoadMessageFile("resources/fr.json")
	if ok := os.Getenv("lang"); ok == "" {
		os.Setenv("lang", "en")
	}
	localizer = i18n.NewLocalizer(bundle, os.Getenv("lang"))
	return localizer
}

func setLangPreferences() {
	if _, ok := os.LookupEnv("lang"); !ok {
		os.Setenv("lang", "en")
	}

}

func Localize(valToLocalize string) string {
	localizeConfig := i18n.LocalizeConfig{
		MessageID: valToLocalize,
	}
	localization, _ := localizer.Localize(&localizeConfig)
	return localization
}
