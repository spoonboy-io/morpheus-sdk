# MorpheusApi.AutomationApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addExecuteSchedules**](AutomationApi.md#addExecuteSchedules) | **POST** /api/execute-schedules | Creates a Execute Schedule
[**executeExecutionRequest**](AutomationApi.md#executeExecutionRequest) | **POST** /api/execution-request/execute | Executes an Execution Request
[**getExecuteSchedules**](AutomationApi.md#getExecuteSchedules) | **GET** /api/execute-schedules/{id} | Retrieves a Specific Execute Schedule
[**getExecutionRequest**](AutomationApi.md#getExecutionRequest) | **GET** /api/execution-request/{uniqueId} | Retrieves a Specific Execution Request
[**listExecuteSchedules**](AutomationApi.md#listExecuteSchedules) | **GET** /api/execute-schedules | Retrieves all Execute Schedules
[**removeExecuteSchedules**](AutomationApi.md#removeExecuteSchedules) | **DELETE** /api/execute-schedules/{id} | Deletes a Execute Schedule
[**updateExecuteSchedules**](AutomationApi.md#updateExecuteSchedules) | **PUT** /api/execute-schedules/{id} | Updates a Execute Schedule



## addExecuteSchedules

> AddExecuteSchedules200Response addExecuteSchedules(opts)

Creates a Execute Schedule

Creates a execute schedule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let opts = {
  'addExecuteSchedulesRequest': new MorpheusApi.AddExecuteSchedulesRequest() // AddExecuteSchedulesRequest | 
};
apiInstance.addExecuteSchedules(opts, (error, data, response) => {
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
 **addExecuteSchedulesRequest** | [**AddExecuteSchedulesRequest**](AddExecuteSchedulesRequest.md)|  | [optional] 

### Return type

[**AddExecuteSchedules200Response**](AddExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## executeExecutionRequest

> ExecuteExecutionRequest200Response executeExecutionRequest(opts)

Executes an Execution Request

Provides API interfaces for executing an arbitrary script or command on an instance, container or host. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let opts = {
  'instanceId': 94, // Number | The Instance ID for Filtering
  'containerId': 135, // Number | The Container ID for Filtering
  'serverId': 97, // Number | The Server ID for Filtering
  'executeExecutionRequestRequest': new MorpheusApi.ExecuteExecutionRequestRequest() // ExecuteExecutionRequestRequest | 
};
apiInstance.executeExecutionRequest(opts, (error, data, response) => {
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
 **instanceId** | **Number**| The Instance ID for Filtering | [optional] 
 **containerId** | **Number**| The Container ID for Filtering | [optional] 
 **serverId** | **Number**| The Server ID for Filtering | [optional] 
 **executeExecutionRequestRequest** | [**ExecuteExecutionRequestRequest**](ExecuteExecutionRequestRequest.md)|  | [optional] 

### Return type

[**ExecuteExecutionRequest200Response**](ExecuteExecutionRequest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getExecuteSchedules

> GetExecuteSchedules200Response getExecuteSchedules(id)

Retrieves a Specific Execute Schedule

Retrieves a specific execute schedule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getExecuteSchedules(id, (error, data, response) => {
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

[**GetExecuteSchedules200Response**](GetExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getExecutionRequest

> GetExecutionRequest200Response getExecutionRequest(uniqueId)

Retrieves a Specific Execution Request

Retrieves a specific execution request. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let uniqueId = c56f75d0-a59a-4566-92e3-4dc716c5d076; // String | The Unique ID of the execution request
apiInstance.getExecutionRequest(uniqueId, (error, data, response) => {
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
 **uniqueId** | **String**| The Unique ID of the execution request | 

### Return type

[**GetExecutionRequest200Response**](GetExecutionRequest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listExecuteSchedules

> ListExecuteSchedules200Response listExecuteSchedules(opts)

Retrieves all Execute Schedules

Retrieves all execute schedules. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listExecuteSchedules(opts, (error, data, response) => {
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

### Return type

[**ListExecuteSchedules200Response**](ListExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeExecuteSchedules

> Model200Success removeExecuteSchedules(id)

Deletes a Execute Schedule

Deletes a specified execute schedule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeExecuteSchedules(id, (error, data, response) => {
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


## updateExecuteSchedules

> AddExecuteSchedules200Response updateExecuteSchedules(id, opts)

Updates a Execute Schedule

Updates a execute schedule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'updateExecuteSchedulesRequest': new MorpheusApi.UpdateExecuteSchedulesRequest() // UpdateExecuteSchedulesRequest | 
};
apiInstance.updateExecuteSchedules(id, opts, (error, data, response) => {
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
 **updateExecuteSchedulesRequest** | [**UpdateExecuteSchedulesRequest**](UpdateExecuteSchedulesRequest.md)|  | [optional] 

### Return type

[**AddExecuteSchedules200Response**](AddExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

