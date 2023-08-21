# CloudsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCloudResourcePool**](CloudsApi.md#addCloudResourcePool) | **POST** /api/zones/{zoneId}/resource-pools | Creates a Specified Resource Pool for Specified Cloud
[**addClouds**](CloudsApi.md#addClouds) | **POST** /api/zones | Creates a Cloud
[**getCloudDatastores**](CloudsApi.md#getCloudDatastores) | **GET** /api/zones/{zoneId}/data-stores/{id} | Retrieves a Datastore for Specified Cloud
[**getCloudFolders**](CloudsApi.md#getCloudFolders) | **GET** /api/zones/{zoneId}/folders/{id} | Retrieves a Resource Folder for Specified Cloud
[**getCloudResourcePools**](CloudsApi.md#getCloudResourcePools) | **GET** /api/zones/{zoneId}/resource-pools/{id} | Retrieves a Resource Pool for Specified Cloud
[**getCloudTypes**](CloudsApi.md#getCloudTypes) | **GET** /api/zone-types/{id} | Retrieves a Specific Cloud Type
[**getClouds**](CloudsApi.md#getClouds) | **GET** /api/zones/{id} | Retrieves a Specific Cloud
[**getWikiCloud**](CloudsApi.md#getWikiCloud) | **GET** /api/zones/{id}/wiki | Retrieves a Cloud Wiki Page
[**listCloudDatastores**](CloudsApi.md#listCloudDatastores) | **GET** /api/zones/{zoneId}/data-stores | Retrieves all Datastores for Specified Cloud
[**listCloudFolders**](CloudsApi.md#listCloudFolders) | **GET** /api/zones/{zoneId}/folders | Retrieves all resource folders for Specified Cloud
[**listCloudResourcePools**](CloudsApi.md#listCloudResourcePools) | **GET** /api/zones/{zoneId}/resource-pools | Retrieves all Resource Pools for Specified Cloud
[**listCloudSecurityGroups**](CloudsApi.md#listCloudSecurityGroups) | **GET** /api/zones/{id}/security-groups | Retrieves all Security Groups for a Cloud
[**listCloudTypes**](CloudsApi.md#listCloudTypes) | **GET** /api/zone-types | Retrieves all Cloud Types
[**listClouds**](CloudsApi.md#listClouds) | **GET** /api/zones | Retrieves all Clouds
[**refreshClouds**](CloudsApi.md#refreshClouds) | **POST** /api/zones/{id}/refresh | Refreshes a Cloud
[**removeCloudResourcePools**](CloudsApi.md#removeCloudResourcePools) | **DELETE** /api/zones/{zoneId}/resource-pools/{id} | Deletes a Resource Pool for Specified Cloud
[**removeClouds**](CloudsApi.md#removeClouds) | **DELETE** /api/zones/{id} | Deletes a Cloud
[**updateCloudDatastores**](CloudsApi.md#updateCloudDatastores) | **PUT** /api/zones/{zoneId}/data-stores/{id} | Updates a Specified Datastore for Specified Cloud
[**updateCloudFolders**](CloudsApi.md#updateCloudFolders) | **PUT** /api/zones/{zoneId}/folders/{id} | Updates a Resource Folder for Specified Cloud
[**updateCloudLogo**](CloudsApi.md#updateCloudLogo) | **POST** /api/zones/{id}/update-logo | Update Logo For Cloud
[**updateCloudResourcePool**](CloudsApi.md#updateCloudResourcePool) | **PUT** /api/zones/{zoneId}/resource-pools/{id} | Updates a Specified Resource Pool for Specified Cloud
[**updateCloudSecurityGroups**](CloudsApi.md#updateCloudSecurityGroups) | **POST** /api/zones/{id}/security-groups | Sets Security Groups for a Cloud
[**updateClouds**](CloudsApi.md#updateClouds) | **PUT** /api/zones/{id} | Updates a Cloud
[**updateWikiCloud**](CloudsApi.md#updateWikiCloud) | **PUT** /api/zones/{id}/wiki | Update a Cloud Wiki Page


<a name="addCloudResourcePool"></a>
# **addCloudResourcePool**
> InlineResponse20024 addCloudResourcePool(zoneId, inlineObject45)

Creates a Specified Resource Pool for Specified Cloud

Creates a resource pool for specified cloud. Only certain types of clouds support creating and deleting resource pools. Configuration options vary by type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    BigDecimal zoneId = new BigDecimal(78); // BigDecimal | The ID of the cloud
    InlineObject45 inlineObject45 = new InlineObject45(); // InlineObject45 | 
    try {
      InlineResponse20024 result = apiInstance.addCloudResourcePool(zoneId, inlineObject45);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#addCloudResourcePool");
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
 **zoneId** | **BigDecimal**| The ID of the cloud |
 **inlineObject45** | [**InlineObject45**](InlineObject45.md)|  | [optional]

### Return type

[**InlineResponse20024**](InlineResponse20024.md)

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

<a name="addClouds"></a>
# **addClouds**
> Object addClouds(inlineObject41)

Creates a Cloud

Creates a cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    InlineObject41 inlineObject41 = new InlineObject41(); // InlineObject41 | 
    try {
      Object result = apiInstance.addClouds(inlineObject41);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#addClouds");
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
 **inlineObject41** | [**InlineObject41**](InlineObject41.md)|  | [optional]

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

<a name="getCloudDatastores"></a>
# **getCloudDatastores**
> Object getCloudDatastores(zoneId, id)

Retrieves a Datastore for Specified Cloud

Data Stores can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves a specific datastore under a cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    BigDecimal zoneId = new BigDecimal(78); // BigDecimal | The ID of the cloud
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getCloudDatastores(zoneId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#getCloudDatastores");
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
 **zoneId** | **BigDecimal**| The ID of the cloud |
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

<a name="getCloudFolders"></a>
# **getCloudFolders**
> Object getCloudFolders(zoneId, id)

Retrieves a Resource Folder for Specified Cloud

Resource Folders can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves a specific folder under a cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    BigDecimal zoneId = new BigDecimal(78); // BigDecimal | The ID of the cloud
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getCloudFolders(zoneId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#getCloudFolders");
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
 **zoneId** | **BigDecimal**| The ID of the cloud |
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

<a name="getCloudResourcePools"></a>
# **getCloudResourcePools**
> Object getCloudResourcePools(zoneId, id)

Retrieves a Resource Pool for Specified Cloud

This endpoint retrieves a specific resource pool. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    BigDecimal zoneId = new BigDecimal(78); // BigDecimal | The ID of the cloud
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getCloudResourcePools(zoneId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#getCloudResourcePools");
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
 **zoneId** | **BigDecimal**| The ID of the cloud |
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

<a name="getCloudTypes"></a>
# **getCloudTypes**
> InlineResponse20020 getCloudTypes(id)

Retrieves a Specific Cloud Type

Retrieves a specific cloud type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20020 result = apiInstance.getCloudTypes(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#getCloudTypes");
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

[**InlineResponse20020**](InlineResponse20020.md)

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

<a name="getClouds"></a>
# **getClouds**
> InlineResponse20021 getClouds(id)

Retrieves a Specific Cloud

Retrieves a specific cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20021 result = apiInstance.getClouds(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#getClouds");
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

[**InlineResponse20021**](InlineResponse20021.md)

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

<a name="getWikiCloud"></a>
# **getWikiCloud**
> InlineResponse200168 getWikiCloud(id)

Retrieves a Cloud Wiki Page

This endpoint retrieves a cloud Wiki page. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200168 result = apiInstance.getWikiCloud(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#getWikiCloud");
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

<a name="listCloudDatastores"></a>
# **listCloudDatastores**
> Object listCloudDatastores(zoneId, name, phrase, max, sort, direction)

Retrieves all Datastores for Specified Cloud

Data Stores can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves all data stores under a cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    BigDecimal zoneId = new BigDecimal(78); // BigDecimal | The ID of the cloud
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    try {
      Object result = apiInstance.listCloudDatastores(zoneId, name, phrase, max, sort, direction);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#listCloudDatastores");
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
 **zoneId** | **BigDecimal**| The ID of the cloud |
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]

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

<a name="listCloudFolders"></a>
# **listCloudFolders**
> Object listCloudFolders(zoneId, name, phrase, max)

Retrieves all resource folders for Specified Cloud

Resource Folders can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves all resource folders under a cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    BigDecimal zoneId = new BigDecimal(78); // BigDecimal | The ID of the cloud
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    try {
      Object result = apiInstance.listCloudFolders(zoneId, name, phrase, max);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#listCloudFolders");
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
 **zoneId** | **BigDecimal**| The ID of the cloud |
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]

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

<a name="listCloudResourcePools"></a>
# **listCloudResourcePools**
> Object listCloudResourcePools(zoneId, name, phrase, max)

Retrieves all Resource Pools for Specified Cloud

Resource Pools can be managed for each Compute Zone (Cloud) in your infrastructure. This endpoint retrieves all resource pools under a cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    BigDecimal zoneId = new BigDecimal(78); // BigDecimal | The ID of the cloud
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    try {
      Object result = apiInstance.listCloudResourcePools(zoneId, name, phrase, max);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#listCloudResourcePools");
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
 **zoneId** | **BigDecimal**| The ID of the cloud |
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]

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

<a name="listCloudSecurityGroups"></a>
# **listCloudSecurityGroups**
> Object listCloudSecurityGroups(id)

Retrieves all Security Groups for a Cloud

Retrieves all security groups for a cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.listCloudSecurityGroups(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#listCloudSecurityGroups");
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

<a name="listCloudTypes"></a>
# **listCloudTypes**
> Object listCloudTypes(max, offset, sort, direction, name, code, phrase, provisionType)

Retrieves all Cloud Types

Fetch a paginated list of available cloud types. This returns the configuration options for each type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String code = "azr"; // String | If specified will return an exact match on code
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String provisionType = "provisionType_example"; // String | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings. 
    try {
      Object result = apiInstance.listCloudTypes(max, offset, sort, direction, name, code, phrase, provisionType);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#listCloudTypes");
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
 **code** | **String**| If specified will return an exact match on code | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **provisionType** | **String**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional] [enum: AKS, Alibaba, Amazon, ARM, Azure, Azure SQL Server, Canonical MaaS, Cloud Foundry App, Cloud Foundry Diego, Cloud Foundry Service, CloudFormation, Digital Ocean, Docker, Docker Swarm, EKS, ESXi, External, Fusion, GKE, Google, Helm, HP Oneview, Huawei, Hyper-V, Hypervisor, IBM Cloud, Kubernetes, KVM, Lumen Edge, lxc, Manual, Manual Kubernetes, Nutanix, Open Telekom, OpenStack, Oracle Cloud, Oracle VM, RDS, SCVMM, SCVMM Hypervisor, SoftLayer, Spoon, Terraform, UCS, Unmanaged, UpCloud, Vagrant, VCD VAPP, VCloud Air, vCloud Director, VirtualBox, Virtustream, VMware, Workflow, Xen]

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

<a name="listClouds"></a>
# **listClouds**
> Object listClouds(lastUpdated, type, groupId, max, offset, sort, direction, phrase, name)

Retrieves all Clouds

Retrieves all clouds. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    String type = "type_example"; // String | If specified will return all zones by cloud type code. Refer to `Zone Types` API for up to date listings. 
    Long groupId = 13L; // Long | If specified will return all zones assigned to a server group by id. Accepts multiple values.
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    try {
      Object result = apiInstance.listClouds(lastUpdated, type, groupId, max, offset, sort, direction, phrase, name);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#listClouds");
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
 **lastUpdated** | **OffsetDateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **type** | **String**| If specified will return all zones by cloud type code. Refer to &#x60;Zone Types&#x60; API for up to date listings.  | [optional] [enum: alibaba, amazon, azure, azurestack, dell, digitalocean, googlecloud, oneview, huawei, hyperv, bluemix, bluemixCloudFoundry, centurylink-edge, macstadium, standard, nutanix, opentelekom, openstack, oraclecloud, oraclevm, platform9, scvmm, softlayer, ucs, upcloud, vcd, vmwareCloudAws, esxi, fusion, vmware, xenserver]
 **groupId** | **Long**| If specified will return all zones assigned to a server group by id. Accepts multiple values. | [optional]
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

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

<a name="refreshClouds"></a>
# **refreshClouds**
> Model200Success refreshClouds(id, inlineObject47)

Refreshes a Cloud

Refreshes a cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject47 inlineObject47 = new InlineObject47(); // InlineObject47 | 
    try {
      Model200Success result = apiInstance.refreshClouds(id, inlineObject47);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#refreshClouds");
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
 **inlineObject47** | [**InlineObject47**](InlineObject47.md)|  | [optional]

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

<a name="removeCloudResourcePools"></a>
# **removeCloudResourcePools**
> Model200Success removeCloudResourcePools(zoneId, id)

Deletes a Resource Pool for Specified Cloud

Deletes a resource pool for specified Cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    BigDecimal zoneId = new BigDecimal(78); // BigDecimal | The ID of the cloud
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.removeCloudResourcePools(zoneId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#removeCloudResourcePools");
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
 **zoneId** | **BigDecimal**| The ID of the cloud |
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

<a name="removeClouds"></a>
# **removeClouds**
> Model200Success removeClouds(id, removeResources)

Deletes a Cloud

Deletes a specified Cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Boolean removeResources = false; // Boolean | Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute.
    try {
      Model200Success result = apiInstance.removeClouds(id, removeResources);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#removeClouds");
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
 **removeResources** | **Boolean**| Removing associated resources will delete the instances and the associated resources underneath.  This includes Virtual Machines and other forms of compute. | [optional] [default to false]

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

<a name="updateCloudDatastores"></a>
# **updateCloudDatastores**
> InlineResponse20022 updateCloudDatastores(zoneId, id, inlineObject43)

Updates a Specified Datastore for Specified Cloud

Updates a datastore for specified cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    BigDecimal zoneId = new BigDecimal(78); // BigDecimal | The ID of the cloud
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject43 inlineObject43 = new InlineObject43(); // InlineObject43 | 
    try {
      InlineResponse20022 result = apiInstance.updateCloudDatastores(zoneId, id, inlineObject43);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#updateCloudDatastores");
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
 **zoneId** | **BigDecimal**| The ID of the cloud |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject43** | [**InlineObject43**](InlineObject43.md)|  | [optional]

### Return type

[**InlineResponse20022**](InlineResponse20022.md)

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

<a name="updateCloudFolders"></a>
# **updateCloudFolders**
> InlineResponse20023 updateCloudFolders(zoneId, id, inlineObject44)

Updates a Resource Folder for Specified Cloud

Updates a resource folder for specified cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    BigDecimal zoneId = new BigDecimal(78); // BigDecimal | The ID of the cloud
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject44 inlineObject44 = new InlineObject44(); // InlineObject44 | 
    try {
      InlineResponse20023 result = apiInstance.updateCloudFolders(zoneId, id, inlineObject44);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#updateCloudFolders");
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
 **zoneId** | **BigDecimal**| The ID of the cloud |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject44** | [**InlineObject44**](InlineObject44.md)|  | [optional]

### Return type

[**InlineResponse20023**](InlineResponse20023.md)

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

<a name="updateCloudLogo"></a>
# **updateCloudLogo**
> Model200Success updateCloudLogo(id, logo, darkLogo)

Update Logo For Cloud

Use this command to update the logo and dark logo images for a cloud. This endpoint expects multipart form data as the request format, not JSON. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    File logo = new File("/path/to/file"); // File | Logo File png,jpg,svg
    File darkLogo = new File("/path/to/file"); // File | Dark Logo File png,jpg,svg
    try {
      Model200Success result = apiInstance.updateCloudLogo(id, logo, darkLogo);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#updateCloudLogo");
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
 **logo** | **File**| Logo File png,jpg,svg | [optional]
 **darkLogo** | **File**| Dark Logo File png,jpg,svg | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateCloudResourcePool"></a>
# **updateCloudResourcePool**
> InlineResponse20024 updateCloudResourcePool(zoneId, id, inlineObject46)

Updates a Specified Resource Pool for Specified Cloud

Updates a resource pool for specified cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    BigDecimal zoneId = new BigDecimal(78); // BigDecimal | The ID of the cloud
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject46 inlineObject46 = new InlineObject46(); // InlineObject46 | 
    try {
      InlineResponse20024 result = apiInstance.updateCloudResourcePool(zoneId, id, inlineObject46);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#updateCloudResourcePool");
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
 **zoneId** | **BigDecimal**| The ID of the cloud |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject46** | [**InlineObject46**](InlineObject46.md)|  | [optional]

### Return type

[**InlineResponse20024**](InlineResponse20024.md)

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

<a name="updateCloudSecurityGroups"></a>
# **updateCloudSecurityGroups**
> Object updateCloudSecurityGroups(id, inlineObject49)

Sets Security Groups for a Cloud

Sets security groups for acloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject49 inlineObject49 = new InlineObject49(); // InlineObject49 | 
    try {
      Object result = apiInstance.updateCloudSecurityGroups(id, inlineObject49);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#updateCloudSecurityGroups");
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
 **inlineObject49** | [**InlineObject49**](InlineObject49.md)|  | [optional]

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

<a name="updateClouds"></a>
# **updateClouds**
> Object updateClouds(id, inlineObject42)

Updates a Cloud

Updates a cloud. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject42 inlineObject42 = new InlineObject42(); // InlineObject42 | 
    try {
      Object result = apiInstance.updateClouds(id, inlineObject42);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#updateClouds");
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
 **inlineObject42** | [**InlineObject42**](InlineObject42.md)|  | [optional]

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

<a name="updateWikiCloud"></a>
# **updateWikiCloud**
> Object updateWikiCloud(id, inlineObject274)

Update a Cloud Wiki Page

Updates a cloud Wiki page. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CloudsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CloudsApi apiInstance = new CloudsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject274 inlineObject274 = new InlineObject274(); // InlineObject274 | 
    try {
      Object result = apiInstance.updateWikiCloud(id, inlineObject274);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CloudsApi#updateWikiCloud");
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
 **inlineObject274** | [**InlineObject274**](InlineObject274.md)|  | [optional]

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

