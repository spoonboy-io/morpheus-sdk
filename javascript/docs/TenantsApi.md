# MorpheusApi.TenantsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addTenant**](TenantsApi.md#addTenant) | **POST** /api/accounts | Create a Tenant
[**addUserTenant**](TenantsApi.md#addUserTenant) | **POST** /api/accounts/{accountId}/users | Create a User For a Tenant
[**createTenantSubtenantGroup**](TenantsApi.md#createTenantSubtenantGroup) | **POST** /api/accounts/{accountId}/groups | Create a Group for Subtenant
[**getTenant**](TenantsApi.md#getTenant) | **GET** /api/accounts/{id} | Get tenant
[**getTenantSubtenantGroup**](TenantsApi.md#getTenantSubtenantGroup) | **GET** /api/accounts/{accountId}/groups/{id} | Get a Specific Group for Subtenant
[**listTenantSubtenantGroups**](TenantsApi.md#listTenantSubtenantGroups) | **GET** /api/accounts/{accountId}/groups | Get Subtenant Groups
[**listTenants**](TenantsApi.md#listTenants) | **GET** /api/accounts | List All Tenants
[**listTenantsAvailableRoles**](TenantsApi.md#listTenantsAvailableRoles) | **GET** /api/accounts/available-roles | List available roles for a tenant
[**removeTenant**](TenantsApi.md#removeTenant) | **DELETE** /api/accounts/{id} | Delete a Specific Tenant
[**removeTenantSubtenantGroup**](TenantsApi.md#removeTenantSubtenantGroup) | **DELETE** /api/accounts/{accountId}/groups/{id} | Delete a Group for Subtenant
[**updateTenant**](TenantsApi.md#updateTenant) | **PUT** /api/accounts/{id} | Update tenant
[**updateTenantSubtenantGroup**](TenantsApi.md#updateTenantSubtenantGroup) | **PUT** /api/accounts/{accountId}/groups/{id} | Updating a Group for Subtenant
[**updateTenantSubtenantGroupZones**](TenantsApi.md#updateTenantSubtenantGroupZones) | **PUT** /api/accounts/{accountId}/groups/{id}/update-zones | Updating Group Zones for Subtenant



## addTenant

> Object addTenant(opts)

Create a Tenant

Create a new tenant. This new account will be a sub-tenant with the master tenant as its parent.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
let opts = {
  'inlineObject238': new MorpheusApi.InlineObject238() // InlineObject238 | 
};
apiInstance.addTenant(opts, (error, data, response) => {
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
 **inlineObject238** | [**InlineObject238**](InlineObject238.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addUserTenant

> Object addUserTenant(accountId, opts)

Create a User For a Tenant

Create a User For a Tenant.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
let accountId = 7; // Number | The ID of the subtenant account
let opts = {
  'inlineObject243': new MorpheusApi.InlineObject243() // InlineObject243 | 
};
apiInstance.addUserTenant(accountId, opts, (error, data, response) => {
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
 **accountId** | **Number**| The ID of the subtenant account | 
 **inlineObject243** | [**InlineObject243**](InlineObject243.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createTenantSubtenantGroup

> InlineResponse200152 createTenantSubtenantGroup(accountId, opts)

Create a Group for Subtenant

Create a Group for Subtenant.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
let accountId = 7; // Number | The ID of the subtenant account
let opts = {
  'inlineObject240': new MorpheusApi.InlineObject240() // InlineObject240 | 
};
apiInstance.createTenantSubtenantGroup(accountId, opts, (error, data, response) => {
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
 **accountId** | **Number**| The ID of the subtenant account | 
 **inlineObject240** | [**InlineObject240**](InlineObject240.md)|  | [optional] 

### Return type

[**InlineResponse200152**](InlineResponse200152.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getTenant

> InlineResponse200150 getTenant(id)

Get tenant

Get details about a tenant

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getTenant(id, (error, data, response) => {
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

[**InlineResponse200150**](InlineResponse200150.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getTenantSubtenantGroup

> InlineResponse200153 getTenantSubtenantGroup(accountId, id)

Get a Specific Group for Subtenant

This endpoint retrieves a specific group.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
let accountId = 7; // Number | The ID of the subtenant account
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getTenantSubtenantGroup(accountId, id, (error, data, response) => {
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
 **accountId** | **Number**| The ID of the subtenant account | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse200153**](InlineResponse200153.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listTenantSubtenantGroups

> Object listTenantSubtenantGroups(accountId, opts)

Get Subtenant Groups

Groups belonging to a subtenant can be managed by the master account.  This endpoint retrieves all groups and a list of zones associated with the group by id. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
let accountId = 7; // Number | The ID of the subtenant account
let opts = {
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'lastUpdated': 2019-03-06T17:52:29+0000 // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
};
apiInstance.listTenantSubtenantGroups(accountId, opts, (error, data, response) => {
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
 **accountId** | **Number**| The ID of the subtenant account | 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listTenants

> Object listTenants(opts)

List All Tenants

Get a list of tenants. A tenant is also referred to as an account.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'lastUpdated': 2019-03-06T17:52:29+0000 // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
};
apiInstance.listTenants(opts, (error, data, response) => {
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
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listTenantsAvailableRoles

> TenantsAvailableRoles listTenantsAvailableRoles()

List available roles for a tenant

Get a list of available roles that can be assigned as the default base role for a sub tenant account.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
apiInstance.listTenantsAvailableRoles((error, data, response) => {
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

[**TenantsAvailableRoles**](TenantsAvailableRoles.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeTenant

> Model200Success removeTenant(id, opts)

Delete a Specific Tenant

Delete an existing tenant. This action is not reversible and will result in the removal of all data pertaining to this tenant as well as potentially any provisioned assets depending on the value of &#x60;removeResources&#x60;.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'removeResources': false // Boolean | Remove Resources. This will delete all the managed resources in the tenant.
};
apiInstance.removeTenant(id, opts, (error, data, response) => {
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
 **removeResources** | **Boolean**| Remove Resources. This will delete all the managed resources in the tenant. | [optional] [default to false]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeTenantSubtenantGroup

> Model200Success removeTenantSubtenantGroup(accountId, id)

Delete a Group for Subtenant

If a group has zones or servers still tied to it, a delete action will fail.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
let accountId = 7; // Number | The ID of the subtenant account
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeTenantSubtenantGroup(accountId, id, (error, data, response) => {
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
 **accountId** | **Number**| The ID of the subtenant account | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateTenant

> InlineResponse200151 updateTenant(id, opts)

Update tenant

Update an existing tenant.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject239': new MorpheusApi.InlineObject239() // InlineObject239 | 
};
apiInstance.updateTenant(id, opts, (error, data, response) => {
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
 **inlineObject239** | [**InlineObject239**](InlineObject239.md)|  | [optional] 

### Return type

[**InlineResponse200151**](InlineResponse200151.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateTenantSubtenantGroup

> InlineResponse200152 updateTenantSubtenantGroup(accountId, id, opts)

Updating a Group for Subtenant

Updating a Group for Subtenant.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
let accountId = 7; // Number | The ID of the subtenant account
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject241': new MorpheusApi.InlineObject241() // InlineObject241 | 
};
apiInstance.updateTenantSubtenantGroup(accountId, id, opts, (error, data, response) => {
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
 **accountId** | **Number**| The ID of the subtenant account | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject241** | [**InlineObject241**](InlineObject241.md)|  | [optional] 

### Return type

[**InlineResponse200152**](InlineResponse200152.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateTenantSubtenantGroupZones

> Model200Success updateTenantSubtenantGroupZones(accountId, id, opts)

Updating Group Zones for Subtenant

This will update the zones that are assigned to the group. Any zones that are not passed in the zones parameter will be removed from the group.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.TenantsApi();
let accountId = 7; // Number | The ID of the subtenant account
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject242': new MorpheusApi.InlineObject242() // InlineObject242 | 
};
apiInstance.updateTenantSubtenantGroupZones(accountId, id, opts, (error, data, response) => {
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
 **accountId** | **Number**| The ID of the subtenant account | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject242** | [**InlineObject242**](InlineObject242.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

