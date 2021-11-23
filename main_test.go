package main

import (
	"testing"
)

func BenchmarkAssertBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AssertBase(&OneOfExternalStruct{"struct_id", "unique_field"})
	}
}

func BenchmarkAssertByTypeSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AssertByTypeSwitch(&OneOfExternalStruct{"struct_id", "unique_field"})
	}
}

func BenchmarkAssertByXMLMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AssertByXMLMarshal(&OneOfExternalStruct{"struct_id", "unique_field"})
	}
}

func BenchmarkAssertByJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AssertByJSONMarshal(&OneOfExternalStruct{"struct_id", "unique_field"})
	}
}

func BenchmarkAssertByReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AssertByReflect(&OneOfExternalStruct{"struct_id", "unique_field"})
	}
}

func BenchmarkAssertByUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AssertByUnsafe(&OneOfExternalStruct{"struct_id", "unique_field"})
	}
}
