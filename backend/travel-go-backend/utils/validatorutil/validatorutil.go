package validatorutil

import (
	"github.com/gin-gonic/gin/binding"
	english "github.com/go-playground/locales/en"
	vietnamese "github.com/go-playground/locales/vi"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	vi_trans "github.com/go-playground/validator/v10/translations/vi"
	"reflect"
	"sync"
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultValidator{}

var TransVi *ut.Translator
var TransEn *ut.Translator

func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		//vietname
		vie := vietnamese.New()
		uni := ut.New(vie, vie)
		transVi, _ := uni.GetTranslator("vi")

		//english
		en := english.New()
		uniEn := ut.New(en, en)
		transEn, _ := uniEn.GetTranslator("en")

		TransVi = &transVi
		TransEn = &transEn

		vi_trans.RegisterDefaultTranslations(v.validate, transVi)

		en_trans.RegisterDefaultTranslations(v.validate, transEn)
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func ArrError(errField []validator.FieldError, lang string) (errStr []string) {
	if lang == "en" {
		for _, e := range errField {
			errStr = append(errStr, e.Translate(*TransEn))
		}
		return
	}

	for _, e := range errField {
		errStr = append(errStr, e.Translate(*TransVi))
	}
	return
}
