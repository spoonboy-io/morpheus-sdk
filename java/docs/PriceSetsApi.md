# PriceSetsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addPriceSets**](PriceSetsApi.md#addPriceSets) | **POST** /api/price-sets | Creates a Price Set
[**deactivatePriceSets**](PriceSetsApi.md#deactivatePriceSets) | **PUT** /api/price-sets/{id}/deactivate | Deactivates a Price Set
[**getPriceSets**](PriceSetsApi.md#getPriceSets) | **GET** /api/price-sets/{id} | Retrieves a Specific Price Set
[**listPriceSets**](PriceSetsApi.md#listPriceSets) | **GET** /api/price-sets | Retrieves all Price Sets
[**updatePriceSets**](PriceSetsApi.md#updatePriceSets) | **PUT** /api/price-sets/{id} | Updates a Price Set


<a name="addPriceSets"></a>
# **addPriceSets**
> Object addPriceSets(inlineObject198)

Creates a Price Set

Creates a price set. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PriceSetsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    PriceSetsApi apiInstance = new PriceSetsApi(defaultClient);
    InlineObject198 inlineObject198 = new InlineObject198(); // InlineObject198 | 
    try {
      Object result = apiInstance.addPriceSets(inlineObject198);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PriceSetsApi#addPriceSets");
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
 **inlineObject198** | [**InlineObject198**](InlineObject198.md)|  | [optional]

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

<a name="deactivatePriceSets"></a>
# **deactivatePriceSets**
> Model200Success deactivatePriceSets(id)

Deactivates a Price Set

Deactivates a price set. This does the same thing as the delete action in the UI, hiding it and making it unavailable to new resources. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PriceSetsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    PriceSetsApi apiInstance = new PriceSetsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deactivatePriceSets(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PriceSetsApi#deactivatePriceSets");
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

<a name="getPriceSets"></a>
# **getPriceSets**
> InlineResponse200123 getPriceSets(id)

Retrieves a Specific Price Set

Retrieves a specific price set. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PriceSetsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    PriceSetsApi apiInstance = new PriceSetsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200123 result = apiInstance.getPriceSets(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PriceSetsApi#getPriceSets");
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

[**InlineResponse200123**](InlineResponse200123.md)

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

<a name="listPriceSets"></a>
# **listPriceSets**
> Object listPriceSets(max, offset, sort, direction, phrase, name, includeInactive, type)

Retrieves all Price Sets

Retrieves all price sets. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PriceSetsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    PriceSetsApi apiInstance = new PriceSetsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    Boolean includeInactive = true; // Boolean | If true, include inactive prices in the results
    String type = "type_example"; // String | Filter by type code
    try {
      Object result = apiInstance.listPriceSets(max, offset, sort, direction, phrase, name, includeInactive, type);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PriceSetsApi#listPriceSets");
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
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **includeInactive** | **Boolean**| If true, include inactive prices in the results | [optional]
 **type** | **String**| Filter by type code | [optional]

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

<a name="updatePriceSets"></a>
# **updatePriceSets**
> Object updatePriceSets(id, inlineObject199)

Updates a Price Set

Updates a price set. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PriceSetsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    PriceSetsApi apiInstance = new PriceSetsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject199 inlineObject199 = new InlineObject199(); // InlineObject199 | 
    try {
      Object result = apiInstance.updatePriceSets(id, inlineObject199);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PriceSetsApi#updatePriceSets");
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
 **inlineObject199** | [**InlineObject199**](InlineObject199.md)|  | [optional]

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

