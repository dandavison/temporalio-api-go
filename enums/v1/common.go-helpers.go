// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package enums

import (
	"fmt"
)

var (
	EncodingType_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Proto3":      1,
		"Json":        2,
	}
	EncodingType_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Proto3",
		2: "Json",
	}
)

// EncodingTypeFromString parses a EncodingType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to EncodingType
func EncodingTypeFromString(s string) (EncodingType, error) {
	if v, ok := EncodingType_value[s]; ok {
		return EncodingType(v), nil
	} else if v, ok := EncodingType_shorthandValue[s]; ok {
		return EncodingType(v), nil
	}
	return EncodingType(0), fmt.Errorf("%s is not a valid EncodingType", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	EncodingType(-1).Shorthand() // will return "", false
func (e EncodingType) Shorthand() (string, bool) {
	if s, ok := EncodingType_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	IndexedValueType_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Text":        1,
		"Keyword":     2,
		"Int":         3,
		"Double":      4,
		"Bool":        5,
		"Datetime":    6,
		"KeywordList": 7,
	}
	IndexedValueType_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Text",
		2: "Keyword",
		3: "Int",
		4: "Double",
		5: "Bool",
		6: "Datetime",
		7: "KeywordList",
	}
)

// IndexedValueTypeFromString parses a IndexedValueType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to IndexedValueType
func IndexedValueTypeFromString(s string) (IndexedValueType, error) {
	if v, ok := IndexedValueType_value[s]; ok {
		return IndexedValueType(v), nil
	} else if v, ok := IndexedValueType_shorthandValue[s]; ok {
		return IndexedValueType(v), nil
	}
	return IndexedValueType(0), fmt.Errorf("%s is not a valid IndexedValueType", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	IndexedValueType(-1).Shorthand() // will return "", false
func (e IndexedValueType) Shorthand() (string, bool) {
	if s, ok := IndexedValueType_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	Severity_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"High":        1,
		"Medium":      2,
		"Low":         3,
	}
	Severity_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "High",
		2: "Medium",
		3: "Low",
	}
)

// SeverityFromString parses a Severity value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to Severity
func SeverityFromString(s string) (Severity, error) {
	if v, ok := Severity_value[s]; ok {
		return Severity(v), nil
	} else if v, ok := Severity_shorthandValue[s]; ok {
		return Severity(v), nil
	}
	return Severity(0), fmt.Errorf("%s is not a valid Severity", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	Severity(-1).Shorthand() // will return "", false
func (e Severity) Shorthand() (string, bool) {
	if s, ok := Severity_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}
