# BillingApi

All URIs are relative to *https://CHANGEME*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getBillingAccount**](BillingApi.md#getBillingAccount) | **GET** /api/billing/account/{id} | This endpoint will retrieve a specific account by id if the user has permission to access it |
| [**getBillingInstancesIdentifier**](BillingApi.md#getBillingInstancesIdentifier) | **GET** /api/billing/instances/{identifier} | Retrieves billing information for an instance in the requestor&#39;s account. Use instanceUUID whenever possible. |
| [**getBillingServersIdentifier**](BillingApi.md#getBillingServersIdentifier) | **GET** /api/billing/servers/{identifier} | Retrieves billing information for a specific server (container host) in the requestor&#39;s account. Use refUUID whenever possible. |
| [**getBillingZoneIdentifier**](BillingApi.md#getBillingZoneIdentifier) | **GET** /api/billing/zones/{identifier} | Retrieves billing information for a specific zone in the requestor&#39;s account. Use zoneUUID whenever possible. |
| [**listBillingAccount**](BillingApi.md#listBillingAccount) | **GET** /api/billing/account | Retrieves billing information for the requesting user&#39;s account. |
| [**listBillingInstances**](BillingApi.md#listBillingInstances) | **GET** /api/billing/instances | Retrieves billing information for all instances on the requestor&#39;s account. |
| [**listBillingServers**](BillingApi.md#listBillingServers) | **GET** /api/billing/servers | Retrieves billing information for all servers (container hosts) on the requestor&#39;s account. |
| [**listBillingZone**](BillingApi.md#listBillingZone) | **GET** /api/billing/zones | Retrieves billing information for all zones on the requestor&#39;s account. |


<a id="getBillingAccount"></a>
# **getBillingAccount**
> ListBillingAccount200Response getBillingAccount(id, startDate, endDate, includeUsages, maxUsages, offsetUsages, includeComputeServers, includeInstances, includeDiscoveredServers, includeLoadBalancers, includeVirtualImages, includeSnapshots)

This endpoint will retrieve a specific account by id if the user has permission to access it

Will retrieve billing information for a specific tenant, if it is the current account or a sub account of the requesting user&#39;s account. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BillingApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BillingApi apiInstance = new BillingApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    Boolean includeUsages = true; // Boolean | Optional ability to suppress the usage records
    Long maxUsages = 56L; // Long | Optional ability to limit the usages returned
    Long offsetUsages = 56L; // Long | Optional ability to offset the usages returned, for use with maxUsages to paginate
    Boolean includeComputeServers = true; // Boolean | Optional ability to exclude compute servers
    Boolean includeInstances = true; // Boolean | Optional ability to exclude instances
    Boolean includeDiscoveredServers = true; // Boolean | Optional ability to exclude discovered servers
    Boolean includeLoadBalancers = true; // Boolean | Optional ability to exclude load balancers
    Boolean includeVirtualImages = true; // Boolean | Optional ability to exclude virtual images
    Boolean includeSnapshots = true; // Boolean | Optional ability to exclude snapshots
    try {
      ListBillingAccount200Response result = apiInstance.getBillingAccount(id, startDate, endDate, includeUsages, maxUsages, offsetUsages, includeComputeServers, includeInstances, includeDiscoveredServers, includeLoadBalancers, includeVirtualImages, includeSnapshots);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BillingApi#getBillingAccount");
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
| **id** | **Long**| Morpheus ID of the Object being referenced | |
| **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] |
| **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] |
| **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true] |
| **maxUsages** | **Long**| Optional ability to limit the usages returned | [optional] |
| **offsetUsages** | **Long**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] |
| **includeComputeServers** | **Boolean**| Optional ability to exclude compute servers | [optional] [default to true] |
| **includeInstances** | **Boolean**| Optional ability to exclude instances | [optional] [default to true] |
| **includeDiscoveredServers** | **Boolean**| Optional ability to exclude discovered servers | [optional] [default to true] |
| **includeLoadBalancers** | **Boolean**| Optional ability to exclude load balancers | [optional] [default to true] |
| **includeVirtualImages** | **Boolean**| Optional ability to exclude virtual images | [optional] [default to true] |
| **includeSnapshots** | **Boolean**| Optional ability to exclude snapshots | [optional] [default to true] |

### Return type

[**ListBillingAccount200Response**](ListBillingAccount200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Array of Billing |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="getBillingInstancesIdentifier"></a>
# **getBillingInstancesIdentifier**
> GetBillingInstancesIdentifier200Response getBillingInstancesIdentifier(identifier, startDate, endDate, includeUsages, maxUsages, offsetUsages, includeTenants, accountId)

Retrieves billing information for an instance in the requestor&#39;s account. Use instanceUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BillingApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BillingApi apiInstance = new BillingApi(defaultClient);
    String identifier = "identifier_example"; // String | Morpheus UUID or ID of the Object being created or referenced
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    Boolean includeUsages = true; // Boolean | Optional ability to suppress the usage records
    Long maxUsages = 56L; // Long | Optional ability to limit the usages returned
    Long offsetUsages = 56L; // Long | Optional ability to offset the usages returned, for use with maxUsages to paginate
    Boolean includeTenants = false; // Boolean | Optional ability to include all subtenant billing information when calling from a master tenant user
    Long accountId = 3L; // Long | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
    try {
      GetBillingInstancesIdentifier200Response result = apiInstance.getBillingInstancesIdentifier(identifier, startDate, endDate, includeUsages, maxUsages, offsetUsages, includeTenants, accountId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BillingApi#getBillingInstancesIdentifier");
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
| **identifier** | **String**| Morpheus UUID or ID of the Object being created or referenced | |
| **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] |
| **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] |
| **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true] |
| **maxUsages** | **Long**| Optional ability to limit the usages returned | [optional] |
| **offsetUsages** | **Long**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] |
| **includeTenants** | **Boolean**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to false] |
| **accountId** | **Long**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] |

### Return type

[**GetBillingInstancesIdentifier200Response**](GetBillingInstancesIdentifier200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Array of Billing |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="getBillingServersIdentifier"></a>
# **getBillingServersIdentifier**
> GetBillingServersIdentifier200Response getBillingServersIdentifier(identifier, startDate, endDate, includeUsages, maxUsages, offsetUsages, includeTenants, accountId)

Retrieves billing information for a specific server (container host) in the requestor&#39;s account. Use refUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BillingApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BillingApi apiInstance = new BillingApi(defaultClient);
    String identifier = "identifier_example"; // String | Morpheus UUID or ID of the Object being created or referenced
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    Boolean includeUsages = true; // Boolean | Optional ability to suppress the usage records
    Long maxUsages = 56L; // Long | Optional ability to limit the usages returned
    Long offsetUsages = 56L; // Long | Optional ability to offset the usages returned, for use with maxUsages to paginate
    Boolean includeTenants = false; // Boolean | Optional ability to include all subtenant billing information when calling from a master tenant user
    Long accountId = 3L; // Long | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
    try {
      GetBillingServersIdentifier200Response result = apiInstance.getBillingServersIdentifier(identifier, startDate, endDate, includeUsages, maxUsages, offsetUsages, includeTenants, accountId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BillingApi#getBillingServersIdentifier");
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
| **identifier** | **String**| Morpheus UUID or ID of the Object being created or referenced | |
| **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] |
| **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] |
| **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true] |
| **maxUsages** | **Long**| Optional ability to limit the usages returned | [optional] |
| **offsetUsages** | **Long**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] |
| **includeTenants** | **Boolean**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to false] |
| **accountId** | **Long**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] |

### Return type

[**GetBillingServersIdentifier200Response**](GetBillingServersIdentifier200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Array of Billing |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="getBillingZoneIdentifier"></a>
# **getBillingZoneIdentifier**
> GetBillingZoneIdentifier200Response getBillingZoneIdentifier(identifier, startDate, endDate, includeUsages, maxUsages, offsetUsages, includeComputeServers, includeInstances, includeDiscoveredServers, includeLoadBalancers, includeVirtualImages, includeSnapshots)

Retrieves billing information for a specific zone in the requestor&#39;s account. Use zoneUUID whenever possible.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BillingApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BillingApi apiInstance = new BillingApi(defaultClient);
    String identifier = "identifier_example"; // String | Morpheus UUID or ID of the Object being created or referenced
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    Boolean includeUsages = true; // Boolean | Optional ability to suppress the usage records
    Long maxUsages = 56L; // Long | Optional ability to limit the usages returned
    Long offsetUsages = 56L; // Long | Optional ability to offset the usages returned, for use with maxUsages to paginate
    Boolean includeComputeServers = true; // Boolean | Optional ability to exclude compute servers
    Boolean includeInstances = true; // Boolean | Optional ability to exclude instances
    Boolean includeDiscoveredServers = true; // Boolean | Optional ability to exclude discovered servers
    Boolean includeLoadBalancers = true; // Boolean | Optional ability to exclude load balancers
    Boolean includeVirtualImages = true; // Boolean | Optional ability to exclude virtual images
    Boolean includeSnapshots = true; // Boolean | Optional ability to exclude snapshots
    try {
      GetBillingZoneIdentifier200Response result = apiInstance.getBillingZoneIdentifier(identifier, startDate, endDate, includeUsages, maxUsages, offsetUsages, includeComputeServers, includeInstances, includeDiscoveredServers, includeLoadBalancers, includeVirtualImages, includeSnapshots);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BillingApi#getBillingZoneIdentifier");
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
| **identifier** | **String**| Morpheus UUID or ID of the Object being created or referenced | |
| **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] |
| **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] |
| **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true] |
| **maxUsages** | **Long**| Optional ability to limit the usages returned | [optional] |
| **offsetUsages** | **Long**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] |
| **includeComputeServers** | **Boolean**| Optional ability to exclude compute servers | [optional] [default to true] |
| **includeInstances** | **Boolean**| Optional ability to exclude instances | [optional] [default to true] |
| **includeDiscoveredServers** | **Boolean**| Optional ability to exclude discovered servers | [optional] [default to true] |
| **includeLoadBalancers** | **Boolean**| Optional ability to exclude load balancers | [optional] [default to true] |
| **includeVirtualImages** | **Boolean**| Optional ability to exclude virtual images | [optional] [default to true] |
| **includeSnapshots** | **Boolean**| Optional ability to exclude snapshots | [optional] [default to true] |

### Return type

[**GetBillingZoneIdentifier200Response**](GetBillingZoneIdentifier200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Array of Billing |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="listBillingAccount"></a>
# **listBillingAccount**
> ListBillingAccount200Response listBillingAccount(startDate, endDate, includeUsages, maxUsages, offsetUsages, includeComputeServers, includeInstances, includeDiscoveredServers, includeLoadBalancers, includeVirtualImages, includeSnapshots)

Retrieves billing information for the requesting user&#39;s account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BillingApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BillingApi apiInstance = new BillingApi(defaultClient);
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    Boolean includeUsages = true; // Boolean | Optional ability to suppress the usage records
    Long maxUsages = 56L; // Long | Optional ability to limit the usages returned
    Long offsetUsages = 56L; // Long | Optional ability to offset the usages returned, for use with maxUsages to paginate
    Boolean includeComputeServers = true; // Boolean | Optional ability to exclude compute servers
    Boolean includeInstances = true; // Boolean | Optional ability to exclude instances
    Boolean includeDiscoveredServers = true; // Boolean | Optional ability to exclude discovered servers
    Boolean includeLoadBalancers = true; // Boolean | Optional ability to exclude load balancers
    Boolean includeVirtualImages = true; // Boolean | Optional ability to exclude virtual images
    Boolean includeSnapshots = true; // Boolean | Optional ability to exclude snapshots
    try {
      ListBillingAccount200Response result = apiInstance.listBillingAccount(startDate, endDate, includeUsages, maxUsages, offsetUsages, includeComputeServers, includeInstances, includeDiscoveredServers, includeLoadBalancers, includeVirtualImages, includeSnapshots);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BillingApi#listBillingAccount");
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
| **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] |
| **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] |
| **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true] |
| **maxUsages** | **Long**| Optional ability to limit the usages returned | [optional] |
| **offsetUsages** | **Long**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] |
| **includeComputeServers** | **Boolean**| Optional ability to exclude compute servers | [optional] [default to true] |
| **includeInstances** | **Boolean**| Optional ability to exclude instances | [optional] [default to true] |
| **includeDiscoveredServers** | **Boolean**| Optional ability to exclude discovered servers | [optional] [default to true] |
| **includeLoadBalancers** | **Boolean**| Optional ability to exclude load balancers | [optional] [default to true] |
| **includeVirtualImages** | **Boolean**| Optional ability to exclude virtual images | [optional] [default to true] |
| **includeSnapshots** | **Boolean**| Optional ability to exclude snapshots | [optional] [default to true] |

### Return type

[**ListBillingAccount200Response**](ListBillingAccount200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Array of Billing |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="listBillingInstances"></a>
# **listBillingInstances**
> ListBillingInstances200Response listBillingInstances(startDate, endDate, includeUsages, maxUsages, offsetUsages, includeTenants, accountId)

Retrieves billing information for all instances on the requestor&#39;s account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BillingApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BillingApi apiInstance = new BillingApi(defaultClient);
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    Boolean includeUsages = true; // Boolean | Optional ability to suppress the usage records
    Long maxUsages = 56L; // Long | Optional ability to limit the usages returned
    Long offsetUsages = 56L; // Long | Optional ability to offset the usages returned, for use with maxUsages to paginate
    Boolean includeTenants = false; // Boolean | Optional ability to include all subtenant billing information when calling from a master tenant user
    Long accountId = 3L; // Long | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
    try {
      ListBillingInstances200Response result = apiInstance.listBillingInstances(startDate, endDate, includeUsages, maxUsages, offsetUsages, includeTenants, accountId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BillingApi#listBillingInstances");
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
| **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] |
| **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] |
| **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true] |
| **maxUsages** | **Long**| Optional ability to limit the usages returned | [optional] |
| **offsetUsages** | **Long**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] |
| **includeTenants** | **Boolean**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to false] |
| **accountId** | **Long**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] |

### Return type

[**ListBillingInstances200Response**](ListBillingInstances200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Array of Billing |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="listBillingServers"></a>
# **listBillingServers**
> ListBillingServers200Response listBillingServers(startDate, endDate, includeUsages, maxUsages, offsetUsages, includeTenants, accountId)

Retrieves billing information for all servers (container hosts) on the requestor&#39;s account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BillingApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BillingApi apiInstance = new BillingApi(defaultClient);
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    Boolean includeUsages = true; // Boolean | Optional ability to suppress the usage records
    Long maxUsages = 56L; // Long | Optional ability to limit the usages returned
    Long offsetUsages = 56L; // Long | Optional ability to offset the usages returned, for use with maxUsages to paginate
    Boolean includeTenants = false; // Boolean | Optional ability to include all subtenant billing information when calling from a master tenant user
    Long accountId = 3L; // Long | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
    try {
      ListBillingServers200Response result = apiInstance.listBillingServers(startDate, endDate, includeUsages, maxUsages, offsetUsages, includeTenants, accountId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BillingApi#listBillingServers");
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
| **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] |
| **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] |
| **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true] |
| **maxUsages** | **Long**| Optional ability to limit the usages returned | [optional] |
| **offsetUsages** | **Long**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] |
| **includeTenants** | **Boolean**| Optional ability to include all subtenant billing information when calling from a master tenant user | [optional] [default to false] |
| **accountId** | **Long**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] |

### Return type

[**ListBillingServers200Response**](ListBillingServers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Array of Billing |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="listBillingZone"></a>
# **listBillingZone**
> ListBillingZone200Response listBillingZone(startDate, endDate, includeUsages, maxUsages, offsetUsages, includeComputeServers, includeInstances, includeDiscoveredServers, includeLoadBalancers, includeVirtualImages, includeSnapshots)

Retrieves billing information for all zones on the requestor&#39;s account.

Provides API interfaces for viewing billing usage information by tenant, zone, instance or server. By default, usage returned is from the beginning of the current month until now. The date range is parameterized but the end date cannot exceed the current date. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BillingApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BillingApi apiInstance = new BillingApi(defaultClient);
    String startDate = "2019-01-01"; // String | Filter by startDate greater than or equal to a specified date
    String endDate = "2020-01-01"; // String | Filter by endDate less than or equal to a specified date
    Boolean includeUsages = true; // Boolean | Optional ability to suppress the usage records
    Long maxUsages = 56L; // Long | Optional ability to limit the usages returned
    Long offsetUsages = 56L; // Long | Optional ability to offset the usages returned, for use with maxUsages to paginate
    Boolean includeComputeServers = true; // Boolean | Optional ability to exclude compute servers
    Boolean includeInstances = true; // Boolean | Optional ability to exclude instances
    Boolean includeDiscoveredServers = true; // Boolean | Optional ability to exclude discovered servers
    Boolean includeLoadBalancers = true; // Boolean | Optional ability to exclude load balancers
    Boolean includeVirtualImages = true; // Boolean | Optional ability to exclude virtual images
    Boolean includeSnapshots = true; // Boolean | Optional ability to exclude snapshots
    try {
      ListBillingZone200Response result = apiInstance.listBillingZone(startDate, endDate, includeUsages, maxUsages, offsetUsages, includeComputeServers, includeInstances, includeDiscoveredServers, includeLoadBalancers, includeVirtualImages, includeSnapshots);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BillingApi#listBillingZone");
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
| **startDate** | **String**| Filter by startDate greater than or equal to a specified date | [optional] |
| **endDate** | **String**| Filter by endDate less than or equal to a specified date | [optional] |
| **includeUsages** | **Boolean**| Optional ability to suppress the usage records | [optional] [default to true] |
| **maxUsages** | **Long**| Optional ability to limit the usages returned | [optional] |
| **offsetUsages** | **Long**| Optional ability to offset the usages returned, for use with maxUsages to paginate | [optional] |
| **includeComputeServers** | **Boolean**| Optional ability to exclude compute servers | [optional] [default to true] |
| **includeInstances** | **Boolean**| Optional ability to exclude instances | [optional] [default to true] |
| **includeDiscoveredServers** | **Boolean**| Optional ability to exclude discovered servers | [optional] [default to true] |
| **includeLoadBalancers** | **Boolean**| Optional ability to exclude load balancers | [optional] [default to true] |
| **includeVirtualImages** | **Boolean**| Optional ability to exclude virtual images | [optional] [default to true] |
| **includeSnapshots** | **Boolean**| Optional ability to exclude snapshots | [optional] [default to true] |

### Return type

[**ListBillingZone200Response**](ListBillingZone200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Array of Billing |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

