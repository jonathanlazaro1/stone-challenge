package helpers

import "reflect"

// CopyIfNotNil copies struct src fields that have a name match on dest, if their src pointer values are not nil
func CopyIfNotNil(src interface{}, dest interface{}) {
	destReflect := reflect.ValueOf(dest).Elem()
	srcReflect := reflect.ValueOf(src).Elem()
	for i := 0; i < srcReflect.NumField(); i++ {
		destField := destReflect.FieldByName(srcReflect.Type().Field(i).Name)
		if !destField.IsValid() {
			continue
		}
		srcFieldValue := reflect.Value(srcReflect.Field(i))
		if srcFieldValue.IsNil() {
			continue
		}
		if srcFieldValue.Kind() == reflect.Ptr {
			destField.Set(srcFieldValue.Elem())
		} else {
			destField.Set(srcFieldValue)
		}
	}
}
