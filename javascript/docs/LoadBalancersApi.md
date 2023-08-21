# MorpheusApi.LoadBalancersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createLoadBalancer**](LoadBalancersApi.md#createLoadBalancer) | **POST** /api/load-balancers | Create a Load Balancer
[**createLoadBalancerMonitor**](LoadBalancersApi.md#createLoadBalancerMonitor) | **POST** /api/load-balancers/{loadBalancerId}/monitors | Create a Load Balancer Monitor
[**createLoadBalancerPool**](LoadBalancersApi.md#createLoadBalancerPool) | **POST** /api/load-balancers/{loadBalancerId}/pools | Create a Load Balancer Pool
[**createLoadBalancerPoolNode**](LoadBalancersApi.md#createLoadBalancerPoolNode) | **POST** /api/load-balancer-pools/{loadBalancerPoolId}/nodes | Create a Load Balancer Pool Node
[**createLoadBalancerProfile**](LoadBalancersApi.md#createLoadBalancerProfile) | **POST** /api/load-balancers/{loadBalancerId}/profiles | Create a Load Balancer Profile
[**createLoadBalancerVirtualServer**](LoadBalancersApi.md#createLoadBalancerVirtualServer) | **POST** /api/load-balancers/{loadBalancerId}/virtual-servers | Create a Load Balancer Virtual Server
[**deleteLoadBalancer**](LoadBalancersApi.md#deleteLoadBalancer) | **DELETE** /api/load-balancers/{loadBalancerId} | Delete a Load Balancer
[**deleteLoadBalancerMonitor**](LoadBalancersApi.md#deleteLoadBalancerMonitor) | **DELETE** /api/load-balancers/{loadBalancerId}/monitors/{id} | Delete a Load Balancer Monitor
[**deleteLoadBalancerPool**](LoadBalancersApi.md#deleteLoadBalancerPool) | **DELETE** /api/load-balancers/{loadBalancerId}/pools/{id} | Delete a Load Balancer Pool
[**deleteLoadBalancerPoolNode**](LoadBalancersApi.md#deleteLoadBalancerPoolNode) | **DELETE** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Delete a Load Balancer Pool Node
[**deleteLoadBalancerProfile**](LoadBalancersApi.md#deleteLoadBalancerProfile) | **DELETE** /api/load-balancers/{loadBalancerId}/profiles/{id} | Delete a Load Balancer Profile
[**deleteLoadBalancerVirtualServer**](LoadBalancersApi.md#deleteLoadBalancerVirtualServer) | **DELETE** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Delete a Load Balancer Virtual Server
[**getLoadBalancer**](LoadBalancersApi.md#getLoadBalancer) | **GET** /api/load-balancers/{loadBalancerId} | Get a Specific Load Balancer
[**getLoadBalancerMonitor**](LoadBalancersApi.md#getLoadBalancerMonitor) | **GET** /api/load-balancers/{loadBalancerId}/monitors/{id} | Get a Specific Load Balancer Monitor
[**getLoadBalancerPool**](LoadBalancersApi.md#getLoadBalancerPool) | **GET** /api/load-balancers/{loadBalancerId}/pools/{id} | Get a Specific Load Balancer Pool
[**getLoadBalancerPoolNode**](LoadBalancersApi.md#getLoadBalancerPoolNode) | **GET** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Get a Specific Load Balancer Pool Node
[**getLoadBalancerProfile**](LoadBalancersApi.md#getLoadBalancerProfile) | **GET** /api/load-balancers/{loadBalancerId}/profiles/{id} | Get a Specific Load Balancer Profile
[**getLoadBalancerType**](LoadBalancersApi.md#getLoadBalancerType) | **GET** /api/load-balancer-types/{id} | Get a Specific Load Balancer Type
[**getLoadBalancerVirtualServer**](LoadBalancersApi.md#getLoadBalancerVirtualServer) | **GET** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Get a Specific Load Balancer Virtual Server
[**listLoadBalancerMonitors**](LoadBalancersApi.md#listLoadBalancerMonitors) | **GET** /api/load-balancers/{loadBalancerId}/monitors | Get All Load Balancer Monitors For Load Balancer
[**listLoadBalancerPoolNodes**](LoadBalancersApi.md#listLoadBalancerPoolNodes) | **GET** /api/load-balancer-pools/{loadBalancerPoolId}/nodes | Get All Load Balancer Pool Nodes For Load Balancer Pool
[**listLoadBalancerPools**](LoadBalancersApi.md#listLoadBalancerPools) | **GET** /api/load-balancers/{loadBalancerId}/pools | Get All Load Balancer Pools For Load Balancer
[**listLoadBalancerProfiles**](LoadBalancersApi.md#listLoadBalancerProfiles) | **GET** /api/load-balancers/{loadBalancerId}/profiles | Get All Load Balancer Profiles For Load Balancer
[**listLoadBalancerTypes**](LoadBalancersApi.md#listLoadBalancerTypes) | **GET** /api/load-balancer-types | Get All Load Balancer Types
[**listLoadBalancerVirtualServers**](LoadBalancersApi.md#listLoadBalancerVirtualServers) | **GET** /api/load-balancers/{loadBalancerId}/virtual-servers | Get All Load Balancer Virtual Servers For Load Balancer
[**listLoadBalancers**](LoadBalancersApi.md#listLoadBalancers) | **GET** /api/load-balancers | Get All Load Balancers
[**refreshLoadBalancer**](LoadBalancersApi.md#refreshLoadBalancer) | **PUT** /api/load-balancers/{loadBalancerId}/refresh | Refresh a Load Balancer
[**updateLoadBalancer**](LoadBalancersApi.md#updateLoadBalancer) | **PUT** /api/load-balancers/{loadBalancerId} | Update a Load Balancer
[**updateLoadBalancerMonitor**](LoadBalancersApi.md#updateLoadBalancerMonitor) | **PUT** /api/load-balancers/{loadBalancerId}/monitors/{id} | Update a Load Balancer Monitor
[**updateLoadBalancerPool**](LoadBalancersApi.md#updateLoadBalancerPool) | **PUT** /api/load-balancers/{loadBalancerId}/pools/{id} | Update a Load Balancer Pool
[**updateLoadBalancerPoolNode**](LoadBalancersApi.md#updateLoadBalancerPoolNode) | **PUT** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Update a Load Balancer Pool Node
[**updateLoadBalancerProfile**](LoadBalancersApi.md#updateLoadBalancerProfile) | **PUT** /api/load-balancers/{loadBalancerId}/profiles/{id} | Update a Load Balancer Profile
[**updateLoadBalancerVirtualServer**](LoadBalancersApi.md#updateLoadBalancerVirtualServer) | **PUT** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Update a Load Balancer Virtual Server



## createLoadBalancer

> InlineResponse20078 createLoadBalancer(opts)

Create a Load Balancer

Available for NSX load balancers only  Use this command to create a load balancer. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let opts = {
  'inlineObject127': new MorpheusApi.InlineObject127() // InlineObject127 | 
};
apiInstance.createLoadBalancer(opts, (error, data, response) => {
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
 **inlineObject127** | [**InlineObject127**](InlineObject127.md)|  | [optional] 

### Return type

[**InlineResponse20078**](InlineResponse20078.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createLoadBalancerMonitor

> Object createLoadBalancerMonitor(loadBalancerId, opts)

Create a Load Balancer Monitor

Use this command to create a load balancer Monitor.  This endpoint allows creating a Load Balancer Monitor. Configuration options vary by Load Balancer Type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let opts = {
  'inlineObject129': new MorpheusApi.InlineObject129() // InlineObject129 | 
};
apiInstance.createLoadBalancerMonitor(loadBalancerId, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **inlineObject129** | [**InlineObject129**](InlineObject129.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createLoadBalancerPool

> Object createLoadBalancerPool(loadBalancerId, opts)

Create a Load Balancer Pool

Use this command to create a load balancer pool.  This endpoint allows creating a Load Balancer Pool. Configuration options vary by Load Balancer Type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let opts = {
  'inlineObject131': new MorpheusApi.InlineObject131() // InlineObject131 | 
};
apiInstance.createLoadBalancerPool(loadBalancerId, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **inlineObject131** | [**InlineObject131**](InlineObject131.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createLoadBalancerPoolNode

> Object createLoadBalancerPoolNode(loadBalancerPoolId, opts)

Create a Load Balancer Pool Node

Use this command to create a load balancer pool node.  This endpoint allows creating a Load Balancer Pool Node. Configuration options vary by Load Balancer Type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerPoolId = 4; // Number | Load Balancer Pool ID
let opts = {
  'inlineObject137': new MorpheusApi.InlineObject137() // InlineObject137 | 
};
apiInstance.createLoadBalancerPoolNode(loadBalancerPoolId, opts, (error, data, response) => {
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
 **loadBalancerPoolId** | **Number**| Load Balancer Pool ID | 
 **inlineObject137** | [**InlineObject137**](InlineObject137.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createLoadBalancerProfile

> Object createLoadBalancerProfile(loadBalancerId, opts)

Create a Load Balancer Profile

Use this command to create a load balancer Profile.  This endpoint allows creating a Load Balancer Profile. Configuration options vary by Load Balancer Type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let opts = {
  'inlineObject133': new MorpheusApi.InlineObject133() // InlineObject133 | 
};
apiInstance.createLoadBalancerProfile(loadBalancerId, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **inlineObject133** | [**InlineObject133**](InlineObject133.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createLoadBalancerVirtualServer

> InlineResponse20082 createLoadBalancerVirtualServer(loadBalancerId, opts)

Create a Load Balancer Virtual Server

Use this command to create a load balancer virtual server.  This endpoint allows creating a Load Balancer Virtual Server. Configuration options vary by Load Balancer Type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let opts = {
  'inlineObject135': new MorpheusApi.InlineObject135() // InlineObject135 | 
};
apiInstance.createLoadBalancerVirtualServer(loadBalancerId, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **inlineObject135** | [**InlineObject135**](InlineObject135.md)|  | [optional] 

### Return type

[**InlineResponse20082**](InlineResponse20082.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteLoadBalancer

> Model200Success deleteLoadBalancer(loadBalancerId)

Delete a Load Balancer

Will delete a Load Balancer from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
apiInstance.deleteLoadBalancer(loadBalancerId, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteLoadBalancerMonitor

> Model200Success deleteLoadBalancerMonitor(loadBalancerId, id)

Delete a Load Balancer Monitor

Will delete a Load Balancer Monitor from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteLoadBalancerMonitor(loadBalancerId, id, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteLoadBalancerPool

> Model200Success deleteLoadBalancerPool(loadBalancerId, id)

Delete a Load Balancer Pool

Will delete a Load Balancer Pool from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteLoadBalancerPool(loadBalancerId, id, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteLoadBalancerPoolNode

> Model200Success deleteLoadBalancerPoolNode(loadBalancerPoolId, id)

Delete a Load Balancer Pool Node

Will delete a Load Balancer Pool Node from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerPoolId = 4; // Number | Load Balancer Pool ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteLoadBalancerPoolNode(loadBalancerPoolId, id, (error, data, response) => {
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
 **loadBalancerPoolId** | **Number**| Load Balancer Pool ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteLoadBalancerProfile

> Model200Success deleteLoadBalancerProfile(loadBalancerId, id)

Delete a Load Balancer Profile

Will delete a Load Balancer Profile from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteLoadBalancerProfile(loadBalancerId, id, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteLoadBalancerVirtualServer

> Model200Success deleteLoadBalancerVirtualServer(loadBalancerId, id)

Delete a Load Balancer Virtual Server

Will delete a Load Balancer Virtual Server from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteLoadBalancerVirtualServer(loadBalancerId, id, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getLoadBalancer

> InlineResponse20078 getLoadBalancer(loadBalancerId)

Get a Specific Load Balancer

This endpoint retrieves a specific Load Balancer.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
apiInstance.getLoadBalancer(loadBalancerId, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 

### Return type

[**InlineResponse20078**](InlineResponse20078.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getLoadBalancerMonitor

> InlineResponse20079 getLoadBalancerMonitor(loadBalancerId, id)

Get a Specific Load Balancer Monitor

This endpoint retrieves a specific Load Balancer Monitor.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getLoadBalancerMonitor(loadBalancerId, id, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20079**](InlineResponse20079.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getLoadBalancerPool

> InlineResponse20080 getLoadBalancerPool(loadBalancerId, id)

Get a Specific Load Balancer Pool

This endpoint retrieves a specific Load Balancer Pool.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getLoadBalancerPool(loadBalancerId, id, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20080**](InlineResponse20080.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getLoadBalancerPoolNode

> InlineResponse20083 getLoadBalancerPoolNode(loadBalancerPoolId, id)

Get a Specific Load Balancer Pool Node

This endpoint retrieves a specific Load Balancer Pool Node.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerPoolId = 4; // Number | Load Balancer Pool ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getLoadBalancerPoolNode(loadBalancerPoolId, id, (error, data, response) => {
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
 **loadBalancerPoolId** | **Number**| Load Balancer Pool ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20083**](InlineResponse20083.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getLoadBalancerProfile

> InlineResponse20081 getLoadBalancerProfile(loadBalancerId, id)

Get a Specific Load Balancer Profile

This endpoint retrieves a specific Load Balancer Profile.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getLoadBalancerProfile(loadBalancerId, id, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20081**](InlineResponse20081.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getLoadBalancerType

> InlineResponse20077 getLoadBalancerType(id)

Get a Specific Load Balancer Type

This endpoint will retrieve a specific load balancer type by id.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getLoadBalancerType(id, (error, data, response) => {
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

[**InlineResponse20077**](InlineResponse20077.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getLoadBalancerVirtualServer

> InlineResponse20082 getLoadBalancerVirtualServer(loadBalancerId, id)

Get a Specific Load Balancer Virtual Server

This endpoint retrieves a specific Load Balancer Virtual Server.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getLoadBalancerVirtualServer(loadBalancerId, id, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20082**](InlineResponse20082.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listLoadBalancerMonitors

> Object listLoadBalancerMonitors(loadBalancerId, opts)

Get All Load Balancer Monitors For Load Balancer

This endpoint retrieves all load balancer monitors associated with a specified load balancer.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listLoadBalancerMonitors(loadBalancerId, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listLoadBalancerPoolNodes

> Object listLoadBalancerPoolNodes(loadBalancerPoolId, opts)

Get All Load Balancer Pool Nodes For Load Balancer Pool

This endpoint retrieves all load balancer pool nodes associated with a specified load balancer pool.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerPoolId = 4; // Number | Load Balancer Pool ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listLoadBalancerPoolNodes(loadBalancerPoolId, opts, (error, data, response) => {
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
 **loadBalancerPoolId** | **Number**| Load Balancer Pool ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listLoadBalancerPools

> Object listLoadBalancerPools(loadBalancerId, opts)

Get All Load Balancer Pools For Load Balancer

This endpoint retrieves all load balancer pools associated with a specified load balancer.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listLoadBalancerPools(loadBalancerId, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listLoadBalancerProfiles

> Object listLoadBalancerProfiles(loadBalancerId, opts)

Get All Load Balancer Profiles For Load Balancer

This endpoint retrieves all load balancer profiles associated with a specified load balancer.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listLoadBalancerProfiles(loadBalancerId, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listLoadBalancerTypes

> Object listLoadBalancerTypes(opts)

Get All Load Balancer Types

This endpoint retrieves all Load Balancer Types.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'optionTypes': false, // Boolean | Pass true to include optionTypes in the response for each entry.
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr // String | If specified will return an exact match on code
};
apiInstance.listLoadBalancerTypes(opts, (error, data, response) => {
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
 **optionTypes** | **Boolean**| Pass true to include optionTypes in the response for each entry. | [optional] [default to false]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **code** | **String**| If specified will return an exact match on code | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listLoadBalancerVirtualServers

> Object listLoadBalancerVirtualServers(loadBalancerId, opts)

Get All Load Balancer Virtual Servers For Load Balancer

This endpoint retrieves load balancer virtual servers associated with a specified load balancer.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'vipName': lb-114, // String | If specified will return an exact match on vipName
  'vipAddress': 192.168.123.44, // String | If specified will return an exact match on vipAddress
  'vipHostname': mylb // String | If specified will return an exact match on vipHostname
};
apiInstance.listLoadBalancerVirtualServers(loadBalancerId, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **vipName** | **String**| If specified will return an exact match on vipName | [optional] 
 **vipAddress** | **String**| If specified will return an exact match on vipAddress | [optional] 
 **vipHostname** | **String**| If specified will return an exact match on vipHostname | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listLoadBalancers

> Object listLoadBalancers(opts)

Get All Load Balancers

This endpoint retrieves all load balancers associated with the account.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listLoadBalancers(opts, (error, data, response) => {
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
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## refreshLoadBalancer

> Object refreshLoadBalancer(loadBalancerId)

Refresh a Load Balancer

Will refresh a Load Balancer.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
apiInstance.refreshLoadBalancer(loadBalancerId, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateLoadBalancer

> InlineResponse20078 updateLoadBalancer(loadBalancerId, opts)

Update a Load Balancer

Available for NSX load balancers only  Use this command to update an existing load balancer. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let opts = {
  'inlineObject128': new MorpheusApi.InlineObject128() // InlineObject128 | 
};
apiInstance.updateLoadBalancer(loadBalancerId, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **inlineObject128** | [**InlineObject128**](InlineObject128.md)|  | [optional] 

### Return type

[**InlineResponse20078**](InlineResponse20078.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateLoadBalancerMonitor

> Object updateLoadBalancerMonitor(loadBalancerId, id, opts)

Update a Load Balancer Monitor

Use this command to update an existing load balancer monitor.  This endpoint allows updating a Load Balancer Monitor. Configuration options vary by Load Balancer Type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject130': new MorpheusApi.InlineObject130() // InlineObject130 | 
};
apiInstance.updateLoadBalancerMonitor(loadBalancerId, id, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject130** | [**InlineObject130**](InlineObject130.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateLoadBalancerPool

> Object updateLoadBalancerPool(loadBalancerId, id, opts)

Update a Load Balancer Pool

Use this command to update an existing load balancer pool.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject132': new MorpheusApi.InlineObject132() // InlineObject132 | 
};
apiInstance.updateLoadBalancerPool(loadBalancerId, id, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject132** | [**InlineObject132**](InlineObject132.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateLoadBalancerPoolNode

> Object updateLoadBalancerPoolNode(loadBalancerPoolId, id, opts)

Update a Load Balancer Pool Node

Use this command to update an existing load balancer pool node.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerPoolId = 4; // Number | Load Balancer Pool ID
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject138': new MorpheusApi.InlineObject138() // InlineObject138 | 
};
apiInstance.updateLoadBalancerPoolNode(loadBalancerPoolId, id, opts, (error, data, response) => {
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
 **loadBalancerPoolId** | **Number**| Load Balancer Pool ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject138** | [**InlineObject138**](InlineObject138.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateLoadBalancerProfile

> Object updateLoadBalancerProfile(loadBalancerId, id, opts)

Update a Load Balancer Profile

Use this command to update an existing load balancer Profile.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject134': new MorpheusApi.InlineObject134() // InlineObject134 | 
};
apiInstance.updateLoadBalancerProfile(loadBalancerId, id, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject134** | [**InlineObject134**](InlineObject134.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateLoadBalancerVirtualServer

> InlineResponse20082 updateLoadBalancerVirtualServer(loadBalancerId, id, opts)

Update a Load Balancer Virtual Server

Use this command to update an existing load balancer virtual server.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LoadBalancersApi();
let loadBalancerId = 4; // Number | Load Balancer ID
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject136': new MorpheusApi.InlineObject136() // InlineObject136 | 
};
apiInstance.updateLoadBalancerVirtualServer(loadBalancerId, id, opts, (error, data, response) => {
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
 **loadBalancerId** | **Number**| Load Balancer ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject136** | [**InlineObject136**](InlineObject136.md)|  | [optional] 

### Return type

[**InlineResponse20082**](InlineResponse20082.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

