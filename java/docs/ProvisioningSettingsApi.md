# ProvisioningSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listProvisioningSettings**](ProvisioningSettingsApi.md#listProvisioningSettings) | **GET** /api/provisioning-settings | Get Provisioning Settings
[**updateProvisioningSettings**](ProvisioningSettingsApi.md#updateProvisioningSettings) | **PUT** /api/provisioning-settings | Update Provisioning Settings


<a name="listProvisioningSettings"></a>
# **listProvisioningSettings**
> InlineResponse200128 listProvisioningSettings()

Get Provisioning Settings

This endpoint retrieves provisioning settings.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProvisioningSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ProvisioningSettingsApi apiInstance = new ProvisioningSettingsApi(defaultClient);
    try {
      InlineResponse200128 result = apiInstance.listProvisioningSettings();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProvisioningSettingsApi#listProvisioningSettings");
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

[**InlineResponse200128**](InlineResponse200128.md)

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

<a name="updateProvisioningSettings"></a>
# **updateProvisioningSettings**
> Model200Success updateProvisioningSettings(inlineObject204)

Update Provisioning Settings

Update Provisioning Settings

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProvisioningSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ProvisioningSettingsApi apiInstance = new ProvisioningSettingsApi(defaultClient);
    InlineObject204 inlineObject204 = new InlineObject204(); // InlineObject204 | 
    try {
      Model200Success result = apiInstance.updateProvisioningSettings(inlineObject204);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProvisioningSettingsApi#updateProvisioningSettings");
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
 **inlineObject204** | [**InlineObject204**](InlineObject204.md)|  | [optional]

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

