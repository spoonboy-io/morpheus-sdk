# MorpheusApi.ServicePlansApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**activateServicePlans**](ServicePlansApi.md#activateServicePlans) | **PUT** /api/service-plans/{id}/activate | Activates a Service Plan
[**addServicePlans**](ServicePlansApi.md#addServicePlans) | **POST** /api/service-plans | Creates a Service Plan
[**deactivateServicePlans**](ServicePlansApi.md#deactivateServicePlans) | **PUT** /api/service-plans/{id}/deactivate | Deactivates a Service Plan
[**getServicePlans**](ServicePlansApi.md#getServicePlans) | **GET** /api/service-plans/{id} | Retrieves a Specific Service Plan
[**listServicePlans**](ServicePlansApi.md#listServicePlans) | **GET** /api/service-plans | Retrieves all Service Plans
[**removeServicePlans**](ServicePlansApi.md#removeServicePlans) | **DELETE** /api/service-plans/{id} | Deletes a Service Plan
[**updateServicePlans**](ServicePlansApi.md#updateServicePlans) | **PUT** /api/service-plans/{id} | Updates a Service Plan



## activateServicePlans

> Model200Success activateServicePlans(id)

Activates a Service Plan

Activates a service plan. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServicePlansApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.activateServicePlans(id, (error, data, response) => {
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


## addServicePlans

> Object addServicePlans(opts)

Creates a Service Plan

Creates a service plan. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServicePlansApi();
let opts = {
  'inlineObject228': new MorpheusApi.InlineObject228() // InlineObject228 | 
};
apiInstance.addServicePlans(opts, (error, data, response) => {
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
 **inlineObject228** | [**InlineObject228**](InlineObject228.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deactivateServicePlans

> Model200Success deactivateServicePlans(id)

Deactivates a Service Plan

Deactivates a service plan. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServicePlansApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deactivateServicePlans(id, (error, data, response) => {
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


## getServicePlans

> InlineResponse200142 getServicePlans(id)

Retrieves a Specific Service Plan

Retrieves a specific service plan. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServicePlansApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getServicePlans(id, (error, data, response) => {
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

[**InlineResponse200142**](InlineResponse200142.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listServicePlans

> Object listServicePlans(opts)

Retrieves all Service Plans

Retrieves all service plans. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServicePlansApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'includeZones': false, // Boolean | Indicates including zones in the service plan response payload
  'provisionTypeId': 22, // Number | Provision type filter, restricts query to only load service plans of specified provision type
  'includeInactive': true // Boolean | If true, include inactive prices in the results
};
apiInstance.listServicePlans(opts, (error, data, response) => {
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
 **includeZones** | **Boolean**| Indicates including zones in the service plan response payload | [optional] [default to false]
 **provisionTypeId** | **Number**| Provision type filter, restricts query to only load service plans of specified provision type | [optional] 
 **includeInactive** | **Boolean**| If true, include inactive prices in the results | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeServicePlans

> Model200Success removeServicePlans(id)

Deletes a Service Plan

Deletes a specified service plan. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServicePlansApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeServicePlans(id, (error, data, response) => {
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


## updateServicePlans

> Object updateServicePlans(id, opts)

Updates a Service Plan

Updates a service plan. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ServicePlansApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject229': new MorpheusApi.InlineObject229() // InlineObject229 | 
};
apiInstance.updateServicePlans(id, opts, (error, data, response) => {
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
 **inlineObject229** | [**InlineObject229**](InlineObject229.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

