# MorpheusApi.ClusterLayoutsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addClusterLayoutClone**](ClusterLayoutsApi.md#addClusterLayoutClone) | **POST** /api/library/cluster-layouts/{id}/clone | Clone a Cluster Layout
[**addClusterLayouts**](ClusterLayoutsApi.md#addClusterLayouts) | **POST** /api/library/cluster-layouts | Create a Cluster Layout
[**deleteClusterLayout**](ClusterLayoutsApi.md#deleteClusterLayout) | **DELETE** /api/library/cluster-layouts/{id} | Delete a Cluster Layout
[**getClusterLayout**](ClusterLayoutsApi.md#getClusterLayout) | **GET** /api/library/cluster-layouts/{id} | Get a Specific Cluster Layout
[**listClusterLayouts**](ClusterLayoutsApi.md#listClusterLayouts) | **GET** /api/library/cluster-layouts | Get All Cluster Layouts
[**updateClusterLayout**](ClusterLayoutsApi.md#updateClusterLayout) | **PUT** /api/library/cluster-layouts/{id} | Update a Cluster Layout



## addClusterLayoutClone

> SuccessId addClusterLayoutClone(id, opts)

Clone a Cluster Layout

Use this command to clone a cluster layout.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClusterLayoutsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'name': New Name, // String | Name of cluster layout. Defaults to Copy of <cloned layout name>
  'description': New Description, // String | Description of cluster layout. Defaults to cloned layout description
  'computeVersion': 2.0 // String | Version of cluster layout. Defaults to cloned layout version
};
apiInstance.addClusterLayoutClone(id, opts, (error, data, response) => {
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
 **name** | **String**| Name of cluster layout. Defaults to Copy of &lt;cloned layout name&gt; | [optional] 
 **description** | **String**| Description of cluster layout. Defaults to cloned layout description | [optional] 
 **computeVersion** | **String**| Version of cluster layout. Defaults to cloned layout version | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## addClusterLayouts

> SuccessId addClusterLayouts(opts)

Create a Cluster Layout

Use this command to create a cluster layout.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClusterLayoutsApi();
let opts = {
  'inlineObject50': new MorpheusApi.InlineObject50() // InlineObject50 | 
};
apiInstance.addClusterLayouts(opts, (error, data, response) => {
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
 **inlineObject50** | [**InlineObject50**](InlineObject50.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteClusterLayout

> Model200Success deleteClusterLayout(id)

Delete a Cluster Layout

Will delete a cluster layout

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClusterLayoutsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteClusterLayout(id, (error, data, response) => {
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


## getClusterLayout

> InlineResponse20025 getClusterLayout(id)

Get a Specific Cluster Layout

This endpoint retrieves a specific cluster layout.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClusterLayoutsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getClusterLayout(id, (error, data, response) => {
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

[**InlineResponse20025**](InlineResponse20025.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClusterLayouts

> Object listClusterLayouts(opts)

Get All Cluster Layouts

This endpoint retrieves all cluster layouts.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClusterLayoutsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'provisionType': "provisionType_example", // String | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings. 
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listClusterLayouts(opts, (error, data, response) => {
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
 **provisionType** | **String**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateClusterLayout

> SuccessId updateClusterLayout(id, opts)

Update a Cluster Layout

Use this command to update an existing cluster layout.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ClusterLayoutsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject51': new MorpheusApi.InlineObject51() // InlineObject51 | 
};
apiInstance.updateClusterLayout(id, opts, (error, data, response) => {
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
 **inlineObject51** | [**InlineObject51**](InlineObject51.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

