package db

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func MapToInterface(object interface{}, result interface{}) error {

	jsonStruct, err := json.Marshal(object)
	if err != nil {
		return err
	}

	json.Unmarshal(jsonStruct, result)

	return nil
}

func InterfaceToMap(object interface{}) (*map[string]interface{}, error) {
	if reflect.ValueOf(object).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("object should be of pointer type")
	}

	result := &map[string]interface{}{}
	rValue := reflect.ValueOf(object).Elem()
	rKind := rValue.Kind()

	switch rKind {

	case reflect.Struct:
		typeOfObject := rValue.Type()

		for i := 0; i < rValue.NumField(); i++ {
			f := rValue.Field(i)
			key := strings.ToLower(typeOfObject.Field(i).Name)
			value := f.Interface()
			(*result)[key] = value
		}
	case reflect.Map:

		if _, ok := object.(*map[string]interface{}); ok {
			result = object.(*map[string]interface{})
		} else {
			return nil, fmt.Errorf("invalid map type, should be *map[string]interface{}")
		}
	default:

		return nil, fmt.Errorf("invalid object type, it should be struct pointer or *map[string]interface{}")
	}

	return result, nil
}
