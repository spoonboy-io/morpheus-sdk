# MorpheusApi.WhitelabelSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getWhitelabelImage**](WhitelabelSettingsApi.md#getWhitelabelImage) | **GET** /api/whitelabel-settings/images/{imageType} | Download Image
[**listWhitelabelSettings**](WhitelabelSettingsApi.md#listWhitelabelSettings) | **GET** /api/whitelabel-settings | Get Whitelabel Settings
[**removeWhitelabelImage**](WhitelabelSettingsApi.md#removeWhitelabelImage) | **DELETE** /api/whitelabel-settings/images/{imageType} | Reset Image
[**updateWhitelabelImages**](WhitelabelSettingsApi.md#updateWhitelabelImages) | **POST** /api/whitelabel-settings/images | Update Images
[**updateWhitelabelSettings**](WhitelabelSettingsApi.md#updateWhitelabelSettings) | **PUT** /api/whitelabel-settings | Update Whitelabel Settings



## getWhitelabelImage

> File getWhitelabelImage(imageType)

Download Image

Downloads the specified image.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.WhitelabelSettingsApi();
let imageType = headerLogo; // String | Valid image types
apiInstance.getWhitelabelImage(imageType, (error, data, response) => {
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
 **imageType** | **String**| Valid image types | 

### Return type

**File**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/ico, image/jpeg, image/png, image/svg+xml, application/json


## listWhitelabelSettings

> InlineResponse200166 listWhitelabelSettings()

Get Whitelabel Settings

This endpoint retrieves whitelabel settings.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.WhitelabelSettingsApi();
apiInstance.listWhitelabelSettings((error, data, response) => {
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

[**InlineResponse200166**](InlineResponse200166.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeWhitelabelImage

> Model200Success removeWhitelabelImage(imageType)

Reset Image

Resets the specified image to the Morpheus default.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.WhitelabelSettingsApi();
let imageType = headerLogo; // String | Valid image types
apiInstance.removeWhitelabelImage(imageType, (error, data, response) => {
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
 **imageType** | **String**| Valid image types | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateWhitelabelImages

> Model200Success updateWhitelabelImages(opts)

Update Images

Uploads whitelabel images. Expects multipart form data as the request format, not JSON.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.WhitelabelSettingsApi();
let opts = {
  'headerLogoFile': "/path/to/file", // File | Header logo image file, valid image types `png|jpg|svg`
  'resetHeaderLogo': true, // Boolean | Resets header logo to default
  'footerLogoFile': "/path/to/file", // File | Footer logo image file, valid image types `png|jpg|svg`
  'resetFooterLogo': true, // Boolean | Resets footer logo to default
  'loginLogoFile': "/path/to/file", // File | Login logo image file, valid image types `png|jpg|svg`
  'resetLoginLogo': true, // Boolean | Resets login logo to default
  'faviconFile': "/path/to/file", // File | Favicon image file, valid image type ico
  'resetFaviconLogo': true // Boolean | Resets favicon logo to default
};
apiInstance.updateWhitelabelImages(opts, (error, data, response) => {
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
 **headerLogoFile** | **File**| Header logo image file, valid image types &#x60;png|jpg|svg&#x60; | [optional] 
 **resetHeaderLogo** | **Boolean**| Resets header logo to default | [optional] 
 **footerLogoFile** | **File**| Footer logo image file, valid image types &#x60;png|jpg|svg&#x60; | [optional] 
 **resetFooterLogo** | **Boolean**| Resets footer logo to default | [optional] 
 **loginLogoFile** | **File**| Login logo image file, valid image types &#x60;png|jpg|svg&#x60; | [optional] 
 **resetLoginLogo** | **Boolean**| Resets login logo to default | [optional] 
 **faviconFile** | **File**| Favicon image file, valid image type ico | [optional] 
 **resetFaviconLogo** | **Boolean**| Resets favicon logo to default | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## updateWhitelabelSettings

> Model200Success updateWhitelabelSettings(opts)

Update Whitelabel Settings

Update Whitelabel Settings

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.WhitelabelSettingsApi();
let opts = {
  'inlineObject265': new MorpheusApi.InlineObject265() // InlineObject265 | 
};
apiInstance.updateWhitelabelSettings(opts, (error, data, response) => {
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
 **inlineObject265** | [**InlineObject265**](InlineObject265.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

