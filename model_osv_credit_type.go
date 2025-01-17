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

// OsvCreditType the model 'OsvCreditType'
type OsvCreditType string

// List of osvCreditType
const (
	OSVCREDITTYPE_UNSPECIFIED OsvCreditType = "UNSPECIFIED"
	OSVCREDITTYPE_OTHER OsvCreditType = "OTHER"
	OSVCREDITTYPE_FINDER OsvCreditType = "FINDER"
	OSVCREDITTYPE_REPORTER OsvCreditType = "REPORTER"
	OSVCREDITTYPE_ANALYST OsvCreditType = "ANALYST"
	OSVCREDITTYPE_COORDINATOR OsvCreditType = "COORDINATOR"
	OSVCREDITTYPE_REMEDIATION_DEVELOPER OsvCreditType = "REMEDIATION_DEVELOPER"
	OSVCREDITTYPE_REMEDIATION_REVIEWER OsvCreditType = "REMEDIATION_REVIEWER"
	OSVCREDITTYPE_REMEDIATION_VERIFIER OsvCreditType = "REMEDIATION_VERIFIER"
	OSVCREDITTYPE_TOOL OsvCreditType = "TOOL"
	OSVCREDITTYPE_SPONSOR OsvCreditType = "SPONSOR"
)

// All allowed values of OsvCreditType enum
var AllowedOsvCreditTypeEnumValues = []OsvCreditType{
	"UNSPECIFIED",
	"OTHER",
	"FINDER",
	"REPORTER",
	"ANALYST",
	"COORDINATOR",
	"REMEDIATION_DEVELOPER",
	"REMEDIATION_REVIEWER",
	"REMEDIATION_VERIFIER",
	"TOOL",
	"SPONSOR",
}

func (v *OsvCreditType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OsvCreditType(value)
	for _, existing := range AllowedOsvCreditTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OsvCreditType", value)
}

// NewOsvCreditTypeFromValue returns a pointer to a valid OsvCreditType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOsvCreditTypeFromValue(v string) (*OsvCreditType, error) {
	ev := OsvCreditType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OsvCreditType: valid values are %v", v, AllowedOsvCreditTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OsvCreditType) IsValid() bool {
	for _, existing := range AllowedOsvCreditTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to osvCreditType value
func (v OsvCreditType) Ptr() *OsvCreditType {
	return &v
}

type NullableOsvCreditType struct {
	value *OsvCreditType
	isSet bool
}

func (v NullableOsvCreditType) Get() *OsvCreditType {
	return v.value
}

func (v *NullableOsvCreditType) Set(val *OsvCreditType) {
	v.value = val
	v.isSet = true
}

func (v NullableOsvCreditType) IsSet() bool {
	return v.isSet
}

func (v *NullableOsvCreditType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsvCreditType(val *OsvCreditType) *NullableOsvCreditType {
	return &NullableOsvCreditType{value: val, isSet: true}
}

func (v NullableOsvCreditType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsvCreditType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

