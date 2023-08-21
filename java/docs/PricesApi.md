# PricesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addPrices**](PricesApi.md#addPrices) | **POST** /api/prices | Creates a Price
[**deactivatePrices**](PricesApi.md#deactivatePrices) | **PUT** /api/prices/{id}/deactivate | Deactivates a Price
[**getPrices**](PricesApi.md#getPrices) | **GET** /api/prices/{id} | Retrieves a Specific Price
[**listPrices**](PricesApi.md#listPrices) | **GET** /api/prices | Retrieves all Prices
[**updatePrices**](PricesApi.md#updatePrices) | **PUT** /api/prices/{id} | Updates a Price


<a name="addPrices"></a>
# **addPrices**
> Object addPrices(inlineObject200)

Creates a Price

Creates a price. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PricesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    PricesApi apiInstance = new PricesApi(defaultClient);
    InlineObject200 inlineObject200 = new InlineObject200(); // InlineObject200 | 
    try {
      Object result = apiInstance.addPrices(inlineObject200);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PricesApi#addPrices");
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
 **inlineObject200** | [**InlineObject200**](InlineObject200.md)|  | [optional]

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

<a name="deactivatePrices"></a>
# **deactivatePrices**
> Model200Success deactivatePrices(id)

Deactivates a Price

Deactivates a price. This does the same thing as the delete action in the UI, hiding it and making it unavailable to new resources. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PricesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    PricesApi apiInstance = new PricesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deactivatePrices(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PricesApi#deactivatePrices");
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

<a name="getPrices"></a>
# **getPrices**
> InlineResponse200124 getPrices(id)

Retrieves a Specific Price

Retrieves a specific price. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PricesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    PricesApi apiInstance = new PricesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200124 result = apiInstance.getPrices(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PricesApi#getPrices");
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

[**InlineResponse200124**](InlineResponse200124.md)

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

<a name="listPrices"></a>
# **listPrices**
> Object listPrices(max, offset, sort, direction, phrase, name, includeInactive, priceType, platform, currency, type)

Retrieves all Prices

Retrieves all prices. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PricesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    PricesApi apiInstance = new PricesApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    Boolean includeInactive = true; // Boolean | If true, include inactive prices in the results
    String priceType = "priceType_example"; // String | Restricts query to only load only prices with specified priceType. * `fixed` - Everything * `compute` - Memory + CPU * `memory` - Memory * `cores` - Cores * `storage` - Storage * `datastore` - Datastore * `platform` - Platform * `software` - Software 
    String platform = "linux"; // String | Restricts query to only load only prices with specified platform
    String currency = "USD"; // String | Restricts query to only load only prices with specified currency
    String type = "type_example"; // String | Filter by type code
    try {
      Object result = apiInstance.listPrices(max, offset, sort, direction, phrase, name, includeInactive, priceType, platform, currency, type);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PricesApi#listPrices");
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
 **priceType** | **String**| Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software  | [optional] [enum: fixed, compute, memory, cores, storage, datastore, platform, software]
 **platform** | **String**| Restricts query to only load only prices with specified platform | [optional]
 **currency** | **String**| Restricts query to only load only prices with specified currency | [optional]
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

<a name="updatePrices"></a>
# **updatePrices**
> Object updatePrices(id, inlineObject201)

Updates a Price

Updates a price. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PricesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    PricesApi apiInstance = new PricesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject201 inlineObject201 = new InlineObject201(); // InlineObject201 | 
    try {
      Object result = apiInstance.updatePrices(id, inlineObject201);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PricesApi#updatePrices");
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
 **inlineObject201** | [**InlineObject201**](InlineObject201.md)|  | [optional]

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

