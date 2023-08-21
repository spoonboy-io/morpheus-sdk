# MorpheusApi.BackupSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listBackupSettings**](BackupSettingsApi.md#listBackupSettings) | **GET** /api/backup-settings | Get Backup Settings
[**updateBackupSettings**](BackupSettingsApi.md#updateBackupSettings) | **PUT** /api/backup-settings | Update Backup Settings



## listBackupSettings

> InlineResponse2009 listBackupSettings()

Get Backup Settings

This endpoint retrieves backup settings.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupSettingsApi();
apiInstance.listBackupSettings((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateBackupSettings

> Model200Success updateBackupSettings(opts)

Update Backup Settings

Update Backup Settings

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupSettingsApi();
let opts = {
  'inlineObject15': new MorpheusApi.InlineObject15() // InlineObject15 | 
};
apiInstance.updateBackupSettings(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject15** | [**InlineObject15**](InlineObject15.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

