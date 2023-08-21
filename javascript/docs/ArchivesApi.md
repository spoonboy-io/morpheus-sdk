# MorpheusApi.ArchivesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addArchiveBucket**](ArchivesApi.md#addArchiveBucket) | **POST** /api/archives/buckets | Create an Archive Bucket
[**addArchiveFile**](ArchivesApi.md#addArchiveFile) | **POST** /api/archives/buckets/{bucket}/files/{filepath} | Upload Archive File
[**addArchiveFileLink**](ArchivesApi.md#addArchiveFileLink) | **POST** /api/archives/files/{id}/links | Create an Archive File Link
[**deleteArchiveBucket**](ArchivesApi.md#deleteArchiveBucket) | **DELETE** /api/archives/buckets/{id} | Delete an Archive Bucket
[**deleteArchiveFile**](ArchivesApi.md#deleteArchiveFile) | **DELETE** /api/archives/files/{id} | Delete Archive File
[**deleteArchiveFileLink**](ArchivesApi.md#deleteArchiveFileLink) | **DELETE** /api/archives/files/{id}/links/{linkId} | Delete an Archive File Link
[**getArchiveBucket**](ArchivesApi.md#getArchiveBucket) | **GET** /api/archives/buckets/{id} | Get a Specific Archive Bucket
[**getArchiveFile**](ArchivesApi.md#getArchiveFile) | **GET** /api/archives/download/{bucket}/{filepath} | Download an Archive File
[**getArchiveFileDetail**](ArchivesApi.md#getArchiveFileDetail) | **GET** /api/archives/files/{id} | Get Archive File Details
[**getArchiveFileLinks**](ArchivesApi.md#getArchiveFileLinks) | **GET** /api/archives/files/{id}/links | Get Archive File Links
[**getArchivePublicFile**](ArchivesApi.md#getArchivePublicFile) | **GET** /api/public-archives/download/{bucket}/{filepath} | Download a Public Archive File
[**getArchivePublicFileLink**](ArchivesApi.md#getArchivePublicFileLink) | **GET** /api/public-archives/link | Download an Archive File Link
[**listArchiveBuckets**](ArchivesApi.md#listArchiveBuckets) | **GET** /api/archives/buckets | Get All Archive Buckets
[**listArchiveFiles**](ArchivesApi.md#listArchiveFiles) | **GET** /api/archives/buckets/{bucket}/files/{filepath} | Get All Archive Files
[**updateArchiveBucket**](ArchivesApi.md#updateArchiveBucket) | **PUT** /api/archives/buckets/{id} | Update an Archive Bucket



## addArchiveBucket

> Object addArchiveBucket(opts)

Create an Archive Bucket

Create an Archive Bucket

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let opts = {
  'inlineObject7': new MorpheusApi.InlineObject7() // InlineObject7 | 
};
apiInstance.addArchiveBucket(opts, (error, data, response) => {
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
 **inlineObject7** | [**InlineObject7**](InlineObject7.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## addArchiveFile

> Object addArchiveFile(bucket, filepath, opts)

Upload Archive File

Upload a file to the specified archive bucket and file path.  This will overwrite the file if it already exists. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let bucket = "bucket_example"; // String | Bucket name
let filepath = /config/environments/; // String | The path to to search for files under. Default is the root directory /.
let opts = {
  'filename': /path/to/file, // String | Specify a filename for archive file. The base filename of the uploaded file is the default.
  'file': "/path/to/file" // File | 
};
apiInstance.addArchiveFile(bucket, filepath, opts, (error, data, response) => {
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
 **bucket** | **String**| Bucket name | 
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]
 **filename** | **String**| Specify a filename for archive file. The base filename of the uploaded file is the default. | [optional] 
 **file** | **File**|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## addArchiveFileLink

> Model200Success addArchiveFileLink(id, opts)

Create an Archive File Link

This returns a secret token that can be used to download the file via a public url, without any other authentication or authorization. File links can be set to expire after a certain amount of time.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'expireSeconds': 600 // Number | Time to live in seconds. 0 means do not expire.
};
apiInstance.addArchiveFileLink(id, opts, (error, data, response) => {
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
 **expireSeconds** | **Number**| Time to live in seconds. 0 means do not expire. | [optional] [default to 0]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteArchiveBucket

> Model200Success deleteArchiveBucket(id)

Delete an Archive Bucket

Will delete an archive bucket from the system and make it no longer usable.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteArchiveBucket(id, (error, data, response) => {
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


## deleteArchiveFile

> Model200Success deleteArchiveFile(id)

Delete Archive File

Permanently delete a file or directory.  Deleting a directory will also delete all the files under it. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteArchiveFile(id, (error, data, response) => {
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


## deleteArchiveFileLink

> Model200Success deleteArchiveFileLink(id, linkId)

Delete an Archive File Link

This will delete the link from the system, so it can no longer be used.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let linkId = 789; // Number | The ID of the archive file link.
apiInstance.deleteArchiveFileLink(id, linkId, (error, data, response) => {
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
 **linkId** | **Number**| The ID of the archive file link. | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getArchiveBucket

> InlineResponse2004 getArchiveBucket(id)

Get a Specific Archive Bucket

This endpoint retrieves a specific archive bucket.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getArchiveBucket(id, (error, data, response) => {
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

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getArchiveFile

> getArchiveFile(bucket, filepath)

Download an Archive File

Download the file as an authorized user with access to the bucket.  Downloading a directory will return a .zip file containing all files under it. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let bucket = "bucket_example"; // String | Bucket name
let filepath = /config/environments/; // String | The path to to search for files under. Default is the root directory /.
apiInstance.getArchiveFile(bucket, filepath, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **String**| Bucket name | 
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]

### Return type

null (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getArchiveFileDetail

> InlineResponse2005 getArchiveFileDetail(id)

Get Archive File Details

Get details about a specific archive file.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getArchiveFileDetail(id, (error, data, response) => {
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

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getArchiveFileLinks

> Object getArchiveFileLinks(id)

Get Archive File Links

This endpoint retrieves the links that have been created for the specified file.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getArchiveFileLinks(id, (error, data, response) => {
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


## getArchivePublicFile

> getArchivePublicFile(bucket, filepath)

Download a Public Archive File

Files in an archive bucket that has Public URL enabled can be downloaded via this endpoint without any authentication, anonymously.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let bucket = "bucket_example"; // String | Bucket name
let filepath = /config/environments/; // String | The path to to search for files under. Default is the root directory /.
apiInstance.getArchivePublicFile(bucket, filepath, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **String**| Bucket name | 
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]

### Return type

null (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getArchivePublicFileLink

> getArchivePublicFileLink(s)

Download an Archive File Link

Download an archive file link.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let s = 45a214fce9a546b9; // String | The secret access token for the archive file being downloaded.
apiInstance.getArchivePublicFileLink(s, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s** | **String**| The secret access token for the archive file being downloaded. | 

### Return type

null (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listArchiveBuckets

> Object listArchiveBuckets(opts)

Get All Archive Buckets

This endpoint retrieves all archive buckets associated with the account.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let opts = {
  'name': example-%, // String | Filter by name, wildcard may be specified as %, eg. example-%
  'phrase': "phrase_example" // String | Search phrase for partial matches on name or description
};
apiInstance.listArchiveBuckets(opts, (error, data, response) => {
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


## listArchiveFiles

> Object listArchiveFiles(bucket, filepath, opts)

Get All Archive Files

This endpoint retrieves all files in an archive bucket under the specified &#x60;filePath&#x60;.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let bucket = "bucket_example"; // String | Bucket name
let filepath = /config/environments/; // String | The path to to search for files under. Default is the root directory /.
let opts = {
  'name': test-config.yaml, // String | If specified will return an exact match on file name
  'phrase': test-%.yaml, // String | Search phrase for partial matches on file name, wildcard may be specified as %, eg. example-% This also searches for files under sub directories too. 
  'fullTree': false // Boolean | Include files under sub directories too. This is always true when searching with phrase.
};
apiInstance.listArchiveFiles(bucket, filepath, opts, (error, data, response) => {
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
 **bucket** | **String**| Bucket name | 
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]
 **name** | **String**| If specified will return an exact match on file name | [optional] 
 **phrase** | **String**| Search phrase for partial matches on file name, wildcard may be specified as %, eg. example-% This also searches for files under sub directories too.  | [optional] 
 **fullTree** | **Boolean**| Include files under sub directories too. This is always true when searching with phrase. | [optional] [default to false]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateArchiveBucket

> Object updateArchiveBucket(id, opts)

Update an Archive Bucket

Update an Archive Bucket

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.ArchivesApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject8': new MorpheusApi.InlineObject8() // InlineObject8 | 
};
apiInstance.updateArchiveBucket(id, opts, (error, data, response) => {
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
 **inlineObject8** | [**InlineObject8**](InlineObject8.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

