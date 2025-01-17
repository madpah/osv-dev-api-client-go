# V1VersionMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpe23** | Pointer to **string** | CPE 2.3. | [optional] 
**OsvIdentifier** | Pointer to [**OsvPackage**](OsvPackage.md) |  | [optional] 
**RepoInfo** | Pointer to [**V1VersionRepositoryInformation**](V1VersionRepositoryInformation.md) |  | [optional] 
**Score** | Pointer to **float64** | Score in the interval (0.0, 1.0] with 1.0 being a perfect match. | [optional] 

## Methods

### NewV1VersionMatch

`func NewV1VersionMatch() *V1VersionMatch`

NewV1VersionMatch instantiates a new V1VersionMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VersionMatchWithDefaults

`func NewV1VersionMatchWithDefaults() *V1VersionMatch`

NewV1VersionMatchWithDefaults instantiates a new V1VersionMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpe23

`func (o *V1VersionMatch) GetCpe23() string`

GetCpe23 returns the Cpe23 field if non-nil, zero value otherwise.

### GetCpe23Ok

`func (o *V1VersionMatch) GetCpe23Ok() (*string, bool)`

GetCpe23Ok returns a tuple with the Cpe23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpe23

`func (o *V1VersionMatch) SetCpe23(v string)`

SetCpe23 sets Cpe23 field to given value.

### HasCpe23

`func (o *V1VersionMatch) HasCpe23() bool`

HasCpe23 returns a boolean if a field has been set.

### GetOsvIdentifier

`func (o *V1VersionMatch) GetOsvIdentifier() OsvPackage`

GetOsvIdentifier returns the OsvIdentifier field if non-nil, zero value otherwise.

### GetOsvIdentifierOk

`func (o *V1VersionMatch) GetOsvIdentifierOk() (*OsvPackage, bool)`

GetOsvIdentifierOk returns a tuple with the OsvIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsvIdentifier

`func (o *V1VersionMatch) SetOsvIdentifier(v OsvPackage)`

SetOsvIdentifier sets OsvIdentifier field to given value.

### HasOsvIdentifier

`func (o *V1VersionMatch) HasOsvIdentifier() bool`

HasOsvIdentifier returns a boolean if a field has been set.

### GetRepoInfo

`func (o *V1VersionMatch) GetRepoInfo() V1VersionRepositoryInformation`

GetRepoInfo returns the RepoInfo field if non-nil, zero value otherwise.

### GetRepoInfoOk

`func (o *V1VersionMatch) GetRepoInfoOk() (*V1VersionRepositoryInformation, bool)`

GetRepoInfoOk returns a tuple with the RepoInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoInfo

`func (o *V1VersionMatch) SetRepoInfo(v V1VersionRepositoryInformation)`

SetRepoInfo sets RepoInfo field to given value.

### HasRepoInfo

`func (o *V1VersionMatch) HasRepoInfo() bool`

HasRepoInfo returns a boolean if a field has been set.

### GetScore

`func (o *V1VersionMatch) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *V1VersionMatch) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *V1VersionMatch) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *V1VersionMatch) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


