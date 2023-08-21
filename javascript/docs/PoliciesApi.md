# MorpheusApi.PoliciesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addPolicies**](PoliciesApi.md#addPolicies) | **POST** /api/policies | Creates a Policy
[**addPoliciesCloud**](PoliciesApi.md#addPoliciesCloud) | **POST** /api/zones/{cloudId}/policies | Creates a Policy for a Cloud
[**addPoliciesGroup**](PoliciesApi.md#addPoliciesGroup) | **POST** /api/groups/{groupId}/policies | Creates a Policy for a Group
[**getPolicies**](PoliciesApi.md#getPolicies) | **GET** /api/policies/{id} | Retrieves a Specific Policy
[**getPoliciesCloud**](PoliciesApi.md#getPoliciesCloud) | **GET** /api/zones/{cloudId}/policies/{id} | Retrieves a Specific Policy for a Cloud
[**getPoliciesGroup**](PoliciesApi.md#getPoliciesGroup) | **GET** /api/groups/{groupId}/policies/{id} | Retrieves a Specific Policy for a Group
[**listPolicies**](PoliciesApi.md#listPolicies) | **GET** /api/policies | Retrieves all Policies
[**listPoliciesCloud**](PoliciesApi.md#listPoliciesCloud) | **GET** /api/zones/{cloudId}/policies | Retrieves Policies for a Cloud
[**listPoliciesGroup**](PoliciesApi.md#listPoliciesGroup) | **GET** /api/groups/{groupId}/policies | Retrieves Policies for a Group
[**listPolicyTypes**](PoliciesApi.md#listPolicyTypes) | **GET** /api/policy-types | Retrieves all Policy Types
[**removePolicies**](PoliciesApi.md#removePolicies) | **DELETE** /api/policies/{id} | Deletes a Policy
[**removePoliciesCloud**](PoliciesApi.md#removePoliciesCloud) | **DELETE** /api/zones/{cloudId}/policies/{id} | Deletes a Policy for a Cloud
[**removePoliciesGroup**](PoliciesApi.md#removePoliciesGroup) | **DELETE** /api/groups/{groupId}/policies/{id} | Deletes a Policy for a Group
[**updatePolicies**](PoliciesApi.md#updatePolicies) | **PUT** /api/policies/{id} | Updates a Policy
[**updatePoliciesCloud**](PoliciesApi.md#updatePoliciesCloud) | **PUT** /api/zones/{cloudId}/policies/{id} | Updates a Policy for a Cloud
[**updatePoliciesGroup**](PoliciesApi.md#updatePoliciesGroup) | **PUT** /api/groups/{groupId}/policies/{id} | Updates a Policy for a Group



## addPolicies

> Object addPolicies(opts)

Creates a Policy

Creates a policy. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let opts = {
  'inlineObject184': new MorpheusApi.InlineObject184() // InlineObject184 | 
};
apiInstance.addPolicies(opts, (error, data, response) => {
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
 **inlineObject184** | [**InlineObject184**](InlineObject184.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addPoliciesCloud

> Object addPoliciesCloud(cloudId, opts)

Creates a Policy for a Cloud

Creates a policy for a Cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let cloudId = 7; // Number | The ID of the cloud
let opts = {
  'inlineObject188': new MorpheusApi.InlineObject188() // InlineObject188 | 
};
apiInstance.addPoliciesCloud(cloudId, opts, (error, data, response) => {
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
 **cloudId** | **Number**| The ID of the cloud | 
 **inlineObject188** | [**InlineObject188**](InlineObject188.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addPoliciesGroup

> Object addPoliciesGroup(groupId, opts)

Creates a Policy for a Group

Creates a policy for a Group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let groupId = 7; // Number | The ID of the group
let opts = {
  'inlineObject186': new MorpheusApi.InlineObject186() // InlineObject186 | 
};
apiInstance.addPoliciesGroup(groupId, opts, (error, data, response) => {
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
 **groupId** | **Number**| The ID of the group | 
 **inlineObject186** | [**InlineObject186**](InlineObject186.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getPolicies

> Object getPolicies(id)

Retrieves a Specific Policy

Retrieves a specific policy. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getPolicies(id, (error, data, response) => {
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

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getPoliciesCloud

> Object getPoliciesCloud(cloudId, id)

Retrieves a Specific Policy for a Cloud

Retrieves a specific policy for a Cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let cloudId = 7; // Number | The ID of the cloud
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getPoliciesCloud(cloudId, id, (error, data, response) => {
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
 **cloudId** | **Number**| The ID of the cloud | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getPoliciesGroup

> Object getPoliciesGroup(groupId, id)

Retrieves a Specific Policy for a Group

Retrieves a specific policy for a Group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let groupId = 7; // Number | The ID of the group
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getPoliciesGroup(groupId, id, (error, data, response) => {
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
 **groupId** | **Number**| The ID of the group | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listPolicies

> Object listPolicies(opts)

Retrieves all Policies

Retrieves all policies. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listPolicies(opts, (error, data, response) => {
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


## listPoliciesCloud

> Object listPoliciesCloud(cloudId, opts)

Retrieves Policies for a Cloud

Retrieves policies for a specific cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let cloudId = 7; // Number | The ID of the cloud
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listPoliciesCloud(cloudId, opts, (error, data, response) => {
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
 **cloudId** | **Number**| The ID of the cloud | 
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


## listPoliciesGroup

> Object listPoliciesGroup(groupId, opts)

Retrieves Policies for a Group

Retrieves policies for a specific group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let groupId = 7; // Number | The ID of the group
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listPoliciesGroup(groupId, opts, (error, data, response) => {
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
 **groupId** | **Number**| The ID of the group | 
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


## listPolicyTypes

> Object listPolicyTypes()

Retrieves all Policy Types

Retrieves all Policy Types 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
apiInstance.listPolicyTypes((error, data, response) => {
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

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removePolicies

> Model200Success removePolicies(id)

Deletes a Policy

Deletes a specified policy. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removePolicies(id, (error, data, response) => {
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


## removePoliciesCloud

> Model200Success removePoliciesCloud(cloudId, id)

Deletes a Policy for a Cloud

Deletes a specified policy for a Cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let cloudId = 7; // Number | The ID of the cloud
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removePoliciesCloud(cloudId, id, (error, data, response) => {
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
 **cloudId** | **Number**| The ID of the cloud | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removePoliciesGroup

> Model200Success removePoliciesGroup(groupId, id)

Deletes a Policy for a Group

Deletes a specified policy for a Group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let groupId = 7; // Number | The ID of the group
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removePoliciesGroup(groupId, id, (error, data, response) => {
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
 **groupId** | **Number**| The ID of the group | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updatePolicies

> Object updatePolicies(id, opts)

Updates a Policy

Updates a policy. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject185': new MorpheusApi.InlineObject185() // InlineObject185 | 
};
apiInstance.updatePolicies(id, opts, (error, data, response) => {
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
 **inlineObject185** | [**InlineObject185**](InlineObject185.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updatePoliciesCloud

> Object updatePoliciesCloud(cloudId, id, opts)

Updates a Policy for a Cloud

Updates a policy for a Cloud. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let cloudId = 7; // Number | The ID of the cloud
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject189': new MorpheusApi.InlineObject189() // InlineObject189 | 
};
apiInstance.updatePoliciesCloud(cloudId, id, opts, (error, data, response) => {
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
 **cloudId** | **Number**| The ID of the cloud | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject189** | [**InlineObject189**](InlineObject189.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updatePoliciesGroup

> Object updatePoliciesGroup(groupId, id, opts)

Updates a Policy for a Group

Updates a policy for a Group. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.PoliciesApi();
let groupId = 7; // Number | The ID of the group
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject187': new MorpheusApi.InlineObject187() // InlineObject187 | 
};
apiInstance.updatePoliciesGroup(groupId, id, opts, (error, data, response) => {
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
 **groupId** | **Number**| The ID of the group | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject187** | [**InlineObject187**](InlineObject187.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

