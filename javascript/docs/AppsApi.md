# MorpheusApi.AppsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addAppInstance**](AppsApi.md#addAppInstance) | **POST** /api/apps/{id}/add-instance | Add Existing Instance to App
[**addAppUndoDelete**](AppsApi.md#addAppUndoDelete) | **PUT** /api/apps/{id}/cancel-removal | Undo Delete of an App
[**addApps**](AppsApi.md#addApps) | **POST** /api/apps | Create an App
[**applyAppState**](AppsApi.md#applyAppState) | **POST** /api/apps/{id}/apply | Apply State of an App
[**deleteApp**](AppsApi.md#deleteApp) | **DELETE** /api/apps/{id} | Delete an App
[**getApp**](AppsApi.md#getApp) | **GET** /api/apps/{id} | Get a Specific App
[**getAppSecurityGroups**](AppsApi.md#getAppSecurityGroups) | **GET** /api/apps/{id}/security-groups | Get Security Groups for an App
[**getAppState**](AppsApi.md#getAppState) | **GET** /api/apps/{id}/state | Get State of an App
[**listApps**](AppsApi.md#listApps) | **GET** /api/apps | Get All Apps
[**prepareAppApply**](AppsApi.md#prepareAppApply) | **GET** /api/apps/{id}/prepare-apply | Prepare To Apply an App
[**refreshAppState**](AppsApi.md#refreshAppState) | **POST** /api/apps/{id}/refresh | Refresh State of an App
[**removeAppInstance**](AppsApi.md#removeAppInstance) | **POST** /api/apps/{id}/remove-instance | Remove Instance from App
[**setAppSecurityGroups**](AppsApi.md#setAppSecurityGroups) | **POST** /api/apps/{id}/security-groups | Set Security Groups for an App
[**updateApp**](AppsApi.md#updateApp) | **PUT** /api/apps/{id} | Updating an App
[**validateAppState**](AppsApi.md#validateAppState) | **POST** /api/apps/{id}/validate-apply | Validate Apply State for an App



## addAppInstance

> GetApp200Response addAppInstance(id, opts)

Add Existing Instance to App

Add Existing Instance to App

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'addAppInstanceRequest': {$ref=../components/examples/appAddInstance.json} // AddAppInstanceRequest | 
};
apiInstance.addAppInstance(id, opts, (error, data, response) => {
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
 **addAppInstanceRequest** | [**AddAppInstanceRequest**](AddAppInstanceRequest.md)|  | [optional] 

### Return type

[**GetApp200Response**](GetApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addAppUndoDelete

> GetApp200Response addAppUndoDelete(id)

Undo Delete of an App

This operation will undo the delete of an app that is pending removal.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.addAppUndoDelete(id, (error, data, response) => {
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

[**GetApp200Response**](GetApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## addApps

> AddApps200Response addApps(opts)

Create an App

Create an App

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let opts = {
  'appCreate': {$ref=../components/examples/appRequest.json} // AppCreate | 
};
apiInstance.addApps(opts, (error, data, response) => {
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
 **appCreate** | [**AppCreate**](AppCreate.md)|  | [optional] 

### Return type

[**AddApps200Response**](AddApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## applyAppState

> Model200Success applyAppState(id, opts)

Apply State of an App

This endpoint provides a way to apply the state of an app. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'applyAppStateRequest': new MorpheusApi.ApplyAppStateRequest() // ApplyAppStateRequest | 
};
apiInstance.applyAppState(id, opts, (error, data, response) => {
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
 **applyAppStateRequest** | [**ApplyAppStateRequest**](ApplyAppStateRequest.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteApp

> Model200Success deleteApp(id, opts)

Delete an App

Will delete an app. Use removeInstances&#x3D;on to also delete the instances in the app and all associated monitors and backups.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'removeInstances': on, // String | Remove Instances
  'preserveVolumes': on, // String | Preserve Volumes
  'keepBackups': on, // String | Preserve copy of backups
  'releaseFloatingIps': off, // String | Release Floating IPs
  'releaseEIPs': off, // String | Alias for releaseFloatingIps
  'force': on // String | Force Delete
};
apiInstance.deleteApp(id, opts, (error, data, response) => {
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
 **removeInstances** | **String**| Remove Instances | [optional] [default to &#39;off&#39;]
 **preserveVolumes** | **String**| Preserve Volumes | [optional] [default to &#39;off&#39;]
 **keepBackups** | **String**| Preserve copy of backups | [optional] [default to &#39;off&#39;]
 **releaseFloatingIps** | **String**| Release Floating IPs | [optional] [default to &#39;on&#39;]
 **releaseEIPs** | **String**| Alias for releaseFloatingIps | [optional] [default to &#39;on&#39;]
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getApp

> GetApp200Response getApp(id)

Get a Specific App

This endpoint retrieves a specific app.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getApp(id, (error, data, response) => {
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

[**GetApp200Response**](GetApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getAppSecurityGroups

> GetAppSecurityGroups200Response getAppSecurityGroups(id)

Get Security Groups for an App

This returns a list of all of the security groups applied to an app and whether the firewall is enabled.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getAppSecurityGroups(id, (error, data, response) => {
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

[**GetAppSecurityGroups200Response**](GetAppSecurityGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getAppState

> GetAppState200Response getAppState(id)

Get State of an App

This endpoint provides a way to view the state of an app. The response includes output and resource planning information from the template provider software such as Terraform. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getAppState(id, (error, data, response) => {
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

[**GetAppState200Response**](GetAppState200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listApps

> ListApps200Response listApps(opts)

Get All Apps

This endpoint retrieves a paginated list of apps. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'createdBy': 63, // Number | The User ID for Filtering
  'showDeleted': true, // Boolean | If true, includes instances in pending removal status.
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listApps(opts, (error, data, response) => {
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
 **createdBy** | **Number**| The User ID for Filtering | [optional] 
 **showDeleted** | **Boolean**| If true, includes instances in pending removal status. | [optional] [default to false]
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

[**ListApps200Response**](ListApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## prepareAppApply

> PrepareAppApply200Response prepareAppApply(id)

Prepare To Apply an App

This endpoint provides a way to view the current app configuration and templateParameter variables available to apply. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.prepareAppApply(id, (error, data, response) => {
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

[**PrepareAppApply200Response**](PrepareAppApply200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## refreshAppState

> Model200Success refreshAppState(id, opts)

Refresh State of an App

This endpoint provides a way to refresh the state of an app. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.   

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'body': {key: null} // Object | 
};
apiInstance.refreshAppState(id, opts, (error, data, response) => {
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
 **body** | **Object**|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## removeAppInstance

> GetApp200Response removeAppInstance(id, opts)

Remove Instance from App

Remove Instance from App

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'removeAppInstanceRequest': {$ref=../components/examples/appRemoveInstance.json} // RemoveAppInstanceRequest | 
};
apiInstance.removeAppInstance(id, opts, (error, data, response) => {
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
 **removeAppInstanceRequest** | [**RemoveAppInstanceRequest**](RemoveAppInstanceRequest.md)|  | [optional] 

### Return type

[**GetApp200Response**](GetApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## setAppSecurityGroups

> SetAppSecurityGroups200Response setAppSecurityGroups(id, opts)

Set Security Groups for an App

Set Security Groups for an App

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'setAppSecurityGroupsRequest': {$ref=../components/examples/appUpdateSecurityGroups.json} // SetAppSecurityGroupsRequest | 
};
apiInstance.setAppSecurityGroups(id, opts, (error, data, response) => {
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
 **setAppSecurityGroupsRequest** | [**SetAppSecurityGroupsRequest**](SetAppSecurityGroupsRequest.md)|  | [optional] 

### Return type

[**SetAppSecurityGroups200Response**](SetAppSecurityGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateApp

> GetApp200Response updateApp(id, opts)

Updating an App

This endpoint provides updating of some basic app settings.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'appUpdate': {$ref=../components/examples/appUpdate.json} // AppUpdate | 
};
apiInstance.updateApp(id, opts, (error, data, response) => {
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
 **appUpdate** | [**AppUpdate**](AppUpdate.md)|  | [optional] 

### Return type

[**GetApp200Response**](GetApp200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## validateAppState

> ValidateAppState200Response validateAppState(id, opts)

Validate Apply State for an App

This endpoint provides a way to validate app configuration and templateParameter variables before executing the apply. This only validates the configuration to see any planned changes and it does not actually apply the changes. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AppsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'applyAppStateRequest': new MorpheusApi.ApplyAppStateRequest() // ApplyAppStateRequest | 
};
apiInstance.validateAppState(id, opts, (error, data, response) => {
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
 **applyAppStateRequest** | [**ApplyAppStateRequest**](ApplyAppStateRequest.md)|  | [optional] 

### Return type

[**ValidateAppState200Response**](ValidateAppState200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

