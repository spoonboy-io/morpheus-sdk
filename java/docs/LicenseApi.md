# LicenseApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getLicense**](LicenseApi.md#getLicense) | **GET** /api/license | Get license
[**installLicense**](LicenseApi.md#installLicense) | **POST** /api/license | Install license key
[**testLicense**](LicenseApi.md#testLicense) | **POST** /api/license/test | Test license key
[**uninstallLicense**](LicenseApi.md#uninstallLicense) | **DELETE** /api/license | Uninstall license key


<a name="getLicense"></a>
# **getLicense**
> License getLicense()

Get license

Get details about the license that is currently installed on the appliance. This returns the license details, but not the key itself. Your license key will never be returned and should always be kept secret.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LicenseApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LicenseApi apiInstance = new LicenseApi(defaultClient);
    try {
      License result = apiInstance.getLicense();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LicenseApi#getLicense");
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

[**License**](License.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Get details about the currently installed license |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="installLicense"></a>
# **installLicense**
> License installLicense(inlineObject125)

Install license key

Install a new license key. This will potentially change the enabled features and capabilities of your Morpheus appliance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LicenseApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LicenseApi apiInstance = new LicenseApi(defaultClient);
    InlineObject125 inlineObject125 = new InlineObject125(); // InlineObject125 | 
    try {
      License result = apiInstance.installLicense(inlineObject125);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LicenseApi#installLicense");
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
 **inlineObject125** | [**InlineObject125**](InlineObject125.md)|  | [optional]

### Return type

[**License**](License.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | License installed successfully. |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="testLicense"></a>
# **testLicense**
> License testLicense(inlineObject126)

Test license key

This endpoint can be used to decode a license to see if it is valid and inspect the license settings, such as who it belongs to and the enabled features. This is only a test, it does not install the key, or make any changes to your appliance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LicenseApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LicenseApi apiInstance = new LicenseApi(defaultClient);
    InlineObject126 inlineObject126 = new InlineObject126(); // InlineObject126 | 
    try {
      License result = apiInstance.testLicense(inlineObject126);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LicenseApi#testLicense");
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
 **inlineObject126** | [**InlineObject126**](InlineObject126.md)|  | [optional]

### Return type

[**License**](License.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | License test succeeded, the license details are returned |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="uninstallLicense"></a>
# **uninstallLicense**
> Model200Success uninstallLicense()

Uninstall license key

Uninstall your appliance license, leaving the appliance with no license installed. This will change the enabled features and capabilities of your Morpheus appliance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LicenseApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LicenseApi apiInstance = new LicenseApi(defaultClient);
    try {
      Model200Success result = apiInstance.uninstallLicense();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LicenseApi#uninstallLicense");
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

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | License uninstalled successfully. |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

