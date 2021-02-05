package test_data

import "time"

type StructWithEmbedded struct {
	InlineS string
	InlineG string
	BaseStruct
}

type BaseStruct struct {
	InlineInlineS string
	InlineInlineG string
}

type TimeTest struct {
	InlineInlineS time.Time
	InlineInlineG time.Time
}

type TestStruct struct {
	TimeTest
	Name    string
	Surname string
	StructWithEmbedded
	BaseStruct
}

type AnotherTestStruct struct {
	TimeTest
	Name    string
	Surname string
	StructWithEmbedded
	BaseStruct
	TestStruct
	Age       int
	IsMarried bool
	FloatTest float32
}

func NewBaseStruct() BaseStruct {
	return BaseStruct{"456", "svsdvd"}
}

func NewTestStruct() TestStruct {
	return TestStruct{TimeTest{time.Now(), time.Now()},
		"23",
		"skidoo",
		StructWithEmbedded{"123", "123345",
			BaseStruct{"456", "svsdvd"}},
		BaseStruct{"cvb", "67"}}
}

func NewAnotherTestStruct() AnotherTestStruct {
	return AnotherTestStruct{TimeTest{time.Now(), time.Now()},
		"23",
		"skidoo",
		StructWithEmbedded{"123", "123345",
			BaseStruct{"456", "svsdvd"}},
		BaseStruct{"cvb", "67"},
		TestStruct{TimeTest{time.Now(), time.Now()},
			"23",
			"skidoo",
			StructWithEmbedded{"123", "123345",
				BaseStruct{"456", "svsdvd"}},
			BaseStruct{"cvb", "67"}},
		12,
		true,
		12.756}
}
