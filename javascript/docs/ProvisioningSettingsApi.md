# MorpheusApi.ProvisioningSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listProvisioningSettings**](ProvisioningSettingsApi.md#listProvisioningSettings) | **GET** /api/provisioning-settings | Get Provisioning Settings
[**updateProvisioningSettings**](ProvisioningSettingsApi.md#updateProvisioningSettings) | **PUT** /api/provisioning-settings | Update Provisioning Settings



## listProvisioningSettings

> InlineResponse200128 listProvisioningSettings()

Get Provisioning Settings

This endpoint retrieves provisioning settings.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ProvisioningSettingsApi();
apiInstance.listProvisioningSettings((error, data, response) => {
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

[**InlineResponse200128**](InlineResponse200128.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateProvisioningSettings

> Model200Success updateProvisioningSettings(opts)

Update Provisioning Settings

Update Provisioning Settings

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ProvisioningSettingsApi();
let opts = {
  'inlineObject204': new MorpheusApi.InlineObject204() // InlineObject204 | 
};
apiInstance.updateProvisioningSettings(opts, (error, data, response) => {
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
 **inlineObject204** | [**InlineObject204**](InlineObject204.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

