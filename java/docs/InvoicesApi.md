# InvoicesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getInvoiceLineItems**](InvoicesApi.md#getInvoiceLineItems) | **GET** /api/invoice-line-items/{id} | Get a Specific Invoice Line Item
[**getInvoices**](InvoicesApi.md#getInvoices) | **GET** /api/invoices/{id} | Get a Specific Invoice
[**listInvoiceLineItems**](InvoicesApi.md#listInvoiceLineItems) | **GET** /api/invoice-line-items | List All Invoice Line Items
[**listInvoices**](InvoicesApi.md#listInvoices) | **GET** /api/invoices | List All Invoices
[**updateInvoices**](InvoicesApi.md#updateInvoices) | **PUT** /api/invoices/{id} | Update Invoice Tags


<a name="getInvoiceLineItems"></a>
# **getInvoiceLineItems**
> Object getInvoiceLineItems(id)

Get a Specific Invoice Line Item

Get details about a specific invoice line item.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InvoicesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InvoicesApi apiInstance = new InvoicesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getInvoiceLineItems(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InvoicesApi#getInvoiceLineItems");
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
**200** | Invoice Line Item Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getInvoices"></a>
# **getInvoices**
> Object getInvoices(id)

Get a Specific Invoice

Get details about a specific invoice.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InvoicesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InvoicesApi apiInstance = new InvoicesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getInvoices(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InvoicesApi#getInvoices");
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
**200** | Invoice Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listInvoiceLineItems"></a>
# **listInvoiceLineItems**
> Object listInvoiceLineItems(max, offset, sort, direction, phrase, name, startDate, endDate, period, refType, refId, zoneId, siteId, instanceId, containerId, serverId, userId, projectId, active, accountId, includeTotals)

List All Invoice Line Items

This endpoint retrieves a list of all invoice line items for the specified parameters.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InvoicesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InvoicesApi apiInstance = new InvoicesApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    String period = "201901"; // String | Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM. 
    String refType = "ComputeSite"; // String | If specified will return an exact match on refType. 
    Long refId = 3L; // Long | If specified will return an exact match on refId
    Long zoneId = 3L; // Long | The Zone ID for Filtering
    Long siteId = 7L; // Long | The Site ID for Filtering
    Long instanceId = 94L; // Long | The Instance ID for Filtering
    Long containerId = 135L; // Long | The Container ID for Filtering
    Long serverId = 97L; // Long | The Server ID for Filtering
    Long userId = 6L; // Long | Filter by User ID
    Long projectId = 1L; // Long | The Project ID for Filtering
    Boolean active = false; // Boolean | True or False flag for Active
    Long accountId = 3L; // Long | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
    Boolean includeTotals = false; // Boolean | Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called `invoiceTotals`. 
    try {
      Object result = apiInstance.listInvoiceLineItems(max, offset, sort, direction, phrase, name, startDate, endDate, period, refType, refId, zoneId, siteId, instanceId, containerId, serverId, userId, projectId, active, accountId, includeTotals);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InvoicesApi#listInvoiceLineItems");
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
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional]
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional]
 **period** | **String**| Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.  | [optional]
 **refType** | **String**| If specified will return an exact match on refType.  | [optional]
 **refId** | **Long**| If specified will return an exact match on refId | [optional]
 **zoneId** | **Long**| The Zone ID for Filtering | [optional]
 **siteId** | **Long**| The Site ID for Filtering | [optional]
 **instanceId** | **Long**| The Instance ID for Filtering | [optional]
 **containerId** | **Long**| The Container ID for Filtering | [optional]
 **serverId** | **Long**| The Server ID for Filtering | [optional]
 **userId** | **Long**| Filter by User ID | [optional]
 **projectId** | **Long**| The Project ID for Filtering | [optional]
 **active** | **Boolean**| True or False flag for Active | [optional]
 **accountId** | **Long**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]
 **includeTotals** | **Boolean**| Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called &#x60;invoiceTotals&#x60;.  | [optional] [default to false]

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
**200** | Array of Invoice Line Items |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listInvoices"></a>
# **listInvoices**
> Object listInvoices(max, offset, sort, direction, phrase, name, startDate, endDate, period, refType, refId, refStatus, zoneId, siteId, instanceId, containerId, serverId, userId, projectId, active, accountId, includeLineItems, includeTotals, tags)

List All Invoices

This endpoint retrieves a list of invoices for the specified parameters.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InvoicesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InvoicesApi apiInstance = new InvoicesApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"refName\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    String period = "201901"; // String | Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM. 
    String refType = "ComputeSite"; // String | If specified will return an exact match on refType. 
    Long refId = 3L; // Long | If specified will return an exact match on refId
    String refStatus = "provisioned"; // String | If specified, will filter on the associated StorageVolume status. This is only applicable whn `refType=StorageVolume`. 
    Long zoneId = 3L; // Long | The Zone ID for Filtering
    Long siteId = 7L; // Long | The Site ID for Filtering
    Long instanceId = 94L; // Long | The Instance ID for Filtering
    Long containerId = 135L; // Long | The Container ID for Filtering
    Long serverId = 97L; // Long | The Server ID for Filtering
    Long userId = 6L; // Long | Filter by User ID
    Long projectId = 1L; // Long | The Project ID for Filtering
    Boolean active = false; // Boolean | True or False flag for Active
    Long accountId = 3L; // Long | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
    Boolean includeLineItems = false; // Boolean | Pass true to include the list of `lineItems` for each invoice. Only `lineItemCount` is returned by default. 
    Boolean includeTotals = false; // Boolean | Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called `invoiceTotals`. 
    String tags = "tags.env=qa&tags.env=test"; // String | Filter by tags (metadata). This allows filtering by a tag name and value(s) 
    try {
      Object result = apiInstance.listInvoices(max, offset, sort, direction, phrase, name, startDate, endDate, period, refType, refId, refStatus, zoneId, siteId, instanceId, containerId, serverId, userId, projectId, active, accountId, includeLineItems, includeTotals, tags);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InvoicesApi#listInvoices");
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
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;refName&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional]
 **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional]
 **period** | **String**| Only return records for period that matches with the specified value. This is an alternative to using startDate/endDate. Format is YYYY or YYYYMM.  | [optional]
 **refType** | **String**| If specified will return an exact match on refType.  | [optional]
 **refId** | **Long**| If specified will return an exact match on refId | [optional]
 **refStatus** | **String**| If specified, will filter on the associated StorageVolume status. This is only applicable whn &#x60;refType&#x3D;StorageVolume&#x60;.  | [optional]
 **zoneId** | **Long**| The Zone ID for Filtering | [optional]
 **siteId** | **Long**| The Site ID for Filtering | [optional]
 **instanceId** | **Long**| The Instance ID for Filtering | [optional]
 **containerId** | **Long**| The Container ID for Filtering | [optional]
 **serverId** | **Long**| The Server ID for Filtering | [optional]
 **userId** | **Long**| Filter by User ID | [optional]
 **projectId** | **Long**| The Project ID for Filtering | [optional]
 **active** | **Boolean**| True or False flag for Active | [optional]
 **accountId** | **Long**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]
 **includeLineItems** | **Boolean**| Pass true to include the list of &#x60;lineItems&#x60; for each invoice. Only &#x60;lineItemCount&#x60; is returned by default.  | [optional] [default to false]
 **includeTotals** | **Boolean**| Pass true to include the summed totals (cost/price values) for all the invoices found in the query. The returned property is called &#x60;invoiceTotals&#x60;.  | [optional] [default to false]
 **tags** | **String**| Filter by tags (metadata). This allows filtering by a tag name and value(s)  | [optional]

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
**200** | Array of Invoices |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateInvoices"></a>
# **updateInvoices**
> Object updateInvoices(id, inlineObject102)

Update Invoice Tags

Update, Add, or Remove invoice tag(s).

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.InvoicesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    InvoicesApi apiInstance = new InvoicesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject102 inlineObject102 = new InlineObject102(); // InlineObject102 | 
    try {
      Object result = apiInstance.updateInvoices(id, inlineObject102);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling InvoicesApi#updateInvoices");
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
 **inlineObject102** | [**InlineObject102**](InlineObject102.md)|  | [optional]

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
**200** | Invoice Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

