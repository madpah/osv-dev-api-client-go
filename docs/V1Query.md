# V1Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to **string** | The commit hash to query for. If specified, &#x60;version&#x60; should not be set. | [optional] 
**Package** | Pointer to [**OsvPackage**](OsvPackage.md) |  | [optional] 
**PageToken** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** | The version string to query for. A fuzzy match is done against upstream versions. If specified, &#x60;commit&#x60; should not be set. | [optional] 

## Methods

### NewV1Query

`func NewV1Query() *V1Query`

NewV1Query instantiates a new V1Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1QueryWithDefaults

`func NewV1QueryWithDefaults() *V1Query`

NewV1QueryWithDefaults instantiates a new V1Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *V1Query) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *V1Query) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *V1Query) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *V1Query) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetPackage

`func (o *V1Query) GetPackage() OsvPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *V1Query) GetPackageOk() (*OsvPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *V1Query) SetPackage(v OsvPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *V1Query) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetPageToken

`func (o *V1Query) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V1Query) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V1Query) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V1Query) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetVersion

`func (o *V1Query) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1Query) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1Query) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1Query) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


