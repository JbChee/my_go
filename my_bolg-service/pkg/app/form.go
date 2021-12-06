package app

import (
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
)

type ValidError struct {
	Key     string
	Message string
}



type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	fmt.Printf("[debug] BindAndValid: v: %v, err: %v, path: %v, \n", v, err, c.FullPath())
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		//fmt.Println(err.Error())
		verrs,ok := err.(validator.ValidationErrors)
		//fmt.Println(err.Error())
		//(validator.ValidationErrors{})
		if !ok {
			return true, nil
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key: key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}