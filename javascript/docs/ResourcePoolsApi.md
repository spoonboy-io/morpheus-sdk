# MorpheusApi.ResourcePoolsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createResourcePoolGroup**](ResourcePoolsApi.md#createResourcePoolGroup) | **POST** /api/resource-pools/groups | Create a Resource Pool Group
[**deleteResourcePoolGroup**](ResourcePoolsApi.md#deleteResourcePoolGroup) | **DELETE** /api/resource-pools/groups/{id} | Delete a Resource Pool Group
[**getResourcePoolGroups**](ResourcePoolsApi.md#getResourcePoolGroups) | **GET** /api/resource-pools/groups | Get all Resource Pool Groups
[**getresourcePoolGroup**](ResourcePoolsApi.md#getresourcePoolGroup) | **GET** /api/resource-pools/groups/{id} | Get a Specific Resource Pool Group
[**updateResourcePoolGroup**](ResourcePoolsApi.md#updateResourcePoolGroup) | **PUT** /api/resource-pools/groups/{id} | Update a Resource Pool Group



## createResourcePoolGroup

> InlineResponse200132 createResourcePoolGroup(opts)

Create a Resource Pool Group

Use this command to create a resource pool group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ResourcePoolsApi();
let opts = {
  'inlineObject206': new MorpheusApi.InlineObject206() // InlineObject206 | 
};
apiInstance.createResourcePoolGroup(opts, (error, data, response) => {
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
 **inlineObject206** | [**InlineObject206**](InlineObject206.md)|  | [optional] 

### Return type

[**InlineResponse200132**](InlineResponse200132.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteResourcePoolGroup

> Model200Success deleteResourcePoolGroup(id)

Delete a Resource Pool Group

Will delete a Resource Pool Group from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ResourcePoolsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteResourcePoolGroup(id, (error, data, response) => {
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


## getResourcePoolGroups

> InlineResponse200131 getResourcePoolGroups()

Get all Resource Pool Groups

This endpoint retrieves all Resource Pool Groups associated with the account. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ResourcePoolsApi();
apiInstance.getResourcePoolGroups((error, data, response) => {
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

[**InlineResponse200131**](InlineResponse200131.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getresourcePoolGroup

> InlineResponse200132 getresourcePoolGroup(id)

Get a Specific Resource Pool Group

This endpoint retrieves a specific Resource Pool Group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ResourcePoolsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getresourcePoolGroup(id, (error, data, response) => {
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

[**InlineResponse200132**](InlineResponse200132.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateResourcePoolGroup

> InlineResponse200132 updateResourcePoolGroup(id, opts)

Update a Resource Pool Group

Use this command to update an existing resource pool Group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ResourcePoolsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject207': new MorpheusApi.InlineObject207() // InlineObject207 | 
};
apiInstance.updateResourcePoolGroup(id, opts, (error, data, response) => {
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
 **inlineObject207** | [**InlineObject207**](InlineObject207.md)|  | [optional] 

### Return type

[**InlineResponse200132**](InlineResponse200132.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

