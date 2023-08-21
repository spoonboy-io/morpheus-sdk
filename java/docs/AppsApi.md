# AppsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addAppInstance**](AppsApi.md#addAppInstance) | **POST** /api/apps/{id}/add-instance | Add Existing Instance to App
[**addAppUndoDelete**](AppsApi.md#addAppUndoDelete) | **PUT** /api/apps/{id}/cancel-removal | Undo Delete of an App
[**addApps**](AppsApi.md#addApps) | **POST** /api/apps | Create an App
[**applyAppState**](AppsApi.md#applyAppState) | **POST** /api/apps/{id}/apply | Apply State of an App
[**deleteApp**](AppsApi.md#deleteApp) | **DELETE** /api/apps/{id} | Delete an App
[**getApp**](AppsApi.md#getApp) | **GET** /api/apps/{id} | Get a Specific App
[**getAppSecurityGroups**](AppsApi.md#getAppSecurityGroups) | **GET** /api/apps/{id}/security-groups | Get Security Groups for an App
[**getAppState**](AppsApi.md#getAppState) | **GET** /api/apps/{id}/state | Get State of an App
[**getWikiApp**](AppsApi.md#getWikiApp) | **GET** /api/apps/{id}/wiki | Retrieves an App Wiki Page
[**listApps**](AppsApi.md#listApps) | **GET** /api/apps | Get All Apps
[**prepareAppApply**](AppsApi.md#prepareAppApply) | **GET** /api/apps/{id}/prepare-apply | Prepare To Apply an App
[**refreshAppState**](AppsApi.md#refreshAppState) | **POST** /api/apps/{id}/refresh | Refresh State of an App
[**removeAppInstance**](AppsApi.md#removeAppInstance) | **POST** /api/apps/{id}/remove-instance | Remove Instance from App
[**setAppSecurityGroups**](AppsApi.md#setAppSecurityGroups) | **POST** /api/apps/{id}/security-groups | Set Security Groups for an App
[**updateApp**](AppsApi.md#updateApp) | **PUT** /api/apps/{id} | Updating an App
[**updateWikiApp**](AppsApi.md#updateWikiApp) | **PUT** /api/apps/{id}/wiki | Update an App Wiki Page
[**validateAppState**](AppsApi.md#validateAppState) | **POST** /api/apps/{id}/validate-apply | Validate Apply State for an App


<a name="addAppInstance"></a>
# **addAppInstance**
> InlineResponse2003 addAppInstance(id, inlineObject3)

Add Existing Instance to App

Add Existing Instance to App

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject3 inlineObject3 = new InlineObject3(); // InlineObject3 | 
    try {
      InlineResponse2003 result = apiInstance.addAppInstance(id, inlineObject3);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#addAppInstance");
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
 **inlineObject3** | [**InlineObject3**](InlineObject3.md)|  | [optional]

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

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

<a name="addAppUndoDelete"></a>
# **addAppUndoDelete**
> InlineResponse2003 addAppUndoDelete(id)

Undo Delete of an App

This operation will undo the delete of an app that is pending removal.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse2003 result = apiInstance.addAppUndoDelete(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#addAppUndoDelete");
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

[**InlineResponse2003**](InlineResponse2003.md)

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

<a name="addApps"></a>
# **addApps**
> InlineResponse2002 addApps(appCreate)

Create an App

Create an App

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    AppCreate appCreate = {"blueprintId":"existing","name":"sample","description":"A sample app","group":{"id":1},"defaultCloud":{"id":19}}; // AppCreate | 
    try {
      InlineResponse2002 result = apiInstance.addApps(appCreate);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#addApps");
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
 **appCreate** | [**AppCreate**](AppCreate.md)|  | [optional]

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

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

<a name="applyAppState"></a>
# **applyAppState**
> Model200Success applyAppState(id, inlineObject4)

Apply State of an App

This endpoint provides a way to apply the state of an app. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject4 inlineObject4 = new InlineObject4(); // InlineObject4 | 
    try {
      Model200Success result = apiInstance.applyAppState(id, inlineObject4);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#applyAppState");
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
 **inlineObject4** | [**InlineObject4**](InlineObject4.md)|  | [optional]

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

<a name="deleteApp"></a>
# **deleteApp**
> Model200Success deleteApp(id, removeInstances, preserveVolumes, keepBackups, releaseFloatingIps, releaseEIPs, force)

Delete an App

Will delete an app. Use removeInstances&#x3D;on to also delete the instances in the app and all associated monitors and backups.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String removeInstances = "\"off\""; // String | Remove Instances
    String preserveVolumes = "\"off\""; // String | Preserve Volumes
    String keepBackups = "\"off\""; // String | Preserve copy of backups
    String releaseFloatingIps = "\"on\""; // String | Release Floating IPs
    String releaseEIPs = "\"on\""; // String | Alias for releaseFloatingIps
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.deleteApp(id, removeInstances, preserveVolumes, keepBackups, releaseFloatingIps, releaseEIPs, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#deleteApp");
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
 **removeInstances** | **String**| Remove Instances | [optional] [default to &quot;off&quot;]
 **preserveVolumes** | **String**| Preserve Volumes | [optional] [default to &quot;off&quot;]
 **keepBackups** | **String**| Preserve copy of backups | [optional] [default to &quot;off&quot;]
 **releaseFloatingIps** | **String**| Release Floating IPs | [optional] [default to &quot;on&quot;]
 **releaseEIPs** | **String**| Alias for releaseFloatingIps | [optional] [default to &quot;on&quot;]
 **force** | **String**| Force Delete | [optional] [default to &quot;off&quot;]

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

<a name="getApp"></a>
# **getApp**
> InlineResponse2003 getApp(id)

Get a Specific App

This endpoint retrieves a specific app.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse2003 result = apiInstance.getApp(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#getApp");
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

[**InlineResponse2003**](InlineResponse2003.md)

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

<a name="getAppSecurityGroups"></a>
# **getAppSecurityGroups**
> Object getAppSecurityGroups(id)

Get Security Groups for an App

This returns a list of all of the security groups applied to an app and whether the firewall is enabled.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getAppSecurityGroups(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#getAppSecurityGroups");
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

<a name="getAppState"></a>
# **getAppState**
> AppState getAppState(id)

Get State of an App

This endpoint provides a way to view the state of an app. The response includes output and resource planning information from the template provider software such as Terraform. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      AppState result = apiInstance.getAppState(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#getAppState");
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

[**AppState**](AppState.md)

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

<a name="getWikiApp"></a>
# **getWikiApp**
> InlineResponse200168 getWikiApp(id)

Retrieves an App Wiki Page

This endpoint retrieves an app Wiki page. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200168 result = apiInstance.getWikiApp(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#getWikiApp");
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

[**InlineResponse200168**](InlineResponse200168.md)

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

<a name="listApps"></a>
# **listApps**
> Object listApps(max, offset, name, phrase, createdBy, showDeleted, labels, allLabels)

Get All Apps

This endpoint retrieves a paginated list of apps. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    Long createdBy = 63L; // Long | The User ID for Filtering
    Boolean showDeleted = false; // Boolean | If true, includes instances in pending removal status.
    String labels = "labels_example"; // String | Filter by label(s), matches records that contain any of the specified labels
    String allLabels = "allLabels_example"; // String | Filter by label(s), matches records that contain all of the specified labels
    try {
      Object result = apiInstance.listApps(max, offset, name, phrase, createdBy, showDeleted, labels, allLabels);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#listApps");
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **createdBy** | **Long**| The User ID for Filtering | [optional]
 **showDeleted** | **Boolean**| If true, includes instances in pending removal status. | [optional] [default to false]
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional]

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

<a name="prepareAppApply"></a>
# **prepareAppApply**
> AppPrepareApply prepareAppApply(id)

Prepare To Apply an App

This endpoint provides a way to view the current app configuration and templateParameter variables available to apply. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      AppPrepareApply result = apiInstance.prepareAppApply(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#prepareAppApply");
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

[**AppPrepareApply**](AppPrepareApply.md)

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

<a name="refreshAppState"></a>
# **refreshAppState**
> Model200Success refreshAppState(id, body)

Refresh State of an App

This endpoint provides a way to refresh the state of an app. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types.   

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Object body = null; // Object | 
    try {
      Model200Success result = apiInstance.refreshAppState(id, body);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#refreshAppState");
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
 **body** | **Object**|  | [optional]

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

<a name="removeAppInstance"></a>
# **removeAppInstance**
> InlineResponse2003 removeAppInstance(id, inlineObject5)

Remove Instance from App

Remove Instance from App

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject5 inlineObject5 = new InlineObject5(); // InlineObject5 | 
    try {
      InlineResponse2003 result = apiInstance.removeAppInstance(id, inlineObject5);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#removeAppInstance");
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
 **inlineObject5** | [**InlineObject5**](InlineObject5.md)|  | [optional]

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

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

<a name="setAppSecurityGroups"></a>
# **setAppSecurityGroups**
> Object setAppSecurityGroups(id, UNKNOWN_BASE_TYPE)

Set Security Groups for an App

Set Security Groups for an App

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE = {"securityGroupIds":[19,2]}; // UNKNOWN_BASE_TYPE | 
    try {
      Object result = apiInstance.setAppSecurityGroups(id, UNKNOWN_BASE_TYPE);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#setAppSecurityGroups");
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional]

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

<a name="updateApp"></a>
# **updateApp**
> InlineResponse2003 updateApp(id, appUpdate)

Updating an App

This endpoint provides updating of some basic app settings.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    AppUpdate appUpdate = {"name":"My Sample App","description":"A new description of this app"}; // AppUpdate | 
    try {
      InlineResponse2003 result = apiInstance.updateApp(id, appUpdate);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#updateApp");
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
 **appUpdate** | [**AppUpdate**](AppUpdate.md)|  | [optional]

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

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

<a name="updateWikiApp"></a>
# **updateWikiApp**
> Object updateWikiApp(id, inlineObject267)

Update an App Wiki Page

Updates an app Wiki page. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject267 inlineObject267 = new InlineObject267(); // InlineObject267 | 
    try {
      Object result = apiInstance.updateWikiApp(id, inlineObject267);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#updateWikiApp");
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
 **inlineObject267** | [**InlineObject267**](InlineObject267.md)|  | [optional]

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

<a name="validateAppState"></a>
# **validateAppState**
> Object validateAppState(id, inlineObject6)

Validate Apply State for an App

This endpoint provides a way to validate app configuration and templateParameter variables before executing the apply. This only validates the configuration to see any planned changes and it does not actually apply the changes. This action only applies to Terraform, CloudFormation and ARM and will return an HTTP 400 error for other types. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.AppsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    AppsApi apiInstance = new AppsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject6 inlineObject6 = new InlineObject6(); // InlineObject6 | 
    try {
      Object result = apiInstance.validateAppState(id, inlineObject6);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling AppsApi#validateAppState");
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
 **inlineObject6** | [**InlineObject6**](InlineObject6.md)|  | [optional]

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

