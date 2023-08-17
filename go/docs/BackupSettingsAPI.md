# \BackupSettingsAPI

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListBackupSettings**](BackupSettingsAPI.md#ListBackupSettings) | **Get** /api/backup-settings | Get Backup Settings
[**UpdateBackupSettings**](BackupSettingsAPI.md#UpdateBackupSettings) | **Put** /api/backup-settings | Update Backup Settings



## ListBackupSettings

> ListBackupSettings200Response ListBackupSettings(ctx).Execute()

Get Backup Settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupSettingsAPI.ListBackupSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupSettingsAPI.ListBackupSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBackupSettings`: ListBackupSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `BackupSettingsAPI.ListBackupSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupSettingsRequest struct via the builder pattern


### Return type

[**ListBackupSettings200Response**](ListBackupSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBackupSettings

> Model200Success UpdateBackupSettings(ctx).UpdateBackupSettingsRequest(updateBackupSettingsRequest).Execute()

Update Backup Settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    updateBackupSettingsRequest := *openapiclient.NewUpdateBackupSettingsRequest() // UpdateBackupSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupSettingsAPI.UpdateBackupSettings(context.Background()).UpdateBackupSettingsRequest(updateBackupSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupSettingsAPI.UpdateBackupSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBackupSettings`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `BackupSettingsAPI.UpdateBackupSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBackupSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateBackupSettingsRequest** | [**UpdateBackupSettingsRequest**](UpdateBackupSettingsRequest.md) |  | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

