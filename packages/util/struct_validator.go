package util

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	validatorPkg "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var errorMessage = map[string]string{
	"required": "The {field} field is required.",
	"email":    "The {field} field must be a valid email address.",
	"max":      "The {field} field must be at most {param} characters.",
	"min":      "The {field} field must be at least {param} characters.",
	"digits":   "The {field} must be {param} digits.",
	"len":      "The {field} must be {param} digits.",
	"number":   "The {field} must be {param} digits.",
	"regex":    "The {field} field is not in the correct format.",
	"thai":     "The {field} field must be only in Thai.",
}

var thaiErrorMessage = map[string]string{
	"required": "จำเป็นต้องกรอกข้อมูล {field}",
	"email":    "{field} ต้องเป็นอีเมลที่ถูกต้อง",
	"max":      "{field} ต้องมีความยาวไม่เกิน {param} ตัวอักษร",
	"min":      "{field} ต้องมีความยาวอย่างน้อย {param} ตัวอักษร",
	"digits":   "{field} ต้องเป็นตัวเลข {param} หลัก",
	"len":      "{field} ต้องมีความยาว {param} หลัก",
	"number":   "{field} ต้องเป็นตัวเลข {param} หลัก",
	"regex":    "รูปแบบของ {field} ไม่ถูกต้อง",
	"thai":     "{field} ต้องเป็นภาษาไทยเท่านั้น",
}

type StructValidationError struct {
	FieldName      string `json:"field_name"`
	ErrorMessageTH string `json:"message_TH"`
	ErrorMessageEN string `json:"message_EN"`
}

func (s *StructValidationError) Error() string {
	return s.ErrorMessageTH
}

func ValidateStruct(structInput interface{}) []StructValidationError {
	validator := validatorPkg.New()
	registerValidation(validator)
	errValidate := validator.Struct(structInput)
	var errorResult []StructValidationError
	if errValidate != nil {
		validationErrors := errValidate.(validatorPkg.ValidationErrors)
		for _, validationError := range validationErrors {
			field, errMsg, errMsgEN := getJSONFieldName(reflect.TypeOf(structInput), validationError.StructNamespace())
			messageTH := formatThaiErrorMessage(validationError.Tag(), field, validationError.Param(), errMsg)
			messageEN := formatErrorMessage(validationError.Tag(), field, validationError.Param(), errMsgEN)
			structError := StructValidationError{
				FieldName:      field,
				ErrorMessageTH: messageTH,
				ErrorMessageEN: messageEN,
			}
			errorResult = append(errorResult, structError)
		}
		return errorResult
	}
	return nil
}

// ค้นหาชื่อ feild จาก feild path
func getJSONFieldName(structType reflect.Type, fieldPath string) (string, string, string) {

	fieldSlice := getFieldByPath(structType, fieldPath)
	var fieldJsonNameSlice []string
	for _, field := range fieldSlice {
		fieldJsonNameSlice = append(fieldJsonNameSlice, field.Tag.Get("json"))
	}
	errMsg := fieldSlice[len(fieldJsonNameSlice)-1].Tag.Get("errMsgTH")
	errMsgEN := fieldSlice[len(fieldJsonNameSlice)-1].Tag.Get("errMsgEN")
	fmt.Println("Test json", fieldJsonNameSlice)
	return strings.Join(fieldJsonNameSlice, "/"), errMsg, errMsgEN

}

func getFieldByPath(structType reflect.Type, fieldPath string) []reflect.StructField {
	fieldPathSlice := strings.Split(fieldPath, ".")
	var structFieldSlice []reflect.StructField
	for _, str := range fieldPathSlice {
		if structType.Kind() == reflect.Ptr {
			structType = structType.Elem()
		}
		field, found := structType.FieldByName(str)
		if found {
			structFieldSlice = append(structFieldSlice, field)
			structType = field.Type
		}
	}
	return structFieldSlice
}

// จักการ error message
func formatErrorMessage(tag string, field string, param string, errMsg string) string {
	tmpl, found := errorMessage[tag]
	if !found {
		tmpl = "Validation failed on the {field} field."
	}
	if errMsg != "" {
		return errMsg
	}

	msg := strings.ReplaceAll(tmpl, "{field}", strings.ReplaceAll(field, "_", " "))
	msg = strings.ReplaceAll(msg, "{param}", param)

	return msg
}

func formatThaiErrorMessage(tag string, field string, param string, customMsg string) string {
	if customMsg != "" {
		return customMsg
	}

	message, exists := thaiErrorMessage[tag]
	if !exists {
		message = "ข้อมูล {field} ไม่ถูกต้อง"
	}

	message = strings.Replace(message, "{field}", field, -1)
	message = strings.Replace(message, "{param}", param, -1)

	return message
}

func registerValidation(validator *validatorPkg.Validate) {
	validator.RegisterValidation("regex", regexValidation)
	validator.RegisterValidation("thai", thaiCharacterValidation)
	validator.RegisterValidation("uuid_default", uuidDefault)
}

func regexValidation(fl validatorPkg.FieldLevel) bool {
	regex := fl.Param()
	value := fl.Field().String()
	re := regexp.MustCompile(regex)
	return re.MatchString(value)
}

func thaiCharacterValidation(fl validatorPkg.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^[\x{0E00}-\x{0E7F}\s]+$`)
	return re.MatchString(value)
}

func uuidDefault(f1 validatorPkg.FieldLevel) bool {
	if f1.Field().String() == "" {
		if f1.Field().CanSet() {
			f1.Field().SetString(uuid.NewString())
		}
	}
	return true
}
func ThaiStringCharacterValidation(value string) bool {
	re := regexp.MustCompile(`^[\x{0E00}-\x{0E7F}.\s]+$`)

	return re.MatchString(value)
}
func IsNumeric(s string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(s)
}
func HasNoAlphabeticChars(s string) bool {
	re := regexp.MustCompile(`[a-zA-Z]`)
	return !re.MatchString(s)
}
