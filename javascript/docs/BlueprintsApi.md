# MorpheusApi.BlueprintsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addBlueprint**](BlueprintsApi.md#addBlueprint) | **POST** /api/blueprints | Create a Blueprint
[**deleteBlueprint**](BlueprintsApi.md#deleteBlueprint) | **DELETE** /api/blueprints/{id} | Delete a Blueprint
[**getBlueprint**](BlueprintsApi.md#getBlueprint) | **GET** /api/blueprints/{id} | Get a Specific Blueprint
[**listBlueprints**](BlueprintsApi.md#listBlueprints) | **GET** /api/blueprints | Get All Blueprints
[**updateBlueprint**](BlueprintsApi.md#updateBlueprint) | **PUT** /api/blueprints/{id} | Updating a Blueprint
[**updateBlueprintImage**](BlueprintsApi.md#updateBlueprintImage) | **POST** /api/blueprints/{id}/image | Update Blueprint Image
[**updateBlueprintPermissions**](BlueprintsApi.md#updateBlueprintPermissions) | **PUT** /api/blueprints/{id}/update-permissions | Update Blueprint Permissions



## addBlueprint

> AddBlueprint200Response addBlueprint(opts)

Create a Blueprint

Create a Blueprint

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BlueprintsApi();
let opts = {
  'addBlueprintRequest': {$ref=../components/examples/blueprintARMCreate.json} // AddBlueprintRequest | 
};
apiInstance.addBlueprint(opts, (error, data, response) => {
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
 **addBlueprintRequest** | [**AddBlueprintRequest**](AddBlueprintRequest.md)|  | [optional] 

### Return type

[**AddBlueprint200Response**](AddBlueprint200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteBlueprint

> Model200Success deleteBlueprint(id)

Delete a Blueprint

Delete a Blueprint

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BlueprintsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteBlueprint(id, (error, data, response) => {
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


## getBlueprint

> GetBlueprint200Response getBlueprint(id)

Get a Specific Blueprint

This endpoint retrieves a specific blueprint.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BlueprintsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getBlueprint(id, (error, data, response) => {
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

[**GetBlueprint200Response**](GetBlueprint200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listBlueprints

> ListBlueprints200Response listBlueprints(opts)

Get All Blueprints

This endpoint retrieves all blueprints.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BlueprintsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listBlueprints(opts, (error, data, response) => {
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
 **max** | **Number**| Maximum number of records to return | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

[**ListBlueprints200Response**](ListBlueprints200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateBlueprint

> GetBlueprint200Response updateBlueprint(id, opts)

Updating a Blueprint

Update a Blueprint. This overwrites the entire config, so the entire blueprint config should be passed.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BlueprintsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'addBlueprintRequest': {$ref=../components/examples/blueprintUpdate.json} // AddBlueprintRequest | 
};
apiInstance.updateBlueprint(id, opts, (error, data, response) => {
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
 **addBlueprintRequest** | [**AddBlueprintRequest**](AddBlueprintRequest.md)|  | [optional] 

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateBlueprintImage

> GetBlueprint200Response updateBlueprintImage(id, opts)

Update Blueprint Image

Update Blueprint Image

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BlueprintsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'templateImage': "/path/to/file" // File | 
};
apiInstance.updateBlueprintImage(id, opts, (error, data, response) => {
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
 **templateImage** | **File**|  | [optional] 

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## updateBlueprintPermissions

> GetBlueprint200Response updateBlueprintPermissions(id, opts)

Update Blueprint Permissions

Update Blueprint Permissions

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BlueprintsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'updateBlueprintPermissionsRequest': {$ref=../components/examples/instanceTypeLayoutPermissionsUpdate.json} // UpdateBlueprintPermissionsRequest | 
};
apiInstance.updateBlueprintPermissions(id, opts, (error, data, response) => {
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
 **updateBlueprintPermissionsRequest** | [**UpdateBlueprintPermissionsRequest**](UpdateBlueprintPermissionsRequest.md)|  | [optional] 

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

