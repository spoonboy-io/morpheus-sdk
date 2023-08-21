# ClusterLayoutsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addClusterLayoutClone**](ClusterLayoutsApi.md#addClusterLayoutClone) | **POST** /api/library/cluster-layouts/{id}/clone | Clone a Cluster Layout
[**addClusterLayouts**](ClusterLayoutsApi.md#addClusterLayouts) | **POST** /api/library/cluster-layouts | Create a Cluster Layout
[**deleteClusterLayout**](ClusterLayoutsApi.md#deleteClusterLayout) | **DELETE** /api/library/cluster-layouts/{id} | Delete a Cluster Layout
[**getClusterLayout**](ClusterLayoutsApi.md#getClusterLayout) | **GET** /api/library/cluster-layouts/{id} | Get a Specific Cluster Layout
[**listClusterLayouts**](ClusterLayoutsApi.md#listClusterLayouts) | **GET** /api/library/cluster-layouts | Get All Cluster Layouts
[**updateClusterLayout**](ClusterLayoutsApi.md#updateClusterLayout) | **PUT** /api/library/cluster-layouts/{id} | Update a Cluster Layout


<a name="addClusterLayoutClone"></a>
# **addClusterLayoutClone**
> SuccessId addClusterLayoutClone(id, name, description, computeVersion)

Clone a Cluster Layout

Use this command to clone a cluster layout.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClusterLayoutsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClusterLayoutsApi apiInstance = new ClusterLayoutsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String name = "New Name"; // String | Name of cluster layout. Defaults to Copy of <cloned layout name>
    String description = "New Description"; // String | Description of cluster layout. Defaults to cloned layout description
    String computeVersion = "2.0"; // String | Version of cluster layout. Defaults to cloned layout version
    try {
      SuccessId result = apiInstance.addClusterLayoutClone(id, name, description, computeVersion);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClusterLayoutsApi#addClusterLayoutClone");
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
 **name** | **String**| Name of cluster layout. Defaults to Copy of &lt;cloned layout name&gt; | [optional]
 **description** | **String**| Description of cluster layout. Defaults to cloned layout description | [optional]
 **computeVersion** | **String**| Version of cluster layout. Defaults to cloned layout version | [optional]

### Return type

[**SuccessId**](SuccessId.md)

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

<a name="addClusterLayouts"></a>
# **addClusterLayouts**
> SuccessId addClusterLayouts(inlineObject50)

Create a Cluster Layout

Use this command to create a cluster layout.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClusterLayoutsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClusterLayoutsApi apiInstance = new ClusterLayoutsApi(defaultClient);
    InlineObject50 inlineObject50 = new InlineObject50(); // InlineObject50 | 
    try {
      SuccessId result = apiInstance.addClusterLayouts(inlineObject50);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClusterLayoutsApi#addClusterLayouts");
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
 **inlineObject50** | [**InlineObject50**](InlineObject50.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

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

<a name="deleteClusterLayout"></a>
# **deleteClusterLayout**
> Model200Success deleteClusterLayout(id)

Delete a Cluster Layout

Will delete a cluster layout

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClusterLayoutsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClusterLayoutsApi apiInstance = new ClusterLayoutsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteClusterLayout(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClusterLayoutsApi#deleteClusterLayout");
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

<a name="getClusterLayout"></a>
# **getClusterLayout**
> InlineResponse20025 getClusterLayout(id)

Get a Specific Cluster Layout

This endpoint retrieves a specific cluster layout.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClusterLayoutsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClusterLayoutsApi apiInstance = new ClusterLayoutsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20025 result = apiInstance.getClusterLayout(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClusterLayoutsApi#getClusterLayout");
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

[**InlineResponse20025**](InlineResponse20025.md)

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

<a name="listClusterLayouts"></a>
# **listClusterLayouts**
> Object listClusterLayouts(max, offset, sort, direction, phrase, provisionType, labels, allLabels)

Get All Cluster Layouts

This endpoint retrieves all cluster layouts.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClusterLayoutsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClusterLayoutsApi apiInstance = new ClusterLayoutsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String provisionType = "provisionType_example"; // String | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings. 
    String labels = "labels_example"; // String | Filter by label(s), matches records that contain any of the specified labels
    String allLabels = "allLabels_example"; // String | Filter by label(s), matches records that contain all of the specified labels
    try {
      Object result = apiInstance.listClusterLayouts(max, offset, sort, direction, phrase, provisionType, labels, allLabels);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClusterLayoutsApi#listClusterLayouts");
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
 **provisionType** | **String**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional] [enum: AKS, Alibaba, Amazon, ARM, Azure, Azure SQL Server, Canonical MaaS, Cloud Foundry App, Cloud Foundry Diego, Cloud Foundry Service, CloudFormation, Digital Ocean, Docker, Docker Swarm, EKS, ESXi, External, Fusion, GKE, Google, Helm, HP Oneview, Huawei, Hyper-V, Hypervisor, IBM Cloud, Kubernetes, KVM, Lumen Edge, lxc, Manual, Manual Kubernetes, Nutanix, Open Telekom, OpenStack, Oracle Cloud, Oracle VM, RDS, SCVMM, SCVMM Hypervisor, SoftLayer, Spoon, Terraform, UCS, Unmanaged, UpCloud, Vagrant, VCD VAPP, VCloud Air, vCloud Director, VirtualBox, Virtustream, VMware, Workflow, Xen]
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional]

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

<a name="updateClusterLayout"></a>
# **updateClusterLayout**
> SuccessId updateClusterLayout(id, inlineObject51)

Update a Cluster Layout

Use this command to update an existing cluster layout.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ClusterLayoutsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ClusterLayoutsApi apiInstance = new ClusterLayoutsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject51 inlineObject51 = new InlineObject51(); // InlineObject51 | 
    try {
      SuccessId result = apiInstance.updateClusterLayout(id, inlineObject51);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ClusterLayoutsApi#updateClusterLayout");
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
 **inlineObject51** | [**InlineObject51**](InlineObject51.md)|  | [optional]

### Return type

[**SuccessId**](SuccessId.md)

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

