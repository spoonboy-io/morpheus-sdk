# MorpheusApi.LogSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addLogSettingsSyslogRules**](LogSettingsApi.md#addLogSettingsSyslogRules) | **POST** /api/log-settings/syslog-rules | Create a New Syslog Rule
[**deleteLogSettingsSyslogRules**](LogSettingsApi.md#deleteLogSettingsSyslogRules) | **DELETE** /api/log-settings/syslog-rules/{id} | Delete a Specific Syslog Rule
[**listLogSettings**](LogSettingsApi.md#listLogSettings) | **GET** /api/log-settings | List All Log Settings
[**updateLogSettings**](LogSettingsApi.md#updateLogSettings) | **PUT** /api/log-settings | Update Log Settings



## addLogSettingsSyslogRules

> Model200Success addLogSettingsSyslogRules(opts)

Create a New Syslog Rule

Creates a new syslog rule. This command will also update existing syslog rule by specified name if already exists

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LogSettingsApi();
let opts = {
  'inlineObject140': new MorpheusApi.InlineObject140() // InlineObject140 | 
};
apiInstance.addLogSettingsSyslogRules(opts, (error, data, response) => {
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
 **inlineObject140** | [**InlineObject140**](InlineObject140.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteLogSettingsSyslogRules

> Model200Success deleteLogSettingsSyslogRules(id)

Delete a Specific Syslog Rule

Will delete the syslog rule matching the specified name.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LogSettingsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteLogSettingsSyslogRules(id, (error, data, response) => {
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
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listLogSettings

> InlineResponse20084 listLogSettings()

List All Log Settings

This endpoint retrieves log settings.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LogSettingsApi();
apiInstance.listLogSettings((error, data, response) => {
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

[**InlineResponse20084**](InlineResponse20084.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateLogSettings

> Model200Success updateLogSettings(opts)

Update Log Settings

Update Log Settings

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LogSettingsApi();
let opts = {
  'inlineObject139': new MorpheusApi.InlineObject139() // InlineObject139 | 
};
apiInstance.updateLogSettings(opts, (error, data, response) => {
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
 **inlineObject139** | [**InlineObject139**](InlineObject139.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

