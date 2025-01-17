# V1FileHash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | Pointer to **string** | The file path inside the repository, relative to the repository root. | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**HashType** | Pointer to **string** |  | [optional] 

## Methods

### NewV1FileHash

`func NewV1FileHash() *V1FileHash`

NewV1FileHash instantiates a new V1FileHash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FileHashWithDefaults

`func NewV1FileHashWithDefaults() *V1FileHash`

NewV1FileHashWithDefaults instantiates a new V1FileHash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePath

`func (o *V1FileHash) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *V1FileHash) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *V1FileHash) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *V1FileHash) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetHash

`func (o *V1FileHash) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *V1FileHash) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *V1FileHash) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *V1FileHash) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetHashType

`func (o *V1FileHash) GetHashType() string`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *V1FileHash) GetHashTypeOk() (*string, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *V1FileHash) SetHashType(v string)`

SetHashType sets HashType field to given value.

### HasHashType

`func (o *V1FileHash) HasHashType() bool`

HasHashType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


