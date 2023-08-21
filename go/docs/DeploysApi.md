# \DeploysApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInstanceDeploy**](DeploysApi.md#AddInstanceDeploy) | **Post** /api/instances/{id}/deploys | Deploy to an Instance
[**Deletedeploy**](DeploysApi.md#Deletedeploy) | **Delete** /api/deploys/{id} | Delete a Deploy
[**GetInstanceDeploys**](DeploysApi.md#GetInstanceDeploys) | **Get** /api/instances/{id}/deploys | Get all Deploys for an Instance
[**ListDeploys**](DeploysApi.md#ListDeploys) | **Get** /api/deploys | Get all Deploys
[**RunDeploy**](DeploysApi.md#RunDeploy) | **Post** /api/deploys/{id}/deploy | Run a Deploy
[**UpdateDeploy**](DeploysApi.md#UpdateDeploy) | **Put** /api/deploys/{id} | Update a Deploy



## AddInstanceDeploy

> InlineResponse20040 AddInstanceDeploy(ctx, id).InlineObject92(inlineObject92).Execute()

Deploy to an Instance



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
    inlineObject92 := *openapiclient.NewInlineObject92() // InlineObject92 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploysApi.AddInstanceDeploy(context.Background(), id).InlineObject92(inlineObject92).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploysApi.AddInstanceDeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInstanceDeploy`: InlineResponse20040
    fmt.Fprintf(os.Stdout, "Response from `DeploysApi.AddInstanceDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddInstanceDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject92** | [**InlineObject92**](InlineObject92.md) |  | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deletedeploy

> Model200Success Deletedeploy(ctx, id).Execute()

Delete a Deploy



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
    resp, r, err := api_client.DeploysApi.Deletedeploy(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploysApi.Deletedeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Deletedeploy`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `DeploysApi.Deletedeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletedeployRequest struct via the builder pattern


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


## GetInstanceDeploys

> map[string]interface{} GetInstanceDeploys(ctx, id).Max(max).Offset(offset).Phrase(phrase).Name(name).DeploymentId(deploymentId).InstanceName(instanceName).InstanceId(instanceId).Version(version).VersionId(versionId).CreatedById(createdById).DeployType(deployType).DateCreated(dateCreated).LastUpdated(lastUpdated).DeployDate(deployDate).Status(status).Execute()

Get all Deploys for an Instance



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
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    deploymentId := int64(5) // int64 | Filter by deployment id (optional)
    instanceName := "instance1" // string | Filter by instance name (optional)
    instanceId := int64(94) // int64 | The Instance ID for Filtering (optional)
    version := int64(5) // int64 | Filter by version number (userVersion) (optional)
    versionId := int64(5) // int64 | Filter by deployment version id (optional)
    createdById := int64(63) // int64 | Filter by owner (user) id (optional)
    deployType := "file" // string | Filter by type (deployType), file, git, fetch (optional)
    dateCreated := "2019-01-01" // string | Filter by dateCreated, the created timestamp is more recent or equal to the date specified (optional)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    deployDate := "2020-01-01" // string | Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified (optional)
    status := "deploying" // string | Filter by status (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploysApi.GetInstanceDeploys(context.Background(), id).Max(max).Offset(offset).Phrase(phrase).Name(name).DeploymentId(deploymentId).InstanceName(instanceName).InstanceId(instanceId).Version(version).VersionId(versionId).CreatedById(createdById).DeployType(deployType).DateCreated(dateCreated).LastUpdated(lastUpdated).DeployDate(deployDate).Status(status).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploysApi.GetInstanceDeploys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceDeploys`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeploysApi.GetInstanceDeploys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceDeploysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **deploymentId** | **int64** | Filter by deployment id | 
 **instanceName** | **string** | Filter by instance name | 
 **instanceId** | **int64** | The Instance ID for Filtering | 
 **version** | **int64** | Filter by version number (userVersion) | 
 **versionId** | **int64** | Filter by deployment version id | 
 **createdById** | **int64** | Filter by owner (user) id | 
 **deployType** | **string** | Filter by type (deployType), file, git, fetch | 
 **dateCreated** | **string** | Filter by dateCreated, the created timestamp is more recent or equal to the date specified | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **deployDate** | **string** | Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified | 
 **status** | **string** | Filter by status | 

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


## ListDeploys

> map[string]interface{} ListDeploys(ctx).Max(max).Offset(offset).Phrase(phrase).Name(name).DeploymentId(deploymentId).InstanceName(instanceName).InstanceId(instanceId).Version(version).VersionId(versionId).CreatedById(createdById).DeployType(deployType).DateCreated(dateCreated).LastUpdated(lastUpdated).DeployDate(deployDate).Status(status).Execute()

Get all Deploys



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
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    deploymentId := int64(5) // int64 | Filter by deployment id (optional)
    instanceName := "instance1" // string | Filter by instance name (optional)
    instanceId := int64(94) // int64 | The Instance ID for Filtering (optional)
    version := int64(5) // int64 | Filter by version number (userVersion) (optional)
    versionId := int64(5) // int64 | Filter by deployment version id (optional)
    createdById := int64(63) // int64 | Filter by owner (user) id (optional)
    deployType := "file" // string | Filter by type (deployType), file, git, fetch (optional)
    dateCreated := "2019-01-01" // string | Filter by dateCreated, the created timestamp is more recent or equal to the date specified (optional)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)
    deployDate := "2020-01-01" // string | Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified (optional)
    status := "deploying" // string | Filter by status (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploysApi.ListDeploys(context.Background()).Max(max).Offset(offset).Phrase(phrase).Name(name).DeploymentId(deploymentId).InstanceName(instanceName).InstanceId(instanceId).Version(version).VersionId(versionId).CreatedById(createdById).DeployType(deployType).DateCreated(dateCreated).LastUpdated(lastUpdated).DeployDate(deployDate).Status(status).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploysApi.ListDeploys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeploys`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeploysApi.ListDeploys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeploysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **deploymentId** | **int64** | Filter by deployment id | 
 **instanceName** | **string** | Filter by instance name | 
 **instanceId** | **int64** | The Instance ID for Filtering | 
 **version** | **int64** | Filter by version number (userVersion) | 
 **versionId** | **int64** | Filter by deployment version id | 
 **createdById** | **int64** | Filter by owner (user) id | 
 **deployType** | **string** | Filter by type (deployType), file, git, fetch | 
 **dateCreated** | **string** | Filter by dateCreated, the created timestamp is more recent or equal to the date specified | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 
 **deployDate** | **string** | Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified | 
 **status** | **string** | Filter by status | 

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


## RunDeploy

> InlineResponse20040 RunDeploy(ctx, id).InlineObject73(inlineObject73).Execute()

Run a Deploy



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
    inlineObject73 := *openapiclient.NewInlineObject73() // InlineObject73 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploysApi.RunDeploy(context.Background(), id).InlineObject73(inlineObject73).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploysApi.RunDeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunDeploy`: InlineResponse20040
    fmt.Fprintf(os.Stdout, "Response from `DeploysApi.RunDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject73** | [**InlineObject73**](InlineObject73.md) |  | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeploy

> InlineResponse20040 UpdateDeploy(ctx, id).InlineObject72(inlineObject72).Execute()

Update a Deploy



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
    inlineObject72 := *openapiclient.NewInlineObject72() // InlineObject72 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploysApi.UpdateDeploy(context.Background(), id).InlineObject72(inlineObject72).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploysApi.UpdateDeploy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeploy`: InlineResponse20040
    fmt.Fprintf(os.Stdout, "Response from `DeploysApi.UpdateDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject72** | [**InlineObject72**](InlineObject72.md) |  | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

