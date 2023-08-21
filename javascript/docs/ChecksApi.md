# MorpheusApi.ChecksApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCheckApps**](ChecksApi.md#addCheckApps) | **POST** /api/monitoring/apps | Create a New Check App
[**addCheckGroups**](ChecksApi.md#addCheckGroups) | **POST** /api/monitoring/groups | Create a New Check Group
[**addChecks**](ChecksApi.md#addChecks) | **POST** /api/monitoring/checks | Create a New Check
[**deleteCheckApps**](ChecksApi.md#deleteCheckApps) | **DELETE** /api/monitoring/apps/{id} | Delete a Specific Check App
[**deleteCheckGroups**](ChecksApi.md#deleteCheckGroups) | **DELETE** /api/monitoring/groups/{id} | Delete a Specific Check Group
[**deleteChecks**](ChecksApi.md#deleteChecks) | **DELETE** /api/monitoring/checks/{id} | Delete a Specific Check
[**getCheckApps**](ChecksApi.md#getCheckApps) | **GET** /api/monitoring/apps/{id} | Get a Specific Check App
[**getCheckGroups**](ChecksApi.md#getCheckGroups) | **GET** /api/monitoring/groups/{id} | Get a Specific Check Group
[**getCheckTypes**](ChecksApi.md#getCheckTypes) | **GET** /api/monitoring/check-types/{id} | Get a Specific Check Type
[**getChecks**](ChecksApi.md#getChecks) | **GET** /api/monitoring/checks/{id} | Get a Specific Check
[**listCheckApps**](ChecksApi.md#listCheckApps) | **GET** /api/monitoring/apps | List All Check Apps
[**listCheckGroups**](ChecksApi.md#listCheckGroups) | **GET** /api/monitoring/groups | List All Check Groups
[**listCheckTypes**](ChecksApi.md#listCheckTypes) | **GET** /api/monitoring/check-types | List All Check Types
[**listChecks**](ChecksApi.md#listChecks) | **GET** /api/monitoring/checks | List All Checks
[**updateCheckApps**](ChecksApi.md#updateCheckApps) | **PUT** /api/monitoring/apps/{id} | Update Check App
[**updateCheckGroups**](ChecksApi.md#updateCheckGroups) | **PUT** /api/monitoring/groups/{id} | Update Check Group
[**updateChecks**](ChecksApi.md#updateChecks) | **PUT** /api/monitoring/checks/{id} | Updates a Check
[**updateMuteAllCheckApps**](ChecksApi.md#updateMuteAllCheckApps) | **PUT** /api/monitoring/apps/mute-all | Mute All Check Apps
[**updateMuteAllCheckGroups**](ChecksApi.md#updateMuteAllCheckGroups) | **PUT** /api/monitoring/groups/mute-all | Mute All Check Groups
[**updateMuteAllChecks**](ChecksApi.md#updateMuteAllChecks) | **PUT** /api/monitoring/checks/mute-all | Mute All Checks
[**updateMuteCheckApps**](ChecksApi.md#updateMuteCheckApps) | **PUT** /api/monitoring/apps/{id}/mute | Mute Check App
[**updateMuteCheckGroups**](ChecksApi.md#updateMuteCheckGroups) | **PUT** /api/monitoring/groups/{id}/mute | Mute Check Group
[**updateMuteChecks**](ChecksApi.md#updateMuteChecks) | **PUT** /api/monitoring/checks/{id}/mute | Mute Check



## addCheckApps

> Object addCheckApps(opts)

Create a New Check App

Create a new check app.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let opts = {
  'inlineObject29': new MorpheusApi.InlineObject29() // InlineObject29 | 
};
apiInstance.addCheckApps(opts, (error, data, response) => {
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
 **inlineObject29** | [**InlineObject29**](InlineObject29.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addCheckGroups

> Object addCheckGroups(opts)

Create a New Check Group

Create a new check group.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let opts = {
  'inlineObject37': new MorpheusApi.InlineObject37() // InlineObject37 | 
};
apiInstance.addCheckGroups(opts, (error, data, response) => {
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
 **inlineObject37** | [**InlineObject37**](InlineObject37.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addChecks

> Object addChecks(opts)

Create a New Check

Create a new monitoring check.  SSH tunneling options allow the different check types to tunnel to a host via a proxy, and execute checks relative to the proxy. A SSH tunnel can use your account generated public and private key-pairs or SSH password. It is strongly recommended to use a key-pair.  To enable SSH tunneling for a check, add &#x60;tunnelOn&#x60;, &#x60;sshHost&#x60;, &#x60;sshUser&#x60;, and optionally, &#x60;sshPort&#x60; and &#x60;sshPassword&#x60; parameters to any check type config as seen earlier in the Check Types section. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let opts = {
  'inlineObject33': new MorpheusApi.InlineObject33() // InlineObject33 | 
};
apiInstance.addChecks(opts, (error, data, response) => {
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
 **inlineObject33** | [**InlineObject33**](InlineObject33.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteCheckApps

> Model200Success deleteCheckApps(id)

Delete a Specific Check App

Delete an existing monitoring check app.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteCheckApps(id, (error, data, response) => {
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


## deleteCheckGroups

> Model200Success deleteCheckGroups(id)

Delete a Specific Check Group

Delete an existing monitoring check group.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteCheckGroups(id, (error, data, response) => {
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


## deleteChecks

> Model200Success deleteChecks(id)

Delete a Specific Check

Delete an existing monitoring check.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteChecks(id, (error, data, response) => {
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


## getCheckApps

> InlineResponse20016 getCheckApps(id)

Get a Specific Check App

Get details about a specific monitoring check app.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getCheckApps(id, (error, data, response) => {
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

[**InlineResponse20016**](InlineResponse20016.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getCheckGroups

> InlineResponse20019 getCheckGroups(id)

Get a Specific Check Group

Get details about a specific monitoring check group.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getCheckGroups(id, (error, data, response) => {
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

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getCheckTypes

> InlineResponse20018 getCheckTypes(id)

Get a Specific Check Type

Get details about a specific monitoring check type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getCheckTypes(id, (error, data, response) => {
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

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getChecks

> InlineResponse20017 getChecks(id)

Get a Specific Check

Get details about a specific monitoring check.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getChecks(id, (error, data, response) => {
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

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCheckApps

> Object listCheckApps(opts)

List All Check Apps

Get a list of check apps.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'status': running, // String | The instance status for filtering.
  'lastUpdated': 2019-03-06T17:52:29+0000, // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
  'deleted': true // Boolean | If true, only deleted resources or instances in pending removal status are returned.
};
apiInstance.listCheckApps(opts, (error, data, response) => {
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **status** | **String**| The instance status for filtering. | [optional] 
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 
 **deleted** | **Boolean**| If true, only deleted resources or instances in pending removal status are returned. | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCheckGroups

> Object listCheckGroups(opts)

List All Check Groups

Get a list of check groups.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'status': running, // String | The instance status for filtering.
  'lastUpdated': 2019-03-06T17:52:29+0000, // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
  'deleted': true // Boolean | If true, only deleted resources or instances in pending removal status are returned.
};
apiInstance.listCheckGroups(opts, (error, data, response) => {
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **status** | **String**| The instance status for filtering. | [optional] 
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 
 **deleted** | **Boolean**| If true, only deleted resources or instances in pending removal status are returned. | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCheckTypes

> Object listCheckTypes(opts)

List All Check Types

Get a list of monitoring check types.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listCheckTypes(opts, (error, data, response) => {
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listChecks

> Object listChecks(opts)

List All Checks

Get a list of monitoring checks.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'status': healthy, // String | Filter by health status.
  'lastUpdated': 2019-03-06T17:52:29+0000, // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
  'deleted': true // Boolean | If true, only deleted resources or instances in pending removal status are returned.
};
apiInstance.listChecks(opts, (error, data, response) => {
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **status** | **String**| Filter by health status. | [optional] 
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 
 **deleted** | **Boolean**| If true, only deleted resources or instances in pending removal status are returned. | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateCheckApps

> Object updateCheckApps(id, opts)

Update Check App

Update an existing monitoring check app.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject31': new MorpheusApi.InlineObject31() // InlineObject31 | 
};
apiInstance.updateCheckApps(id, opts, (error, data, response) => {
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
 **inlineObject31** | [**InlineObject31**](InlineObject31.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateCheckGroups

> Object updateCheckGroups(id, opts)

Update Check Group

Update an existing monitoring check group.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject38': new MorpheusApi.InlineObject38() // InlineObject38 | 
};
apiInstance.updateCheckGroups(id, opts, (error, data, response) => {
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
 **inlineObject38** | [**InlineObject38**](InlineObject38.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateChecks

> Object updateChecks(id, opts)

Updates a Check

Updates a monitoring check.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject35': new MorpheusApi.InlineObject35() // InlineObject35 | 
};
apiInstance.updateChecks(id, opts, (error, data, response) => {
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
 **inlineObject35** | [**InlineObject35**](InlineObject35.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateMuteAllCheckApps

> Object updateMuteAllCheckApps(opts)

Mute All Check Apps

Mute all existing check apps.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let opts = {
  'inlineObject30': new MorpheusApi.InlineObject30() // InlineObject30 | 
};
apiInstance.updateMuteAllCheckApps(opts, (error, data, response) => {
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
 **inlineObject30** | [**InlineObject30**](InlineObject30.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateMuteAllCheckGroups

> Object updateMuteAllCheckGroups(opts)

Mute All Check Groups

Mute all existing check groups.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let opts = {
  'inlineObject40': new MorpheusApi.InlineObject40() // InlineObject40 | 
};
apiInstance.updateMuteAllCheckGroups(opts, (error, data, response) => {
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
 **inlineObject40** | [**InlineObject40**](InlineObject40.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateMuteAllChecks

> Object updateMuteAllChecks(opts)

Mute All Checks

Mute all existing checks.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let opts = {
  'inlineObject34': new MorpheusApi.InlineObject34() // InlineObject34 | 
};
apiInstance.updateMuteAllChecks(opts, (error, data, response) => {
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
 **inlineObject34** | [**InlineObject34**](InlineObject34.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateMuteCheckApps

> Object updateMuteCheckApps(id, opts)

Mute Check App

Mute an existing check app.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject32': new MorpheusApi.InlineObject32() // InlineObject32 | 
};
apiInstance.updateMuteCheckApps(id, opts, (error, data, response) => {
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
 **inlineObject32** | [**InlineObject32**](InlineObject32.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateMuteCheckGroups

> Object updateMuteCheckGroups(id, opts)

Mute Check Group

Mute an existing check group.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject39': new MorpheusApi.InlineObject39() // InlineObject39 | 
};
apiInstance.updateMuteCheckGroups(id, opts, (error, data, response) => {
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
 **inlineObject39** | [**InlineObject39**](InlineObject39.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateMuteChecks

> Object updateMuteChecks(id, opts)

Mute Check

Mute an existing check.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ChecksApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject36': new MorpheusApi.InlineObject36() // InlineObject36 | 
};
apiInstance.updateMuteChecks(id, opts, (error, data, response) => {
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
 **inlineObject36** | [**InlineObject36**](InlineObject36.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

