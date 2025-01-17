# V1VersionRepositoryInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Source address of the repository. | [optional] 
**Commit** | Pointer to **string** | Commit hash. | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**V1VersionRepositoryInformationRepoType**](V1VersionRepositoryInformationRepoType.md) |  | [optional] [default to V1VERSIONREPOSITORYINFORMATIONREPOTYPE_UNSPECIFIED]
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewV1VersionRepositoryInformation

`func NewV1VersionRepositoryInformation() *V1VersionRepositoryInformation`

NewV1VersionRepositoryInformation instantiates a new V1VersionRepositoryInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VersionRepositoryInformationWithDefaults

`func NewV1VersionRepositoryInformationWithDefaults() *V1VersionRepositoryInformation`

NewV1VersionRepositoryInformationWithDefaults instantiates a new V1VersionRepositoryInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *V1VersionRepositoryInformation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V1VersionRepositoryInformation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V1VersionRepositoryInformation) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V1VersionRepositoryInformation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCommit

`func (o *V1VersionRepositoryInformation) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *V1VersionRepositoryInformation) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *V1VersionRepositoryInformation) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *V1VersionRepositoryInformation) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetTag

`func (o *V1VersionRepositoryInformation) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V1VersionRepositoryInformation) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V1VersionRepositoryInformation) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V1VersionRepositoryInformation) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *V1VersionRepositoryInformation) GetType() V1VersionRepositoryInformationRepoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1VersionRepositoryInformation) GetTypeOk() (*V1VersionRepositoryInformationRepoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1VersionRepositoryInformation) SetType(v V1VersionRepositoryInformationRepoType)`

SetType sets Type field to given value.

### HasType

`func (o *V1VersionRepositoryInformation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *V1VersionRepositoryInformation) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1VersionRepositoryInformation) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1VersionRepositoryInformation) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1VersionRepositoryInformation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


