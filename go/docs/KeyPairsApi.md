# \KeyPairsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKeyPairs**](KeyPairsApi.md#AddKeyPairs) | **Post** /api/key-pairs | Creates a Key Pair
[**GenerateKeyPairs**](KeyPairsApi.md#GenerateKeyPairs) | **Post** /api/key-pairs/generate | Generates a Key Pair
[**GetKeyPairs**](KeyPairsApi.md#GetKeyPairs) | **Get** /api/key-pairs/{id} | Retrieves a Specific Key Pair
[**RemoveKeyPairs**](KeyPairsApi.md#RemoveKeyPairs) | **Delete** /api/key-pairs/{id} | Deletes a Key Pair



## AddKeyPairs

> map[string]interface{} AddKeyPairs(ctx).InlineObject105(inlineObject105).Execute()

Creates a Key Pair



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
    inlineObject105 := *openapiclient.NewInlineObject105(*openapiclient.NewApiKeyPairsKeyPair("Name_example", "PublicKey_example")) // InlineObject105 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyPairsApi.AddKeyPairs(context.Background()).InlineObject105(inlineObject105).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.AddKeyPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddKeyPairs`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.AddKeyPairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddKeyPairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject105** | [**InlineObject105**](InlineObject105.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateKeyPairs

> InlineResponse20066 GenerateKeyPairs(ctx).InlineObject106(inlineObject106).Execute()

Generates a Key Pair



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
    inlineObject106 := *openapiclient.NewInlineObject106(*openapiclient.NewApiKeyPairsGenerateKeyPair("Name_example")) // InlineObject106 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyPairsApi.GenerateKeyPairs(context.Background()).InlineObject106(inlineObject106).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.GenerateKeyPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateKeyPairs`: InlineResponse20066
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.GenerateKeyPairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateKeyPairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject106** | [**InlineObject106**](InlineObject106.md) |  | 

### Return type

[**InlineResponse20066**](inline_response_200_66.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyPairs

> InlineResponse20067 GetKeyPairs(ctx, id).Execute()

Retrieves a Specific Key Pair



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
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyPairsApi.GetKeyPairs(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.GetKeyPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyPairs`: InlineResponse20067
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.GetKeyPairs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyPairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20067**](inline_response_200_67.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveKeyPairs

> Model200Success RemoveKeyPairs(ctx, id).Execute()

Deletes a Key Pair



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
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeyPairsApi.RemoveKeyPairs(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairsApi.RemoveKeyPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveKeyPairs`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `KeyPairsApi.RemoveKeyPairs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveKeyPairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

