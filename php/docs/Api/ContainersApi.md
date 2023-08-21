# OpenAPI\Client\ContainersApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**cloneImageContainerAction()**](ContainersApi.md#cloneImageContainerAction) | **PUT** /api/containers/{id}/clone-image | Clone Specific Container to Image
[**containersAttachFloatingIp()**](ContainersApi.md#containersAttachFloatingIp) | **PUT** /api/containers/{id}/attach-floating-ip | Attach Floating IP to Container
[**containersDetachFloatingIp()**](ContainersApi.md#containersDetachFloatingIp) | **PUT** /api/containers/{id}/detach-floating-ip | Detach Floating IP from Container
[**ejectContainerAction()**](ContainersApi.md#ejectContainerAction) | **PUT** /api/containers/{id}/eject | Eject a Specific Container
[**executeContainerAction()**](ContainersApi.md#executeContainerAction) | **PUT** /api/containers/{id}/action | Execute Container Action
[**getContainer()**](ContainersApi.md#getContainer) | **GET** /api/containers/{id} | Get a Specific Container
[**getContainerActions()**](ContainersApi.md#getContainerActions) | **GET** /api/containers/{id}/actions | List Container Actions
[**importContainerAction()**](ContainersApi.md#importContainerAction) | **PUT** /api/containers/{id}/import | Import a Specific Container
[**restartContainerAction()**](ContainersApi.md#restartContainerAction) | **PUT** /api/containers/{id}/restart | Restart a Specific Container
[**startContainerAction()**](ContainersApi.md#startContainerAction) | **PUT** /api/containers/{id}/start | Start a Specific Container
[**stopContainerAction()**](ContainersApi.md#stopContainerAction) | **PUT** /api/containers/{id}/stop | Stop a Specific Container
[**suspendContainerAction()**](ContainersApi.md#suspendContainerAction) | **PUT** /api/containers/{id}/suspend | Suspend a Specific Container


## `cloneImageContainerAction()`

```php
cloneImageContainerAction($id, $instances_clone_image): \OpenAPI\Client\Model\SuccessMessage
```

Clone Specific Container to Image

This endpoint clones and converts a Container in its current state to image in the source Cloud and adds Virtual Image Record with metadata matching the source configuration.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ContainersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$instances_clone_image = {"templateName":"Example Image","zoneFolder":"group-v81"}; // \OpenAPI\Client\Model\InstancesCloneImage

try {
    $result = $apiInstance->cloneImageContainerAction($id, $instances_clone_image);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContainersApi->cloneImageContainerAction: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **instances_clone_image** | [**\OpenAPI\Client\Model\InstancesCloneImage**](../Model/InstancesCloneImage.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessMessage**](../Model/SuccessMessage.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `containersAttachFloatingIp()`

```php
containersAttachFloatingIp($id, $inline_object63): \OpenAPI\Client\Model\Model200Success
```

Attach Floating IP to Container

This endpoint attaches a floating IP to a container (node/VM).

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ContainersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object63 = new \OpenAPI\Client\Model\InlineObject63(); // \OpenAPI\Client\Model\InlineObject63

try {
    $result = $apiInstance->containersAttachFloatingIp($id, $inline_object63);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContainersApi->containersAttachFloatingIp: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object63** | [**\OpenAPI\Client\Model\InlineObject63**](../Model/InlineObject63.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `containersDetachFloatingIp()`

```php
containersDetachFloatingIp($id): \OpenAPI\Client\Model\Model200Success
```

Detach Floating IP from Container

This endpoint detaches a floating IP from a container (node/VM).

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ContainersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->containersDetachFloatingIp($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContainersApi->containersDetachFloatingIp: ', $e->getMessage(), PHP_EOL;
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

## `ejectContainerAction()`

```php
ejectContainerAction($id): \OpenAPI\Client\Model\SuccessMessage
```

Eject a Specific Container

This endpoint ejects attached disk/iso of a specific container.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ContainersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->ejectContainerAction($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContainersApi->ejectContainerAction: ', $e->getMessage(), PHP_EOL;
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

## `executeContainerAction()`

```php
executeContainerAction($id, $code): \OpenAPI\Client\Model\SuccessMessage
```

Execute Container Action

This endpoint executes a container action on specific container.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ContainersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$code = apache-remove-node; // string | Action code to be executed on a specific container.

try {
    $result = $apiInstance->executeContainerAction($id, $code);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContainersApi->executeContainerAction: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **code** | **string**| Action code to be executed on a specific container. |

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

## `getContainer()`

```php
getContainer($id): \OpenAPI\Client\Model\InlineResponse20034
```

Get a Specific Container

This endpoint retrieves a specific container.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ContainersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getContainer($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContainersApi->getContainer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20034**](../Model/InlineResponse20034.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getContainerActions()`

```php
getContainerActions($id): \OpenAPI\Client\Model\InlineResponse20035
```

List Container Actions

This endpoint lists available actions for specific container.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ContainersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getContainerActions($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContainersApi->getContainerActions: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20035**](../Model/InlineResponse20035.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `importContainerAction()`

```php
importContainerAction($id, $inline_object62): \OpenAPI\Client\Model\SuccessMessage
```

Import a Specific Container

This endpoint clones and exports a Container in its current state to target Storage provider and adds Virtual Image Record with metadata matching the source configuration.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ContainersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object62 = new \OpenAPI\Client\Model\InlineObject62(); // \OpenAPI\Client\Model\InlineObject62

try {
    $result = $apiInstance->importContainerAction($id, $inline_object62);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContainersApi->importContainerAction: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object62** | [**\OpenAPI\Client\Model\InlineObject62**](../Model/InlineObject62.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessMessage**](../Model/SuccessMessage.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `restartContainerAction()`

```php
restartContainerAction($id): \OpenAPI\Client\Model\SuccessMessage
```

Restart a Specific Container

This endpoint restarts a specific container.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ContainersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->restartContainerAction($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContainersApi->restartContainerAction: ', $e->getMessage(), PHP_EOL;
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

## `startContainerAction()`

```php
startContainerAction($id): \OpenAPI\Client\Model\SuccessMessage
```

Start a Specific Container

This endpoint starts a specific container.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ContainersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->startContainerAction($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContainersApi->startContainerAction: ', $e->getMessage(), PHP_EOL;
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

## `stopContainerAction()`

```php
stopContainerAction($id): \OpenAPI\Client\Model\SuccessMessage
```

Stop a Specific Container

This endpoint stops a specific container.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ContainersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->stopContainerAction($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContainersApi->stopContainerAction: ', $e->getMessage(), PHP_EOL;
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

## `suspendContainerAction()`

```php
suspendContainerAction($id): \OpenAPI\Client\Model\SuccessMessage
```

Suspend a Specific Container

This endpoint suspends a specific container.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\ContainersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->suspendContainerAction($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContainersApi->suspendContainerAction: ', $e->getMessage(), PHP_EOL;
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
