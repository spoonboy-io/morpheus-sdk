# AutomationApi

All URIs are relative to *https://CHANGEME*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addExecuteSchedules**](AutomationApi.md#addExecuteSchedules) | **POST** /api/execute-schedules | Creates a Execute Schedule |
| [**executeExecutionRequest**](AutomationApi.md#executeExecutionRequest) | **POST** /api/execution-request/execute | Executes an Execution Request |
| [**getExecuteSchedules**](AutomationApi.md#getExecuteSchedules) | **GET** /api/execute-schedules/{id} | Retrieves a Specific Execute Schedule |
| [**getExecutionRequest**](AutomationApi.md#getExecutionRequest) | **GET** /api/execution-request/{uniqueId} | Retrieves a Specific Execution Request |
| [**listExecuteSchedules**](AutomationApi.md#listExecuteSchedules) | **GET** /api/execute-schedules | Retrieves all Execute Schedules |
| [**removeExecuteSchedules**](AutomationApi.md#removeExecuteSchedules) | **DELETE** /api/execute-schedules/{id} | Deletes a Execute Schedule |
| [**updateExecuteSchedules**](AutomationApi.md#updateExecuteSchedules) | **PUT** /api/execute-schedules/{id} | Updates a Execute Schedule |


<a id="addExecuteSchedules"></a>
# **addExecuteSchedules**
> AddExecuteSchedules200Response addExecuteSchedules(addExecuteSchedulesRequest)

Creates a Execute Schedule

Creates a execute schedule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AutomationApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AutomationApi apiInstance = new AutomationApi(defaultClient);
    AddExecuteSchedulesRequest addExecuteSchedulesRequest = new AddExecuteSchedulesRequest(); // AddExecuteSchedulesRequest | 
    try {
      AddExecuteSchedules200Response result = apiInstance.addExecuteSchedules(addExecuteSchedulesRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AutomationApi#addExecuteSchedules");
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
| **addExecuteSchedulesRequest** | [**AddExecuteSchedulesRequest**](AddExecuteSchedulesRequest.md)|  | [optional] |

### Return type

[**AddExecuteSchedules200Response**](AddExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="executeExecutionRequest"></a>
# **executeExecutionRequest**
> ExecuteExecutionRequest200Response executeExecutionRequest(instanceId, containerId, serverId, executeExecutionRequestRequest)

Executes an Execution Request

Provides API interfaces for executing an arbitrary script or command on an instance, container or host. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AutomationApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AutomationApi apiInstance = new AutomationApi(defaultClient);
    Long instanceId = 94L; // Long | The Instance ID for Filtering
    Long containerId = 135L; // Long | The Container ID for Filtering
    Long serverId = 97L; // Long | The Server ID for Filtering
    ExecuteExecutionRequestRequest executeExecutionRequestRequest = new ExecuteExecutionRequestRequest(); // ExecuteExecutionRequestRequest | 
    try {
      ExecuteExecutionRequest200Response result = apiInstance.executeExecutionRequest(instanceId, containerId, serverId, executeExecutionRequestRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AutomationApi#executeExecutionRequest");
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
| **instanceId** | **Long**| The Instance ID for Filtering | [optional] |
| **containerId** | **Long**| The Container ID for Filtering | [optional] |
| **serverId** | **Long**| The Server ID for Filtering | [optional] |
| **executeExecutionRequestRequest** | [**ExecuteExecutionRequestRequest**](ExecuteExecutionRequestRequest.md)|  | [optional] |

### Return type

[**ExecuteExecutionRequest200Response**](ExecuteExecutionRequest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Execution Request Response |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="getExecuteSchedules"></a>
# **getExecuteSchedules**
> GetExecuteSchedules200Response getExecuteSchedules(id)

Retrieves a Specific Execute Schedule

Retrieves a specific execute schedule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AutomationApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AutomationApi apiInstance = new AutomationApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      GetExecuteSchedules200Response result = apiInstance.getExecuteSchedules(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AutomationApi#getExecuteSchedules");
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
| **id** | **Long**| Morpheus ID of the Object being referenced | |

### Return type

[**GetExecuteSchedules200Response**](GetExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="getExecutionRequest"></a>
# **getExecutionRequest**
> GetExecutionRequest200Response getExecutionRequest(uniqueId)

Retrieves a Specific Execution Request

Retrieves a specific execution request. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AutomationApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AutomationApi apiInstance = new AutomationApi(defaultClient);
    String uniqueId = "c56f75d0-a59a-4566-92e3-4dc716c5d076"; // String | The Unique ID of the execution request
    try {
      GetExecutionRequest200Response result = apiInstance.getExecutionRequest(uniqueId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AutomationApi#getExecutionRequest");
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
| **uniqueId** | **String**| The Unique ID of the execution request | |

### Return type

[**GetExecutionRequest200Response**](GetExecutionRequest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="listExecuteSchedules"></a>
# **listExecuteSchedules**
> ListExecuteSchedules200Response listExecuteSchedules(max, offset, sort, direction, phrase, name)

Retrieves all Execute Schedules

Retrieves all execute schedules. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AutomationApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AutomationApi apiInstance = new AutomationApi(defaultClient);
    Long max = 25L; // Long | Maximum number of records to return
    Long offset = 0L; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "name"; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    try {
      ListExecuteSchedules200Response result = apiInstance.listExecuteSchedules(max, offset, sort, direction, phrase, name);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AutomationApi#listExecuteSchedules");
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
| **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc] |
| **phrase** | **String**| Search phrase for partial matches on name or description | [optional] |
| **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] |

### Return type

[**ListExecuteSchedules200Response**](ListExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="removeExecuteSchedules"></a>
# **removeExecuteSchedules**
> Model200Success removeExecuteSchedules(id)

Deletes a Execute Schedule

Deletes a specified execute schedule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AutomationApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AutomationApi apiInstance = new AutomationApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.removeExecuteSchedules(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AutomationApi#removeExecuteSchedules");
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
| **id** | **Long**| Morpheus ID of the Object being referenced | |

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
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="updateExecuteSchedules"></a>
# **updateExecuteSchedules**
> AddExecuteSchedules200Response updateExecuteSchedules(id, updateExecuteSchedulesRequest)

Updates a Execute Schedule

Updates a execute schedule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AutomationApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AutomationApi apiInstance = new AutomationApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    UpdateExecuteSchedulesRequest updateExecuteSchedulesRequest = new UpdateExecuteSchedulesRequest(); // UpdateExecuteSchedulesRequest | 
    try {
      AddExecuteSchedules200Response result = apiInstance.updateExecuteSchedules(id, updateExecuteSchedulesRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AutomationApi#updateExecuteSchedules");
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
| **id** | **Long**| Morpheus ID of the Object being referenced | |
| **updateExecuteSchedulesRequest** | [**UpdateExecuteSchedulesRequest**](UpdateExecuteSchedulesRequest.md)|  | [optional] |

### Return type

[**AddExecuteSchedules200Response**](AddExecuteSchedules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

