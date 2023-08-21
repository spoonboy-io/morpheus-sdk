# MorpheusApi.GuidanceApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**executeGuidances**](GuidanceApi.md#executeGuidances) | **PUT** /api/guidance/{id}/execute | Executes a Specific Guidance Recommendation
[**getGuidanceStats**](GuidanceApi.md#getGuidanceStats) | **GET** /api/guidance/stats | Retrieves Guidance Stats
[**getGuidanceTypes**](GuidanceApi.md#getGuidanceTypes) | **GET** /api/guidance/types | Retrieves Guidance Types
[**getGuidances**](GuidanceApi.md#getGuidances) | **GET** /api/guidance/{id} | Retrieves a Specific Guidance Recommendation
[**ignoreGuidances**](GuidanceApi.md#ignoreGuidances) | **PUT** /api/guidance/{id}/ignore | Ignores a Specific Guidance Recommendation
[**listGuidances**](GuidanceApi.md#listGuidances) | **GET** /api/guidance | Retrieves all Guidance Recommendations



## executeGuidances

> Model200Success executeGuidances(id)

Executes a Specific Guidance Recommendation

Executes a specific guidance recommendation. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.GuidanceApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.executeGuidances(id, (error, data, response) => {
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


## getGuidanceStats

> InlineResponse20045 getGuidanceStats()

Retrieves Guidance Stats

This endpoint retrieves a summary of actionable guidance. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.GuidanceApi();
apiInstance.getGuidanceStats((error, data, response) => {
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

[**InlineResponse20045**](InlineResponse20045.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getGuidanceTypes

> InlineResponse20046 getGuidanceTypes()

Retrieves Guidance Types

This endpoint retrieves all guidance types. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.GuidanceApi();
apiInstance.getGuidanceTypes((error, data, response) => {
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

[**InlineResponse20046**](InlineResponse20046.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getGuidances

> InlineResponse20044 getGuidances(id)

Retrieves a Specific Guidance Recommendation

Retrieves a specific guidance recommendation. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.GuidanceApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getGuidances(id, (error, data, response) => {
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

[**InlineResponse20044**](InlineResponse20044.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## ignoreGuidances

> Model200Success ignoreGuidances(id)

Ignores a Specific Guidance Recommendation

Ignores a specific guidance recommendation. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.GuidanceApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.ignoreGuidances(id, (error, data, response) => {
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


## listGuidances

> Object listGuidances(opts)

Retrieves all Guidance Recommendations

Retrieves all guidance recommendations. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.GuidanceApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'severity': "severity_example", // String | Filter by severity
  'type': "type_example", // String | Filter by Guidance Type.  See `Retrieves Guidance Types` for most up to date list of types.
  'state': "state_example" // String | Filter by state, restricts query to only load discoveries by state
};
apiInstance.listGuidances(opts, (error, data, response) => {
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
 **severity** | **String**| Filter by severity | [optional] 
 **type** | **String**| Filter by Guidance Type.  See &#x60;Retrieves Guidance Types&#x60; for most up to date list of types. | [optional] 
 **state** | **String**| Filter by state, restricts query to only load discoveries by state | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

