/*
OSV

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osvdev

import (
	"encoding/json"
)

// checks if the OsvEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsvEvent{}

// OsvEvent Version events.
type OsvEvent struct {
	// The version/commit that this vulnerability was fixed in.
	Fixed *string `json:"fixed,omitempty"`
	// The earliest version/commit where this vulnerability was introduced in.
	Introduced *string `json:"introduced,omitempty"`
	// The last affected version.
	LastAffected *string `json:"lastAffected,omitempty"`
	// The limit to apply to the range.
	Limit *string `json:"limit,omitempty"`
}

// NewOsvEvent instantiates a new OsvEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsvEvent() *OsvEvent {
	this := OsvEvent{}
	return &this
}

// NewOsvEventWithDefaults instantiates a new OsvEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsvEventWithDefaults() *OsvEvent {
	this := OsvEvent{}
	return &this
}

// GetFixed returns the Fixed field value if set, zero value otherwise.
func (o *OsvEvent) GetFixed() string {
	if o == nil || IsNil(o.Fixed) {
		var ret string
		return ret
	}
	return *o.Fixed
}

// GetFixedOk returns a tuple with the Fixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsvEvent) GetFixedOk() (*string, bool) {
	if o == nil || IsNil(o.Fixed) {
		return nil, false
	}
	return o.Fixed, true
}

// HasFixed returns a boolean if a field has been set.
func (o *OsvEvent) HasFixed() bool {
	if o != nil && !IsNil(o.Fixed) {
		return true
	}

	return false
}

// SetFixed gets a reference to the given string and assigns it to the Fixed field.
func (o *OsvEvent) SetFixed(v string) {
	o.Fixed = &v
}

// GetIntroduced returns the Introduced field value if set, zero value otherwise.
func (o *OsvEvent) GetIntroduced() string {
	if o == nil || IsNil(o.Introduced) {
		var ret string
		return ret
	}
	return *o.Introduced
}

// GetIntroducedOk returns a tuple with the Introduced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsvEvent) GetIntroducedOk() (*string, bool) {
	if o == nil || IsNil(o.Introduced) {
		return nil, false
	}
	return o.Introduced, true
}

// HasIntroduced returns a boolean if a field has been set.
func (o *OsvEvent) HasIntroduced() bool {
	if o != nil && !IsNil(o.Introduced) {
		return true
	}

	return false
}

// SetIntroduced gets a reference to the given string and assigns it to the Introduced field.
func (o *OsvEvent) SetIntroduced(v string) {
	o.Introduced = &v
}

// GetLastAffected returns the LastAffected field value if set, zero value otherwise.
func (o *OsvEvent) GetLastAffected() string {
	if o == nil || IsNil(o.LastAffected) {
		var ret string
		return ret
	}
	return *o.LastAffected
}

// GetLastAffectedOk returns a tuple with the LastAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsvEvent) GetLastAffectedOk() (*string, bool) {
	if o == nil || IsNil(o.LastAffected) {
		return nil, false
	}
	return o.LastAffected, true
}

// HasLastAffected returns a boolean if a field has been set.
func (o *OsvEvent) HasLastAffected() bool {
	if o != nil && !IsNil(o.LastAffected) {
		return true
	}

	return false
}

// SetLastAffected gets a reference to the given string and assigns it to the LastAffected field.
func (o *OsvEvent) SetLastAffected(v string) {
	o.LastAffected = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *OsvEvent) GetLimit() string {
	if o == nil || IsNil(o.Limit) {
		var ret string
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsvEvent) GetLimitOk() (*string, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *OsvEvent) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given string and assigns it to the Limit field.
func (o *OsvEvent) SetLimit(v string) {
	o.Limit = &v
}

func (o OsvEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsvEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fixed) {
		toSerialize["fixed"] = o.Fixed
	}
	if !IsNil(o.Introduced) {
		toSerialize["introduced"] = o.Introduced
	}
	if !IsNil(o.LastAffected) {
		toSerialize["lastAffected"] = o.LastAffected
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

type NullableOsvEvent struct {
	value *OsvEvent
	isSet bool
}

func (v NullableOsvEvent) Get() *OsvEvent {
	return v.value
}

func (v *NullableOsvEvent) Set(val *OsvEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableOsvEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableOsvEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsvEvent(val *OsvEvent) *NullableOsvEvent {
	return &NullableOsvEvent{value: val, isSet: true}
}

func (v NullableOsvEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsvEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


