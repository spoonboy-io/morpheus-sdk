# MorpheusApi.SetupApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**setup**](SetupApi.md#setup) | **POST** /api/setup | Setup appliance



## setup

> Setup setup(opts)

Setup appliance

Initialize a freshly installed appliance to create the master tenant and System Admin user.  Authorization is not required.  This operation can only be executed successfully once. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';

let apiInstance = new MorpheusApi.SetupApi();
let opts = {
  'UNKNOWN_BASE_TYPE': {"applianceName":"The Matrix","accountName":"Meta Cortex Corporation","firstName":"Thomas","lastName":"Anderson","email":"tanderson@mccorp.com","username":"tanderson","password":"QnW}cg}8}<~:P9YU"} // UNKNOWN_BASE_TYPE | 
};
apiInstance.setup(opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

[**Setup**](Setup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

