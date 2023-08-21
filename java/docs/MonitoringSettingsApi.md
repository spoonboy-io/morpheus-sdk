# MonitoringSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getMonitoringSettings**](MonitoringSettingsApi.md#getMonitoringSettings) | **GET** /api/monitoring-settings | Get Monitoring Settings
[**updateMonitoringSettings**](MonitoringSettingsApi.md#updateMonitoringSettings) | **PUT** /api/monitoring-settings | Update Monitoring Settings


<a name="getMonitoringSettings"></a>
# **getMonitoringSettings**
> InlineResponse20085 getMonitoringSettings()

Get Monitoring Settings

This endpoint retrieves monitoring settings.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.MonitoringSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    MonitoringSettingsApi apiInstance = new MonitoringSettingsApi(defaultClient);
    try {
      InlineResponse20085 result = apiInstance.getMonitoringSettings();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling MonitoringSettingsApi#getMonitoringSettings");
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

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Monitoring Settings |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateMonitoringSettings"></a>
# **updateMonitoringSettings**
> Model200Success updateMonitoringSettings(inlineObject141)

Update Monitoring Settings

Update Monitoring Settings

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.MonitoringSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    MonitoringSettingsApi apiInstance = new MonitoringSettingsApi(defaultClient);
    InlineObject141 inlineObject141 = new InlineObject141(); // InlineObject141 | 
    try {
      Model200Success result = apiInstance.updateMonitoringSettings(inlineObject141);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling MonitoringSettingsApi#updateMonitoringSettings");
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
 **inlineObject141** | [**InlineObject141**](InlineObject141.md)|  | [optional]

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
**200** | Update Monitoring Settings Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

