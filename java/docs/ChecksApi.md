# ChecksApi

All URIs are relative to *https://CHANGEME*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addCheckApps**](ChecksApi.md#addCheckApps) | **POST** /api/monitoring/apps | Create a New Check App |
| [**listCheckApps**](ChecksApi.md#listCheckApps) | **GET** /api/monitoring/apps | List All Check Apps |


<a id="addCheckApps"></a>
# **addCheckApps**
> AddCheckApps200Response addCheckApps(addCheckAppsRequest)

Create a New Check App

Create a new check app.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ChecksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ChecksApi apiInstance = new ChecksApi(defaultClient);
    AddCheckAppsRequest addCheckAppsRequest = new AddCheckAppsRequest(); // AddCheckAppsRequest | 
    try {
      AddCheckApps200Response result = apiInstance.addCheckApps(addCheckAppsRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#addCheckApps");
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
| **addCheckAppsRequest** | [**AddCheckAppsRequest**](AddCheckAppsRequest.md)|  | [optional] |

### Return type

[**AddCheckApps200Response**](AddCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Check App Object |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="listCheckApps"></a>
# **listCheckApps**
> ListCheckApps200Response listCheckApps(max, offset, sort, name, phrase, status, lastUpdated, deleted)

List All Check Apps

Get a list of check apps.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ChecksApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ChecksApi apiInstance = new ChecksApi(defaultClient);
    Long max = 25L; // Long | Maximum number of records to return
    Long offset = 0L; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "name"; // String | Sort order, the name of the property to sort by
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String status = "running"; // String | The instance status for filtering.
    OffsetDateTime lastUpdated = OffsetDateTime.parse("2019-03-06T17:52:29+0000"); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    Boolean deleted = true; // Boolean | If true, only deleted resources or instances in pending removal status are returned.
    try {
      ListCheckApps200Response result = apiInstance.listCheckApps(max, offset, sort, name, phrase, status, lastUpdated, deleted);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#listCheckApps");
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
| **max** | **Long**| Maximum number of records to return | [optional] [default to 25] |
| **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0] |
| **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to name] |
| **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] |
| **phrase** | **String**| Search phrase for partial matches on name or description | [optional] |
| **status** | **String**| The instance status for filtering. | [optional] |
| **lastUpdated** | **OffsetDateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] |
| **deleted** | **Boolean**| If true, only deleted resources or instances in pending removal status are returned. | [optional] |

### Return type

[**ListCheckApps200Response**](ListCheckApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Array of Check Apps |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

