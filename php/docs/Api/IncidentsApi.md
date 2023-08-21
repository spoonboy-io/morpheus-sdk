# OpenAPI\Client\IncidentsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addIncident()**](IncidentsApi.md#addIncident) | **POST** /api/monitoring/incidents | Create a New Incident
[**deleteIncidents()**](IncidentsApi.md#deleteIncidents) | **DELETE** /api/monitoring/incidents/{id} | Close a Specific Incident
[**getIncidents()**](IncidentsApi.md#getIncidents) | **GET** /api/monitoring/incidents/{id} | Get a Specific Incident
[**listIncidents()**](IncidentsApi.md#listIncidents) | **GET** /api/monitoring/incidents | List All Incidents
[**updateIncidents()**](IncidentsApi.md#updateIncidents) | **PUT** /api/monitoring/incidents/{id} | Update Incident
[**updateIncidentsReopen()**](IncidentsApi.md#updateIncidentsReopen) | **GET** /api/monitoring/incidents/{id}/reopen | Reopen a Specific Incident
[**updateMuteAllIncidents()**](IncidentsApi.md#updateMuteAllIncidents) | **PUT** /api/monitoring/incidents/mute-all | Mute All Incidents
[**updateMuteIncidents()**](IncidentsApi.md#updateMuteIncidents) | **PUT** /api/monitoring/incidents/{id}/mute | Mute Incident


## `addIncident()`

```php
addIncident($inline_object87): object
```

Create a New Incident

Create a new incident.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IncidentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object87 = new \OpenAPI\Client\Model\InlineObject87(); // \OpenAPI\Client\Model\InlineObject87

try {
    $result = $apiInstance->addIncident($inline_object87);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IncidentsApi->addIncident: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object87** | [**\OpenAPI\Client\Model\InlineObject87**](../Model/InlineObject87.md)|  | [optional]

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

## `deleteIncidents()`

```php
deleteIncidents($id): \OpenAPI\Client\Model\Model200Success
```

Close a Specific Incident

Close an existing monitoring incident.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IncidentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteIncidents($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IncidentsApi->deleteIncidents: ', $e->getMessage(), PHP_EOL;
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

## `getIncidents()`

```php
getIncidents($id): \OpenAPI\Client\Model\InlineResponse20055
```

Get a Specific Incident

Get details about a specific incident.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IncidentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getIncidents($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IncidentsApi->getIncidents: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20055**](../Model/InlineResponse20055.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listIncidents()`

```php
listIncidents($max, $offset, $status, $severity): object
```

List All Incidents

Get a list of monitoring incidents.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IncidentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$status = running; // string | The instance status for filtering.
$severity = 'severity_example'; // string | Filter by severity

try {
    $result = $apiInstance->listIncidents($max, $offset, $status, $severity);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IncidentsApi->listIncidents: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **status** | **string**| The instance status for filtering. | [optional]
 **severity** | **string**| Filter by severity | [optional]

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

## `updateIncidents()`

```php
updateIncidents($id, $inline_object88): object
```

Update Incident

Update an existing incident.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IncidentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object88 = new \OpenAPI\Client\Model\InlineObject88(); // \OpenAPI\Client\Model\InlineObject88

try {
    $result = $apiInstance->updateIncidents($id, $inline_object88);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IncidentsApi->updateIncidents: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object88** | [**\OpenAPI\Client\Model\InlineObject88**](../Model/InlineObject88.md)|  | [optional]

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

## `updateIncidentsReopen()`

```php
updateIncidentsReopen($id): \OpenAPI\Client\Model\SuccessMessage
```

Reopen a Specific Incident

Get details about a specific incident.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IncidentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->updateIncidentsReopen($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IncidentsApi->updateIncidentsReopen: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\SuccessMessage**](../Model/SuccessMessage.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateMuteAllIncidents()`

```php
updateMuteAllIncidents($inline_object90): object
```

Mute All Incidents

Mute all existing incident.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IncidentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object90 = new \OpenAPI\Client\Model\InlineObject90(); // \OpenAPI\Client\Model\InlineObject90

try {
    $result = $apiInstance->updateMuteAllIncidents($inline_object90);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IncidentsApi->updateMuteAllIncidents: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object90** | [**\OpenAPI\Client\Model\InlineObject90**](../Model/InlineObject90.md)|  | [optional]

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

## `updateMuteIncidents()`

```php
updateMuteIncidents($id, $inline_object89): object
```

Mute Incident

Mute an existing incident.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\IncidentsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object89 = new \OpenAPI\Client\Model\InlineObject89(); // \OpenAPI\Client\Model\InlineObject89

try {
    $result = $apiInstance->updateMuteIncidents($id, $inline_object89);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling IncidentsApi->updateMuteIncidents: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object89** | [**\OpenAPI\Client\Model\InlineObject89**](../Model/InlineObject89.md)|  | [optional]

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
