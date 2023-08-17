# MorpheusApi.AlertsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addAlerts**](AlertsApi.md#addAlerts) | **POST** /api/monitoring/alerts | Create a New Alert
[**deleteAlerts**](AlertsApi.md#deleteAlerts) | **DELETE** /api/monitoring/alerts/{id} | Delete a Specific Alert
[**getAlerts**](AlertsApi.md#getAlerts) | **GET** /api/monitoring/alerts/{id} | Get a Specific Alert
[**listAlerts**](AlertsApi.md#listAlerts) | **GET** /api/monitoring/alerts | List All Alerts
[**updateAlerts**](AlertsApi.md#updateAlerts) | **PUT** /api/monitoring/alerts/{id} | Update Alert



## addAlerts

> AddAlerts200Response addAlerts(opts)

Create a New Alert

Create a new monitoring alert.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AlertsApi();
let opts = {
  'addAlertsRequest': new MorpheusApi.AddAlertsRequest() // AddAlertsRequest | 
};
apiInstance.addAlerts(opts, (error, data, response) => {
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
 **addAlertsRequest** | [**AddAlertsRequest**](AddAlertsRequest.md)|  | [optional] 

### Return type

[**AddAlerts200Response**](AddAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteAlerts

> Model200Success deleteAlerts(id)

Delete a Specific Alert

Delete an existing monitoring alert.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AlertsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteAlerts(id, (error, data, response) => {
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


## getAlerts

> GetAlerts200Response getAlerts(id)

Get a Specific Alert

Get details about a specific monitoring alert.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AlertsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getAlerts(id, (error, data, response) => {
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

[**GetAlerts200Response**](GetAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listAlerts

> ListAlerts200Response listAlerts(opts)

List All Alerts

Get a list of monitoring alerts.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AlertsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'lastUpdated': 2019-03-06T17:52:29+0000 // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
};
apiInstance.listAlerts(opts, (error, data, response) => {
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
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 

### Return type

[**ListAlerts200Response**](ListAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateAlerts

> UpdateAlerts200Response updateAlerts(id, opts)

Update Alert

Update an existing monitoring alert.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.AlertsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'updateAlertsRequest': new MorpheusApi.UpdateAlertsRequest() // UpdateAlertsRequest | 
};
apiInstance.updateAlerts(id, opts, (error, data, response) => {
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
 **updateAlertsRequest** | [**UpdateAlertsRequest**](UpdateAlertsRequest.md)|  | [optional] 

### Return type

[**UpdateAlerts200Response**](UpdateAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

