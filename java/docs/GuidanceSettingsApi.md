# GuidanceSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getGuidanceSettings**](GuidanceSettingsApi.md#getGuidanceSettings) | **GET** /api/guidance-settings | Get Guidance Settings
[**updateGuidanceSettings**](GuidanceSettingsApi.md#updateGuidanceSettings) | **PUT** /api/guidance-settings | Update Guidance Settings


<a name="getGuidanceSettings"></a>
# **getGuidanceSettings**
> InlineResponse20047 getGuidanceSettings()

Get Guidance Settings

This endpoint retrieves guidance settings.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.GuidanceSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    GuidanceSettingsApi apiInstance = new GuidanceSettingsApi(defaultClient);
    try {
      InlineResponse20047 result = apiInstance.getGuidanceSettings();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling GuidanceSettingsApi#getGuidanceSettings");
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

[**InlineResponse20047**](InlineResponse20047.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Guidance Settings |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateGuidanceSettings"></a>
# **updateGuidanceSettings**
> Model200Success updateGuidanceSettings(inlineObject79)

Update Guidance Settings

Update Guidance Settings

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.GuidanceSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    GuidanceSettingsApi apiInstance = new GuidanceSettingsApi(defaultClient);
    InlineObject79 inlineObject79 = new InlineObject79(); // InlineObject79 | 
    try {
      Model200Success result = apiInstance.updateGuidanceSettings(inlineObject79);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling GuidanceSettingsApi#updateGuidanceSettings");
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
 **inlineObject79** | [**InlineObject79**](InlineObject79.md)|  | [optional]

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
**200** | Update Guidance Settings Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

