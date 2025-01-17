# OsvEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fixed** | Pointer to **string** | The version/commit that this vulnerability was fixed in. | [optional] 
**Introduced** | Pointer to **string** | The earliest version/commit where this vulnerability was introduced in. | [optional] 
**LastAffected** | Pointer to **string** | The last affected version. | [optional] 
**Limit** | Pointer to **string** | The limit to apply to the range. | [optional] 

## Methods

### NewOsvEvent

`func NewOsvEvent() *OsvEvent`

NewOsvEvent instantiates a new OsvEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsvEventWithDefaults

`func NewOsvEventWithDefaults() *OsvEvent`

NewOsvEventWithDefaults instantiates a new OsvEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFixed

`func (o *OsvEvent) GetFixed() string`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *OsvEvent) GetFixedOk() (*string, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *OsvEvent) SetFixed(v string)`

SetFixed sets Fixed field to given value.

### HasFixed

`func (o *OsvEvent) HasFixed() bool`

HasFixed returns a boolean if a field has been set.

### GetIntroduced

`func (o *OsvEvent) GetIntroduced() string`

GetIntroduced returns the Introduced field if non-nil, zero value otherwise.

### GetIntroducedOk

`func (o *OsvEvent) GetIntroducedOk() (*string, bool)`

GetIntroducedOk returns a tuple with the Introduced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduced

`func (o *OsvEvent) SetIntroduced(v string)`

SetIntroduced sets Introduced field to given value.

### HasIntroduced

`func (o *OsvEvent) HasIntroduced() bool`

HasIntroduced returns a boolean if a field has been set.

### GetLastAffected

`func (o *OsvEvent) GetLastAffected() string`

GetLastAffected returns the LastAffected field if non-nil, zero value otherwise.

### GetLastAffectedOk

`func (o *OsvEvent) GetLastAffectedOk() (*string, bool)`

GetLastAffectedOk returns a tuple with the LastAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAffected

`func (o *OsvEvent) SetLastAffected(v string)`

SetLastAffected sets LastAffected field to given value.

### HasLastAffected

`func (o *OsvEvent) HasLastAffected() bool`

HasLastAffected returns a boolean if a field has been set.

### GetLimit

`func (o *OsvEvent) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *OsvEvent) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *OsvEvent) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *OsvEvent) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


