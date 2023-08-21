# MorpheusApi.RolesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addRoles**](RolesApi.md#addRoles) | **POST** /api/roles | Create role
[**deleteRole**](RolesApi.md#deleteRole) | **DELETE** /api/roles/{id} | Delete role
[**getRole**](RolesApi.md#getRole) | **GET** /api/roles/{id} | Get role
[**listRoles**](RolesApi.md#listRoles) | **GET** /api/roles | List roles
[**updateRole**](RolesApi.md#updateRole) | **PUT** /api/roles/{id} | Update role
[**updateRoleBlueprintAccess**](RolesApi.md#updateRoleBlueprintAccess) | **PUT** /api/roles/{id}/update-blueprint | Customizing Blueprint Access
[**updateRoleCatalogItemTypeAccess**](RolesApi.md#updateRoleCatalogItemTypeAccess) | **PUT** /api/roles/{id}/update-catalog-item-type | Customizing Catalog Item Type Access
[**updateRoleCloudAccess**](RolesApi.md#updateRoleCloudAccess) | **PUT** /api/roles/{id}/update-cloud | Customizing Cloud Access
[**updateRoleGroupAccess**](RolesApi.md#updateRoleGroupAccess) | **PUT** /api/roles/{id}/update-group | Customizing Group Access
[**updateRoleInstanceTypeAccess**](RolesApi.md#updateRoleInstanceTypeAccess) | **PUT** /api/roles/{id}/update-instance-type | Customizing Instance Type Access
[**updateRolePermission**](RolesApi.md#updateRolePermission) | **PUT** /api/roles/{id}/update-permission | Updating Role Permissions
[**updateRolePersonaAccess**](RolesApi.md#updateRolePersonaAccess) | **PUT** /api/roles/{id}/update-persona | Customizing Persona Access
[**updateRoleReportTypeAccess**](RolesApi.md#updateRoleReportTypeAccess) | **PUT** /api/roles/{id}/update-report-type | Customizing Report Type Access
[**updateRoleTaskAccess**](RolesApi.md#updateRoleTaskAccess) | **PUT** /api/roles/{id}/update-task | Customizing Task Access
[**updateRoleVDIPoolAccess**](RolesApi.md#updateRoleVDIPoolAccess) | **PUT** /api/roles/{id}/update-vdi-pool | Customizing VDI Pool Access
[**updateRoleWorkflowAccess**](RolesApi.md#updateRoleWorkflowAccess) | **PUT** /api/roles/{id}/update-task-set | Customizing Workflow Access



## addRoles

> Role addRoles(opts)

Create role

Create a new role.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let opts = {
  'includeDefaultAccess': true, // Boolean | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none` 
  'inlineObject208': new MorpheusApi.InlineObject208() // InlineObject208 | 
};
apiInstance.addRoles(opts, (error, data, response) => {
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
 **includeDefaultAccess** | **Boolean**| Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60;  | [optional] 
 **inlineObject208** | [**InlineObject208**](InlineObject208.md)|  | [optional] 

### Return type

[**Role**](Role.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteRole

> Model200Success deleteRole(id)

Delete role

Delete an existing role. A role cannot be deleted while it is still in use.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteRole(id, (error, data, response) => {
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


## getRole

> Role getRole(id, opts)

Get role

Get details about a role

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'includeDefaultAccess': true // Boolean | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none` 
};
apiInstance.getRole(id, opts, (error, data, response) => {
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
 **includeDefaultAccess** | **Boolean**| Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60;  | [optional] 

### Return type

[**Role**](Role.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listRoles

> Object listRoles(opts)

List roles

Get a list of roles.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let opts = {
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'authority': "authority_example", // String | Filter by authority
  'roleType': "roleType_example" // String | Filter by roleType
};
apiInstance.listRoles(opts, (error, data, response) => {
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
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **authority** | **String**| Filter by authority | [optional] 
 **roleType** | **String**| Filter by roleType | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateRole

> Role updateRole(id, opts)

Update role

Update an existing role.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'includeDefaultAccess': true, // Boolean | Pass true to include all resource permissions in the response including those with access set to `default`. Only resources with access specific levels are returned by default. eg. `full`, `read` or `none` 
  'inlineObject209': new MorpheusApi.InlineObject209() // InlineObject209 | 
};
apiInstance.updateRole(id, opts, (error, data, response) => {
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
 **includeDefaultAccess** | **Boolean**| Pass true to include all resource permissions in the response including those with access set to &#x60;default&#x60;. Only resources with access specific levels are returned by default. eg. &#x60;full&#x60;, &#x60;read&#x60; or &#x60;none&#x60;  | [optional] 
 **inlineObject209** | [**InlineObject209**](InlineObject209.md)|  | [optional] 

### Return type

[**Role**](Role.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateRoleBlueprintAccess

> Object updateRoleBlueprintAccess(id, opts)

Customizing Blueprint Access

Customizing Blueprint Access

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'UNKNOWN_BASE_TYPE': new MorpheusApi.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateRoleBlueprintAccess(id, opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateRoleCatalogItemTypeAccess

> Object updateRoleCatalogItemTypeAccess(id, opts)

Customizing Catalog Item Type Access

Customizing Catalog Item Type Access

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'UNKNOWN_BASE_TYPE': new MorpheusApi.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateRoleCatalogItemTypeAccess(id, opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateRoleCloudAccess

> Object updateRoleCloudAccess(id, opts)

Customizing Cloud Access

Customizing Cloud Access

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'UNKNOWN_BASE_TYPE': new MorpheusApi.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateRoleCloudAccess(id, opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateRoleGroupAccess

> Object updateRoleGroupAccess(id, opts)

Customizing Group Access

Customizing Group Access

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'UNKNOWN_BASE_TYPE': new MorpheusApi.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateRoleGroupAccess(id, opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateRoleInstanceTypeAccess

> Object updateRoleInstanceTypeAccess(id, opts)

Customizing Instance Type Access

Customizing Instance Type Access

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'UNKNOWN_BASE_TYPE': new MorpheusApi.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateRoleInstanceTypeAccess(id, opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateRolePermission

> Object updateRolePermission(id, opts)

Updating Role Permissions

Update a feature permission or default permission category (group, cloud, persona, ect.)

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'UNKNOWN_BASE_TYPE': {"permissionCode":"admin-users","access":"read"} // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateRolePermission(id, opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateRolePersonaAccess

> Object updateRolePersonaAccess(id, opts)

Customizing Persona Access

Customizing Persona Access

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'UNKNOWN_BASE_TYPE': new MorpheusApi.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateRolePersonaAccess(id, opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateRoleReportTypeAccess

> Object updateRoleReportTypeAccess(id, opts)

Customizing Report Type Access

Customizing Report Type Access

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'UNKNOWN_BASE_TYPE': new MorpheusApi.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateRoleReportTypeAccess(id, opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateRoleTaskAccess

> Object updateRoleTaskAccess(id, opts)

Customizing Task Access

Customizing Task Access

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'UNKNOWN_BASE_TYPE': new MorpheusApi.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateRoleTaskAccess(id, opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateRoleVDIPoolAccess

> Object updateRoleVDIPoolAccess(id, opts)

Customizing VDI Pool Access

Customizing VDI Pool Access

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject210': new MorpheusApi.InlineObject210() // InlineObject210 | 
};
apiInstance.updateRoleVDIPoolAccess(id, opts, (error, data, response) => {
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
 **inlineObject210** | [**InlineObject210**](InlineObject210.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateRoleWorkflowAccess

> Object updateRoleWorkflowAccess(id, opts)

Customizing Workflow Access

Customizing Workflow Access

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.RolesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'UNKNOWN_BASE_TYPE': new MorpheusApi.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateRoleWorkflowAccess(id, opts, (error, data, response) => {
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

