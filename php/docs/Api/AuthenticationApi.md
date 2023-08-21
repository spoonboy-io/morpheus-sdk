# OpenAPI\Client\AuthenticationApi

All URIs are relative to https://CHANGEME.

Method | HTTP request | Description
------------- | ------------- | -------------
[**forgotPassword()**](AuthenticationApi.md#forgotPassword) | **POST** /api/forgot/send-email | Request a reset password email
[**getAccessToken()**](AuthenticationApi.md#getAccessToken) | **POST** /oauth/token | Provides authentication via username and password
[**resetPassword()**](AuthenticationApi.md#resetPassword) | **POST** /api/forgot/reset-password | Reset user password
[**whoami()**](AuthenticationApi.md#whoami) | **GET** /api/whoami | Retrieves information about current user roles and permissions


## `forgotPassword()`

```php
forgotPassword($inline_object11): \OpenAPI\Client\Model\InlineResponse2007
```

Request a reset password email

This endpoint will trigger the Reset your password email to be sent to the specified user.  The User is identified by `username` and, if they exist, will be notified via their configured email address. The email notification will indicate a Reset Password Request was made and it will include a token.  Once you obtain the token from the email, it may be used to reset the password of your user.  **Note**: This is an unauthorized endpoint and the response will always appear successful, it is not possible determine from the response whether the user exists or if an email was sent.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');



$apiInstance = new OpenAPI\Client\Api\AuthenticationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$inline_object11 = new \OpenAPI\Client\Model\InlineObject11(); // \OpenAPI\Client\Model\InlineObject11

try {
    $result = $apiInstance->forgotPassword($inline_object11);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AuthenticationApi->forgotPassword: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object11** | [**\OpenAPI\Client\Model\InlineObject11**](../Model/InlineObject11.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse2007**](../Model/InlineResponse2007.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `getAccessToken()`

```php
getAccessToken($client_id, $grant_type, $scope, $username, $password, $refresh_token): \OpenAPI\Client\Model\AccessToken
```

Provides authentication via username and password

This endpoint provides authentication via `username` and `password` of a Morpheus User. The response includes a valid access token. If your current token is expired, a new one will be created and returned.  Subtenant users will need to pass their subdomain prefix like subdomain\\username. The default subdomain is the tenant account ID.  This endpoint also allows refreshing your current access token to get a new token. This is done by passing your current `refresh_token`. This provides a way to renew your client’s session with the API, and extend the expiration date.  This will render your current access token invalid, so you will need to update any scripts relying on it.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');



$apiInstance = new OpenAPI\Client\Api\AuthenticationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$client_id = 'client_id_example'; // string | Client ID, use morph-api. Users may only have one access token per Client ID. The CLI uses morph-cli.
$grant_type = 'grant_type_example'; // string | OAuth Grant Type, use password.
$scope = 'scope_example'; // string | OAuth token scope, use write.
$username = 'username_example'; // string | Specified as \\\"username\\\" or \\\"tenantId\\\\username\\\" with proper HTML encoding and used in conjunction with `password`. Not utilized with `refresh_token`.
$password = 'password_example'; // string | The Password for defined `username`. Must have proper HTML encoding
$refresh_token = NULL; // mixed | The `refresh_token` value from a previous API generation can be utilized instead of `username` and `password`.

try {
    $result = $apiInstance->getAccessToken($client_id, $grant_type, $scope, $username, $password, $refresh_token);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AuthenticationApi->getAccessToken: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **client_id** | **string**| Client ID, use morph-api. Users may only have one access token per Client ID. The CLI uses morph-cli. |
 **grant_type** | **string**| OAuth Grant Type, use password. |
 **scope** | **string**| OAuth token scope, use write. |
 **username** | **string**| Specified as \\\&quot;username\\\&quot; or \\\&quot;tenantId\\\\username\\\&quot; with proper HTML encoding and used in conjunction with &#x60;password&#x60;. Not utilized with &#x60;refresh_token&#x60;. | [optional]
 **password** | **string**| The Password for defined &#x60;username&#x60;. Must have proper HTML encoding | [optional]
 **refresh_token** | [**mixed**](../Model/mixed.md)| The &#x60;refresh_token&#x60; value from a previous API generation can be utilized instead of &#x60;username&#x60; and &#x60;password&#x60;. | [optional]

### Return type

[**\OpenAPI\Client\Model\AccessToken**](../Model/AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/x-www-form-urlencoded`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `resetPassword()`

```php
resetPassword($inline_object10): \OpenAPI\Client\Model\InlineResponse2006
```

Reset user password

This endpoint will reset the password for a user, updating it to the specified value. A secret token must be passed to identify the user who is being updated.  **Note**: You can obtain this token by inspecting the URL of the “Click here to reset” link seen in the email.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');



$apiInstance = new OpenAPI\Client\Api\AuthenticationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$inline_object10 = new \OpenAPI\Client\Model\InlineObject10(); // \OpenAPI\Client\Model\InlineObject10

try {
    $result = $apiInstance->resetPassword($inline_object10);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AuthenticationApi->resetPassword: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inline_object10** | [**\OpenAPI\Client\Model\InlineObject10**](../Model/InlineObject10.md)|  | [optional]

### Return type

[**\OpenAPI\Client\Model\InlineResponse2006**](../Model/InlineResponse2006.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `whoami()`

```php
whoami(): \OpenAPI\Client\Model\InlineResponse200167
```

Retrieves information about current user roles and permissions

Provides API to retrieve information about yourself, including your roles and permissions.  The appliance build version is also returned.

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure Bearer authorization: bearerAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setAccessToken('YOUR_ACCESS_TOKEN');


$apiInstance = new OpenAPI\Client\Api\AuthenticationApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->whoami();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AuthenticationApi->whoami: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\InlineResponse200167**](../Model/InlineResponse200167.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
