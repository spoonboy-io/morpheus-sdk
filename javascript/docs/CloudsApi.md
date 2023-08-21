# MorpheusApi.CloudsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCloudResourcePool**](CloudsApi.md#addCloudResourcePool) | **POST** /api/zones/{zoneId}/resource-pools | Creates a Specified Resource Pool for Specified Cloud
[**addClouds**](CloudsApi.md#addClouds) | **POST** /api/zones | Creates a Cloud
[**getCloudDatastores**](CloudsApi.md#getCloudDatastores) | **GET** /api/zones/{zoneId}/data-stores/{id} | Retrieves a Datastore for Specified Cloud
[**getCloudFolders**](CloudsApi.md#getCloudFolders) | **GET** /api/zones/{zoneId}/folders/{id} | Retrieves a Resource Folder for Specified Cloud
[**getCloudResourcePools**](CloudsApi.md#getCloudResourcePools) | **GET** /api/zones/{zoneId}/resource-pools/{id} | Retrieves a Resource Pool for Specified Cloud
[**getCloudTypes**](CloudsApi.md#getCloudTypes) | **GET** /api/zone-types/{id} | Retrieves a Specific Cloud Type
[**getClouds**](CloudsApi.md#getClouds) | **GET** /api/zones/{id} | Retrieves a Specific Cloud
[**getWikiCloud**](CloudsApi.md#getWikiCloud) | **GET** /api/zones/{id}/wiki | Retrieves a Cloud Wiki Page
[**listCloudDatastores**](CloudsApi.md#listCloudDatastores) | **GET** /api/zones/{zoneId}/data-stores | Retrieves all Datastores for Specified Cloud
[**listCloudFolders**](CloudsApi.md#listCloudFolders) | **GET** /api/zones/{zoneId}/folders | Retrieves all resource folders for Specified Cloud
[**listCloudResourcePools**](CloudsApi.md#listCloudResourcePools) | **GET** /api/zones/{zoneId}/resource-pools | Retrieves all Resource Pools for Specified Cloud
[**listCloudSecurityGroups**](CloudsApi.md#listCloudSecurityGroups) | **GET** /api/zones/{id}/security-groups | Retrieves all Security Groups for a Cloud
[**listCloudTypes**](CloudsApi.md#listCloudTypes) | **GET** /api/zone-types | Retrieves all Cloud Types
[**listClouds**](CloudsApi.md#listClouds) | **GET** /api/zones | Retrieves all Clouds
[**refreshClouds**](CloudsApi.md#refreshClouds) | **POST** /api/zones/{id}/refresh | Refreshes a Cloud
[**removeCloudResourcePools**](CloudsApi.md#removeCloudResourcePools) | **DELETE** /api/zones/{zoneId}/resource-pools/{id} | Deletes a Resource Pool for Specified Cloud
[**removeClouds**](CloudsApi.md#removeClouds) | **DELETE** /api/zones/{id} | Deletes a Cloud
[**updateCloudDatastores**](CloudsApi.md#updateCloudDatastores) | **PUT** /api/zones/{zoneId}/data-stores/{id} | Updates a Specified Datastore for Specified Cloud
[**updateCloudFolders**](CloudsApi.md#updateCloudFolders) | **PUT** /api/zones/{zoneId}/folders/{id} | Updates a Resource Folder for Specified Cloud
[**updateCloudLogo**](CloudsApi.md#updateCloudLogo) | **POST** /api/zones/{id}/update-logo | Update Logo For Cloud
[**updateCloudResourcePool**](CloudsApi.md#updateCloudResourcePool) | **PUT** /api/zones/{zoneId}/resource-pools/{id} | Updates a Specified Resource Pool for Specified Cloud
[**updateCloudSecurityGroups**](CloudsApi.md#updateCloudSecurityGroups) | **POST** /api/zones/{id}/security-groups | Sets Security Groups for a Cloud
[**updateClouds**](CloudsApi.md#updateClouds) | **PUT** /api/zones/{id} | Updates a Cloud
[**updateWikiCloud**](CloudsApi.md#updateWikiCloud) | **PUT** /api/zones/{id}/wiki | Update a Cloud Wiki Page



## addCloudResourcePool

> InlineResponse20024 addCloudResourcePool(zoneId, opts)

Creates a Specified Resource Pool for Specified Cloud

Creates a resource pool for specified cloud. Only certain types of clouds support creating and deleting resource pools. Configuration options vary by type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let zoneId = 7; // Number | The ID of the cloud
let opts = {
  'inlineObject45': new MorpheusApi.InlineObject45() // InlineObject45 | 
};
apiInstance.addCloudResourcePool(zoneId, opts, (error, data, response) => {
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
 **zoneId** | **Number**| The ID of the cloud | 
 **inlineObject45** | [**InlineObject45**](InlineObject45.md)|  | [optional] 

### Return type

[**InlineResponse20024**](InlineResponse20024.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addClouds

> Object addClouds(opts)

Creates a Cloud

Creates a cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let opts = {
  'inlineObject41': new MorpheusApi.InlineObject41() // InlineObject41 | 
};
apiInstance.addClouds(opts, (error, data, response) => {
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
 **inlineObject41** | [**InlineObject41**](InlineObject41.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getCloudDatastores

> Object getCloudDatastores(zoneId, id)

Retrieves a Datastore for Specified Cloud

Data Stores can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves a specific datastore under a cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let zoneId = 7; // Number | The ID of the cloud
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getCloudDatastores(zoneId, id, (error, data, response) => {
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
 **zoneId** | **Number**| The ID of the cloud | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getCloudFolders

> Object getCloudFolders(zoneId, id)

Retrieves a Resource Folder for Specified Cloud

Resource Folders can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves a specific folder under a cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let zoneId = 7; // Number | The ID of the cloud
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getCloudFolders(zoneId, id, (error, data, response) => {
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
 **zoneId** | **Number**| The ID of the cloud | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getCloudResourcePools

> Object getCloudResourcePools(zoneId, id)

Retrieves a Resource Pool for Specified Cloud

This endpoint retrieves a specific resource pool. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let zoneId = 7; // Number | The ID of the cloud
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getCloudResourcePools(zoneId, id, (error, data, response) => {
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
 **zoneId** | **Number**| The ID of the cloud | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getCloudTypes

> InlineResponse20020 getCloudTypes(id)

Retrieves a Specific Cloud Type

Retrieves a specific cloud type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getCloudTypes(id, (error, data, response) => {
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

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getClouds

> InlineResponse20021 getClouds(id)

Retrieves a Specific Cloud

Retrieves a specific cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getClouds(id, (error, data, response) => {
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

[**InlineResponse20021**](InlineResponse20021.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getWikiCloud

> InlineResponse200168 getWikiCloud(id)

Retrieves a Cloud Wiki Page

This endpoint retrieves a cloud Wiki page. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getWikiCloud(id, (error, data, response) => {
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

[**InlineResponse200168**](InlineResponse200168.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCloudDatastores

> Object listCloudDatastores(zoneId, opts)

Retrieves all Datastores for Specified Cloud

Data Stores can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves all data stores under a cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let zoneId = 7; // Number | The ID of the cloud
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc // String | Sort direction, use 'desc' to reverse sort
};
apiInstance.listCloudDatastores(zoneId, opts, (error, data, response) => {
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
 **zoneId** | **Number**| The ID of the cloud | 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCloudFolders

> Object listCloudFolders(zoneId, opts)

Retrieves all resource folders for Specified Cloud

Resource Folders can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves all resource folders under a cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let zoneId = 7; // Number | The ID of the cloud
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'max': 25 // Number | Maximum number of records to return, -1 can be used to fetch all records
};
apiInstance.listCloudFolders(zoneId, opts, (error, data, response) => {
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
 **zoneId** | **Number**| The ID of the cloud | 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCloudResourcePools

> Object listCloudResourcePools(zoneId, opts)

Retrieves all Resource Pools for Specified Cloud

Resource Pools can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves all resource pools under a cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let zoneId = 7; // Number | The ID of the cloud
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'max': 25 // Number | Maximum number of records to return, -1 can be used to fetch all records
};
apiInstance.listCloudResourcePools(zoneId, opts, (error, data, response) => {
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
 **zoneId** | **Number**| The ID of the cloud | 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCloudSecurityGroups

> Object listCloudSecurityGroups(id)

Retrieves all Security Groups for a Cloud

Retrieves all security groups for a cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.listCloudSecurityGroups(id, (error, data, response) => {
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


## listCloudTypes

> Object listCloudTypes(opts)

Retrieves all Cloud Types

Fetch a paginated list of available cloud types. This returns the configuration options for each type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'provisionType': "provisionType_example" // String | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings. 
};
apiInstance.listCloudTypes(opts, (error, data, response) => {
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **code** | **String**| If specified will return an exact match on code | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **provisionType** | **String**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listClouds

> Object listClouds(opts)

Retrieves all Clouds

Retrieves all clouds. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let opts = {
  'lastUpdated': 2019-03-06T17:52:29+0000, // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
  'type': "type_example", // String | If specified will return all zones by cloud type code. Refer to `Zone Types` API for up to date listings. 
  'groupId': 13, // Number | If specified will return all zones assigned to a server group by id. Accepts multiple values.
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listClouds(opts, (error, data, response) => {
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
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 
 **type** | **String**| If specified will return all zones by cloud type code. Refer to &#x60;Zone Types&#x60; API for up to date listings.  | [optional] 
 **groupId** | **Number**| If specified will return all zones assigned to a server group by id. Accepts multiple values. | [optional] 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## refreshClouds

> Model200Success refreshClouds(id, opts)

Refreshes a Cloud

Refreshes a cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject47': new MorpheusApi.InlineObject47() // InlineObject47 | 
};
apiInstance.refreshClouds(id, opts, (error, data, response) => {
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
 **inlineObject47** | [**InlineObject47**](InlineObject47.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## removeCloudResourcePools

> Model200Success removeCloudResourcePools(zoneId, id)

Deletes a Resource Pool for Specified Cloud

Deletes a resource pool for specified Cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let zoneId = 7; // Number | The ID of the cloud
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeCloudResourcePools(zoneId, id, (error, data, response) => {
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
 **zoneId** | **Number**| The ID of the cloud | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeClouds

> Model200Success removeClouds(id, opts)

Deletes a Cloud

Deletes a specified Cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'removeResources': false // Boolean | Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute.
};
apiInstance.removeClouds(id, opts, (error, data, response) => {
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
 **removeResources** | **Boolean**| Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute. | [optional] [default to false]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateCloudDatastores

> InlineResponse20022 updateCloudDatastores(zoneId, id, opts)

Updates a Specified Datastore for Specified Cloud

Updates a datastore for specified cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let zoneId = 7; // Number | The ID of the cloud
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject43': new MorpheusApi.InlineObject43() // InlineObject43 | 
};
apiInstance.updateCloudDatastores(zoneId, id, opts, (error, data, response) => {
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
 **zoneId** | **Number**| The ID of the cloud | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject43** | [**InlineObject43**](InlineObject43.md)|  | [optional] 

### Return type

[**InlineResponse20022**](InlineResponse20022.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateCloudFolders

> InlineResponse20023 updateCloudFolders(zoneId, id, opts)

Updates a Resource Folder for Specified Cloud

Updates a resource folder for specified cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let zoneId = 7; // Number | The ID of the cloud
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject44': new MorpheusApi.InlineObject44() // InlineObject44 | 
};
apiInstance.updateCloudFolders(zoneId, id, opts, (error, data, response) => {
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
 **zoneId** | **Number**| The ID of the cloud | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject44** | [**InlineObject44**](InlineObject44.md)|  | [optional] 

### Return type

[**InlineResponse20023**](InlineResponse20023.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateCloudLogo

> Model200Success updateCloudLogo(id, opts)

Update Logo For Cloud

Use this command to update the logo and dark logo images for a cloud. This endpoint expects multipart form data as the request format, not JSON. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'logo': "/path/to/file", // File | Logo File png,jpg,svg
  'darkLogo': "/path/to/file" // File | Dark Logo File png,jpg,svg
};
apiInstance.updateCloudLogo(id, opts, (error, data, response) => {
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
 **logo** | **File**| Logo File png,jpg,svg | [optional] 
 **darkLogo** | **File**| Dark Logo File png,jpg,svg | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## updateCloudResourcePool

> InlineResponse20024 updateCloudResourcePool(zoneId, id, opts)

Updates a Specified Resource Pool for Specified Cloud

Updates a resource pool for specified cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let zoneId = 7; // Number | The ID of the cloud
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject46': new MorpheusApi.InlineObject46() // InlineObject46 | 
};
apiInstance.updateCloudResourcePool(zoneId, id, opts, (error, data, response) => {
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
 **zoneId** | **Number**| The ID of the cloud | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject46** | [**InlineObject46**](InlineObject46.md)|  | [optional] 

### Return type

[**InlineResponse20024**](InlineResponse20024.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateCloudSecurityGroups

> Object updateCloudSecurityGroups(id, opts)

Sets Security Groups for a Cloud

Sets security groups for acloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject49': new MorpheusApi.InlineObject49() // InlineObject49 | 
};
apiInstance.updateCloudSecurityGroups(id, opts, (error, data, response) => {
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
 **inlineObject49** | [**InlineObject49**](InlineObject49.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateClouds

> Object updateClouds(id, opts)

Updates a Cloud

Updates a cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject42': new MorpheusApi.InlineObject42() // InlineObject42 | 
};
apiInstance.updateClouds(id, opts, (error, data, response) => {
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
 **inlineObject42** | [**InlineObject42**](InlineObject42.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateWikiCloud

> Object updateWikiCloud(id, opts)

Update a Cloud Wiki Page

Updates a cloud Wiki page. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CloudsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject274': new MorpheusApi.InlineObject274() // InlineObject274 | 
};
apiInstance.updateWikiCloud(id, opts, (error, data, response) => {
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
 **inlineObject274** | [**InlineObject274**](InlineObject274.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

