# HistoryApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getHistory**](HistoryApi.md#getHistory) | **GET** /api/processes/{id} | Retrieves a Specific Process
[**getHistoryEvent**](HistoryApi.md#getHistoryEvent) | **GET** /api/processes/events/{id} | Retrieves a Specific Process Event
[**listHistory**](HistoryApi.md#listHistory) | **GET** /api/processes | Retrieves Process History


<a name="getHistory"></a>
# **getHistory**
> InlineResponse20048 getHistory(id)

Retrieves a Specific Process

Retrieves a specific process. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HistoryApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HistoryApi apiInstance = new HistoryApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20048 result = apiInstance.getHistory(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HistoryApi#getHistory");
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

[**InlineResponse20048**](InlineResponse20048.md)

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

<a name="getHistoryEvent"></a>
# **getHistoryEvent**
> InlineResponse20049 getHistoryEvent(id)

Retrieves a Specific Process Event

Retrieves a specific process event. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HistoryApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HistoryApi apiInstance = new HistoryApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20049 result = apiInstance.getHistoryEvent(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HistoryApi#getHistoryEvent");
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

[**InlineResponse20049**](InlineResponse20049.md)

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

<a name="listHistory"></a>
# **listHistory**
> Object listHistory(instanceId, containerId, serverId, zoneId, appId, phrase)

Retrieves Process History

Retrieves process history for objects 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HistoryApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HistoryApi apiInstance = new HistoryApi(defaultClient);
    Long instanceId = 94L; // Long | The Instance ID for Filtering
    Long containerId = 135L; // Long | The Container ID for Filtering
    Long serverId = 97L; // Long | The Server ID for Filtering
    Long zoneId = 3L; // Long | The Zone ID for Filtering
    Long appId = 5L; // Long | The App ID for Filtering
    String phrase = "phrase_example"; // String | Search phrase for partial matches on message, displayName, output, event.message, event.output or event.error
    try {
      Object result = apiInstance.listHistory(instanceId, containerId, serverId, zoneId, appId, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HistoryApi#listHistory");
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
 **instanceId** | **Long**| The Instance ID for Filtering | [optional]
 **containerId** | **Long**| The Container ID for Filtering | [optional]
 **serverId** | **Long**| The Server ID for Filtering | [optional]
 **zoneId** | **Long**| The Zone ID for Filtering | [optional]
 **appId** | **Long**| The App ID for Filtering | [optional]
 **phrase** | **String**| Search phrase for partial matches on message, displayName, output, event.message, event.output or event.error | [optional]

### Return type

**Object**

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

