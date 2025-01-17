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

// checks if the V1VersionQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1VersionQuery{}

// V1VersionQuery The version query.
type V1VersionQuery struct {
	FileHashes []V1FileHash `json:"fileHashes,omitempty"`
	// The name of the dependency. Can be empty.
	Name *string `json:"name,omitempty"`
}

// NewV1VersionQuery instantiates a new V1VersionQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1VersionQuery() *V1VersionQuery {
	this := V1VersionQuery{}
	return &this
}

// NewV1VersionQueryWithDefaults instantiates a new V1VersionQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1VersionQueryWithDefaults() *V1VersionQuery {
	this := V1VersionQuery{}
	return &this
}

// GetFileHashes returns the FileHashes field value if set, zero value otherwise.
func (o *V1VersionQuery) GetFileHashes() []V1FileHash {
	if o == nil || IsNil(o.FileHashes) {
		var ret []V1FileHash
		return ret
	}
	return o.FileHashes
}

// GetFileHashesOk returns a tuple with the FileHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VersionQuery) GetFileHashesOk() ([]V1FileHash, bool) {
	if o == nil || IsNil(o.FileHashes) {
		return nil, false
	}
	return o.FileHashes, true
}

// HasFileHashes returns a boolean if a field has been set.
func (o *V1VersionQuery) HasFileHashes() bool {
	if o != nil && !IsNil(o.FileHashes) {
		return true
	}

	return false
}

// SetFileHashes gets a reference to the given []V1FileHash and assigns it to the FileHashes field.
func (o *V1VersionQuery) SetFileHashes(v []V1FileHash) {
	o.FileHashes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1VersionQuery) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VersionQuery) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1VersionQuery) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1VersionQuery) SetName(v string) {
	o.Name = &v
}

func (o V1VersionQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1VersionQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileHashes) {
		toSerialize["fileHashes"] = o.FileHashes
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV1VersionQuery struct {
	value *V1VersionQuery
	isSet bool
}

func (v NullableV1VersionQuery) Get() *V1VersionQuery {
	return v.value
}

func (v *NullableV1VersionQuery) Set(val *V1VersionQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableV1VersionQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableV1VersionQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1VersionQuery(val *V1VersionQuery) *NullableV1VersionQuery {
	return &NullableV1VersionQuery{value: val, isSet: true}
}

func (v NullableV1VersionQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1VersionQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


