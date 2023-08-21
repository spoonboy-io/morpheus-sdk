# MorpheusApi.SSLCertificatesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCertificate**](SSLCertificatesApi.md#addCertificate) | **POST** /api/certificates | Create a Certificate
[**deleteCertificate**](SSLCertificatesApi.md#deleteCertificate) | **DELETE** /api/certificates/{id} | Delete a Certificate
[**getCertificate**](SSLCertificatesApi.md#getCertificate) | **GET** /api/certificates/{id} | Get a Specific Certificate
[**listCertificates**](SSLCertificatesApi.md#listCertificates) | **GET** /api/certificates | Get All SSL Certificates
[**updateCertificate**](SSLCertificatesApi.md#updateCertificate) | **PUT** /api/certificates/{id} | Update a Certificate



## addCertificate

> Object addCertificate(opts)

Create a Certificate

Create a Certificate

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SSLCertificatesApi();
let opts = {
  'inlineObject230': new MorpheusApi.InlineObject230() // InlineObject230 | 
};
apiInstance.addCertificate(opts, (error, data, response) => {
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
 **inlineObject230** | [**InlineObject230**](InlineObject230.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteCertificate

> Model200Success deleteCertificate(id)

Delete a Certificate

Will delete a certificate from the system and make it no longer usable.  If a certificate is actively in use, a delete will fail.  

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SSLCertificatesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteCertificate(id, (error, data, response) => {
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


## getCertificate

> InlineResponse200144 getCertificate(id)

Get a Specific Certificate

This endpoint retrieves a specific certificate.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SSLCertificatesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getCertificate(id, (error, data, response) => {
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

[**InlineResponse200144**](InlineResponse200144.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCertificates

> InlineResponse200143 listCertificates(opts)

Get All SSL Certificates

This endpoint retrieves all SSL certificates associated with the account.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SSLCertificatesApi();
let opts = {
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listCertificates(opts, (error, data, response) => {
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 

### Return type

[**InlineResponse200143**](InlineResponse200143.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateCertificate

> InlineResponse200144 updateCertificate(id, opts)

Update a Certificate

Update a Certificate.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SSLCertificatesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject231': new MorpheusApi.InlineObject231() // InlineObject231 | 
};
apiInstance.updateCertificate(id, opts, (error, data, response) => {
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
 **inlineObject231** | [**InlineObject231**](InlineObject231.md)|  | [optional] 

### Return type

[**InlineResponse200144**](InlineResponse200144.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

