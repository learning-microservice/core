package validator

import (
	"encoding/json"
	"reflect"
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

func NewStructValidator() StructValidator {
	return &structValidator{}
}

type StructValidator interface {
	Validate(obj interface{}) error
	ValidateStruct(obj interface{}) error
}

type structValidator struct {
	once     sync.Once
	trans    ut.Translator
	validate *validator.Validate
}

func (v *structValidator) Validate(obj interface{}) error {
	return v.ValidateStruct(obj)
}

func (v *structValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return &validatorError{
				err:   err,
				trans: v.trans,
			}
		}
	}
	return nil
}

func (v *structValidator) lazyinit() {
	v.once.Do(func() {
		en := en.New()
		uni := ut.New(en, en)

		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		trans, _ := uni.GetTranslator("en")

		validate := validator.New()
		validate.SetTagName("binding")
		en_translations.RegisterDefaultTranslations(validate, trans)

		v.trans = trans
		v.validate = validate
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

type validatorError struct {
	err   error
	trans ut.Translator
}

func (v *validatorError) Error() string {
	if v.err != nil {
		return v.err.Error()
	}
	return "unknown error"
}

func (v *validatorError) MarshalJSON() ([]byte, error) {
	var messages []string
	for _, e := range v.err.(validator.ValidationErrors) {
		messages = append(messages, e.Translate(v.trans))
	}
	return json.Marshal(messages)
}
