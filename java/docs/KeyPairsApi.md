# KeyPairsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addKeyPairs**](KeyPairsApi.md#addKeyPairs) | **POST** /api/key-pairs | Creates a Key Pair
[**generateKeyPairs**](KeyPairsApi.md#generateKeyPairs) | **POST** /api/key-pairs/generate | Generates a Key Pair
[**getKeyPairs**](KeyPairsApi.md#getKeyPairs) | **GET** /api/key-pairs/{id} | Retrieves a Specific Key Pair
[**removeKeyPairs**](KeyPairsApi.md#removeKeyPairs) | **DELETE** /api/key-pairs/{id} | Deletes a Key Pair


<a name="addKeyPairs"></a>
# **addKeyPairs**
> Object addKeyPairs(inlineObject105)

Creates a Key Pair

Creates a Key Pair.  Private keys are typically optional. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.KeyPairsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    KeyPairsApi apiInstance = new KeyPairsApi(defaultClient);
    InlineObject105 inlineObject105 = new InlineObject105(); // InlineObject105 | 
    try {
      Object result = apiInstance.addKeyPairs(inlineObject105);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling KeyPairsApi#addKeyPairs");
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
 **inlineObject105** | [**InlineObject105**](InlineObject105.md)|  | [optional]

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

<a name="generateKeyPairs"></a>
# **generateKeyPairs**
> InlineResponse20066 generateKeyPairs(inlineObject106)

Generates a Key Pair

Generates a Key Pair. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.KeyPairsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    KeyPairsApi apiInstance = new KeyPairsApi(defaultClient);
    InlineObject106 inlineObject106 = new InlineObject106(); // InlineObject106 | 
    try {
      InlineResponse20066 result = apiInstance.generateKeyPairs(inlineObject106);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling KeyPairsApi#generateKeyPairs");
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
 **inlineObject106** | [**InlineObject106**](InlineObject106.md)|  | [optional]

### Return type

[**InlineResponse20066**](InlineResponse20066.md)

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

<a name="getKeyPairs"></a>
# **getKeyPairs**
> InlineResponse20067 getKeyPairs(id)

Retrieves a Specific Key Pair

Retrieves a specific key pair. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.KeyPairsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    KeyPairsApi apiInstance = new KeyPairsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20067 result = apiInstance.getKeyPairs(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling KeyPairsApi#getKeyPairs");
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

[**InlineResponse20067**](InlineResponse20067.md)

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

<a name="removeKeyPairs"></a>
# **removeKeyPairs**
> Model200Success removeKeyPairs(id)

Deletes a Key Pair

Deletes a specified key pair. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.KeyPairsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    KeyPairsApi apiInstance = new KeyPairsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.removeKeyPairs(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling KeyPairsApi#removeKeyPairs");
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

