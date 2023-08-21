# MorpheusApi.HealthApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**acknowledgeHealthAlarm**](HealthApi.md#acknowledgeHealthAlarm) | **PUT** /api/health/alarms/{id}/acknowledge | Acknowledge a Health Alarm
[**acknowledgeHealthAlarms**](HealthApi.md#acknowledgeHealthAlarms) | **PUT** /api/health/alarms/acknowledge | Acknowledge Many Health Alarms
[**exportHealthLogs**](HealthApi.md#exportHealthLogs) | **GET** /api/health/logs/export | Export Appliance Health Logs
[**getHealthAlarms**](HealthApi.md#getHealthAlarms) | **GET** /api/health/alarms/{id} | Retrieves a Specific Appliance Health Alarm
[**listHealth**](HealthApi.md#listHealth) | **GET** /api/health | Retrieves Appliance Health
[**listHealthAlarms**](HealthApi.md#listHealthAlarms) | **GET** /api/health/alarms | Retrieves Appliance Health Alarms
[**listHealthLogs**](HealthApi.md#listHealthLogs) | **GET** /api/health/logs | Retrieves Appliance Health Logs
[**ping**](HealthApi.md#ping) | **GET** /api/ping | Basic information about current Morpheus Installation



## acknowledgeHealthAlarm

> Model200Success acknowledgeHealthAlarm(id, opts)

Acknowledge a Health Alarm

Acknowledge a specific health alarm.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HealthApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject81': new MorpheusApi.InlineObject81() // InlineObject81 | 
};
apiInstance.acknowledgeHealthAlarm(id, opts, (error, data, response) => {
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
 **inlineObject81** | [**InlineObject81**](InlineObject81.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## acknowledgeHealthAlarms

> Model200Success acknowledgeHealthAlarms(opts)

Acknowledge Many Health Alarms

Acknowledge health alarms.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HealthApi();
let opts = {
  'inlineObject80': new MorpheusApi.InlineObject80() // InlineObject80 | 
};
apiInstance.acknowledgeHealthAlarms(opts, (error, data, response) => {
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
 **inlineObject80** | [**InlineObject80**](InlineObject80.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## exportHealthLogs

> File exportHealthLogs(opts)

Export Appliance Health Logs

This endpoint downloads the morpheus appliance logs as a file attachment. By default, the most recent 10,000 log entries are returned, with the newest at the end of the file. The format for each log entry is &#x60;timestamp&#x60; &#x60;level&#x60; &#x60;message&#x60;.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HealthApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'acknowledged': false, // Boolean | True or False flag for Acknowledged items
  'startDate': 2019-01-01, // String | Filter by startDate greater than or equal to a specified date
  'endDate': 2020-01-01, // String | Filter by endDate less than or equal to a specified date
  'reverse': true // Boolean | Reverse order of records. This `true` by default when sort and direction are not passed, but `false` by default if either is passed. This means that by default the newest log entries are the bottom of the file. 
};
apiInstance.exportHealthLogs(opts, (error, data, response) => {
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
 **acknowledged** | **Boolean**| True or False flag for Acknowledged items | [optional] 
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 
 **reverse** | **Boolean**| Reverse order of records. This &#x60;true&#x60; by default when sort and direction are not passed, but &#x60;false&#x60; by default if either is passed. This means that by default the newest log entries are the bottom of the file.  | [optional] 

### Return type

**File**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json


## getHealthAlarms

> Object getHealthAlarms(id)

Retrieves a Specific Appliance Health Alarm

This endpoint will retrieve a specific health alarm by ID.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HealthApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getHealthAlarms(id, (error, data, response) => {
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


## listHealth

> Object listHealth()

Retrieves Appliance Health

This endpoint retrieves health info about the appliance such as cpu, memory and database usage. Elasticsearch statistics and queue usage are also returned.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HealthApi();
apiInstance.listHealth((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listHealthAlarms

> Object listHealthAlarms(opts)

Retrieves Appliance Health Alarms

This endpoint retrieves all health alarms, which are Operation notifications from Cloud and other Service Integrations. These alarms are not generated by the appliance, but synced and displayed for visibility. By default only open alarms are returned. Open alarms are those that have not yet been acknowledged.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HealthApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'acknowledged': false // Boolean | True or False flag for Acknowledged items
};
apiInstance.listHealthAlarms(opts, (error, data, response) => {
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
 **acknowledged** | **Boolean**| True or False flag for Acknowledged items | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listHealthLogs

> Object listHealthLogs(opts)

Retrieves Appliance Health Logs

This endpoint retrieves all health logs. These are the logs of the remote appliance itself. These logs show all ui activity and are useful for troubleshooting and auditing. Stack traces are filtered for Morpheus services. Complete stack traces can be found in &#x60;/var/log/morpheus/morpheus-ui/current&#x60;.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HealthApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'acknowledged': false, // Boolean | True or False flag for Acknowledged items
  'startDate': 2019-01-01, // String | Filter by startDate greater than or equal to a specified date
  'endDate': 2020-01-01 // String | Filter by endDate less than or equal to a specified date
};
apiInstance.listHealthLogs(opts, (error, data, response) => {
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
 **acknowledged** | **Boolean**| True or False flag for Acknowledged items | [optional] 
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] 
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## ping

> Ping ping()

Basic information about current Morpheus Installation

This endpoint can be used to check the remote appliance build version and some other basic information.  This is an unsecured endpoint and does not require authorization. However, build version will not be returned unless you are authenticated with a valid access token. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.HealthApi();
apiInstance.ping((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**Ping**](Ping.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

