# MorpheusApi.DeploymentsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addDeploymentFile**](DeploymentsApi.md#addDeploymentFile) | **POST** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | Upload a Deployment File
[**addDeploymentVersion**](DeploymentsApi.md#addDeploymentVersion) | **POST** /api/deployments/{deploymentId}/versions | Create a new Deployment Version
[**addDeployments**](DeploymentsApi.md#addDeployments) | **POST** /api/deployments | Create a new Deployment
[**deleteDeployment**](DeploymentsApi.md#deleteDeployment) | **DELETE** /api/deployments/{deploymentId} | Delete a Deployment
[**deleteDeploymentFile**](DeploymentsApi.md#deleteDeploymentFile) | **DELETE** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | Delete a Deployment File
[**deleteDeploymentVersion**](DeploymentsApi.md#deleteDeploymentVersion) | **DELETE** /api/deployments/{deploymentId}/versions/{id} | Delete a Deployment Version
[**getDeployment**](DeploymentsApi.md#getDeployment) | **GET** /api/deployments/{deploymentId} | Get a Specific Deployment
[**getDeploymentVersion**](DeploymentsApi.md#getDeploymentVersion) | **GET** /api/deployments/{deploymentId}/versions/{id} | Get a Specific Deployment Version
[**listDeploymentFiles**](DeploymentsApi.md#listDeploymentFiles) | **GET** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | List Deployment Files
[**listDeploymentVersions**](DeploymentsApi.md#listDeploymentVersions) | **GET** /api/deployments/{deploymentId}/versions | Get All Versions For a Deployment
[**listDeployments**](DeploymentsApi.md#listDeployments) | **GET** /api/deployments | Get All Deployments
[**updateDeployment**](DeploymentsApi.md#updateDeployment) | **PUT** /api/deployments/{deploymentId} | Updating a Deployment
[**updateDeploymentVersion**](DeploymentsApi.md#updateDeploymentVersion) | **PUT** /api/deployments/{deploymentId}/versions/{id} | Updating a Deployment Version



## addDeploymentFile

> Model200Success addDeploymentFile(deploymentId, id, filepath, opts)

Upload a Deployment File

This endpoint will upload a file for a specific deployment version. This will overwrite the file if one with the same name exists already.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let deploymentId = 4; // Number | Deployment ID
let id = 1; // Number | Morpheus ID of the Object being referenced
let filepath = /config/environments/; // String | The path to to search for files under. Default is the root directory /.
let opts = {
  'file': "/path/to/file" // File | 
};
apiInstance.addDeploymentFile(deploymentId, id, filepath, opts, (error, data, response) => {
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
 **deploymentId** | **Number**| Deployment ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]
 **file** | **File**|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## addDeploymentVersion

> Object addDeploymentVersion(deploymentId, opts)

Create a new Deployment Version

This endpoint will create a new deployment version that is ready to have files uploaded to it. The default type is file, which has files directly uploaded via Morpheus. Alternatively, the type git or fetch can be used to just point to a repository or remote url.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let deploymentId = 4; // Number | Deployment ID
let opts = {
  'inlineObject69': new MorpheusApi.InlineObject69() // InlineObject69 | 
};
apiInstance.addDeploymentVersion(deploymentId, opts, (error, data, response) => {
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
 **deploymentId** | **Number**| Deployment ID | 
 **inlineObject69** | [**InlineObject69**](InlineObject69.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addDeployments

> Object addDeployments(opts)

Create a new Deployment

This endpoint will create a new deployment that is ready to have versions added to it.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let opts = {
  'inlineObject67': new MorpheusApi.InlineObject67() // InlineObject67 | 
};
apiInstance.addDeployments(opts, (error, data, response) => {
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
 **inlineObject67** | [**InlineObject67**](InlineObject67.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteDeployment

> Model200Success deleteDeployment(deploymentId)

Delete a Deployment

This endpoint will delete an existing deployment.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let deploymentId = 4; // Number | Deployment ID
apiInstance.deleteDeployment(deploymentId, (error, data, response) => {
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
 **deploymentId** | **Number**| Deployment ID | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteDeploymentFile

> Model200Success deleteDeploymentFile(deploymentId, id, filepath, opts)

Delete a Deployment File

This endpoint will delete an existing deployment file. To recursively delete a directory and all of its contents, the force parameter must be specified.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let deploymentId = 4; // Number | Deployment ID
let id = 1; // Number | Morpheus ID of the Object being referenced
let filepath = /config/environments/; // String | The path to to search for files under. Default is the root directory /.
let opts = {
  'force': on // String | Force Delete
};
apiInstance.deleteDeploymentFile(deploymentId, id, filepath, opts, (error, data, response) => {
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
 **deploymentId** | **Number**| Deployment ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]
 **force** | **String**| Force Delete | [optional] [default to &#39;off&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteDeploymentVersion

> Model200Success deleteDeploymentVersion(deploymentId, id)

Delete a Deployment Version

This endpoint will delete an existing deployment version.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let deploymentId = 4; // Number | Deployment ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteDeploymentVersion(deploymentId, id, (error, data, response) => {
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
 **deploymentId** | **Number**| Deployment ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getDeployment

> InlineResponse20038 getDeployment(deploymentId, opts)

Get a Specific Deployment

This endpoint retrieves a specific deployment. By default the 5 most recent versions are returned, more can be returned by specifying the maxVersions parameter.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let deploymentId = 4; // Number | Deployment ID
let opts = {
  'maxVersions': 6 // Number | Max number of recent versions to return.
};
apiInstance.getDeployment(deploymentId, opts, (error, data, response) => {
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
 **deploymentId** | **Number**| Deployment ID | 
 **maxVersions** | **Number**| Max number of recent versions to return. | [optional] 

### Return type

[**InlineResponse20038**](InlineResponse20038.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getDeploymentVersion

> InlineResponse20039 getDeploymentVersion(deploymentId, id)

Get a Specific Deployment Version

This endpoint retrieves a specific deployment version.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let deploymentId = 4; // Number | Deployment ID
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getDeploymentVersion(deploymentId, id, (error, data, response) => {
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
 **deploymentId** | **Number**| Deployment ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse20039**](InlineResponse20039.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listDeploymentFiles

> Object listDeploymentFiles(deploymentId, id, filepath, opts)

List Deployment Files

This endpoint returns a list of files for a specific deployment version. This only applies to deploy type file. Files are sorted alphabetically, with directories appearing at the beginning of the list.  The filepath parameter can be specified to search for specific files or directories. To list files under a directory, use a trailing / in the filepath parameter.  To list a specific file, provide it&#39;s full path. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let deploymentId = 4; // Number | Deployment ID
let id = 1; // Number | Morpheus ID of the Object being referenced
let filepath = /config/environments/; // String | The path to to search for files under. Default is the root directory /.
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'version': 5 // Number | Filter by version number (userVersion)
};
apiInstance.listDeploymentFiles(deploymentId, id, filepath, opts, (error, data, response) => {
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
 **deploymentId** | **Number**| Deployment ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **version** | **Number**| Filter by version number (userVersion) | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listDeploymentVersions

> Object listDeploymentVersions(deploymentId, opts)

Get All Versions For a Deployment

This endpoint returns a paginated list of versions for a specific deployment.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let deploymentId = 4; // Number | Deployment ID
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'version': 5, // Number | Filter by version number (userVersion)
  'type': file, // String | Filter by type (deployType), file, git, fetch
  'dateCreated': 2019-01-01, // String | Filter by dateCreated, the created timestamp is more recent or equal to the date specified
  'lastUpdated': 2019-03-06T17:52:29+0000 // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
};
apiInstance.listDeploymentVersions(deploymentId, opts, (error, data, response) => {
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
 **deploymentId** | **Number**| Deployment ID | 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **version** | **Number**| Filter by version number (userVersion) | [optional] 
 **type** | **String**| Filter by type (deployType), file, git, fetch | [optional] 
 **dateCreated** | **String**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional] 
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listDeployments

> Object listDeployments(opts)

Get All Deployments

This endpoint returns a paginated list of deployments.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'description': The desription of my cool object, // String | Filter by description, wildcard may be specified as %. eg. `example-%`
  'dateCreated': 2019-01-01, // String | Filter by dateCreated, the created timestamp is more recent or equal to the date specified
  'lastUpdated': 2019-03-06T17:52:29+0000 // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
};
apiInstance.listDeployments(opts, (error, data, response) => {
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
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **description** | **String**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional] 
 **dateCreated** | **String**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional] 
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateDeployment

> Object updateDeployment(deploymentId, opts)

Updating a Deployment

This endpoint will update an existing deployment.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let deploymentId = 4; // Number | Deployment ID
let opts = {
  'inlineObject68': new MorpheusApi.InlineObject68() // InlineObject68 | 
};
apiInstance.updateDeployment(deploymentId, opts, (error, data, response) => {
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
 **deploymentId** | **Number**| Deployment ID | 
 **inlineObject68** | [**InlineObject68**](InlineObject68.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateDeploymentVersion

> Object updateDeploymentVersion(deploymentId, id, opts)

Updating a Deployment Version

This endpoint will update an existing deployment version.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploymentsApi();
let deploymentId = 4; // Number | Deployment ID
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject70': new MorpheusApi.InlineObject70() // InlineObject70 | 
};
apiInstance.updateDeploymentVersion(deploymentId, id, opts, (error, data, response) => {
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
 **deploymentId** | **Number**| Deployment ID | 
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject70** | [**InlineObject70**](InlineObject70.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

