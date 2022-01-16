package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		//修改Gin框架中的validator引擎属性，实现定制
		if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
			//注册一个获取json的tag自定义方法
			validate.RegisterTagNameFunc(func(field reflect.StructField) string {
				name := strings.Split(field.Tag.Get("json"), ",")[0]
				//处理特殊的json 例如： json:"-"方法，这种不应该处理
				if name == "-" {
					return ""
				}
				return name
			})
			zhT := zh.New() //中文翻译器
			enT := en.New() //英文翻译器
			//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
			uni := ut.New(enT, zhT, enT)

			locale := c.GetHeader("locale")
			if locale == "" { //如果header里面没有指定语言，则默认使用英文
				locale = "en"
			}
			trans, _ := uni.GetTranslator(locale)
			switch locale {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(validate, trans)
			case "en":
				_ = en_translations.RegisterDefaultTranslations(validate, trans)
			default:
				_ = en_translations.RegisterDefaultTranslations(validate, trans)
			}
			c.Set("trans", trans) //保存，用于后续使用
		}
		c.Next()
	}
}
