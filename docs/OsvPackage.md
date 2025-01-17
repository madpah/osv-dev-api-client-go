# OsvPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ecosystem** | Pointer to **string** | Required. The ecosystem for this package.  For the complete list of valid ecosystem names, see &lt;https://ossf.github.io/osv-schema/#affectedpackage-field&gt;. | [optional] 
**Name** | Pointer to **string** | Required. Name of the package. Should match the name used in the package ecosystem (e.g. the npm package name). For C/C++ projects integrated in OSS-Fuzz, this is the name used for the integration. | [optional] 
**Purl** | Pointer to **string** | Optional. The package URL for this package. | [optional] 

## Methods

### NewOsvPackage

`func NewOsvPackage() *OsvPackage`

NewOsvPackage instantiates a new OsvPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsvPackageWithDefaults

`func NewOsvPackageWithDefaults() *OsvPackage`

NewOsvPackageWithDefaults instantiates a new OsvPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcosystem

`func (o *OsvPackage) GetEcosystem() string`

GetEcosystem returns the Ecosystem field if non-nil, zero value otherwise.

### GetEcosystemOk

`func (o *OsvPackage) GetEcosystemOk() (*string, bool)`

GetEcosystemOk returns a tuple with the Ecosystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcosystem

`func (o *OsvPackage) SetEcosystem(v string)`

SetEcosystem sets Ecosystem field to given value.

### HasEcosystem

`func (o *OsvPackage) HasEcosystem() bool`

HasEcosystem returns a boolean if a field has been set.

### GetName

`func (o *OsvPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsvPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsvPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsvPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPurl

`func (o *OsvPackage) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *OsvPackage) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *OsvPackage) SetPurl(v string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *OsvPackage) HasPurl() bool`

HasPurl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


