# ChecksApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCheckApps**](ChecksApi.md#addCheckApps) | **POST** /api/monitoring/apps | Create a New Check App
[**addCheckGroups**](ChecksApi.md#addCheckGroups) | **POST** /api/monitoring/groups | Create a New Check Group
[**addChecks**](ChecksApi.md#addChecks) | **POST** /api/monitoring/checks | Create a New Check
[**deleteCheckApps**](ChecksApi.md#deleteCheckApps) | **DELETE** /api/monitoring/apps/{id} | Delete a Specific Check App
[**deleteCheckGroups**](ChecksApi.md#deleteCheckGroups) | **DELETE** /api/monitoring/groups/{id} | Delete a Specific Check Group
[**deleteChecks**](ChecksApi.md#deleteChecks) | **DELETE** /api/monitoring/checks/{id} | Delete a Specific Check
[**getCheckApps**](ChecksApi.md#getCheckApps) | **GET** /api/monitoring/apps/{id} | Get a Specific Check App
[**getCheckGroups**](ChecksApi.md#getCheckGroups) | **GET** /api/monitoring/groups/{id} | Get a Specific Check Group
[**getCheckTypes**](ChecksApi.md#getCheckTypes) | **GET** /api/monitoring/check-types/{id} | Get a Specific Check Type
[**getChecks**](ChecksApi.md#getChecks) | **GET** /api/monitoring/checks/{id} | Get a Specific Check
[**listCheckApps**](ChecksApi.md#listCheckApps) | **GET** /api/monitoring/apps | List All Check Apps
[**listCheckGroups**](ChecksApi.md#listCheckGroups) | **GET** /api/monitoring/groups | List All Check Groups
[**listCheckTypes**](ChecksApi.md#listCheckTypes) | **GET** /api/monitoring/check-types | List All Check Types
[**listChecks**](ChecksApi.md#listChecks) | **GET** /api/monitoring/checks | List All Checks
[**updateCheckApps**](ChecksApi.md#updateCheckApps) | **PUT** /api/monitoring/apps/{id} | Update Check App
[**updateCheckGroups**](ChecksApi.md#updateCheckGroups) | **PUT** /api/monitoring/groups/{id} | Update Check Group
[**updateChecks**](ChecksApi.md#updateChecks) | **PUT** /api/monitoring/checks/{id} | Updates a Check
[**updateMuteAllCheckApps**](ChecksApi.md#updateMuteAllCheckApps) | **PUT** /api/monitoring/apps/mute-all | Mute All Check Apps
[**updateMuteAllCheckGroups**](ChecksApi.md#updateMuteAllCheckGroups) | **PUT** /api/monitoring/groups/mute-all | Mute All Check Groups
[**updateMuteAllChecks**](ChecksApi.md#updateMuteAllChecks) | **PUT** /api/monitoring/checks/mute-all | Mute All Checks
[**updateMuteCheckApps**](ChecksApi.md#updateMuteCheckApps) | **PUT** /api/monitoring/apps/{id}/mute | Mute Check App
[**updateMuteCheckGroups**](ChecksApi.md#updateMuteCheckGroups) | **PUT** /api/monitoring/groups/{id}/mute | Mute Check Group
[**updateMuteChecks**](ChecksApi.md#updateMuteChecks) | **PUT** /api/monitoring/checks/{id}/mute | Mute Check


<a name="addCheckApps"></a>
# **addCheckApps**
> Object addCheckApps(inlineObject29)

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
    InlineObject29 inlineObject29 = new InlineObject29(); // InlineObject29 | 
    try {
      Object result = apiInstance.addCheckApps(inlineObject29);
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject29** | [**InlineObject29**](InlineObject29.md)|  | [optional]

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
**200** | Check App Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="addCheckGroups"></a>
# **addCheckGroups**
> Object addCheckGroups(inlineObject37)

Create a New Check Group

Create a new check group.

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
    InlineObject37 inlineObject37 = new InlineObject37(); // InlineObject37 | 
    try {
      Object result = apiInstance.addCheckGroups(inlineObject37);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#addCheckGroups");
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
 **inlineObject37** | [**InlineObject37**](InlineObject37.md)|  | [optional]

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
**200** | Check Group Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="addChecks"></a>
# **addChecks**
> Object addChecks(inlineObject33)

Create a New Check

Create a new monitoring check.  SSH tunneling options allow the different check types to tunnel to a host via a proxy, and execute checks relative to the proxy. A SSH tunnel can use your account generated public and private key-pairs or SSH password. It is strongly recommended to use a key-pair.  To enable SSH tunneling for a check, add &#x60;tunnelOn&#x60;, &#x60;sshHost&#x60;, &#x60;sshUser&#x60;, and optionally, &#x60;sshPort&#x60; and &#x60;sshPassword&#x60; parameters to any check type config as seen earlier in the Check Types section. 

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
    InlineObject33 inlineObject33 = new InlineObject33(); // InlineObject33 | 
    try {
      Object result = apiInstance.addChecks(inlineObject33);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#addChecks");
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
 **inlineObject33** | [**InlineObject33**](InlineObject33.md)|  | [optional]

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
**200** | Check Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteCheckApps"></a>
# **deleteCheckApps**
> Model200Success deleteCheckApps(id)

Delete a Specific Check App

Delete an existing monitoring check app.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteCheckApps(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#deleteCheckApps");
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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteCheckGroups"></a>
# **deleteCheckGroups**
> Model200Success deleteCheckGroups(id)

Delete a Specific Check Group

Delete an existing monitoring check group.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteCheckGroups(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#deleteCheckGroups");
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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteChecks"></a>
# **deleteChecks**
> Model200Success deleteChecks(id)

Delete a Specific Check

Delete an existing monitoring check.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteChecks(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#deleteChecks");
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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getCheckApps"></a>
# **getCheckApps**
> InlineResponse20016 getCheckApps(id)

Get a Specific Check App

Get details about a specific monitoring check app.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20016 result = apiInstance.getCheckApps(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#getCheckApps");
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

[**InlineResponse20016**](InlineResponse20016.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check App Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getCheckGroups"></a>
# **getCheckGroups**
> InlineResponse20019 getCheckGroups(id)

Get a Specific Check Group

Get details about a specific monitoring check group.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20019 result = apiInstance.getCheckGroups(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#getCheckGroups");
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

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check Group Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getCheckTypes"></a>
# **getCheckTypes**
> InlineResponse20018 getCheckTypes(id)

Get a Specific Check Type

Get details about a specific monitoring check type.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20018 result = apiInstance.getCheckTypes(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#getCheckTypes");
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

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Check Types Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getChecks"></a>
# **getChecks**
> InlineResponse20017 getChecks(id)

Get a Specific Check

Get details about a specific monitoring check.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20017 result = apiInstance.getChecks(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#getChecks");
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

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Checks Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listCheckApps"></a>
# **listCheckApps**
> Object listCheckApps(max, offset, sort, name, phrase, status, lastUpdated, deleted)

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
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String status = "running"; // String | The instance status for filtering.
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    Boolean deleted = true; // Boolean | If true, only deleted resources or instances in pending removal status are returned.
    try {
      Object result = apiInstance.listCheckApps(max, offset, sort, name, phrase, status, lastUpdated, deleted);
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **status** | **String**| The instance status for filtering. | [optional]
 **lastUpdated** | **OffsetDateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deleted** | **Boolean**| If true, only deleted resources or instances in pending removal status are returned. | [optional]

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
**200** | Array of Check Apps |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listCheckGroups"></a>
# **listCheckGroups**
> Object listCheckGroups(max, offset, sort, name, phrase, status, lastUpdated, deleted)

List All Check Groups

Get a list of check groups.

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
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String status = "running"; // String | The instance status for filtering.
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    Boolean deleted = true; // Boolean | If true, only deleted resources or instances in pending removal status are returned.
    try {
      Object result = apiInstance.listCheckGroups(max, offset, sort, name, phrase, status, lastUpdated, deleted);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#listCheckGroups");
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **status** | **String**| The instance status for filtering. | [optional]
 **lastUpdated** | **OffsetDateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deleted** | **Boolean**| If true, only deleted resources or instances in pending removal status are returned. | [optional]

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
**200** | Array of Check Groups |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listCheckTypes"></a>
# **listCheckTypes**
> Object listCheckTypes(max, offset, sort, name, phrase)

List All Check Types

Get a list of monitoring check types.

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
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      Object result = apiInstance.listCheckTypes(max, offset, sort, name, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#listCheckTypes");
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
**200** | Array of Check Types |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listChecks"></a>
# **listChecks**
> Object listChecks(max, offset, sort, name, phrase, status, lastUpdated, deleted)

List All Checks

Get a list of monitoring checks.

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
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String status = "healthy"; // String | Filter by health status.
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    Boolean deleted = true; // Boolean | If true, only deleted resources or instances in pending removal status are returned.
    try {
      Object result = apiInstance.listChecks(max, offset, sort, name, phrase, status, lastUpdated, deleted);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#listChecks");
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **status** | **String**| Filter by health status. | [optional] [enum: error, healthy, warning, muted]
 **lastUpdated** | **OffsetDateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deleted** | **Boolean**| If true, only deleted resources or instances in pending removal status are returned. | [optional]

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
**200** | Array of Checks |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateCheckApps"></a>
# **updateCheckApps**
> Object updateCheckApps(id, inlineObject31)

Update Check App

Update an existing monitoring check app.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject31 inlineObject31 = new InlineObject31(); // InlineObject31 | 
    try {
      Object result = apiInstance.updateCheckApps(id, inlineObject31);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#updateCheckApps");
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
 **inlineObject31** | [**InlineObject31**](InlineObject31.md)|  | [optional]

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
**200** | Check App Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateCheckGroups"></a>
# **updateCheckGroups**
> Object updateCheckGroups(id, inlineObject38)

Update Check Group

Update an existing monitoring check group.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject38 inlineObject38 = new InlineObject38(); // InlineObject38 | 
    try {
      Object result = apiInstance.updateCheckGroups(id, inlineObject38);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#updateCheckGroups");
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
 **inlineObject38** | [**InlineObject38**](InlineObject38.md)|  | [optional]

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
**200** | Check Group Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateChecks"></a>
# **updateChecks**
> Object updateChecks(id, inlineObject35)

Updates a Check

Updates a monitoring check.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject35 inlineObject35 = new InlineObject35(); // InlineObject35 | 
    try {
      Object result = apiInstance.updateChecks(id, inlineObject35);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#updateChecks");
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
 **inlineObject35** | [**InlineObject35**](InlineObject35.md)|  | [optional]

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
**200** | Check Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateMuteAllCheckApps"></a>
# **updateMuteAllCheckApps**
> Object updateMuteAllCheckApps(inlineObject30)

Mute All Check Apps

Mute all existing check apps.

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
    InlineObject30 inlineObject30 = new InlineObject30(); // InlineObject30 | 
    try {
      Object result = apiInstance.updateMuteAllCheckApps(inlineObject30);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#updateMuteAllCheckApps");
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
 **inlineObject30** | [**InlineObject30**](InlineObject30.md)|  | [optional]

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
**200** | Mute All Check Apps Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateMuteAllCheckGroups"></a>
# **updateMuteAllCheckGroups**
> Object updateMuteAllCheckGroups(inlineObject40)

Mute All Check Groups

Mute all existing check groups.

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
    InlineObject40 inlineObject40 = new InlineObject40(); // InlineObject40 | 
    try {
      Object result = apiInstance.updateMuteAllCheckGroups(inlineObject40);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#updateMuteAllCheckGroups");
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
 **inlineObject40** | [**InlineObject40**](InlineObject40.md)|  | [optional]

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
**200** | Mute All Check Groups Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateMuteAllChecks"></a>
# **updateMuteAllChecks**
> Object updateMuteAllChecks(inlineObject34)

Mute All Checks

Mute all existing checks.

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
    InlineObject34 inlineObject34 = new InlineObject34(); // InlineObject34 | 
    try {
      Object result = apiInstance.updateMuteAllChecks(inlineObject34);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#updateMuteAllChecks");
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
 **inlineObject34** | [**InlineObject34**](InlineObject34.md)|  | [optional]

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
**200** | Mute All Checks Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateMuteCheckApps"></a>
# **updateMuteCheckApps**
> Object updateMuteCheckApps(id, inlineObject32)

Mute Check App

Mute an existing check app.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject32 inlineObject32 = new InlineObject32(); // InlineObject32 | 
    try {
      Object result = apiInstance.updateMuteCheckApps(id, inlineObject32);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#updateMuteCheckApps");
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
 **inlineObject32** | [**InlineObject32**](InlineObject32.md)|  | [optional]

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
**200** | Mute Check App Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateMuteCheckGroups"></a>
# **updateMuteCheckGroups**
> Object updateMuteCheckGroups(id, inlineObject39)

Mute Check Group

Mute an existing check group.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject39 inlineObject39 = new InlineObject39(); // InlineObject39 | 
    try {
      Object result = apiInstance.updateMuteCheckGroups(id, inlineObject39);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#updateMuteCheckGroups");
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
 **inlineObject39** | [**InlineObject39**](InlineObject39.md)|  | [optional]

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
**200** | Mute Check Group Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateMuteChecks"></a>
# **updateMuteChecks**
> Object updateMuteChecks(id, inlineObject36)

Mute Check

Mute an existing check.

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
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject36 inlineObject36 = new InlineObject36(); // InlineObject36 | 
    try {
      Object result = apiInstance.updateMuteChecks(id, inlineObject36);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ChecksApi#updateMuteChecks");
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
 **inlineObject36** | [**InlineObject36**](InlineObject36.md)|  | [optional]

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
**200** | Check Mute Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

