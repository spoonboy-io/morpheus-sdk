# MorpheusApi.MonitoringSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getMonitoringSettings**](MonitoringSettingsApi.md#getMonitoringSettings) | **GET** /api/monitoring-settings | Get Monitoring Settings
[**updateMonitoringSettings**](MonitoringSettingsApi.md#updateMonitoringSettings) | **PUT** /api/monitoring-settings | Update Monitoring Settings



## getMonitoringSettings

> InlineResponse20085 getMonitoringSettings()

Get Monitoring Settings

This endpoint retrieves monitoring settings.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.MonitoringSettingsApi();
apiInstance.getMonitoringSettings((error, data, response) => {
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

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateMonitoringSettings

> Model200Success updateMonitoringSettings(opts)

Update Monitoring Settings

Update Monitoring Settings

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.MonitoringSettingsApi();
let opts = {
  'inlineObject141': new MorpheusApi.InlineObject141() // InlineObject141 | 
};
apiInstance.updateMonitoringSettings(opts, (error, data, response) => {
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
 **inlineObject141** | [**InlineObject141**](InlineObject141.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

