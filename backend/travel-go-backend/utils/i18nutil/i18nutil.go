package i18nutil

import (
	"github.com/gin-gonic/gin"
	i18n_bundle "github.com/nicksnyder/go-i18n/v2/i18n"
	i18n_lib "github.com/nicksnyder/go-i18n/v2/i18n"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_request"
	"golang.org/x/text/language"
	"sync"
)

var (
	bundle  *i18n_bundle.Bundle
	syncOne sync.Once
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func NewBundle() {
	syncOne.Do(func() {
		bundle = i18n_bundle.NewBundle(language.English)
		englishDict := mergeDict()
		vietnamDict := mergeDict()

		for _, v := range englishDict {
			must(bundle.AddMessages(language.English, &v))
		}

		for _, v := range vietnamDict {
			must(bundle.AddMessages(language.Vietnamese, &v))
		}
	})
}

func mergeDict(args ...[]i18n_bundle.Message) []i18n_bundle.Message {
	mergedDict := make([]i18n_bundle.Message, 0)
	for _, oneDict := range args {
		mergedDict = append(mergedDict, oneDict...)
	}

	return mergedDict
}

func NewLocalizer(langs ...string) *i18n_bundle.Localizer {
	return i18n_bundle.NewLocalizer(bundle, langs...)
}

func GetI18nMessage(id string, c *gin.Context) (string, error) {
	localizer := NewLocalizer(api_request.GetRequestLanguage(c))
	return localizer.Localize(&i18n_lib.LocalizeConfig{MessageID: id})
}
