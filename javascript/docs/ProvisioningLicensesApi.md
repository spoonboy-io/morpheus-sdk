# MorpheusApi.ProvisioningLicensesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addProvisioningLicense**](ProvisioningLicensesApi.md#addProvisioningLicense) | **POST** /api/provisioning-licenses | Create a License
[**getProvisioningLicense**](ProvisioningLicensesApi.md#getProvisioningLicense) | **GET** /api/provisioning-licenses/{id} | Get a Specific License
[**getProvisioningLicenseReservations**](ProvisioningLicensesApi.md#getProvisioningLicenseReservations) | **GET** /api/provisioning-licenses/{id}/reservations | Get Reservations for Specific License
[**listProvisioningLicenses**](ProvisioningLicensesApi.md#listProvisioningLicenses) | **GET** /api/provisioning-licenses | Get All Licenses
[**removeProvisioningLicense**](ProvisioningLicensesApi.md#removeProvisioningLicense) | **DELETE** /api/provisioning-licenses/{id} | Delete a License
[**updateProvisioningLicense**](ProvisioningLicensesApi.md#updateProvisioningLicense) | **PUT** /api/provisioning-licenses/{id} | Update a License



## addProvisioningLicense

> Model200Success addProvisioningLicense(opts)

Create a License

Use this command to create a new license.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ProvisioningLicensesApi();
let opts = {
  'inlineObject202': new MorpheusApi.InlineObject202() // InlineObject202 | 
};
apiInstance.addProvisioningLicense(opts, (error, data, response) => {
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
 **inlineObject202** | [**InlineObject202**](InlineObject202.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getProvisioningLicense

> InlineResponse200126 getProvisioningLicense(id)

Get a Specific License

This endpoint retrieves a specific license.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ProvisioningLicensesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getProvisioningLicense(id, (error, data, response) => {
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

[**InlineResponse200126**](InlineResponse200126.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getProvisioningLicenseReservations

> InlineResponse200127 getProvisioningLicenseReservations(id)

Get Reservations for Specific License

This endpoint retrieves all reservations for a specific license. Each time a license is applied to a new server, a reservation is created, reducing the available copies for the license.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ProvisioningLicensesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getProvisioningLicenseReservations(id, (error, data, response) => {
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

[**InlineResponse200127**](InlineResponse200127.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listProvisioningLicenses

> Object listProvisioningLicenses(opts)

Get All Licenses

This endpoint retrieves all licenses.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ProvisioningLicensesApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'licenseType': win, // String | If specified will return an exact match on licenseType code
  'licenseVersion': 2012 R2 Standard, // String | If specified will return an exact match on licenseVersion
  'orgName': Acme Motors, // String | If specified will return an exact match on orgName
  'fullName': Bugs Bunny // String | If specified will return an exact match on fullName
};
apiInstance.listProvisioningLicenses(opts, (error, data, response) => {
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
 **licenseType** | **String**| If specified will return an exact match on licenseType code | [optional] 
 **licenseVersion** | **String**| If specified will return an exact match on licenseVersion | [optional] 
 **orgName** | **String**| If specified will return an exact match on orgName | [optional] 
 **fullName** | **String**| If specified will return an exact match on fullName | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeProvisioningLicense

> Model200Success removeProvisioningLicense(id)

Delete a License

Will delete a license.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ProvisioningLicensesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeProvisioningLicense(id, (error, data, response) => {
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


## updateProvisioningLicense

> Model200Success updateProvisioningLicense(id, opts)

Update a License

Use this command to update an existing license.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ProvisioningLicensesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject203': new MorpheusApi.InlineObject203() // InlineObject203 | 
};
apiInstance.updateProvisioningLicense(id, opts, (error, data, response) => {
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
 **inlineObject203** | [**InlineObject203**](InlineObject203.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

