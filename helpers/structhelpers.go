package helpers

import "reflect"

// CopyIfNotNil copies struct src fields that have a match on dest, if their pointer values are not nil
func CopyIfNotNil(src interface{}, dest interface{}) {
	destReflect := reflect.ValueOf(dest).Elem()
	srcReflect := reflect.ValueOf(src).Elem()
	for i := 0; i < srcReflect.NumField(); i++ {
		destField := destReflect.Field(i)
		srcFieldValue := reflect.Value(srcReflect.Field(i))
		if !srcFieldValue.IsNil() {
			destField.Set(srcFieldValue)
		}
	}
}
