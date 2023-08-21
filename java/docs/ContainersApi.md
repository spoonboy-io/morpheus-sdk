# ContainersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**cloneImageContainerAction**](ContainersApi.md#cloneImageContainerAction) | **PUT** /api/containers/{id}/clone-image | Clone Specific Container to Image
[**containersAttachFloatingIp**](ContainersApi.md#containersAttachFloatingIp) | **PUT** /api/containers/{id}/attach-floating-ip | Attach Floating IP to Container
[**containersDetachFloatingIp**](ContainersApi.md#containersDetachFloatingIp) | **PUT** /api/containers/{id}/detach-floating-ip | Detach Floating IP from Container
[**ejectContainerAction**](ContainersApi.md#ejectContainerAction) | **PUT** /api/containers/{id}/eject | Eject a Specific Container
[**executeContainerAction**](ContainersApi.md#executeContainerAction) | **PUT** /api/containers/{id}/action | Execute Container Action
[**getContainer**](ContainersApi.md#getContainer) | **GET** /api/containers/{id} | Get a Specific Container
[**getContainerActions**](ContainersApi.md#getContainerActions) | **GET** /api/containers/{id}/actions | List Container Actions
[**importContainerAction**](ContainersApi.md#importContainerAction) | **PUT** /api/containers/{id}/import | Import a Specific Container
[**restartContainerAction**](ContainersApi.md#restartContainerAction) | **PUT** /api/containers/{id}/restart | Restart a Specific Container
[**startContainerAction**](ContainersApi.md#startContainerAction) | **PUT** /api/containers/{id}/start | Start a Specific Container
[**stopContainerAction**](ContainersApi.md#stopContainerAction) | **PUT** /api/containers/{id}/stop | Stop a Specific Container
[**suspendContainerAction**](ContainersApi.md#suspendContainerAction) | **PUT** /api/containers/{id}/suspend | Suspend a Specific Container


<a name="cloneImageContainerAction"></a>
# **cloneImageContainerAction**
> SuccessMessage cloneImageContainerAction(id, instancesCloneImage)

Clone Specific Container to Image

This endpoint clones and converts a Container in its current state to image in the source Cloud and adds Virtual Image Record with metadata matching the source configuration.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ContainersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ContainersApi apiInstance = new ContainersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InstancesCloneImage instancesCloneImage = {"templateName":"Example Image","zoneFolder":"group-v81"}; // InstancesCloneImage | 
    try {
      SuccessMessage result = apiInstance.cloneImageContainerAction(id, instancesCloneImage);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ContainersApi#cloneImageContainerAction");
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
 **instancesCloneImage** | [**InstancesCloneImage**](InstancesCloneImage.md)|  | [optional]

### Return type

[**SuccessMessage**](SuccessMessage.md)

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

<a name="containersAttachFloatingIp"></a>
# **containersAttachFloatingIp**
> Model200Success containersAttachFloatingIp(id, inlineObject63)

Attach Floating IP to Container

This endpoint attaches a floating IP to a container (node/VM).

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ContainersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ContainersApi apiInstance = new ContainersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject63 inlineObject63 = new InlineObject63(); // InlineObject63 | 
    try {
      Model200Success result = apiInstance.containersAttachFloatingIp(id, inlineObject63);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ContainersApi#containersAttachFloatingIp");
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
 **inlineObject63** | [**InlineObject63**](InlineObject63.md)|  | [optional]

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
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="containersDetachFloatingIp"></a>
# **containersDetachFloatingIp**
> Model200Success containersDetachFloatingIp(id)

Detach Floating IP from Container

This endpoint detaches a floating IP from a container (node/VM).

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ContainersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ContainersApi apiInstance = new ContainersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.containersDetachFloatingIp(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ContainersApi#containersDetachFloatingIp");
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

<a name="ejectContainerAction"></a>
# **ejectContainerAction**
> SuccessMessage ejectContainerAction(id)

Eject a Specific Container

This endpoint ejects attached disk/iso of a specific container.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ContainersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ContainersApi apiInstance = new ContainersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      SuccessMessage result = apiInstance.ejectContainerAction(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ContainersApi#ejectContainerAction");
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

[**SuccessMessage**](SuccessMessage.md)

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

<a name="executeContainerAction"></a>
# **executeContainerAction**
> SuccessMessage executeContainerAction(id, code)

Execute Container Action

This endpoint executes a container action on specific container.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ContainersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ContainersApi apiInstance = new ContainersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String code = "apache-remove-node"; // String | Action code to be executed on a specific container.
    try {
      SuccessMessage result = apiInstance.executeContainerAction(id, code);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ContainersApi#executeContainerAction");
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
 **code** | **String**| Action code to be executed on a specific container. |

### Return type

[**SuccessMessage**](SuccessMessage.md)

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

<a name="getContainer"></a>
# **getContainer**
> InlineResponse20034 getContainer(id)

Get a Specific Container

This endpoint retrieves a specific container.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ContainersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ContainersApi apiInstance = new ContainersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20034 result = apiInstance.getContainer(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ContainersApi#getContainer");
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

[**InlineResponse20034**](InlineResponse20034.md)

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

<a name="getContainerActions"></a>
# **getContainerActions**
> InlineResponse20035 getContainerActions(id)

List Container Actions

This endpoint lists available actions for specific container.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ContainersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ContainersApi apiInstance = new ContainersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20035 result = apiInstance.getContainerActions(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ContainersApi#getContainerActions");
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

[**InlineResponse20035**](InlineResponse20035.md)

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

<a name="importContainerAction"></a>
# **importContainerAction**
> SuccessMessage importContainerAction(id, inlineObject62)

Import a Specific Container

This endpoint clones and exports a Container in its current state to target Storage provider and adds Virtual Image Record with metadata matching the source configuration.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ContainersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ContainersApi apiInstance = new ContainersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject62 inlineObject62 = new InlineObject62(); // InlineObject62 | 
    try {
      SuccessMessage result = apiInstance.importContainerAction(id, inlineObject62);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ContainersApi#importContainerAction");
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
 **inlineObject62** | [**InlineObject62**](InlineObject62.md)|  | [optional]

### Return type

[**SuccessMessage**](SuccessMessage.md)

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

<a name="restartContainerAction"></a>
# **restartContainerAction**
> SuccessMessage restartContainerAction(id)

Restart a Specific Container

This endpoint restarts a specific container.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ContainersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ContainersApi apiInstance = new ContainersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      SuccessMessage result = apiInstance.restartContainerAction(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ContainersApi#restartContainerAction");
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

[**SuccessMessage**](SuccessMessage.md)

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

<a name="startContainerAction"></a>
# **startContainerAction**
> SuccessMessage startContainerAction(id)

Start a Specific Container

This endpoint starts a specific container.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ContainersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ContainersApi apiInstance = new ContainersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      SuccessMessage result = apiInstance.startContainerAction(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ContainersApi#startContainerAction");
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

[**SuccessMessage**](SuccessMessage.md)

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

<a name="stopContainerAction"></a>
# **stopContainerAction**
> SuccessMessage stopContainerAction(id)

Stop a Specific Container

This endpoint stops a specific container.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ContainersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ContainersApi apiInstance = new ContainersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      SuccessMessage result = apiInstance.stopContainerAction(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ContainersApi#stopContainerAction");
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

[**SuccessMessage**](SuccessMessage.md)

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

<a name="suspendContainerAction"></a>
# **suspendContainerAction**
> SuccessMessage suspendContainerAction(id)

Suspend a Specific Container

This endpoint suspends a specific container.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ContainersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ContainersApi apiInstance = new ContainersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      SuccessMessage result = apiInstance.suspendContainerAction(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ContainersApi#suspendContainerAction");
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

[**SuccessMessage**](SuccessMessage.md)

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

