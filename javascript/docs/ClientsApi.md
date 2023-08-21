# MorpheusApi.ClientsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addClient**](ClientsApi.md#addClient) | **POST** /api/clients | Create an Oauth Client
[**getClients**](ClientsApi.md#getClients) | **GET** /api/clients/{id} | Retrieves a Specific Oauth Client
[**listClients**](ClientsApi.md#listClients) | **GET** /api/clients | Get All Oauth Clients
[**removeClients**](ClientsApi.md#removeClients) | **DELETE** /api/clients/{id} | Deletes an Oauth Client
[**updateClients**](ClientsApi.md#updateClients) | **PUT** /api/clients/{id} | Updates an Oauth Client



## addClient

> Object addClient(opts)

Create an Oauth Client

Create a new Oauth Client.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClientsApi();
let opts = {
  'inlineObject27': new MorpheusApi.InlineObject27() // InlineObject27 | 
};
apiInstance.addClient(opts, (error, data, response) => {
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
 **inlineObject27** | [**InlineObject27**](InlineObject27.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getClients

> Object getClients(id)

Retrieves a Specific Oauth Client

Retrieves a specific oauth client. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClientsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getClients(id, (error, data, response) => {
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


## listClients

> Object listClients(opts)

Get All Oauth Clients

This endpoint retrieves a paginated list of oauth clients.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClientsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'clientId'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on clientId
  'clientId': "clientId_example" // String | Search phrase for partial matches on clientId
};
apiInstance.listClients(opts, (error, data, response) => {
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
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;clientId&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on clientId | [optional] 
 **clientId** | **String**| Search phrase for partial matches on clientId | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeClients

> Model200Success removeClients(id)

Deletes an Oauth Client

Deletes a specified oauth client. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClientsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeClients(id, (error, data, response) => {
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


## updateClients

> Object updateClients(id, opts)

Updates an Oauth Client

Updates an oauth client. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClientsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject28': new MorpheusApi.InlineObject28() // InlineObject28 | 
};
apiInstance.updateClients(id, opts, (error, data, response) => {
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
 **inlineObject28** | [**InlineObject28**](InlineObject28.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

