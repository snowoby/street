package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("sid", sid)
		panic(err)
	}
}

func sid(fl validator.FieldLevel) bool {
	str, ok := fl.Field().Interface().(string)
	if ok {
		if len(str) < 10 {
			for _, r := range str {
				return !(r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122)
			}
		}
	}
	return true
}
