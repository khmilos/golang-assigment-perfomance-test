package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
	"unsafe"
)

// Represents one of multiple external structures
type OneOfExternalStruct struct {
	ID          string `xml:"Id" json:"id"`
	UniqueField string `xml:"UniqueField" json:"uniqueField"`
}

// Represent needed fields
type LocalStruct struct {
	ID string `xml:"Id" json:"id"`
}

func main() {
	value := &OneOfExternalStruct{"struct_id", "unique_field"}
	fmt.Println(AssertBase(value))
	fmt.Println(AssertByTypeSwitch(value))
	fmt.Println(AssertByXMLMarshal(value))
	fmt.Println(AssertByJSONMarshal(value))
	fmt.Println(AssertByReflect(value))
	fmt.Println(AssertByUnsafe(value))
}

// Only for comparison
func AssertBase(i interface{}) string {
	return i.(*OneOfExternalStruct).ID
}

// Only for comparison
func AssertByTypeSwitch(i interface{}) string {
	switch v := i.(type) {
	case OneOfExternalStruct:
		return v.ID
	default:
		return ""
	}
}

func AssertByXMLMarshal(i interface{}) string {
	data, _ := xml.Marshal(i)
	var local LocalStruct
	xml.Unmarshal(data, &local)
	return local.ID
}

func AssertByJSONMarshal(i interface{}) string {
	data, _ := json.Marshal(i)
	var local LocalStruct
	json.Unmarshal(data, &local)
	return local.ID
}

func AssertByReflect(i interface{}) string {
	v := reflect.ValueOf(i)
	return v.Elem().FieldByName("ID").String()
}

var size = unsafe.Sizeof(LocalStruct{})

func AssertByUnsafe(i interface{}) string {
	return (*LocalStruct)(unsafe.Pointer((*[2]uintptr)(unsafe.Pointer(&i))[1])).ID
}
