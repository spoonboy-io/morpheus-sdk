# MorpheusApi.GuidanceSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getGuidanceSettings**](GuidanceSettingsApi.md#getGuidanceSettings) | **GET** /api/guidance-settings | Get Guidance Settings
[**updateGuidanceSettings**](GuidanceSettingsApi.md#updateGuidanceSettings) | **PUT** /api/guidance-settings | Update Guidance Settings



## getGuidanceSettings

> InlineResponse20047 getGuidanceSettings()

Get Guidance Settings

This endpoint retrieves guidance settings.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.GuidanceSettingsApi();
apiInstance.getGuidanceSettings((error, data, response) => {
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

[**InlineResponse20047**](InlineResponse20047.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateGuidanceSettings

> Model200Success updateGuidanceSettings(opts)

Update Guidance Settings

Update Guidance Settings

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.GuidanceSettingsApi();
let opts = {
  'inlineObject79': new MorpheusApi.InlineObject79() // InlineObject79 | 
};
apiInstance.updateGuidanceSettings(opts, (error, data, response) => {
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
 **inlineObject79** | [**InlineObject79**](InlineObject79.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

