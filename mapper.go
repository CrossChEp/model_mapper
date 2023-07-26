package model_mapper

import (
	"fmt"
	"reflect"
)

func Map(to interface{}, from interface{}, skipNulls bool) error {
	if !skipNulls {
		if err := mapWithNullFields(to, from); err != nil {
			return err
		}
		return nil
	}
	return mapWithoutNullFields(to, from)
}

func ConvertToJson(structure interface{}) (map[string]interface{}, error) {
	return convertToJson(structure)
}

func SetValueToObjectField(obj interface{}, fName string, value interface{}) error {
	reflElem := reflect.ValueOf(obj)
	if reflElem.Kind() != reflect.Ptr {
		return fmt.Errorf("you have to input only reference type values")
	}
	return setField(reflElem.Elem(), fName, value)
}
