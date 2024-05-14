package helpers

import "github.com/go-playground/validator/v10"

var validate = validator.New()
type ErrorResponse struct{
	FailedFields string `json:"failed_fields"`
	Tag string `json:"tag"`
	Value string `json:"value"`
}
func ValidateStruct(param any) []*ErrorResponse {
var errors []*ErrorResponse
err := validate.Struct(param)
if err != nil {
	for _, err := range err.(validator.ValidationErrors){
		element:=ErrorResponse{
			FailedFields: err.StructNamespace(),
			Tag: err.Tag(),
			Value: err.Param(),
		}
		errors = append(errors, &element)
	}
}
return errors
}