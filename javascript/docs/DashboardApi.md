# MorpheusApi.DashboardApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**dashboard**](DashboardApi.md#dashboard) | **GET** /api/dashboard | Overview of Dashboard information



## dashboard

> Dashboard dashboard()

Overview of Dashboard information

This endpoint can be used to view dashboard information about the remote Morpheus appliance. This is an overview and summary of data available to the user that can be used to render a dashboard. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DashboardApi();
apiInstance.dashboard((error, data, response) => {
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

[**Dashboard**](Dashboard.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

