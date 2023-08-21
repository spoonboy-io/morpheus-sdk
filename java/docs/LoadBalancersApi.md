# LoadBalancersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createLoadBalancer**](LoadBalancersApi.md#createLoadBalancer) | **POST** /api/load-balancers | Create a Load Balancer
[**createLoadBalancerMonitor**](LoadBalancersApi.md#createLoadBalancerMonitor) | **POST** /api/load-balancers/{loadBalancerId}/monitors | Create a Load Balancer Monitor
[**createLoadBalancerPool**](LoadBalancersApi.md#createLoadBalancerPool) | **POST** /api/load-balancers/{loadBalancerId}/pools | Create a Load Balancer Pool
[**createLoadBalancerPoolNode**](LoadBalancersApi.md#createLoadBalancerPoolNode) | **POST** /api/load-balancer-pools/{loadBalancerPoolId}/nodes | Create a Load Balancer Pool Node
[**createLoadBalancerProfile**](LoadBalancersApi.md#createLoadBalancerProfile) | **POST** /api/load-balancers/{loadBalancerId}/profiles | Create a Load Balancer Profile
[**createLoadBalancerVirtualServer**](LoadBalancersApi.md#createLoadBalancerVirtualServer) | **POST** /api/load-balancers/{loadBalancerId}/virtual-servers | Create a Load Balancer Virtual Server
[**deleteLoadBalancer**](LoadBalancersApi.md#deleteLoadBalancer) | **DELETE** /api/load-balancers/{loadBalancerId} | Delete a Load Balancer
[**deleteLoadBalancerMonitor**](LoadBalancersApi.md#deleteLoadBalancerMonitor) | **DELETE** /api/load-balancers/{loadBalancerId}/monitors/{id} | Delete a Load Balancer Monitor
[**deleteLoadBalancerPool**](LoadBalancersApi.md#deleteLoadBalancerPool) | **DELETE** /api/load-balancers/{loadBalancerId}/pools/{id} | Delete a Load Balancer Pool
[**deleteLoadBalancerPoolNode**](LoadBalancersApi.md#deleteLoadBalancerPoolNode) | **DELETE** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Delete a Load Balancer Pool Node
[**deleteLoadBalancerProfile**](LoadBalancersApi.md#deleteLoadBalancerProfile) | **DELETE** /api/load-balancers/{loadBalancerId}/profiles/{id} | Delete a Load Balancer Profile
[**deleteLoadBalancerVirtualServer**](LoadBalancersApi.md#deleteLoadBalancerVirtualServer) | **DELETE** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Delete a Load Balancer Virtual Server
[**getLoadBalancer**](LoadBalancersApi.md#getLoadBalancer) | **GET** /api/load-balancers/{loadBalancerId} | Get a Specific Load Balancer
[**getLoadBalancerMonitor**](LoadBalancersApi.md#getLoadBalancerMonitor) | **GET** /api/load-balancers/{loadBalancerId}/monitors/{id} | Get a Specific Load Balancer Monitor
[**getLoadBalancerPool**](LoadBalancersApi.md#getLoadBalancerPool) | **GET** /api/load-balancers/{loadBalancerId}/pools/{id} | Get a Specific Load Balancer Pool
[**getLoadBalancerPoolNode**](LoadBalancersApi.md#getLoadBalancerPoolNode) | **GET** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Get a Specific Load Balancer Pool Node
[**getLoadBalancerProfile**](LoadBalancersApi.md#getLoadBalancerProfile) | **GET** /api/load-balancers/{loadBalancerId}/profiles/{id} | Get a Specific Load Balancer Profile
[**getLoadBalancerType**](LoadBalancersApi.md#getLoadBalancerType) | **GET** /api/load-balancer-types/{id} | Get a Specific Load Balancer Type
[**getLoadBalancerVirtualServer**](LoadBalancersApi.md#getLoadBalancerVirtualServer) | **GET** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Get a Specific Load Balancer Virtual Server
[**listLoadBalancerMonitors**](LoadBalancersApi.md#listLoadBalancerMonitors) | **GET** /api/load-balancers/{loadBalancerId}/monitors | Get All Load Balancer Monitors For Load Balancer
[**listLoadBalancerPoolNodes**](LoadBalancersApi.md#listLoadBalancerPoolNodes) | **GET** /api/load-balancer-pools/{loadBalancerPoolId}/nodes | Get All Load Balancer Pool Nodes For Load Balancer Pool
[**listLoadBalancerPools**](LoadBalancersApi.md#listLoadBalancerPools) | **GET** /api/load-balancers/{loadBalancerId}/pools | Get All Load Balancer Pools For Load Balancer
[**listLoadBalancerProfiles**](LoadBalancersApi.md#listLoadBalancerProfiles) | **GET** /api/load-balancers/{loadBalancerId}/profiles | Get All Load Balancer Profiles For Load Balancer
[**listLoadBalancerTypes**](LoadBalancersApi.md#listLoadBalancerTypes) | **GET** /api/load-balancer-types | Get All Load Balancer Types
[**listLoadBalancerVirtualServers**](LoadBalancersApi.md#listLoadBalancerVirtualServers) | **GET** /api/load-balancers/{loadBalancerId}/virtual-servers | Get All Load Balancer Virtual Servers For Load Balancer
[**listLoadBalancers**](LoadBalancersApi.md#listLoadBalancers) | **GET** /api/load-balancers | Get All Load Balancers
[**refreshLoadBalancer**](LoadBalancersApi.md#refreshLoadBalancer) | **PUT** /api/load-balancers/{loadBalancerId}/refresh | Refresh a Load Balancer
[**updateLoadBalancer**](LoadBalancersApi.md#updateLoadBalancer) | **PUT** /api/load-balancers/{loadBalancerId} | Update a Load Balancer
[**updateLoadBalancerMonitor**](LoadBalancersApi.md#updateLoadBalancerMonitor) | **PUT** /api/load-balancers/{loadBalancerId}/monitors/{id} | Update a Load Balancer Monitor
[**updateLoadBalancerPool**](LoadBalancersApi.md#updateLoadBalancerPool) | **PUT** /api/load-balancers/{loadBalancerId}/pools/{id} | Update a Load Balancer Pool
[**updateLoadBalancerPoolNode**](LoadBalancersApi.md#updateLoadBalancerPoolNode) | **PUT** /api/load-balancer-pools/{loadBalancerPoolId}/nodes/{id} | Update a Load Balancer Pool Node
[**updateLoadBalancerProfile**](LoadBalancersApi.md#updateLoadBalancerProfile) | **PUT** /api/load-balancers/{loadBalancerId}/profiles/{id} | Update a Load Balancer Profile
[**updateLoadBalancerVirtualServer**](LoadBalancersApi.md#updateLoadBalancerVirtualServer) | **PUT** /api/load-balancers/{loadBalancerId}/virtual-servers/{id} | Update a Load Balancer Virtual Server


<a name="createLoadBalancer"></a>
# **createLoadBalancer**
> InlineResponse20078 createLoadBalancer(inlineObject127)

Create a Load Balancer

Available for NSX load balancers only  Use this command to create a load balancer. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    InlineObject127 inlineObject127 = new InlineObject127(); // InlineObject127 | 
    try {
      InlineResponse20078 result = apiInstance.createLoadBalancer(inlineObject127);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#createLoadBalancer");
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
 **inlineObject127** | [**InlineObject127**](InlineObject127.md)|  | [optional]

### Return type

[**InlineResponse20078**](InlineResponse20078.md)

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

<a name="createLoadBalancerMonitor"></a>
# **createLoadBalancerMonitor**
> Object createLoadBalancerMonitor(loadBalancerId, inlineObject129)

Create a Load Balancer Monitor

Use this command to create a load balancer Monitor.  This endpoint allows creating a Load Balancer Monitor. Configuration options vary by Load Balancer Type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    InlineObject129 inlineObject129 = new InlineObject129(); // InlineObject129 | 
    try {
      Object result = apiInstance.createLoadBalancerMonitor(loadBalancerId, inlineObject129);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#createLoadBalancerMonitor");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **inlineObject129** | [**InlineObject129**](InlineObject129.md)|  | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createLoadBalancerPool"></a>
# **createLoadBalancerPool**
> Object createLoadBalancerPool(loadBalancerId, inlineObject131)

Create a Load Balancer Pool

Use this command to create a load balancer pool.  This endpoint allows creating a Load Balancer Pool. Configuration options vary by Load Balancer Type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    InlineObject131 inlineObject131 = new InlineObject131(); // InlineObject131 | 
    try {
      Object result = apiInstance.createLoadBalancerPool(loadBalancerId, inlineObject131);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#createLoadBalancerPool");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **inlineObject131** | [**InlineObject131**](InlineObject131.md)|  | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createLoadBalancerPoolNode"></a>
# **createLoadBalancerPoolNode**
> Object createLoadBalancerPoolNode(loadBalancerPoolId, inlineObject137)

Create a Load Balancer Pool Node

Use this command to create a load balancer pool node.  This endpoint allows creating a Load Balancer Pool Node. Configuration options vary by Load Balancer Type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerPoolId = new BigDecimal(78); // BigDecimal | Load Balancer Pool ID
    InlineObject137 inlineObject137 = new InlineObject137(); // InlineObject137 | 
    try {
      Object result = apiInstance.createLoadBalancerPoolNode(loadBalancerPoolId, inlineObject137);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#createLoadBalancerPoolNode");
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
 **loadBalancerPoolId** | **BigDecimal**| Load Balancer Pool ID |
 **inlineObject137** | [**InlineObject137**](InlineObject137.md)|  | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createLoadBalancerProfile"></a>
# **createLoadBalancerProfile**
> Object createLoadBalancerProfile(loadBalancerId, inlineObject133)

Create a Load Balancer Profile

Use this command to create a load balancer Profile.  This endpoint allows creating a Load Balancer Profile. Configuration options vary by Load Balancer Type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    InlineObject133 inlineObject133 = new InlineObject133(); // InlineObject133 | 
    try {
      Object result = apiInstance.createLoadBalancerProfile(loadBalancerId, inlineObject133);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#createLoadBalancerProfile");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **inlineObject133** | [**InlineObject133**](InlineObject133.md)|  | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createLoadBalancerVirtualServer"></a>
# **createLoadBalancerVirtualServer**
> InlineResponse20082 createLoadBalancerVirtualServer(loadBalancerId, inlineObject135)

Create a Load Balancer Virtual Server

Use this command to create a load balancer virtual server.  This endpoint allows creating a Load Balancer Virtual Server. Configuration options vary by Load Balancer Type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    InlineObject135 inlineObject135 = new InlineObject135(); // InlineObject135 | 
    try {
      InlineResponse20082 result = apiInstance.createLoadBalancerVirtualServer(loadBalancerId, inlineObject135);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#createLoadBalancerVirtualServer");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **inlineObject135** | [**InlineObject135**](InlineObject135.md)|  | [optional]

### Return type

[**InlineResponse20082**](InlineResponse20082.md)

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

<a name="deleteLoadBalancer"></a>
# **deleteLoadBalancer**
> Model200Success deleteLoadBalancer(loadBalancerId)

Delete a Load Balancer

Will delete a Load Balancer from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    try {
      Model200Success result = apiInstance.deleteLoadBalancer(loadBalancerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#deleteLoadBalancer");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |

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

<a name="deleteLoadBalancerMonitor"></a>
# **deleteLoadBalancerMonitor**
> Model200Success deleteLoadBalancerMonitor(loadBalancerId, id)

Delete a Load Balancer Monitor

Will delete a Load Balancer Monitor from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteLoadBalancerMonitor(loadBalancerId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#deleteLoadBalancerMonitor");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
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

<a name="deleteLoadBalancerPool"></a>
# **deleteLoadBalancerPool**
> Model200Success deleteLoadBalancerPool(loadBalancerId, id)

Delete a Load Balancer Pool

Will delete a Load Balancer Pool from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteLoadBalancerPool(loadBalancerId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#deleteLoadBalancerPool");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
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

<a name="deleteLoadBalancerPoolNode"></a>
# **deleteLoadBalancerPoolNode**
> Model200Success deleteLoadBalancerPoolNode(loadBalancerPoolId, id)

Delete a Load Balancer Pool Node

Will delete a Load Balancer Pool Node from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerPoolId = new BigDecimal(78); // BigDecimal | Load Balancer Pool ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteLoadBalancerPoolNode(loadBalancerPoolId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#deleteLoadBalancerPoolNode");
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
 **loadBalancerPoolId** | **BigDecimal**| Load Balancer Pool ID |
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

<a name="deleteLoadBalancerProfile"></a>
# **deleteLoadBalancerProfile**
> Model200Success deleteLoadBalancerProfile(loadBalancerId, id)

Delete a Load Balancer Profile

Will delete a Load Balancer Profile from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteLoadBalancerProfile(loadBalancerId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#deleteLoadBalancerProfile");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
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

<a name="deleteLoadBalancerVirtualServer"></a>
# **deleteLoadBalancerVirtualServer**
> Model200Success deleteLoadBalancerVirtualServer(loadBalancerId, id)

Delete a Load Balancer Virtual Server

Will delete a Load Balancer Virtual Server from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteLoadBalancerVirtualServer(loadBalancerId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#deleteLoadBalancerVirtualServer");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
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

<a name="getLoadBalancer"></a>
# **getLoadBalancer**
> InlineResponse20078 getLoadBalancer(loadBalancerId)

Get a Specific Load Balancer

This endpoint retrieves a specific Load Balancer.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    try {
      InlineResponse20078 result = apiInstance.getLoadBalancer(loadBalancerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#getLoadBalancer");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |

### Return type

[**InlineResponse20078**](InlineResponse20078.md)

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

<a name="getLoadBalancerMonitor"></a>
# **getLoadBalancerMonitor**
> InlineResponse20079 getLoadBalancerMonitor(loadBalancerId, id)

Get a Specific Load Balancer Monitor

This endpoint retrieves a specific Load Balancer Monitor.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20079 result = apiInstance.getLoadBalancerMonitor(loadBalancerId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#getLoadBalancerMonitor");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20079**](InlineResponse20079.md)

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

<a name="getLoadBalancerPool"></a>
# **getLoadBalancerPool**
> InlineResponse20080 getLoadBalancerPool(loadBalancerId, id)

Get a Specific Load Balancer Pool

This endpoint retrieves a specific Load Balancer Pool.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20080 result = apiInstance.getLoadBalancerPool(loadBalancerId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#getLoadBalancerPool");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20080**](InlineResponse20080.md)

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

<a name="getLoadBalancerPoolNode"></a>
# **getLoadBalancerPoolNode**
> InlineResponse20083 getLoadBalancerPoolNode(loadBalancerPoolId, id)

Get a Specific Load Balancer Pool Node

This endpoint retrieves a specific Load Balancer Pool Node.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerPoolId = new BigDecimal(78); // BigDecimal | Load Balancer Pool ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20083 result = apiInstance.getLoadBalancerPoolNode(loadBalancerPoolId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#getLoadBalancerPoolNode");
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
 **loadBalancerPoolId** | **BigDecimal**| Load Balancer Pool ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20083**](InlineResponse20083.md)

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

<a name="getLoadBalancerProfile"></a>
# **getLoadBalancerProfile**
> InlineResponse20081 getLoadBalancerProfile(loadBalancerId, id)

Get a Specific Load Balancer Profile

This endpoint retrieves a specific Load Balancer Profile.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20081 result = apiInstance.getLoadBalancerProfile(loadBalancerId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#getLoadBalancerProfile");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20081**](InlineResponse20081.md)

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

<a name="getLoadBalancerType"></a>
# **getLoadBalancerType**
> InlineResponse20077 getLoadBalancerType(id)

Get a Specific Load Balancer Type

This endpoint will retrieve a specific load balancer type by id.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20077 result = apiInstance.getLoadBalancerType(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#getLoadBalancerType");
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

[**InlineResponse20077**](InlineResponse20077.md)

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

<a name="getLoadBalancerVirtualServer"></a>
# **getLoadBalancerVirtualServer**
> InlineResponse20082 getLoadBalancerVirtualServer(loadBalancerId, id)

Get a Specific Load Balancer Virtual Server

This endpoint retrieves a specific Load Balancer Virtual Server.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20082 result = apiInstance.getLoadBalancerVirtualServer(loadBalancerId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#getLoadBalancerVirtualServer");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20082**](InlineResponse20082.md)

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

<a name="listLoadBalancerMonitors"></a>
# **listLoadBalancerMonitors**
> Object listLoadBalancerMonitors(loadBalancerId, max, offset, sort, direction, name, phrase)

Get All Load Balancer Monitors For Load Balancer

This endpoint retrieves all load balancer monitors associated with a specified load balancer.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.listLoadBalancerMonitors(loadBalancerId, max, offset, sort, direction, name, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#listLoadBalancerMonitors");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

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

<a name="listLoadBalancerPoolNodes"></a>
# **listLoadBalancerPoolNodes**
> Object listLoadBalancerPoolNodes(loadBalancerPoolId, max, offset, sort, direction, phrase)

Get All Load Balancer Pool Nodes For Load Balancer Pool

This endpoint retrieves all load balancer pool nodes associated with a specified load balancer pool.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerPoolId = new BigDecimal(78); // BigDecimal | Load Balancer Pool ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.listLoadBalancerPoolNodes(loadBalancerPoolId, max, offset, sort, direction, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#listLoadBalancerPoolNodes");
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
 **loadBalancerPoolId** | **BigDecimal**| Load Balancer Pool ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

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

<a name="listLoadBalancerPools"></a>
# **listLoadBalancerPools**
> Object listLoadBalancerPools(loadBalancerId, max, offset, sort, direction, name, phrase)

Get All Load Balancer Pools For Load Balancer

This endpoint retrieves all load balancer pools associated with a specified load balancer.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.listLoadBalancerPools(loadBalancerId, max, offset, sort, direction, name, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#listLoadBalancerPools");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

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

<a name="listLoadBalancerProfiles"></a>
# **listLoadBalancerProfiles**
> Object listLoadBalancerProfiles(loadBalancerId, max, offset, sort, direction, name, phrase)

Get All Load Balancer Profiles For Load Balancer

This endpoint retrieves all load balancer profiles associated with a specified load balancer.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.listLoadBalancerProfiles(loadBalancerId, max, offset, sort, direction, name, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#listLoadBalancerProfiles");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

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

<a name="listLoadBalancerTypes"></a>
# **listLoadBalancerTypes**
> Object listLoadBalancerTypes(max, offset, sort, direction, optionTypes, phrase, name, code)

Get All Load Balancer Types

This endpoint retrieves all Load Balancer Types.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    Boolean optionTypes = false; // Boolean | Pass true to include optionTypes in the response for each entry.
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String code = "azr"; // String | If specified will return an exact match on code
    try {
      Object result = apiInstance.listLoadBalancerTypes(max, offset, sort, direction, optionTypes, phrase, name, code);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#listLoadBalancerTypes");
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
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **optionTypes** | **Boolean**| Pass true to include optionTypes in the response for each entry. | [optional] [default to false]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **String**| If specified will return an exact match on code | [optional]

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

<a name="listLoadBalancerVirtualServers"></a>
# **listLoadBalancerVirtualServers**
> Object listLoadBalancerVirtualServers(loadBalancerId, max, offset, sort, direction, phrase, vipName, vipAddress, vipHostname)

Get All Load Balancer Virtual Servers For Load Balancer

This endpoint retrieves load balancer virtual servers associated with a specified load balancer.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String vipName = "lb-114"; // String | If specified will return an exact match on vipName
    String vipAddress = "192.168.123.44"; // String | If specified will return an exact match on vipAddress
    String vipHostname = "mylb"; // String | If specified will return an exact match on vipHostname
    try {
      Object result = apiInstance.listLoadBalancerVirtualServers(loadBalancerId, max, offset, sort, direction, phrase, vipName, vipAddress, vipHostname);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#listLoadBalancerVirtualServers");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **vipName** | **String**| If specified will return an exact match on vipName | [optional]
 **vipAddress** | **String**| If specified will return an exact match on vipAddress | [optional]
 **vipHostname** | **String**| If specified will return an exact match on vipHostname | [optional]

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

<a name="listLoadBalancers"></a>
# **listLoadBalancers**
> Object listLoadBalancers(max, offset, sort, direction, name, phrase)

Get All Load Balancers

This endpoint retrieves all load balancers associated with the account.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.listLoadBalancers(max, offset, sort, direction, name, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#listLoadBalancers");
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
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

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

<a name="refreshLoadBalancer"></a>
# **refreshLoadBalancer**
> Object refreshLoadBalancer(loadBalancerId)

Refresh a Load Balancer

Will refresh a Load Balancer.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    try {
      Object result = apiInstance.refreshLoadBalancer(loadBalancerId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#refreshLoadBalancer");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |

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

<a name="updateLoadBalancer"></a>
# **updateLoadBalancer**
> InlineResponse20078 updateLoadBalancer(loadBalancerId, inlineObject128)

Update a Load Balancer

Available for NSX load balancers only  Use this command to update an existing load balancer. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    InlineObject128 inlineObject128 = new InlineObject128(); // InlineObject128 | 
    try {
      InlineResponse20078 result = apiInstance.updateLoadBalancer(loadBalancerId, inlineObject128);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#updateLoadBalancer");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **inlineObject128** | [**InlineObject128**](InlineObject128.md)|  | [optional]

### Return type

[**InlineResponse20078**](InlineResponse20078.md)

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

<a name="updateLoadBalancerMonitor"></a>
# **updateLoadBalancerMonitor**
> Object updateLoadBalancerMonitor(loadBalancerId, id, inlineObject130)

Update a Load Balancer Monitor

Use this command to update an existing load balancer monitor.  This endpoint allows updating a Load Balancer Monitor. Configuration options vary by Load Balancer Type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject130 inlineObject130 = new InlineObject130(); // InlineObject130 | 
    try {
      Object result = apiInstance.updateLoadBalancerMonitor(loadBalancerId, id, inlineObject130);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#updateLoadBalancerMonitor");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject130** | [**InlineObject130**](InlineObject130.md)|  | [optional]

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

<a name="updateLoadBalancerPool"></a>
# **updateLoadBalancerPool**
> Object updateLoadBalancerPool(loadBalancerId, id, inlineObject132)

Update a Load Balancer Pool

Use this command to update an existing load balancer pool.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject132 inlineObject132 = new InlineObject132(); // InlineObject132 | 
    try {
      Object result = apiInstance.updateLoadBalancerPool(loadBalancerId, id, inlineObject132);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#updateLoadBalancerPool");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject132** | [**InlineObject132**](InlineObject132.md)|  | [optional]

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

<a name="updateLoadBalancerPoolNode"></a>
# **updateLoadBalancerPoolNode**
> Object updateLoadBalancerPoolNode(loadBalancerPoolId, id, inlineObject138)

Update a Load Balancer Pool Node

Use this command to update an existing load balancer pool node.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerPoolId = new BigDecimal(78); // BigDecimal | Load Balancer Pool ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject138 inlineObject138 = new InlineObject138(); // InlineObject138 | 
    try {
      Object result = apiInstance.updateLoadBalancerPoolNode(loadBalancerPoolId, id, inlineObject138);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#updateLoadBalancerPoolNode");
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
 **loadBalancerPoolId** | **BigDecimal**| Load Balancer Pool ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject138** | [**InlineObject138**](InlineObject138.md)|  | [optional]

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

<a name="updateLoadBalancerProfile"></a>
# **updateLoadBalancerProfile**
> Object updateLoadBalancerProfile(loadBalancerId, id, inlineObject134)

Update a Load Balancer Profile

Use this command to update an existing load balancer Profile.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject134 inlineObject134 = new InlineObject134(); // InlineObject134 | 
    try {
      Object result = apiInstance.updateLoadBalancerProfile(loadBalancerId, id, inlineObject134);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#updateLoadBalancerProfile");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject134** | [**InlineObject134**](InlineObject134.md)|  | [optional]

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

<a name="updateLoadBalancerVirtualServer"></a>
# **updateLoadBalancerVirtualServer**
> InlineResponse20082 updateLoadBalancerVirtualServer(loadBalancerId, id, inlineObject136)

Update a Load Balancer Virtual Server

Use this command to update an existing load balancer virtual server.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LoadBalancersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LoadBalancersApi apiInstance = new LoadBalancersApi(defaultClient);
    BigDecimal loadBalancerId = new BigDecimal(78); // BigDecimal | Load Balancer ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject136 inlineObject136 = new InlineObject136(); // InlineObject136 | 
    try {
      InlineResponse20082 result = apiInstance.updateLoadBalancerVirtualServer(loadBalancerId, id, inlineObject136);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LoadBalancersApi#updateLoadBalancerVirtualServer");
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
 **loadBalancerId** | **BigDecimal**| Load Balancer ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject136** | [**InlineObject136**](InlineObject136.md)|  | [optional]

### Return type

[**InlineResponse20082**](InlineResponse20082.md)

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

