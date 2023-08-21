# \ArchivesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddArchiveBucket**](ArchivesApi.md#AddArchiveBucket) | **Post** /api/archives/buckets | Create an Archive Bucket
[**AddArchiveFile**](ArchivesApi.md#AddArchiveFile) | **Post** /api/archives/buckets/{bucket}/files/{filepath} | Upload Archive File
[**AddArchiveFileLink**](ArchivesApi.md#AddArchiveFileLink) | **Post** /api/archives/files/{id}/links | Create an Archive File Link
[**DeleteArchiveBucket**](ArchivesApi.md#DeleteArchiveBucket) | **Delete** /api/archives/buckets/{id} | Delete an Archive Bucket
[**DeleteArchiveFile**](ArchivesApi.md#DeleteArchiveFile) | **Delete** /api/archives/files/{id} | Delete Archive File
[**DeleteArchiveFileLink**](ArchivesApi.md#DeleteArchiveFileLink) | **Delete** /api/archives/files/{id}/links/{linkId} | Delete an Archive File Link
[**GetArchiveBucket**](ArchivesApi.md#GetArchiveBucket) | **Get** /api/archives/buckets/{id} | Get a Specific Archive Bucket
[**GetArchiveFile**](ArchivesApi.md#GetArchiveFile) | **Get** /api/archives/download/{bucket}/{filepath} | Download an Archive File
[**GetArchiveFileDetail**](ArchivesApi.md#GetArchiveFileDetail) | **Get** /api/archives/files/{id} | Get Archive File Details
[**GetArchiveFileLinks**](ArchivesApi.md#GetArchiveFileLinks) | **Get** /api/archives/files/{id}/links | Get Archive File Links
[**GetArchivePublicFile**](ArchivesApi.md#GetArchivePublicFile) | **Get** /api/public-archives/download/{bucket}/{filepath} | Download a Public Archive File
[**GetArchivePublicFileLink**](ArchivesApi.md#GetArchivePublicFileLink) | **Get** /api/public-archives/link | Download an Archive File Link
[**ListArchiveBuckets**](ArchivesApi.md#ListArchiveBuckets) | **Get** /api/archives/buckets | Get All Archive Buckets
[**ListArchiveFiles**](ArchivesApi.md#ListArchiveFiles) | **Get** /api/archives/buckets/{bucket}/files/{filepath} | Get All Archive Files
[**UpdateArchiveBucket**](ArchivesApi.md#UpdateArchiveBucket) | **Put** /api/archives/buckets/{id} | Update an Archive Bucket



## AddArchiveBucket

> map[string]interface{} AddArchiveBucket(ctx).InlineObject7(inlineObject7).Execute()

Create an Archive Bucket



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
    inlineObject7 := *openapiclient.NewInlineObject7() // InlineObject7 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivesApi.AddArchiveBucket(context.Background()).InlineObject7(inlineObject7).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.AddArchiveBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddArchiveBucket`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.AddArchiveBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddArchiveBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject7** | [**InlineObject7**](InlineObject7.md) |  | 

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


## AddArchiveFile

> map[string]interface{} AddArchiveFile(ctx, bucket, filepath).Filename(filename).File(file).Execute()

Upload Archive File



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
    bucket := "bucket_example" // string | Bucket name
    filepath := "/config/environments/" // string | The path to to search for files under. Default is the root directory /. (default to "/")
    filename := "/path/to/file" // string | Specify a filename for archive file. The base filename of the uploaded file is the default. (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivesApi.AddArchiveFile(context.Background(), bucket, filepath).Filename(filename).File(file).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.AddArchiveFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddArchiveFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.AddArchiveFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket name | 
**filepath** | **string** | The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiAddArchiveFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filename** | **string** | Specify a filename for archive file. The base filename of the uploaded file is the default. | 
 **file** | ***os.File** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddArchiveFileLink

> Model200Success AddArchiveFileLink(ctx, id).ExpireSeconds(expireSeconds).Execute()

Create an Archive File Link



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
    expireSeconds := int64(600) // int64 | Time to live in seconds. 0 means do not expire. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivesApi.AddArchiveFileLink(context.Background(), id).ExpireSeconds(expireSeconds).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.AddArchiveFileLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddArchiveFileLink`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.AddArchiveFileLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddArchiveFileLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expireSeconds** | **int64** | Time to live in seconds. 0 means do not expire. | [default to 0]

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


## DeleteArchiveBucket

> Model200Success DeleteArchiveBucket(ctx, id).Execute()

Delete an Archive Bucket



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
    resp, r, err := api_client.ArchivesApi.DeleteArchiveBucket(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.DeleteArchiveBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteArchiveBucket`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.DeleteArchiveBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArchiveBucketRequest struct via the builder pattern


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


## DeleteArchiveFile

> Model200Success DeleteArchiveFile(ctx, id).Execute()

Delete Archive File



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
    resp, r, err := api_client.ArchivesApi.DeleteArchiveFile(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.DeleteArchiveFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteArchiveFile`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.DeleteArchiveFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArchiveFileRequest struct via the builder pattern


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


## DeleteArchiveFileLink

> Model200Success DeleteArchiveFileLink(ctx, id, linkId).Execute()

Delete an Archive File Link



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
    linkId := int64(789) // int64 | The ID of the archive file link.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivesApi.DeleteArchiveFileLink(context.Background(), id, linkId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.DeleteArchiveFileLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteArchiveFileLink`: Model200Success
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.DeleteArchiveFileLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 
**linkId** | **int64** | The ID of the archive file link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArchiveFileLinkRequest struct via the builder pattern


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


## GetArchiveBucket

> InlineResponse2004 GetArchiveBucket(ctx, id).Execute()

Get a Specific Archive Bucket



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
    resp, r, err := api_client.ArchivesApi.GetArchiveBucket(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.GetArchiveBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArchiveBucket`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.GetArchiveBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchiveBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchiveFile

> GetArchiveFile(ctx, bucket, filepath).Execute()

Download an Archive File



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
    bucket := "bucket_example" // string | Bucket name
    filepath := "/config/environments/" // string | The path to to search for files under. Default is the root directory /. (default to "/")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivesApi.GetArchiveFile(context.Background(), bucket, filepath).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.GetArchiveFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket name | 
**filepath** | **string** | The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchiveFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchiveFileDetail

> InlineResponse2005 GetArchiveFileDetail(ctx, id).Execute()

Get Archive File Details



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
    resp, r, err := api_client.ArchivesApi.GetArchiveFileDetail(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.GetArchiveFileDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArchiveFileDetail`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.GetArchiveFileDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchiveFileDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchiveFileLinks

> map[string]interface{} GetArchiveFileLinks(ctx, id).Execute()

Get Archive File Links



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
    resp, r, err := api_client.ArchivesApi.GetArchiveFileLinks(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.GetArchiveFileLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArchiveFileLinks`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.GetArchiveFileLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchiveFileLinksRequest struct via the builder pattern


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


## GetArchivePublicFile

> GetArchivePublicFile(ctx, bucket, filepath).Execute()

Download a Public Archive File



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
    bucket := "bucket_example" // string | Bucket name
    filepath := "/config/environments/" // string | The path to to search for files under. Default is the root directory /. (default to "/")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivesApi.GetArchivePublicFile(context.Background(), bucket, filepath).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.GetArchivePublicFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket name | 
**filepath** | **string** | The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchivePublicFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchivePublicFileLink

> GetArchivePublicFileLink(ctx).S(s).Execute()

Download an Archive File Link



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
    s := "45a214fce9a546b9" // string | The secret access token for the archive file being downloaded.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivesApi.GetArchivePublicFileLink(context.Background()).S(s).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.GetArchivePublicFileLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetArchivePublicFileLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s** | **string** | The secret access token for the archive file being downloaded. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArchiveBuckets

> map[string]interface{} ListArchiveBuckets(ctx).Name(name).Phrase(phrase).Execute()

Get All Archive Buckets



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
    name := "example-%" // string | Filter by name, wildcard may be specified as %, eg. example-% (optional)
    phrase := "phrase_example" // string | Search phrase for partial matches on name or description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivesApi.ListArchiveBuckets(context.Background()).Name(name).Phrase(phrase).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.ListArchiveBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArchiveBuckets`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.ListArchiveBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListArchiveBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by name, wildcard may be specified as %, eg. example-% | 
 **phrase** | **string** | Search phrase for partial matches on name or description | 

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


## ListArchiveFiles

> map[string]interface{} ListArchiveFiles(ctx, bucket, filepath).Name(name).Phrase(phrase).FullTree(fullTree).Execute()

Get All Archive Files



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
    bucket := "bucket_example" // string | Bucket name
    filepath := "/config/environments/" // string | The path to to search for files under. Default is the root directory /. (default to "/")
    name := "test-config.yaml" // string | If specified will return an exact match on file name (optional)
    phrase := "test-%.yaml" // string | Search phrase for partial matches on file name, wildcard may be specified as %, eg. example-% This also searches for files under sub directories too.  (optional)
    fullTree := true // bool | Include files under sub directories too. This is always true when searching with phrase. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivesApi.ListArchiveFiles(context.Background(), bucket, filepath).Name(name).Phrase(phrase).FullTree(fullTree).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.ListArchiveFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArchiveFiles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.ListArchiveFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket name | 
**filepath** | **string** | The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiListArchiveFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **string** | If specified will return an exact match on file name | 
 **phrase** | **string** | Search phrase for partial matches on file name, wildcard may be specified as %, eg. example-% This also searches for files under sub directories too.  | 
 **fullTree** | **bool** | Include files under sub directories too. This is always true when searching with phrase. | [default to false]

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


## UpdateArchiveBucket

> map[string]interface{} UpdateArchiveBucket(ctx, id).InlineObject8(inlineObject8).Execute()

Update an Archive Bucket



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
    inlineObject8 := *openapiclient.NewInlineObject8() // InlineObject8 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivesApi.UpdateArchiveBucket(context.Background(), id).InlineObject8(inlineObject8).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivesApi.UpdateArchiveBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateArchiveBucket`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArchivesApi.UpdateArchiveBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Morpheus ID of the Object being referenced | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArchiveBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject8** | [**InlineObject8**](InlineObject8.md) |  | 

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

