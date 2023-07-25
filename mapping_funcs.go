package model_mapper

import (
	"errors"
	"fmt"
	"reflect"
)

func mapWithNullFields(to interface{}, from interface{}) error {
	toRefl := reflect.ValueOf(to)
	fromRefl := reflect.ValueOf(from)
	if toRefl.Kind() != reflect.Ptr || fromRefl.Kind() != reflect.Ptr {
		return fmt.Errorf("you have to input only reference type values")
	}
	toRefl = reflect.ValueOf(to).Elem()
	fromRefl = reflect.ValueOf(from).Elem()
	for i := 0; i < toRefl.NumField(); i++ {
		fieldName := toRefl.Type().Field(i).Name
		fromStructField := fromRefl.FieldByName(fieldName)
		var fromValue interface{}
		if !fromStructField.IsValid() {
			fromValue = ""
		} else {
			fromValue = fromStructField.Interface()
		}
		toValue := toRefl.Field(i).Interface()
		if reflect.TypeOf(fromValue) != reflect.TypeOf(toValue) {
			fromValue = ""
		}
		if err := setField(toRefl, fieldName, fromValue); err != nil {
			return err
		}
	}
	return nil
}

func setField(structValue reflect.Value, name string, value interface{}) error {
	//structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)
	if !structFieldValue.IsValid() {
		return fmt.Errorf("couldn't find %s field in obj", name)
	}
	if !structFieldValue.CanSet() {
		return fmt.Errorf("couldn't set %s field value", name)
	}
	structFieldType := structFieldValue.Type()
	if reflect.ValueOf(value).Type() != structFieldType {
		return errors.New("provided value type doesn't match input obj field type")
	}
	structFieldValue.Set(reflect.ValueOf(value))
	return nil
}
