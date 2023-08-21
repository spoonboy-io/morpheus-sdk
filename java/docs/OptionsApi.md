# OptionsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getOptionSourceData**](OptionsApi.md#getOptionSourceData) | **GET** /api/options/{optionSource} | Get Option Source Data
[**listCodeRepositories**](OptionsApi.md#listCodeRepositories) | **GET** /api/options/codeRepositories | Retrieves a list of Code/GIT Repositories
[**listOptionNetworkOptions**](OptionsApi.md#listOptionNetworkOptions) | **GET** /api/options/zoneNetworkOptions | Retrieves network options by zone/cloud
[**listOptionValues**](OptionsApi.md#listOptionValues) | **GET** /api/options/list | Retrieves input option values


<a name="getOptionSourceData"></a>
# **getOptionSourceData**
> Object getOptionSourceData(optionSource)

Get Option Source Data

Returns a list of name/value pairs for option-type models. Some option-types depend on input data for proper representation. This typically includes zoneId or siteId for the item being provisioned as request parameters or sometimes previous option type parameters. Each option returned has a &#x60;value&#x60;, which is often the &#x60;id&#x60;, but may be a &#x60;code&#x60; or other attribute. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.OptionsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    OptionsApi apiInstance = new OptionsApi(defaultClient);
    String optionSource = "keyPairs"; // String | `optionSource` to be listed
    try {
      Object result = apiInstance.getOptionSourceData(optionSource);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling OptionsApi#getOptionSourceData");
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
 **optionSource** | **String**| &#x60;optionSource&#x60; to be listed |

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

<a name="listCodeRepositories"></a>
# **listCodeRepositories**
> Model200Success listCodeRepositories(integrationId)

Retrieves a list of Code/GIT Repositories

Retrieves a list of Code/GIT Repositories 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.OptionsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    OptionsApi apiInstance = new OptionsApi(defaultClient);
    Long integrationId = 33L; // Long | Filter by an integration Id.
    try {
      Model200Success result = apiInstance.listCodeRepositories(integrationId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling OptionsApi#listCodeRepositories");
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
 **integrationId** | **Long**| Filter by an integration Id. | [optional]

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

<a name="listOptionNetworkOptions"></a>
# **listOptionNetworkOptions**
> ZoneNetworkOptions listOptionNetworkOptions(zoneId, provisionTypeId)

Retrieves network options by zone/cloud

This endpoint can be used to see which network options are available for a given cloud (zoneId) and provision type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.OptionsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    OptionsApi apiInstance = new OptionsApi(defaultClient);
    Long zoneId = 3L; // Long | The Zone ID for Filtering
    Long provisionTypeId = 22L; // Long | Provision type filter, restricts query to only load service plans of specified provision type
    try {
      ZoneNetworkOptions result = apiInstance.listOptionNetworkOptions(zoneId, provisionTypeId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling OptionsApi#listOptionNetworkOptions");
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
 **zoneId** | **Long**| The Zone ID for Filtering | [optional]
 **provisionTypeId** | **Long**| Provision type filter, restricts query to only load service plans of specified provision type | [optional]

### Return type

[**ZoneNetworkOptions**](ZoneNetworkOptions.md)

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

<a name="listOptionValues"></a>
# **listOptionValues**
> Object listOptionValues(optionTypeId, config)

Retrieves input option values

Retrieves all input option values.  Can be used with parameters to supply dependent input values. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.OptionsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    OptionsApi apiInstance = new OptionsApi(defaultClient);
    Long optionTypeId = 56L; // Long | Input or Option Type ID
    Object config = {"config.fieldName1":"value1"}; // Object | Input parameters are required if the input is dependent on them.  Fields must be prefixed with `config.`
    try {
      Object result = apiInstance.listOptionValues(optionTypeId, config);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling OptionsApi#listOptionValues");
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
 **optionTypeId** | **Long**| Input or Option Type ID |
 **config** | [**Object**](.md)| Input parameters are required if the input is dependent on them.  Fields must be prefixed with &#x60;config.&#x60; | [optional]

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

