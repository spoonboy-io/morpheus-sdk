# LogSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addLogSettingsSyslogRules**](LogSettingsApi.md#addLogSettingsSyslogRules) | **POST** /api/log-settings/syslog-rules | Create a New Syslog Rule
[**deleteLogSettingsSyslogRules**](LogSettingsApi.md#deleteLogSettingsSyslogRules) | **DELETE** /api/log-settings/syslog-rules/{id} | Delete a Specific Syslog Rule
[**listLogSettings**](LogSettingsApi.md#listLogSettings) | **GET** /api/log-settings | List All Log Settings
[**updateLogSettings**](LogSettingsApi.md#updateLogSettings) | **PUT** /api/log-settings | Update Log Settings


<a name="addLogSettingsSyslogRules"></a>
# **addLogSettingsSyslogRules**
> Model200Success addLogSettingsSyslogRules(inlineObject140)

Create a New Syslog Rule

Creates a new syslog rule. This command will also update existing syslog rule by specified name if already exists

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LogSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LogSettingsApi apiInstance = new LogSettingsApi(defaultClient);
    InlineObject140 inlineObject140 = new InlineObject140(); // InlineObject140 | 
    try {
      Model200Success result = apiInstance.addLogSettingsSyslogRules(inlineObject140);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LogSettingsApi#addLogSettingsSyslogRules");
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
 **inlineObject140** | [**InlineObject140**](InlineObject140.md)|  | [optional]

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
**200** | Syslog Rule Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteLogSettingsSyslogRules"></a>
# **deleteLogSettingsSyslogRules**
> Model200Success deleteLogSettingsSyslogRules(id)

Delete a Specific Syslog Rule

Will delete the syslog rule matching the specified name.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LogSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LogSettingsApi apiInstance = new LogSettingsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteLogSettingsSyslogRules(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LogSettingsApi#deleteLogSettingsSyslogRules");
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
 **id** | **Long**| Morpheus ID of the Object being referenced |

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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listLogSettings"></a>
# **listLogSettings**
> InlineResponse20084 listLogSettings()

List All Log Settings

This endpoint retrieves log settings.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LogSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LogSettingsApi apiInstance = new LogSettingsApi(defaultClient);
    try {
      InlineResponse20084 result = apiInstance.listLogSettings();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LogSettingsApi#listLogSettings");
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

[**InlineResponse20084**](InlineResponse20084.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Log Settings |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateLogSettings"></a>
# **updateLogSettings**
> Model200Success updateLogSettings(inlineObject139)

Update Log Settings

Update Log Settings

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LogSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LogSettingsApi apiInstance = new LogSettingsApi(defaultClient);
    InlineObject139 inlineObject139 = new InlineObject139(); // InlineObject139 | 
    try {
      Model200Success result = apiInstance.updateLogSettings(inlineObject139);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LogSettingsApi#updateLogSettings");
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
 **inlineObject139** | [**InlineObject139**](InlineObject139.md)|  | [optional]

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
**200** | Update Log Settings Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

