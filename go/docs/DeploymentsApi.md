# \DeploymentsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDeploymentFile**](DeploymentsApi.md#AddDeploymentFile) | **Post** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | Upload a Deployment File
[**AddDeploymentVersion**](DeploymentsApi.md#AddDeploymentVersion) | **Post** /api/deployments/{deploymentId}/versions | Create a new Deployment Version
[**AddDeployments**](DeploymentsApi.md#AddDeployments) | **Post** /api/deployments | Create a new Deployment
[**DeleteDeployment**](DeploymentsApi.md#DeleteDeployment) | **Delete** /api/deployments/{deploymentId} | Delete a Deployment
[**DeleteDeploymentFile**](DeploymentsApi.md#DeleteDeploymentFile) | **Delete** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | Delete a Deployment File
[**DeleteDeploymentVersion**](DeploymentsApi.md#DeleteDeploymentVersion) | **Delete** /api/deployments/{deploymentId}/versions/{id} | Delete a Deployment Version
[**GetDeployment**](DeploymentsApi.md#GetDeployment) | **Get** /api/deployments/{deploymentId} | Get a Specific Deployment
[**GetDeploymentVersion**](DeploymentsApi.md#GetDeploymentVersion) | **Get** /api/deployments/{deploymentId}/versions/{id} | Get a Specific Deployment Version
[**ListDeploymentFiles**](DeploymentsApi.md#ListDeploymentFiles) | **Get** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | List Deployment Files
[**ListDeploymentVersions**](DeploymentsApi.md#ListDeploymentVersions) | **Get** /api/deployments/{deploymentId}/versions | Get All Versions For a Deployment
[**ListDeployments**](DeploymentsApi.md#ListDeployments) | **Get** /api/deployments | Get All Deployments
[**UpdateDeployment**](DeploymentsApi.md#UpdateDeployment) | **Put** /api/deployments/{deploymentId} | Updating a Deployment
[**UpdateDeploymentVersion**](DeploymentsApi.md#UpdateDeploymentVersion) | **Put** /api/deployments/{deploymentId}/versions/{id} | Updating a Deployment Version



## AddDeploymentFile

> Model200Success AddDeploymentFile(ctx, deploymentId, id, filepath).File(file).Execute()

Upload a Deployment File



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
    deploymentId := int64(4) // int64 | Deployment ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    filepath := "/config/environments/" // string | The path to to search for files under. Default is the root directory /. (default to "/")
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.AddDeploymentFile(context.Background(), deploymentId, id, filepath).File(file).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.AddDeploymentFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDeploymentFile`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.AddDeploymentFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | Deployment ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 
**filepath** | **string** | The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiAddDeploymentFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **file** | ***os.File** |  | 

### Return type

[**Model200Success**](200-success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDeploymentVersion

> map[string]interface{} AddDeploymentVersion(ctx, deploymentId).InlineObject69(inlineObject69).Execute()

Create a new Deployment Version



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
    deploymentId := int64(4) // int64 | Deployment ID
    inlineObject69 := *openapiclient.NewInlineObject69() // InlineObject69 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.AddDeploymentVersion(context.Background(), deploymentId).InlineObject69(inlineObject69).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.AddDeploymentVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDeploymentVersion`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.AddDeploymentVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | Deployment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDeploymentVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject69** | [**InlineObject69**](InlineObject69.md) |  | 

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


## AddDeployments

> map[string]interface{} AddDeployments(ctx).InlineObject67(inlineObject67).Execute()

Create a new Deployment



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
    inlineObject67 := *openapiclient.NewInlineObject67() // InlineObject67 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.AddDeployments(context.Background()).InlineObject67(inlineObject67).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.AddDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDeployments`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.AddDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject67** | [**InlineObject67**](InlineObject67.md) |  | 

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


## DeleteDeployment

> Model200Success DeleteDeployment(ctx, deploymentId).Execute()

Delete a Deployment



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
    deploymentId := int64(4) // int64 | Deployment ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.DeleteDeployment(context.Background(), deploymentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.DeleteDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDeployment`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.DeleteDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | Deployment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentRequest struct via the builder pattern


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


## DeleteDeploymentFile

> Model200Success DeleteDeploymentFile(ctx, deploymentId, id, filepath).Force(force).Execute()

Delete a Deployment File



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
    deploymentId := int64(4) // int64 | Deployment ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    filepath := "/config/environments/" // string | The path to to search for files under. Default is the root directory /. (default to "/")
    force := "on" // string | Force Delete (optional) (default to "off")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.DeleteDeploymentFile(context.Background(), deploymentId, id, filepath).Force(force).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.DeleteDeploymentFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDeploymentFile`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.DeleteDeploymentFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | Deployment ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 
**filepath** | **string** | The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## DeleteDeploymentVersion

> Model200Success DeleteDeploymentVersion(ctx, deploymentId, id).Execute()

Delete a Deployment Version



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
    deploymentId := int64(4) // int64 | Deployment ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.DeleteDeploymentVersion(context.Background(), deploymentId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.DeleteDeploymentVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDeploymentVersion`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.DeleteDeploymentVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | Deployment ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentVersionRequest struct via the builder pattern


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


## GetDeployment

> InlineResponse20038 GetDeployment(ctx, deploymentId).MaxVersions(maxVersions).Execute()

Get a Specific Deployment



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
    deploymentId := int64(4) // int64 | Deployment ID
    maxVersions := int64(6) // int64 | Max number of recent versions to return. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetDeployment(context.Background(), deploymentId).MaxVersions(maxVersions).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeployment`: InlineResponse20038
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | Deployment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxVersions** | **int64** | Max number of recent versions to return. | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentVersion

> InlineResponse20039 GetDeploymentVersion(ctx, deploymentId, id).Execute()

Get a Specific Deployment Version



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
    deploymentId := int64(4) // int64 | Deployment ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.GetDeploymentVersion(context.Background(), deploymentId, id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeploymentVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentVersion`: InlineResponse20039
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeploymentVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | Deployment ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeploymentFiles

> map[string]interface{} ListDeploymentFiles(ctx, deploymentId, id, filepath).Max(max).Offset(offset).Phrase(phrase).Version(version).Execute()

List Deployment Files



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
    deploymentId := int64(4) // int64 | Deployment ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    filepath := "/config/environments/" // string | The path to to search for files under. Default is the root directory /. (default to "/")
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    version := int64(5) // int64 | Filter by version number (userVersion) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ListDeploymentFiles(context.Background(), deploymentId, id, filepath).Max(max).Offset(offset).Phrase(phrase).Version(version).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ListDeploymentFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeploymentFiles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ListDeploymentFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | Deployment ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 
**filepath** | **string** | The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiListDeploymentFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **version** | **int64** | Filter by version number (userVersion) | 

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


## ListDeploymentVersions

> map[string]interface{} ListDeploymentVersions(ctx, deploymentId).Max(max).Offset(offset).Phrase(phrase).Version(version).Type_(type_).DateCreated(dateCreated).LastUpdated(lastUpdated).Execute()

Get All Versions For a Deployment



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
    deploymentId := int64(4) // int64 | Deployment ID
    max := int64(789) // int64 | Maximum number of records to return, -1 can be used to fetch all records (optional) (default to 25)
    offset := int64(789) // int64 | Offset records, the number of records to skip, for paginating requests (optional) (default to 0)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)
    version := int64(5) // int64 | Filter by version number (userVersion) (optional)
    type_ := "file" // string | Filter by type (deployType), file, git, fetch (optional)
    dateCreated := "2019-01-01" // string | Filter by dateCreated, the created timestamp is more recent or equal to the date specified (optional)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ListDeploymentVersions(context.Background(), deploymentId).Max(max).Offset(offset).Phrase(phrase).Version(version).Type_(type_).DateCreated(dateCreated).LastUpdated(lastUpdated).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ListDeploymentVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeploymentVersions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ListDeploymentVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | Deployment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeploymentVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **version** | **int64** | Filter by version number (userVersion) | 
 **type_** | **string** | Filter by type (deployType), file, git, fetch | 
 **dateCreated** | **string** | Filter by dateCreated, the created timestamp is more recent or equal to the date specified | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 

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


## ListDeployments

> map[string]interface{} ListDeployments(ctx).Max(max).Offset(offset).Phrase(phrase).Name(name).Description(description).DateCreated(dateCreated).LastUpdated(lastUpdated).Execute()

Get All Deployments



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
    description := "The desription of my cool object" // string | Filter by description, wildcard may be specified as %. eg. `example-%` (optional)
    dateCreated := "2019-01-01" // string | Filter by dateCreated, the created timestamp is more recent or equal to the date specified (optional)
    lastUpdated := time.Now() // time.Time | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.ListDeployments(context.Background()).Max(max).Offset(offset).Phrase(phrase).Name(name).Description(description).DateCreated(dateCreated).LastUpdated(lastUpdated).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ListDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeployments`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ListDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **int64** | Maximum number of records to return, -1 can be used to fetch all records | [default to 25]
 **offset** | **int64** | Offset records, the number of records to skip, for paginating requests | [default to 0]
 **phrase** | **string** | Search phrase for partial matches on name or description | 
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **description** | **string** | Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | 
 **dateCreated** | **string** | Filter by dateCreated, the created timestamp is more recent or equal to the date specified | 
 **lastUpdated** | **time.Time** | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | 

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


## UpdateDeployment

> map[string]interface{} UpdateDeployment(ctx, deploymentId).InlineObject68(inlineObject68).Execute()

Updating a Deployment



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
    deploymentId := int64(4) // int64 | Deployment ID
    inlineObject68 := *openapiclient.NewInlineObject68() // InlineObject68 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.UpdateDeployment(context.Background(), deploymentId).InlineObject68(inlineObject68).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.UpdateDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeployment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.UpdateDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | Deployment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject68** | [**InlineObject68**](InlineObject68.md) |  | 

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


## UpdateDeploymentVersion

> map[string]interface{} UpdateDeploymentVersion(ctx, deploymentId, id).InlineObject70(inlineObject70).Execute()

Updating a Deployment Version



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
    deploymentId := int64(4) // int64 | Deployment ID
    id := int64(1) // int64 | Morpheus ID of the Object being referenced
    inlineObject70 := *openapiclient.NewInlineObject70() // InlineObject70 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeploymentsApi.UpdateDeploymentVersion(context.Background(), deploymentId, id).InlineObject70(inlineObject70).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.UpdateDeploymentVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeploymentVersion`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.UpdateDeploymentVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **int64** | Deployment ID | 
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject70** | [**InlineObject70**](InlineObject70.md) |  | 

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

