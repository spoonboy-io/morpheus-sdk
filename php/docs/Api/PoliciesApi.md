# OpenAPI\Client\PoliciesApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addPolicies()**](PoliciesApi.md#addPolicies) | **POST** /api/policies | Creates a Policy
[**addPoliciesCloud()**](PoliciesApi.md#addPoliciesCloud) | **POST** /api/zones/{cloudId}/policies | Creates a Policy for a Cloud
[**addPoliciesGroup()**](PoliciesApi.md#addPoliciesGroup) | **POST** /api/groups/{groupId}/policies | Creates a Policy for a Group
[**getPolicies()**](PoliciesApi.md#getPolicies) | **GET** /api/policies/{id} | Retrieves a Specific Policy
[**getPoliciesCloud()**](PoliciesApi.md#getPoliciesCloud) | **GET** /api/zones/{cloudId}/policies/{id} | Retrieves a Specific Policy for a Cloud
[**getPoliciesGroup()**](PoliciesApi.md#getPoliciesGroup) | **GET** /api/groups/{groupId}/policies/{id} | Retrieves a Specific Policy for a Group
[**listPolicies()**](PoliciesApi.md#listPolicies) | **GET** /api/policies | Retrieves all Policies
[**listPoliciesCloud()**](PoliciesApi.md#listPoliciesCloud) | **GET** /api/zones/{cloudId}/policies | Retrieves Policies for a Cloud
[**listPoliciesGroup()**](PoliciesApi.md#listPoliciesGroup) | **GET** /api/groups/{groupId}/policies | Retrieves Policies for a Group
[**listPolicyTypes()**](PoliciesApi.md#listPolicyTypes) | **GET** /api/policy-types | Retrieves all Policy Types
[**removePolicies()**](PoliciesApi.md#removePolicies) | **DELETE** /api/policies/{id} | Deletes a Policy
[**removePoliciesCloud()**](PoliciesApi.md#removePoliciesCloud) | **DELETE** /api/zones/{cloudId}/policies/{id} | Deletes a Policy for a Cloud
[**removePoliciesGroup()**](PoliciesApi.md#removePoliciesGroup) | **DELETE** /api/groups/{groupId}/policies/{id} | Deletes a Policy for a Group
[**updatePolicies()**](PoliciesApi.md#updatePolicies) | **PUT** /api/policies/{id} | Updates a Policy
[**updatePoliciesCloud()**](PoliciesApi.md#updatePoliciesCloud) | **PUT** /api/zones/{cloudId}/policies/{id} | Updates a Policy for a Cloud
[**updatePoliciesGroup()**](PoliciesApi.md#updatePoliciesGroup) | **PUT** /api/groups/{groupId}/policies/{id} | Updates a Policy for a Group


## `addPolicies()`

```php
addPolicies($inline_object184): object
```

Creates a Policy

Creates a policy.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object184 = new \OpenAPI\Client\Model\InlineObject184(); // \OpenAPI\Client\Model\InlineObject184

try {
    $result = $apiInstance->addPolicies($inline_object184);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->addPolicies: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object184** | [**\OpenAPI\Client\Model\InlineObject184**](../Model/InlineObject184.md)|  | [optional]

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

## `addPoliciesCloud()`

```php
addPoliciesCloud($cloud_id, $inline_object188): object
```

Creates a Policy for a Cloud

Creates a policy for a Cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cloud_id = 7; // int | The ID of the cloud
$inline_object188 = new \OpenAPI\Client\Model\InlineObject188(); // \OpenAPI\Client\Model\InlineObject188

try {
    $result = $apiInstance->addPoliciesCloud($cloud_id, $inline_object188);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->addPoliciesCloud: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_id** | **int**| The ID of the cloud |
 **inline_object188** | [**\OpenAPI\Client\Model\InlineObject188**](../Model/InlineObject188.md)|  | [optional]

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

## `addPoliciesGroup()`

```php
addPoliciesGroup($group_id, $inline_object186): object
```

Creates a Policy for a Group

Creates a policy for a Group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$group_id = 7; // int | The ID of the group
$inline_object186 = new \OpenAPI\Client\Model\InlineObject186(); // \OpenAPI\Client\Model\InlineObject186

try {
    $result = $apiInstance->addPoliciesGroup($group_id, $inline_object186);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->addPoliciesGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **int**| The ID of the group |
 **inline_object186** | [**\OpenAPI\Client\Model\InlineObject186**](../Model/InlineObject186.md)|  | [optional]

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

## `getPolicies()`

```php
getPolicies($id): object
```

Retrieves a Specific Policy

Retrieves a specific policy.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getPolicies($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->getPolicies: ', $e->getMessage(), PHP_EOL;
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

## `getPoliciesCloud()`

```php
getPoliciesCloud($cloud_id, $id): object
```

Retrieves a Specific Policy for a Cloud

Retrieves a specific policy for a Cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cloud_id = 7; // int | The ID of the cloud
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getPoliciesCloud($cloud_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->getPoliciesCloud: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_id** | **int**| The ID of the cloud |
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

## `getPoliciesGroup()`

```php
getPoliciesGroup($group_id, $id): object
```

Retrieves a Specific Policy for a Group

Retrieves a specific policy for a Group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$group_id = 7; // int | The ID of the group
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getPoliciesGroup($group_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->getPoliciesGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **int**| The ID of the group |
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

## `listPolicies()`

```php
listPolicies($max, $offset, $sort, $direction, $phrase, $name): object
```

Retrieves all Policies

Retrieves all policies.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%

try {
    $result = $apiInstance->listPolicies($max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->listPolicies: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

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

## `listPoliciesCloud()`

```php
listPoliciesCloud($cloud_id, $max, $offset, $sort, $direction, $phrase, $name): object
```

Retrieves Policies for a Cloud

Retrieves policies for a specific cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cloud_id = 7; // int | The ID of the cloud
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%

try {
    $result = $apiInstance->listPoliciesCloud($cloud_id, $max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->listPoliciesCloud: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_id** | **int**| The ID of the cloud |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

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

## `listPoliciesGroup()`

```php
listPoliciesGroup($group_id, $max, $offset, $sort, $direction, $phrase, $name): object
```

Retrieves Policies for a Group

Retrieves policies for a specific group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$group_id = 7; // int | The ID of the group
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%

try {
    $result = $apiInstance->listPoliciesGroup($group_id, $max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->listPoliciesGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **int**| The ID of the group |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

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

## `listPolicyTypes()`

```php
listPolicyTypes(): object
```

Retrieves all Policy Types

Retrieves all Policy Types

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->listPolicyTypes();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->listPolicyTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

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

## `removePolicies()`

```php
removePolicies($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Policy

Deletes a specified policy.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removePolicies($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->removePolicies: ', $e->getMessage(), PHP_EOL;
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

## `removePoliciesCloud()`

```php
removePoliciesCloud($cloud_id, $id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Policy for a Cloud

Deletes a specified policy for a Cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cloud_id = 7; // int | The ID of the cloud
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removePoliciesCloud($cloud_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->removePoliciesCloud: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_id** | **int**| The ID of the cloud |
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

## `removePoliciesGroup()`

```php
removePoliciesGroup($group_id, $id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Policy for a Group

Deletes a specified policy for a Group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$group_id = 7; // int | The ID of the group
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removePoliciesGroup($group_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->removePoliciesGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **int**| The ID of the group |
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

## `updatePolicies()`

```php
updatePolicies($id, $inline_object185): object
```

Updates a Policy

Updates a policy.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object185 = new \OpenAPI\Client\Model\InlineObject185(); // \OpenAPI\Client\Model\InlineObject185

try {
    $result = $apiInstance->updatePolicies($id, $inline_object185);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->updatePolicies: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object185** | [**\OpenAPI\Client\Model\InlineObject185**](../Model/InlineObject185.md)|  | [optional]

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

## `updatePoliciesCloud()`

```php
updatePoliciesCloud($cloud_id, $id, $inline_object189): object
```

Updates a Policy for a Cloud

Updates a policy for a Cloud.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cloud_id = 7; // int | The ID of the cloud
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object189 = new \OpenAPI\Client\Model\InlineObject189(); // \OpenAPI\Client\Model\InlineObject189

try {
    $result = $apiInstance->updatePoliciesCloud($cloud_id, $id, $inline_object189);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->updatePoliciesCloud: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_id** | **int**| The ID of the cloud |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object189** | [**\OpenAPI\Client\Model\InlineObject189**](../Model/InlineObject189.md)|  | [optional]

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

## `updatePoliciesGroup()`

```php
updatePoliciesGroup($group_id, $id, $inline_object187): object
```

Updates a Policy for a Group

Updates a policy for a Group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\PoliciesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$group_id = 7; // int | The ID of the group
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object187 = new \OpenAPI\Client\Model\InlineObject187(); // \OpenAPI\Client\Model\InlineObject187

try {
    $result = $apiInstance->updatePoliciesGroup($group_id, $id, $inline_object187);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PoliciesApi->updatePoliciesGroup: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **int**| The ID of the group |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object187** | [**\OpenAPI\Client\Model\InlineObject187**](../Model/InlineObject187.md)|  | [optional]

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
