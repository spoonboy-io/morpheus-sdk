# MorpheusApi.IncidentsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addIncident**](IncidentsApi.md#addIncident) | **POST** /api/monitoring/incidents | Create a New Incident
[**deleteIncidents**](IncidentsApi.md#deleteIncidents) | **DELETE** /api/monitoring/incidents/{id} | Close a Specific Incident
[**getIncidents**](IncidentsApi.md#getIncidents) | **GET** /api/monitoring/incidents/{id} | Get a Specific Incident
[**listIncidents**](IncidentsApi.md#listIncidents) | **GET** /api/monitoring/incidents | List All Incidents
[**updateIncidents**](IncidentsApi.md#updateIncidents) | **PUT** /api/monitoring/incidents/{id} | Update Incident
[**updateIncidentsReopen**](IncidentsApi.md#updateIncidentsReopen) | **GET** /api/monitoring/incidents/{id}/reopen | Reopen a Specific Incident
[**updateMuteAllIncidents**](IncidentsApi.md#updateMuteAllIncidents) | **PUT** /api/monitoring/incidents/mute-all | Mute All Incidents
[**updateMuteIncidents**](IncidentsApi.md#updateMuteIncidents) | **PUT** /api/monitoring/incidents/{id}/mute | Mute Incident



## addIncident

> Object addIncident(opts)

Create a New Incident

Create a new incident.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IncidentsApi();
let opts = {
  'inlineObject87': new MorpheusApi.InlineObject87() // InlineObject87 | 
};
apiInstance.addIncident(opts, (error, data, response) => {
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
 **inlineObject87** | [**InlineObject87**](InlineObject87.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteIncidents

> Model200Success deleteIncidents(id)

Close a Specific Incident

Close an existing monitoring incident.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IncidentsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteIncidents(id, (error, data, response) => {
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


## getIncidents

> InlineResponse20055 getIncidents(id)

Get a Specific Incident

Get details about a specific incident.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IncidentsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getIncidents(id, (error, data, response) => {
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

[**InlineResponse20055**](InlineResponse20055.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listIncidents

> Object listIncidents(opts)

List All Incidents

Get a list of monitoring incidents.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IncidentsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'status': running, // String | The instance status for filtering.
  'severity': "severity_example" // String | Filter by severity
};
apiInstance.listIncidents(opts, (error, data, response) => {
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
 **status** | **String**| The instance status for filtering. | [optional] 
 **severity** | **String**| Filter by severity | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateIncidents

> Object updateIncidents(id, opts)

Update Incident

Update an existing incident.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IncidentsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject88': new MorpheusApi.InlineObject88() // InlineObject88 | 
};
apiInstance.updateIncidents(id, opts, (error, data, response) => {
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
 **inlineObject88** | [**InlineObject88**](InlineObject88.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateIncidentsReopen

> SuccessMessage updateIncidentsReopen(id)

Reopen a Specific Incident

Get details about a specific incident.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IncidentsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.updateIncidentsReopen(id, (error, data, response) => {
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


## updateMuteAllIncidents

> Object updateMuteAllIncidents(opts)

Mute All Incidents

Mute all existing incident.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IncidentsApi();
let opts = {
  'inlineObject90': new MorpheusApi.InlineObject90() // InlineObject90 | 
};
apiInstance.updateMuteAllIncidents(opts, (error, data, response) => {
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
 **inlineObject90** | [**InlineObject90**](InlineObject90.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateMuteIncidents

> Object updateMuteIncidents(id, opts)

Mute Incident

Mute an existing incident.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.IncidentsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject89': new MorpheusApi.InlineObject89() // InlineObject89 | 
};
apiInstance.updateMuteIncidents(id, opts, (error, data, response) => {
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
 **inlineObject89** | [**InlineObject89**](InlineObject89.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

