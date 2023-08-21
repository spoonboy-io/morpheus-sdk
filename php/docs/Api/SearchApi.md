# OpenAPI\Client\SearchApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**search()**](SearchApi.md#search) | **GET** /api/search | Provides global search for all types of objects


## `search()`

```php
search($max, $offset, $phrase, $query): \OpenAPI\Client\Model\Search
```

Provides global search for all types of objects

This endpoint provides global search for all types of objects, with the results sorted in order of relevance.  The `phrase` parameter can be specified as part of the URL or as a query parameter. If phrase is not specified, then 0 results (hits) will be returned.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\SearchApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$max = 25; // int | Maximum number of records to return, -1 can be used to fetch all records
$offset = 0; // int | Offset records, the number of records to skip, for paginating requests
$phrase = 'phrase_example'; // string | Search phrase for partial matches on name or description
$query = 'query_example'; // string | Alias for phrase

try {
    $result = $apiInstance->search($max, $offset, $phrase, $query);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SearchApi->search: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **phrase** | **string**| Search phrase for partial matches on name or description | [optional]
 **query** | **string**| Alias for phrase | [optional]

### Return type

[**\OpenAPI\Client\Model\Search**](../Model/Search.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
