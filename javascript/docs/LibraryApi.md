# MorpheusApi.LibraryApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addFileTemplate**](LibraryApi.md#addFileTemplate) | **POST** /api/library/container-templates | Create a File Template
[**addInstanceType**](LibraryApi.md#addInstanceType) | **POST** /api/library/instance-types | Create an Instance Type
[**addLayout**](LibraryApi.md#addLayout) | **POST** /api/library/instance-types/{instanceTypeId}/layouts | Create a Layout
[**addNodeType**](LibraryApi.md#addNodeType) | **POST** /api/library/container-types | Create a Node Type
[**addOptionList**](LibraryApi.md#addOptionList) | **POST** /api/library/option-type-lists | Create an Option List
[**addOptionType**](LibraryApi.md#addOptionType) | **POST** /api/library/option-types | Create an Input
[**addScript**](LibraryApi.md#addScript) | **POST** /api/library/container-scripts | Create a Script
[**addSpecTemplate**](LibraryApi.md#addSpecTemplate) | **POST** /api/library/spec-templates | Create a Spec Template
[**addVirtualImage**](LibraryApi.md#addVirtualImage) | **POST** /api/virtual-images | Create a Virtual Image
[**addVirtualImageFile**](LibraryApi.md#addVirtualImageFile) | **POST** /api/virtual-images/{virtualImageId}/upload | Upload Virtual Image File
[**deleteFileTemplate**](LibraryApi.md#deleteFileTemplate) | **DELETE** /api/library/container-templates/{id} | Delete a File Template
[**deleteInstanceType**](LibraryApi.md#deleteInstanceType) | **DELETE** /api/library/instance-types/{instanceTypeId} | Delete an Instance Type
[**deleteLayout**](LibraryApi.md#deleteLayout) | **DELETE** /api/library/layouts/{id} | Delete a Layout
[**deleteNodeType**](LibraryApi.md#deleteNodeType) | **DELETE** /api/library/container-types/{id} | Delete a Node Type
[**deleteOptionList**](LibraryApi.md#deleteOptionList) | **DELETE** /api/library/option-type-lists/{id} | Delete an Option List
[**deleteOptionType**](LibraryApi.md#deleteOptionType) | **DELETE** /api/library/option-types/{id} | Delete an Input
[**deleteScript**](LibraryApi.md#deleteScript) | **DELETE** /api/library/container-scripts/{id} | Delete a Script
[**deleteSpecTemplate**](LibraryApi.md#deleteSpecTemplate) | **DELETE** /api/library/spec-templates/{id} | Delete a Spec Template
[**getFileTemplate**](LibraryApi.md#getFileTemplate) | **GET** /api/library/container-templates/{id} | Get a Specific File Template
[**getInput**](LibraryApi.md#getInput) | **GET** /api/library/option-types/{id} | Get A Specific Input
[**getInstanceType**](LibraryApi.md#getInstanceType) | **GET** /api/library/instance-types/{instanceTypeId} | Get a Specific Instance Type
[**getLayout**](LibraryApi.md#getLayout) | **GET** /api/library/layouts/{id} | Get a Specific Layout
[**getNodeType**](LibraryApi.md#getNodeType) | **GET** /api/library/container-types/{id} | Get a Specific Node Type
[**getOptionList**](LibraryApi.md#getOptionList) | **GET** /api/library/option-type-lists/{id} | Get a Specific Option List
[**getOptionListItems**](LibraryApi.md#getOptionListItems) | **GET** /api/library/option-type-lists/{id}/items | List Items for a Specific Option List
[**getScript**](LibraryApi.md#getScript) | **GET** /api/library/container-scripts/{id} | Get a Specific Script
[**getSecurityPackageType**](LibraryApi.md#getSecurityPackageType) | **GET** /api/security-package-types/{id} | Retrieves a Specific Security Package Type
[**getSpecTemplate**](LibraryApi.md#getSpecTemplate) | **GET** /api/library/spec-templates/{id} | Get a Specific Spec Template
[**getVirtualImage**](LibraryApi.md#getVirtualImage) | **GET** /api/virtual-images/{virtualImageId} | Get a Specific Virtual Image
[**listFileTemplates**](LibraryApi.md#listFileTemplates) | **GET** /api/library/container-templates | Get All File Templates
[**listInputs**](LibraryApi.md#listInputs) | **GET** /api/library/option-types | Get All Inputs
[**listInstanceTypes**](LibraryApi.md#listInstanceTypes) | **GET** /api/library/instance-types | Get All Instance Types
[**listLayouts**](LibraryApi.md#listLayouts) | **GET** /api/library/layouts | Get All Layouts
[**listLayoutsForInstanceType**](LibraryApi.md#listLayoutsForInstanceType) | **GET** /api/library/instance-types/{instanceTypeId}/layouts | Get All Layouts For an Instance Type
[**listNodeTypes**](LibraryApi.md#listNodeTypes) | **GET** /api/library/container-types | Get All Node Types
[**listOptionLists**](LibraryApi.md#listOptionLists) | **GET** /api/library/option-type-lists | Get All Option Lists
[**listScripts**](LibraryApi.md#listScripts) | **GET** /api/library/container-scripts | Get All Scripts
[**listSecurityPackageTypes**](LibraryApi.md#listSecurityPackageTypes) | **GET** /api/security-package-types | Retrieves all Security Package Types
[**listSpecTemplates**](LibraryApi.md#listSpecTemplates) | **GET** /api/library/spec-templates | Get All Spec Templates
[**listVirtualImageLocations**](LibraryApi.md#listVirtualImageLocations) | **GET** /api/virtual-images/{virtualImageId}/locations | Get a List of Virtual Image Locations
[**listVirtualImages**](LibraryApi.md#listVirtualImages) | **GET** /api/virtual-images | Get List of Virtual Images
[**removeSecurityScans**](LibraryApi.md#removeSecurityScans) | **DELETE** /api/security-scans/{id} | Deletes a Security Scan
[**removeVirtualImage**](LibraryApi.md#removeVirtualImage) | **DELETE** /api/virtual-images/{virtualImageId} | Delete a Virtual Image
[**removeVirtualImageFile**](LibraryApi.md#removeVirtualImageFile) | **DELETE** /api/virtual-images/{virtualImageId}/files | Remove Virtual Image File
[**removeVirtualImageLocation**](LibraryApi.md#removeVirtualImageLocation) | **DELETE** /api/virtual-images/{virtualImageId}/locations/{id} | Delete a Virtual Image Location
[**setInstanceTypeFeatured**](LibraryApi.md#setInstanceTypeFeatured) | **PUT** /api/library/instance-types/{instanceTypeId}/toggle-featured | Toggle Featured For Instance Type
[**updateFileTemplate**](LibraryApi.md#updateFileTemplate) | **PUT** /api/library/container-templates/{id} | Update a File Template
[**updateInstanceType**](LibraryApi.md#updateInstanceType) | **PUT** /api/library/instance-types/{instanceTypeId} | Update an Instance Type
[**updateInstanceTypeLogo**](LibraryApi.md#updateInstanceTypeLogo) | **POST** /api/library/instance-types/{instanceTypeId}/update-logo | Update Logo For Instance Type
[**updateLayout**](LibraryApi.md#updateLayout) | **PUT** /api/library/layouts/{id} | Update a Layout
[**updateLayoutPermissions**](LibraryApi.md#updateLayoutPermissions) | **PUT** /api/library/layouts/{id}/permissions | Update Layout Permissions
[**updateNodeType**](LibraryApi.md#updateNodeType) | **PUT** /api/library/container-types/{id} | Update a Node Type
[**updateOptionList**](LibraryApi.md#updateOptionList) | **PUT** /api/library/option-type-lists/{id} | Update an Option List
[**updateOptionType**](LibraryApi.md#updateOptionType) | **PUT** /api/library/option-types/{id} | Update an Input
[**updateScript**](LibraryApi.md#updateScript) | **PUT** /api/library/container-scripts/{id} | Update a Script
[**updateSpecTemplate**](LibraryApi.md#updateSpecTemplate) | **PUT** /api/library/spec-templates/{id} | Update a Spec Template
[**updateVirtualImage**](LibraryApi.md#updateVirtualImage) | **PUT** /api/virtual-images/{virtualImageId} | Update a Virtual Image



## addFileTemplate

> SuccessId addFileTemplate(opts)

Create a File Template

Use this command to create a file template.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'inlineObject111': new MorpheusApi.InlineObject111() // InlineObject111 | 
};
apiInstance.addFileTemplate(opts, (error, data, response) => {
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
 **inlineObject111** | [**InlineObject111**](InlineObject111.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addInstanceType

> Model200Success addInstanceType(opts)

Create an Instance Type

Use this command to create an instance type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'inlineObject113': new MorpheusApi.InlineObject113() // InlineObject113 | 
};
apiInstance.addInstanceType(opts, (error, data, response) => {
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
 **inlineObject113** | [**InlineObject113**](InlineObject113.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addLayout

> SuccessId addLayout(instanceTypeId, opts)

Create a Layout

Use this command to create a layout.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let instanceTypeId = 2; // Number | The ID of the instance type
let opts = {
  'inlineObject115': new MorpheusApi.InlineObject115() // InlineObject115 | 
};
apiInstance.addLayout(instanceTypeId, opts, (error, data, response) => {
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
 **instanceTypeId** | **Number**| The ID of the instance type | 
 **inlineObject115** | [**InlineObject115**](InlineObject115.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addNodeType

> Object addNodeType(opts)

Create a Node Type

Use this command to create a node type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'inlineObject109': new MorpheusApi.InlineObject109() // InlineObject109 | 
};
apiInstance.addNodeType(opts, (error, data, response) => {
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
 **inlineObject109** | [**InlineObject109**](InlineObject109.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addOptionList

> Model200Success addOptionList(opts)

Create an Option List

Use this command to create an option list.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'inlineObject119': new MorpheusApi.InlineObject119() // InlineObject119 | 
};
apiInstance.addOptionList(opts, (error, data, response) => {
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
 **inlineObject119** | [**InlineObject119**](InlineObject119.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addOptionType

> SuccessId addOptionType(opts)

Create an Input

Use this command to create an option type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'inlineObject121': new MorpheusApi.InlineObject121() // InlineObject121 | 
};
apiInstance.addOptionType(opts, (error, data, response) => {
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
 **inlineObject121** | [**InlineObject121**](InlineObject121.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addScript

> ScriptSuccessId addScript(opts)

Create a Script

Use this command to create a script.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'inlineObject107': new MorpheusApi.InlineObject107() // InlineObject107 | 
};
apiInstance.addScript(opts, (error, data, response) => {
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
 **inlineObject107** | [**InlineObject107**](InlineObject107.md)|  | [optional] 

### Return type

[**ScriptSuccessId**](ScriptSuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addSpecTemplate

> SuccessId addSpecTemplate(opts)

Create a Spec Template

Use this command to create a spec template.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'inlineObject123': new MorpheusApi.InlineObject123() // InlineObject123 | 
};
apiInstance.addSpecTemplate(opts, (error, data, response) => {
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
 **inlineObject123** | [**InlineObject123**](InlineObject123.md)|  | [optional] 

### Return type

[**SuccessId**](SuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addVirtualImage

> Object addVirtualImage(opts)

Create a Virtual Image

This endpoint creates a new virtual image, without any files yet.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'inlineObject263': new MorpheusApi.InlineObject263() // InlineObject263 | 
};
apiInstance.addVirtualImage(opts, (error, data, response) => {
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
 **inlineObject263** | [**InlineObject263**](InlineObject263.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addVirtualImageFile

> Model200Success addVirtualImageFile(virtualImageId, opts)

Upload Virtual Image File

This will upload the file and associate it to the Virtual Image.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let virtualImageId = 4; // Number | Virtual Image ID
let opts = {
  'filename': testimage.ovf, // String | The name of the file
  'url': https://example.com/testimage.ovf, // String | Download the file from a remote url. This can be used instead of uploading a local file.
  'body': "/path/to/file" // File | 
};
apiInstance.addVirtualImageFile(virtualImageId, opts, (error, data, response) => {
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
 **virtualImageId** | **Number**| Virtual Image ID | 
 **filename** | **String**| The name of the file | [optional] 
 **url** | **String**| Download the file from a remote url. This can be used instead of uploading a local file. | [optional] 
 **body** | **File**|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json


## deleteFileTemplate

> Model200Success deleteFileTemplate(id)

Delete a File Template

Will delete a file template

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteFileTemplate(id, (error, data, response) => {
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


## deleteInstanceType

> Model200Success deleteInstanceType(instanceTypeId)

Delete an Instance Type

Will delete an instance type

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let instanceTypeId = 2; // Number | The ID of the instance type
apiInstance.deleteInstanceType(instanceTypeId, (error, data, response) => {
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
 **instanceTypeId** | **Number**| The ID of the instance type | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteLayout

> Model200Success deleteLayout(id)

Delete a Layout

Will delete a layout

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteLayout(id, (error, data, response) => {
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


## deleteNodeType

> Model200Success deleteNodeType(id)

Delete a Node Type

Will delete a node type

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteNodeType(id, (error, data, response) => {
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


## deleteOptionList

> Model200Success deleteOptionList(id)

Delete an Option List

Will delete an option list.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteOptionList(id, (error, data, response) => {
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


## deleteOptionType

> Model200Success deleteOptionType(id)

Delete an Input

Will delete an option type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteOptionType(id, (error, data, response) => {
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


## deleteScript

> Model200Success deleteScript(id)

Delete a Script

Will delete a script

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteScript(id, (error, data, response) => {
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


## deleteSpecTemplate

> Model200Success deleteSpecTemplate(id)

Delete a Spec Template

Will delete a spec template

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteSpecTemplate(id, (error, data, response) => {
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


## getFileTemplate

> InlineResponse20070 getFileTemplate(id)

Get a Specific File Template

This endpoint retrieves a specific file template.  The value of template will be masked as ************ for system owned file templates. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getFileTemplate(id, (error, data, response) => {
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

[**InlineResponse20070**](InlineResponse20070.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getInput

> InlineResponse20075 getInput(id)

Get A Specific Input

This endpoint retrieves a specific option type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getInput(id, (error, data, response) => {
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

[**InlineResponse20075**](InlineResponse20075.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getInstanceType

> InlineResponse20071 getInstanceType(instanceTypeId)

Get a Specific Instance Type

This endpoint retrieves a specific instance type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let instanceTypeId = 2; // Number | The ID of the instance type
apiInstance.getInstanceType(instanceTypeId, (error, data, response) => {
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
 **instanceTypeId** | **Number**| The ID of the instance type | 

### Return type

[**InlineResponse20071**](InlineResponse20071.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getLayout

> InlineResponse20072 getLayout(id)

Get a Specific Layout

This endpoint retrieves a specific layout.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getLayout(id, (error, data, response) => {
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

[**InlineResponse20072**](InlineResponse20072.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getNodeType

> InlineResponse20069 getNodeType(id)

Get a Specific Node Type

This endpoint retrieves a specific node type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getNodeType(id, (error, data, response) => {
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

[**InlineResponse20069**](InlineResponse20069.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getOptionList

> InlineResponse20073 getOptionList(id)

Get a Specific Option List

This endpoint retrieves a specific option list.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getOptionList(id, (error, data, response) => {
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

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getOptionListItems

> InlineResponse20074 getOptionListItems(id)

List Items for a Specific Option List

This endpoint retrieves the items for a specific option list.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getOptionListItems(id, (error, data, response) => {
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

[**InlineResponse20074**](InlineResponse20074.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getScript

> InlineResponse20068 getScript(id)

Get a Specific Script

This endpoint retrieves a specific script.  The value of script will be masked as ************ for system owned scripts. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getScript(id, (error, data, response) => {
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

[**InlineResponse20068**](InlineResponse20068.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getSecurityPackageType

> InlineResponse200136 getSecurityPackageType(id)

Retrieves a Specific Security Package Type

Retrieves a specific security package type. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getSecurityPackageType(id, (error, data, response) => {
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

[**InlineResponse200136**](InlineResponse200136.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getSpecTemplate

> InlineResponse20076 getSpecTemplate(id)

Get a Specific Spec Template

This endpoint retrieves a specific spec template.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getSpecTemplate(id, (error, data, response) => {
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

[**InlineResponse20076**](InlineResponse20076.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getVirtualImage

> InlineResponse200165 getVirtualImage(virtualImageId)

Get a Specific Virtual Image

This endpoint retrieves a specific virtual image and its files.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let virtualImageId = 4; // Number | Virtual Image ID
apiInstance.getVirtualImage(virtualImageId, (error, data, response) => {
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
 **virtualImageId** | **Number**| Virtual Image ID | 

### Return type

[**InlineResponse200165**](InlineResponse200165.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listFileTemplates

> Object listFileTemplates(opts)

Get All File Templates

This endpoint retrieves all file templates.  The value of template will be masked as ************ for system owned file templates. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example", // String | Filter by label(s), matches records that contain all of the specified labels
  'fileName': testtext.txt // String | Filename filter, restricts query to only load file template matching fileName specified
};
apiInstance.listFileTemplates(opts, (error, data, response) => {
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
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 
 **fileName** | **String**| Filename filter, restricts query to only load file template matching fileName specified | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listInputs

> Object listInputs(opts)

Get All Inputs

This endpoint retrieves all option types. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example", // String | Filter by label(s), matches records that contain all of the specified labels
  'fieldName': cloudName, // String | Field Name filter, restricts query to only load type matching fieldName specified
  'fieldContext': config.customOptions, // String | Field Context filter, restricts query to only load type matching fieldContext specified
  'fieldLabel': DB Version // String | Field Label filter, restricts query to only load type matching fieldLabel specified
};
apiInstance.listInputs(opts, (error, data, response) => {
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
 **code** | **String**| If specified will return an exact match on code | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 
 **fieldName** | **String**| Field Name filter, restricts query to only load type matching fieldName specified | [optional] 
 **fieldContext** | **String**| Field Context filter, restricts query to only load type matching fieldContext specified | [optional] 
 **fieldLabel** | **String**| Field Label filter, restricts query to only load type matching fieldLabel specified | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listInstanceTypes

> Object listInstanceTypes(opts)

Get All Instance Types

This endpoint retrieves all instance types. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'featured': false, // Boolean | Filter by featured
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example", // String | Filter by label(s), matches records that contain all of the specified labels
  'details': true // Boolean | Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default.
};
apiInstance.listInstanceTypes(opts, (error, data, response) => {
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
 **code** | **String**| If specified will return an exact match on code | [optional] 
 **featured** | **Boolean**| Filter by featured | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 
 **details** | **Boolean**| Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listLayouts

> Object listLayouts(opts)

Get All Layouts

This endpoint retrieves all layouts. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'provisionType': "provisionType_example", // String | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings. 
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listLayouts(opts, (error, data, response) => {
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
 **code** | **String**| If specified will return an exact match on code | [optional] 
 **provisionType** | **String**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listLayoutsForInstanceType

> Object listLayoutsForInstanceType(instanceTypeId, opts)

Get All Layouts For an Instance Type

This endpoint retrieves all layouts for a specific instance type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let instanceTypeId = 2; // Number | The ID of the instance type
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'provisionType': "provisionType_example", // String | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings. 
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listLayoutsForInstanceType(instanceTypeId, opts, (error, data, response) => {
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
 **instanceTypeId** | **Number**| The ID of the instance type | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **code** | **String**| If specified will return an exact match on code | [optional] 
 **provisionType** | **String**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listNodeTypes

> Object listNodeTypes(opts)

Get All Node Types

This endpoint retrieves all node types.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'provisionType': "provisionType_example", // String | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings. 
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listNodeTypes(opts, (error, data, response) => {
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
 **code** | **String**| If specified will return an exact match on code | [optional] 
 **provisionType** | **String**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listOptionLists

> Object listOptionLists(opts)

Get All Option Lists

This endpoint retrieves all option lists.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listOptionLists(opts, (error, data, response) => {
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
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listScripts

> Object listScripts(opts)

Get All Scripts

This endpoint retrieves all scripts.  The value of script will be masked as ************ for system owned scripts. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example", // String | Filter by label(s), matches records that contain all of the specified labels
  'scriptType': "scriptType_example", // String | Script type code filter, restricts query to only load scripts of specified type
  'scriptPhase': "scriptPhase_example" // String | Script phase filter, restricts query to only load scripts of specified phase
};
apiInstance.listScripts(opts, (error, data, response) => {
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
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 
 **scriptType** | **String**| Script type code filter, restricts query to only load scripts of specified type | [optional] 
 **scriptPhase** | **String**| Script phase filter, restricts query to only load scripts of specified phase | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listSecurityPackageTypes

> Object listSecurityPackageTypes()

Retrieves all Security Package Types

Retrieves all Security Package Types 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
apiInstance.listSecurityPackageTypes((error, data, response) => {
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


## listSpecTemplates

> Object listSpecTemplates(opts)

Get All Spec Templates

This endpoint retrieves all spec templates.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listSpecTemplates(opts, (error, data, response) => {
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
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listVirtualImageLocations

> Object listVirtualImageLocations(virtualImageId)

Get a List of Virtual Image Locations

This endpoint retrieves a specific virtual image and its files.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let virtualImageId = 4; // Number | Virtual Image ID
apiInstance.listVirtualImageLocations(virtualImageId, (error, data, response) => {
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
 **virtualImageId** | **Number**| Virtual Image ID | 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listVirtualImages

> Object listVirtualImages(opts)

Get List of Virtual Images

This endpoint retrieves a list of virtual images for the specified filter.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'lastUpdated': 2019-03-06T17:52:29+0000, // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
  'filterType': System, // String | Filter by type, \"User\", \"System\", \"Synced\", or \"All\"
  'imageType': vmware, // String | Filter by image type code, \"vmware\", \"ami\", etc
  'tags': tags.env=qa&tags.env=test, // String | Filter by tags (metadata). This allows filtering by a tag name and value(s) 
  'labels': "labels_example", // String | Filter by label(s), matches records that contain any of the specified labels
  'allLabels': "allLabels_example" // String | Filter by label(s), matches records that contain all of the specified labels
};
apiInstance.listVirtualImages(opts, (error, data, response) => {
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 
 **filterType** | **String**| Filter by type, \&quot;User\&quot;, \&quot;System\&quot;, \&quot;Synced\&quot;, or \&quot;All\&quot; | [optional] [default to &#39;User&#39;]
 **imageType** | **String**| Filter by image type code, \&quot;vmware\&quot;, \&quot;ami\&quot;, etc | [optional] 
 **tags** | **String**| Filter by tags (metadata). This allows filtering by a tag name and value(s)  | [optional] 
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] 
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeSecurityScans

> Model200Success removeSecurityScans(id)

Deletes a Security Scan

Deletes a specified security scan. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeSecurityScans(id, (error, data, response) => {
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


## removeVirtualImage

> Model200Success removeVirtualImage(virtualImageId)

Delete a Virtual Image

Will delete a virtual image and any associated files, use removeFromCloud&#x3D;true to also delete image locations from all clouds.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let virtualImageId = 4; // Number | Virtual Image ID
apiInstance.removeVirtualImage(virtualImageId, (error, data, response) => {
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
 **virtualImageId** | **Number**| Virtual Image ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeVirtualImageFile

> Model200Success removeVirtualImageFile(virtualImageId, opts)

Remove Virtual Image File

Remove Virtual Image File

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let virtualImageId = 4; // Number | Virtual Image ID
let opts = {
  'filename': testimage.ovf // String | The name of the file
};
apiInstance.removeVirtualImageFile(virtualImageId, opts, (error, data, response) => {
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
 **virtualImageId** | **Number**| Virtual Image ID | 
 **filename** | **String**| The name of the file | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeVirtualImageLocation

> Model200Success removeVirtualImageLocation(virtualImageId, id)

Delete a Virtual Image Location

Will delete a virtual image location, use removeFromCloud&#x3D;true to all also delete image locations from all clouds as well.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let virtualImageId = 4; // Number | Virtual Image ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeVirtualImageLocation(virtualImageId, id, (error, data, response) => {
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
 **virtualImageId** | **Number**| Virtual Image ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## setInstanceTypeFeatured

> Model200Success setInstanceTypeFeatured(instanceTypeId)

Toggle Featured For Instance Type

Use this command to toggle the featured flag for an existing instance type. This will change the value from false to true, or from true to false. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let instanceTypeId = 2; // Number | The ID of the instance type
apiInstance.setInstanceTypeFeatured(instanceTypeId, (error, data, response) => {
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
 **instanceTypeId** | **Number**| The ID of the instance type | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateFileTemplate

> Model200Success updateFileTemplate(id, opts)

Update a File Template

Use this command to update an existing file template.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject112': new MorpheusApi.InlineObject112() // InlineObject112 | 
};
apiInstance.updateFileTemplate(id, opts, (error, data, response) => {
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
 **inlineObject112** | [**InlineObject112**](InlineObject112.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateInstanceType

> Model200Success updateInstanceType(instanceTypeId, opts)

Update an Instance Type

Use this command to update an existing instance type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let instanceTypeId = 2; // Number | The ID of the instance type
let opts = {
  'inlineObject114': new MorpheusApi.InlineObject114() // InlineObject114 | 
};
apiInstance.updateInstanceType(instanceTypeId, opts, (error, data, response) => {
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
 **instanceTypeId** | **Number**| The ID of the instance type | 
 **inlineObject114** | [**InlineObject114**](InlineObject114.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateInstanceTypeLogo

> Model200Success updateInstanceTypeLogo(instanceTypeId, opts)

Update Logo For Instance Type

Use this command to update the logo and dark logo images for an existing instance type. This endpoint expects multipart form data as the request format, not JSON. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let instanceTypeId = 2; // Number | The ID of the instance type
let opts = {
  'logo': "/path/to/file", // File | Logo File png,jpg,svg
  'darkLogo': "/path/to/file" // File | Dark Logo File png,jpg,svg
};
apiInstance.updateInstanceTypeLogo(instanceTypeId, opts, (error, data, response) => {
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
 **instanceTypeId** | **Number**| The ID of the instance type | 
 **logo** | **File**| Logo File png,jpg,svg | [optional] 
 **darkLogo** | **File**| Dark Logo File png,jpg,svg | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## updateLayout

> Model200Success updateLayout(id, opts)

Update a Layout

Use this command to update an existing layout.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject117': new MorpheusApi.InlineObject117() // InlineObject117 | 
};
apiInstance.updateLayout(id, opts, (error, data, response) => {
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
 **inlineObject117** | [**InlineObject117**](InlineObject117.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateLayoutPermissions

> Model200Success updateLayoutPermissions(id, opts)

Update Layout Permissions

Use this command to update permissions for an existing layout.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject118': new MorpheusApi.InlineObject118() // InlineObject118 | 
};
apiInstance.updateLayoutPermissions(id, opts, (error, data, response) => {
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
 **inlineObject118** | [**InlineObject118**](InlineObject118.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateNodeType

> Model200Success updateNodeType(id, opts)

Update a Node Type

Use this command to update an existing node type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject110': new MorpheusApi.InlineObject110() // InlineObject110 | 
};
apiInstance.updateNodeType(id, opts, (error, data, response) => {
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
 **inlineObject110** | [**InlineObject110**](InlineObject110.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateOptionList

> Model200Success updateOptionList(id, opts)

Update an Option List

Use this command to update an existing option list.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject120': new MorpheusApi.InlineObject120() // InlineObject120 | 
};
apiInstance.updateOptionList(id, opts, (error, data, response) => {
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
 **inlineObject120** | [**InlineObject120**](InlineObject120.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateOptionType

> Model200Success updateOptionType(id, opts)

Update an Input

Use this command to update an existing option type.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject122': new MorpheusApi.InlineObject122() // InlineObject122 | 
};
apiInstance.updateOptionType(id, opts, (error, data, response) => {
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
 **inlineObject122** | [**InlineObject122**](InlineObject122.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateScript

> ScriptSuccessId updateScript(id, opts)

Update a Script

Use this command to update an existing script.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject108': new MorpheusApi.InlineObject108() // InlineObject108 | 
};
apiInstance.updateScript(id, opts, (error, data, response) => {
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
 **inlineObject108** | [**InlineObject108**](InlineObject108.md)|  | [optional] 

### Return type

[**ScriptSuccessId**](ScriptSuccessId.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateSpecTemplate

> Model200Success updateSpecTemplate(id, opts)

Update a Spec Template

Use this command to update an existing spec template.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject124': new MorpheusApi.InlineObject124() // InlineObject124 | 
};
apiInstance.updateSpecTemplate(id, opts, (error, data, response) => {
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
 **inlineObject124** | [**InlineObject124**](InlineObject124.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateVirtualImage

> Object updateVirtualImage(virtualImageId, opts)

Update a Virtual Image

This endpoint updates an existing virtual image.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.LibraryApi();
let virtualImageId = 4; // Number | Virtual Image ID
let opts = {
  'inlineObject264': new MorpheusApi.InlineObject264() // InlineObject264 | 
};
apiInstance.updateVirtualImage(virtualImageId, opts, (error, data, response) => {
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
 **virtualImageId** | **Number**| Virtual Image ID | 
 **inlineObject264** | [**InlineObject264**](InlineObject264.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

