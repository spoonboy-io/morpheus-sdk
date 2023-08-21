# \SetupApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Setup**](SetupApi.md#Setup) | **Post** /api/setup | Setup appliance



## Setup

> Setup Setup(ctx).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()

Setup appliance



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
    uNKNOWNBASETYPE := TODO // UNKNOWN_BASE_TYPE |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SetupApi.Setup(context.Background()).UNKNOWNBASETYPE(uNKNOWNBASETYPE).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.Setup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Setup`: Setup
    fmt.Fprintf(os.Stdout, "Response from `SetupApi.Setup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md) |  | 

### Return type

[**Setup**](setup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

