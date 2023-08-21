# MorpheusApi.ContainersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**cloneImageContainerAction**](ContainersApi.md#cloneImageContainerAction) | **PUT** /api/containers/{id}/clone-image | Clone Specific Container to Image
[**containersAttachFloatingIp**](ContainersApi.md#containersAttachFloatingIp) | **PUT** /api/containers/{id}/attach-floating-ip | Attach Floating IP to Container
[**containersDetachFloatingIp**](ContainersApi.md#containersDetachFloatingIp) | **PUT** /api/containers/{id}/detach-floating-ip | Detach Floating IP from Container
[**ejectContainerAction**](ContainersApi.md#ejectContainerAction) | **PUT** /api/containers/{id}/eject | Eject a Specific Container
[**executeContainerAction**](ContainersApi.md#executeContainerAction) | **PUT** /api/containers/{id}/action | Execute Container Action
[**getContainer**](ContainersApi.md#getContainer) | **GET** /api/containers/{id} | Get a Specific Container
[**getContainerActions**](ContainersApi.md#getContainerActions) | **GET** /api/containers/{id}/actions | List Container Actions
[**importContainerAction**](ContainersApi.md#importContainerAction) | **PUT** /api/containers/{id}/import | Import a Specific Container
[**restartContainerAction**](ContainersApi.md#restartContainerAction) | **PUT** /api/containers/{id}/restart | Restart a Specific Container
[**startContainerAction**](ContainersApi.md#startContainerAction) | **PUT** /api/containers/{id}/start | Start a Specific Container
[**stopContainerAction**](ContainersApi.md#stopContainerAction) | **PUT** /api/containers/{id}/stop | Stop a Specific Container
[**suspendContainerAction**](ContainersApi.md#suspendContainerAction) | **PUT** /api/containers/{id}/suspend | Suspend a Specific Container



## cloneImageContainerAction

> SuccessMessage cloneImageContainerAction(id, opts)

Clone Specific Container to Image

This endpoint clones and converts a Container in its current state to image in the source Cloud and adds Virtual Image Record with metadata matching the source configuration.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ContainersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'instancesCloneImage': {"templateName":"Example Image","zoneFolder":"group-v81"} // InstancesCloneImage | 
};
apiInstance.cloneImageContainerAction(id, opts, (error, data, response) => {
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
 **instancesCloneImage** | [**InstancesCloneImage**](InstancesCloneImage.md)|  | [optional] 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## containersAttachFloatingIp

> Model200Success containersAttachFloatingIp(id, opts)

Attach Floating IP to Container

This endpoint attaches a floating IP to a container (node/VM).

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ContainersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject63': new MorpheusApi.InlineObject63() // InlineObject63 | 
};
apiInstance.containersAttachFloatingIp(id, opts, (error, data, response) => {
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
 **inlineObject63** | [**InlineObject63**](InlineObject63.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## containersDetachFloatingIp

> Model200Success containersDetachFloatingIp(id)

Detach Floating IP from Container

This endpoint detaches a floating IP from a container (node/VM).

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ContainersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.containersDetachFloatingIp(id, (error, data, response) => {
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


## ejectContainerAction

> SuccessMessage ejectContainerAction(id)

Eject a Specific Container

This endpoint ejects attached disk/iso of a specific container.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ContainersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.ejectContainerAction(id, (error, data, response) => {
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

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## executeContainerAction

> SuccessMessage executeContainerAction(id, code)

Execute Container Action

This endpoint executes a container action on specific container.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ContainersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let code = apache-remove-node; // String | Action code to be executed on a specific container.
apiInstance.executeContainerAction(id, code, (error, data, response) => {
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
 **code** | **String**| Action code to be executed on a specific container. | 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getContainer

> InlineResponse20034 getContainer(id)

Get a Specific Container

This endpoint retrieves a specific container.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ContainersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getContainer(id, (error, data, response) => {
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

[**InlineResponse20034**](InlineResponse20034.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getContainerActions

> InlineResponse20035 getContainerActions(id)

List Container Actions

This endpoint lists available actions for specific container.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ContainersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getContainerActions(id, (error, data, response) => {
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

[**InlineResponse20035**](InlineResponse20035.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## importContainerAction

> SuccessMessage importContainerAction(id, opts)

Import a Specific Container

This endpoint clones and exports a Container in its current state to target Storage provider and adds Virtual Image Record with metadata matching the source configuration.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ContainersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject62': new MorpheusApi.InlineObject62() // InlineObject62 | 
};
apiInstance.importContainerAction(id, opts, (error, data, response) => {
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
 **inlineObject62** | [**InlineObject62**](InlineObject62.md)|  | [optional] 

### Return type

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## restartContainerAction

> SuccessMessage restartContainerAction(id)

Restart a Specific Container

This endpoint restarts a specific container.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ContainersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.restartContainerAction(id, (error, data, response) => {
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

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## startContainerAction

> SuccessMessage startContainerAction(id)

Start a Specific Container

This endpoint starts a specific container.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ContainersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.startContainerAction(id, (error, data, response) => {
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

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## stopContainerAction

> SuccessMessage stopContainerAction(id)

Stop a Specific Container

This endpoint stops a specific container.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ContainersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.stopContainerAction(id, (error, data, response) => {
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

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## suspendContainerAction

> SuccessMessage suspendContainerAction(id)

Suspend a Specific Container

This endpoint suspends a specific container.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ContainersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.suspendContainerAction(id, (error, data, response) => {
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

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

