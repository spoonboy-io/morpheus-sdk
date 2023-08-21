# MorpheusApi.OptionsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getOptionSourceData**](OptionsApi.md#getOptionSourceData) | **GET** /api/options/{optionSource} | Get Option Source Data
[**listCodeRepositories**](OptionsApi.md#listCodeRepositories) | **GET** /api/options/codeRepositories | Retrieves a list of Code/GIT Repositories
[**listOptionNetworkOptions**](OptionsApi.md#listOptionNetworkOptions) | **GET** /api/options/zoneNetworkOptions | Retrieves network options by zone/cloud
[**listOptionValues**](OptionsApi.md#listOptionValues) | **GET** /api/options/list | Retrieves input option values



## getOptionSourceData

> Object getOptionSourceData(optionSource)

Get Option Source Data

Returns a list of name/value pairs for option-type models. Some option-types depend on input data for proper representation. This typically includes zoneId or siteId for the item being provisioned as request parameters or sometimes previous option type parameters. Each option returned has a &#x60;value&#x60;, which is often the &#x60;id&#x60;, but may be a &#x60;code&#x60; or other attribute. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.OptionsApi();
let optionSource = keyPairs; // String | `optionSource` to be listed
apiInstance.getOptionSourceData(optionSource, (error, data, response) => {
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
 **optionSource** | **String**| &#x60;optionSource&#x60; to be listed | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCodeRepositories

> Model200Success listCodeRepositories(opts)

Retrieves a list of Code/GIT Repositories

Retrieves a list of Code/GIT Repositories 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.OptionsApi();
let opts = {
  'integrationId': 33 // Number | Filter by an integration Id.
};
apiInstance.listCodeRepositories(opts, (error, data, response) => {
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
 **integrationId** | **Number**| Filter by an integration Id. | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listOptionNetworkOptions

> ZoneNetworkOptions listOptionNetworkOptions(opts)

Retrieves network options by zone/cloud

This endpoint can be used to see which network options are available for a given cloud (zoneId) and provision type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.OptionsApi();
let opts = {
  'zoneId': 3, // Number | The Zone ID for Filtering
  'provisionTypeId': 22 // Number | Provision type filter, restricts query to only load service plans of specified provision type
};
apiInstance.listOptionNetworkOptions(opts, (error, data, response) => {
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
 **zoneId** | **Number**| The Zone ID for Filtering | [optional] 
 **provisionTypeId** | **Number**| Provision type filter, restricts query to only load service plans of specified provision type | [optional] 

### Return type

[**ZoneNetworkOptions**](ZoneNetworkOptions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listOptionValues

> Object listOptionValues(optionTypeId, opts)

Retrieves input option values

Retrieves all input option values.  Can be used with parameters to supply dependent input values. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.OptionsApi();
let optionTypeId = 789; // Number | Input or Option Type ID
let opts = {
  'config': {"config.fieldName1":"value1"} // Object | Input parameters are required if the input is dependent on them.  Fields must be prefixed with `config.`
};
apiInstance.listOptionValues(optionTypeId, opts, (error, data, response) => {
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
 **optionTypeId** | **Number**| Input or Option Type ID | 
 **config** | [**Object**](.md)| Input parameters are required if the input is dependent on them.  Fields must be prefixed with &#x60;config.&#x60; | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

