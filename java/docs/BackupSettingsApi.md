# BackupSettingsApi

All URIs are relative to *https://CHANGEME*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**listBackupSettings**](BackupSettingsApi.md#listBackupSettings) | **GET** /api/backup-settings | Get Backup Settings |
| [**updateBackupSettings**](BackupSettingsApi.md#updateBackupSettings) | **PUT** /api/backup-settings | Update Backup Settings |


<a id="listBackupSettings"></a>
# **listBackupSettings**
> ListBackupSettings200Response listBackupSettings()

Get Backup Settings

This endpoint retrieves backup settings.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BackupSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BackupSettingsApi apiInstance = new BackupSettingsApi(defaultClient);
    try {
      ListBackupSettings200Response result = apiInstance.listBackupSettings();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BackupSettingsApi#listBackupSettings");
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

[**ListBackupSettings200Response**](ListBackupSettings200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="updateBackupSettings"></a>
# **updateBackupSettings**
> Model200Success updateBackupSettings(updateBackupSettingsRequest)

Update Backup Settings

Update Backup Settings

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BackupSettingsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BackupSettingsApi apiInstance = new BackupSettingsApi(defaultClient);
    UpdateBackupSettingsRequest updateBackupSettingsRequest = new UpdateBackupSettingsRequest(); // UpdateBackupSettingsRequest | 
    try {
      Model200Success result = apiInstance.updateBackupSettings(updateBackupSettingsRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BackupSettingsApi#updateBackupSettings");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **updateBackupSettingsRequest** | [**UpdateBackupSettingsRequest**](UpdateBackupSettingsRequest.md)|  | [optional] |

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
| **200** | Successful Response |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

