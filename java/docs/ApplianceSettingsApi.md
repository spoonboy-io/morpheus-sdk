# ApplianceSettingsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listApplianceSettings**](ApplianceSettingsApi.md#listApplianceSettings) | **GET** /api/appliance-settings | Get Appliance Settings
[**reindex**](ApplianceSettingsApi.md#reindex) | **POST** /api/appliance-settings/reindex | Reindex Search
[**setApplianceSettingsMaintenanceMode**](ApplianceSettingsApi.md#setApplianceSettingsMaintenanceMode) | **POST** /api/appliance-settings/maintenance | Toggle Maintenance Mode
[**updateApplianceSettings**](ApplianceSettingsApi.md#updateApplianceSettings) | **PUT** /api/appliance-settings | Update Appliance Settings


<a name="listApplianceSettings"></a>
# **listApplianceSettings**
> InlineResponse200 listApplianceSettings()

Get Appliance Settings

This endpoint retrieves appliance settings.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ApplianceSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ApplianceSettingsApi apiInstance = new ApplianceSettingsApi(defaultClient);
    try {
      InlineResponse200 result = apiInstance.listApplianceSettings();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ApplianceSettingsApi#listApplianceSettings");
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

[**InlineResponse200**](InlineResponse200.md)

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

<a name="reindex"></a>
# **reindex**
> Model200Success reindex()

Reindex Search

Reindex all searchable data

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ApplianceSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ApplianceSettingsApi apiInstance = new ApplianceSettingsApi(defaultClient);
    try {
      Model200Success result = apiInstance.reindex();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ApplianceSettingsApi#reindex");
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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="setApplianceSettingsMaintenanceMode"></a>
# **setApplianceSettingsMaintenanceMode**
> Model200Success setApplianceSettingsMaintenanceMode(enabled)

Toggle Maintenance Mode

This endpoint allows toggling the appliance maintenance mode.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ApplianceSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ApplianceSettingsApi apiInstance = new ApplianceSettingsApi(defaultClient);
    Boolean enabled = true; // Boolean | Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa.
    try {
      Model200Success result = apiInstance.setApplianceSettingsMaintenanceMode(enabled);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ApplianceSettingsApi#setApplianceSettingsMaintenanceMode");
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
 **enabled** | **Boolean**| Pass true to turn on maintenance mode, or false to turn it off. If no value is given then it will be toggled from off to on or vice versa. | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateApplianceSettings"></a>
# **updateApplianceSettings**
> Model200Success updateApplianceSettings(inlineObject2)

Update Appliance Settings

Update Appliance Settings

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ApplianceSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ApplianceSettingsApi apiInstance = new ApplianceSettingsApi(defaultClient);
    InlineObject2 inlineObject2 = new InlineObject2(); // InlineObject2 | 
    try {
      Model200Success result = apiInstance.updateApplianceSettings(inlineObject2);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ApplianceSettingsApi#updateApplianceSettings");
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
 **inlineObject2** | [**InlineObject2**](InlineObject2.md)|  | [optional]

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

