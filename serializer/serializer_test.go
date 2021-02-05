package serializer

import (
	"JsonSerializer/serializer/test_data"
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSerialize_WithBaseStruct_ShouldReturnValidJson(t *testing.T) {
	jsn := Serialize(test_data.NewBaseStruct())
	assert.Equal(t, json.Valid([]byte(jsn)), true)
}

func TestSerialize_WithComplexStruct_ShouldReturnValidJson(t *testing.T) {
	jsn := Serialize(test_data.NewTestStruct())
	assert.Equal(t, json.Valid([]byte(jsn)), true)
}

func TestSerialize_WithStructWithManyEmbeddedStructs_ShouldReturnValidJson(t *testing.T) {
	jsn := Serialize(test_data.NewAnotherTestStruct())
	assert.Equal(t, json.Valid([]byte(jsn)), true)
}

func BenchmarkSerialize_BaseStruct(b *testing.B) {
	var testSet []test_data.BaseStruct
	for i := 0; i < b.N; i++ {
		testSet = append(testSet, test_data.NewBaseStruct())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Serialize(test_data.NewAnotherTestStruct())
	}
}

func BenchmarkSerialize_ComplexStruct(b *testing.B) {
	var testSet []test_data.TestStruct
	for i := 0; i < b.N; i++ {
		testSet = append(testSet, test_data.NewTestStruct())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Serialize(test_data.NewAnotherTestStruct())
	}
}

func BenchmarkSerialize_StructWithManyEmbeddedStructs(b *testing.B) {
	var testSet []test_data.AnotherTestStruct
	for i := 0; i < b.N; i++ {
		testSet = append(testSet, test_data.NewAnotherTestStruct())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Serialize(test_data.NewAnotherTestStruct())
	}
}
