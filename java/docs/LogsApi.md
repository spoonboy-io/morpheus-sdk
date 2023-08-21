# LogsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listLogs**](LogsApi.md#listLogs) | **GET** /api/logs | Retrieves Logs


<a name="listLogs"></a>
# **listLogs**
> Log listLogs(max, offset, sort, order, query, message, sourceType, typeCode, objectId, token, level, startMs, endMs, startDateTime, endDateTime, containers, servers, clusterId)

Retrieves Logs

Retrieves logs based on filters provided. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.LogsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    LogsApi apiInstance = new LogsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String order = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String query = "query_example"; // String | Alias for phrase
    String message = "message_example"; // String | Filter by message
    String sourceType = "sourceType_example"; // String | Filter by source type
    String typeCode = "typeCode_example"; // String | Filter by code type
    Long objectId = 15L; // Long | Filter by objectId
    String token = "token_example"; // String | Filter by token
    String level = "ERROR"; // String | Filter by log level. Multiple values can be passed pipe delimited.
    Long startMs = 1657309873L; // Long | Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified.
    Long endMs = 1657309873L; // Long | Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified.
    OffsetDateTime startDateTime = OffsetDateTime.now(); // OffsetDateTime | Start Date timestamp (ISO 8601)
    OffsetDateTime endDateTime = OffsetDateTime.now(); // OffsetDateTime | End Date timestamp (ISO 8601)
    Long containers = 135L; // Long | The Container ID(s) for filtering. Accepts multiple values.
    Long servers = 135L; // Long | The Server ID(s) for filtering. Accepts multiple values.
    Long clusterId = 135L; // Long | The Cluster ID(s) for filtering. Accepts multiple values.
    try {
      Log result = apiInstance.listLogs(max, offset, sort, order, query, message, sourceType, typeCode, objectId, token, level, startMs, endMs, startDateTime, endDateTime, containers, servers, clusterId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling LogsApi#listLogs");
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
 **order** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **query** | **String**| Alias for phrase | [optional]
 **message** | **String**| Filter by message | [optional]
 **sourceType** | **String**| Filter by source type | [optional]
 **typeCode** | **String**| Filter by code type | [optional]
 **objectId** | **Long**| Filter by objectId | [optional]
 **token** | **String**| Filter by token | [optional]
 **level** | **String**| Filter by log level. Multiple values can be passed pipe delimited. | [optional] [enum: DEBUG, INFO, ERROR]
 **startMs** | **Long**| Date filter in milliseconds (unix epoch), restricts query to only load logs updated more recently than the time specified. | [optional]
 **endMs** | **Long**| Date filter in milliseconds (unix epoch), restricts query to only load logs updated before the time specified. | [optional]
 **startDateTime** | **OffsetDateTime**| Start Date timestamp (ISO 8601) | [optional]
 **endDateTime** | **OffsetDateTime**| End Date timestamp (ISO 8601) | [optional]
 **containers** | **Long**| The Container ID(s) for filtering. Accepts multiple values. | [optional]
 **servers** | **Long**| The Server ID(s) for filtering. Accepts multiple values. | [optional]
 **clusterId** | **Long**| The Cluster ID(s) for filtering. Accepts multiple values. | [optional]

### Return type

[**Log**](Log.md)

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

