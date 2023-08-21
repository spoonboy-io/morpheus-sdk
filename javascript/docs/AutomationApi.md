# MorpheusApi.AutomationApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addExecuteSchedules**](AutomationApi.md#addExecuteSchedules) | **POST** /api/execute-schedules | Creates a Execute Schedule
[**addPowerScheduleInstances**](AutomationApi.md#addPowerScheduleInstances) | **PUT** /api/power-schedules/{id}/add-instances | Add Instances to a Power Schedule
[**addPowerScheduleServers**](AutomationApi.md#addPowerScheduleServers) | **PUT** /api/power-schedules/{id}/add-servers | Add Servers to a Power Schedule
[**addPowerSchedules**](AutomationApi.md#addPowerSchedules) | **POST** /api/power-schedules | Creates a Power Schedule
[**addScaleThresholds**](AutomationApi.md#addScaleThresholds) | **POST** /api/scale-thresholds | Creates a Scale Threshold
[**addTasks**](AutomationApi.md#addTasks) | **POST** /api/tasks | Creates a Task
[**addWorkflows**](AutomationApi.md#addWorkflows) | **POST** /api/task-sets | Creates a Workflow
[**deletePowerScheduleInstances**](AutomationApi.md#deletePowerScheduleInstances) | **PUT** /api/power-schedules/{id}/remove-instances | Remove Instances from a Power Schedule
[**deletePowerScheduleServers**](AutomationApi.md#deletePowerScheduleServers) | **PUT** /api/power-schedules/{id}/remove-servers | Remove Servers from a Power Schedule
[**executeExecutionRequest**](AutomationApi.md#executeExecutionRequest) | **POST** /api/execution-request/execute | Executes an Execution Request
[**executeTasks**](AutomationApi.md#executeTasks) | **POST** /api/tasks/{id}/execute | Executes a Task
[**executeWorkflows**](AutomationApi.md#executeWorkflows) | **POST** /api/task-sets/{id}/execute | Executes a Workflow
[**getExecuteSchedules**](AutomationApi.md#getExecuteSchedules) | **GET** /api/execute-schedules/{id} | Retrieves a Specific Execute Schedule
[**getExecutionRequest**](AutomationApi.md#getExecutionRequest) | **GET** /api/execution-request/{uniqueId} | Retrieves a Specific Execution Request
[**getPowerSchedules**](AutomationApi.md#getPowerSchedules) | **GET** /api/power-schedules/{id} | Retrieves a Specific Power Schedule
[**getScaleThresholds**](AutomationApi.md#getScaleThresholds) | **GET** /api/scale-thresholds/{id} | Retrieves a Specific Scale Threshold
[**getTaskTypes**](AutomationApi.md#getTaskTypes) | **GET** /api/task-types/{id} | Retrieves a Specific Task Type
[**getTasks**](AutomationApi.md#getTasks) | **GET** /api/tasks/{id} | Retrieves a Specific Task
[**getWorkflows**](AutomationApi.md#getWorkflows) | **GET** /api/task-sets/{id} | Retrieves a Specific Workflow
[**listExecuteSchedules**](AutomationApi.md#listExecuteSchedules) | **GET** /api/execute-schedules | Retrieves all Execute Schedules
[**listPowerSchedules**](AutomationApi.md#listPowerSchedules) | **GET** /api/power-schedules | Retrieves all Power Schedules
[**listScaleThresholds**](AutomationApi.md#listScaleThresholds) | **GET** /api/scale-thresholds | Retrieves all Scale Thresholds
[**listTaskTypes**](AutomationApi.md#listTaskTypes) | **GET** /api/task-types | Retrieves all Task Types
[**listTasks**](AutomationApi.md#listTasks) | **GET** /api/tasks | Retrieves all Tasks
[**listWorkflows**](AutomationApi.md#listWorkflows) | **GET** /api/task-sets | Retrieves all Workflows
[**removeExecuteSchedules**](AutomationApi.md#removeExecuteSchedules) | **DELETE** /api/execute-schedules/{id} | Deletes a Execute Schedule
[**removePowerSchedules**](AutomationApi.md#removePowerSchedules) | **DELETE** /api/power-schedules/{id} | Deletes a Power Schedule
[**removeScaleThresholds**](AutomationApi.md#removeScaleThresholds) | **DELETE** /api/scale-thresholds/{id} | Deletes a Scale Threshold
[**removeTasks**](AutomationApi.md#removeTasks) | **DELETE** /api/tasks/{id} | Deletes a Task
[**removeWorkflows**](AutomationApi.md#removeWorkflows) | **DELETE** /api/task-sets/{id} | Deletes a Workflow
[**updateExecuteSchedules**](AutomationApi.md#updateExecuteSchedules) | **PUT** /api/execute-schedules/{id} | Updates a Execute Schedule
[**updatePowerSchedules**](AutomationApi.md#updatePowerSchedules) | **PUT** /api/power-schedules/{id} | Updates a Power Schedule
[**updateScaleThresholds**](AutomationApi.md#updateScaleThresholds) | **PUT** /api/scale-thresholds/{id} | Updates a Scale Threshold
[**updateTasks**](AutomationApi.md#updateTasks) | **PUT** /api/tasks/{id} | Updates a Task
[**updateWorkflows**](AutomationApi.md#updateWorkflows) | **PUT** /api/task-sets/{id} | Updates a Workflow



## addExecuteSchedules

> Object addExecuteSchedules(opts)

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
  'inlineObject12': new MorpheusApi.InlineObject12() // InlineObject12 | 
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
 **inlineObject12** | [**InlineObject12**](InlineObject12.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addPowerScheduleInstances

> Model200SuccessExpanded addPowerScheduleInstances(id, opts)

Add Instances to a Power Schedule

Add Instances to a Power Schedule. 

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
  'inlineObject192': new MorpheusApi.InlineObject192() // InlineObject192 | 
};
apiInstance.addPowerScheduleInstances(id, opts, (error, data, response) => {
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
 **inlineObject192** | [**InlineObject192**](InlineObject192.md)|  | [optional] 

### Return type

[**Model200SuccessExpanded**](Model200SuccessExpanded.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addPowerScheduleServers

> Model200SuccessExpanded addPowerScheduleServers(id, opts)

Add Servers to a Power Schedule

Add Servers to a Power Schedule. 

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
  'inlineObject193': new MorpheusApi.InlineObject193() // InlineObject193 | 
};
apiInstance.addPowerScheduleServers(id, opts, (error, data, response) => {
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
 **inlineObject193** | [**InlineObject193**](InlineObject193.md)|  | [optional] 

### Return type

[**Model200SuccessExpanded**](Model200SuccessExpanded.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addPowerSchedules

> Object addPowerSchedules(opts)

Creates a Power Schedule

Creates a power schedule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let opts = {
  'inlineObject190': new MorpheusApi.InlineObject190() // InlineObject190 | 
};
apiInstance.addPowerSchedules(opts, (error, data, response) => {
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
 **inlineObject190** | [**InlineObject190**](InlineObject190.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addScaleThresholds

> Object addScaleThresholds(opts)

Creates a Scale Threshold

Creates a scale threshold. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let opts = {
  'inlineObject211': new MorpheusApi.InlineObject211() // InlineObject211 | 
};
apiInstance.addScaleThresholds(opts, (error, data, response) => {
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
 **inlineObject211** | [**InlineObject211**](InlineObject211.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addTasks

> Object addTasks(opts)

Creates a Task

Creates a task. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let opts = {
  'inlineObject246': new MorpheusApi.InlineObject246() // InlineObject246 | 
};
apiInstance.addTasks(opts, (error, data, response) => {
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
 **inlineObject246** | [**InlineObject246**](InlineObject246.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addWorkflows

> Object addWorkflows(opts)

Creates a Workflow

Creates a workflow. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let opts = {
  'inlineObject249': new MorpheusApi.InlineObject249() // InlineObject249 | 
};
apiInstance.addWorkflows(opts, (error, data, response) => {
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
 **inlineObject249** | [**InlineObject249**](InlineObject249.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deletePowerScheduleInstances

> Model200SuccessExpanded deletePowerScheduleInstances(id, opts)

Remove Instances from a Power Schedule

Remove Instances from a Power Schedule. 

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
  'inlineObject194': new MorpheusApi.InlineObject194() // InlineObject194 | 
};
apiInstance.deletePowerScheduleInstances(id, opts, (error, data, response) => {
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
 **inlineObject194** | [**InlineObject194**](InlineObject194.md)|  | [optional] 

### Return type

[**Model200SuccessExpanded**](Model200SuccessExpanded.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deletePowerScheduleServers

> Model200SuccessExpanded deletePowerScheduleServers(id, opts)

Remove Servers from a Power Schedule

Remove Servers from a Power Schedule. 

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
  'inlineObject195': new MorpheusApi.InlineObject195() // InlineObject195 | 
};
apiInstance.deletePowerScheduleServers(id, opts, (error, data, response) => {
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
 **inlineObject195** | [**InlineObject195**](InlineObject195.md)|  | [optional] 

### Return type

[**Model200SuccessExpanded**](Model200SuccessExpanded.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## executeExecutionRequest

> InlineResponse2008 executeExecutionRequest(opts)

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
  'inlineObject14': new MorpheusApi.InlineObject14() // InlineObject14 | 
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
 **inlineObject14** | [**InlineObject14**](InlineObject14.md)|  | [optional] 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## executeTasks

> Object executeTasks(id, opts)

Executes a Task

Executes a task. 

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
  'inlineObject248': new MorpheusApi.InlineObject248() // InlineObject248 | 
};
apiInstance.executeTasks(id, opts, (error, data, response) => {
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
 **inlineObject248** | [**InlineObject248**](InlineObject248.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## executeWorkflows

> Object executeWorkflows(id, opts)

Executes a Workflow

Executes a workflow. 

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
  'inlineObject251': new MorpheusApi.InlineObject251() // InlineObject251 | 
};
apiInstance.executeWorkflows(id, opts, (error, data, response) => {
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
 **inlineObject251** | [**InlineObject251**](InlineObject251.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getExecuteSchedules

> Object getExecuteSchedules(id)

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

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getExecutionRequest

> Object getExecutionRequest(uniqueId)

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

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getPowerSchedules

> Object getPowerSchedules(id)

Retrieves a Specific Power Schedule

Retrieves a specific power schedule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getPowerSchedules(id, (error, data, response) => {
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


## getScaleThresholds

> Object getScaleThresholds(id)

Retrieves a Specific Scale Threshold

Retrieves a specific scale threshold. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getScaleThresholds(id, (error, data, response) => {
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


## getTaskTypes

> InlineResponse200156 getTaskTypes(id)

Retrieves a Specific Task Type

Retrieves a specific task type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getTaskTypes(id, (error, data, response) => {
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

[**InlineResponse200156**](InlineResponse200156.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getTasks

> Object getTasks(id)

Retrieves a Specific Task

Retrieves a specific task. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getTasks(id, (error, data, response) => {
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


## getWorkflows

> Object getWorkflows(id)

Retrieves a Specific Workflow

Retrieves a specific workflow. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getWorkflows(id, (error, data, response) => {
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


## listExecuteSchedules

> Object listExecuteSchedules(opts)

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
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
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


## listPowerSchedules

> Object listPowerSchedules(opts)

Retrieves all Power Schedules

Retrieves all power schedules. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listPowerSchedules(opts, (error, data, response) => {
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

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listScaleThresholds

> Object listScaleThresholds(opts)

Retrieves all Scale Thresholds

Retrieves all scale thresholds. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listScaleThresholds(opts, (error, data, response) => {
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

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listTaskTypes

> InlineResponse200155 listTaskTypes(opts)

Retrieves all Task Types

A Task Type is a type of automation task. Each type defines its own set of options to be configured for each task. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr // String | If specified will return an exact match on code
};
apiInstance.listTaskTypes(opts, (error, data, response) => {
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **code** | **String**| If specified will return an exact match on code | [optional] 

### Return type

[**InlineResponse200155**](InlineResponse200155.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listTasks

> Object listTasks(opts)

Retrieves all Tasks

Retrieves all tasks. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let opts = {
  'type': "type_example", // String | If specified will return all tasks by `task type` code. Refer to `Task Types` API for up to date listings. 
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listTasks(opts, (error, data, response) => {
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
 **type** | **String**| If specified will return all tasks by &#x60;task type&#x60; code. Refer to &#x60;Task Types&#x60; API for up to date listings.  | [optional] 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listWorkflows

> Object listWorkflows(opts)

Retrieves all Workflows

Retrieves all workflows. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example", // String | Filter by label(s), matches records that contain all of the specified labels
  'type': provision // String | Filter by Workflow Type
};
apiInstance.listWorkflows(opts, (error, data, response) => {
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
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 
 **type** | **String**| Filter by Workflow Type | [optional] 

### Return type

**Object**

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


## removePowerSchedules

> Model200Success removePowerSchedules(id)

Deletes a Power Schedule

Deletes a specified power schedule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removePowerSchedules(id, (error, data, response) => {
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


## removeScaleThresholds

> Model200Success removeScaleThresholds(id)

Deletes a Scale Threshold

Deletes a specified scale threshold. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeScaleThresholds(id, (error, data, response) => {
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


## removeTasks

> Model200Success removeTasks(id)

Deletes a Task

Deletes a specified task. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeTasks(id, (error, data, response) => {
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


## removeWorkflows

> Model200Success removeWorkflows(id)

Deletes a Workflow

Deletes a specified workflow. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AutomationApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeWorkflows(id, (error, data, response) => {
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

> Object updateExecuteSchedules(id, opts)

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
  'inlineObject13': new MorpheusApi.InlineObject13() // InlineObject13 | 
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
 **inlineObject13** | [**InlineObject13**](InlineObject13.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updatePowerSchedules

> Object updatePowerSchedules(id, opts)

Updates a Power Schedule

Updates a power schedule. 

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
  'inlineObject191': new MorpheusApi.InlineObject191() // InlineObject191 | 
};
apiInstance.updatePowerSchedules(id, opts, (error, data, response) => {
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
 **inlineObject191** | [**InlineObject191**](InlineObject191.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateScaleThresholds

> Object updateScaleThresholds(id, opts)

Updates a Scale Threshold

Updates a scale threshold. 

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
  'inlineObject212': new MorpheusApi.InlineObject212() // InlineObject212 | 
};
apiInstance.updateScaleThresholds(id, opts, (error, data, response) => {
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
 **inlineObject212** | [**InlineObject212**](InlineObject212.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateTasks

> Object updateTasks(id, opts)

Updates a Task

Updates a task. 

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
  'inlineObject247': new MorpheusApi.InlineObject247() // InlineObject247 | 
};
apiInstance.updateTasks(id, opts, (error, data, response) => {
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
 **inlineObject247** | [**InlineObject247**](InlineObject247.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateWorkflows

> Object updateWorkflows(id, opts)

Updates a Workflow

Updates a workflow. 

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
  'inlineObject250': new MorpheusApi.InlineObject250() // InlineObject250 | 
};
apiInstance.updateWorkflows(id, opts, (error, data, response) => {
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
 **inlineObject250** | [**InlineObject250**](InlineObject250.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

