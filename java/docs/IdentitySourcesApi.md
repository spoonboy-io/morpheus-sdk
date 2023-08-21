# IdentitySourcesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addIdentitySources**](IdentitySourcesApi.md#addIdentitySources) | **POST** /api/user-sources | Creates an Identity Source
[**getIdentitySources**](IdentitySourcesApi.md#getIdentitySources) | **GET** /api/user-sources/{id} | Retrieves a Specific Identity Source
[**listIdentitySources**](IdentitySourcesApi.md#listIdentitySources) | **GET** /api/user-sources | Retrieves all Identity Sources
[**removeIdentitySources**](IdentitySourcesApi.md#removeIdentitySources) | **DELETE** /api/user-sources/{id} | Deletes an Identity Source
[**updateIdentitySourceSubdomains**](IdentitySourcesApi.md#updateIdentitySourceSubdomains) | **PUT** /api/user-sources/{id}/subdomain | Updates an Identity Source Subdomain
[**updateIdentitySources**](IdentitySourcesApi.md#updateIdentitySources) | **PUT** /api/user-sources/{id} | Updates an Identity Source


<a name="addIdentitySources"></a>
# **addIdentitySources**
> Object addIdentitySources(accountId, userSourceCreate)

Creates an Identity Source

Creates an identity source. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IdentitySourcesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IdentitySourcesApi apiInstance = new IdentitySourcesApi(defaultClient);
    Long accountId = 3L; // Long | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
    UserSourceCreate userSourceCreate = new UserSourceCreate(); // UserSourceCreate | 
    try {
      Object result = apiInstance.addIdentitySources(accountId, userSourceCreate);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IdentitySourcesApi#addIdentitySources");
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
 **accountId** | **Long**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]
 **userSourceCreate** | [**UserSourceCreate**](UserSourceCreate.md)|  | [optional]

### Return type

**Object**

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

<a name="getIdentitySources"></a>
# **getIdentitySources**
> InlineResponse20051 getIdentitySources(id)

Retrieves a Specific Identity Source

Retrieves a specific identity source. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IdentitySourcesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IdentitySourcesApi apiInstance = new IdentitySourcesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20051 result = apiInstance.getIdentitySources(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IdentitySourcesApi#getIdentitySources");
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

[**InlineResponse20051**](InlineResponse20051.md)

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

<a name="listIdentitySources"></a>
# **listIdentitySources**
> Object listIdentitySources(type, max, offset, sort, direction, phrase, name, accountId)

Retrieves all Identity Sources

Retrieves all identity sources. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IdentitySourcesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IdentitySourcesApi apiInstance = new IdentitySourcesApi(defaultClient);
    String type = "type_example"; // String | If specified will return all tasks by `task type` code. Refer to `Task Types` API for up to date listings. 
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    Long accountId = 3L; // Long | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
    try {
      Object result = apiInstance.listIdentitySources(type, max, offset, sort, direction, phrase, name, accountId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IdentitySourcesApi#listIdentitySources");
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
 **type** | **String**| If specified will return all tasks by &#x60;task type&#x60; code. Refer to &#x60;Task Types&#x60; API for up to date listings.  | [optional] [enum: ansibleTask, ansibleTowerTask, chefTask, email, groovyTask, httpTask, javascriptTask, jrubyTask, containerScript, containerTemplate, localScript, nestedWorkflow, winrmTask, puppetTask, jythonTask, restartTask, script, sshTask, vro, writeAttributes]
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **accountId** | **Long**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]

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

<a name="removeIdentitySources"></a>
# **removeIdentitySources**
> Model200Success removeIdentitySources(id)

Deletes an Identity Source

Deletes a specified identity source. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IdentitySourcesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IdentitySourcesApi apiInstance = new IdentitySourcesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.removeIdentitySources(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IdentitySourcesApi#removeIdentitySources");
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

<a name="updateIdentitySourceSubdomains"></a>
# **updateIdentitySourceSubdomains**
> Object updateIdentitySourceSubdomains(id, inlineObject82)

Updates an Identity Source Subdomain

Updates an identity source subdomain. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IdentitySourcesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IdentitySourcesApi apiInstance = new IdentitySourcesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject82 inlineObject82 = new InlineObject82(); // InlineObject82 | 
    try {
      Object result = apiInstance.updateIdentitySourceSubdomains(id, inlineObject82);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IdentitySourcesApi#updateIdentitySourceSubdomains");
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
 **inlineObject82** | [**InlineObject82**](InlineObject82.md)|  | [optional]

### Return type

**Object**

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

<a name="updateIdentitySources"></a>
# **updateIdentitySources**
> Object updateIdentitySources(id, userSourceCreate)

Updates an Identity Source

Updates an identity source. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IdentitySourcesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IdentitySourcesApi apiInstance = new IdentitySourcesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    UserSourceCreate userSourceCreate = new UserSourceCreate(); // UserSourceCreate | 
    try {
      Object result = apiInstance.updateIdentitySources(id, userSourceCreate);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IdentitySourcesApi#updateIdentitySources");
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
 **userSourceCreate** | [**UserSourceCreate**](UserSourceCreate.md)|  | [optional]

### Return type

**Object**

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

