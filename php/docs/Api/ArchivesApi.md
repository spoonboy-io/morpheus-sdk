# OpenAPI\Client\ArchivesApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addArchiveBucket()**](ArchivesApi.md#addArchiveBucket) | **POST** /api/archives/buckets | Create an Archive Bucket
[**addArchiveFile()**](ArchivesApi.md#addArchiveFile) | **POST** /api/archives/buckets/{bucket}/files/{filepath} | Upload Archive File
[**addArchiveFileLink()**](ArchivesApi.md#addArchiveFileLink) | **POST** /api/archives/files/{id}/links | Create an Archive File Link
[**deleteArchiveBucket()**](ArchivesApi.md#deleteArchiveBucket) | **DELETE** /api/archives/buckets/{id} | Delete an Archive Bucket
[**deleteArchiveFile()**](ArchivesApi.md#deleteArchiveFile) | **DELETE** /api/archives/files/{id} | Delete Archive File
[**deleteArchiveFileLink()**](ArchivesApi.md#deleteArchiveFileLink) | **DELETE** /api/archives/files/{id}/links/{linkId} | Delete an Archive File Link
[**getArchiveBucket()**](ArchivesApi.md#getArchiveBucket) | **GET** /api/archives/buckets/{id} | Get a Specific Archive Bucket
[**getArchiveFile()**](ArchivesApi.md#getArchiveFile) | **GET** /api/archives/download/{bucket}/{filepath} | Download an Archive File
[**getArchiveFileDetail()**](ArchivesApi.md#getArchiveFileDetail) | **GET** /api/archives/files/{id} | Get Archive File Details
[**getArchiveFileLinks()**](ArchivesApi.md#getArchiveFileLinks) | **GET** /api/archives/files/{id}/links | Get Archive File Links
[**getArchivePublicFile()**](ArchivesApi.md#getArchivePublicFile) | **GET** /api/public-archives/download/{bucket}/{filepath} | Download a Public Archive File
[**getArchivePublicFileLink()**](ArchivesApi.md#getArchivePublicFileLink) | **GET** /api/public-archives/link | Download an Archive File Link
[**listArchiveBuckets()**](ArchivesApi.md#listArchiveBuckets) | **GET** /api/archives/buckets | Get All Archive Buckets
[**listArchiveFiles()**](ArchivesApi.md#listArchiveFiles) | **GET** /api/archives/buckets/{bucket}/files/{filepath} | Get All Archive Files
[**updateArchiveBucket()**](ArchivesApi.md#updateArchiveBucket) | **PUT** /api/archives/buckets/{id} | Update an Archive Bucket


## `addArchiveBucket()`

```php
addArchiveBucket($inline_object7): object
```

Create an Archive Bucket

Create an Archive Bucket

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object7 = new \OpenAPI\Client\Model\InlineObject7(); // \OpenAPI\Client\Model\InlineObject7

try {
    $result = $apiInstance->addArchiveBucket($inline_object7);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->addArchiveBucket: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object7** | [**\OpenAPI\Client\Model\InlineObject7**](../Model/InlineObject7.md)|  | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addArchiveFile()`

```php
addArchiveFile($bucket, $filepath, $filename, $file): object
```

Upload Archive File

Upload a file to the specified archive bucket and file path.  This will overwrite the file if it already exists.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$bucket = 'bucket_example'; // string | Bucket name
$filepath = /config/environments/; // string | The path to to search for files under. Default is the root directory /.
$filename = /path/to/file; // string | Specify a filename for archive file. The base filename of the uploaded file is the default.
$file = "/path/to/file.txt"; // \SplFileObject

try {
    $result = $apiInstance->addArchiveFile($bucket, $filepath, $filename, $file);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->addArchiveFile: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **string**| Bucket name |
 **filepath** | **string**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]
 **filename** | **string**| Specify a filename for archive file. The base filename of the uploaded file is the default. | [optional]
 **file** | **\SplFileObject****\SplFileObject**|  | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addArchiveFileLink()`

```php
addArchiveFileLink($id, $expire_seconds): Model200Success
```

Create an Archive File Link

This returns a secret token that can be used to download the file via a public url, without any other authentication or authorization. File links can be set to expire after a certain amount of time.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$expire_seconds = 600; // int | Time to live in seconds. 0 means do not expire.

try {
    $result = $apiInstance->addArchiveFileLink($id, $expire_seconds);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->addArchiveFileLink: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **expire_seconds** | **int**| Time to live in seconds. 0 means do not expire. | [optional] [default to 0]

### Return type

[**Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteArchiveBucket()`

```php
deleteArchiveBucket($id): \OpenAPI\Client\Model\Model200Success
```

Delete an Archive Bucket

Will delete an archive bucket from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteArchiveBucket($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->deleteArchiveBucket: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteArchiveFile()`

```php
deleteArchiveFile($id): \OpenAPI\Client\Model\Model200Success
```

Delete Archive File

Permanently delete a file or directory.  Deleting a directory will also delete all the files under it.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteArchiveFile($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->deleteArchiveFile: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteArchiveFileLink()`

```php
deleteArchiveFileLink($id, $link_id): \OpenAPI\Client\Model\Model200Success
```

Delete an Archive File Link

This will delete the link from the system, so it can no longer be used.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$link_id = 56; // int | The ID of the archive file link.

try {
    $result = $apiInstance->deleteArchiveFileLink($id, $link_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->deleteArchiveFileLink: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **link_id** | **int**| The ID of the archive file link. |

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getArchiveBucket()`

```php
getArchiveBucket($id): \OpenAPI\Client\Model\InlineResponse2004
```

Get a Specific Archive Bucket

This endpoint retrieves a specific archive bucket.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getArchiveBucket($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->getArchiveBucket: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse2004**](../Model/InlineResponse2004.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getArchiveFile()`

```php
getArchiveFile($bucket, $filepath)
```

Download an Archive File

Download the file as an authorized user with access to the bucket.  Downloading a directory will return a .zip file containing all files under it.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$bucket = 'bucket_example'; // string | Bucket name
$filepath = /config/environments/; // string | The path to to search for files under. Default is the root directory /.

try {
    $apiInstance->getArchiveFile($bucket, $filepath);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->getArchiveFile: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **string**| Bucket name |
 **filepath** | **string**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]

### Return type

void (empty response body)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getArchiveFileDetail()`

```php
getArchiveFileDetail($id): \OpenAPI\Client\Model\InlineResponse2005
```

Get Archive File Details

Get details about a specific archive file.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getArchiveFileDetail($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->getArchiveFileDetail: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse2005**](../Model/InlineResponse2005.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getArchiveFileLinks()`

```php
getArchiveFileLinks($id): object
```

Get Archive File Links

This endpoint retrieves the links that have been created for the specified file.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getArchiveFileLinks($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->getArchiveFileLinks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getArchivePublicFile()`

```php
getArchivePublicFile($bucket, $filepath)
```

Download a Public Archive File

Files in an archive bucket that has Public URL enabled can be downloaded via this endpoint without any authentication, anonymously.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$bucket = 'bucket_example'; // string | Bucket name
$filepath = /config/environments/; // string | The path to to search for files under. Default is the root directory /.

try {
    $apiInstance->getArchivePublicFile($bucket, $filepath);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->getArchivePublicFile: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **string**| Bucket name |
 **filepath** | **string**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]

### Return type

void (empty response body)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getArchivePublicFileLink()`

```php
getArchivePublicFileLink($s)
```

Download an Archive File Link

Download an archive file link.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$s = 45a214fce9a546b9; // string | The secret access token for the archive file being downloaded.

try {
    $apiInstance->getArchivePublicFileLink($s);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->getArchivePublicFileLink: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s** | **string**| The secret access token for the archive file being downloaded. |

### Return type

void (empty response body)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listArchiveBuckets()`

```php
listArchiveBuckets($name, $phrase): object
```

Get All Archive Buckets

This endpoint retrieves all archive buckets associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listArchiveBuckets($name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->listArchiveBuckets: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listArchiveFiles()`

```php
listArchiveFiles($bucket, $filepath, $name, $phrase, $full_tree): object
```

Get All Archive Files

This endpoint retrieves all files in an archive bucket under the specified `filePath`.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$bucket = 'bucket_example'; // string | Bucket name
$filepath = /config/environments/; // string | The path to to search for files under. Default is the root directory /.
$name = test-config.yaml; // string | If specified will return an exact match on file name
$phrase = test-%.yaml; // string | Search phrase for partial matches on file name, wildcard may be specified as %, eg. example-% This also searches for files under sub directories too.
$full_tree = false; // bool | Include files under sub directories too. This is always true when searching with phrase.

try {
    $result = $apiInstance->listArchiveFiles($bucket, $filepath, $name, $phrase, $full_tree);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->listArchiveFiles: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **string**| Bucket name |
 **filepath** | **string**| The path to to search for files under. Default is the root directory /. | [default to &#39;/&#39;]
 **name** | **string**| If specified will return an exact match on file name | [optional]
 **phrase** | **string**| Search phrase for partial matches on file name, wildcard may be specified as %, eg. example-% This also searches for files under sub directories too. | [optional]
 **full_tree** | **bool**| Include files under sub directories too. This is always true when searching with phrase. | [optional] [default to false]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateArchiveBucket()`

```php
updateArchiveBucket($id, $inline_object8): object
```

Update an Archive Bucket

Update an Archive Bucket

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ArchivesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object8 = new \OpenAPI\Client\Model\InlineObject8(); // \OpenAPI\Client\Model\InlineObject8

try {
    $result = $apiInstance->updateArchiveBucket($id, $inline_object8);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ArchivesApi->updateArchiveBucket: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object8** | [**\OpenAPI\Client\Model\InlineObject8**](../Model/InlineObject8.md)|  | [optional]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
