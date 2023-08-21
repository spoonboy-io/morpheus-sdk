# \BackupSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListBackupSettings**](BackupSettingsApi.md#ListBackupSettings) | **Get** /api/backup-settings | Get Backup Settings
[**UpdateBackupSettings**](BackupSettingsApi.md#UpdateBackupSettings) | **Put** /api/backup-settings | Update Backup Settings



## ListBackupSettings

> InlineResponse2009 ListBackupSettings(ctx).Execute()

Get Backup Settings



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupSettingsApi.ListBackupSettings(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupSettingsApi.ListBackupSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBackupSettings`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `BackupSettingsApi.ListBackupSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupSettingsRequest struct via the builder pattern


### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBackupSettings

> Model200Success UpdateBackupSettings(ctx).InlineObject15(inlineObject15).Execute()

Update Backup Settings



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
    inlineObject15 := *openapiclient.NewInlineObject15() // InlineObject15 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupSettingsApi.UpdateBackupSettings(context.Background()).InlineObject15(inlineObject15).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupSettingsApi.UpdateBackupSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBackupSettings`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `BackupSettingsApi.UpdateBackupSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBackupSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject15** | [**InlineObject15**](InlineObject15.md) |  | 

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

