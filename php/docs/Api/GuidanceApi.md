# OpenAPI\Client\GuidanceApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**executeGuidances()**](GuidanceApi.md#executeGuidances) | **PUT** /api/guidance/{id}/execute | Executes a Specific Guidance Recommendation
[**getGuidanceStats()**](GuidanceApi.md#getGuidanceStats) | **GET** /api/guidance/stats | Retrieves Guidance Stats
[**getGuidanceTypes()**](GuidanceApi.md#getGuidanceTypes) | **GET** /api/guidance/types | Retrieves Guidance Types
[**getGuidances()**](GuidanceApi.md#getGuidances) | **GET** /api/guidance/{id} | Retrieves a Specific Guidance Recommendation
[**ignoreGuidances()**](GuidanceApi.md#ignoreGuidances) | **PUT** /api/guidance/{id}/ignore | Ignores a Specific Guidance Recommendation
[**listGuidances()**](GuidanceApi.md#listGuidances) | **GET** /api/guidance | Retrieves all Guidance Recommendations


## `executeGuidances()`

```php
executeGuidances($id): \OpenAPI\Client\Model\Model200Success
```

Executes a Specific Guidance Recommendation

Executes a specific guidance recommendation.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\GuidanceApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->executeGuidances($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GuidanceApi->executeGuidances: ', $e->getMessage(), PHP_EOL;
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

## `getGuidanceStats()`

```php
getGuidanceStats(): \OpenAPI\Client\Model\InlineResponse20045
```

Retrieves Guidance Stats

This endpoint retrieves a summary of actionable guidance.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\GuidanceApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->getGuidanceStats();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GuidanceApi->getGuidanceStats: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse20045**](../Model/InlineResponse20045.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getGuidanceTypes()`

```php
getGuidanceTypes(): \OpenAPI\Client\Model\InlineResponse20046
```

Retrieves Guidance Types

This endpoint retrieves all guidance types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\GuidanceApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->getGuidanceTypes();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GuidanceApi->getGuidanceTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse20046**](../Model/InlineResponse20046.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getGuidances()`

```php
getGuidances($id): \OpenAPI\Client\Model\InlineResponse20044
```

Retrieves a Specific Guidance Recommendation

Retrieves a specific guidance recommendation.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\GuidanceApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getGuidances($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GuidanceApi->getGuidances: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20044**](../Model/InlineResponse20044.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `ignoreGuidances()`

```php
ignoreGuidances($id): \OpenAPI\Client\Model\Model200Success
```

Ignores a Specific Guidance Recommendation

Ignores a specific guidance recommendation.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\GuidanceApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->ignoreGuidances($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GuidanceApi->ignoreGuidances: ', $e->getMessage(), PHP_EOL;
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

## `listGuidances()`

```php
listGuidances($max, $offset, $sort, $direction, $phrase, $severity, $type, $state): object
```

Retrieves all Guidance Recommendations

Retrieves all guidance recommendations.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\GuidanceApi(
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
$severity = 'severity_example'; // string | Filter by severity
$type = 'type_example'; // string | Filter by Guidance Type.  See `Retrieves Guidance Types` for most up to date list of types.
$state = 'state_example'; // string | Filter by state, restricts query to only load discoveries by state

try {
    $result = $apiInstance->listGuidances($max, $offset, $sort, $direction, $phrase, $severity, $type, $state);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GuidanceApi->listGuidances: ', $e->getMessage(), PHP_EOL;
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
 **severity** | **string**| Filter by severity | [optional]
 **type** | **string**| Filter by Guidance Type.  See &#x60;Retrieves Guidance Types&#x60; for most up to date list of types. | [optional]
 **state** | **string**| Filter by state, restricts query to only load discoveries by state | [optional]

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
