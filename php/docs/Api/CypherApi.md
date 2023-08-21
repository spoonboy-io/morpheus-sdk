# OpenAPI\Client\CypherApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCypherKey()**](CypherApi.md#addCypherKey) | **POST** /api/cypher/{cypherPath} | Write a Cypher
[**getCypherKey()**](CypherApi.md#getCypherKey) | **GET** /api/cypher/{cypherPath} | Read or Create a Cypher Key
[**listCypherKeys()**](CypherApi.md#listCypherKeys) | **GET** /api/cypher | List Cypher Keys
[**removeCypher()**](CypherApi.md#removeCypher) | **DELETE** /api/cypher/{cypherPath} | Delete a Cypher


## `addCypherKey()`

```php
addCypherKey($cypher_path, $ttl, $value, $type, $inline_object66): object
```

Write a Cypher

This endpoint will create or update a cypher key.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CypherApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cypher_path = 'cypher_path_example'; // string | The key includes a mount prefix separated by a /. For example, the key secret/foo uses the secret mount.  Available Mounts  <table>   <tr>     <th>Mount</th>     <th>Description</th>     <th>Example</th>   </tr>   <tr>     <td>password</td>     <td>Generates a secure password of specified character length in the key pattern (or 15) with symbols, numbers, upper case, and lower case letters (i.e. password/15/mypass generates a 15 character password).</td>     <td>password/15/mypass</td>   </tr>   <tr>     <td>tfvars</td>     <td>This is a module to store a tfvars file for terraform.</td>     <td>tfvars/mytfvar</td>   </tr>   <tr>     <td>secret</td>     <td>This is the standard secret module that stores a key/value in encrypted form. Capable of storing entire JSON object or a String.</td>     <td>secret/foo</td>   </tr>   <tr>     <td>uuid</td>     <td>Returns a new UUID by key name when requested and stores the generated UUID by key name for a given lease timeout period.</td>     <td>uuid/autoMac1</td>   </tr>   <tr>     <td>key</td>     <td>Generates a Base 64 encoded AES Key of specified bit length in the key pattern (i.e. key/128/mykey generates a 128-bit key)</td>     <td>key/128/mykey</td>   </tr> </table>
$ttl = 30d; // OneOfStringLong | Time to Live. The lease duration in seconds, or a human readable format eg. '15m', 8h, '7d'.  0 means no expiry.
$value = {"foo":"bar", "zip":"zap"}; // string | The secret value to be stored. Only required for certain mounts. Some mounts generate their own value and do not require a value to be passed. eg. `uuid`, `key` and `password`.
$type = 'type_example'; // string | The type of data being stored, `string` or `object`. The data type depends on the cypher mount being used. Most mounts use `string` as their data type, but `secret` uses `object` by default. You can store a string instead by passing `type=string`. This means the `data` value returned by the API will be a string instead of an object.
$inline_object66 = new \OpenAPI\Client\Model\InlineObject66(); // \OpenAPI\Client\Model\InlineObject66

try {
    $result = $apiInstance->addCypherKey($cypher_path, $ttl, $value, $type, $inline_object66);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CypherApi->addCypherKey: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cypher_path** | **string**| The key includes a mount prefix separated by a /. For example, the key secret/foo uses the secret mount.  Available Mounts  &lt;table&gt;   &lt;tr&gt;     &lt;th&gt;Mount&lt;/th&gt;     &lt;th&gt;Description&lt;/th&gt;     &lt;th&gt;Example&lt;/th&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;password&lt;/td&gt;     &lt;td&gt;Generates a secure password of specified character length in the key pattern (or 15) with symbols, numbers, upper case, and lower case letters (i.e. password/15/mypass generates a 15 character password).&lt;/td&gt;     &lt;td&gt;password/15/mypass&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;tfvars&lt;/td&gt;     &lt;td&gt;This is a module to store a tfvars file for terraform.&lt;/td&gt;     &lt;td&gt;tfvars/mytfvar&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;secret&lt;/td&gt;     &lt;td&gt;This is the standard secret module that stores a key/value in encrypted form. Capable of storing entire JSON object or a String.&lt;/td&gt;     &lt;td&gt;secret/foo&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;uuid&lt;/td&gt;     &lt;td&gt;Returns a new UUID by key name when requested and stores the generated UUID by key name for a given lease timeout period.&lt;/td&gt;     &lt;td&gt;uuid/autoMac1&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;key&lt;/td&gt;     &lt;td&gt;Generates a Base 64 encoded AES Key of specified bit length in the key pattern (i.e. key/128/mykey generates a 128-bit key)&lt;/td&gt;     &lt;td&gt;key/128/mykey&lt;/td&gt;   &lt;/tr&gt; &lt;/table&gt; |
 **ttl** | [**OneOfStringLong**](../Model/.md)| Time to Live. The lease duration in seconds, or a human readable format eg. &#39;15m&#39;, 8h, &#39;7d&#39;.  0 means no expiry. | [optional]
 **value** | **string**| The secret value to be stored. Only required for certain mounts. Some mounts generate their own value and do not require a value to be passed. eg. &#x60;uuid&#x60;, &#x60;key&#x60; and &#x60;password&#x60;. | [optional]
 **type** | **string**| The type of data being stored, &#x60;string&#x60; or &#x60;object&#x60;. The data type depends on the cypher mount being used. Most mounts use &#x60;string&#x60; as their data type, but &#x60;secret&#x60; uses &#x60;object&#x60; by default. You can store a string instead by passing &#x60;type&#x3D;string&#x60;. This means the &#x60;data&#x60; value returned by the API will be a string instead of an object. | [optional]
 **inline_object66** | [**\OpenAPI\Client\Model\InlineObject66**](../Model/InlineObject66.md)|  | [optional]

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

## `getCypherKey()`

```php
getCypherKey($cypher_path, $lease_token, $sort, $direction): Model200Success
```

Read or Create a Cypher Key

This endpoint retrieves a specific cypher key. The value of the key is decrypted and returned as data. It may be a String or an object with many {\"key\":\"value\"} pairs.  The type depends on the cypher mount's capabilities and what type of data was written to the key.  For example the `secret/` mount allows either a string or an object, while the `password/` mount will always store and return a string. This endpoint can also create a key. This only applies to mount types `uuid`, `key`, `password`.  Refer to the `POST` endpoint for more information.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');

// Configure API key authorization: cypherAuth-XCToken
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('X-Cypher-Token', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('X-Cypher-Token', 'Bearer');

// Configure API key authorization: cypherAuth-XMLease
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('X-Morpheus-Lease', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('X-Morpheus-Lease', 'Bearer');

// Configure API key authorization: cypherAuth-XVToken
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('X-Vault-Token', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('X-Vault-Token', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\CypherApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cypher_path = password/15/mypass; // string | The cypher key including the mount prefix.
$lease_token = 5bd808dc-82bb-4974-b699-5f1cdd3cc503; // string | An execution lease token.
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort

try {
    $result = $apiInstance->getCypherKey($cypher_path, $lease_token, $sort, $direction);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CypherApi->getCypherKey: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cypher_path** | **string**| The cypher key including the mount prefix. |
 **lease_token** | **string**| An execution lease token. | [optional]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]

### Return type

[**Model200Success**](../Model/Model200Success.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth), [cypherAuth-XCToken](../../README.md#cypherAuth-XCToken), [cypherAuth-XMLease](../../README.md#cypherAuth-XMLease), [cypherAuth-XVToken](../../README.md#cypherAuth-XVToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `listCypherKeys()`

```php
listCypherKeys($lease_token, $list, $phrase, $max, $offset, $sort, $direction): object
```

List Cypher Keys

This endpoint retrieves all cypher keys associated with the account, or user.  This method can be used to list keys as well, by passing the query parameter list=true.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');

// Configure API key authorization: cypherAuth-XCToken
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('X-Cypher-Token', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('X-Cypher-Token', 'Bearer');

// Configure API key authorization: cypherAuth-XMLease
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('X-Morpheus-Lease', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('X-Morpheus-Lease', 'Bearer');

// Configure API key authorization: cypherAuth-XVToken
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('X-Vault-Token', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('X-Vault-Token', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\CypherApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$lease_token = 5bd808dc-82bb-4974-b699-5f1cdd3cc503; // string | An execution lease token.
$list = false; // bool | This endpoint is available via the http method LIST. The GET method can be used to list keys as well, by passing the query parameter list=true.
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$sort = 'name'; // string | Sort order, the name of the property to sort by
$direction = asc; // string | Sort direction, use 'desc' to reverse sort

try {
    $result = $apiInstance->listCypherKeys($lease_token, $list, $phrase, $max, $offset, $sort, $direction);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CypherApi->listCypherKeys: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lease_token** | **string**| An execution lease token. | [optional]
 **list** | **bool**| This endpoint is available via the http method LIST. The GET method can be used to list keys as well, by passing the query parameter list&#x3D;true. | [optional] [default to false]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **string**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **string**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]

### Return type

**object**

### Authorization

[bearerAuth](../../README.md#bearerAuth), [cypherAuth-XCToken](../../README.md#cypherAuth-XCToken), [cypherAuth-XMLease](../../README.md#cypherAuth-XMLease), [cypherAuth-XVToken](../../README.md#cypherAuth-XVToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `removeCypher()`

```php
removeCypher($cypher_path): \OpenAPI\Client\Model\Model200Success
```

Delete a Cypher

Will delete a cypher from the system and make it no longer usable.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\CypherApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$cypher_path = password/15/mypass; // string | The cypher key including the mount prefix.

try {
    $result = $apiInstance->removeCypher($cypher_path);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CypherApi->removeCypher: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cypher_path** | **string**| The cypher key including the mount prefix. |

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
