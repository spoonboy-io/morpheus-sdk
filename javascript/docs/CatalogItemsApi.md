# MorpheusApi.CatalogItemsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCatalogItemType**](CatalogItemsApi.md#addCatalogItemType) | **POST** /api/catalog-item-types | Create a Catalog Item Type
[**getCatalogItemType**](CatalogItemsApi.md#getCatalogItemType) | **GET** /api/catalog-item-types/{id} | Get a Specific Catalog Item Type
[**listCatalogItemTypes**](CatalogItemsApi.md#listCatalogItemTypes) | **GET** /api/catalog-item-types | Get All Catalog Item Types
[**removeCatalogItemType**](CatalogItemsApi.md#removeCatalogItemType) | **DELETE** /api/catalog-item-types/{id} | Delete a Catalog Item Type
[**updateCatalogItemType**](CatalogItemsApi.md#updateCatalogItemType) | **PUT** /api/catalog-item-types/{id} | Update a Catalog Item Type
[**updateCatalogItemTypeLogo**](CatalogItemsApi.md#updateCatalogItemTypeLogo) | **PUT** /api/catalog-item-types/{id}/update-logo | Update Logo For Catalog Item Type



## addCatalogItemType

> AddCatalogItemType200Response addCatalogItemType(opts)

Create a Catalog Item Type

Use this command to create a catalog item type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CatalogItemsApi();
let opts = {
  'addCatalogItemTypeRequest': {$ref=../components/examples/catalogItemTypeInstanceCreate.json} // AddCatalogItemTypeRequest | 
};
apiInstance.addCatalogItemType(opts, (error, data, response) => {
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
 **addCatalogItemTypeRequest** | [**AddCatalogItemTypeRequest**](AddCatalogItemTypeRequest.md)|  | [optional] 

### Return type

[**AddCatalogItemType200Response**](AddCatalogItemType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json


## getCatalogItemType

> GetCatalogItemType200Response getCatalogItemType(id)

Get a Specific Catalog Item Type

This endpoint retrieves a specific category item type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CatalogItemsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getCatalogItemType(id, (error, data, response) => {
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

[**GetCatalogItemType200Response**](GetCatalogItemType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCatalogItemTypes

> ListCatalogItemTypes200Response listCatalogItemTypes(opts)

Get All Catalog Item Types

This endpoint retrieves all catalog item types.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CatalogItemsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'description': The desription of my cool object, // String | Filter by description, wildcard may be specified as %. eg. `example-%`
  'enabled': false, // Boolean | Filter by enabled
  'featured': false, // Boolean | Filter by featured
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example", // String | Filter by label(s), matches records that contain all of the specified labels
  'code': azr // String | If specified will return an exact match on code
};
apiInstance.listCatalogItemTypes(opts, (error, data, response) => {
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
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **description** | **String**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional] 
 **enabled** | **Boolean**| Filter by enabled | [optional] 
 **featured** | **Boolean**| Filter by featured | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 
 **code** | **String**| If specified will return an exact match on code | [optional] 

### Return type

[**ListCatalogItemTypes200Response**](ListCatalogItemTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeCatalogItemType

> Model200Success removeCatalogItemType(id)

Delete a Catalog Item Type

Will delete a catalog item type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CatalogItemsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeCatalogItemType(id, (error, data, response) => {
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


## updateCatalogItemType

> UpdateCatalogItemType200Response updateCatalogItemType(id, opts)

Update a Catalog Item Type

Use this command to update an existing catalog item type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CatalogItemsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'updateCatalogItemTypeRequest': {$ref=../components/examples/catalogItemTypeUpdate.json} // UpdateCatalogItemTypeRequest | 
};
apiInstance.updateCatalogItemType(id, opts, (error, data, response) => {
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
 **updateCatalogItemTypeRequest** | [**UpdateCatalogItemTypeRequest**](UpdateCatalogItemTypeRequest.md)|  | [optional] 

### Return type

[**UpdateCatalogItemType200Response**](UpdateCatalogItemType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json


## updateCatalogItemTypeLogo

> UpdateCatalogItemType200Response updateCatalogItemTypeLogo(id, opts)

Update Logo For Catalog Item Type

Use this command to update the logo and dark logo images for an existing catalog item type. This endpoint expects multipart form data as the request format, not JSON.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CatalogItemsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'catalogItemTypeLogo': "/path/to/file", // File | Logo File png,jpg,svg
  'catalogItemTypeDarkLogo': "/path/to/file" // File | Dark Logo File png,jpg,svg
};
apiInstance.updateCatalogItemTypeLogo(id, opts, (error, data, response) => {
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
 **catalogItemTypeLogo** | **File**| Logo File png,jpg,svg | [optional] 
 **catalogItemTypeDarkLogo** | **File**| Dark Logo File png,jpg,svg | [optional] 

### Return type

[**UpdateCatalogItemType200Response**](UpdateCatalogItemType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

