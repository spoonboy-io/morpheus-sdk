# MorpheusApi.ApplianceSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listApplianceSettings**](ApplianceSettingsApi.md#listApplianceSettings) | **GET** /api/appliance-settings | Get Appliance Settings
[**reindex**](ApplianceSettingsApi.md#reindex) | **POST** /api/appliance-settings/reindex | Reindex Search
[**setApplianceSettingsMaintenanceMode**](ApplianceSettingsApi.md#setApplianceSettingsMaintenanceMode) | **POST** /api/appliance-settings/maintenance | Toggle Maintenance Mode
[**updateApplianceSettings**](ApplianceSettingsApi.md#updateApplianceSettings) | **PUT** /api/appliance-settings | Update Appliance Settings



## listApplianceSettings

> InlineResponse200 listApplianceSettings()

Get Appliance Settings

This endpoint retrieves appliance settings.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ApplianceSettingsApi();
apiInstance.listApplianceSettings((error, data, response) => {
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

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## reindex

> Model200Success reindex()

Reindex Search

Reindex all searchable data

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ApplianceSettingsApi();
apiInstance.reindex((error, data, response) => {
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

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## setApplianceSettingsMaintenanceMode

> Model200Success setApplianceSettingsMaintenanceMode(opts)

Toggle Maintenance Mode

This endpoint allows toggling the appliance maintenance mode.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ApplianceSettingsApi();
let opts = {
  'enabled': true // Boolean | Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa.
};
apiInstance.setApplianceSettingsMaintenanceMode(opts, (error, data, response) => {
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
 **enabled** | **Boolean**| Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa. | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateApplianceSettings

> Model200Success updateApplianceSettings(opts)

Update Appliance Settings

Update Appliance Settings

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ApplianceSettingsApi();
let opts = {
  'inlineObject2': new MorpheusApi.InlineObject2() // InlineObject2 | 
};
apiInstance.updateApplianceSettings(opts, (error, data, response) => {
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
 **inlineObject2** | [**InlineObject2**](InlineObject2.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

