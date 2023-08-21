# MorpheusApi.SecurityScansApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getSecurityScans**](SecurityScansApi.md#getSecurityScans) | **GET** /api/security-scans/{id} | Retrieves a Specific Security Scan
[**listSecurityScans**](SecurityScansApi.md#listSecurityScans) | **GET** /api/security-scans | Retrieves all Security Scans



## getSecurityScans

> Object getSecurityScans(id, opts)

Retrieves a Specific Security Scan

Retrieves a specific security scan. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityScansApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'results': true // Boolean | Include the `results` object in the response under the security scan. This is a potentially very large object containing the raw results of the scan.
};
apiInstance.getSecurityScans(id, opts, (error, data, response) => {
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
 **results** | **Boolean**| Include the &#x60;results&#x60; object in the response under the security scan. This is a potentially very large object containing the raw results of the scan. | [optional] [default to false]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listSecurityScans

> Object listSecurityScans(opts)

Retrieves all Security Scans

Retrieves all security scans. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityScansApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'scanDate'", // String | Sort order, the name of the property to sort by
  'direction': desc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description of security package
  'securityPackageId': 789, // Number | Filter results by security package id(s). This parameter can be passed multiple times to match more than one id.
  'serverId': 97, // Number | The Server ID for Filtering
  'results': true // Boolean | Include the `results` object in the response under each security scan. This is a potentially very large object containing the raw results of the scan.
};
apiInstance.listSecurityScans(opts, (error, data, response) => {
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
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;scanDate&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;desc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description of security package | [optional] 
 **securityPackageId** | **Number**| Filter results by security package id(s). This parameter can be passed multiple times to match more than one id. | [optional] 
 **serverId** | **Number**| The Server ID for Filtering | [optional] 
 **results** | **Boolean**| Include the &#x60;results&#x60; object in the response under each security scan. This is a potentially very large object containing the raw results of the scan. | [optional] [default to false]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

