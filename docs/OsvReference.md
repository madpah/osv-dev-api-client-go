# OsvReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**OsvReferenceType**](OsvReferenceType.md) |  | [optional] [default to OSVREFERENCETYPE_NONE]
**Url** | Pointer to **string** | Required. The URL. | [optional] 

## Methods

### NewOsvReference

`func NewOsvReference() *OsvReference`

NewOsvReference instantiates a new OsvReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsvReferenceWithDefaults

`func NewOsvReferenceWithDefaults() *OsvReference`

NewOsvReferenceWithDefaults instantiates a new OsvReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OsvReference) GetType() OsvReferenceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsvReference) GetTypeOk() (*OsvReferenceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsvReference) SetType(v OsvReferenceType)`

SetType sets Type field to given value.

### HasType

`func (o *OsvReference) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *OsvReference) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OsvReference) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OsvReference) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OsvReference) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


