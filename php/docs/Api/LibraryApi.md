# OpenAPI\Client\LibraryApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addFileTemplate()**](LibraryApi.md#addFileTemplate) | **POST** /api/library/container-templates | Create a File Template
[**addInstanceType()**](LibraryApi.md#addInstanceType) | **POST** /api/library/instance-types | Create an Instance Type
[**addLayout()**](LibraryApi.md#addLayout) | **POST** /api/library/instance-types/{instanceTypeId}/layouts | Create a Layout
[**addNodeType()**](LibraryApi.md#addNodeType) | **POST** /api/library/container-types | Create a Node Type
[**addOptionList()**](LibraryApi.md#addOptionList) | **POST** /api/library/option-type-lists | Create an Option List
[**addOptionType()**](LibraryApi.md#addOptionType) | **POST** /api/library/option-types | Create an Input
[**addScript()**](LibraryApi.md#addScript) | **POST** /api/library/container-scripts | Create a Script
[**addSpecTemplate()**](LibraryApi.md#addSpecTemplate) | **POST** /api/library/spec-templates | Create a Spec Template
[**addVirtualImage()**](LibraryApi.md#addVirtualImage) | **POST** /api/virtual-images | Create a Virtual Image
[**addVirtualImageFile()**](LibraryApi.md#addVirtualImageFile) | **POST** /api/virtual-images/{virtualImageId}/upload | Upload Virtual Image File
[**deleteFileTemplate()**](LibraryApi.md#deleteFileTemplate) | **DELETE** /api/library/container-templates/{id} | Delete a File Template
[**deleteInstanceType()**](LibraryApi.md#deleteInstanceType) | **DELETE** /api/library/instance-types/{instanceTypeId} | Delete an Instance Type
[**deleteLayout()**](LibraryApi.md#deleteLayout) | **DELETE** /api/library/layouts/{id} | Delete a Layout
[**deleteNodeType()**](LibraryApi.md#deleteNodeType) | **DELETE** /api/library/container-types/{id} | Delete a Node Type
[**deleteOptionList()**](LibraryApi.md#deleteOptionList) | **DELETE** /api/library/option-type-lists/{id} | Delete an Option List
[**deleteOptionType()**](LibraryApi.md#deleteOptionType) | **DELETE** /api/library/option-types/{id} | Delete an Input
[**deleteScript()**](LibraryApi.md#deleteScript) | **DELETE** /api/library/container-scripts/{id} | Delete a Script
[**deleteSpecTemplate()**](LibraryApi.md#deleteSpecTemplate) | **DELETE** /api/library/spec-templates/{id} | Delete a Spec Template
[**getFileTemplate()**](LibraryApi.md#getFileTemplate) | **GET** /api/library/container-templates/{id} | Get a Specific File Template
[**getInput()**](LibraryApi.md#getInput) | **GET** /api/library/option-types/{id} | Get A Specific Input
[**getInstanceType()**](LibraryApi.md#getInstanceType) | **GET** /api/library/instance-types/{instanceTypeId} | Get a Specific Instance Type
[**getLayout()**](LibraryApi.md#getLayout) | **GET** /api/library/layouts/{id} | Get a Specific Layout
[**getNodeType()**](LibraryApi.md#getNodeType) | **GET** /api/library/container-types/{id} | Get a Specific Node Type
[**getOptionList()**](LibraryApi.md#getOptionList) | **GET** /api/library/option-type-lists/{id} | Get a Specific Option List
[**getOptionListItems()**](LibraryApi.md#getOptionListItems) | **GET** /api/library/option-type-lists/{id}/items | List Items for a Specific Option List
[**getScript()**](LibraryApi.md#getScript) | **GET** /api/library/container-scripts/{id} | Get a Specific Script
[**getSecurityPackageType()**](LibraryApi.md#getSecurityPackageType) | **GET** /api/security-package-types/{id} | Retrieves a Specific Security Package Type
[**getSpecTemplate()**](LibraryApi.md#getSpecTemplate) | **GET** /api/library/spec-templates/{id} | Get a Specific Spec Template
[**getVirtualImage()**](LibraryApi.md#getVirtualImage) | **GET** /api/virtual-images/{virtualImageId} | Get a Specific Virtual Image
[**listFileTemplates()**](LibraryApi.md#listFileTemplates) | **GET** /api/library/container-templates | Get All File Templates
[**listInputs()**](LibraryApi.md#listInputs) | **GET** /api/library/option-types | Get All Inputs
[**listInstanceTypes()**](LibraryApi.md#listInstanceTypes) | **GET** /api/library/instance-types | Get All Instance Types
[**listLayouts()**](LibraryApi.md#listLayouts) | **GET** /api/library/layouts | Get All Layouts
[**listLayoutsForInstanceType()**](LibraryApi.md#listLayoutsForInstanceType) | **GET** /api/library/instance-types/{instanceTypeId}/layouts | Get All Layouts For an Instance Type
[**listNodeTypes()**](LibraryApi.md#listNodeTypes) | **GET** /api/library/container-types | Get All Node Types
[**listOptionLists()**](LibraryApi.md#listOptionLists) | **GET** /api/library/option-type-lists | Get All Option Lists
[**listScripts()**](LibraryApi.md#listScripts) | **GET** /api/library/container-scripts | Get All Scripts
[**listSecurityPackageTypes()**](LibraryApi.md#listSecurityPackageTypes) | **GET** /api/security-package-types | Retrieves all Security Package Types
[**listSpecTemplates()**](LibraryApi.md#listSpecTemplates) | **GET** /api/library/spec-templates | Get All Spec Templates
[**listVirtualImageLocations()**](LibraryApi.md#listVirtualImageLocations) | **GET** /api/virtual-images/{virtualImageId}/locations | Get a List of Virtual Image Locations
[**listVirtualImages()**](LibraryApi.md#listVirtualImages) | **GET** /api/virtual-images | Get List of Virtual Images
[**removeSecurityScans()**](LibraryApi.md#removeSecurityScans) | **DELETE** /api/security-scans/{id} | Deletes a Security Scan
[**removeVirtualImage()**](LibraryApi.md#removeVirtualImage) | **DELETE** /api/virtual-images/{virtualImageId} | Delete a Virtual Image
[**removeVirtualImageFile()**](LibraryApi.md#removeVirtualImageFile) | **DELETE** /api/virtual-images/{virtualImageId}/files | Remove Virtual Image File
[**removeVirtualImageLocation()**](LibraryApi.md#removeVirtualImageLocation) | **DELETE** /api/virtual-images/{virtualImageId}/locations/{id} | Delete a Virtual Image Location
[**setInstanceTypeFeatured()**](LibraryApi.md#setInstanceTypeFeatured) | **PUT** /api/library/instance-types/{instanceTypeId}/toggle-featured | Toggle Featured For Instance Type
[**updateFileTemplate()**](LibraryApi.md#updateFileTemplate) | **PUT** /api/library/container-templates/{id} | Update a File Template
[**updateInstanceType()**](LibraryApi.md#updateInstanceType) | **PUT** /api/library/instance-types/{instanceTypeId} | Update an Instance Type
[**updateInstanceTypeLogo()**](LibraryApi.md#updateInstanceTypeLogo) | **POST** /api/library/instance-types/{instanceTypeId}/update-logo | Update Logo For Instance Type
[**updateLayout()**](LibraryApi.md#updateLayout) | **PUT** /api/library/layouts/{id} | Update a Layout
[**updateLayoutPermissions()**](LibraryApi.md#updateLayoutPermissions) | **PUT** /api/library/layouts/{id}/permissions | Update Layout Permissions
[**updateNodeType()**](LibraryApi.md#updateNodeType) | **PUT** /api/library/container-types/{id} | Update a Node Type
[**updateOptionList()**](LibraryApi.md#updateOptionList) | **PUT** /api/library/option-type-lists/{id} | Update an Option List
[**updateOptionType()**](LibraryApi.md#updateOptionType) | **PUT** /api/library/option-types/{id} | Update an Input
[**updateScript()**](LibraryApi.md#updateScript) | **PUT** /api/library/container-scripts/{id} | Update a Script
[**updateSpecTemplate()**](LibraryApi.md#updateSpecTemplate) | **PUT** /api/library/spec-templates/{id} | Update a Spec Template
[**updateVirtualImage()**](LibraryApi.md#updateVirtualImage) | **PUT** /api/virtual-images/{virtualImageId} | Update a Virtual Image


## `addFileTemplate()`

```php
addFileTemplate($inline_object111): \OpenAPI\Client\Model\SuccessId
```

Create a File Template

Use this command to create a file template.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object111 = new \OpenAPI\Client\Model\InlineObject111(); // \OpenAPI\Client\Model\InlineObject111

try {
    $result = $apiInstance->addFileTemplate($inline_object111);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->addFileTemplate: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object111** | [**\OpenAPI\Client\Model\InlineObject111**](../Model/InlineObject111.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addInstanceType()`

```php
addInstanceType($inline_object113): \OpenAPI\Client\Model\Model200Success
```

Create an Instance Type

Use this command to create an instance type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object113 = new \OpenAPI\Client\Model\InlineObject113(); // \OpenAPI\Client\Model\InlineObject113

try {
    $result = $apiInstance->addInstanceType($inline_object113);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->addInstanceType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object113** | [**\OpenAPI\Client\Model\InlineObject113**](../Model/InlineObject113.md)|  | [optional]

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

## `addLayout()`

```php
addLayout($instance_type_id, $inline_object115): \OpenAPI\Client\Model\SuccessId
```

Create a Layout

Use this command to create a layout.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$instance_type_id = 2; // float | The ID of the instance type
$inline_object115 = new \OpenAPI\Client\Model\InlineObject115(); // \OpenAPI\Client\Model\InlineObject115

try {
    $result = $apiInstance->addLayout($instance_type_id, $inline_object115);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->addLayout: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |
 **inline_object115** | [**\OpenAPI\Client\Model\InlineObject115**](../Model/InlineObject115.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addNodeType()`

```php
addNodeType($inline_object109): object
```

Create a Node Type

Use this command to create a node type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object109 = new \OpenAPI\Client\Model\InlineObject109(); // \OpenAPI\Client\Model\InlineObject109

try {
    $result = $apiInstance->addNodeType($inline_object109);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->addNodeType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object109** | [**\OpenAPI\Client\Model\InlineObject109**](../Model/InlineObject109.md)|  | [optional]

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

## `addOptionList()`

```php
addOptionList($inline_object119): \OpenAPI\Client\Model\Model200Success
```

Create an Option List

Use this command to create an option list.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object119 = new \OpenAPI\Client\Model\InlineObject119(); // \OpenAPI\Client\Model\InlineObject119

try {
    $result = $apiInstance->addOptionList($inline_object119);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->addOptionList: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object119** | [**\OpenAPI\Client\Model\InlineObject119**](../Model/InlineObject119.md)|  | [optional]

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

## `addOptionType()`

```php
addOptionType($inline_object121): \OpenAPI\Client\Model\SuccessId
```

Create an Input

Use this command to create an option type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object121 = new \OpenAPI\Client\Model\InlineObject121(); // \OpenAPI\Client\Model\InlineObject121

try {
    $result = $apiInstance->addOptionType($inline_object121);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->addOptionType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object121** | [**\OpenAPI\Client\Model\InlineObject121**](../Model/InlineObject121.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addScript()`

```php
addScript($inline_object107): \OpenAPI\Client\Model\ScriptSuccessId
```

Create a Script

Use this command to create a script.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object107 = new \OpenAPI\Client\Model\InlineObject107(); // \OpenAPI\Client\Model\InlineObject107

try {
    $result = $apiInstance->addScript($inline_object107);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->addScript: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object107** | [**\OpenAPI\Client\Model\InlineObject107**](../Model/InlineObject107.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\ScriptSuccessId**](../Model/ScriptSuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addSpecTemplate()`

```php
addSpecTemplate($inline_object123): \OpenAPI\Client\Model\SuccessId
```

Create a Spec Template

Use this command to create a spec template.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object123 = new \OpenAPI\Client\Model\InlineObject123(); // \OpenAPI\Client\Model\InlineObject123

try {
    $result = $apiInstance->addSpecTemplate($inline_object123);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->addSpecTemplate: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object123** | [**\OpenAPI\Client\Model\InlineObject123**](../Model/InlineObject123.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\SuccessId**](../Model/SuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `addVirtualImage()`

```php
addVirtualImage($inline_object263): object
```

Create a Virtual Image

This endpoint creates a new virtual image, without any files yet.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$inline_object263 = new \OpenAPI\Client\Model\InlineObject263(); // \OpenAPI\Client\Model\InlineObject263

try {
    $result = $apiInstance->addVirtualImage($inline_object263);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->addVirtualImage: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object263** | [**\OpenAPI\Client\Model\InlineObject263**](../Model/InlineObject263.md)|  | [optional]

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

## `addVirtualImageFile()`

```php
addVirtualImageFile($virtual_image_id, $filename, $url, $body): \OpenAPI\Client\Model\Model200Success
```

Upload Virtual Image File

This will upload the file and associate it to the Virtual Image.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$virtual_image_id = 4; // float | Virtual Image ID
$filename = testimage.ovf; // string | The name of the file
$url = https://example.com/testimage.ovf; // string | Download the file from a remote url. This can be used instead of uploading a local file.
$body = "/path/to/file.txt"; // \SplFileObject

try {
    $result = $apiInstance->addVirtualImageFile($virtual_image_id, $filename, $url, $body);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->addVirtualImageFile: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |
 **filename** | **string**| The name of the file | [optional]
 **url** | **string**| Download the file from a remote url. This can be used instead of uploading a local file. | [optional]
 **body** | **\SplFileObject****\SplFileObject**|  | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/octet-stream`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `deleteFileTemplate()`

```php
deleteFileTemplate($id): \OpenAPI\Client\Model\Model200Success
```

Delete a File Template

Will delete a file template

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteFileTemplate($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->deleteFileTemplate: ', $e->getMessage(), PHP_EOL;
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

## `deleteInstanceType()`

```php
deleteInstanceType($instance_type_id): \OpenAPI\Client\Model\Model200Success
```

Delete an Instance Type

Will delete an instance type

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$instance_type_id = 2; // float | The ID of the instance type

try {
    $result = $apiInstance->deleteInstanceType($instance_type_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->deleteInstanceType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |

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

## `deleteLayout()`

```php
deleteLayout($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Layout

Will delete a layout

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteLayout($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->deleteLayout: ', $e->getMessage(), PHP_EOL;
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

## `deleteNodeType()`

```php
deleteNodeType($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Node Type

Will delete a node type

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteNodeType($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->deleteNodeType: ', $e->getMessage(), PHP_EOL;
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

## `deleteOptionList()`

```php
deleteOptionList($id): \OpenAPI\Client\Model\Model200Success
```

Delete an Option List

Will delete an option list.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteOptionList($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->deleteOptionList: ', $e->getMessage(), PHP_EOL;
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

## `deleteOptionType()`

```php
deleteOptionType($id): \OpenAPI\Client\Model\Model200Success
```

Delete an Input

Will delete an option type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteOptionType($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->deleteOptionType: ', $e->getMessage(), PHP_EOL;
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

## `deleteScript()`

```php
deleteScript($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Script

Will delete a script

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteScript($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->deleteScript: ', $e->getMessage(), PHP_EOL;
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

## `deleteSpecTemplate()`

```php
deleteSpecTemplate($id): \OpenAPI\Client\Model\Model200Success
```

Delete a Spec Template

Will delete a spec template

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->deleteSpecTemplate($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->deleteSpecTemplate: ', $e->getMessage(), PHP_EOL;
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

## `getFileTemplate()`

```php
getFileTemplate($id): \OpenAPI\Client\Model\InlineResponse20070
```

Get a Specific File Template

This endpoint retrieves a specific file template.  The value of template will be masked as ************ for system owned file templates.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getFileTemplate($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->getFileTemplate: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20070**](../Model/InlineResponse20070.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getInput()`

```php
getInput($id): \OpenAPI\Client\Model\InlineResponse20075
```

Get A Specific Input

This endpoint retrieves a specific option type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getInput($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->getInput: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20075**](../Model/InlineResponse20075.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getInstanceType()`

```php
getInstanceType($instance_type_id): \OpenAPI\Client\Model\InlineResponse20071
```

Get a Specific Instance Type

This endpoint retrieves a specific instance type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$instance_type_id = 2; // float | The ID of the instance type

try {
    $result = $apiInstance->getInstanceType($instance_type_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->getInstanceType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20071**](../Model/InlineResponse20071.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getLayout()`

```php
getLayout($id): \OpenAPI\Client\Model\InlineResponse20072
```

Get a Specific Layout

This endpoint retrieves a specific layout.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getLayout($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->getLayout: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20072**](../Model/InlineResponse20072.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getNodeType()`

```php
getNodeType($id): \OpenAPI\Client\Model\InlineResponse20069
```

Get a Specific Node Type

This endpoint retrieves a specific node type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getNodeType($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->getNodeType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20069**](../Model/InlineResponse20069.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getOptionList()`

```php
getOptionList($id): \OpenAPI\Client\Model\InlineResponse20073
```

Get a Specific Option List

This endpoint retrieves a specific option list.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getOptionList($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->getOptionList: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20073**](../Model/InlineResponse20073.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getOptionListItems()`

```php
getOptionListItems($id): \OpenAPI\Client\Model\InlineResponse20074
```

List Items for a Specific Option List

This endpoint retrieves the items for a specific option list.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getOptionListItems($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->getOptionListItems: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20074**](../Model/InlineResponse20074.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getScript()`

```php
getScript($id): \OpenAPI\Client\Model\InlineResponse20068
```

Get a Specific Script

This endpoint retrieves a specific script.  The value of script will be masked as ************ for system owned scripts.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getScript($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->getScript: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20068**](../Model/InlineResponse20068.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getSecurityPackageType()`

```php
getSecurityPackageType($id): \OpenAPI\Client\Model\InlineResponse200136
```

Retrieves a Specific Security Package Type

Retrieves a specific security package type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getSecurityPackageType($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->getSecurityPackageType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200136**](../Model/InlineResponse200136.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getSpecTemplate()`

```php
getSpecTemplate($id): \OpenAPI\Client\Model\InlineResponse20076
```

Get a Specific Spec Template

This endpoint retrieves a specific spec template.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->getSpecTemplate($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->getSpecTemplate: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |

### Return type

[**\OpenAPI\Client\Model\InlineResponse20076**](../Model/InlineResponse20076.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getVirtualImage()`

```php
getVirtualImage($virtual_image_id): \OpenAPI\Client\Model\InlineResponse200165
```

Get a Specific Virtual Image

This endpoint retrieves a specific virtual image and its files.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$virtual_image_id = 4; // float | Virtual Image ID

try {
    $result = $apiInstance->getVirtualImage($virtual_image_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->getVirtualImage: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |

### Return type

[**\OpenAPI\Client\Model\InlineResponse200165**](../Model/InlineResponse200165.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listFileTemplates()`

```php
listFileTemplates($max, $offset, $sort, $direction, $phrase, $name, $labels, $all_labels, $file_name): object
```

Get All File Templates

This endpoint retrieves all file templates.  The value of template will be masked as ************ for system owned file templates.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
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
$file_name = testtext.txt; // string | Filename filter, restricts query to only load file template matching fileName specified

try {
    $result = $apiInstance->listFileTemplates($max, $offset, $sort, $direction, $phrase, $name, $labels, $all_labels, $file_name);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->listFileTemplates: ', $e->getMessage(), PHP_EOL;
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
 **file_name** | **string**| Filename filter, restricts query to only load file template matching fileName specified | [optional]

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

## `listInputs()`

```php
listInputs($max, $offset, $sort, $direction, $phrase, $name, $code, $labels, $all_labels, $field_name, $field_context, $field_label): object
```

Get All Inputs

This endpoint retrieves all option types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
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
$code = azr; // string | If specified will return an exact match on code
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels
$field_name = cloudName; // string | Field Name filter, restricts query to only load type matching fieldName specified
$field_context = config.customOptions; // string | Field Context filter, restricts query to only load type matching fieldContext specified
$field_label = DB Version; // string | Field Label filter, restricts query to only load type matching fieldLabel specified

try {
    $result = $apiInstance->listInputs($max, $offset, $sort, $direction, $phrase, $name, $code, $labels, $all_labels, $field_name, $field_context, $field_label);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->listInputs: ', $e->getMessage(), PHP_EOL;
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
 **code** | **string**| If specified will return an exact match on code | [optional]
 **labels** | **string**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **string**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **field_name** | **string**| Field Name filter, restricts query to only load type matching fieldName specified | [optional]
 **field_context** | **string**| Field Context filter, restricts query to only load type matching fieldContext specified | [optional]
 **field_label** | **string**| Field Label filter, restricts query to only load type matching fieldLabel specified | [optional]

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

## `listInstanceTypes()`

```php
listInstanceTypes($max, $offset, $sort, $direction, $phrase, $name, $code, $featured, $labels, $all_labels, $details): object
```

Get All Instance Types

This endpoint retrieves all instance types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
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
$code = azr; // string | If specified will return an exact match on code
$featured = false; // bool | Filter by featured
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels
$details = true; // bool | Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default.

try {
    $result = $apiInstance->listInstanceTypes($max, $offset, $sort, $direction, $phrase, $name, $code, $featured, $labels, $all_labels, $details);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->listInstanceTypes: ', $e->getMessage(), PHP_EOL;
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
 **code** | **string**| If specified will return an exact match on code | [optional]
 **featured** | **bool**| Filter by featured | [optional]
 **labels** | **string**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **string**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **details** | **bool**| Load full details including optionTypes, environmentVariables, etc. These properties are excluded by default. | [optional]

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

## `listLayouts()`

```php
listLayouts($max, $offset, $sort, $direction, $phrase, $name, $code, $provision_type, $labels, $all_labels): object
```

Get All Layouts

This endpoint retrieves all layouts.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
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
$code = azr; // string | If specified will return an exact match on code
$provision_type = 'provision_type_example'; // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels

try {
    $result = $apiInstance->listLayouts($max, $offset, $sort, $direction, $phrase, $name, $code, $provision_type, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->listLayouts: ', $e->getMessage(), PHP_EOL;
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
 **code** | **string**| If specified will return an exact match on code | [optional]
 **provision_type** | **string**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings. | [optional]
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

## `listLayoutsForInstanceType()`

```php
listLayoutsForInstanceType($instance_type_id, $max, $offset, $sort, $direction, $phrase, $name, $code, $provision_type, $labels, $all_labels): object
```

Get All Layouts For an Instance Type

This endpoint retrieves all layouts for a specific instance type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$instance_type_id = 2; // float | The ID of the instance type
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$code = azr; // string | If specified will return an exact match on code
$provision_type = 'provision_type_example'; // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels

try {
    $result = $apiInstance->listLayoutsForInstanceType($instance_type_id, $max, $offset, $sort, $direction, $phrase, $name, $code, $provision_type, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->listLayoutsForInstanceType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **string**| If specified will return an exact match on code | [optional]
 **provision_type** | **string**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings. | [optional]
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

## `listNodeTypes()`

```php
listNodeTypes($max, $offset, $sort, $direction, $phrase, $name, $code, $provision_type, $labels, $all_labels): object
```

Get All Node Types

This endpoint retrieves all node types.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
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
$code = azr; // string | If specified will return an exact match on code
$provision_type = 'provision_type_example'; // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels

try {
    $result = $apiInstance->listNodeTypes($max, $offset, $sort, $direction, $phrase, $name, $code, $provision_type, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->listNodeTypes: ', $e->getMessage(), PHP_EOL;
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
 **code** | **string**| If specified will return an exact match on code | [optional]
 **provision_type** | **string**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings. | [optional]
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

## `listOptionLists()`

```php
listOptionLists($max, $offset, $sort, $direction, $phrase, $name, $labels, $all_labels): object
```

Get All Option Lists

This endpoint retrieves all option lists.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
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

try {
    $result = $apiInstance->listOptionLists($max, $offset, $sort, $direction, $phrase, $name, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->listOptionLists: ', $e->getMessage(), PHP_EOL;
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

## `listScripts()`

```php
listScripts($max, $offset, $sort, $direction, $phrase, $labels, $all_labels, $script_type, $script_phase): object
```

Get All Scripts

This endpoint retrieves all scripts.  The value of script will be masked as ************ for system owned scripts.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
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
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels
$script_type = 'script_type_example'; // string | Script type code filter, restricts query to only load scripts of specified type
$script_phase = 'script_phase_example'; // string | Script phase filter, restricts query to only load scripts of specified phase

try {
    $result = $apiInstance->listScripts($max, $offset, $sort, $direction, $phrase, $labels, $all_labels, $script_type, $script_phase);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->listScripts: ', $e->getMessage(), PHP_EOL;
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
 **labels** | **string**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **all_labels** | **string**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **script_type** | **string**| Script type code filter, restricts query to only load scripts of specified type | [optional]
 **script_phase** | **string**| Script phase filter, restricts query to only load scripts of specified phase | [optional]

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

## `listSecurityPackageTypes()`

```php
listSecurityPackageTypes(): object
```

Retrieves all Security Package Types

Retrieves all Security Package Types

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->listSecurityPackageTypes();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->listSecurityPackageTypes: ', $e->getMessage(), PHP_EOL;
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

## `listSpecTemplates()`

```php
listSpecTemplates($max, $offset, $sort, $direction, $phrase, $name, $labels, $all_labels): object
```

Get All Spec Templates

This endpoint retrieves all spec templates.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
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

try {
    $result = $apiInstance->listSpecTemplates($max, $offset, $sort, $direction, $phrase, $name, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->listSpecTemplates: ', $e->getMessage(), PHP_EOL;
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

## `listVirtualImageLocations()`

```php
listVirtualImageLocations($virtual_image_id): object
```

Get a List of Virtual Image Locations

This endpoint retrieves a specific virtual image and its files.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$virtual_image_id = 4; // float | Virtual Image ID

try {
    $result = $apiInstance->listVirtualImageLocations($virtual_image_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->listVirtualImageLocations: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |

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

## `listVirtualImages()`

```php
listVirtualImages($max, $offset, $name, $phrase, $last_updated, $filter_type, $image_type, $tags, $labels, $all_labels): object
```

Get List of Virtual Images

This endpoint retrieves a list of virtual images for the specified filter.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$name = example-%; // string | Filter by name, wildcard may be specified as %, eg. example-%
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$last_updated = 2019-03-06T17:52:29+0000; // \DateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
$filter_type = System; // string | Filter by type, \"User\", \"System\", \"Synced\", or \"All\"
$image_type = vmware; // string | Filter by image type code, \"vmware\", \"ami\", etc
$tags = tags.env=qa&tags.env=test; // string | Filter by tags (metadata). This allows filtering by a tag name and value(s)
$labels = 'labels_example'; // string | Filter by label(s), matches records that contain any of the specified labels
$all_labels = 'all_labels_example'; // string | Filter by label(s), matches records that contain all of the specified labels

try {
    $result = $apiInstance->listVirtualImages($max, $offset, $name, $phrase, $last_updated, $filter_type, $image_type, $tags, $labels, $all_labels);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->listVirtualImages: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **name** | **string**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **last_updated** | **\DateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **filter_type** | **string**| Filter by type, \&quot;User\&quot;, \&quot;System\&quot;, \&quot;Synced\&quot;, or \&quot;All\&quot; | [optional] [default to &#39;User&#39;]
 **image_type** | **string**| Filter by image type code, \&quot;vmware\&quot;, \&quot;ami\&quot;, etc | [optional]
 **tags** | **string**| Filter by tags (metadata). This allows filtering by a tag name and value(s) | [optional]
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

## `removeSecurityScans()`

```php
removeSecurityScans($id): \OpenAPI\Client\Model\Model200Success
```

Deletes a Security Scan

Deletes a specified security scan.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeSecurityScans($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->removeSecurityScans: ', $e->getMessage(), PHP_EOL;
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

## `removeVirtualImage()`

```php
removeVirtualImage($virtual_image_id): \OpenAPI\Client\Model\Model200Success
```

Delete a Virtual Image

Will delete a virtual image and any associated files, use removeFromCloud=true to also delete image locations from all clouds.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$virtual_image_id = 4; // float | Virtual Image ID

try {
    $result = $apiInstance->removeVirtualImage($virtual_image_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->removeVirtualImage: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |

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

## `removeVirtualImageFile()`

```php
removeVirtualImageFile($virtual_image_id, $filename): \OpenAPI\Client\Model\Model200Success
```

Remove Virtual Image File

Remove Virtual Image File

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$virtual_image_id = 4; // float | Virtual Image ID
$filename = testimage.ovf; // string | The name of the file

try {
    $result = $apiInstance->removeVirtualImageFile($virtual_image_id, $filename);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->removeVirtualImageFile: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |
 **filename** | **string**| The name of the file | [optional]

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

## `removeVirtualImageLocation()`

```php
removeVirtualImageLocation($virtual_image_id, $id): \OpenAPI\Client\Model\Model200Success
```

Delete a Virtual Image Location

Will delete a virtual image location, use removeFromCloud=true to all also delete image locations from all clouds as well.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$virtual_image_id = 4; // float | Virtual Image ID
$id = 1; // int | Morpheus ID of the Object being referenced

try {
    $result = $apiInstance->removeVirtualImageLocation($virtual_image_id, $id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->removeVirtualImageLocation: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |
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

## `setInstanceTypeFeatured()`

```php
setInstanceTypeFeatured($instance_type_id): \OpenAPI\Client\Model\Model200Success
```

Toggle Featured For Instance Type

Use this command to toggle the featured flag for an existing instance type. This will change the value from false to true, or from true to false.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$instance_type_id = 2; // float | The ID of the instance type

try {
    $result = $apiInstance->setInstanceTypeFeatured($instance_type_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->setInstanceTypeFeatured: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |

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

## `updateFileTemplate()`

```php
updateFileTemplate($id, $inline_object112): \OpenAPI\Client\Model\Model200Success
```

Update a File Template

Use this command to update an existing file template.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object112 = new \OpenAPI\Client\Model\InlineObject112(); // \OpenAPI\Client\Model\InlineObject112

try {
    $result = $apiInstance->updateFileTemplate($id, $inline_object112);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->updateFileTemplate: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object112** | [**\OpenAPI\Client\Model\InlineObject112**](../Model/InlineObject112.md)|  | [optional]

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

## `updateInstanceType()`

```php
updateInstanceType($instance_type_id, $inline_object114): \OpenAPI\Client\Model\Model200Success
```

Update an Instance Type

Use this command to update an existing instance type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$instance_type_id = 2; // float | The ID of the instance type
$inline_object114 = new \OpenAPI\Client\Model\InlineObject114(); // \OpenAPI\Client\Model\InlineObject114

try {
    $result = $apiInstance->updateInstanceType($instance_type_id, $inline_object114);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->updateInstanceType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |
 **inline_object114** | [**\OpenAPI\Client\Model\InlineObject114**](../Model/InlineObject114.md)|  | [optional]

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

## `updateInstanceTypeLogo()`

```php
updateInstanceTypeLogo($instance_type_id, $logo, $dark_logo): \OpenAPI\Client\Model\Model200Success
```

Update Logo For Instance Type

Use this command to update the logo and dark logo images for an existing instance type. This endpoint expects multipart form data as the request format, not JSON.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$instance_type_id = 2; // float | The ID of the instance type
$logo = "/path/to/file.txt"; // \SplFileObject | Logo File png,jpg,svg
$dark_logo = "/path/to/file.txt"; // \SplFileObject | Dark Logo File png,jpg,svg

try {
    $result = $apiInstance->updateInstanceTypeLogo($instance_type_id, $logo, $dark_logo);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->updateInstanceTypeLogo: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance_type_id** | **float**| The ID of the instance type |
 **logo** | **\SplFileObject****\SplFileObject**| Logo File png,jpg,svg | [optional]
 **dark_logo** | **\SplFileObject****\SplFileObject**| Dark Logo File png,jpg,svg | [optional]

### Return type

[**\OpenAPI\Client\Model\Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `multipart/form-data`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateLayout()`

```php
updateLayout($id, $inline_object117): \OpenAPI\Client\Model\Model200Success
```

Update a Layout

Use this command to update an existing layout.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object117 = new \OpenAPI\Client\Model\InlineObject117(); // \OpenAPI\Client\Model\InlineObject117

try {
    $result = $apiInstance->updateLayout($id, $inline_object117);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->updateLayout: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object117** | [**\OpenAPI\Client\Model\InlineObject117**](../Model/InlineObject117.md)|  | [optional]

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

## `updateLayoutPermissions()`

```php
updateLayoutPermissions($id, $inline_object118): \OpenAPI\Client\Model\Model200Success
```

Update Layout Permissions

Use this command to update permissions for an existing layout.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object118 = new \OpenAPI\Client\Model\InlineObject118(); // \OpenAPI\Client\Model\InlineObject118

try {
    $result = $apiInstance->updateLayoutPermissions($id, $inline_object118);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->updateLayoutPermissions: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object118** | [**\OpenAPI\Client\Model\InlineObject118**](../Model/InlineObject118.md)|  | [optional]

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

## `updateNodeType()`

```php
updateNodeType($id, $inline_object110): \OpenAPI\Client\Model\Model200Success
```

Update a Node Type

Use this command to update an existing node type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object110 = new \OpenAPI\Client\Model\InlineObject110(); // \OpenAPI\Client\Model\InlineObject110

try {
    $result = $apiInstance->updateNodeType($id, $inline_object110);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->updateNodeType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object110** | [**\OpenAPI\Client\Model\InlineObject110**](../Model/InlineObject110.md)|  | [optional]

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

## `updateOptionList()`

```php
updateOptionList($id, $inline_object120): \OpenAPI\Client\Model\Model200Success
```

Update an Option List

Use this command to update an existing option list.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object120 = new \OpenAPI\Client\Model\InlineObject120(); // \OpenAPI\Client\Model\InlineObject120

try {
    $result = $apiInstance->updateOptionList($id, $inline_object120);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->updateOptionList: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object120** | [**\OpenAPI\Client\Model\InlineObject120**](../Model/InlineObject120.md)|  | [optional]

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

## `updateOptionType()`

```php
updateOptionType($id, $inline_object122): \OpenAPI\Client\Model\Model200Success
```

Update an Input

Use this command to update an existing option type.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object122 = new \OpenAPI\Client\Model\InlineObject122(); // \OpenAPI\Client\Model\InlineObject122

try {
    $result = $apiInstance->updateOptionType($id, $inline_object122);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->updateOptionType: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object122** | [**\OpenAPI\Client\Model\InlineObject122**](../Model/InlineObject122.md)|  | [optional]

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

## `updateScript()`

```php
updateScript($id, $inline_object108): \OpenAPI\Client\Model\ScriptSuccessId
```

Update a Script

Use this command to update an existing script.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object108 = new \OpenAPI\Client\Model\InlineObject108(); // \OpenAPI\Client\Model\InlineObject108

try {
    $result = $apiInstance->updateScript($id, $inline_object108);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->updateScript: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object108** | [**\OpenAPI\Client\Model\InlineObject108**](../Model/InlineObject108.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\ScriptSuccessId**](../Model/ScriptSuccessId.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `updateSpecTemplate()`

```php
updateSpecTemplate($id, $inline_object124): \OpenAPI\Client\Model\Model200Success
```

Update a Spec Template

Use this command to update an existing spec template.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$id = 1; // int | Morpheus ID of the Object being referenced
$inline_object124 = new \OpenAPI\Client\Model\InlineObject124(); // \OpenAPI\Client\Model\InlineObject124

try {
    $result = $apiInstance->updateSpecTemplate($id, $inline_object124);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->updateSpecTemplate: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| Morpheus ID of the Object being referenced |
 **inline_object124** | [**\OpenAPI\Client\Model\InlineObject124**](../Model/InlineObject124.md)|  | [optional]

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

## `updateVirtualImage()`

```php
updateVirtualImage($virtual_image_id, $inline_object264): object
```

Update a Virtual Image

This endpoint updates an existing virtual image.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\LibraryApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$virtual_image_id = 4; // float | Virtual Image ID
$inline_object264 = new \OpenAPI\Client\Model\InlineObject264(); // \OpenAPI\Client\Model\InlineObject264

try {
    $result = $apiInstance->updateVirtualImage($virtual_image_id, $inline_object264);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling LibraryApi->updateVirtualImage: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtual_image_id** | **float**| Virtual Image ID |
 **inline_object264** | [**\OpenAPI\Client\Model\InlineObject264**](../Model/InlineObject264.md)|  | [optional]

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
