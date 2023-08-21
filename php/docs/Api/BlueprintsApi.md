# OpenAPI\Client\BlueprintsApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addBlueprint()**](BlueprintsApi.md#addBlueprint) | **POST** /api/blueprints | Create a Blueprint
[**deleteBlueprint()**](BlueprintsApi.md#deleteBlueprint) | **DELETE** /api/blueprints/{id} | Delete a Blueprint
[**getBlueprint()**](BlueprintsApi.md#getBlueprint) | **GET** /api/blueprints/{id} | Get a Specific Blueprint
[**listBlueprints()**](BlueprintsApi.md#listBlueprints) | **GET** /api/blueprints | Get All Blueprints
[**updateBlueprint()**](BlueprintsApi.md#updateBlueprint) | **PUT** /api/blueprints/{id} | Updating a Blueprint
[**updateBlueprintImage()**](BlueprintsApi.md#updateBlueprintImage) | **POST** /api/blueprints/{id}/image | Update Blueprint Image
[**updateBlueprintPermissions()**](BlueprintsApi.md#updateBlueprintPermissions) | **PUT** /api/blueprints/{id}/update-permissions | Update Blueprint Permissions


## `addBlueprint()`

```php
addBlueprint($unknown_base_type): object
```

Create a Blueprint

Create a Blueprint

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BlueprintsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$unknown_base_type = {"image":"/assets/apps/template.png","name":"ARM: Ubuntu 16Test","type":"arm","arm":{"json":"{\n    \"$schema\": \"http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#\",\n    \"contentVersion\": \"1.0.0.0\",\n    \"parameters\": {\n        \"virtualMachineName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"DMUBARM<%=sequence%>\",\n            \"metadata\": {\n                \"description\": \"Name of Virtual Machine\"\n            }\n        },\n        \"networkInterfaceName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"demo-arm-ubuntu-nic\",\n            \"metadata\": {\n                \"description\": \"Name of Network Interface\"\n            }\n        },\n        \"publicIpAddressName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"demo-arm-ubuntu-ip\",\n            \"metadata\": {\n                \"description\": \"Name of IP Address\"\n            }\n        },\n        \"adminUsername\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"morpheus\",\n            \"metadata\": {\n                \"description\": \"OS Username\"\n            }\n        },\n        \"adminPassword\": {\n            \"type\": \"securestring\",\n            \"metadata\": {\n                \"description\": \"OS Password\"\n            }\n        },\n        \"location\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"westus\"\n        },\n        \"virtualMachineSize\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"Standard_A1\"\n        },\n        \"virtualNetworkName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"morphdemoVnet\"\n        },    \n        \"networkSecurityGroupName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"morph-demo-sg\"\n        },\n        \"diagnosticsStorageAccountName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"morphdemosa\"\n        },\n        \"subnetName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"morphDemoSubnet\"\n        },\n        \"publicIpAddressType\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"Dynamic\"\n        },\n        \"publicIpAddressSku\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"Basic\"\n        }\n    },\n    \"variables\": {\n        \"vnetId\": \"[resourceId('morph-demo','Microsoft.Network/virtualNetworks', parameters('virtualNetworkName'))]\",\n        \"subnetRef\": \"[concat(variables('vnetId'), '/subnets/', parameters('subnetName'))]\"\n    },\n    \"resources\": [\n        {\n            \"name\": \"[parameters('virtualMachineName')]\",\n            \"type\": \"Microsoft.Compute/virtualMachines\",\n            \"apiVersion\": \"2016-04-30-preview\",\n            \"location\": \"[parameters('location')]\",\n            \"dependsOn\": [\n                \"[concat('Microsoft.Network/networkInterfaces/', parameters('networkInterfaceName'))]\"\n            ],\n            \"properties\": {\n                \"osProfile\": {\n                    \"computerName\": \"[parameters('virtualMachineName')]\",\n                    \"adminUsername\": \"[parameters('adminUsername')]\",\n                    \"adminPassword\": \"[parameters('adminPassword')]\"\n                },\n                \"hardwareProfile\": {\n                    \"vmSize\": \"[parameters('virtualMachineSize')]\"\n                },\n                \"storageProfile\": {\n                    \"imageReference\": {\n                        \"publisher\": \"Canonical\",\n                        \"offer\": \"UbuntuServer\",\n                        \"sku\": \"16.04-LTS\",\n                        \"version\": \"latest\"\n                    },\n                    \"osDisk\": {\n                        \"createOption\": \"fromImage\",\n                        \"managedDisk\": {\n                            \"storageAccountType\": \"Standard_LRS\"\n                        }\n                    },\n                    \"dataDisks\": []\n                },\n                \"networkProfile\": {\n                    \"networkInterfaces\": [\n                        {\n                            \"id\": \"[resourceId('Microsoft.Network/networkInterfaces', parameters('networkInterfaceName'))]\"\n                        }\n                    ]\n                },\n                \"diagnosticsProfile\": {\n                    \"bootDiagnostics\": {\n                        \"enabled\": true,\n                        \"storageUri\": \"[reference(resourceId('morph-demo', 'Microsoft.Storage/storageAccounts', parameters('diagnosticsStorageAccountName')), '2015-06-15').primaryEndpoints['blob']]\"\n                    }\n                }\n            }\n        },\n        {\n            \"name\": \"[parameters('networkInterfaceName')]\",\n            \"type\": \"Microsoft.Network/networkInterfaces\",\n            \"apiVersion\": \"2016-09-01\",\n            \"location\": \"[parameters('location')]\",\n            \"dependsOn\": [\n                \"[concat('Microsoft.Network/publicIpAddresses/', parameters('publicIpAddressName'))]\"\n            ],\n            \"properties\": {\n                \"ipConfigurations\": [\n                    {\n                        \"name\": \"ipconfig1\",\n                        \"properties\": {\n                            \"subnet\": {\n                                \"id\": \"[variables('subnetRef')]\"\n                            },\n                            \"privateIPAllocationMethod\": \"Dynamic\",\n                            \"publicIpAddress\": {\n                                \"id\": \"[resourceId('morph-demo','Microsoft.Network/publicIpAddresses', parameters('publicIpAddressName'))]\"\n                            }\n                        }\n                    }\n                ],\n                \"networkSecurityGroup\": {\n                    \"id\": \"[resourceId('morph-demo', 'Microsoft.Network/networkSecurityGroups', parameters('networkSecurityGroupName'))]\"\n                }\n            }\n        },\n        {\n            \"name\": \"[parameters('publicIpAddressName')]\",\n            \"type\": \"Microsoft.Network/publicIpAddresses\",\n            \"apiVersion\": \"2017-08-01\",\n            \"location\": \"[parameters('location')]\",\n            \"properties\": {\n                \"publicIpAllocationMethod\": \"[parameters('publicIpAddressType')]\"\n            },\n            \"sku\": {\n                \"name\": \"[parameters('publicIpAddressSku')]\"\n            }\n        }\n    ],\n    \"outputs\": {\n        \"adminUsername\": {\n            \"type\": \"string\",\n            \"value\": \"[parameters('adminUsername')]\"\n        }\n    }\n}","configType":"json","installAgent":true}}; // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->addBlueprint($unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BlueprintsApi->addBlueprint: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

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

## `deleteBlueprint()`

```php
deleteBlueprint($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Blueprint

Delete a Blueprint

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BlueprintsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteBlueprint($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BlueprintsApi->deleteBlueprint: ', $e->getMessage(), PHP_EOL;
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

## `getBlueprint()`

```php
getBlueprint($id): \OpenAPI\Client\Model\InlineResponse20014
```

Get a Specific Blueprint

This endpoint retrieves a specific blueprint.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BlueprintsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getBlueprint($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BlueprintsApi->getBlueprint: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20014**](../Model/InlineResponse20014.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listBlueprints()`

```php
listBlueprints($max, $offset, $name, $phrase, $labels, $all_labels): object
```

Get All Blueprints

This endpoint retrieves all blueprints.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BlueprintsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels

try {
    $result = $apiInstance->listBlueprints($max, $offset, $name, $phrase, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BlueprintsApi->listBlueprints: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
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

## `updateBlueprint()`

```php
updateBlueprint($id, $unknown_base_type): \OpenAPI\Client\Model\InlineResponse20014
```

Updating a Blueprint

Update a Blueprint. This overwrites the entire config, so the entire blueprint config should be passed.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BlueprintsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$unknown_base_type = {"name":"sample","description":"A sample nginx blueprint","type":"morpheus","tiers":{"Web":{"linkedTiers":[],"tierIndex":1,"instances":[{"instance":{"type":"nginx"},"groups":{"My Group":{"clouds":{"My Cloud":{"instance":{"layout":{"code":"nginx-vmware-1.9-single","id":179},"name":"test-nginx-${sequence}","allowExisting":false,"createUser":"on","type":"nginx","userGroup":{"id":""}},"networkInterfaces":[{"ipMode":"","primaryInterface":true,"network":{"id":"","hasPool":false},"networkInterfaceTypeId":4,"networkInterfaceTypeIdName":"VMXNET 3"}],"volumes":[{"vId":255,"controllerMountPoint":"46:0:4:0","size":10,"maxIOPS":null,"name":"root","rootVolume":true,"storageType":1,"datastoreId":"autoCluster","maxStorage":0}],"config":{"resourcePoolId":"resgroup-123","createUser":true},"plan":{"code":"vm-1024","id":76}}}}}}]}}}; // \OpenAPI\Client\Model\UNKNOWN_BASE_TYPE

try {
    $result = $apiInstance->updateBlueprint($id, $unknown_base_type);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BlueprintsApi->updateBlueprint: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **unknown_base_type** | [**\OpenAPI\Client\Model\UNKNOWN_BASE_TYPE**](../Model/UNKNOWN_BASE_TYPE.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20014**](../Model/InlineResponse20014.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateBlueprintImage()`

```php
updateBlueprintImage($id, $template_image): \OpenAPI\Client\Model\InlineResponse20014
```

Update Blueprint Image

Update Blueprint Image

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BlueprintsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$template_image = "/path/to/file.txt"; // \SplFileObject

try {
    $result = $apiInstance->updateBlueprintImage($id, $template_image);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BlueprintsApi->updateBlueprintImage: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **template_image** | **\SplFileObject****\SplFileObject**|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20014**](../Model/InlineResponse20014.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateBlueprintPermissions()`

```php
updateBlueprintPermissions($id, $inline_object21): \OpenAPI\Client\Model\InlineResponse20014
```

Update Blueprint Permissions

Update Blueprint Permissions

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\BlueprintsApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object21 = new \OpenAPI\Client\Model\InlineObject21(); // \OpenAPI\Client\Model\InlineObject21

try {
    $result = $apiInstance->updateBlueprintPermissions($id, $inline_object21);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BlueprintsApi->updateBlueprintPermissions: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object21** | [**\OpenAPI\Client\Model\InlineObject21**](../Model/InlineObject21.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse20014**](../Model/InlineResponse20014.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
