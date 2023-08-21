# MorpheusApi.LicenseApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getLicense**](LicenseApi.md#getLicense) | **GET** /api/license | Get license
[**installLicense**](LicenseApi.md#installLicense) | **POST** /api/license | Install license key
[**testLicense**](LicenseApi.md#testLicense) | **POST** /api/license/test | Test license key
[**uninstallLicense**](LicenseApi.md#uninstallLicense) | **DELETE** /api/license | Uninstall license key



## getLicense

> License getLicense()

Get license

Get details about the license that is currently installed on the appliance. This returns the license details, but not the key itself. Your license key will never be returned and should always be kept secret.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LicenseApi();
apiInstance.getLicense((error, data, response) => {
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

[**License**](License.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## installLicense

> License installLicense(opts)

Install license key

Install a new license key. This will potentially change the enabled features and capabilities of your Morpheus appliance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LicenseApi();
let opts = {
  'inlineObject125': new MorpheusApi.InlineObject125() // InlineObject125 | 
};
apiInstance.installLicense(opts, (error, data, response) => {
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
 **inlineObject125** | [**InlineObject125**](InlineObject125.md)|  | [optional] 

### Return type

[**License**](License.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## testLicense

> License testLicense(opts)

Test license key

This endpoint can be used to decode a license to see if it is valid and inspect the license settings, such as who it belongs to and the enabled features. This is only a test, it does not install the key, or make any changes to your appliance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LicenseApi();
let opts = {
  'inlineObject126': new MorpheusApi.InlineObject126() // InlineObject126 | 
};
apiInstance.testLicense(opts, (error, data, response) => {
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
 **inlineObject126** | [**InlineObject126**](InlineObject126.md)|  | [optional] 

### Return type

[**License**](License.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## uninstallLicense

> Model200Success uninstallLicense()

Uninstall license key

Uninstall your appliance license, leaving the appliance with no license installed. This will change the enabled features and capabilities of your Morpheus appliance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LicenseApi();
apiInstance.uninstallLicense((error, data, response) => {
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

