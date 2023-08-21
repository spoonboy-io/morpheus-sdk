# ResourcePoolsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createResourcePoolGroup**](ResourcePoolsApi.md#createResourcePoolGroup) | **POST** /api/resource-pools/groups | Create a Resource Pool Group
[**deleteResourcePoolGroup**](ResourcePoolsApi.md#deleteResourcePoolGroup) | **DELETE** /api/resource-pools/groups/{id} | Delete a Resource Pool Group
[**getResourcePoolGroups**](ResourcePoolsApi.md#getResourcePoolGroups) | **GET** /api/resource-pools/groups | Get all Resource Pool Groups
[**getresourcePoolGroup**](ResourcePoolsApi.md#getresourcePoolGroup) | **GET** /api/resource-pools/groups/{id} | Get a Specific Resource Pool Group
[**updateResourcePoolGroup**](ResourcePoolsApi.md#updateResourcePoolGroup) | **PUT** /api/resource-pools/groups/{id} | Update a Resource Pool Group


<a name="createResourcePoolGroup"></a>
# **createResourcePoolGroup**
> InlineResponse200132 createResourcePoolGroup(inlineObject206)

Create a Resource Pool Group

Use this command to create a resource pool group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ResourcePoolsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ResourcePoolsApi apiInstance = new ResourcePoolsApi(defaultClient);
    InlineObject206 inlineObject206 = new InlineObject206(); // InlineObject206 | 
    try {
      InlineResponse200132 result = apiInstance.createResourcePoolGroup(inlineObject206);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ResourcePoolsApi#createResourcePoolGroup");
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
 **inlineObject206** | [**InlineObject206**](InlineObject206.md)|  | [optional]

### Return type

[**InlineResponse200132**](InlineResponse200132.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteResourcePoolGroup"></a>
# **deleteResourcePoolGroup**
> Model200Success deleteResourcePoolGroup(id)

Delete a Resource Pool Group

Will delete a Resource Pool Group from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ResourcePoolsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ResourcePoolsApi apiInstance = new ResourcePoolsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteResourcePoolGroup(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ResourcePoolsApi#deleteResourcePoolGroup");
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
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getResourcePoolGroups"></a>
# **getResourcePoolGroups**
> InlineResponse200131 getResourcePoolGroups()

Get all Resource Pool Groups

This endpoint retrieves all Resource Pool Groups associated with the account. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ResourcePoolsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ResourcePoolsApi apiInstance = new ResourcePoolsApi(defaultClient);
    try {
      InlineResponse200131 result = apiInstance.getResourcePoolGroups();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ResourcePoolsApi#getResourcePoolGroups");
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

[**InlineResponse200131**](InlineResponse200131.md)

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

<a name="getresourcePoolGroup"></a>
# **getresourcePoolGroup**
> InlineResponse200132 getresourcePoolGroup(id)

Get a Specific Resource Pool Group

This endpoint retrieves a specific Resource Pool Group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ResourcePoolsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ResourcePoolsApi apiInstance = new ResourcePoolsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200132 result = apiInstance.getresourcePoolGroup(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ResourcePoolsApi#getresourcePoolGroup");
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

[**InlineResponse200132**](InlineResponse200132.md)

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

<a name="updateResourcePoolGroup"></a>
# **updateResourcePoolGroup**
> InlineResponse200132 updateResourcePoolGroup(id, inlineObject207)

Update a Resource Pool Group

Use this command to update an existing resource pool Group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ResourcePoolsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ResourcePoolsApi apiInstance = new ResourcePoolsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject207 inlineObject207 = new InlineObject207(); // InlineObject207 | 
    try {
      InlineResponse200132 result = apiInstance.updateResourcePoolGroup(id, inlineObject207);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ResourcePoolsApi#updateResourcePoolGroup");
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
 **inlineObject207** | [**InlineObject207**](InlineObject207.md)|  | [optional]

### Return type

[**InlineResponse200132**](InlineResponse200132.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

