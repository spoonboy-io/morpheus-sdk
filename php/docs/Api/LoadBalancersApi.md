# OpenAPI\Client\LoadBalancersApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**createLoadBalancer()**](LoadBalancersApi.md#createLoadBalancer) | **POST** /api/load-balancers | Create a Load Balancer
[**createLoadBalancerMonitor()**](LoadBalancersApi.md#createLoadBalancerMonitor) | **POST** /api/load-balancers/{loadBalancerId}/monitors | Create a Load Balancer Monitor
[**createLoadBalancerPool()**](LoadBalancersApi.md#createLoadBalancerPool) | **POST** /api/load-balancers/{loadBalancerId}/pools | Create a Load Balancer Pool
[**createLoadBalancerPoolNode()**](LoadBalancersApi.md#createLoadBalancerPoolNode) | **POST** /api/load-balancer-pools/{loadBalancerPoolId}/nodes | Create a Load Balancer Pool Node
[**createLoadBalancerProfile()**](LoadBalancersApi.md#createLoadBalancerProfile) | **POST** /api/load-balancers/{loadBalancerId}/profiles | Create a Load Balancer Profile
[**createLoadBalancerVirtualServer()**](LoadBalancersApi.md#createLoadBalancerVirtualServer) | **POST** /api/load-balancers/{loadBalancerId}/virtual-servers | Create a Load Balancer Virtual Server
[**deleteLoadBalancer()**](LoadBalancersApi.md#deleteLoadBalancer) | **DELETE** /api/load-balancers/{loadBalancerId} | Delete a Load Balancer
[**deleteLoadBalancerMonitor()**](LoadBalancersApi.md#deleteLoadBalancerMonitor) | **DELETE** /api/load-balancers/{loadBalancerId}/monitors/{id} | Delete a Load Balancer Monitor
[**deleteLoadBalancerPool()**](LoadBalancersApi.md#deleteLoadBalancerPool) | **DELETE** /api/load-balancers/{loadBalancerId}/pools/{id} | Delete a Load Balancer Pool
[**deleteLoadBalancerPoolNode()**](LoadBalancersApi.md#deleteLoadBalancerPoolNode) | **DELETE** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Delete a Load Balancer Pool Node
[**deleteLoadBalancerProfile()**](LoadBalancersApi.md#deleteLoadBalancerProfile) | **DELETE** /api/load-balancers/{loadBalancerId}/profiles/{id} | Delete a Load Balancer Profile
[**deleteLoadBalancerVirtualServer()**](LoadBalancersApi.md#deleteLoadBalancerVirtualServer) | **DELETE** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Delete a Load Balancer Virtual Server
[**getLoadBalancer()**](LoadBalancersApi.md#getLoadBalancer) | **GET** /api/load-balancers/{loadBalancerId} | Get a Specific Load Balancer
[**getLoadBalancerMonitor()**](LoadBalancersApi.md#getLoadBalancerMonitor) | **GET** /api/load-balancers/{loadBalancerId}/monitors/{id} | Get a Specific Load Balancer Monitor
[**getLoadBalancerPool()**](LoadBalancersApi.md#getLoadBalancerPool) | **GET** /api/load-balancers/{loadBalancerId}/pools/{id} | Get a Specific Load Balancer Pool
[**getLoadBalancerPoolNode()**](LoadBalancersApi.md#getLoadBalancerPoolNode) | **GET** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Get a Specific Load Balancer Pool Node
[**getLoadBalancerProfile()**](LoadBalancersApi.md#getLoadBalancerProfile) | **GET** /api/load-balancers/{loadBalancerId}/profiles/{id} | Get a Specific Load Balancer Profile
[**getLoadBalancerType()**](LoadBalancersApi.md#getLoadBalancerType) | **GET** /api/load-balancer-types/{id} | Get a Specific Load Balancer Type
[**getLoadBalancerVirtualServer()**](LoadBalancersApi.md#getLoadBalancerVirtualServer) | **GET** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Get a Specific Load Balancer Virtual Server
[**listLoadBalancerMonitors()**](LoadBalancersApi.md#listLoadBalancerMonitors) | **GET** /api/load-balancers/{loadBalancerId}/monitors | Get All Load Balancer Monitors For Load Balancer
[**listLoadBalancerPoolNodes()**](LoadBalancersApi.md#listLoadBalancerPoolNodes) | **GET** /api/load-balancer-pools/{loadBalancerPoolId}/nodes | Get All Load Balancer Pool Nodes For Load Balancer Pool
[**listLoadBalancerPools()**](LoadBalancersApi.md#listLoadBalancerPools) | **GET** /api/load-balancers/{loadBalancerId}/pools | Get All Load Balancer Pools For Load Balancer
[**listLoadBalancerProfiles()**](LoadBalancersApi.md#listLoadBalancerProfiles) | **GET** /api/load-balancers/{loadBalancerId}/profiles | Get All Load Balancer Profiles For Load Balancer
[**listLoadBalancerTypes()**](LoadBalancersApi.md#listLoadBalancerTypes) | **GET** /api/load-balancer-types | Get All Load Balancer Types
[**listLoadBalancerVirtualServers()**](LoadBalancersApi.md#listLoadBalancerVirtualServers) | **GET** /api/load-balancers/{loadBalancerId}/virtual-servers | Get All Load Balancer Virtual Servers For Load Balancer
[**listLoadBalancers()**](LoadBalancersApi.md#listLoadBalancers) | **GET** /api/load-balancers | Get All Load Balancers
[**refreshLoadBalancer()**](LoadBalancersApi.md#refreshLoadBalancer) | **PUT** /api/load-balancers/{loadBalancerId}/refresh | Refresh a Load Balancer
[**updateLoadBalancer()**](LoadBalancersApi.md#updateLoadBalancer) | **PUT** /api/load-balancers/{loadBalancerId} | Update a Load Balancer
[**updateLoadBalancerMonitor()**](LoadBalancersApi.md#updateLoadBalancerMonitor) | **PUT** /api/load-balancers/{loadBalancerId}/monitors/{id} | Update a Load Balancer Monitor
[**updateLoadBalancerPool()**](LoadBalancersApi.md#updateLoadBalancerPool) | **PUT** /api/load-balancers/{loadBalancerId}/pools/{id} | Update a Load Balancer Pool
[**updateLoadBalancerPoolNode()**](LoadBalancersApi.md#updateLoadBalancerPoolNode) | **PUT** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Update a Load Balancer Pool Node
[**updateLoadBalancerProfile()**](LoadBalancersApi.md#updateLoadBalancerProfile) | **PUT** /api/load-balancers/{loadBalancerId}/profiles/{id} | Update a Load Balancer Profile
[**updateLoadBalancerVirtualServer()**](LoadBalancersApi.md#updateLoadBalancerVirtualServer) | **PUT** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Update a Load Balancer Virtual Server


## `createLoadBalancer()`

```php
createLoadBalancer($inline_object127): \OpenAPI\Client\Model\InlineResponse20078
```

Create a Load Balancer

Available for NSX load balancers only  Use this command to create a load balancer.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object127 = new \OpenAPI\Client\Model\InlineObject127(); // \OpenAPI\Client\Model\InlineObject127

try {
    $result = $apiInstance->createLoadBalancer($inline_object127);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->createLoadBalancer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object127** | [**\OpenAPI\Client\Model\InlineObject127**](../Model/InlineObject127.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20078**](../Model/InlineResponse20078.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `createLoadBalancerMonitor()`

```php
createLoadBalancerMonitor($load_balancer_id, $inline_object129): object
```

Create a Load Balancer Monitor

Use this command to create a load balancer Monitor.  This endpoint allows creating a Load Balancer Monitor. Configuration options vary by Load Balancer Type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$inline_object129 = new \OpenAPI\Client\Model\InlineObject129(); // \OpenAPI\Client\Model\InlineObject129

try {
    $result = $apiInstance->createLoadBalancerMonitor($load_balancer_id, $inline_object129);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->createLoadBalancerMonitor: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **inline_object129** | [**\OpenAPI\Client\Model\InlineObject129**](../Model/InlineObject129.md)|  | [optional]

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

## `createLoadBalancerPool()`

```php
createLoadBalancerPool($load_balancer_id, $inline_object131): object
```

Create a Load Balancer Pool

Use this command to create a load balancer pool.  This endpoint allows creating a Load Balancer Pool. Configuration options vary by Load Balancer Type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$inline_object131 = new \OpenAPI\Client\Model\InlineObject131(); // \OpenAPI\Client\Model\InlineObject131

try {
    $result = $apiInstance->createLoadBalancerPool($load_balancer_id, $inline_object131);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->createLoadBalancerPool: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **inline_object131** | [**\OpenAPI\Client\Model\InlineObject131**](../Model/InlineObject131.md)|  | [optional]

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

## `createLoadBalancerPoolNode()`

```php
createLoadBalancerPoolNode($load_balancer_pool_id, $inline_object137): object
```

Create a Load Balancer Pool Node

Use this command to create a load balancer pool node.  This endpoint allows creating a Load Balancer Pool Node. Configuration options vary by Load Balancer Type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_pool_id = 4; // float | Load Balancer Pool ID
$inline_object137 = new \OpenAPI\Client\Model\InlineObject137(); // \OpenAPI\Client\Model\InlineObject137

try {
    $result = $apiInstance->createLoadBalancerPoolNode($load_balancer_pool_id, $inline_object137);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->createLoadBalancerPoolNode: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_pool_id** | **float**| Load Balancer Pool ID |
 **inline_object137** | [**\OpenAPI\Client\Model\InlineObject137**](../Model/InlineObject137.md)|  | [optional]

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

## `createLoadBalancerProfile()`

```php
createLoadBalancerProfile($load_balancer_id, $inline_object133): object
```

Create a Load Balancer Profile

Use this command to create a load balancer Profile.  This endpoint allows creating a Load Balancer Profile. Configuration options vary by Load Balancer Type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$inline_object133 = new \OpenAPI\Client\Model\InlineObject133(); // \OpenAPI\Client\Model\InlineObject133

try {
    $result = $apiInstance->createLoadBalancerProfile($load_balancer_id, $inline_object133);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->createLoadBalancerProfile: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **inline_object133** | [**\OpenAPI\Client\Model\InlineObject133**](../Model/InlineObject133.md)|  | [optional]

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

## `createLoadBalancerVirtualServer()`

```php
createLoadBalancerVirtualServer($load_balancer_id, $inline_object135): \OpenAPI\Client\Model\InlineResponse20082
```

Create a Load Balancer Virtual Server

Use this command to create a load balancer virtual server.  This endpoint allows creating a Load Balancer Virtual Server. Configuration options vary by Load Balancer Type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$inline_object135 = new \OpenAPI\Client\Model\InlineObject135(); // \OpenAPI\Client\Model\InlineObject135

try {
    $result = $apiInstance->createLoadBalancerVirtualServer($load_balancer_id, $inline_object135);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->createLoadBalancerVirtualServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **inline_object135** | [**\OpenAPI\Client\Model\InlineObject135**](../Model/InlineObject135.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20082**](../Model/InlineResponse20082.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteLoadBalancer()`

```php
deleteLoadBalancer($load_balancer_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Load Balancer

Will delete a Load Balancer from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID

try {
    $result = $apiInstance->deleteLoadBalancer($load_balancer_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->deleteLoadBalancer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |

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

## `deleteLoadBalancerMonitor()`

```php
deleteLoadBalancerMonitor($load_balancer_id, $id): \OpenAPI\Client\Model\Model200Success
```

Delete a Load Balancer Monitor

Will delete a Load Balancer Monitor from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteLoadBalancerMonitor($load_balancer_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->deleteLoadBalancerMonitor: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
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

## `deleteLoadBalancerPool()`

```php
deleteLoadBalancerPool($load_balancer_id, $id): \OpenAPI\Client\Model\Model200Success
```

Delete a Load Balancer Pool

Will delete a Load Balancer Pool from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteLoadBalancerPool($load_balancer_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->deleteLoadBalancerPool: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
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

## `deleteLoadBalancerPoolNode()`

```php
deleteLoadBalancerPoolNode($load_balancer_pool_id, $id): \OpenAPI\Client\Model\Model200Success
```

Delete a Load Balancer Pool Node

Will delete a Load Balancer Pool Node from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_pool_id = 4; // float | Load Balancer Pool ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteLoadBalancerPoolNode($load_balancer_pool_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->deleteLoadBalancerPoolNode: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_pool_id** | **float**| Load Balancer Pool ID |
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

## `deleteLoadBalancerProfile()`

```php
deleteLoadBalancerProfile($load_balancer_id, $id): \OpenAPI\Client\Model\Model200Success
```

Delete a Load Balancer Profile

Will delete a Load Balancer Profile from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteLoadBalancerProfile($load_balancer_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->deleteLoadBalancerProfile: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
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

## `deleteLoadBalancerVirtualServer()`

```php
deleteLoadBalancerVirtualServer($load_balancer_id, $id): \OpenAPI\Client\Model\Model200Success
```

Delete a Load Balancer Virtual Server

Will delete a Load Balancer Virtual Server from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteLoadBalancerVirtualServer($load_balancer_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->deleteLoadBalancerVirtualServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
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

## `getLoadBalancer()`

```php
getLoadBalancer($load_balancer_id): \OpenAPI\Client\Model\InlineResponse20078
```

Get a Specific Load Balancer

This endpoint retrieves a specific Load Balancer.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID

try {
    $result = $apiInstance->getLoadBalancer($load_balancer_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->getLoadBalancer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20078**](../Model/InlineResponse20078.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getLoadBalancerMonitor()`

```php
getLoadBalancerMonitor($load_balancer_id, $id): \OpenAPI\Client\Model\InlineResponse20079
```

Get a Specific Load Balancer Monitor

This endpoint retrieves a specific Load Balancer Monitor.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getLoadBalancerMonitor($load_balancer_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->getLoadBalancerMonitor: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20079**](../Model/InlineResponse20079.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getLoadBalancerPool()`

```php
getLoadBalancerPool($load_balancer_id, $id): \OpenAPI\Client\Model\InlineResponse20080
```

Get a Specific Load Balancer Pool

This endpoint retrieves a specific Load Balancer Pool.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getLoadBalancerPool($load_balancer_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->getLoadBalancerPool: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20080**](../Model/InlineResponse20080.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getLoadBalancerPoolNode()`

```php
getLoadBalancerPoolNode($load_balancer_pool_id, $id): \OpenAPI\Client\Model\InlineResponse20083
```

Get a Specific Load Balancer Pool Node

This endpoint retrieves a specific Load Balancer Pool Node.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_pool_id = 4; // float | Load Balancer Pool ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getLoadBalancerPoolNode($load_balancer_pool_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->getLoadBalancerPoolNode: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_pool_id** | **float**| Load Balancer Pool ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20083**](../Model/InlineResponse20083.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getLoadBalancerProfile()`

```php
getLoadBalancerProfile($load_balancer_id, $id): \OpenAPI\Client\Model\InlineResponse20081
```

Get a Specific Load Balancer Profile

This endpoint retrieves a specific Load Balancer Profile.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getLoadBalancerProfile($load_balancer_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->getLoadBalancerProfile: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20081**](../Model/InlineResponse20081.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getLoadBalancerType()`

```php
getLoadBalancerType($id): \OpenAPI\Client\Model\InlineResponse20077
```

Get a Specific Load Balancer Type

This endpoint will retrieve a specific load balancer type by id.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getLoadBalancerType($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->getLoadBalancerType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20077**](../Model/InlineResponse20077.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getLoadBalancerVirtualServer()`

```php
getLoadBalancerVirtualServer($load_balancer_id, $id): \OpenAPI\Client\Model\InlineResponse20082
```

Get a Specific Load Balancer Virtual Server

This endpoint retrieves a specific Load Balancer Virtual Server.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getLoadBalancerVirtualServer($load_balancer_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->getLoadBalancerVirtualServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20082**](../Model/InlineResponse20082.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listLoadBalancerMonitors()`

```php
listLoadBalancerMonitors($load_balancer_id, $max, $offset, $sort, $direction, $name, $phrase): object
```

Get All Load Balancer Monitors For Load Balancer

This endpoint retrieves all load balancer monitors associated with a specified load balancer.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listLoadBalancerMonitors($load_balancer_id, $max, $offset, $sort, $direction, $name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->listLoadBalancerMonitors: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
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

## `listLoadBalancerPoolNodes()`

```php
listLoadBalancerPoolNodes($load_balancer_pool_id, $max, $offset, $sort, $direction, $phrase): object
```

Get All Load Balancer Pool Nodes For Load Balancer Pool

This endpoint retrieves all load balancer pool nodes associated with a specified load balancer pool.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_pool_id = 4; // float | Load Balancer Pool ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listLoadBalancerPoolNodes($load_balancer_pool_id, $max, $offset, $sort, $direction, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->listLoadBalancerPoolNodes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_pool_id** | **float**| Load Balancer Pool ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
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

## `listLoadBalancerPools()`

```php
listLoadBalancerPools($load_balancer_id, $max, $offset, $sort, $direction, $name, $phrase): object
```

Get All Load Balancer Pools For Load Balancer

This endpoint retrieves all load balancer pools associated with a specified load balancer.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listLoadBalancerPools($load_balancer_id, $max, $offset, $sort, $direction, $name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->listLoadBalancerPools: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
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

## `listLoadBalancerProfiles()`

```php
listLoadBalancerProfiles($load_balancer_id, $max, $offset, $sort, $direction, $name, $phrase): object
```

Get All Load Balancer Profiles For Load Balancer

This endpoint retrieves all load balancer profiles associated with a specified load balancer.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listLoadBalancerProfiles($load_balancer_id, $max, $offset, $sort, $direction, $name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->listLoadBalancerProfiles: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
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

## `listLoadBalancerTypes()`

```php
listLoadBalancerTypes($max, $offset, $sort, $direction, $option_types, $phrase, $name, $code): object
```

Get All Load Balancer Types

This endpoint retrieves all Load Balancer Types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$option_types = false; // bool | Pass true to include optionTypes in the response for each entry.
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$code = azr; // string | If specified will return an exact match on code

try {
    $result = $apiInstance->listLoadBalancerTypes($max, $offset, $sort, $direction, $option_types, $phrase, $name, $code);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->listLoadBalancerTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **option_types** | **bool**| Pass true to include optionTypes in the response for each entry. | [optional] [default to false]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **string**| If specified will return an exact match on code | [optional]

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

## `listLoadBalancerVirtualServers()`

```php
listLoadBalancerVirtualServers($load_balancer_id, $max, $offset, $sort, $direction, $phrase, $vip_name, $vip_address, $vip_hostname): object
```

Get All Load Balancer Virtual Servers For Load Balancer

This endpoint retrieves load balancer virtual servers associated with a specified load balancer.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$vip_name = lb-114; // string | If specified will return an exact match on vipName
$vip_address = 192.168.123.44; // string | If specified will return an exact match on vipAddress
$vip_hostname = mylb; // string | If specified will return an exact match on vipHostname

try {
    $result = $apiInstance->listLoadBalancerVirtualServers($load_balancer_id, $max, $offset, $sort, $direction, $phrase, $vip_name, $vip_address, $vip_hostname);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->listLoadBalancerVirtualServers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **vip_name** | **string**| If specified will return an exact match on vipName | [optional]
 **vip_address** | **string**| If specified will return an exact match on vipAddress | [optional]
 **vip_hostname** | **string**| If specified will return an exact match on vipHostname | [optional]

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

## `listLoadBalancers()`

```php
listLoadBalancers($max, $offset, $sort, $direction, $name, $phrase): object
```

Get All Load Balancers

This endpoint retrieves all load balancers associated with the account.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description

try {
    $result = $apiInstance->listLoadBalancers($max, $offset, $sort, $direction, $name, $phrase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->listLoadBalancers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
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

## `refreshLoadBalancer()`

```php
refreshLoadBalancer($load_balancer_id): object
```

Refresh a Load Balancer

Will refresh a Load Balancer.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID

try {
    $result = $apiInstance->refreshLoadBalancer($load_balancer_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->refreshLoadBalancer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |

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

## `updateLoadBalancer()`

```php
updateLoadBalancer($load_balancer_id, $inline_object128): \OpenAPI\Client\Model\InlineResponse20078
```

Update a Load Balancer

Available for NSX load balancers only  Use this command to update an existing load balancer.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$inline_object128 = new \OpenAPI\Client\Model\InlineObject128(); // \OpenAPI\Client\Model\InlineObject128

try {
    $result = $apiInstance->updateLoadBalancer($load_balancer_id, $inline_object128);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->updateLoadBalancer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **inline_object128** | [**\OpenAPI\Client\Model\InlineObject128**](../Model/InlineObject128.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20078**](../Model/InlineResponse20078.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateLoadBalancerMonitor()`

```php
updateLoadBalancerMonitor($load_balancer_id, $id, $inline_object130): object
```

Update a Load Balancer Monitor

Use this command to update an existing load balancer monitor.  This endpoint allows updating a Load Balancer Monitor. Configuration options vary by Load Balancer Type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object130 = new \OpenAPI\Client\Model\InlineObject130(); // \OpenAPI\Client\Model\InlineObject130

try {
    $result = $apiInstance->updateLoadBalancerMonitor($load_balancer_id, $id, $inline_object130);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->updateLoadBalancerMonitor: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object130** | [**\OpenAPI\Client\Model\InlineObject130**](../Model/InlineObject130.md)|  | [optional]

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

## `updateLoadBalancerPool()`

```php
updateLoadBalancerPool($load_balancer_id, $id, $inline_object132): object
```

Update a Load Balancer Pool

Use this command to update an existing load balancer pool.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object132 = new \OpenAPI\Client\Model\InlineObject132(); // \OpenAPI\Client\Model\InlineObject132

try {
    $result = $apiInstance->updateLoadBalancerPool($load_balancer_id, $id, $inline_object132);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->updateLoadBalancerPool: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object132** | [**\OpenAPI\Client\Model\InlineObject132**](../Model/InlineObject132.md)|  | [optional]

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

## `updateLoadBalancerPoolNode()`

```php
updateLoadBalancerPoolNode($load_balancer_pool_id, $id, $inline_object138): object
```

Update a Load Balancer Pool Node

Use this command to update an existing load balancer pool node.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_pool_id = 4; // float | Load Balancer Pool ID
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object138 = new \OpenAPI\Client\Model\InlineObject138(); // \OpenAPI\Client\Model\InlineObject138

try {
    $result = $apiInstance->updateLoadBalancerPoolNode($load_balancer_pool_id, $id, $inline_object138);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->updateLoadBalancerPoolNode: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_pool_id** | **float**| Load Balancer Pool ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object138** | [**\OpenAPI\Client\Model\InlineObject138**](../Model/InlineObject138.md)|  | [optional]

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

## `updateLoadBalancerProfile()`

```php
updateLoadBalancerProfile($load_balancer_id, $id, $inline_object134): object
```

Update a Load Balancer Profile

Use this command to update an existing load balancer Profile.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object134 = new \OpenAPI\Client\Model\InlineObject134(); // \OpenAPI\Client\Model\InlineObject134

try {
    $result = $apiInstance->updateLoadBalancerProfile($load_balancer_id, $id, $inline_object134);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->updateLoadBalancerProfile: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object134** | [**\OpenAPI\Client\Model\InlineObject134**](../Model/InlineObject134.md)|  | [optional]

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

## `updateLoadBalancerVirtualServer()`

```php
updateLoadBalancerVirtualServer($load_balancer_id, $id, $inline_object136): \OpenAPI\Client\Model\InlineResponse20082
```

Update a Load Balancer Virtual Server

Use this command to update an existing load balancer virtual server.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LoadBalancersApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$load_balancer_id = 4; // float | Load Balancer ID
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object136 = new \OpenAPI\Client\Model\InlineObject136(); // \OpenAPI\Client\Model\InlineObject136

try {
    $result = $apiInstance->updateLoadBalancerVirtualServer($load_balancer_id, $id, $inline_object136);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LoadBalancersApi->updateLoadBalancerVirtualServer: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **load_balancer_id** | **float**| Load Balancer ID |
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object136** | [**\OpenAPI\Client\Model\InlineObject136**](../Model/InlineObject136.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20082**](../Model/InlineResponse20082.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
