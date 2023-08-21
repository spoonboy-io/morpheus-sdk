# MorpheusApi.SecurityGroupsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addSecurityGroupLocations**](SecurityGroupsApi.md#addSecurityGroupLocations) | **POST** /api/security-groups/{id}/locations | Creates a Security Group Location
[**addSecurityGroupRules**](SecurityGroupsApi.md#addSecurityGroupRules) | **POST** /api/security-groups/{id}/rules | Creates a Security Group Rule
[**addSecurityGroups**](SecurityGroupsApi.md#addSecurityGroups) | **POST** /api/security-groups | Creates a Security Group
[**getSecurityGroupRules**](SecurityGroupsApi.md#getSecurityGroupRules) | **GET** /api/security-groups/{id}/rules/{sgId} | Retrieves a Specific Security Group Rule
[**getSecurityGroups**](SecurityGroupsApi.md#getSecurityGroups) | **GET** /api/security-groups/{id} | Retrieves a Specific Security Group
[**listSecurityGroupRules**](SecurityGroupsApi.md#listSecurityGroupRules) | **GET** /api/security-groups/{id}/rules | Retrieves all Security Group Rules
[**listSecurityGroups**](SecurityGroupsApi.md#listSecurityGroups) | **GET** /api/security-groups | Retrieves all Security Groups
[**removeSecurityGroupLocations**](SecurityGroupsApi.md#removeSecurityGroupLocations) | **DELETE** /api/security-groups/{id}/locations/{locationId} | Deletes a Security Group Location
[**removeSecurityGroupRules**](SecurityGroupsApi.md#removeSecurityGroupRules) | **DELETE** /api/security-groups/{id}/rules/{sgId} | Deletes a Security Group Rule
[**removeSecurityGroups**](SecurityGroupsApi.md#removeSecurityGroups) | **DELETE** /api/security-groups/{id} | Deletes a Security Group
[**updateSecurityGroupRules**](SecurityGroupsApi.md#updateSecurityGroupRules) | **PUT** /api/security-groups/{id}/rules/{sgId} | Updates a Security Group Rule
[**updateSecurityGroups**](SecurityGroupsApi.md#updateSecurityGroups) | **PUT** /api/security-groups/{id} | Updating a Security Group



## addSecurityGroupLocations

> Object addSecurityGroupLocations(id, opts)

Creates a Security Group Location

Creates a security group location. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityGroupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject215': new MorpheusApi.InlineObject215() // InlineObject215 | 
};
apiInstance.addSecurityGroupLocations(id, opts, (error, data, response) => {
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
 **inlineObject215** | [**InlineObject215**](InlineObject215.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addSecurityGroupRules

> Object addSecurityGroupRules(id, opts)

Creates a Security Group Rule

Creates a security group rule on specified security group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityGroupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject216': new MorpheusApi.InlineObject216() // InlineObject216 | 
};
apiInstance.addSecurityGroupRules(id, opts, (error, data, response) => {
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
 **inlineObject216** | [**InlineObject216**](InlineObject216.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addSecurityGroups

> InlineResponse200133 addSecurityGroups(opts)

Creates a Security Group

Creates a security group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityGroupsApi();
let opts = {
  'inlineObject213': new MorpheusApi.InlineObject213() // InlineObject213 | 
};
apiInstance.addSecurityGroups(opts, (error, data, response) => {
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
 **inlineObject213** | [**InlineObject213**](InlineObject213.md)|  | [optional] 

### Return type

[**InlineResponse200133**](InlineResponse200133.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getSecurityGroupRules

> InlineResponse200135 getSecurityGroupRules(id, sgId)

Retrieves a Specific Security Group Rule

Retrieves a specific security group rule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityGroupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let sgId = 2352; // Number | Morpheus ID of the security group rule being referenced
apiInstance.getSecurityGroupRules(id, sgId, (error, data, response) => {
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
 **sgId** | **Number**| Morpheus ID of the security group rule being referenced | 

### Return type

[**InlineResponse200135**](InlineResponse200135.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getSecurityGroups

> InlineResponse200134 getSecurityGroups(id)

Retrieves a Specific Security Group

Retrieves a specific security group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityGroupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getSecurityGroups(id, (error, data, response) => {
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

[**InlineResponse200134**](InlineResponse200134.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listSecurityGroupRules

> Object listSecurityGroupRules(id, opts)

Retrieves all Security Group Rules

Retrieves all security group rules for specified security group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityGroupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listSecurityGroupRules(id, opts, (error, data, response) => {
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
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listSecurityGroups

> Object listSecurityGroups(opts)

Retrieves all Security Groups

This endpoint retrieves all security groups and their JSON encoded configuration attributes. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityGroupsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listSecurityGroups(opts, (error, data, response) => {
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

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeSecurityGroupLocations

> Model200Success removeSecurityGroupLocations(id, locationId)

Deletes a Security Group Location

Deletes a security group location. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityGroupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let locationId = 2; // Number | The ID of the location
apiInstance.removeSecurityGroupLocations(id, locationId, (error, data, response) => {
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
 **locationId** | **Number**| The ID of the location | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeSecurityGroupRules

> Model200Success removeSecurityGroupRules(id, sgId)

Deletes a Security Group Rule

Deletes a security group rule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityGroupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let sgId = 2352; // Number | Morpheus ID of the security group rule being referenced
apiInstance.removeSecurityGroupRules(id, sgId, (error, data, response) => {
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
 **sgId** | **Number**| Morpheus ID of the security group rule being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeSecurityGroups

> Model200Success removeSecurityGroups(id)

Deletes a Security Group

Deletes a specified security group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityGroupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeSecurityGroups(id, (error, data, response) => {
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


## updateSecurityGroupRules

> Object updateSecurityGroupRules(id, sgId, opts)

Updates a Security Group Rule

Updates a security group rule. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityGroupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let sgId = 2352; // Number | Morpheus ID of the security group rule being referenced
let opts = {
  'inlineObject217': new MorpheusApi.InlineObject217() // InlineObject217 | 
};
apiInstance.updateSecurityGroupRules(id, sgId, opts, (error, data, response) => {
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
 **sgId** | **Number**| Morpheus ID of the security group rule being referenced | 
 **inlineObject217** | [**InlineObject217**](InlineObject217.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateSecurityGroups

> InlineResponse200133 updateSecurityGroups(id, opts)

Updating a Security Group

Updating a Security Group 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.SecurityGroupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject214': new MorpheusApi.InlineObject214() // InlineObject214 | 
};
apiInstance.updateSecurityGroups(id, opts, (error, data, response) => {
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
 **inlineObject214** | [**InlineObject214**](InlineObject214.md)|  | [optional] 

### Return type

[**InlineResponse200133**](InlineResponse200133.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

