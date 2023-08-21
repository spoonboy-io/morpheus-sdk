# WhitelabelSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getWhitelabelImage**](WhitelabelSettingsApi.md#getWhitelabelImage) | **GET** /api/whitelabel-settings/images/{imageType} | Download Image
[**listWhitelabelSettings**](WhitelabelSettingsApi.md#listWhitelabelSettings) | **GET** /api/whitelabel-settings | Get Whitelabel Settings
[**removeWhitelabelImage**](WhitelabelSettingsApi.md#removeWhitelabelImage) | **DELETE** /api/whitelabel-settings/images/{imageType} | Reset Image
[**updateWhitelabelImages**](WhitelabelSettingsApi.md#updateWhitelabelImages) | **POST** /api/whitelabel-settings/images | Update Images
[**updateWhitelabelSettings**](WhitelabelSettingsApi.md#updateWhitelabelSettings) | **PUT** /api/whitelabel-settings | Update Whitelabel Settings


<a name="getWhitelabelImage"></a>
# **getWhitelabelImage**
> File getWhitelabelImage(imageType)

Download Image

Downloads the specified image.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.WhitelabelSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    WhitelabelSettingsApi apiInstance = new WhitelabelSettingsApi(defaultClient);
    String imageType = "headerLogo"; // String | Valid image types
    try {
      File result = apiInstance.getWhitelabelImage(imageType);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling WhitelabelSettingsApi#getWhitelabelImage");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageType** | **String**| Valid image types | [enum: headerLogo, footerLogo, loginLogo, favicon]

### Return type

[**File**](File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/ico, image/jpeg, image/png, image/svg+xml, application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listWhitelabelSettings"></a>
# **listWhitelabelSettings**
> InlineResponse200166 listWhitelabelSettings()

Get Whitelabel Settings

This endpoint retrieves whitelabel settings.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.WhitelabelSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    WhitelabelSettingsApi apiInstance = new WhitelabelSettingsApi(defaultClient);
    try {
      InlineResponse200166 result = apiInstance.listWhitelabelSettings();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling WhitelabelSettingsApi#listWhitelabelSettings");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200166**](InlineResponse200166.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="removeWhitelabelImage"></a>
# **removeWhitelabelImage**
> Model200Success removeWhitelabelImage(imageType)

Reset Image

Resets the specified image to the Morpheus default.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.WhitelabelSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    WhitelabelSettingsApi apiInstance = new WhitelabelSettingsApi(defaultClient);
    String imageType = "headerLogo"; // String | Valid image types
    try {
      Model200Success result = apiInstance.removeWhitelabelImage(imageType);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling WhitelabelSettingsApi#removeWhitelabelImage");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageType** | **String**| Valid image types | [enum: headerLogo, footerLogo, loginLogo, favicon]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateWhitelabelImages"></a>
# **updateWhitelabelImages**
> Model200Success updateWhitelabelImages(headerLogoFile, resetHeaderLogo, footerLogoFile, resetFooterLogo, loginLogoFile, resetLoginLogo, faviconFile, resetFaviconLogo)

Update Images

Uploads whitelabel images. Expects multipart form data as the request format, not JSON.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.WhitelabelSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    WhitelabelSettingsApi apiInstance = new WhitelabelSettingsApi(defaultClient);
    File headerLogoFile = new File("/path/to/file"); // File | Header logo image file, valid image types `png|jpg|svg`
    Boolean resetHeaderLogo = true; // Boolean | Resets header logo to default
    File footerLogoFile = new File("/path/to/file"); // File | Footer logo image file, valid image types `png|jpg|svg`
    Boolean resetFooterLogo = true; // Boolean | Resets footer logo to default
    File loginLogoFile = new File("/path/to/file"); // File | Login logo image file, valid image types `png|jpg|svg`
    Boolean resetLoginLogo = true; // Boolean | Resets login logo to default
    File faviconFile = new File("/path/to/file"); // File | Favicon image file, valid image type ico
    Boolean resetFaviconLogo = true; // Boolean | Resets favicon logo to default
    try {
      Model200Success result = apiInstance.updateWhitelabelImages(headerLogoFile, resetHeaderLogo, footerLogoFile, resetFooterLogo, loginLogoFile, resetLoginLogo, faviconFile, resetFaviconLogo);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling WhitelabelSettingsApi#updateWhitelabelImages");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **headerLogoFile** | **File**| Header logo image file, valid image types &#x60;png|jpg|svg&#x60; | [optional]
 **resetHeaderLogo** | **Boolean**| Resets header logo to default | [optional]
 **footerLogoFile** | **File**| Footer logo image file, valid image types &#x60;png|jpg|svg&#x60; | [optional]
 **resetFooterLogo** | **Boolean**| Resets footer logo to default | [optional]
 **loginLogoFile** | **File**| Login logo image file, valid image types &#x60;png|jpg|svg&#x60; | [optional]
 **resetLoginLogo** | **Boolean**| Resets login logo to default | [optional]
 **faviconFile** | **File**| Favicon image file, valid image type ico | [optional]
 **resetFaviconLogo** | **Boolean**| Resets favicon logo to default | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateWhitelabelSettings"></a>
# **updateWhitelabelSettings**
> Model200Success updateWhitelabelSettings(inlineObject265)

Update Whitelabel Settings

Update Whitelabel Settings

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.WhitelabelSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    WhitelabelSettingsApi apiInstance = new WhitelabelSettingsApi(defaultClient);
    InlineObject265 inlineObject265 = new InlineObject265(); // InlineObject265 | 
    try {
      Model200Success result = apiInstance.updateWhitelabelSettings(inlineObject265);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling WhitelabelSettingsApi#updateWhitelabelSettings");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject265** | [**InlineObject265**](InlineObject265.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

