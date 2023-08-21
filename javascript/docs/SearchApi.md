# MorpheusApi.SearchApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**search**](SearchApi.md#search) | **GET** /api/search | Provides global search for all types of objects



## search

> Search search(opts)

Provides global search for all types of objects

This endpoint provides global search for all types of objects, with the results sorted in order of relevance.  The &#x60;phrase&#x60; parameter can be specified as part of the URL or as a query parameter. If phrase is not specified, then 0 results (hits) will be returned. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SearchApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'query': "query_example" // String | Alias for phrase
};
apiInstance.search(opts, (error, data, response) => {
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
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **query** | **String**| Alias for phrase | [optional] 

### Return type

[**Search**](Search.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

