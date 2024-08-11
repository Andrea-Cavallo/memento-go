package utils

import (
	"github.com/fatih/structs"
	"mememento-go/internal/models"
	"mememento-go/internal/utils/constants"
	"reflect"
	"time"
)

// NewPerson creates a new instance of models.Person
func NewPerson(email string) models.Person {
	return models.Person{
		Name:      constants.VALIDNAME,
		Age:       constants.ETA,
		Email:     email,
		LocalDate: time.Now().UnixMilli(),
	}
}

// convertToMap converts the provided input to a map[string]interface{}.
// It iterates over the fields of the input struct using reflection and
// adds each field name and value to the output map. Only fields of
// struct types are considered for conversion.
// The function returns the resulting map.
// If the input is not a struct, an empty map is returned.
func ConvertToMap(input interface{}) map[string]interface{} {
	outputMap := make(map[string]interface{})

	valueOf := reflect.ValueOf(input)
	typeOf := reflect.TypeOf(input)

	if typeOf.Kind() == reflect.Struct {
		for i := 0; i < valueOf.NumField(); i++ {
			field := typeOf.Field(i)
			fieldName := field.Name
			fieldValue := valueOf.Field(i).Interface()
			outputMap[fieldName] = fieldValue
		}
	}

	return outputMap
}

// coolConvertToMap takes an input and converts it to a map[string]interface{} if the input is a struct.
// If the input is not a struct, an empty map is returned.
func coolConvertToMap(input interface{}) map[string]interface{} {

	if structs.IsStruct(input) {
		return structs.Map(input)
	}
	return make(map[string]interface{})
}
