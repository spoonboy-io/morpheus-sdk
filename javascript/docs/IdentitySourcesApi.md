# MorpheusApi.IdentitySourcesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addIdentitySources**](IdentitySourcesApi.md#addIdentitySources) | **POST** /api/user-sources | Creates an Identity Source
[**getIdentitySources**](IdentitySourcesApi.md#getIdentitySources) | **GET** /api/user-sources/{id} | Retrieves a Specific Identity Source
[**listIdentitySources**](IdentitySourcesApi.md#listIdentitySources) | **GET** /api/user-sources | Retrieves all Identity Sources
[**removeIdentitySources**](IdentitySourcesApi.md#removeIdentitySources) | **DELETE** /api/user-sources/{id} | Deletes an Identity Source
[**updateIdentitySourceSubdomains**](IdentitySourcesApi.md#updateIdentitySourceSubdomains) | **PUT** /api/user-sources/{id}/subdomain | Updates an Identity Source Subdomain
[**updateIdentitySources**](IdentitySourcesApi.md#updateIdentitySources) | **PUT** /api/user-sources/{id} | Updates an Identity Source



## addIdentitySources

> Object addIdentitySources(opts)

Creates an Identity Source

Creates an identity source. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IdentitySourcesApi();
let opts = {
  'accountId': 3, // Number | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
  'userSourceCreate': new MorpheusApi.UserSourceCreate() // UserSourceCreate | 
};
apiInstance.addIdentitySources(opts, (error, data, response) => {
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
 **accountId** | **Number**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 
 **userSourceCreate** | [**UserSourceCreate**](UserSourceCreate.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getIdentitySources

> InlineResponse20051 getIdentitySources(id)

Retrieves a Specific Identity Source

Retrieves a specific identity source. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IdentitySourcesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getIdentitySources(id, (error, data, response) => {
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

[**InlineResponse20051**](InlineResponse20051.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listIdentitySources

> Object listIdentitySources(opts)

Retrieves all Identity Sources

Retrieves all identity sources. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IdentitySourcesApi();
let opts = {
  'type': "type_example", // String | If specified will return all tasks by `task type` code. Refer to `Task Types` API for up to date listings. 
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'accountId': 3 // Number | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
};
apiInstance.listIdentitySources(opts, (error, data, response) => {
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
 **type** | **String**| If specified will return all tasks by &#x60;task type&#x60; code. Refer to &#x60;Task Types&#x60; API for up to date listings.  | [optional] 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **accountId** | **Number**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeIdentitySources

> Model200Success removeIdentitySources(id)

Deletes an Identity Source

Deletes a specified identity source. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IdentitySourcesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeIdentitySources(id, (error, data, response) => {
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


## updateIdentitySourceSubdomains

> Object updateIdentitySourceSubdomains(id, opts)

Updates an Identity Source Subdomain

Updates an identity source subdomain. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IdentitySourcesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject82': new MorpheusApi.InlineObject82() // InlineObject82 | 
};
apiInstance.updateIdentitySourceSubdomains(id, opts, (error, data, response) => {
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
 **inlineObject82** | [**InlineObject82**](InlineObject82.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateIdentitySources

> Object updateIdentitySources(id, opts)

Updates an Identity Source

Updates an identity source. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IdentitySourcesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'userSourceCreate': new MorpheusApi.UserSourceCreate() // UserSourceCreate | 
};
apiInstance.updateIdentitySources(id, opts, (error, data, response) => {
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
 **userSourceCreate** | [**UserSourceCreate**](UserSourceCreate.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

