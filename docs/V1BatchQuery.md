# V1BatchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | Pointer to [**[]V1Query**](V1Query.md) | The queries that form this batch query. | [optional] 

## Methods

### NewV1BatchQuery

`func NewV1BatchQuery() *V1BatchQuery`

NewV1BatchQuery instantiates a new V1BatchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BatchQueryWithDefaults

`func NewV1BatchQueryWithDefaults() *V1BatchQuery`

NewV1BatchQueryWithDefaults instantiates a new V1BatchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *V1BatchQuery) GetQueries() []V1Query`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *V1BatchQuery) GetQueriesOk() (*[]V1Query, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *V1BatchQuery) SetQueries(v []V1Query)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *V1BatchQuery) HasQueries() bool`

HasQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


