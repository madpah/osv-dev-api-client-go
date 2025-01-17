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

// checks if the V1Query type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1Query{}

// V1Query Query format.
type V1Query struct {
	// The commit hash to query for. If specified, `version` should not be set.
	Commit *string `json:"commit,omitempty"`
	Package *OsvPackage `json:"package,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
	// The version string to query for. A fuzzy match is done against upstream versions. If specified, `commit` should not be set.
	Version *string `json:"version,omitempty"`
}

// NewV1Query instantiates a new V1Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Query() *V1Query {
	this := V1Query{}
	return &this
}

// NewV1QueryWithDefaults instantiates a new V1Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1QueryWithDefaults() *V1Query {
	this := V1Query{}
	return &this
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *V1Query) GetCommit() string {
	if o == nil || IsNil(o.Commit) {
		var ret string
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Query) GetCommitOk() (*string, bool) {
	if o == nil || IsNil(o.Commit) {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *V1Query) HasCommit() bool {
	if o != nil && !IsNil(o.Commit) {
		return true
	}

	return false
}

// SetCommit gets a reference to the given string and assigns it to the Commit field.
func (o *V1Query) SetCommit(v string) {
	o.Commit = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *V1Query) GetPackage() OsvPackage {
	if o == nil || IsNil(o.Package) {
		var ret OsvPackage
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Query) GetPackageOk() (*OsvPackage, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *V1Query) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given OsvPackage and assigns it to the Package field.
func (o *V1Query) SetPackage(v OsvPackage) {
	o.Package = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *V1Query) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Query) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *V1Query) HasPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *V1Query) SetPageToken(v string) {
	o.PageToken = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *V1Query) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Query) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *V1Query) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *V1Query) SetVersion(v string) {
	o.Version = &v
}

func (o V1Query) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1Query) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commit) {
		toSerialize["commit"] = o.Commit
	}
	if !IsNil(o.Package) {
		toSerialize["package"] = o.Package
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableV1Query struct {
	value *V1Query
	isSet bool
}

func (v NullableV1Query) Get() *V1Query {
	return v.value
}

func (v *NullableV1Query) Set(val *V1Query) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Query) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Query) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Query(val *V1Query) *NullableV1Query {
	return &NullableV1Query{value: val, isSet: true}
}

func (v NullableV1Query) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Query) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


