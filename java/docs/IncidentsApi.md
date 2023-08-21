# IncidentsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addIncident**](IncidentsApi.md#addIncident) | **POST** /api/monitoring/incidents | Create a New Incident
[**deleteIncidents**](IncidentsApi.md#deleteIncidents) | **DELETE** /api/monitoring/incidents/{id} | Close a Specific Incident
[**getIncidents**](IncidentsApi.md#getIncidents) | **GET** /api/monitoring/incidents/{id} | Get a Specific Incident
[**listIncidents**](IncidentsApi.md#listIncidents) | **GET** /api/monitoring/incidents | List All Incidents
[**updateIncidents**](IncidentsApi.md#updateIncidents) | **PUT** /api/monitoring/incidents/{id} | Update Incident
[**updateIncidentsReopen**](IncidentsApi.md#updateIncidentsReopen) | **GET** /api/monitoring/incidents/{id}/reopen | Reopen a Specific Incident
[**updateMuteAllIncidents**](IncidentsApi.md#updateMuteAllIncidents) | **PUT** /api/monitoring/incidents/mute-all | Mute All Incidents
[**updateMuteIncidents**](IncidentsApi.md#updateMuteIncidents) | **PUT** /api/monitoring/incidents/{id}/mute | Mute Incident


<a name="addIncident"></a>
# **addIncident**
> Object addIncident(inlineObject87)

Create a New Incident

Create a new incident.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IncidentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IncidentsApi apiInstance = new IncidentsApi(defaultClient);
    InlineObject87 inlineObject87 = new InlineObject87(); // InlineObject87 | 
    try {
      Object result = apiInstance.addIncident(inlineObject87);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IncidentsApi#addIncident");
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
 **inlineObject87** | [**InlineObject87**](InlineObject87.md)|  | [optional]

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
**200** | Incident Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteIncidents"></a>
# **deleteIncidents**
> Model200Success deleteIncidents(id)

Close a Specific Incident

Close an existing monitoring incident.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IncidentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IncidentsApi apiInstance = new IncidentsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteIncidents(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IncidentsApi#deleteIncidents");
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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getIncidents"></a>
# **getIncidents**
> InlineResponse20055 getIncidents(id)

Get a Specific Incident

Get details about a specific incident.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IncidentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IncidentsApi apiInstance = new IncidentsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20055 result = apiInstance.getIncidents(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IncidentsApi#getIncidents");
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

[**InlineResponse20055**](InlineResponse20055.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Incident Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listIncidents"></a>
# **listIncidents**
> Object listIncidents(max, offset, status, severity)

List All Incidents

Get a list of monitoring incidents.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IncidentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IncidentsApi apiInstance = new IncidentsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String status = "running"; // String | The instance status for filtering.
    String severity = "severity_example"; // String | Filter by severity
    try {
      Object result = apiInstance.listIncidents(max, offset, status, severity);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IncidentsApi#listIncidents");
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
 **status** | **String**| The instance status for filtering. | [optional]
 **severity** | **String**| Filter by severity | [optional] [enum: INFO, WARNING, CRITICAL]

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
**200** | Array of Incidents |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateIncidents"></a>
# **updateIncidents**
> Object updateIncidents(id, inlineObject88)

Update Incident

Update an existing incident.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IncidentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IncidentsApi apiInstance = new IncidentsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject88 inlineObject88 = new InlineObject88(); // InlineObject88 | 
    try {
      Object result = apiInstance.updateIncidents(id, inlineObject88);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IncidentsApi#updateIncidents");
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
 **inlineObject88** | [**InlineObject88**](InlineObject88.md)|  | [optional]

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
**200** | Incident Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateIncidentsReopen"></a>
# **updateIncidentsReopen**
> SuccessMessage updateIncidentsReopen(id)

Reopen a Specific Incident

Get details about a specific incident.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IncidentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IncidentsApi apiInstance = new IncidentsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      SuccessMessage result = apiInstance.updateIncidentsReopen(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IncidentsApi#updateIncidentsReopen");
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

[**SuccessMessage**](SuccessMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Incident Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateMuteAllIncidents"></a>
# **updateMuteAllIncidents**
> Object updateMuteAllIncidents(inlineObject90)

Mute All Incidents

Mute all existing incident.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IncidentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IncidentsApi apiInstance = new IncidentsApi(defaultClient);
    InlineObject90 inlineObject90 = new InlineObject90(); // InlineObject90 | 
    try {
      Object result = apiInstance.updateMuteAllIncidents(inlineObject90);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IncidentsApi#updateMuteAllIncidents");
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
 **inlineObject90** | [**InlineObject90**](InlineObject90.md)|  | [optional]

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
**200** | Mute All Incidents Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateMuteIncidents"></a>
# **updateMuteIncidents**
> Object updateMuteIncidents(id, inlineObject89)

Mute Incident

Mute an existing incident.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IncidentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IncidentsApi apiInstance = new IncidentsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject89 inlineObject89 = new InlineObject89(); // InlineObject89 | 
    try {
      Object result = apiInstance.updateMuteIncidents(id, inlineObject89);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IncidentsApi#updateMuteIncidents");
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
 **inlineObject89** | [**InlineObject89**](InlineObject89.md)|  | [optional]

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
**200** | Incident Mute Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

