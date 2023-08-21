# MorpheusApi.LogsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listLogs**](LogsApi.md#listLogs) | **GET** /api/logs | Retrieves Logs



## listLogs

> Log listLogs(opts)

Retrieves Logs

Retrieves logs based on filters provided. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LogsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'order': "'asc'", // String | Sort direction, use 'desc' to reverse sort
  'query': "query_example", // String | Alias for phrase
  'message': "message_example", // String | Filter by message
  'sourceType': "sourceType_example", // String | Filter by source type
  'typeCode': "typeCode_example", // String | Filter by code type
  'objectId': 15, // Number | Filter by objectId
  'token': "token_example", // String | Filter by token
  'level': ERROR, // String | Filter by log level. Multiple values can be passed pipe delimited.
  'startMs': 1657309873, // Number | Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified.
  'endMs': 1657309873, // Number | Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified.
  'startDateTime': 2019-03-06T17:52:29+0000, // Date | Start Date timestamp (ISO 8601)
  'endDateTime': 2019-03-06T17:52:29+0000, // Date | End Date timestamp (ISO 8601)
  'containers': 135, // Number | The Container ID(s) for filtering. Accepts multiple values.
  'servers': 135, // Number | The Server ID(s) for filtering. Accepts multiple values.
  'clusterId': 135 // Number | The Cluster ID(s) for filtering. Accepts multiple values.
};
apiInstance.listLogs(opts, (error, data, response) => {
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
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **query** | **String**| Alias for phrase | [optional] 
 **message** | **String**| Filter by message | [optional] 
 **sourceType** | **String**| Filter by source type | [optional] 
 **typeCode** | **String**| Filter by code type | [optional] 
 **objectId** | **Number**| Filter by objectId | [optional] 
 **token** | **String**| Filter by token | [optional] 
 **level** | **String**| Filter by log level. Multiple values can be passed pipe delimited. | [optional] 
 **startMs** | **Number**| Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified. | [optional] 
 **endMs** | **Number**| Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified. | [optional] 
 **startDateTime** | **Date**| Start Date timestamp (ISO 8601) | [optional] 
 **endDateTime** | **Date**| End Date timestamp (ISO 8601) | [optional] 
 **containers** | **Number**| The Container ID(s) for filtering. Accepts multiple values. | [optional] 
 **servers** | **Number**| The Server ID(s) for filtering. Accepts multiple values. | [optional] 
 **clusterId** | **Number**| The Cluster ID(s) for filtering. Accepts multiple values. | [optional] 

### Return type

[**Log**](Log.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

