# MorpheusApi.BackupsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addBackupJobs**](BackupsApi.md#addBackupJobs) | **POST** /api/backups/jobs | Creates a Backup Job
[**addBackups**](BackupsApi.md#addBackups) | **POST** /api/backups | Creates a Backup
[**executeBackupJobs**](BackupsApi.md#executeBackupJobs) | **POST** /api/backups/jobs/{id}/execute | Executes a Backup Job
[**executeBackups**](BackupsApi.md#executeBackups) | **POST** /api/backups/{id}/execute | Executes a Backup
[**getBackupJobs**](BackupsApi.md#getBackupJobs) | **GET** /api/backups/jobs/{id} | Retrieves a Specific Backup Job
[**getBackupRestores**](BackupsApi.md#getBackupRestores) | **GET** /api/backups/restores/{id} | Retrieves a Specific Backup Restore
[**getBackupResults**](BackupsApi.md#getBackupResults) | **GET** /api/backups/results/{id} | Retrieves a Specific Backup Result
[**getBackups**](BackupsApi.md#getBackups) | **GET** /api/backups/{id} | Retrieves a Specific Backup
[**listBackupJobs**](BackupsApi.md#listBackupJobs) | **GET** /api/backups/jobs | Retrieves all Backup Jobs
[**listBackupRestores**](BackupsApi.md#listBackupRestores) | **GET** /api/backups/restores | Retrieves all Backup Restores
[**listBackupResults**](BackupsApi.md#listBackupResults) | **GET** /api/backups/results | Retrieves all Backup Results
[**listBackups**](BackupsApi.md#listBackups) | **GET** /api/backups | Retrieves all Backups
[**removeBackupJobs**](BackupsApi.md#removeBackupJobs) | **DELETE** /api/backups/jobs/{id} | Deletes a Backup Job
[**removeBackupRestores**](BackupsApi.md#removeBackupRestores) | **DELETE** /api/backups/restores/{id} | Deletes a Backup Restore
[**removeBackupResults**](BackupsApi.md#removeBackupResults) | **DELETE** /api/backups/results/{id} | Deletes a Backup Result
[**removeBackups**](BackupsApi.md#removeBackups) | **DELETE** /api/backups/{id} | Deletes a Backup
[**updateBackupJobs**](BackupsApi.md#updateBackupJobs) | **PUT** /api/backups/jobs/{id} | Updates a Backup Job
[**updateBackups**](BackupsApi.md#updateBackups) | **PUT** /api/backups/{id} | Updates a Backup



## addBackupJobs

> Object addBackupJobs(opts)

Creates a Backup Job

This endpoint allows creating a Backup Job. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let opts = {
  'inlineObject18': new MorpheusApi.InlineObject18() // InlineObject18 | 
};
apiInstance.addBackupJobs(opts, (error, data, response) => {
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
 **inlineObject18** | [**InlineObject18**](InlineObject18.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addBackups

> Object addBackups(opts)

Creates a Backup

This endpoint allows creating a Backup. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let opts = {
  'inlineObject16': new MorpheusApi.InlineObject16() // InlineObject16 | 
};
apiInstance.addBackups(opts, (error, data, response) => {
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
 **inlineObject16** | [**InlineObject16**](InlineObject16.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## executeBackupJobs

> Object executeBackupJobs(id, opts)

Executes a Backup Job

Executes a backup job. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'body': null // Object | 
};
apiInstance.executeBackupJobs(id, opts, (error, data, response) => {
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
 **body** | **Object**|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## executeBackups

> Object executeBackups(id, opts)

Executes a Backup

Executes a backup. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'body': null // Object | 
};
apiInstance.executeBackups(id, opts, (error, data, response) => {
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
 **body** | **Object**|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getBackupJobs

> InlineResponse20011 getBackupJobs(id)

Retrieves a Specific Backup Job

Retrieves a specific backup job. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getBackupJobs(id, (error, data, response) => {
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

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getBackupRestores

> InlineResponse20013 getBackupRestores(id)

Retrieves a Specific Backup Restore

Retrieves a specific backup restore. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getBackupRestores(id, (error, data, response) => {
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

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getBackupResults

> InlineResponse20012 getBackupResults(id)

Retrieves a Specific Backup Result

Retrieves a specific backup result. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getBackupResults(id, (error, data, response) => {
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

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getBackups

> InlineResponse20010 getBackups(id)

Retrieves a Specific Backup

Retrieves a specific backup. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getBackups(id, (error, data, response) => {
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

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listBackupJobs

> Object listBackupJobs(opts)

Retrieves all Backup Jobs

Retrieves all backup jobs. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'code': azr, // String | If specified will return an exact match on code
  'externalId': "externalId_example" // String | Filter by External ID
};
apiInstance.listBackupJobs(opts, (error, data, response) => {
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
 **externalId** | **String**| Filter by External ID | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listBackupRestores

> Object listBackupRestores(opts)

Retrieves all Backup Restores

Retrieves all backup restores. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "sort_example", // String | Sort order, the name of the property to sort by. The default sort order is `startDate` descending
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'name': test-backup, // String | Filter by backup name
  'phrase': test, // String | Search phrase for partial matches on backup name, wildcard may be specified as %, eg. example-% 
  'backupId': 789, // Number | Filter by backup ID
  'backupResultId': "backupResultId_example", // String | Filter by backup result ID
  'containerId': 789 // Number | Filter by container ID
};
apiInstance.listBackupRestores(opts, (error, data, response) => {
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
 **sort** | **String**| Sort order, the name of the property to sort by. The default sort order is &#x60;startDate&#x60; descending | [optional] 
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **name** | **String**| Filter by backup name | [optional] 
 **phrase** | **String**| Search phrase for partial matches on backup name, wildcard may be specified as %, eg. example-%  | [optional] 
 **backupId** | **Number**| Filter by backup ID | [optional] 
 **backupResultId** | **String**| Filter by backup result ID | [optional] 
 **containerId** | **Number**| Filter by container ID | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listBackupResults

> Object listBackupResults(opts)

Retrieves all Backup Results

Retrieves all backup results. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'name': test-backup, // String | Filter by backupName
  'phrase': test, // String | Search phrase for partial matches on backupName, wildcard may be specified as %, eg. example-% 
  'backupId': 789, // Number | Filter by backup ID
  'backupSetId': "backupSetId_example", // String | Filter by backupSetId
  'instanceId': 789, // Number | Filter by instance ID
  'containerId': 789, // Number | Filter by container ID
  'serverId': 789 // Number | Filter by server ID
};
apiInstance.listBackupResults(opts, (error, data, response) => {
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
 **name** | **String**| Filter by backupName | [optional] 
 **phrase** | **String**| Search phrase for partial matches on backupName, wildcard may be specified as %, eg. example-%  | [optional] 
 **backupId** | **Number**| Filter by backup ID | [optional] 
 **backupSetId** | **String**| Filter by backupSetId | [optional] 
 **instanceId** | **Number**| Filter by instance ID | [optional] 
 **containerId** | **Number**| Filter by container ID | [optional] 
 **serverId** | **Number**| Filter by server ID | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listBackups

> Object listBackups(opts)

Retrieves all Backups

Retrieves all backups. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'name': example-% // String | Filter by name, wildcard may be specified as %, eg. example-%
};
apiInstance.listBackups(opts, (error, data, response) => {
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


## removeBackupJobs

> Model200Success removeBackupJobs(id)

Deletes a Backup Job

Deletes a specified backup job. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeBackupJobs(id, (error, data, response) => {
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


## removeBackupRestores

> Model200Success removeBackupRestores(id)

Deletes a Backup Restore

Deletes a specified backup restore. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeBackupRestores(id, (error, data, response) => {
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


## removeBackupResults

> Model200Success removeBackupResults(id)

Deletes a Backup Result

Deletes a specified backup result. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeBackupResults(id, (error, data, response) => {
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


## removeBackups

> Model200Success removeBackups(id)

Deletes a Backup

Deletes a specified backup. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.removeBackups(id, (error, data, response) => {
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


## updateBackupJobs

> Object updateBackupJobs(id, opts)

Updates a Backup Job

Updates a backup job. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject19': new MorpheusApi.InlineObject19() // InlineObject19 | 
};
apiInstance.updateBackupJobs(id, opts, (error, data, response) => {
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
 **inlineObject19** | [**InlineObject19**](InlineObject19.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateBackups

> Object updateBackups(id, opts)

Updates a Backup

Updates a backup. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.BackupsApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject17': new MorpheusApi.InlineObject17() // InlineObject17 | 
};
apiInstance.updateBackups(id, opts, (error, data, response) => {
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
 **inlineObject17** | [**InlineObject17**](InlineObject17.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

