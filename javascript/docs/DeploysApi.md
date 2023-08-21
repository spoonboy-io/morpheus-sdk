# MorpheusApi.DeploysApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addInstanceDeploy**](DeploysApi.md#addInstanceDeploy) | **POST** /api/instances/{id}/deploys | Deploy to an Instance
[**deletedeploy**](DeploysApi.md#deletedeploy) | **DELETE** /api/deploys/{id} | Delete a Deploy
[**getInstanceDeploys**](DeploysApi.md#getInstanceDeploys) | **GET** /api/instances/{id}/deploys | Get all Deploys for an Instance
[**listDeploys**](DeploysApi.md#listDeploys) | **GET** /api/deploys | Get all Deploys
[**runDeploy**](DeploysApi.md#runDeploy) | **POST** /api/deploys/{id}/deploy | Run a Deploy
[**updateDeploy**](DeploysApi.md#updateDeploy) | **PUT** /api/deploys/{id} | Update a Deploy



## addInstanceDeploy

> InlineResponse20040 addInstanceDeploy(id, opts)

Deploy to an Instance

This endpoint will deploy the specified deployment version to specified instance. The version to deploy can be identified with deploymentId and version or with versionId alone.  By default, the deployment is executed right away. To prevent this so that it can be run manually later on. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploysApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject92': new MorpheusApi.InlineObject92() // InlineObject92 | 
};
apiInstance.addInstanceDeploy(id, opts, (error, data, response) => {
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
 **inlineObject92** | [**InlineObject92**](InlineObject92.md)|  | [optional] 

### Return type

[**InlineResponse20040**](InlineResponse20040.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deletedeploy

> Model200Success deletedeploy(id)

Delete a Deploy

This endpoint will delete an archived instance deploy.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploysApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deletedeploy(id, (error, data, response) => {
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


## getInstanceDeploys

> Object getInstanceDeploys(id, opts)

Get all Deploys for an Instance

This endpoint retrieves all deploys for a specific instance.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploysApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'deploymentId': 5, // Number | Filter by deployment id
  'instanceName': instance1, // String | Filter by instance name
  'instanceId': 94, // Number | The Instance ID for Filtering
  'version': 5, // Number | Filter by version number (userVersion)
  'versionId': 5, // Number | Filter by deployment version id
  'createdById': 63, // Number | Filter by owner (user) id
  'deployType': file, // String | Filter by type (deployType), file, git, fetch
  'dateCreated': 2019-01-01, // String | Filter by dateCreated, the created timestamp is more recent or equal to the date specified
  'lastUpdated': 2019-03-06T17:52:29+0000, // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
  'deployDate': 2020-01-01, // String | Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified
  'status': deploying // String | Filter by status
};
apiInstance.getInstanceDeploys(id, opts, (error, data, response) => {
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
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] 
 **deploymentId** | **Number**| Filter by deployment id | [optional] 
 **instanceName** | **String**| Filter by instance name | [optional] 
 **instanceId** | **Number**| The Instance ID for Filtering | [optional] 
 **version** | **Number**| Filter by version number (userVersion) | [optional] 
 **versionId** | **Number**| Filter by deployment version id | [optional] 
 **createdById** | **Number**| Filter by owner (user) id | [optional] 
 **deployType** | **String**| Filter by type (deployType), file, git, fetch | [optional] 
 **dateCreated** | **String**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional] 
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 
 **deployDate** | **String**| Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified | [optional] 
 **status** | **String**| Filter by status | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listDeploys

> Object listDeploys(opts)

Get all Deploys

This endpoint retrieves all deploys.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploysApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'deploymentId': 5, // Number | Filter by deployment id
  'instanceName': instance1, // String | Filter by instance name
  'instanceId': 94, // Number | The Instance ID for Filtering
  'version': 5, // Number | Filter by version number (userVersion)
  'versionId': 5, // Number | Filter by deployment version id
  'createdById': 63, // Number | Filter by owner (user) id
  'deployType': file, // String | Filter by type (deployType), file, git, fetch
  'dateCreated': 2019-01-01, // String | Filter by dateCreated, the created timestamp is more recent or equal to the date specified
  'lastUpdated': 2019-03-06T17:52:29+0000, // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
  'deployDate': 2020-01-01, // String | Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified
  'status': deploying // String | Filter by status
};
apiInstance.listDeploys(opts, (error, data, response) => {
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
 **deploymentId** | **Number**| Filter by deployment id | [optional] 
 **instanceName** | **String**| Filter by instance name | [optional] 
 **instanceId** | **Number**| The Instance ID for Filtering | [optional] 
 **version** | **Number**| Filter by version number (userVersion) | [optional] 
 **versionId** | **Number**| Filter by deployment version id | [optional] 
 **createdById** | **Number**| Filter by owner (user) id | [optional] 
 **deployType** | **String**| Filter by type (deployType), file, git, fetch | [optional] 
 **dateCreated** | **String**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional] 
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 
 **deployDate** | **String**| Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified | [optional] 
 **status** | **String**| Filter by status | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## runDeploy

> InlineResponse20040 runDeploy(id, opts)

Run a Deploy

This endpoint will run an existing instance deploy. This is for running a new staged deploy or to rollback to previous version by re-running a deploy that is archived.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploysApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject73': new MorpheusApi.InlineObject73() // InlineObject73 | 
};
apiInstance.runDeploy(id, opts, (error, data, response) => {
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
 **inlineObject73** | [**InlineObject73**](InlineObject73.md)|  | [optional] 

### Return type

[**InlineResponse20040**](InlineResponse20040.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateDeploy

> InlineResponse20040 updateDeploy(id, opts)

Update a Deploy

This endpoint will update an existing deploy. This is typically only needed to change settings on a deploy that is staged, before it is run.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.DeploysApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject72': new MorpheusApi.InlineObject72() // InlineObject72 | 
};
apiInstance.updateDeploy(id, opts, (error, data, response) => {
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
 **inlineObject72** | [**InlineObject72**](InlineObject72.md)|  | [optional] 

### Return type

[**InlineResponse20040**](InlineResponse20040.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

