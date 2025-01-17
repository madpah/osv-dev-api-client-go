# OsvAffected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseSpecific** | Pointer to **map[string]interface{}** | Optional. JSON object holding additional information about the vulnerability as defined by the database for which the record applies. | [optional] 
**EcosystemSpecific** | Pointer to **map[string]interface{}** | Optional. JSON object holding additional information about the vulnerability as defined by the ecosystem for which the record applies. | [optional] 
**Package** | Pointer to [**OsvPackage**](OsvPackage.md) |  | [optional] 
**Ranges** | Pointer to [**[]OsvRange**](OsvRange.md) | Required. Range information. | [optional] 
**Severity** | Pointer to [**[]OsvSeverity**](OsvSeverity.md) | Optional. Severity of the vulnerability for this package. | [optional] 
**Versions** | Pointer to **[]string** | Optional. List of affected versions. | [optional] 

## Methods

### NewOsvAffected

`func NewOsvAffected() *OsvAffected`

NewOsvAffected instantiates a new OsvAffected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsvAffectedWithDefaults

`func NewOsvAffectedWithDefaults() *OsvAffected`

NewOsvAffectedWithDefaults instantiates a new OsvAffected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseSpecific

`func (o *OsvAffected) GetDatabaseSpecific() map[string]interface{}`

GetDatabaseSpecific returns the DatabaseSpecific field if non-nil, zero value otherwise.

### GetDatabaseSpecificOk

`func (o *OsvAffected) GetDatabaseSpecificOk() (*map[string]interface{}, bool)`

GetDatabaseSpecificOk returns a tuple with the DatabaseSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSpecific

`func (o *OsvAffected) SetDatabaseSpecific(v map[string]interface{})`

SetDatabaseSpecific sets DatabaseSpecific field to given value.

### HasDatabaseSpecific

`func (o *OsvAffected) HasDatabaseSpecific() bool`

HasDatabaseSpecific returns a boolean if a field has been set.

### GetEcosystemSpecific

`func (o *OsvAffected) GetEcosystemSpecific() map[string]interface{}`

GetEcosystemSpecific returns the EcosystemSpecific field if non-nil, zero value otherwise.

### GetEcosystemSpecificOk

`func (o *OsvAffected) GetEcosystemSpecificOk() (*map[string]interface{}, bool)`

GetEcosystemSpecificOk returns a tuple with the EcosystemSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcosystemSpecific

`func (o *OsvAffected) SetEcosystemSpecific(v map[string]interface{})`

SetEcosystemSpecific sets EcosystemSpecific field to given value.

### HasEcosystemSpecific

`func (o *OsvAffected) HasEcosystemSpecific() bool`

HasEcosystemSpecific returns a boolean if a field has been set.

### GetPackage

`func (o *OsvAffected) GetPackage() OsvPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *OsvAffected) GetPackageOk() (*OsvPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *OsvAffected) SetPackage(v OsvPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *OsvAffected) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetRanges

`func (o *OsvAffected) GetRanges() []OsvRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *OsvAffected) GetRangesOk() (*[]OsvRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *OsvAffected) SetRanges(v []OsvRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *OsvAffected) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetSeverity

`func (o *OsvAffected) GetSeverity() []OsvSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *OsvAffected) GetSeverityOk() (*[]OsvSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *OsvAffected) SetSeverity(v []OsvSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *OsvAffected) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetVersions

`func (o *OsvAffected) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *OsvAffected) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *OsvAffected) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *OsvAffected) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


