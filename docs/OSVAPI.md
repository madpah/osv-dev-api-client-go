# \OSVAPI

All URIs are relative to *http://api.osv.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OSVDetermineVersion**](OSVAPI.md#OSVDetermineVersion) | **Post** /v1experimental/determineversion | Determine the version of the provided hash values.
[**OSVGetVulnById**](OSVAPI.md#OSVGetVulnById) | **Get** /v1/vulns/{id} | Return a &#x60;Vulnerability&#x60; object for a given OSV ID.
[**OSVQueryAffected**](OSVAPI.md#OSVQueryAffected) | **Post** /v1/query | Query vulnerabilities for a particular project at a given commit or version.
[**OSVQueryAffectedBatch**](OSVAPI.md#OSVQueryAffectedBatch) | **Post** /v1/querybatch | Query vulnerabilities (batched) for given package versions and commits. This currently allows a maximum of 1000 package versions to be included in a single query.



## OSVDetermineVersion

> V1VersionMatchList OSVDetermineVersion(ctx).Body(body).Execute()

Determine the version of the provided hash values.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	osvdev "github.com/madpah/osv-dev-api-client-go"
)

func main() {
	body := *osvdev.NewV1VersionQuery() // V1VersionQuery | 

	configuration := osvdev.NewConfiguration()
	apiClient := osvdev.NewAPIClient(configuration)
	resp, r, err := apiClient.OSVAPI.OSVDetermineVersion(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSVAPI.OSVDetermineVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OSVDetermineVersion`: V1VersionMatchList
	fmt.Fprintf(os.Stdout, "Response from `OSVAPI.OSVDetermineVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOSVDetermineVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1VersionQuery**](V1VersionQuery.md) |  | 

### Return type

[**V1VersionMatchList**](V1VersionMatchList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OSVGetVulnById

> OsvVulnerability OSVGetVulnById(ctx, id).Execute()

Return a `Vulnerability` object for a given OSV ID.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	osvdev "github.com/madpah/osv-dev-api-client-go"
)

func main() {
	id := "id_example" // string | 

	configuration := osvdev.NewConfiguration()
	apiClient := osvdev.NewAPIClient(configuration)
	resp, r, err := apiClient.OSVAPI.OSVGetVulnById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSVAPI.OSVGetVulnById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OSVGetVulnById`: OsvVulnerability
	fmt.Fprintf(os.Stdout, "Response from `OSVAPI.OSVGetVulnById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOSVGetVulnByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OsvVulnerability**](OsvVulnerability.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OSVQueryAffected

> V1VulnerabilityList OSVQueryAffected(ctx).Body(body).Execute()

Query vulnerabilities for a particular project at a given commit or version.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	osvdev "github.com/madpah/osv-dev-api-client-go"
)

func main() {
	body := *osvdev.NewV1Query() // V1Query | 

	configuration := osvdev.NewConfiguration()
	apiClient := osvdev.NewAPIClient(configuration)
	resp, r, err := apiClient.OSVAPI.OSVQueryAffected(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSVAPI.OSVQueryAffected``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OSVQueryAffected`: V1VulnerabilityList
	fmt.Fprintf(os.Stdout, "Response from `OSVAPI.OSVQueryAffected`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOSVQueryAffectedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1Query**](V1Query.md) |  | 

### Return type

[**V1VulnerabilityList**](V1VulnerabilityList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OSVQueryAffectedBatch

> V1BatchVulnerabilityList OSVQueryAffectedBatch(ctx).Body(body).Execute()

Query vulnerabilities (batched) for given package versions and commits. This currently allows a maximum of 1000 package versions to be included in a single query.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	osvdev "github.com/madpah/osv-dev-api-client-go"
)

func main() {
	body := *osvdev.NewV1BatchQuery() // V1BatchQuery | 

	configuration := osvdev.NewConfiguration()
	apiClient := osvdev.NewAPIClient(configuration)
	resp, r, err := apiClient.OSVAPI.OSVQueryAffectedBatch(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSVAPI.OSVQueryAffectedBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OSVQueryAffectedBatch`: V1BatchVulnerabilityList
	fmt.Fprintf(os.Stdout, "Response from `OSVAPI.OSVQueryAffectedBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOSVQueryAffectedBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1BatchQuery**](V1BatchQuery.md) |  | 

### Return type

[**V1BatchVulnerabilityList**](V1BatchVulnerabilityList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

