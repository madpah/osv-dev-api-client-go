# OsvRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]OsvEvent**](OsvEvent.md) | Required. Version event information. | [optional] 
**Repo** | Pointer to **string** | Required if type is GIT. The publicly accessible URL of the repo that can be directly passed to clone commands. | [optional] 
**Type** | Pointer to [**OsvRangeType**](OsvRangeType.md) |  | [optional] [default to OSVRANGETYPE_UNSPECIFIED]

## Methods

### NewOsvRange

`func NewOsvRange() *OsvRange`

NewOsvRange instantiates a new OsvRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsvRangeWithDefaults

`func NewOsvRangeWithDefaults() *OsvRange`

NewOsvRangeWithDefaults instantiates a new OsvRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *OsvRange) GetEvents() []OsvEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *OsvRange) GetEventsOk() (*[]OsvEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *OsvRange) SetEvents(v []OsvEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *OsvRange) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetRepo

`func (o *OsvRange) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *OsvRange) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *OsvRange) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *OsvRange) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetType

`func (o *OsvRange) GetType() OsvRangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsvRange) GetTypeOk() (*OsvRangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsvRange) SetType(v OsvRangeType)`

SetType sets Type field to given value.

### HasType

`func (o *OsvRange) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


