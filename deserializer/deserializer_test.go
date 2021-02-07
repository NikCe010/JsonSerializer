package deserializer

import (
	"JsonSerializer/deserializer/test_data"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestInvoke(t *testing.T) {
	var data test_data.TestData
	SetValues(&data, 12, "12")
	assert.Equal(t, data.Name, "12")
	assert.Equal(t, data.Num, 12)
}

func TestInvoke_WithComplexStruct(t *testing.T) {
	var data test_data.TestComplex
	SetValues(&data, 12, "12", 12, "12")
	assert.Equal(t, data.Name, "12")
	assert.Equal(t, data.Num, 12)
	assert.Equal(t, data.NameC, "12")
	assert.Equal(t, data.NumC, 12)
}
