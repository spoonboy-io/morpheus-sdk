# ActivityApi

All URIs are relative to *https://CHANGEME*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**listActivity**](ActivityApi.md#listActivity) | **GET** /api/activity | Retrieves Activity |


<a id="listActivity"></a>
# **listActivity**
> ListActivity200Response listActivity(max, offset, sort, order, phrase, name, userId, tenantId, timeframe, start, end)

Retrieves Activity

Retrieves activity. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ActivityApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ActivityApi apiInstance = new ActivityApi(defaultClient);
    Long max = 25L; // Long | Maximum number of records to return
    Long offset = 0L; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "name"; // String | Sort order, the name of the property to sort by
    String order = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    Long userId = 6L; // Long | Filter by User ID
    BigDecimal tenantId = new BigDecimal("3"); // BigDecimal | Filter by Tenant ID. Only available to the master account.
    String timeframe = "month"; // String | Filter by a timeframe (ex - today, yesterday, week, month, 3months)
    OffsetDateTime start = OffsetDateTime.now(); // OffsetDateTime | Filter by activity on or after a date(time). Default is 1 month prior
    OffsetDateTime end = OffsetDateTime.now(); // OffsetDateTime | Filter by activity on or before a date(time). Default is current date
    try {
      ListActivity200Response result = apiInstance.listActivity(max, offset, sort, order, phrase, name, userId, tenantId, timeframe, start, end);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ActivityApi#listActivity");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **max** | **Long**| Maximum number of records to return | [optional] [default to 25] |
| **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0] |
| **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to name] |
| **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc] |
| **phrase** | **String**| Search phrase for partial matches on name or description | [optional] |
| **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] |
| **userId** | **Long**| Filter by User ID | [optional] |
| **tenantId** | **BigDecimal**| Filter by Tenant ID. Only available to the master account. | [optional] |
| **timeframe** | **String**| Filter by a timeframe (ex - today, yesterday, week, month, 3months) | [optional] [default to month] |
| **start** | **OffsetDateTime**| Filter by activity on or after a date(time). Default is 1 month prior | [optional] |
| **end** | **OffsetDateTime**| Filter by activity on or before a date(time). Default is current date | [optional] |

### Return type

[**ListActivity200Response**](ListActivity200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

