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

// checks if the V1BatchQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BatchQuery{}

// V1BatchQuery Batch query format.
type V1BatchQuery struct {
	// The queries that form this batch query.
	Queries []V1Query `json:"queries,omitempty"`
}

// NewV1BatchQuery instantiates a new V1BatchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BatchQuery() *V1BatchQuery {
	this := V1BatchQuery{}
	return &this
}

// NewV1BatchQueryWithDefaults instantiates a new V1BatchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BatchQueryWithDefaults() *V1BatchQuery {
	this := V1BatchQuery{}
	return &this
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *V1BatchQuery) GetQueries() []V1Query {
	if o == nil || IsNil(o.Queries) {
		var ret []V1Query
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BatchQuery) GetQueriesOk() ([]V1Query, bool) {
	if o == nil || IsNil(o.Queries) {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *V1BatchQuery) HasQueries() bool {
	if o != nil && !IsNil(o.Queries) {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []V1Query and assigns it to the Queries field.
func (o *V1BatchQuery) SetQueries(v []V1Query) {
	o.Queries = v
}

func (o V1BatchQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BatchQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Queries) {
		toSerialize["queries"] = o.Queries
	}
	return toSerialize, nil
}

type NullableV1BatchQuery struct {
	value *V1BatchQuery
	isSet bool
}

func (v NullableV1BatchQuery) Get() *V1BatchQuery {
	return v.value
}

func (v *NullableV1BatchQuery) Set(val *V1BatchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BatchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BatchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BatchQuery(val *V1BatchQuery) *NullableV1BatchQuery {
	return &NullableV1BatchQuery{value: val, isSet: true}
}

func (v NullableV1BatchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BatchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


