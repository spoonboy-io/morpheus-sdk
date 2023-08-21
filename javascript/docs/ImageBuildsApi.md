# MorpheusApi.ImageBuildsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addBootScript**](ImageBuildsApi.md#addBootScript) | **POST** /api/boot-scripts | Create a Boot Script
[**addImageBuild**](ImageBuildsApi.md#addImageBuild) | **POST** /api/image-builds | Create an Image Build
[**addPreseedScript**](ImageBuildsApi.md#addPreseedScript) | **POST** /api/preseed-scripts | Create a Preseed Script
[**deleteBootScript**](ImageBuildsApi.md#deleteBootScript) | **DELETE** /api/boot-scripts/{id} | Delete a Boot Script
[**deleteImageBuild**](ImageBuildsApi.md#deleteImageBuild) | **DELETE** /api/image-builds/{id} | Delete an Image Build
[**deletePreseedScript**](ImageBuildsApi.md#deletePreseedScript) | **DELETE** /api/preseed-scripts/{id} | Delete a Preseed Script
[**executeImageBuild**](ImageBuildsApi.md#executeImageBuild) | **POST** /api/image-builds/{id}/run | Run an Image Build
[**getBootScript**](ImageBuildsApi.md#getBootScript) | **GET** /api/boot-scripts/{id} | Get a Specific Boot Script
[**getImageBuild**](ImageBuildsApi.md#getImageBuild) | **GET** /api/image-builds/{id} | Get a Specific Image Build
[**getImageBuildExecutions**](ImageBuildsApi.md#getImageBuildExecutions) | **GET** /api/image-builds/{id}/list-executions | List Image Build Executions
[**getPreseedScript**](ImageBuildsApi.md#getPreseedScript) | **GET** /api/preseed-scripts/{id} | Get a Specific Preseed Script
[**listBootScripts**](ImageBuildsApi.md#listBootScripts) | **GET** /api/boot-scripts | Boot Scripts
[**listImageBuilds**](ImageBuildsApi.md#listImageBuilds) | **GET** /api/image-builds | Get All Image Builds
[**listPreseedScripts**](ImageBuildsApi.md#listPreseedScripts) | **GET** /api/preseed-scripts | Preseed Scripts
[**updateBootScript**](ImageBuildsApi.md#updateBootScript) | **PUT** /api/boot-scripts/{id} | Update a Boot Script
[**updateImageBuild**](ImageBuildsApi.md#updateImageBuild) | **PUT** /api/image-builds/{id} | Update an Image Build
[**updatePreseedScript**](ImageBuildsApi.md#updatePreseedScript) | **PUT** /api/preseed-scripts/{id} | Update a Preseed Script



## addBootScript

> Object addBootScript(opts)

Create a Boot Script

Create a Boot Script

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let opts = {
  'inlineObject83': new MorpheusApi.InlineObject83() // InlineObject83 | 
};
apiInstance.addBootScript(opts, (error, data, response) => {
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
 **inlineObject83** | [**InlineObject83**](InlineObject83.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addImageBuild

> Object addImageBuild(opts)

Create an Image Build

Create an Image Build

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let opts = {
  'inlineObject85': new MorpheusApi.InlineObject85() // InlineObject85 | 
};
apiInstance.addImageBuild(opts, (error, data, response) => {
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
 **inlineObject85** | [**InlineObject85**](InlineObject85.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addPreseedScript

> Object addPreseedScript(opts)

Create a Preseed Script

Create a Preseed Script

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let opts = {
  'inlineObject196': new MorpheusApi.InlineObject196() // InlineObject196 | 
};
apiInstance.addPreseedScript(opts, (error, data, response) => {
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
 **inlineObject196** | [**InlineObject196**](InlineObject196.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteBootScript

> Model200Success deleteBootScript(id)

Delete a Boot Script

Will delete a boot script from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteBootScript(id, (error, data, response) => {
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


## deleteImageBuild

> Model200Success deleteImageBuild(id)

Delete an Image Build

Will delete an image build from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteImageBuild(id, (error, data, response) => {
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


## deletePreseedScript

> Model200Success deletePreseedScript(id)

Delete a Preseed Script

Will delete a preseed script from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deletePreseedScript(id, (error, data, response) => {
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


## executeImageBuild

> Model200Success executeImageBuild(id)

Run an Image Build

Running an image build is done asynchronously. This api will kick off the new execution and update the image build status.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.executeImageBuild(id, (error, data, response) => {
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


## getBootScript

> InlineResponse20052 getBootScript(id)

Get a Specific Boot Script

This endpoint retrieves a specific boot script.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getBootScript(id, (error, data, response) => {
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

[**InlineResponse20052**](InlineResponse20052.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getImageBuild

> InlineResponse20053 getImageBuild(id)

Get a Specific Image Build

This endpoint retrieves a specific image build.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getImageBuild(id, (error, data, response) => {
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

[**InlineResponse20053**](InlineResponse20053.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getImageBuildExecutions

> Object getImageBuildExecutions(id, opts)

List Image Build Executions

List all executions for an image build. This same info is also returned by the image build GET api, which returns the 100 most recent executions.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'buildNumber': 4, // Number | If specified will return an exact match on buildNumber
  'status': running // String | Filter by status
};
apiInstance.getImageBuildExecutions(id, opts, (error, data, response) => {
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
 **buildNumber** | **Number**| If specified will return an exact match on buildNumber | [optional] 
 **status** | **String**| Filter by status | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getPreseedScript

> InlineResponse200122 getPreseedScript(id)

Get a Specific Preseed Script

This endpoint retrieves a specific preseed script.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getPreseedScript(id, (error, data, response) => {
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

[**InlineResponse200122**](InlineResponse200122.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listBootScripts

> Object listBootScripts(opts)

Boot Scripts

Boot Scripts are used in the Image Builder service. See Image Builds.  This endpoint retrieves all boot scripts associated with the account. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listBootScripts(opts, (error, data, response) => {
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listImageBuilds

> Object listImageBuilds(opts)

Get All Image Builds

This endpoint retrieves all image builds associated with the account.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listImageBuilds(opts, (error, data, response) => {
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listPreseedScripts

> Object listPreseedScripts(opts)

Preseed Scripts

Preseed Scripts are used in the Image Builder service. See Image Builds.  This endpoint retrieves all preseed scripts associated with the account. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listPreseedScripts(opts, (error, data, response) => {
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateBootScript

> Object updateBootScript(id, opts)

Update a Boot Script

Update a Boot Script

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject84': new MorpheusApi.InlineObject84() // InlineObject84 | 
};
apiInstance.updateBootScript(id, opts, (error, data, response) => {
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
 **inlineObject84** | [**InlineObject84**](InlineObject84.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateImageBuild

> InlineResponse20054 updateImageBuild(id, opts)

Update an Image Build

Update an Image Build

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject86': new MorpheusApi.InlineObject86() // InlineObject86 | 
};
apiInstance.updateImageBuild(id, opts, (error, data, response) => {
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
 **inlineObject86** | [**InlineObject86**](InlineObject86.md)|  | [optional] 

### Return type

[**InlineResponse20054**](InlineResponse20054.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updatePreseedScript

> Object updatePreseedScript(id, opts)

Update a Preseed Script

Update a Preseed Script

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ImageBuildsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject197': new MorpheusApi.InlineObject197() // InlineObject197 | 
};
apiInstance.updatePreseedScript(id, opts, (error, data, response) => {
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
 **inlineObject197** | [**InlineObject197**](InlineObject197.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

