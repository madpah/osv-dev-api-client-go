/*
OSV

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osvdev

import (
	"encoding/json"
	"fmt"
)

// OsvReferenceType the model 'OsvReferenceType'
type OsvReferenceType string

// List of osvReferenceType
const (
	OSVREFERENCETYPE_NONE OsvReferenceType = "NONE"
	OSVREFERENCETYPE_WEB OsvReferenceType = "WEB"
	OSVREFERENCETYPE_ADVISORY OsvReferenceType = "ADVISORY"
	OSVREFERENCETYPE_REPORT OsvReferenceType = "REPORT"
	OSVREFERENCETYPE_FIX OsvReferenceType = "FIX"
	OSVREFERENCETYPE_PACKAGE OsvReferenceType = "PACKAGE"
	OSVREFERENCETYPE_ARTICLE OsvReferenceType = "ARTICLE"
	OSVREFERENCETYPE_EVIDENCE OsvReferenceType = "EVIDENCE"
)

// All allowed values of OsvReferenceType enum
var AllowedOsvReferenceTypeEnumValues = []OsvReferenceType{
	"NONE",
	"WEB",
	"ADVISORY",
	"REPORT",
	"FIX",
	"PACKAGE",
	"ARTICLE",
	"EVIDENCE",
}

func (v *OsvReferenceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OsvReferenceType(value)
	for _, existing := range AllowedOsvReferenceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OsvReferenceType", value)
}

// NewOsvReferenceTypeFromValue returns a pointer to a valid OsvReferenceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOsvReferenceTypeFromValue(v string) (*OsvReferenceType, error) {
	ev := OsvReferenceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OsvReferenceType: valid values are %v", v, AllowedOsvReferenceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OsvReferenceType) IsValid() bool {
	for _, existing := range AllowedOsvReferenceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to osvReferenceType value
func (v OsvReferenceType) Ptr() *OsvReferenceType {
	return &v
}

type NullableOsvReferenceType struct {
	value *OsvReferenceType
	isSet bool
}

func (v NullableOsvReferenceType) Get() *OsvReferenceType {
	return v.value
}

func (v *NullableOsvReferenceType) Set(val *OsvReferenceType) {
	v.value = val
	v.isSet = true
}

func (v NullableOsvReferenceType) IsSet() bool {
	return v.isSet
}

func (v *NullableOsvReferenceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsvReferenceType(val *OsvReferenceType) *NullableOsvReferenceType {
	return &NullableOsvReferenceType{value: val, isSet: true}
}

func (v NullableOsvReferenceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsvReferenceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

