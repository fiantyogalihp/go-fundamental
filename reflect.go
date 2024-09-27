package main

import (
	"fmt"
	"reflect"
)

type PersonV2 struct {
	Name string `required:"true" min:"4"`
}

type People struct {
	Name string // * tidak ada tag
	Desc string
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data) // ? cek dlu type data nya
	// ? lalu iterasi jumlah datanya
	for i := 0; i < t.NumField(); i++ {
		// ? ambil semua field nya
		field := t.Field(i)
		// ? validasi value pada tag pada field nya
		if field.Tag.Get("required") == "true" {
			// ? data dibawah tidak boleh string kosong
			return reflect.ValueOf(data).Field(i).Interface() != ""
			// ? cara dibawah adalah versi detail dari code diatas
			// value := reflect.ValueOf(data).Field(i).Interface()
			// if value == "" {
			//   return false
			// }
		}
	}
	// ? jika proses sudah selesai return true
	return true
}

func ReflectModule() {
	person := PersonV2{"Slamet"}

	personType := reflect.TypeOf(person)

	fmt.Println(personType.NumField())
	fmt.Println(personType.Field(0).Name)
	fmt.Println(personType.Field(0).Tag.Get("required"))
	fmt.Println(personType.Field(0).Tag.Get("min"))

	// person.Name = ""
	// ! jika isi struct berisi string kosong, maka func `IsValid` akan tetap true, karena tidak ada tag dalam field nya
	fmt.Println("Person:", IsValid(person))

	people := People{
		Desc: "kopleng",
		Name: "Slamet",
	}
	fmt.Println("`People` akan tetap true, karena tidak ada tag dalam field nya \nPeople:", IsValid(people))
}
