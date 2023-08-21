# HostsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getHost**](HostsApi.md#getHost) | **GET** /api/servers/{id} | Get a Specific Host
[**getHostSnpshots**](HostsApi.md#getHostSnpshots) | **GET** /api/servers/{id}/snapshots | Get list of snapshots for a Host
[**getHostType**](HostsApi.md#getHostType) | **GET** /api/server-types/{id} | Get a Specific Host Type
[**getWikiServer**](HostsApi.md#getWikiServer) | **GET** /api/servers/{id}/wiki | Retrieves a Server Wiki Page
[**listHostTypes**](HostsApi.md#listHostTypes) | **GET** /api/server-types | Host Types
[**listHosts**](HostsApi.md#listHosts) | **GET** /api/servers | Get All Hosts
[**listServerServicePlans**](HostsApi.md#listServerServicePlans) | **GET** /api/servers/service-plans | Get Available Service Plans for a Host
[**removeHost**](HostsApi.md#removeHost) | **DELETE** /api/servers/{id} | Delete a Host
[**restartHost**](HostsApi.md#restartHost) | **PUT** /api/servers/{id}/restart | Restart a Host
[**startHost**](HostsApi.md#startHost) | **PUT** /api/servers/{id}/start | Start a Host
[**stopHost**](HostsApi.md#stopHost) | **PUT** /api/servers/{id}/stop | Stop a Host
[**updateHost**](HostsApi.md#updateHost) | **PUT** /api/servers/{id} | Updating a Host
[**updateHostAssignTenant**](HostsApi.md#updateHostAssignTenant) | **PUT** /api/servers/{id}/assign-account | Assign To Tenant
[**updateHostCloud**](HostsApi.md#updateHostCloud) | **PUT** /api/servers/change-cloud | Change Server Cloud
[**updateHostExecuteWorkflow**](HostsApi.md#updateHostExecuteWorkflow) | **PUT** /api/servers/{id}/workflow | Run Workflow on a Host
[**updateHostInstallAgent**](HostsApi.md#updateHostInstallAgent) | **PUT** /api/servers/{id}/install-agent | Install Agent
[**updateHostManaged**](HostsApi.md#updateHostManaged) | **PUT** /api/servers/{id}/make-managed | Convert To Managed
[**updateHostResize**](HostsApi.md#updateHostResize) | **PUT** /api/servers/{id}/resize | Resize a Host
[**updateHostUpgradeAgent**](HostsApi.md#updateHostUpgradeAgent) | **PUT** /api/servers/{id}/upgrade | Upgrade Agent
[**updateServerNetworkInterface**](HostsApi.md#updateServerNetworkInterface) | **PUT** /api/servers/{id}/networkInterfaces/{networkInterfaceId} | Updating a label for a Server&#39;s Network
[**updateWikiServer**](HostsApi.md#updateWikiServer) | **PUT** /api/servers/{id}/wiki | Update a Server Wiki Page


<a name="getHost"></a>
# **getHost**
> InlineResponse200137 getHost(id)

Get a Specific Host

This endpoint retrieves a specific host.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200137 result = apiInstance.getHost(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#getHost");
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

[**InlineResponse200137**](InlineResponse200137.md)

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

<a name="getHostSnpshots"></a>
# **getHostSnpshots**
> InlineResponse200138 getHostSnpshots(id)

Get list of snapshots for a Host

Get list of snapshots for a Host

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200138 result = apiInstance.getHostSnpshots(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#getHostSnpshots");
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

[**InlineResponse200138**](InlineResponse200138.md)

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

<a name="getHostType"></a>
# **getHostType**
> InlineResponse20050 getHostType(id)

Get a Specific Host Type

This endpoint will retrieve a specific host type by id

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20050 result = apiInstance.getHostType(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#getHostType");
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

[**InlineResponse20050**](InlineResponse20050.md)

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

<a name="getWikiServer"></a>
# **getWikiServer**
> InlineResponse200168 getWikiServer(id)

Retrieves a Server Wiki Page

This endpoint retrieves a server Wiki page. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200168 result = apiInstance.getWikiServer(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#getWikiServer");
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

[**InlineResponse200168**](InlineResponse200168.md)

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

<a name="listHostTypes"></a>
# **listHostTypes**
> Object listHostTypes(max, offset, sort, direction, name, code, phrase, provisionType, zoneType, creatable)

Host Types

Fetch a paginated list of available host types. This returns the configuration options for each type.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String code = "azr"; // String | If specified will return an exact match on code
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String provisionType = "provisionType_example"; // String | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings. 
    String zoneType = "zoneType_example"; // String | Filter by Cloud Type code.
    Boolean creatable = true; // Boolean | Filter by creatable flag. This is whether or not it can be provisioned.
    try {
      Object result = apiInstance.listHostTypes(max, offset, sort, direction, name, code, phrase, provisionType, zoneType, creatable);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#listHostTypes");
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **code** | **String**| If specified will return an exact match on code | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **provisionType** | **String**| Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | [optional] [enum: AKS, Alibaba, Amazon, ARM, Azure, Azure SQL Server, Canonical MaaS, Cloud Foundry App, Cloud Foundry Diego, Cloud Foundry Service, CloudFormation, Digital Ocean, Docker, Docker Swarm, EKS, ESXi, External, Fusion, GKE, Google, Helm, HP Oneview, Huawei, Hyper-V, Hypervisor, IBM Cloud, Kubernetes, KVM, Lumen Edge, lxc, Manual, Manual Kubernetes, Nutanix, Open Telekom, OpenStack, Oracle Cloud, Oracle VM, RDS, SCVMM, SCVMM Hypervisor, SoftLayer, Spoon, Terraform, UCS, Unmanaged, UpCloud, Vagrant, VCD VAPP, VCloud Air, vCloud Director, VirtualBox, Virtustream, VMware, Workflow, Xen]
 **zoneType** | **String**| Filter by Cloud Type code. | [optional] [enum: alibaba, amazon, azure, azurestack, dell, digitalocean, googlecloud, oneview, huawei, hyperv, bluemix, bluemixCloudFoundry, centurylink-edge, macstadium, standard, nutanix, opentelekom, openstack, oraclecloud, oraclevm, platform9, scvmm, softlayer, ucs, upcloud, vcd, vmwareCloudAws, esxi, fusion, vmware, xenserver]
 **creatable** | **Boolean**| Filter by creatable flag. This is whether or not it can be provisioned. | [optional]

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

<a name="listHosts"></a>
# **listHosts**
> Object listHosts(name, phrase, zoneId, siteId, clusterId, managed, serverType, powerState, ip, vm, vmHypervisor, bareMetalHost, status, agentInstalled, max, offset, lastUpdated, createdBy, labels, allLabels, tags, metadata, uuid, externalId, internalId, externalUniquelId)

Get All Hosts

This endpoint retrieves a paginated list of hosts.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    Long zoneId = 3L; // Long | The Zone ID for Filtering
    Long siteId = 7L; // Long | The Site ID for Filtering
    Long clusterId = 135L; // Long | The Cluster ID(s) for filtering. Accepts multiple values.
    Boolean managed = false; // Boolean | Filter by managed (true) or unmanaged (false)
    String serverType = "vmwareHypervisor"; // String | Filter by server type code
    String powerState = "off"; // String | Filter by power status
    String ip = "192.168.1.45"; // String | Filter by IP address
    Boolean vm = false; // Boolean | Filter to show only Virtual Machines (true)
    Boolean vmHypervisor = false; // Boolean | Filter to show only VM Hypervisors (true)
    Boolean bareMetalHost = false; // Boolean | Filter to show only Baremetal Servers
    String status = "status_example"; // String | Filter by status
    Boolean agentInstalled = true; // Boolean | Filter by agent installed (true)
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    Long createdBy = 63L; // Long | The User ID for Filtering
    String labels = "labels_example"; // String | Filter by label(s), matches records that contain any of the specified labels
    String allLabels = "allLabels_example"; // String | Filter by label(s), matches records that contain all of the specified labels
    String tags = "tags.env=qa&tags.env=test"; // String | Filter by tags (metadata). This allows filtering by a tag name and value(s) 
    String metadata = "tags.env=qa&tags.env=test"; // String | Alias for tags
    String uuid = "uuid_example"; // String | Filter by UUID
    String externalId = "externalId_example"; // String | Filter by External ID
    String internalId = "internalId_example"; // String | Filter by Internal ID
    String externalUniquelId = "externalUniquelId_example"; // String | Filter by External Unique ID
    try {
      Object result = apiInstance.listHosts(name, phrase, zoneId, siteId, clusterId, managed, serverType, powerState, ip, vm, vmHypervisor, bareMetalHost, status, agentInstalled, max, offset, lastUpdated, createdBy, labels, allLabels, tags, metadata, uuid, externalId, internalId, externalUniquelId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#listHosts");
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
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **zoneId** | **Long**| The Zone ID for Filtering | [optional]
 **siteId** | **Long**| The Site ID for Filtering | [optional]
 **clusterId** | **Long**| The Cluster ID(s) for filtering. Accepts multiple values. | [optional]
 **managed** | **Boolean**| Filter by managed (true) or unmanaged (false) | [optional]
 **serverType** | **String**| Filter by server type code | [optional]
 **powerState** | **String**| Filter by power status | [optional]
 **ip** | **String**| Filter by IP address | [optional]
 **vm** | **Boolean**| Filter to show only Virtual Machines (true) | [optional]
 **vmHypervisor** | **Boolean**| Filter to show only VM Hypervisors (true) | [optional]
 **bareMetalHost** | **Boolean**| Filter to show only Baremetal Servers | [optional]
 **status** | **String**| Filter by status | [optional]
 **agentInstalled** | **Boolean**| Filter by agent installed (true) | [optional]
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **lastUpdated** | **OffsetDateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **createdBy** | **Long**| The User ID for Filtering | [optional]
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional]
 **tags** | **String**| Filter by tags (metadata). This allows filtering by a tag name and value(s)  | [optional]
 **metadata** | **String**| Alias for tags | [optional]
 **uuid** | **String**| Filter by UUID | [optional]
 **externalId** | **String**| Filter by External ID | [optional]
 **internalId** | **String**| Filter by Internal ID | [optional]
 **externalUniquelId** | **String**| Filter by External Unique ID | [optional]

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

<a name="listServerServicePlans"></a>
# **listServerServicePlans**
> InlineResponse200141 listServerServicePlans(zoneId, serverTypeId, siteId)

Get Available Service Plans for a Host

This endpoint retrieves all the Service Plans available for the specified cloud and host type. It may be used to get the list of available plans when creating a new host or resizing an existing host.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long zoneId = 3L; // Long | The Zone ID for Filtering
    Long serverTypeId = 5L; // Long | The ID of the Host Type
    Long siteId = 7L; // Long | The Site ID for Filtering
    try {
      InlineResponse200141 result = apiInstance.listServerServicePlans(zoneId, serverTypeId, siteId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#listServerServicePlans");
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
 **zoneId** | **Long**| The Zone ID for Filtering |
 **serverTypeId** | **Long**| The ID of the Host Type | [optional]
 **siteId** | **Long**| The Site ID for Filtering | [optional]

### Return type

[**InlineResponse200141**](InlineResponse200141.md)

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

<a name="removeHost"></a>
# **removeHost**
> Model200Success removeHost(id, removeResources, removeInstances, preserveVolumes, releaseFloatingIps, releaseEIPs, force)

Delete a Host

Will delete a host asynchronously.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String removeResources = "\"on\""; // String | Remove Resources
    String removeInstances = "\"off\""; // String | Remove Instances
    String preserveVolumes = "\"off\""; // String | Preserve Volumes
    String releaseFloatingIps = "\"on\""; // String | Release Floating IPs
    String releaseEIPs = "\"on\""; // String | Alias for releaseFloatingIps
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.removeHost(id, removeResources, removeInstances, preserveVolumes, releaseFloatingIps, releaseEIPs, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#removeHost");
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
 **removeResources** | **String**| Remove Resources | [optional] [default to &quot;on&quot;]
 **removeInstances** | **String**| Remove Instances | [optional] [default to &quot;off&quot;]
 **preserveVolumes** | **String**| Preserve Volumes | [optional] [default to &quot;off&quot;]
 **releaseFloatingIps** | **String**| Release Floating IPs | [optional] [default to &quot;on&quot;]
 **releaseEIPs** | **String**| Alias for releaseFloatingIps | [optional] [default to &quot;on&quot;]
 **force** | **String**| Force Delete | [optional] [default to &quot;off&quot;]

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

<a name="restartHost"></a>
# **restartHost**
> Object restartHost(id)

Restart a Host

This will restart a host.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.restartHost(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#restartHost");
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

<a name="startHost"></a>
# **startHost**
> Model200Success startHost(id)

Start a Host

This will start a host.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.startHost(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#startHost");
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

<a name="stopHost"></a>
# **stopHost**
> Model200Success stopHost(id)

Stop a Host

This will stop a host.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.stopHost(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#stopHost");
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

<a name="updateHost"></a>
# **updateHost**
> InlineResponse200137 updateHost(id, inlineObject220)

Updating a Host

Updating a Host

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject220 inlineObject220 = new InlineObject220(); // InlineObject220 | 
    try {
      InlineResponse200137 result = apiInstance.updateHost(id, inlineObject220);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#updateHost");
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
 **inlineObject220** | [**InlineObject220**](InlineObject220.md)|  | [optional]

### Return type

[**InlineResponse200137**](InlineResponse200137.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateHostAssignTenant"></a>
# **updateHostAssignTenant**
> Object updateHostAssignTenant(id, accountId, inlineObject221)

Assign To Tenant

This will change the ownership of the host to the specified Tenant account. This is only available to Master Tenant users.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Long accountId = 3L; // Long | ID of the Tenant
    InlineObject221 inlineObject221 = new InlineObject221(); // InlineObject221 | 
    try {
      Object result = apiInstance.updateHostAssignTenant(id, accountId, inlineObject221);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#updateHostAssignTenant");
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
 **accountId** | **Long**| ID of the Tenant | [optional]
 **inlineObject221** | [**InlineObject221**](InlineObject221.md)|  | [optional]

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
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateHostCloud"></a>
# **updateHostCloud**
> Object updateHostCloud(inlineObject226)

Change Server Cloud

This api call is reserved for migrating servers from one cloud to another. This could be due to moving clusters or resource pool scoping of a server without losing the data.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    InlineObject226 inlineObject226 = new InlineObject226(); // InlineObject226 | 
    try {
      Object result = apiInstance.updateHostCloud(inlineObject226);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#updateHostCloud");
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
 **inlineObject226** | [**InlineObject226**](InlineObject226.md)|  | [optional]

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
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateHostExecuteWorkflow"></a>
# **updateHostExecuteWorkflow**
> Model200Success updateHostExecuteWorkflow(id, workflowId, workflowName, inlineObject225)

Run Workflow on a Host

This will run a provisioning workflow on a host.  For operational workflows, see Execute a Workflow. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Long workflowId = 3L; // Long | ID of the workflow to execute
    String workflowName = "myworkflow"; // String | Name of the workflow to execute
    InlineObject225 inlineObject225 = new InlineObject225(); // InlineObject225 | 
    try {
      Model200Success result = apiInstance.updateHostExecuteWorkflow(id, workflowId, workflowName, inlineObject225);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#updateHostExecuteWorkflow");
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
 **workflowId** | **Long**| ID of the workflow to execute | [optional]
 **workflowName** | **String**| Name of the workflow to execute | [optional]
 **inlineObject225** | [**InlineObject225**](InlineObject225.md)|  | [optional]

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

<a name="updateHostInstallAgent"></a>
# **updateHostInstallAgent**
> Object updateHostInstallAgent(id, inlineObject222)

Install Agent

This will make the host a managed server, and install the agent.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject222 inlineObject222 = new InlineObject222(); // InlineObject222 | 
    try {
      Object result = apiInstance.updateHostInstallAgent(id, inlineObject222);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#updateHostInstallAgent");
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
 **inlineObject222** | [**InlineObject222**](InlineObject222.md)|  | [optional]

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
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateHostManaged"></a>
# **updateHostManaged**
> Object updateHostManaged(id, inlineObject223)

Convert To Managed

This will make the host a managed server, and install the agent.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject223 inlineObject223 = new InlineObject223(); // InlineObject223 | 
    try {
      Object result = apiInstance.updateHostManaged(id, inlineObject223);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#updateHostManaged");
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
 **inlineObject223** | [**InlineObject223**](InlineObject223.md)|  | [optional]

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
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateHostResize"></a>
# **updateHostResize**
> Object updateHostResize(id, inlineObject224)

Resize a Host

Will resize a host asynchronously. This endpoint also allows for NIC reconfiguration by passing a new array of &#x60;networkInterfaces&#x60;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject224 inlineObject224 = new InlineObject224(); // InlineObject224 | 
    try {
      Object result = apiInstance.updateHostResize(id, inlineObject224);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#updateHostResize");
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
 **inlineObject224** | [**InlineObject224**](InlineObject224.md)|  | [optional]

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
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateHostUpgradeAgent"></a>
# **updateHostUpgradeAgent**
> Model200Success updateHostUpgradeAgent(id)

Upgrade Agent

This will upgrade the version of the agent installed on the host.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.updateHostUpgradeAgent(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#updateHostUpgradeAgent");
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

<a name="updateServerNetworkInterface"></a>
# **updateServerNetworkInterface**
> Object updateServerNetworkInterface(id, networkInterfaceId, networkInterfaceUpdate)

Updating a label for a Server&#39;s Network

Updating a Server&#39;s Network&#39;s Label

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal networkInterfaceId = new BigDecimal(78); // BigDecimal | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced
    NetworkInterfaceUpdate networkInterfaceUpdate = {"name":"newLabelName"}; // NetworkInterfaceUpdate | 
    try {
      Object result = apiInstance.updateServerNetworkInterface(id, networkInterfaceId, networkInterfaceUpdate);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#updateServerNetworkInterface");
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
 **networkInterfaceId** | **BigDecimal**| NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced |
 **networkInterfaceUpdate** | [**NetworkInterfaceUpdate**](NetworkInterfaceUpdate.md)|  | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateWikiServer"></a>
# **updateWikiServer**
> Object updateWikiServer(id, inlineObject271)

Update a Server Wiki Page

Updates a server Wiki page. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.HostsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    HostsApi apiInstance = new HostsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject271 inlineObject271 = new InlineObject271(); // InlineObject271 | 
    try {
      Object result = apiInstance.updateWikiServer(id, inlineObject271);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling HostsApi#updateWikiServer");
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
 **inlineObject271** | [**InlineObject271**](InlineObject271.md)|  | [optional]

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
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

