# OpenAPI\Client\AutomationApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addExecuteSchedules()**](AutomationApi.md#addExecuteSchedules) | **POST** /api/execute-schedules | Creates a Execute Schedule
[**addPowerScheduleInstances()**](AutomationApi.md#addPowerScheduleInstances) | **PUT** /api/power-schedules/{id}/add-instances | Add Instances to a Power Schedule
[**addPowerScheduleServers()**](AutomationApi.md#addPowerScheduleServers) | **PUT** /api/power-schedules/{id}/add-servers | Add Servers to a Power Schedule
[**addPowerSchedules()**](AutomationApi.md#addPowerSchedules) | **POST** /api/power-schedules | Creates a Power Schedule
[**addScaleThresholds()**](AutomationApi.md#addScaleThresholds) | **POST** /api/scale-thresholds | Creates a Scale Threshold
[**addTasks()**](AutomationApi.md#addTasks) | **POST** /api/tasks | Creates a Task
[**addWorkflows()**](AutomationApi.md#addWorkflows) | **POST** /api/task-sets | Creates a Workflow
[**deletePowerScheduleInstances()**](AutomationApi.md#deletePowerScheduleInstances) | **PUT** /api/power-schedules/{id}/remove-instances | Remove Instances from a Power Schedule
[**deletePowerScheduleServers()**](AutomationApi.md#deletePowerScheduleServers) | **PUT** /api/power-schedules/{id}/remove-servers | Remove Servers from a Power Schedule
[**executeExecutionRequest()**](AutomationApi.md#executeExecutionRequest) | **POST** /api/execution-request/execute | Executes an Execution Request
[**executeTasks()**](AutomationApi.md#executeTasks) | **POST** /api/tasks/{id}/execute | Executes a Task
[**executeWorkflows()**](AutomationApi.md#executeWorkflows) | **POST** /api/task-sets/{id}/execute | Executes a Workflow
[**getExecuteSchedules()**](AutomationApi.md#getExecuteSchedules) | **GET** /api/execute-schedules/{id} | Retrieves a Specific Execute Schedule
[**getExecutionRequest()**](AutomationApi.md#getExecutionRequest) | **GET** /api/execution-request/{uniqueId} | Retrieves a Specific Execution Request
[**getPowerSchedules()**](AutomationApi.md#getPowerSchedules) | **GET** /api/power-schedules/{id} | Retrieves a Specific Power Schedule
[**getScaleThresholds()**](AutomationApi.md#getScaleThresholds) | **GET** /api/scale-thresholds/{id} | Retrieves a Specific Scale Threshold
[**getTaskTypes()**](AutomationApi.md#getTaskTypes) | **GET** /api/task-types/{id} | Retrieves a Specific Task Type
[**getTasks()**](AutomationApi.md#getTasks) | **GET** /api/tasks/{id} | Retrieves a Specific Task
[**getWorkflows()**](AutomationApi.md#getWorkflows) | **GET** /api/task-sets/{id} | Retrieves a Specific Workflow
[**listExecuteSchedules()**](AutomationApi.md#listExecuteSchedules) | **GET** /api/execute-schedules | Retrieves all Execute Schedules
[**listPowerSchedules()**](AutomationApi.md#listPowerSchedules) | **GET** /api/power-schedules | Retrieves all Power Schedules
[**listScaleThresholds()**](AutomationApi.md#listScaleThresholds) | **GET** /api/scale-thresholds | Retrieves all Scale Thresholds
[**listTaskTypes()**](AutomationApi.md#listTaskTypes) | **GET** /api/task-types | Retrieves all Task Types
[**listTasks()**](AutomationApi.md#listTasks) | **GET** /api/tasks | Retrieves all Tasks
[**listWorkflows()**](AutomationApi.md#listWorkflows) | **GET** /api/task-sets | Retrieves all Workflows
[**removeExecuteSchedules()**](AutomationApi.md#removeExecuteSchedules) | **DELETE** /api/execute-schedules/{id} | Deletes a Execute Schedule
[**removePowerSchedules()**](AutomationApi.md#removePowerSchedules) | **DELETE** /api/power-schedules/{id} | Deletes a Power Schedule
[**removeScaleThresholds()**](AutomationApi.md#removeScaleThresholds) | **DELETE** /api/scale-thresholds/{id} | Deletes a Scale Threshold
[**removeTasks()**](AutomationApi.md#removeTasks) | **DELETE** /api/tasks/{id} | Deletes a Task
[**removeWorkflows()**](AutomationApi.md#removeWorkflows) | **DELETE** /api/task-sets/{id} | Deletes a Workflow
[**updateExecuteSchedules()**](AutomationApi.md#updateExecuteSchedules) | **PUT** /api/execute-schedules/{id} | Updates a Execute Schedule
[**updatePowerSchedules()**](AutomationApi.md#updatePowerSchedules) | **PUT** /api/power-schedules/{id} | Updates a Power Schedule
[**updateScaleThresholds()**](AutomationApi.md#updateScaleThresholds) | **PUT** /api/scale-thresholds/{id} | Updates a Scale Threshold
[**updateTasks()**](AutomationApi.md#updateTasks) | **PUT** /api/tasks/{id} | Updates a Task
[**updateWorkflows()**](AutomationApi.md#updateWorkflows) | **PUT** /api/task-sets/{id} | Updates a Workflow


## `addExecuteSchedules()`

```php
addExecuteSchedules($inline_object12): object
```

Creates a Execute Schedule

Creates a execute schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object12 = new \OpenAPI\Client\Model\InlineObject12(); // \OpenAPI\Client\Model\InlineObject12

try {
    $result = $apiInstance->addExecuteSchedules($inline_object12);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->addExecuteSchedules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object12** | [**\OpenAPI\Client\Model\InlineObject12**](../Model/InlineObject12.md)|  | [optional]

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

## `addPowerScheduleInstances()`

```php
addPowerScheduleInstances($id, $inline_object192): \OpenAPI\Client\Model\Model200SuccessExpanded
```

Add Instances to a Power Schedule

Add Instances to a Power Schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object192 = new \OpenAPI\Client\Model\InlineObject192(); // \OpenAPI\Client\Model\InlineObject192

try {
    $result = $apiInstance->addPowerScheduleInstances($id, $inline_object192);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->addPowerScheduleInstances: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object192** | [**\OpenAPI\Client\Model\InlineObject192**](../Model/InlineObject192.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200SuccessExpanded**](../Model/Model200SuccessExpanded.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addPowerScheduleServers()`

```php
addPowerScheduleServers($id, $inline_object193): \OpenAPI\Client\Model\Model200SuccessExpanded
```

Add Servers to a Power Schedule

Add Servers to a Power Schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object193 = new \OpenAPI\Client\Model\InlineObject193(); // \OpenAPI\Client\Model\InlineObject193

try {
    $result = $apiInstance->addPowerScheduleServers($id, $inline_object193);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->addPowerScheduleServers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object193** | [**\OpenAPI\Client\Model\InlineObject193**](../Model/InlineObject193.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200SuccessExpanded**](../Model/Model200SuccessExpanded.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addPowerSchedules()`

```php
addPowerSchedules($inline_object190): object
```

Creates a Power Schedule

Creates a power schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object190 = new \OpenAPI\Client\Model\InlineObject190(); // \OpenAPI\Client\Model\InlineObject190

try {
    $result = $apiInstance->addPowerSchedules($inline_object190);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->addPowerSchedules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object190** | [**\OpenAPI\Client\Model\InlineObject190**](../Model/InlineObject190.md)|  | [optional]

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

## `addScaleThresholds()`

```php
addScaleThresholds($inline_object211): object
```

Creates a Scale Threshold

Creates a scale threshold.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object211 = new \OpenAPI\Client\Model\InlineObject211(); // \OpenAPI\Client\Model\InlineObject211

try {
    $result = $apiInstance->addScaleThresholds($inline_object211);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->addScaleThresholds: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object211** | [**\OpenAPI\Client\Model\InlineObject211**](../Model/InlineObject211.md)|  | [optional]

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

## `addTasks()`

```php
addTasks($inline_object246): object
```

Creates a Task

Creates a task.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object246 = new \OpenAPI\Client\Model\InlineObject246(); // \OpenAPI\Client\Model\InlineObject246

try {
    $result = $apiInstance->addTasks($inline_object246);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->addTasks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object246** | [**\OpenAPI\Client\Model\InlineObject246**](../Model/InlineObject246.md)|  | [optional]

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

## `addWorkflows()`

```php
addWorkflows($inline_object249): object
```

Creates a Workflow

Creates a workflow.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object249 = new \OpenAPI\Client\Model\InlineObject249(); // \OpenAPI\Client\Model\InlineObject249

try {
    $result = $apiInstance->addWorkflows($inline_object249);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->addWorkflows: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object249** | [**\OpenAPI\Client\Model\InlineObject249**](../Model/InlineObject249.md)|  | [optional]

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

## `deletePowerScheduleInstances()`

```php
deletePowerScheduleInstances($id, $inline_object194): \OpenAPI\Client\Model\Model200SuccessExpanded
```

Remove Instances from a Power Schedule

Remove Instances from a Power Schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object194 = new \OpenAPI\Client\Model\InlineObject194(); // \OpenAPI\Client\Model\InlineObject194

try {
    $result = $apiInstance->deletePowerScheduleInstances($id, $inline_object194);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->deletePowerScheduleInstances: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object194** | [**\OpenAPI\Client\Model\InlineObject194**](../Model/InlineObject194.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200SuccessExpanded**](../Model/Model200SuccessExpanded.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deletePowerScheduleServers()`

```php
deletePowerScheduleServers($id, $inline_object195): \OpenAPI\Client\Model\Model200SuccessExpanded
```

Remove Servers from a Power Schedule

Remove Servers from a Power Schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object195 = new \OpenAPI\Client\Model\InlineObject195(); // \OpenAPI\Client\Model\InlineObject195

try {
    $result = $apiInstance->deletePowerScheduleServers($id, $inline_object195);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->deletePowerScheduleServers: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object195** | [**\OpenAPI\Client\Model\InlineObject195**](../Model/InlineObject195.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200SuccessExpanded**](../Model/Model200SuccessExpanded.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `executeExecutionRequest()`

```php
executeExecutionRequest($instance_id, $container_id, $server_id, $inline_object14): \OpenAPI\Client\Model\InlineResponse2008
```

Executes an Execution Request

Provides API interfaces for executing an arbitrary script or command on an instance, container or host.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$instance_id = 94; // int | The Instance ID for Filtering
$container_id = 135; // int | The Container ID for Filtering
$server_id = 97; // int | The Server ID for Filtering
$inline_object14 = new \OpenAPI\Client\Model\InlineObject14(); // \OpenAPI\Client\Model\InlineObject14

try {
    $result = $apiInstance->executeExecutionRequest($instance_id, $container_id, $server_id, $inline_object14);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->executeExecutionRequest: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_id** | **int**| The Instance ID for Filtering | [optional]
 **container_id** | **int**| The Container ID for Filtering | [optional]
 **server_id** | **int**| The Server ID for Filtering | [optional]
 **inline_object14** | [**\OpenAPI\Client\Model\InlineObject14**](../Model/InlineObject14.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse2008**](../Model/InlineResponse2008.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `executeTasks()`

```php
executeTasks($id, $inline_object248): object
```

Executes a Task

Executes a task.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object248 = new \OpenAPI\Client\Model\InlineObject248(); // \OpenAPI\Client\Model\InlineObject248

try {
    $result = $apiInstance->executeTasks($id, $inline_object248);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->executeTasks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object248** | [**\OpenAPI\Client\Model\InlineObject248**](../Model/InlineObject248.md)|  | [optional]

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

## `executeWorkflows()`

```php
executeWorkflows($id, $inline_object251): object
```

Executes a Workflow

Executes a workflow.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object251 = new \OpenAPI\Client\Model\InlineObject251(); // \OpenAPI\Client\Model\InlineObject251

try {
    $result = $apiInstance->executeWorkflows($id, $inline_object251);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->executeWorkflows: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object251** | [**\OpenAPI\Client\Model\InlineObject251**](../Model/InlineObject251.md)|  | [optional]

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

## `getExecuteSchedules()`

```php
getExecuteSchedules($id): object
```

Retrieves a Specific Execute Schedule

Retrieves a specific execute schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getExecuteSchedules($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->getExecuteSchedules: ', $e->getMessage(), PHP_EOL;
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

## `getExecutionRequest()`

```php
getExecutionRequest($unique_id): object
```

Retrieves a Specific Execution Request

Retrieves a specific execution request.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$unique_id = c56f75d0-a59a-4566-92e3-4dc716c5d076; // string | The Unique ID of the execution request

try {
    $result = $apiInstance->getExecutionRequest($unique_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->getExecutionRequest: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unique_id** | **string**| The Unique ID of the execution request |

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

## `getPowerSchedules()`

```php
getPowerSchedules($id): object
```

Retrieves a Specific Power Schedule

Retrieves a specific power schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getPowerSchedules($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->getPowerSchedules: ', $e->getMessage(), PHP_EOL;
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

## `getScaleThresholds()`

```php
getScaleThresholds($id): object
```

Retrieves a Specific Scale Threshold

Retrieves a specific scale threshold.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getScaleThresholds($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->getScaleThresholds: ', $e->getMessage(), PHP_EOL;
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

## `getTaskTypes()`

```php
getTaskTypes($id): \OpenAPI\Client\Model\InlineResponse200156
```

Retrieves a Specific Task Type

Retrieves a specific task type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getTaskTypes($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->getTaskTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200156**](../Model/InlineResponse200156.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getTasks()`

```php
getTasks($id): object
```

Retrieves a Specific Task

Retrieves a specific task.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getTasks($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->getTasks: ', $e->getMessage(), PHP_EOL;
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

## `getWorkflows()`

```php
getWorkflows($id): object
```

Retrieves a Specific Workflow

Retrieves a specific workflow.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getWorkflows($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->getWorkflows: ', $e->getMessage(), PHP_EOL;
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

## `listExecuteSchedules()`

```php
listExecuteSchedules($max, $offset, $sort, $direction, $phrase, $name): object
```

Retrieves all Execute Schedules

Retrieves all execute schedules.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
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
    $result = $apiInstance->listExecuteSchedules($max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->listExecuteSchedules: ', $e->getMessage(), PHP_EOL;
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

## `listPowerSchedules()`

```php
listPowerSchedules($max, $offset, $sort, $direction, $phrase, $name): object
```

Retrieves all Power Schedules

Retrieves all power schedules.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
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
    $result = $apiInstance->listPowerSchedules($max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->listPowerSchedules: ', $e->getMessage(), PHP_EOL;
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

## `listScaleThresholds()`

```php
listScaleThresholds($max, $offset, $sort, $direction, $phrase, $name): object
```

Retrieves all Scale Thresholds

Retrieves all scale thresholds.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
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
    $result = $apiInstance->listScaleThresholds($max, $offset, $sort, $direction, $phrase, $name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->listScaleThresholds: ', $e->getMessage(), PHP_EOL;
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

## `listTaskTypes()`

```php
listTaskTypes($name, $code): \OpenAPI\Client\Model\InlineResponse200155
```

Retrieves all Task Types

A Task Type is a type of automation task. Each type defines its own set of options to be configured for each task.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$code = azr; // string | If specified will return an exact match on code

try {
    $result = $apiInstance->listTaskTypes($name, $code);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->listTaskTypes: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **string**| If specified will return an exact match on code | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse200155**](../Model/InlineResponse200155.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listTasks()`

```php
listTasks($type, $max, $offset, $sort, $direction, $phrase, $name, $labels, $all_labels): object
```

Retrieves all Tasks

Retrieves all tasks.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$type = 'type_example'; // string | If specified will return all tasks by `task type` code. Refer to `Task Types` API for up to date listings.
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels

try {
    $result = $apiInstance->listTasks($type, $max, $offset, $sort, $direction, $phrase, $name, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->listTasks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type** | **string**| If specified will return all tasks by &#x60;task type&#x60; code. Refer to &#x60;Task Types&#x60; API for up to date listings. | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **labels** | **string**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **string**| Filter by label(s), matches records that contain all of the specified labels | [optional]

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

## `listWorkflows()`

```php
listWorkflows($max, $offset, $sort, $direction, $phrase, $name, $labels, $all_labels, $type): object
```

Retrieves all Workflows

Retrieves all workflows.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
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
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels
$type = provision; // string | Filter by Workflow Type

try {
    $result = $apiInstance->listWorkflows($max, $offset, $sort, $direction, $phrase, $name, $labels, $all_labels, $type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->listWorkflows: ', $e->getMessage(), PHP_EOL;
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
 **labels** | **string**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **string**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **type** | **string**| Filter by Workflow Type | [optional]

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

## `removeExecuteSchedules()`

```php
removeExecuteSchedules($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Execute Schedule

Deletes a specified execute schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeExecuteSchedules($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->removeExecuteSchedules: ', $e->getMessage(), PHP_EOL;
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

## `removePowerSchedules()`

```php
removePowerSchedules($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Power Schedule

Deletes a specified power schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removePowerSchedules($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->removePowerSchedules: ', $e->getMessage(), PHP_EOL;
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

## `removeScaleThresholds()`

```php
removeScaleThresholds($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Scale Threshold

Deletes a specified scale threshold.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeScaleThresholds($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->removeScaleThresholds: ', $e->getMessage(), PHP_EOL;
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

## `removeTasks()`

```php
removeTasks($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Task

Deletes a specified task.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeTasks($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->removeTasks: ', $e->getMessage(), PHP_EOL;
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

## `removeWorkflows()`

```php
removeWorkflows($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Workflow

Deletes a specified workflow.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeWorkflows($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->removeWorkflows: ', $e->getMessage(), PHP_EOL;
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

## `updateExecuteSchedules()`

```php
updateExecuteSchedules($id, $inline_object13): object
```

Updates a Execute Schedule

Updates a execute schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object13 = new \OpenAPI\Client\Model\InlineObject13(); // \OpenAPI\Client\Model\InlineObject13

try {
    $result = $apiInstance->updateExecuteSchedules($id, $inline_object13);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->updateExecuteSchedules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object13** | [**\OpenAPI\Client\Model\InlineObject13**](../Model/InlineObject13.md)|  | [optional]

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

## `updatePowerSchedules()`

```php
updatePowerSchedules($id, $inline_object191): object
```

Updates a Power Schedule

Updates a power schedule.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object191 = new \OpenAPI\Client\Model\InlineObject191(); // \OpenAPI\Client\Model\InlineObject191

try {
    $result = $apiInstance->updatePowerSchedules($id, $inline_object191);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->updatePowerSchedules: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object191** | [**\OpenAPI\Client\Model\InlineObject191**](../Model/InlineObject191.md)|  | [optional]

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

## `updateScaleThresholds()`

```php
updateScaleThresholds($id, $inline_object212): object
```

Updates a Scale Threshold

Updates a scale threshold.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object212 = new \OpenAPI\Client\Model\InlineObject212(); // \OpenAPI\Client\Model\InlineObject212

try {
    $result = $apiInstance->updateScaleThresholds($id, $inline_object212);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->updateScaleThresholds: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object212** | [**\OpenAPI\Client\Model\InlineObject212**](../Model/InlineObject212.md)|  | [optional]

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

## `updateTasks()`

```php
updateTasks($id, $inline_object247): object
```

Updates a Task

Updates a task.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object247 = new \OpenAPI\Client\Model\InlineObject247(); // \OpenAPI\Client\Model\InlineObject247

try {
    $result = $apiInstance->updateTasks($id, $inline_object247);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->updateTasks: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object247** | [**\OpenAPI\Client\Model\InlineObject247**](../Model/InlineObject247.md)|  | [optional]

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

## `updateWorkflows()`

```php
updateWorkflows($id, $inline_object250): object
```

Updates a Workflow

Updates a workflow.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AutomationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object250 = new \OpenAPI\Client\Model\InlineObject250(); // \OpenAPI\Client\Model\InlineObject250

try {
    $result = $apiInstance->updateWorkflows($id, $inline_object250);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AutomationApi->updateWorkflows: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object250** | [**\OpenAPI\Client\Model\InlineObject250**](../Model/InlineObject250.md)|  | [optional]

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
