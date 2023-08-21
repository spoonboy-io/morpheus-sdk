# \HostsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHost**](HostsApi.md#GetHost) | **Get** /api/servers/{id} | Get a Specific Host
[**GetHostSnpshots**](HostsApi.md#GetHostSnpshots) | **Get** /api/servers/{id}/snapshots | Get list of snapshots for a Host
[**GetHostType**](HostsApi.md#GetHostType) | **Get** /api/server-types/{id} | Get a Specific Host Type
[**GetWikiServer**](HostsApi.md#GetWikiServer) | **Get** /api/servers/{id}/wiki | Retrieves a Server Wiki Page
[**ListHostTypes**](HostsApi.md#ListHostTypes) | **Get** /api/server-types | Host Types
[**ListHosts**](HostsApi.md#ListHosts) | **Get** /api/servers | Get All Hosts
[**ListServerServicePlans**](HostsApi.md#ListServerServicePlans) | **Get** /api/servers/service-plans | Get Available Service Plans for a Host
[**RemoveHost**](HostsApi.md#RemoveHost) | **Delete** /api/servers/{id} | Delete a Host
[**RestartHost**](HostsApi.md#RestartHost) | **Put** /api/servers/{id}/restart | Restart a Host
[**StartHost**](HostsApi.md#StartHost) | **Put** /api/servers/{id}/start | Start a Host
[**StopHost**](HostsApi.md#StopHost) | **Put** /api/servers/{id}/stop | Stop a Host
[**UpdateHost**](HostsApi.md#UpdateHost) | **Put** /api/servers/{id} | Updating a Host
[**UpdateHostAssignTenant**](HostsApi.md#UpdateHostAssignTenant) | **Put** /api/servers/{id}/assign-account | Assign To Tenant
[**UpdateHostCloud**](HostsApi.md#UpdateHostCloud) | **Put** /api/servers/change-cloud | Change Server Cloud
[**UpdateHostExecuteWorkflow**](HostsApi.md#UpdateHostExecuteWorkflow) | **Put** /api/servers/{id}/workflow | Run Workflow on a Host
[**UpdateHostInstallAgent**](HostsApi.md#UpdateHostInstallAgent) | **Put** /api/servers/{id}/install-agent | Install Agent
[**UpdateHostManaged**](HostsApi.md#UpdateHostManaged) | **Put** /api/servers/{id}/make-managed | Convert To Managed
[**UpdateHostResize**](HostsApi.md#UpdateHostResize) | **Put** /api/servers/{id}/resize | Resize a Host
[**UpdateHostUpgradeAgent**](HostsApi.md#UpdateHostUpgradeAgent) | **Put** /api/servers/{id}/upgrade | Upgrade Agent
[**UpdateServerNetworkInterface**](HostsApi.md#UpdateServerNetworkInterface) | **Put** /api/servers/{id}/networkInterfaces/{networkInterfaceId} | Updating a label for a Server&#39;s Network
[**UpdateWikiServer**](HostsApi.md#UpdateWikiServer) | **Put** /api/servers/{id}/wiki | Update a Server Wiki Page



## GetHost

> InlineResponse200137 GetHost(ctx, id).Execute()

Get a Specific Host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.GetHost(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.GetHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHost`: InlineResponse200137
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.GetHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200137**](inline_response_200_137.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostSnpshots

> InlineResponse200138 GetHostSnpshots(ctx, id).Execute()

Get list of snapshots for a Host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.GetHostSnpshots(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.GetHostSnpshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostSnpshots`: InlineResponse200138
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.GetHostSnpshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostSnpshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200138**](inline_response_200_138.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostType

> InlineResponse20050 GetHostType(ctx, id).Execute()

Get a Specific Host Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.GetHostType(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.GetHostType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostType`: InlineResponse20050
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.GetHostType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWikiServer

> InlineResponse200168 GetWikiServer(ctx, id).Execute()

Retrieves a Server Wiki Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.GetWikiServer(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.GetWikiServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWikiServer`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.GetWikiServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200168**](inline_response_200_168.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHostTypes

> map[string]interface{} ListHostTypes(ctx).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Code(code).Phrase(phrase).ProvisionType(provisionType).ZoneType(zoneType).Creatable(creatable).Execute()

Host Types



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    sort := "sort_example" // string | Sort order, the name of the property to sort by (optional) (default to "name")
    direction := "asc" // string | Sort direction, use 'desc' to reverse sort (optional) (default to "asc")
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    code := "azr" // string | If specified will return an exact match on code (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    provisionType := "provisionType_example" // string | Filter by `Provision Type` code. Refer to `Provision Types` API for up to date listings.  (optional)
    zoneType := "zoneType_example" // string | Filter by Cloud Type code. (optional)
    creatable := true // bool | Filter by creatable flag. This is whether or not it can be provisioned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ListHostTypes(context.Background()).Max(max).Offset(offset).Sort(sort).Direction(direction).Name(name).Code(code).Phrase(phrase).ProvisionType(provisionType).ZoneType(zoneType).Creatable(creatable).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ListHostTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHostTypes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ListHostTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **sort** | **string** | Sort order, the name of the property to sort by | [default to &quot;name&quot;]
 **direction** | **string** | Sort direction, use &#39;desc&#39; to reverse sort | [default to &quot;asc&quot;]
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **code** | **string** | If specified will return an exact match on code | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **provisionType** | **string** | Filter by &#x60;Provision Type&#x60; code. Refer to &#x60;Provision Types&#x60; API for up to date listings.  | 
 **zoneType** | **string** | Filter by Cloud Type code. | 
 **creatable** | **bool** | Filter by creatable flag. This is whether or not it can be provisioned. | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosts

> map[string]interface{} ListHosts(ctx).Name(name).Phrase(phrase).ZoneId(zoneId).SiteId(siteId).ClusterId(clusterId).Managed(managed).ServerType(serverType).PowerState(powerState).Ip(ip).Vm(vm).VmHypervisor(vmHypervisor).BareMetalHost(bareMetalHost).Status(status).AgentInstalled(agentInstalled).Max(max).Offset(offset).LastUpdated(lastUpdated).CreatedBy(createdBy).Labels(labels).AllLabels(allLabels).Tags(tags).Metadata(metadata).Uuid(uuid).ExternalId(externalId).InternalId(internalId).ExternalUniquelId(externalUniquelId).Execute()

Get All Hosts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    zoneId := int64(3) // int64 | The Zone ID for Filtering (optional)
    siteId := int64(7) // int64 | The Site ID for Filtering (optional)
    clusterId := int64(135) // int64 | The Cluster ID(s) for filtering. Accepts multiple values. (optional)
    managed := false // bool | Filter by managed (true) or unmanaged (false) (optional)
    serverType := "vmwareHypervisor" // string | Filter by server type code (optional)
    powerState := "off" // string | Filter by power status (optional)
    ip := "192.168.1.45" // string | Filter by IP address (optional)
    vm := false // bool | Filter to show only Virtual Machines (true) (optional)
    vmHypervisor := false // bool | Filter to show only VM Hypervisors (true) (optional)
    bareMetalHost := false // bool | Filter to show only Baremetal Servers (optional)
    status := "status_example" // string | Filter by status (optional)
    agentInstalled := true // bool | Filter by agent installed (true) (optional)
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    createdBy := int64(63) // int64 | The User ID for Filtering (optional)
    labels := "labels_example" // string | Filter by label(s), matches records that contain any of the specified labels (optional)
    allLabels := "allLabels_example" // string | Filter by label(s), matches records that contain all of the specified labels (optional)
    tags := "tags.env=qa&tags.env=test" // string | Filter by tags (metadata). This allows filtering by a tag name and value(s)  (optional)
    metadata := "tags.env=qa&tags.env=test" // string | Alias for tags (optional)
    uuid := "uuid_example" // string | Filter by UUID (optional)
    externalId := "externalId_example" // string | Filter by External ID (optional)
    internalId := "internalId_example" // string | Filter by Internal ID (optional)
    externalUniquelId := "externalUniquelId_example" // string | Filter by External Unique ID (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ListHosts(context.Background()).Name(name).Phrase(phrase).ZoneId(zoneId).SiteId(siteId).ClusterId(clusterId).Managed(managed).ServerType(serverType).PowerState(powerState).Ip(ip).Vm(vm).VmHypervisor(vmHypervisor).BareMetalHost(bareMetalHost).Status(status).AgentInstalled(agentInstalled).Max(max).Offset(offset).LastUpdated(lastUpdated).CreatedBy(createdBy).Labels(labels).AllLabels(allLabels).Tags(tags).Metadata(metadata).Uuid(uuid).ExternalId(externalId).InternalId(internalId).ExternalUniquelId(externalUniquelId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ListHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHosts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ListHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **zoneId** | **int64** | The Zone ID for Filtering | 
 **siteId** | **int64** | The Site ID for Filtering | 
 **clusterId** | **int64** | The Cluster ID(s) for filtering. Accepts multiple values. | 
 **managed** | **bool** | Filter by managed (true) or unmanaged (false) | 
 **serverType** | **string** | Filter by server type code | 
 **powerState** | **string** | Filter by power status | 
 **ip** | **string** | Filter by IP address | 
 **vm** | **bool** | Filter to show only Virtual Machines (true) | 
 **vmHypervisor** | **bool** | Filter to show only VM Hypervisors (true) | 
 **bareMetalHost** | **bool** | Filter to show only Baremetal Servers | 
 **status** | **string** | Filter by status | 
 **agentInstalled** | **bool** | Filter by agent installed (true) | 
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **createdBy** | **int64** | The User ID for Filtering | 
 **labels** | **string** | Filter by label(s), matches records that contain any of the specified labels | 
 **allLabels** | **string** | Filter by label(s), matches records that contain all of the specified labels | 
 **tags** | **string** | Filter by tags (metadata). This allows filtering by a tag name and value(s)  | 
 **metadata** | **string** | Alias for tags | 
 **uuid** | **string** | Filter by UUID | 
 **externalId** | **string** | Filter by External ID | 
 **internalId** | **string** | Filter by Internal ID | 
 **externalUniquelId** | **string** | Filter by External Unique ID | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServerServicePlans

> InlineResponse200141 ListServerServicePlans(ctx).ZoneId(zoneId).ServerTypeId(serverTypeId).SiteId(siteId).Execute()

Get Available Service Plans for a Host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    zoneId := int64(3) // int64 | The Zone ID for Filtering
    serverTypeId := int64(5) // int64 | The ID of the Host Type (optional)
    siteId := int64(7) // int64 | The Site ID for Filtering (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ListServerServicePlans(context.Background()).ZoneId(zoneId).ServerTypeId(serverTypeId).SiteId(siteId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ListServerServicePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServerServicePlans`: InlineResponse200141
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ListServerServicePlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServerServicePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **int64** | The Zone ID for Filtering | 
 **serverTypeId** | **int64** | The ID of the Host Type | 
 **siteId** | **int64** | The Site ID for Filtering | 

### Return type

[**InlineResponse200141**](inline_response_200_141.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveHost

> Model200Success RemoveHost(ctx, id).RemoveResources(removeResources).RemoveInstances(removeInstances).PreserveVolumes(preserveVolumes).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).Force(force).Execute()

Delete a Host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    removeResources := "off" // string | Remove Resources (optional) (default to "on")
    removeInstances := "on" // string | Remove Instances (optional) (default to "off")
    preserveVolumes := "on" // string | Preserve Volumes (optional) (default to "off")
    releaseFloatingIps := "off" // string | Release Floating IPs (optional) (default to "on")
    releaseEIPs := "off" // string | Alias for releaseFloatingIps (optional) (default to "on")
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.RemoveHost(context.Background(), id).RemoveResources(removeResources).RemoveInstances(removeInstances).PreserveVolumes(preserveVolumes).ReleaseFloatingIps(releaseFloatingIps).ReleaseEIPs(releaseEIPs).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.RemoveHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveHost`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.RemoveHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeResources** | **string** | Remove Resources | [default to &quot;on&quot;]
 **removeInstances** | **string** | Remove Instances | [default to &quot;off&quot;]
 **preserveVolumes** | **string** | Preserve Volumes | [default to &quot;off&quot;]
 **releaseFloatingIps** | **string** | Release Floating IPs | [default to &quot;on&quot;]
 **releaseEIPs** | **string** | Alias for releaseFloatingIps | [default to &quot;on&quot;]
 **force** | **string** | Force Delete | [default to &quot;off&quot;]

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartHost

> map[string]interface{} RestartHost(ctx, id).Execute()

Restart a Host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.RestartHost(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.RestartHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartHost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.RestartHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartHost

> Model200Success StartHost(ctx, id).Execute()

Start a Host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.StartHost(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.StartHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartHost`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.StartHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopHost

> Model200Success StopHost(ctx, id).Execute()

Stop a Host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.StopHost(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.StopHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopHost`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.StopHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHost

> InlineResponse200137 UpdateHost(ctx, id).InlineObject220(inlineObject220).Execute()

Updating a Host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject220 := *openapiclient.NewInlineObject220() // InlineObject220 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.UpdateHost(context.Background(), id).InlineObject220(inlineObject220).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.UpdateHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHost`: InlineResponse200137
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.UpdateHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject220** | [**InlineObject220**](InlineObject220.md) |  | 

### Return type

[**InlineResponse200137**](inline_response_200_137.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostAssignTenant

> map[string]interface{} UpdateHostAssignTenant(ctx, id).AccountId(accountId).InlineObject221(inlineObject221).Execute()

Assign To Tenant



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    accountId := int64(3) // int64 | ID of the Tenant (optional)
    inlineObject221 := *openapiclient.NewInlineObject221() // InlineObject221 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.UpdateHostAssignTenant(context.Background(), id).AccountId(accountId).InlineObject221(inlineObject221).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.UpdateHostAssignTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHostAssignTenant`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.UpdateHostAssignTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostAssignTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **int64** | ID of the Tenant | 
 **inlineObject221** | [**InlineObject221**](InlineObject221.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostCloud

> map[string]interface{} UpdateHostCloud(ctx).InlineObject226(inlineObject226).Execute()

Change Server Cloud



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject226 := *openapiclient.NewInlineObject226() // InlineObject226 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.UpdateHostCloud(context.Background()).InlineObject226(inlineObject226).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.UpdateHostCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHostCloud`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.UpdateHostCloud`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostCloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject226** | [**InlineObject226**](InlineObject226.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostExecuteWorkflow

> Model200Success UpdateHostExecuteWorkflow(ctx, id).WorkflowId(workflowId).WorkflowName(workflowName).InlineObject225(inlineObject225).Execute()

Run Workflow on a Host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    workflowId := int64(3) // int64 | ID of the workflow to execute (optional)
    workflowName := "myworkflow" // string | Name of the workflow to execute (optional)
    inlineObject225 := *openapiclient.NewInlineObject225() // InlineObject225 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.UpdateHostExecuteWorkflow(context.Background(), id).WorkflowId(workflowId).WorkflowName(workflowName).InlineObject225(inlineObject225).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.UpdateHostExecuteWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHostExecuteWorkflow`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.UpdateHostExecuteWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostExecuteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowId** | **int64** | ID of the workflow to execute | 
 **workflowName** | **string** | Name of the workflow to execute | 
 **inlineObject225** | [**InlineObject225**](InlineObject225.md) |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostInstallAgent

> map[string]interface{} UpdateHostInstallAgent(ctx, id).InlineObject222(inlineObject222).Execute()

Install Agent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject222 := *openapiclient.NewInlineObject222() // InlineObject222 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.UpdateHostInstallAgent(context.Background(), id).InlineObject222(inlineObject222).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.UpdateHostInstallAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHostInstallAgent`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.UpdateHostInstallAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostInstallAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject222** | [**InlineObject222**](InlineObject222.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostManaged

> map[string]interface{} UpdateHostManaged(ctx, id).InlineObject223(inlineObject223).Execute()

Convert To Managed



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject223 := *openapiclient.NewInlineObject223() // InlineObject223 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.UpdateHostManaged(context.Background(), id).InlineObject223(inlineObject223).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.UpdateHostManaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHostManaged`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.UpdateHostManaged`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostManagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject223** | [**InlineObject223**](InlineObject223.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostResize

> map[string]interface{} UpdateHostResize(ctx, id).InlineObject224(inlineObject224).Execute()

Resize a Host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject224 := *openapiclient.NewInlineObject224() // InlineObject224 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.UpdateHostResize(context.Background(), id).InlineObject224(inlineObject224).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.UpdateHostResize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHostResize`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.UpdateHostResize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostResizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject224** | [**InlineObject224**](InlineObject224.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostUpgradeAgent

> Model200Success UpdateHostUpgradeAgent(ctx, id).Execute()

Upgrade Agent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.UpdateHostUpgradeAgent(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.UpdateHostUpgradeAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHostUpgradeAgent`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.UpdateHostUpgradeAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostUpgradeAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerNetworkInterface

> map[string]interface{} UpdateServerNetworkInterface(ctx, id, networkInterfaceId).NetworkInterfaceUpdate(networkInterfaceUpdate).Execute()

Updating a label for a Server's Network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    networkInterfaceId := float32(7) // float32 | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced
    networkInterfaceUpdate := *openapiclient.NewNetworkInterfaceUpdate() // NetworkInterfaceUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.UpdateServerNetworkInterface(context.Background(), id, networkInterfaceId).NetworkInterfaceUpdate(networkInterfaceUpdate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.UpdateServerNetworkInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServerNetworkInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.UpdateServerNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**networkInterfaceId** | **float32** | NetworkInterface (ComputeServerInterface) ID of the Object being created or referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkInterfaceUpdate** | [**NetworkInterfaceUpdate**](NetworkInterfaceUpdate.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWikiServer

> map[string]interface{} UpdateWikiServer(ctx, id).InlineObject271(inlineObject271).Execute()

Update a Server Wiki Page



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject271 := *openapiclient.NewInlineObject271() // InlineObject271 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.UpdateWikiServer(context.Background(), id).InlineObject271(inlineObject271).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.UpdateWikiServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWikiServer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.UpdateWikiServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWikiServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject271** | [**InlineObject271**](InlineObject271.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

