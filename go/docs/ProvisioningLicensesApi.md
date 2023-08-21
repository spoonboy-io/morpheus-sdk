# \ProvisioningLicensesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProvisioningLicense**](ProvisioningLicensesApi.md#AddProvisioningLicense) | **Post** /api/provisioning-licenses | Create a License
[**GetProvisioningLicense**](ProvisioningLicensesApi.md#GetProvisioningLicense) | **Get** /api/provisioning-licenses/{id} | Get a Specific License
[**GetProvisioningLicenseReservations**](ProvisioningLicensesApi.md#GetProvisioningLicenseReservations) | **Get** /api/provisioning-licenses/{id}/reservations | Get Reservations for Specific License
[**ListProvisioningLicenses**](ProvisioningLicensesApi.md#ListProvisioningLicenses) | **Get** /api/provisioning-licenses | Get All Licenses
[**RemoveProvisioningLicense**](ProvisioningLicensesApi.md#RemoveProvisioningLicense) | **Delete** /api/provisioning-licenses/{id} | Delete a License
[**UpdateProvisioningLicense**](ProvisioningLicensesApi.md#UpdateProvisioningLicense) | **Put** /api/provisioning-licenses/{id} | Update a License



## AddProvisioningLicense

> Model200Success AddProvisioningLicense(ctx).InlineObject202(inlineObject202).Execute()

Create a License



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
    inlineObject202 := *openapiclient.NewInlineObject202() // InlineObject202 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvisioningLicensesApi.AddProvisioningLicense(context.Background()).InlineObject202(inlineObject202).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningLicensesApi.AddProvisioningLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProvisioningLicense`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningLicensesApi.AddProvisioningLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddProvisioningLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject202** | [**InlineObject202**](InlineObject202.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningLicense

> InlineResponse200126 GetProvisioningLicense(ctx, id).Execute()

Get a Specific License



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
    resp, r, err := api_client.ProvisioningLicensesApi.GetProvisioningLicense(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningLicensesApi.GetProvisioningLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvisioningLicense`: InlineResponse200126
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningLicensesApi.GetProvisioningLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200126**](inline_response_200_126.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningLicenseReservations

> InlineResponse200127 GetProvisioningLicenseReservations(ctx, id).Execute()

Get Reservations for Specific License



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
    resp, r, err := api_client.ProvisioningLicensesApi.GetProvisioningLicenseReservations(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningLicensesApi.GetProvisioningLicenseReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvisioningLicenseReservations`: InlineResponse200127
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningLicensesApi.GetProvisioningLicenseReservations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningLicenseReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200127**](inline_response_200_127.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProvisioningLicenses

> map[string]interface{} ListProvisioningLicenses(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).LicenseType(licenseType).LicenseVersion(licenseVersion).OrgName(orgName).FullName(fullName).Execute()

Get All Licenses



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
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    licenseType := "win" // string | If specified will return an exact match on licenseType code (optional)
    licenseVersion := "2012 R2 Standard" // string | If specified will return an exact match on licenseVersion (optional)
    orgName := "Acme Motors" // string | If specified will return an exact match on orgName (optional)
    fullName := "Bugs Bunny" // string | If specified will return an exact match on fullName (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvisioningLicensesApi.ListProvisioningLicenses(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Phrase(phrase).Name(name).LicenseType(licenseType).LicenseVersion(licenseVersion).OrgName(orgName).FullName(fullName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningLicensesApi.ListProvisioningLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProvisioningLicenses`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningLicensesApi.ListProvisioningLicenses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProvisioningLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **licenseType** | **string** | If specified will return an exact match on licenseType code | 
 **licenseVersion** | **string** | If specified will return an exact match on licenseVersion | 
 **orgName** | **string** | If specified will return an exact match on orgName | 
 **fullName** | **string** | If specified will return an exact match on fullName | 

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


## RemoveProvisioningLicense

> Model200Success RemoveProvisioningLicense(ctx, id).Execute()

Delete a License



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
    resp, r, err := api_client.ProvisioningLicensesApi.RemoveProvisioningLicense(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningLicensesApi.RemoveProvisioningLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveProvisioningLicense`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningLicensesApi.RemoveProvisioningLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProvisioningLicenseRequest struct via the builder pattern


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


## UpdateProvisioningLicense

> Model200Success UpdateProvisioningLicense(ctx, id).InlineObject203(inlineObject203).Execute()

Update a License



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
    inlineObject203 := *openapiclient.NewInlineObject203() // InlineObject203 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProvisioningLicensesApi.UpdateProvisioningLicense(context.Background(), id).InlineObject203(inlineObject203).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningLicensesApi.UpdateProvisioningLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProvisioningLicense`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningLicensesApi.UpdateProvisioningLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProvisioningLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject203** | [**InlineObject203**](InlineObject203.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

