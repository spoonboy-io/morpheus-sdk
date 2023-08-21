# MorpheusApi.VDIApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addVDIApps**](VDIApi.md#addVDIApps) | **POST** /api/vdi-apps | Creates a VDI App
[**addVDIGateways**](VDIApi.md#addVDIGateways) | **POST** /api/vdi-gateways | Creates a VDI Gateway
[**addVDIPools**](VDIApi.md#addVDIPools) | **POST** /api/vdi-pools | Creates a VDI Pool
[**addVdiAllocation**](VDIApi.md#addVdiAllocation) | **POST** /api/vdi/{id}/allocate | Allocate Virtual Desktop
[**getVDIAllocations**](VDIApi.md#getVDIAllocations) | **GET** /api/vdi-allocations/{id} | Retrieves a Specific VDI Allocation
[**getVDIApps**](VDIApi.md#getVDIApps) | **GET** /api/vdi-apps/{id} | Retrieves a Specific VDI App
[**getVDIGateways**](VDIApi.md#getVDIGateways) | **GET** /api/vdi-gateways/{id} | Retrieves a Specific VDI Gateway
[**getVDIPools**](VDIApi.md#getVDIPools) | **GET** /api/vdi-pools/{id} | Retrieves a Specific VDI Pool
[**getVdi**](VDIApi.md#getVdi) | **GET** /api/vdi/{id} | Get a Specific Virtual Desktop
[**listVDIAllocations**](VDIApi.md#listVDIAllocations) | **GET** /api/vdi-allocations | Retrieves all VDI Allocations
[**listVDIApps**](VDIApi.md#listVDIApps) | **GET** /api/vdi-apps | Retrieves all VDI Apps
[**listVDIGateways**](VDIApi.md#listVDIGateways) | **GET** /api/vdi-gateways | Retrieves all VDI Gateways
[**listVDIPools**](VDIApi.md#listVDIPools) | **GET** /api/vdi-pools | Retrieves all VDI Pools
[**listVdi**](VDIApi.md#listVdi) | **GET** /api/vdi | List Virtual Desktops
[**removeVDIApps**](VDIApi.md#removeVDIApps) | **DELETE** /api/vdi-apps/{id} | Deletes a VDI App
[**removeVDIGateways**](VDIApi.md#removeVDIGateways) | **DELETE** /api/vdi-gateways/{id} | Deletes a VDI Gateway
[**removeVDIPools**](VDIApi.md#removeVDIPools) | **DELETE** /api/vdi-pools/{id} | Deletes a VDI Pool
[**updateVDIApps**](VDIApi.md#updateVDIApps) | **PUT** /api/vdi-apps/{id} | Updates a VDI App Configuration or Icon
[**updateVDIGateways**](VDIApi.md#updateVDIGateways) | **PUT** /api/vdi-gateways/{id} | Updates a VDI Gateway Configuration
[**updateVDIPools**](VDIApi.md#updateVDIPools) | **PUT** /api/vdi-pools/{id} | Updates a VDI Pool Configuration or Icon



## addVDIApps

> AnyOfobject200Success addVDIApps(opts)

Creates a VDI App

Creates a VDI app. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let opts = {
  'inlineObject257': new MorpheusApi.InlineObject257() // InlineObject257 | 
};
apiInstance.addVDIApps(opts, (error, data, response) => {
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
 **inlineObject257** | [**InlineObject257**](InlineObject257.md)|  | [optional] 

### Return type

[**AnyOfobject200Success**](AnyOfobject200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json


## addVDIGateways

> AnyOfobject200Success addVDIGateways(opts)

Creates a VDI Gateway

Creates a VDI gateway. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let opts = {
  'inlineObject259': new MorpheusApi.InlineObject259() // InlineObject259 | 
};
apiInstance.addVDIGateways(opts, (error, data, response) => {
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
 **inlineObject259** | [**InlineObject259**](InlineObject259.md)|  | [optional] 

### Return type

[**AnyOfobject200Success**](AnyOfobject200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addVDIPools

> AnyOfobject200Success addVDIPools(opts)

Creates a VDI Pool

Creates a VDI pool. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let opts = {
  'inlineObject261': new MorpheusApi.InlineObject261() // InlineObject261 | 
};
apiInstance.addVDIPools(opts, (error, data, response) => {
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
 **inlineObject261** | [**InlineObject261**](InlineObject261.md)|  | [optional] 

### Return type

[**AnyOfobject200Success**](AnyOfobject200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json


## addVdiAllocation

> Object addVdiAllocation(id)

Allocate Virtual Desktop

This endpoint allocates a specific virtual desktop for use by your user. It will return the desktop and its allocation for your user, or an error if allocation fails, which will occur if the desktop is fully allocated already. If your user already has an allocation, the desktop and allocation will still be returned succesfully and the server does not make any changes. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.addVdiAllocation(id, (error, data, response) => {
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


## getVDIAllocations

> InlineResponse200163 getVDIAllocations(id)

Retrieves a Specific VDI Allocation

Retrieves a specific VDI allocation. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getVDIAllocations(id, (error, data, response) => {
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

[**InlineResponse200163**](InlineResponse200163.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getVDIApps

> InlineResponse200160 getVDIApps(id)

Retrieves a Specific VDI App

Retrieves a specific VDI app. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getVDIApps(id, (error, data, response) => {
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

[**InlineResponse200160**](InlineResponse200160.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getVDIGateways

> InlineResponse200161 getVDIGateways(id)

Retrieves a Specific VDI Gateway

Retrieves a specific VDI gateway. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getVDIGateways(id, (error, data, response) => {
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

[**InlineResponse200161**](InlineResponse200161.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getVDIPools

> InlineResponse200162 getVDIPools(id)

Retrieves a Specific VDI Pool

Retrieves a specific VDI pool. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getVDIPools(id, (error, data, response) => {
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

[**InlineResponse200162**](InlineResponse200162.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getVdi

> InlineResponse200164 getVdi(id)

Get a Specific Virtual Desktop

This endpoint retrieves a specific virtual desktop along with the allocation for your user if one exists. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getVdi(id, (error, data, response) => {
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

[**InlineResponse200164**](InlineResponse200164.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listVDIAllocations

> Object listVDIAllocations(opts)

Retrieves all VDI Allocations

Retrieves all VDI allocations. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'id': preparing, // String | Filter by allocation ID
  'status': preparing, // String | Filter by allocation status
  'poolId': 1, // Number | Filter by `VDI Pool` ID
  'userId': 6 // Number | Filter by User ID
};
apiInstance.listVDIAllocations(opts, (error, data, response) => {
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
 **id** | **String**| Filter by allocation ID | [optional] 
 **status** | **String**| Filter by allocation status | [optional] 
 **poolId** | **Number**| Filter by &#x60;VDI Pool&#x60; ID | [optional] 
 **userId** | **Number**| Filter by User ID | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listVDIApps

> Object listVDIApps(opts)

Retrieves all VDI Apps

Retrieves all VDI apps. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'description': The desription of my cool object // String | Filter by description, wildcard may be specified as %. eg. `example-%`
};
apiInstance.listVDIApps(opts, (error, data, response) => {
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
 **description** | **String**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listVDIGateways

> Object listVDIGateways(opts)

Retrieves all VDI Gateways

Retrieves all VDI gateways. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'description': The desription of my cool object // String | Filter by description, wildcard may be specified as %. eg. `example-%`
};
apiInstance.listVDIGateways(opts, (error, data, response) => {
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
 **description** | **String**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listVDIPools

> Object listVDIPools(opts)

Retrieves all VDI Pools

Retrieves all VDI pools. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'description': The desription of my cool object, // String | Filter by description, wildcard may be specified as %. eg. `example-%`
  'enabled': false // Boolean | Filter by enabled
};
apiInstance.listVDIPools(opts, (error, data, response) => {
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
 **description** | **String**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional] 
 **enabled** | **Boolean**| Filter by enabled | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listVdi

> Object listVdi(opts)

List Virtual Desktops

This endpoint retrieves all virtual desktops along with the allocation for your user if one exists. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let opts = {
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'description': The desription of my cool object // String | Filter by description, wildcard may be specified as %. eg. `example-%`
};
apiInstance.listVdi(opts, (error, data, response) => {
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
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **description** | **String**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeVDIApps

> Model200Success removeVDIApps(id)

Deletes a VDI App

Deletes a specified VDI App. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeVDIApps(id, (error, data, response) => {
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


## removeVDIGateways

> Model200Success removeVDIGateways(id)

Deletes a VDI Gateway

Deletes a specified VDI Gateway. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeVDIGateways(id, (error, data, response) => {
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


## removeVDIPools

> Model200Success removeVDIPools(id)

Deletes a VDI Pool

Deletes a specified VDI Pool. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeVDIPools(id, (error, data, response) => {
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


## updateVDIApps

> AnyOfobject200Success updateVDIApps(id, opts)

Updates a VDI App Configuration or Icon

Updates a VDI App configuration or icon. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject258': new MorpheusApi.InlineObject258() // InlineObject258 | 
};
apiInstance.updateVDIApps(id, opts, (error, data, response) => {
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
 **inlineObject258** | [**InlineObject258**](InlineObject258.md)|  | [optional] 

### Return type

[**AnyOfobject200Success**](AnyOfobject200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json


## updateVDIGateways

> AnyOfobject200Success updateVDIGateways(id, opts)

Updates a VDI Gateway Configuration

Updates a VDI Gateway configuration or icon. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject260': new MorpheusApi.InlineObject260() // InlineObject260 | 
};
apiInstance.updateVDIGateways(id, opts, (error, data, response) => {
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
 **inlineObject260** | [**InlineObject260**](InlineObject260.md)|  | [optional] 

### Return type

[**AnyOfobject200Success**](AnyOfobject200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateVDIPools

> AnyOfobject200Success updateVDIPools(id, opts)

Updates a VDI Pool Configuration or Icon

Updates a VDI Pool configuration or icon. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.VDIApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject262': new MorpheusApi.InlineObject262() // InlineObject262 | 
};
apiInstance.updateVDIPools(id, opts, (error, data, response) => {
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
 **inlineObject262** | [**InlineObject262**](InlineObject262.md)|  | [optional] 

### Return type

[**AnyOfobject200Success**](AnyOfobject200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

