package validate

import "github.com/go-playground/validator"

var validate = validator.New()

// ErrorResponse 异常错误声明
type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

// Do 验证请求数据
func Do(params interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(params)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
