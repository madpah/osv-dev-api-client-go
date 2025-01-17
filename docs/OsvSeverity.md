# OsvSeverity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **string** | The quantitative score. | [optional] 
**Type** | Pointer to [**OsvSeverityType**](OsvSeverityType.md) |  | [optional] [default to OSVSEVERITYTYPE_UNSPECIFIED]

## Methods

### NewOsvSeverity

`func NewOsvSeverity() *OsvSeverity`

NewOsvSeverity instantiates a new OsvSeverity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsvSeverityWithDefaults

`func NewOsvSeverityWithDefaults() *OsvSeverity`

NewOsvSeverityWithDefaults instantiates a new OsvSeverity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *OsvSeverity) GetScore() string`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *OsvSeverity) GetScoreOk() (*string, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *OsvSeverity) SetScore(v string)`

SetScore sets Score field to given value.

### HasScore

`func (o *OsvSeverity) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetType

`func (o *OsvSeverity) GetType() OsvSeverityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsvSeverity) GetTypeOk() (*OsvSeverityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsvSeverity) SetType(v OsvSeverityType)`

SetType sets Type field to given value.

### HasType

`func (o *OsvSeverity) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


