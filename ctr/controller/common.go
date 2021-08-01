package controller

import (
	"errors"
	"reflect"
	"strings"

	"github.com/ahojcn/ecloud/ctr/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		uni := ut.New(zh.New())
		trans, _ = uni.GetTranslator("zh")
		//注册翻译器
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})
	}
}

type g struct {
	C *gin.Context
}

func newGin(c *gin.Context) *g {
	return &g{C: c}
}

type responseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// https://www.cnblogs.com/xiao-xue-di/p/14451923.html
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func (g *g) response(statusCode int, msg string, data interface{}) {
	errs, ok := data.(validator.ValidationErrors)
	if ok {
		data = removeTopStruct(errs.Translate(trans))
	}
	err, ok := data.(error)
	if ok {
		data = err.Error()
	}

	g.C.JSON(statusCode, responseData{
		Code: statusCode,
		Msg:  msg,
		Data: data,
	})
}

func (g *g) loginRequired() (*model.User, error) {
	sess := sessions.Default(g.C)
	username := sess.Get("username")
	if username == nil {
		return nil, errors.New("登录信息过期")
		//username = "ahojcn"
	}

	user, has := model.UserOne(map[string]interface{}{"username": username})
	if !has {
		return nil, errors.New("未查找到对应用户")
	}

	return user, nil
}
