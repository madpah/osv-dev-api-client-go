# V1VersionQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileHashes** | Pointer to [**[]V1FileHash**](V1FileHash.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the dependency. Can be empty. | [optional] 

## Methods

### NewV1VersionQuery

`func NewV1VersionQuery() *V1VersionQuery`

NewV1VersionQuery instantiates a new V1VersionQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VersionQueryWithDefaults

`func NewV1VersionQueryWithDefaults() *V1VersionQuery`

NewV1VersionQueryWithDefaults instantiates a new V1VersionQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileHashes

`func (o *V1VersionQuery) GetFileHashes() []V1FileHash`

GetFileHashes returns the FileHashes field if non-nil, zero value otherwise.

### GetFileHashesOk

`func (o *V1VersionQuery) GetFileHashesOk() (*[]V1FileHash, bool)`

GetFileHashesOk returns a tuple with the FileHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHashes

`func (o *V1VersionQuery) SetFileHashes(v []V1FileHash)`

SetFileHashes sets FileHashes field to given value.

### HasFileHashes

`func (o *V1VersionQuery) HasFileHashes() bool`

HasFileHashes returns a boolean if a field has been set.

### GetName

`func (o *V1VersionQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1VersionQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1VersionQuery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1VersionQuery) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


