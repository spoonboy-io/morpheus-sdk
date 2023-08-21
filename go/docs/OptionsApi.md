# \OptionsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOptionSourceData**](OptionsApi.md#GetOptionSourceData) | **Get** /api/options/{optionSource} | Get Option Source Data
[**ListCodeRepositories**](OptionsApi.md#ListCodeRepositories) | **Get** /api/options/codeRepositories | Retrieves a list of Code/GIT Repositories
[**ListOptionNetworkOptions**](OptionsApi.md#ListOptionNetworkOptions) | **Get** /api/options/zoneNetworkOptions | Retrieves network options by zone/cloud
[**ListOptionValues**](OptionsApi.md#ListOptionValues) | **Get** /api/options/list | Retrieves input option values



## GetOptionSourceData

> map[string]interface{} GetOptionSourceData(ctx, optionSource).Execute()

Get Option Source Data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    optionSource := "keyPairs" // string | `optionSource` to be listed

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OptionsApi.GetOptionSourceData(context.Background(), optionSource).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionsApi.GetOptionSourceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOptionSourceData`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OptionsApi.GetOptionSourceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optionSource** | **string** | &#x60;optionSource&#x60; to be listed | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionSourceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCodeRepositories

> Model200Success ListCodeRepositories(ctx).IntegrationId(integrationId).Execute()

Retrieves a list of Code/GIT Repositories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    integrationId := int64(33) // int64 | Filter by an integration Id. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OptionsApi.ListCodeRepositories(context.Background()).IntegrationId(integrationId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionsApi.ListCodeRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCodeRepositories`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `OptionsApi.ListCodeRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCodeRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationId** | **int64** | Filter by an integration Id. | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOptionNetworkOptions

> ZoneNetworkOptions ListOptionNetworkOptions(ctx).ZoneId(zoneId).ProvisionTypeId(provisionTypeId).Execute()

Retrieves network options by zone/cloud



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    zoneId := int64(3) // int64 | The Zone ID for Filtering (optional)
    provisionTypeId := int64(22) // int64 | Provision type filter, restricts query to only load service plans of specified provision type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OptionsApi.ListOptionNetworkOptions(context.Background()).ZoneId(zoneId).ProvisionTypeId(provisionTypeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionsApi.ListOptionNetworkOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOptionNetworkOptions`: ZoneNetworkOptions
    fmt.Fprintf(os.Stdout, "Response from `OptionsApi.ListOptionNetworkOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOptionNetworkOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **int64** | The Zone ID for Filtering | 
 **provisionTypeId** | **int64** | Provision type filter, restricts query to only load service plans of specified provision type | 

### Return type

[**ZoneNetworkOptions**](zoneNetworkOptions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOptionValues

> map[string]interface{} ListOptionValues(ctx).OptionTypeId(optionTypeId).Config(config).Execute()

Retrieves input option values



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    optionTypeId := int64(789) // int64 | Input or Option Type ID
    config := TODO // map[string]interface{} | Input parameters are required if the input is dependent on them.  Fields must be prefixed with `config.` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OptionsApi.ListOptionValues(context.Background()).OptionTypeId(optionTypeId).Config(config).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionsApi.ListOptionValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOptionValues`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OptionsApi.ListOptionValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOptionValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optionTypeId** | **int64** | Input or Option Type ID | 
 **config** | [**map[string]interface{}**](map[string]interface{}.md) | Input parameters are required if the input is dependent on them.  Fields must be prefixed with &#x60;config.&#x60; | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

