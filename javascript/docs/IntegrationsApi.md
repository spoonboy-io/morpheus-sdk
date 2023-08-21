# MorpheusApi.IntegrationsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addIntegrationSnowObjects**](IntegrationsApi.md#addIntegrationSnowObjects) | **POST** /api/integrations/{id}/objects | Creates an Exposed ServiceNow Catalog Item
[**addIntegrations**](IntegrationsApi.md#addIntegrations) | **POST** /api/integrations | Creates an Integration
[**getIntegrationInventory**](IntegrationsApi.md#getIntegrationInventory) | **GET** /api/integrations/{id}/inventory/{inventoryId} | Get a Specific Integration Inventory
[**getIntegrationObjects**](IntegrationsApi.md#getIntegrationObjects) | **GET** /api/integrations/{id}/objects/{objectId} | Get a Specific ServiceNow Integration Object
[**getIntegrationTypeOptionTypes**](IntegrationsApi.md#getIntegrationTypeOptionTypes) | **GET** /api/integration-types/{id}/option-types | Retrieves a Option Types for a Specific Integration Type
[**getIntegrationTypes**](IntegrationsApi.md#getIntegrationTypes) | **GET** /api/integration-types/{id} | Retrieves a Specific Integration Type
[**getIntegrations**](IntegrationsApi.md#getIntegrations) | **GET** /api/integrations/{id} | Retrieves a Specific Integration
[**listIntegrationInventory**](IntegrationsApi.md#listIntegrationInventory) | **GET** /api/integrations/{id}/inventory | Get All Integration Inventory
[**listIntegrationObjects**](IntegrationsApi.md#listIntegrationObjects) | **GET** /api/integrations/{id}/objects | Get ServiceNow Integration Objects
[**listIntegrationTypes**](IntegrationsApi.md#listIntegrationTypes) | **GET** /api/integration-types | Retrieves all Integration Types
[**listIntegrations**](IntegrationsApi.md#listIntegrations) | **GET** /api/integrations | Retrieves all Integrations
[**refreshIntegrations**](IntegrationsApi.md#refreshIntegrations) | **POST** /api/integrations/{id}/refresh | Refresh an Integration
[**removeIntegrationObjects**](IntegrationsApi.md#removeIntegrationObjects) | **DELETE** /api/integrations/{id}/objects/{objectId} | Deletes a ServiceNow Integration object
[**removeIntegrations**](IntegrationsApi.md#removeIntegrations) | **DELETE** /api/integrations/{id} | Deletes an Integration
[**updateIntegrationInventory**](IntegrationsApi.md#updateIntegrationInventory) | **PUT** /api/integrations/{id}/inventory/{inventoryId} | Updating an Integration Inventory
[**updateIntegrations**](IntegrationsApi.md#updateIntegrations) | **PUT** /api/integrations/{id} | Updates an Integration



## addIntegrationSnowObjects

> Object addIntegrationSnowObjects(id, opts)

Creates an Exposed ServiceNow Catalog Item

This endpoint creates an Exposed Catalog Item. This is an integration object of type &#x60;catalog&#x60; that references a &#x60;Catalog Item Type.&#x60; 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject100': new MorpheusApi.InlineObject100() // InlineObject100 | 
};
apiInstance.addIntegrationSnowObjects(id, opts, (error, data, response) => {
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
 **inlineObject100** | [**InlineObject100**](InlineObject100.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addIntegrations

> Object addIntegrations(opts)

Creates an Integration

Creates an integration. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let opts = {
  'UNKNOWN_BASE_TYPE': new MorpheusApi.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.addIntegrations(opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getIntegrationInventory

> InlineResponse20065 getIntegrationInventory(id, inventoryId)

Get a Specific Integration Inventory

This endpoint retrieves a specific integration inventory. Only certain types of integrations support this operation, such as Ansible Tower. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 9; // Number | Morpheus ID of the Integration
let inventoryId = 1; // Number | Morpheus ID of the Integration Inventory
apiInstance.getIntegrationInventory(id, inventoryId, (error, data, response) => {
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
 **id** | **Number**| Morpheus ID of the Integration | 
 **inventoryId** | **Number**| Morpheus ID of the Integration Inventory | 

### Return type

[**InlineResponse20065**](InlineResponse20065.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getIntegrationObjects

> InlineResponse20064 getIntegrationObjects(id, objectId)

Get a Specific ServiceNow Integration Object

This endpoint retrieves a specific ServiceNow integration object.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let objectId = 122; // Number | Morpheus ID of the Object being created or referenced
apiInstance.getIntegrationObjects(id, objectId, (error, data, response) => {
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
 **objectId** | **Number**| Morpheus ID of the Object being created or referenced | 

### Return type

[**InlineResponse20064**](InlineResponse20064.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getIntegrationTypeOptionTypes

> InlineResponse20062 getIntegrationTypeOptionTypes(id)

Retrieves a Option Types for a Specific Integration Type

This endpoint will retrieve the list of option types for a specific integration type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getIntegrationTypeOptionTypes(id, (error, data, response) => {
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

[**InlineResponse20062**](InlineResponse20062.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getIntegrationTypes

> InlineResponse20061 getIntegrationTypes(id, opts)

Retrieves a Specific Integration Type

Retrieves a specific integration type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'optiontypes': false // Boolean | Pass `true` to include optionTypes in the response for each integration type
};
apiInstance.getIntegrationTypes(id, opts, (error, data, response) => {
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
 **optiontypes** | **Boolean**| Pass &#x60;true&#x60; to include optionTypes in the response for each integration type | [optional] [default to false]

### Return type

[**InlineResponse20061**](InlineResponse20061.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getIntegrations

> Object getIntegrations(id)

Retrieves a Specific Integration

Retrieves a specific integration. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getIntegrations(id, (error, data, response) => {
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


## listIntegrationInventory

> Object listIntegrationInventory(id, opts)

Get All Integration Inventory

This endpoint retrieves a list of inventory for a specific integration. Only certain types of integrations support this operation, such as Ansible Tower. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 9; // Number | Morpheus ID of the Integration
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listIntegrationInventory(id, opts, (error, data, response) => {
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
 **id** | **Number**| Morpheus ID of the Integration | 
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


## listIntegrationObjects

> InlineResponse20063 listIntegrationObjects(id, opts)

Get ServiceNow Integration Objects

This endpoint retrieves a list of exposed objects for a specific ServiceNow integration. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'value': "value_example", // String | The type of data being stored, string or object. The data type depends on the cypher mount being used. Most mounts use string as their data type, but secret uses object by default.  You can store a string instead by passing `type=string`. This means the data value returned by the API will be a string instead of an object. 
  'refId': 3 // Number | If specified will return an exact match on refId
};
apiInstance.listIntegrationObjects(id, opts, (error, data, response) => {
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
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **value** | **String**| The type of data being stored, string or object. The data type depends on the cypher mount being used. Most mounts use string as their data type, but secret uses object by default.  You can store a string instead by passing &#x60;type&#x3D;string&#x60;. This means the data value returned by the API will be a string instead of an object.  | [optional] 
 **refId** | **Number**| If specified will return an exact match on refId | [optional] 

### Return type

[**InlineResponse20063**](InlineResponse20063.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listIntegrationTypes

> Object listIntegrationTypes(opts)

Retrieves all Integration Types

An Integration Type is specific third party software that the appliance can be configured to integrate with, such as Ansible, Chef, Git, ServiceNOW, etc. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'optiontypes': false, // Boolean | Pass `true` to include optionTypes in the response for each integration type
  'description': The desription of my cool object, // String | Filter by description, wildcard may be specified as %. eg. `example-%`
  'category': "category_example", // String | If specified will return an exact match on category
  'creatable': true, // Boolean | Filter by creatable
  'enabled': false, // Boolean | Filter by enabled
  'hasCMDB': true, // Boolean | Filter by hasCMDB
  'hasCMDBDiscovery': true, // Boolean | Filter by hasCMDBDiscovery
  'hasCM': true, // Boolean | Filter by hasCM
  'hasDNS': true, // Boolean | Filter by hasDNS
  'hasApprovals': true, // Boolean | Filter by hasApprovals
  'isPlugin': true // Boolean | Filter by isPlugin
};
apiInstance.listIntegrationTypes(opts, (error, data, response) => {
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
 **code** | **String**| If specified will return an exact match on code | [optional] 
 **optiontypes** | **Boolean**| Pass &#x60;true&#x60; to include optionTypes in the response for each integration type | [optional] [default to false]
 **description** | **String**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional] 
 **category** | **String**| If specified will return an exact match on category | [optional] 
 **creatable** | **Boolean**| Filter by creatable | [optional] 
 **enabled** | **Boolean**| Filter by enabled | [optional] 
 **hasCMDB** | **Boolean**| Filter by hasCMDB | [optional] 
 **hasCMDBDiscovery** | **Boolean**| Filter by hasCMDBDiscovery | [optional] 
 **hasCM** | **Boolean**| Filter by hasCM | [optional] 
 **hasDNS** | **Boolean**| Filter by hasDNS | [optional] 
 **hasApprovals** | **Boolean**| Filter by hasApprovals | [optional] 
 **isPlugin** | **Boolean**| Filter by isPlugin | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listIntegrations

> AnyOfobjectmeta listIntegrations(opts)

Retrieves all Integrations

Retrieves all integrations. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'id': 7, // Number | Morpheus ID of the Object being created or referenced
  'url': https://example.com/testimage.ovf, // String | Download the file from a remote url. This can be used instead of uploading a local file.
  'host': "host_example", // String | Filter by integration Host
  'port': "port_example", // String | Filter by integration Port
  'username': administrator, // String | Username
  'version': 5, // Number | Filter by version number (userVersion)
  'windowsVersion': "windowsVersion_example", // String | Filter by integration Windows Version
  'status': running, // String | The instance status for filtering.
  'objects': false // Boolean | Include `objects=true` to return the Integration Objects for each integration.  Available in api version 5.2.8/5.3.2. 
};
apiInstance.listIntegrations(opts, (error, data, response) => {
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
 **id** | **Number**| Morpheus ID of the Object being created or referenced | [optional] 
 **url** | **String**| Download the file from a remote url. This can be used instead of uploading a local file. | [optional] 
 **host** | **String**| Filter by integration Host | [optional] 
 **port** | **String**| Filter by integration Port | [optional] 
 **username** | **String**| Username | [optional] 
 **version** | **Number**| Filter by version number (userVersion) | [optional] 
 **windowsVersion** | **String**| Filter by integration Windows Version | [optional] 
 **status** | **String**| The instance status for filtering. | [optional] 
 **objects** | **Boolean**| Include &#x60;objects&#x3D;true&#x60; to return the Integration Objects for each integration.  Available in api version 5.2.8/5.3.2.  | [optional] [default to false]

### Return type

[**AnyOfobjectmeta**](AnyOfobjectmeta.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## refreshIntegrations

> Object refreshIntegrations(id)

Refresh an Integration

This endpoint will refresh an existing Integration. Only some types support this and will actually perform an action as a result. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.refreshIntegrations(id, (error, data, response) => {
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


## removeIntegrationObjects

> Model200Success removeIntegrationObjects(id, objectId)

Deletes a ServiceNow Integration object

Deletes a specified ServiceNow integration object. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let objectId = 122; // Number | Morpheus ID of the Object being created or referenced
apiInstance.removeIntegrationObjects(id, objectId, (error, data, response) => {
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
 **objectId** | **Number**| Morpheus ID of the Object being created or referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeIntegrations

> Model200Success removeIntegrations(id)

Deletes an Integration

Deletes a specified integration. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeIntegrations(id, (error, data, response) => {
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


## updateIntegrationInventory

> Object updateIntegrationInventory(id, inventoryId, opts)

Updating an Integration Inventory

This endpoint provides updating of integration inventory

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 9; // Number | Morpheus ID of the Integration
let inventoryId = 1; // Number | Morpheus ID of the Integration Inventory
let opts = {
  'inlineObject101': new MorpheusApi.InlineObject101() // InlineObject101 | 
};
apiInstance.updateIntegrationInventory(id, inventoryId, opts, (error, data, response) => {
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
 **id** | **Number**| Morpheus ID of the Integration | 
 **inventoryId** | **Number**| Morpheus ID of the Integration Inventory | 
 **inlineObject101** | [**InlineObject101**](InlineObject101.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateIntegrations

> Object updateIntegrations(id, opts)

Updates an Integration

Updates an integration. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IntegrationsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'UNKNOWN_BASE_TYPE': new MorpheusApi.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateIntegrations(id, opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

