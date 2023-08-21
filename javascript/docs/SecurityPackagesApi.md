# MorpheusApi.SecurityPackagesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addSecurityPackages**](SecurityPackagesApi.md#addSecurityPackages) | **POST** /api/security-packages | Creates a Security Package
[**getSecurityPackages**](SecurityPackagesApi.md#getSecurityPackages) | **GET** /api/security-packages/{id} | Retrieves a Specific Security Package
[**listSecurityPackages**](SecurityPackagesApi.md#listSecurityPackages) | **GET** /api/security-packages | Retrieves all Security Packages
[**removeSecurityPackages**](SecurityPackagesApi.md#removeSecurityPackages) | **DELETE** /api/security-packages/{id} | Deletes a Security Package
[**updateSecurityPackages**](SecurityPackagesApi.md#updateSecurityPackages) | **PUT** /api/security-packages/{id} | Updates a Security Package



## addSecurityPackages

> Object addSecurityPackages(opts)

Creates a Security Package

Creates a security package. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityPackagesApi();
let opts = {
  'inlineObject218': new MorpheusApi.InlineObject218() // InlineObject218 | 
};
apiInstance.addSecurityPackages(opts, (error, data, response) => {
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
 **inlineObject218** | [**InlineObject218**](InlineObject218.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getSecurityPackages

> Object getSecurityPackages(id)

Retrieves a Specific Security Package

Retrieves a specific security package. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityPackagesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getSecurityPackages(id, (error, data, response) => {
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

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listSecurityPackages

> Object listSecurityPackages(opts)

Retrieves all Security Packages

Retrieves all security packages. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityPackagesApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listSecurityPackages(opts, (error, data, response) => {
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
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeSecurityPackages

> Model200Success removeSecurityPackages(id)

Deletes a Security Package

Deletes a specified security package. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityPackagesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeSecurityPackages(id, (error, data, response) => {
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


## updateSecurityPackages

> Object updateSecurityPackages(id, opts)

Updates a Security Package

Updates a security package. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityPackagesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject219': new MorpheusApi.InlineObject219() // InlineObject219 | 
};
apiInstance.updateSecurityPackages(id, opts, (error, data, response) => {
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
 **inlineObject219** | [**InlineObject219**](InlineObject219.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

