# MorpheusApi.ActivityApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listActivity**](ActivityApi.md#listActivity) | **GET** /api/activity | Retrieves Activity



## listActivity

> ListActivity200Response listActivity(opts)

Retrieves Activity

Retrieves activity. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ActivityApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'order': "'asc'", // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'userId': 6, // Number | Filter by User ID
  'tenantId': 3, // Number | Filter by Tenant ID. Only available to the master account.
  'timeframe': "'month'", // String | Filter by a timeframe (ex - today, yesterday, week, month, 3months)
  'start': new Date("2013-10-20T19:20:30+01:00"), // Date | Filter by activity on or after a date(time). Default is 1 month prior
  'end': new Date("2013-10-20T19:20:30+01:00") // Date | Filter by activity on or before a date(time). Default is current date
};
apiInstance.listActivity(opts, (error, data, response) => {
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
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **userId** | **Number**| Filter by User ID | [optional] 
 **tenantId** | **Number**| Filter by Tenant ID. Only available to the master account. | [optional] 
 **timeframe** | **String**| Filter by a timeframe (ex - today, yesterday, week, month, 3months) | [optional] [default to &#39;month&#39;]
 **start** | **Date**| Filter by activity on or after a date(time). Default is 1 month prior | [optional] 
 **end** | **Date**| Filter by activity on or before a date(time). Default is current date | [optional] 

### Return type

[**ListActivity200Response**](ListActivity200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

