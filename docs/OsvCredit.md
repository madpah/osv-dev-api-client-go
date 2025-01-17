# OsvCredit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to **[]string** | Contact methods (URLs). | [optional] 
**Name** | Pointer to **string** | The name to give credit to. | [optional] 
**Type** | Pointer to [**OsvCreditType**](OsvCreditType.md) |  | [optional] [default to OSVCREDITTYPE_UNSPECIFIED]

## Methods

### NewOsvCredit

`func NewOsvCredit() *OsvCredit`

NewOsvCredit instantiates a new OsvCredit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsvCreditWithDefaults

`func NewOsvCreditWithDefaults() *OsvCredit`

NewOsvCreditWithDefaults instantiates a new OsvCredit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *OsvCredit) GetContact() []string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *OsvCredit) GetContactOk() (*[]string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *OsvCredit) SetContact(v []string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *OsvCredit) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetName

`func (o *OsvCredit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsvCredit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsvCredit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsvCredit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *OsvCredit) GetType() OsvCreditType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsvCredit) GetTypeOk() (*OsvCreditType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsvCredit) SetType(v OsvCreditType)`

SetType sets Type field to given value.

### HasType

`func (o *OsvCredit) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


