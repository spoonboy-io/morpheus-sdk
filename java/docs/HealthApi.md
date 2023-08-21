# HealthApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**acknowledgeHealthAlarm**](HealthApi.md#acknowledgeHealthAlarm) | **PUT** /api/health/alarms/{id}/acknowledge | Acknowledge a Health Alarm
[**acknowledgeHealthAlarms**](HealthApi.md#acknowledgeHealthAlarms) | **PUT** /api/health/alarms/acknowledge | Acknowledge Many Health Alarms
[**exportHealthLogs**](HealthApi.md#exportHealthLogs) | **GET** /api/health/logs/export | Export Appliance Health Logs
[**getHealthAlarms**](HealthApi.md#getHealthAlarms) | **GET** /api/health/alarms/{id} | Retrieves a Specific Appliance Health Alarm
[**listHealth**](HealthApi.md#listHealth) | **GET** /api/health | Retrieves Appliance Health
[**listHealthAlarms**](HealthApi.md#listHealthAlarms) | **GET** /api/health/alarms | Retrieves Appliance Health Alarms
[**listHealthLogs**](HealthApi.md#listHealthLogs) | **GET** /api/health/logs | Retrieves Appliance Health Logs
[**ping**](HealthApi.md#ping) | **GET** /api/ping | Basic information about current Morpheus Installation


<a name="acknowledgeHealthAlarm"></a>
# **acknowledgeHealthAlarm**
> Model200Success acknowledgeHealthAlarm(id, inlineObject81)

Acknowledge a Health Alarm

Acknowledge a specific health alarm.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HealthApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HealthApi apiInstance = new HealthApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject81 inlineObject81 = new InlineObject81(); // InlineObject81 | 
    try {
      Model200Success result = apiInstance.acknowledgeHealthAlarm(id, inlineObject81);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HealthApi#acknowledgeHealthAlarm");
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
 **inlineObject81** | [**InlineObject81**](InlineObject81.md)|  | [optional]

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

<a name="acknowledgeHealthAlarms"></a>
# **acknowledgeHealthAlarms**
> Model200Success acknowledgeHealthAlarms(inlineObject80)

Acknowledge Many Health Alarms

Acknowledge health alarms.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HealthApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HealthApi apiInstance = new HealthApi(defaultClient);
    InlineObject80 inlineObject80 = new InlineObject80(); // InlineObject80 | 
    try {
      Model200Success result = apiInstance.acknowledgeHealthAlarms(inlineObject80);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HealthApi#acknowledgeHealthAlarms");
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
 **inlineObject80** | [**InlineObject80**](InlineObject80.md)|  | [optional]

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

<a name="exportHealthLogs"></a>
# **exportHealthLogs**
> File exportHealthLogs(max, offset, sort, direction, phrase, name, acknowledged, startDate, endDate, reverse)

Export Appliance Health Logs

This endpoint downloads the morpheus appliance logs as a file attachment. By default, the most recent 10,000 log entries are returned, with the newest at the end of the file. The format for each log entry is &#x60;timestamp&#x60; &#x60;level&#x60; &#x60;message&#x60;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HealthApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HealthApi apiInstance = new HealthApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    Boolean acknowledged = false; // Boolean | True or False flag for Acknowledged items
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    Boolean reverse = true; // Boolean | Reverse order of records. This `true` by default when sort and direction are not passed, but `false` by default if either is passed. This means that by default the newest log entries are the bottom of the file. 
    try {
      File result = apiInstance.exportHealthLogs(max, offset, sort, direction, phrase, name, acknowledged, startDate, endDate, reverse);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HealthApi#exportHealthLogs");
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
 **acknowledged** | **Boolean**| True or False flag for Acknowledged items | [optional]
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional]
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional]
 **reverse** | **Boolean**| Reverse order of records. This &#x60;true&#x60; by default when sort and direction are not passed, but &#x60;false&#x60; by default if either is passed. This means that by default the newest log entries are the bottom of the file.  | [optional]

### Return type

[**File**](File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream, application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  * Content-Disposition -  <br>  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getHealthAlarms"></a>
# **getHealthAlarms**
> Object getHealthAlarms(id)

Retrieves a Specific Appliance Health Alarm

This endpoint will retrieve a specific health alarm by ID.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HealthApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HealthApi apiInstance = new HealthApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getHealthAlarms(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HealthApi#getHealthAlarms");
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

<a name="listHealth"></a>
# **listHealth**
> Object listHealth()

Retrieves Appliance Health

This endpoint retrieves health info about the appliance such as cpu, memory and database usage. Elasticsearch statistics and queue usage are also returned.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HealthApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HealthApi apiInstance = new HealthApi(defaultClient);
    try {
      Object result = apiInstance.listHealth();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HealthApi#listHealth");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

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

<a name="listHealthAlarms"></a>
# **listHealthAlarms**
> Object listHealthAlarms(max, offset, sort, direction, phrase, name, acknowledged)

Retrieves Appliance Health Alarms

This endpoint retrieves all health alarms, which are Operation notifications from Cloud and other Service Integrations. These alarms are not generated by the appliance, but synced and displayed for visibility. By default only open alarms are returned. Open alarms are those that have not yet been acknowledged.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HealthApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HealthApi apiInstance = new HealthApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    Boolean acknowledged = false; // Boolean | True or False flag for Acknowledged items
    try {
      Object result = apiInstance.listHealthAlarms(max, offset, sort, direction, phrase, name, acknowledged);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HealthApi#listHealthAlarms");
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
 **acknowledged** | **Boolean**| True or False flag for Acknowledged items | [optional]

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

<a name="listHealthLogs"></a>
# **listHealthLogs**
> Object listHealthLogs(max, offset, sort, direction, phrase, name, acknowledged, startDate, endDate)

Retrieves Appliance Health Logs

This endpoint retrieves all health logs. These are the logs of the remote appliance itself. These logs show all ui activity and are useful for troubleshooting and auditing. Stack traces are filtered for Morpheus services. Complete stack traces can be found in &#x60;/var/log/morpheus/morpheus-ui/current&#x60;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HealthApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HealthApi apiInstance = new HealthApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    Boolean acknowledged = false; // Boolean | True or False flag for Acknowledged items
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    try {
      Object result = apiInstance.listHealthLogs(max, offset, sort, direction, phrase, name, acknowledged, startDate, endDate);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HealthApi#listHealthLogs");
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
 **acknowledged** | **Boolean**| True or False flag for Acknowledged items | [optional]
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional]
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional]

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

<a name="ping"></a>
# **ping**
> Ping ping()

Basic information about current Morpheus Installation

This endpoint can be used to check the remote appliance build version and some other basic information.  This is an unsecured endpoint and does not require authorization. However, build version will not be returned unless you are authenticated with a valid access token. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HealthApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HealthApi apiInstance = new HealthApi(defaultClient);
    try {
      Ping result = apiInstance.ping();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HealthApi#ping");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**Ping**](Ping.md)

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

