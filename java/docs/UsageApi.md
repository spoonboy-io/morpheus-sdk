# UsageApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listUsages**](UsageApi.md#listUsages) | **GET** /api/usage | Retrieves Usage Records


<a name="listUsages"></a>
# **listUsages**
> Usages listUsages(max, offset, sort, direction, phrase, startDate, endDate)

Retrieves Usage Records

Retrieves a paginated list of usage records. The usages are scoped to only include resources you have access to. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsageApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsageApi apiInstance = new UsageApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    try {
      Usages result = apiInstance.listUsages(max, offset, sort, direction, phrase, startDate, endDate);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsageApi#listUsages");
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
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional]
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional]

### Return type

[**Usages**](Usages.md)

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

