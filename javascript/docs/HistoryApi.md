# MorpheusApi.HistoryApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getHistory**](HistoryApi.md#getHistory) | **GET** /api/processes/{id} | Retrieves a Specific Process
[**getHistoryEvent**](HistoryApi.md#getHistoryEvent) | **GET** /api/processes/events/{id} | Retrieves a Specific Process Event
[**listHistory**](HistoryApi.md#listHistory) | **GET** /api/processes | Retrieves Process History



## getHistory

> InlineResponse20048 getHistory(id)

Retrieves a Specific Process

Retrieves a specific process. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HistoryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getHistory(id, (error, data, response) => {
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

[**InlineResponse20048**](InlineResponse20048.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getHistoryEvent

> InlineResponse20049 getHistoryEvent(id)

Retrieves a Specific Process Event

Retrieves a specific process event. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HistoryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getHistoryEvent(id, (error, data, response) => {
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

[**InlineResponse20049**](InlineResponse20049.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listHistory

> Object listHistory(opts)

Retrieves Process History

Retrieves process history for objects 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HistoryApi();
let opts = {
  'instanceId': 94, // Number | The Instance ID for Filtering
  'containerId': 135, // Number | The Container ID for Filtering
  'serverId': 97, // Number | The Server ID for Filtering
  'zoneId': 3, // Number | The Zone ID for Filtering
  'appId': 5, // Number | The App ID for Filtering
  'phrase': "phrase_example" // String | Search phrase for partial matches on message, displayName, output, event.message, event.output or event.error
};
apiInstance.listHistory(opts, (error, data, response) => {
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
 **zoneId** | **Number**| The Zone ID for Filtering | [optional] 
 **appId** | **Number**| The App ID for Filtering | [optional] 
 **phrase** | **String**| Search phrase for partial matches on message, displayName, output, event.message, event.output or event.error | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

