package reflection

import "reflect"

//Reflection in computing is the ability of a program to examine its own structure,
//particularly through types; it's a form of metaprogramming.

func walk(x interface{}, fn func(input string)) {
	//fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		//fn(field.String())

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
