# SecurityScansApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getSecurityScans**](SecurityScansApi.md#getSecurityScans) | **GET** /api/security-scans/{id} | Retrieves a Specific Security Scan
[**listSecurityScans**](SecurityScansApi.md#listSecurityScans) | **GET** /api/security-scans | Retrieves all Security Scans


<a name="getSecurityScans"></a>
# **getSecurityScans**
> Object getSecurityScans(id, results)

Retrieves a Specific Security Scan

Retrieves a specific security scan. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityScansApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityScansApi apiInstance = new SecurityScansApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Boolean results = false; // Boolean | Include the `results` object in the response under the security scan. This is a potentially very large object containing the raw results of the scan.
    try {
      Object result = apiInstance.getSecurityScans(id, results);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityScansApi#getSecurityScans");
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
 **results** | **Boolean**| Include the &#x60;results&#x60; object in the response under the security scan. This is a potentially very large object containing the raw results of the scan. | [optional] [default to false]

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

<a name="listSecurityScans"></a>
# **listSecurityScans**
> Object listSecurityScans(max, offset, sort, direction, phrase, securityPackageId, serverId, results)

Retrieves all Security Scans

Retrieves all security scans. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityScansApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityScansApi apiInstance = new SecurityScansApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"scanDate\""; // String | Sort order, the name of the property to sort by
    String direction = "desc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description of security package
    Long securityPackageId = 56L; // Long | Filter results by security package id(s). This parameter can be passed multiple times to match more than one id.
    Long serverId = 97L; // Long | The Server ID for Filtering
    Boolean results = false; // Boolean | Include the `results` object in the response under each security scan. This is a potentially very large object containing the raw results of the scan.
    try {
      Object result = apiInstance.listSecurityScans(max, offset, sort, direction, phrase, securityPackageId, serverId, results);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityScansApi#listSecurityScans");
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
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;scanDate&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to desc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description of security package | [optional]
 **securityPackageId** | **Long**| Filter results by security package id(s). This parameter can be passed multiple times to match more than one id. | [optional]
 **serverId** | **Long**| The Server ID for Filtering | [optional]
 **results** | **Boolean**| Include the &#x60;results&#x60; object in the response under each security scan. This is a potentially very large object containing the raw results of the scan. | [optional] [default to false]

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

