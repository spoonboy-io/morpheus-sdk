# OpenAPI\Client\SecurityGroupsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addSecurityGroupLocations()**](SecurityGroupsApi.md#addSecurityGroupLocations) | **POST** /api/security-groups/{id}/locations | Creates a Security Group Location
[**addSecurityGroupRules()**](SecurityGroupsApi.md#addSecurityGroupRules) | **POST** /api/security-groups/{id}/rules | Creates a Security Group Rule
[**addSecurityGroups()**](SecurityGroupsApi.md#addSecurityGroups) | **POST** /api/security-groups | Creates a Security Group
[**getSecurityGroupRules()**](SecurityGroupsApi.md#getSecurityGroupRules) | **GET** /api/security-groups/{id}/rules/{sgId} | Retrieves a Specific Security Group Rule
[**getSecurityGroups()**](SecurityGroupsApi.md#getSecurityGroups) | **GET** /api/security-groups/{id} | Retrieves a Specific Security Group
[**listSecurityGroupRules()**](SecurityGroupsApi.md#listSecurityGroupRules) | **GET** /api/security-groups/{id}/rules | Retrieves all Security Group Rules
[**listSecurityGroups()**](SecurityGroupsApi.md#listSecurityGroups) | **GET** /api/security-groups | Retrieves all Security Groups
[**removeSecurityGroupLocations()**](SecurityGroupsApi.md#removeSecurityGroupLocations) | **DELETE** /api/security-groups/{id}/locations/{locationId} | Deletes a Security Group Location
[**removeSecurityGroupRules()**](SecurityGroupsApi.md#removeSecurityGroupRules) | **DELETE** /api/security-groups/{id}/rules/{sgId} | Deletes a Security Group Rule
[**removeSecurityGroups()**](SecurityGroupsApi.md#removeSecurityGroups) | **DELETE** /api/security-groups/{id} | Deletes a Security Group
[**updateSecurityGroupRules()**](SecurityGroupsApi.md#updateSecurityGroupRules) | **PUT** /api/security-groups/{id}/rules/{sgId} | Updates a Security Group Rule
[**updateSecurityGroups()**](SecurityGroupsApi.md#updateSecurityGroups) | **PUT** /api/security-groups/{id} | Updating a Security Group


## `addSecurityGroupLocations()`

```php
addSecurityGroupLocations($id, $inline_object215): object
```

Creates a Security Group Location

Creates a security group location.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityGroupsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object215 = new \OpenAPI\Client\Model\InlineObject215(); // \OpenAPI\Client\Model\InlineObject215

try {
    $result = $apiInstance->addSecurityGroupLocations($id, $inline_object215);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityGroupsApi->addSecurityGroupLocations: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object215** | [**\OpenAPI\Client\Model\InlineObject215**](../Model/InlineObject215.md)|  | [optional]

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

## `addSecurityGroupRules()`

```php
addSecurityGroupRules($id, $inline_object216): object
```

Creates a Security Group Rule

Creates a security group rule on specified security group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityGroupsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object216 = new \OpenAPI\Client\Model\InlineObject216(); // \OpenAPI\Client\Model\InlineObject216

try {
    $result = $apiInstance->addSecurityGroupRules($id, $inline_object216);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityGroupsApi->addSecurityGroupRules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object216** | [**\OpenAPI\Client\Model\InlineObject216**](../Model/InlineObject216.md)|  | [optional]

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

## `addSecurityGroups()`

```php
addSecurityGroups($inline_object213): \OpenAPI\Client\Model\InlineResponse200133
```

Creates a Security Group

Creates a security group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityGroupsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object213 = new \OpenAPI\Client\Model\InlineObject213(); // \OpenAPI\Client\Model\InlineObject213

try {
    $result = $apiInstance->addSecurityGroups($inline_object213);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityGroupsApi->addSecurityGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object213** | [**\OpenAPI\Client\Model\InlineObject213**](../Model/InlineObject213.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200133**](../Model/InlineResponse200133.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getSecurityGroupRules()`

```php
getSecurityGroupRules($id, $sg_id): \OpenAPI\Client\Model\InlineResponse200135
```

Retrieves a Specific Security Group Rule

Retrieves a specific security group rule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityGroupsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$sg_id = 2352; // float | Morpheus ID of the security group rule being referenced

try {
    $result = $apiInstance->getSecurityGroupRules($id, $sg_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityGroupsApi->getSecurityGroupRules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **sg_id** | **float**| Morpheus ID of the security group rule being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200135**](../Model/InlineResponse200135.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getSecurityGroups()`

```php
getSecurityGroups($id): \OpenAPI\Client\Model\InlineResponse200134
```

Retrieves a Specific Security Group

Retrieves a specific security group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityGroupsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getSecurityGroups($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityGroupsApi->getSecurityGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200134**](../Model/InlineResponse200134.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listSecurityGroupRules()`

```php
listSecurityGroupRules($id, $max, $offset, $sort, $direction, $phrase, $name): object
```

Retrieves all Security Group Rules

Retrieves all security group rules for specified security group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityGroupsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%

try {
    $result = $apiInstance->listSecurityGroupRules($id, $max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityGroupsApi->listSecurityGroupRules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
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

## `listSecurityGroups()`

```php
listSecurityGroups($max, $offset, $sort, $direction, $phrase, $name): object
```

Retrieves all Security Groups

This endpoint retrieves all security groups and their JSON encoded configuration attributes.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityGroupsApi(
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
    $result = $apiInstance->listSecurityGroups($max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityGroupsApi->listSecurityGroups: ', $e->getMessage(), PHP_EOL;
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

## `removeSecurityGroupLocations()`

```php
removeSecurityGroupLocations($id, $location_id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Security Group Location

Deletes a security group location.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityGroupsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$location_id = 2; // float | The ID of the location

try {
    $result = $apiInstance->removeSecurityGroupLocations($id, $location_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityGroupsApi->removeSecurityGroupLocations: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **location_id** | **float**| The ID of the location |

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

## `removeSecurityGroupRules()`

```php
removeSecurityGroupRules($id, $sg_id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Security Group Rule

Deletes a security group rule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityGroupsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$sg_id = 2352; // float | Morpheus ID of the security group rule being referenced

try {
    $result = $apiInstance->removeSecurityGroupRules($id, $sg_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityGroupsApi->removeSecurityGroupRules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **sg_id** | **float**| Morpheus ID of the security group rule being referenced |

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

## `removeSecurityGroups()`

```php
removeSecurityGroups($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Security Group

Deletes a specified security group.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityGroupsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeSecurityGroups($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityGroupsApi->removeSecurityGroups: ', $e->getMessage(), PHP_EOL;
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

## `updateSecurityGroupRules()`

```php
updateSecurityGroupRules($id, $sg_id, $inline_object217): object
```

Updates a Security Group Rule

Updates a security group rule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityGroupsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$sg_id = 2352; // float | Morpheus ID of the security group rule being referenced
$inline_object217 = new \OpenAPI\Client\Model\InlineObject217(); // \OpenAPI\Client\Model\InlineObject217

try {
    $result = $apiInstance->updateSecurityGroupRules($id, $sg_id, $inline_object217);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityGroupsApi->updateSecurityGroupRules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **sg_id** | **float**| Morpheus ID of the security group rule being referenced |
 **inline_object217** | [**\OpenAPI\Client\Model\InlineObject217**](../Model/InlineObject217.md)|  | [optional]

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

## `updateSecurityGroups()`

```php
updateSecurityGroups($id, $inline_object214): \OpenAPI\Client\Model\InlineResponse200133
```

Updating a Security Group

Updating a Security Group

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SecurityGroupsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object214 = new \OpenAPI\Client\Model\InlineObject214(); // \OpenAPI\Client\Model\InlineObject214

try {
    $result = $apiInstance->updateSecurityGroups($id, $inline_object214);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SecurityGroupsApi->updateSecurityGroups: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object214** | [**\OpenAPI\Client\Model\InlineObject214**](../Model/InlineObject214.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200133**](../Model/InlineResponse200133.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
