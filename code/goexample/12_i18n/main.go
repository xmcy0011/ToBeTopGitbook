package main

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

func main() {
	// new bundle
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// load from file
	langs := []language.Tag{
		language.AmericanEnglish,   // en-US
		language.SimplifiedChinese, // zh-Hans
	}
	for _, tag := range langs {
		bundle.MustLoadMessageFile(fmt.Sprintf("12_i18n/%s.json", tag.String()))
	}

	// new Localizer for every language
	local := map[string]*i18n.Localizer{
		language.AmericanEnglish.String():   i18n.NewLocalizer(bundle, language.AmericanEnglish.String()),
		language.SimplifiedChinese.String(): i18n.NewLocalizer(bundle, language.SimplifiedChinese.String()),
	}

	// get localizer from lang
	deviceLang := "zh-Hans"
	if localizer, ok := local[deviceLang]; ok {

		// translate
		content, err := localizer.Localize(&i18n.LocalizeConfig{
			TemplateData: map[string]interface{}{
				"Nickname2": "小王",
			},
			DefaultMessage: &i18n.Message{
				ID:         "friend.request.add",
				LeftDelim:  "{{",
				RightDelim: "}}",
			},
		})

		if err != nil {
			panic(err)
		} else {
			log.Println(content)
		}
	} else {
		log.Println("language ", deviceLang, " not found")
	}
}
