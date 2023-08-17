# ArchivesApi

All URIs are relative to *https://CHANGEME*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addArchiveBucket**](ArchivesApi.md#addArchiveBucket) | **POST** /api/archives/buckets | Create an Archive Bucket |
| [**addArchiveFile**](ArchivesApi.md#addArchiveFile) | **POST** /api/archives/buckets/{bucket}/files/{filepath} | Upload Archive File |
| [**addArchiveFileLink**](ArchivesApi.md#addArchiveFileLink) | **POST** /api/archives/files/{id}/links | Create an Archive File Link |
| [**deleteArchiveBucket**](ArchivesApi.md#deleteArchiveBucket) | **DELETE** /api/archives/buckets/{id} | Delete an Archive Bucket |
| [**deleteArchiveFile**](ArchivesApi.md#deleteArchiveFile) | **DELETE** /api/archives/files/{id} | Delete Archive File |
| [**deleteArchiveFileLink**](ArchivesApi.md#deleteArchiveFileLink) | **DELETE** /api/archives/files/{id}/links/{linkId} | Delete an Archive File Link |
| [**getArchiveBucket**](ArchivesApi.md#getArchiveBucket) | **GET** /api/archives/buckets/{id} | Get a Specific Archive Bucket |
| [**getArchiveFile**](ArchivesApi.md#getArchiveFile) | **GET** /api/archives/download/{bucket}/{filepath} | Download an Archive File |
| [**getArchiveFileDetail**](ArchivesApi.md#getArchiveFileDetail) | **GET** /api/archives/files/{id} | Get Archive File Details |
| [**getArchiveFileLinks**](ArchivesApi.md#getArchiveFileLinks) | **GET** /api/archives/files/{id}/links | Get Archive File Links |
| [**listArchiveBuckets**](ArchivesApi.md#listArchiveBuckets) | **GET** /api/archives/buckets | Get All Archive Buckets |
| [**listArchiveFiles**](ArchivesApi.md#listArchiveFiles) | **GET** /api/archives/buckets/{bucket}/files/{filepath} | Get All Archive Files |
| [**updateArchiveBucket**](ArchivesApi.md#updateArchiveBucket) | **PUT** /api/archives/buckets/{id} | Update an Archive Bucket |


<a id="addArchiveBucket"></a>
# **addArchiveBucket**
> AddArchiveBucket200Response addArchiveBucket(addArchiveBucketRequest)

Create an Archive Bucket

Create an Archive Bucket

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    AddArchiveBucketRequest addArchiveBucketRequest = new AddArchiveBucketRequest(); // AddArchiveBucketRequest | 
    try {
      AddArchiveBucket200Response result = apiInstance.addArchiveBucket(addArchiveBucketRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#addArchiveBucket");
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
| **addArchiveBucketRequest** | [**AddArchiveBucketRequest**](AddArchiveBucketRequest.md)|  | [optional] |

### Return type

[**AddArchiveBucket200Response**](AddArchiveBucket200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="addArchiveFile"></a>
# **addArchiveFile**
> AddArchiveFile200Response addArchiveFile(bucket, filepath, filename, _file)

Upload Archive File

Upload a file to the specified archive bucket and file path.  This will overwrite the file if it already exists. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    String bucket = "bucket_example"; // String | Bucket name
    String filepath = "/"; // String | The path to to search for files under. Default is the root directory /.
    String filename = "/path/to/file"; // String | Specify a filename for archive file. The base filename of the uploaded file is the default.
    File _file = new File("/path/to/file"); // File | 
    try {
      AddArchiveFile200Response result = apiInstance.addArchiveFile(bucket, filepath, filename, _file);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#addArchiveFile");
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
| **bucket** | **String**| Bucket name | |
| **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to /] |
| **filename** | **String**| Specify a filename for archive file. The base filename of the uploaded file is the default. | [optional] |
| **_file** | **File**|  | [optional] |

### Return type

[**AddArchiveFile200Response**](AddArchiveFile200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="addArchiveFileLink"></a>
# **addArchiveFileLink**
> AddArchiveFileLink200Response addArchiveFileLink(id, expireSeconds)

Create an Archive File Link

This returns a secret token that can be used to download the file via a public url, without any other authentication or authorization. File links can be set to expire after a certain amount of time.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Long expireSeconds = 0L; // Long | Time to live in seconds. 0 means do not expire.
    try {
      AddArchiveFileLink200Response result = apiInstance.addArchiveFileLink(id, expireSeconds);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#addArchiveFileLink");
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
| **expireSeconds** | **Long**| Time to live in seconds. 0 means do not expire. | [optional] [default to 0] |

### Return type

[**AddArchiveFileLink200Response**](AddArchiveFileLink200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="deleteArchiveBucket"></a>
# **deleteArchiveBucket**
> Model200Success deleteArchiveBucket(id)

Delete an Archive Bucket

Will delete an archive bucket from the system and make it no longer usable.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteArchiveBucket(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#deleteArchiveBucket");
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
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="deleteArchiveFile"></a>
# **deleteArchiveFile**
> Model200Success deleteArchiveFile(id)

Delete Archive File

Permanently delete a file or directory.  Deleting a directory will also delete all the files under it. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteArchiveFile(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#deleteArchiveFile");
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
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="deleteArchiveFileLink"></a>
# **deleteArchiveFileLink**
> Model200Success deleteArchiveFileLink(id, linkId)

Delete an Archive File Link

This will delete the link from the system, so it can no longer be used.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Long linkId = 56L; // Long | The ID of the archive file link.
    try {
      Model200Success result = apiInstance.deleteArchiveFileLink(id, linkId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#deleteArchiveFileLink");
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
| **linkId** | **Long**| The ID of the archive file link. | |

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
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="getArchiveBucket"></a>
# **getArchiveBucket**
> GetArchiveBucket200Response getArchiveBucket(id)

Get a Specific Archive Bucket

This endpoint retrieves a specific archive bucket.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      GetArchiveBucket200Response result = apiInstance.getArchiveBucket(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#getArchiveBucket");
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

### Return type

[**GetArchiveBucket200Response**](GetArchiveBucket200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="getArchiveFile"></a>
# **getArchiveFile**
> getArchiveFile(bucket, filepath)

Download an Archive File

Download the file as an authorized user with access to the bucket.  Downloading a directory will return a .zip file containing all files under it. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    String bucket = "bucket_example"; // String | Bucket name
    String filepath = "/"; // String | The path to to search for files under. Default is the root directory /.
    try {
      apiInstance.getArchiveFile(bucket, filepath);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#getArchiveFile");
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
| **bucket** | **String**| Bucket name | |
| **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to /] |

### Return type

null (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Returns the contents of the specified file as an attachment with Content-Type dictated by the file |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="getArchiveFileDetail"></a>
# **getArchiveFileDetail**
> GetArchiveFileDetail200Response getArchiveFileDetail(id)

Get Archive File Details

Get details about a specific archive file.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      GetArchiveFileDetail200Response result = apiInstance.getArchiveFileDetail(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#getArchiveFileDetail");
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

### Return type

[**GetArchiveFileDetail200Response**](GetArchiveFileDetail200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="getArchiveFileLinks"></a>
# **getArchiveFileLinks**
> GetArchiveFileLinks200Response getArchiveFileLinks(id)

Get Archive File Links

This endpoint retrieves the links that have been created for the specified file.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      GetArchiveFileLinks200Response result = apiInstance.getArchiveFileLinks(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#getArchiveFileLinks");
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

### Return type

[**GetArchiveFileLinks200Response**](GetArchiveFileLinks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="listArchiveBuckets"></a>
# **listArchiveBuckets**
> ListArchiveBuckets200Response listArchiveBuckets(name, phrase)

Get All Archive Buckets

This endpoint retrieves all archive buckets associated with the account.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    try {
      ListArchiveBuckets200Response result = apiInstance.listArchiveBuckets(name, phrase);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#listArchiveBuckets");
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
| **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] |
| **phrase** | **String**| Search phrase for partial matches on name or description | [optional] |

### Return type

[**ListArchiveBuckets200Response**](ListArchiveBuckets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="listArchiveFiles"></a>
# **listArchiveFiles**
> ListArchiveFiles200Response listArchiveFiles(bucket, filepath, name, phrase, fullTree)

Get All Archive Files

This endpoint retrieves all files in an archive bucket under the specified &#x60;filePath&#x60;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    String bucket = "bucket_example"; // String | Bucket name
    String filepath = "/"; // String | The path to to search for files under. Default is the root directory /.
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    Boolean fullTree = false; // Boolean | Include files under sub directories too. This is always true when searching with phrase.
    try {
      ListArchiveFiles200Response result = apiInstance.listArchiveFiles(bucket, filepath, name, phrase, fullTree);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#listArchiveFiles");
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
| **bucket** | **String**| Bucket name | |
| **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to /] |
| **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] |
| **phrase** | **String**| Search phrase for partial matches on name or description | [optional] |
| **fullTree** | **Boolean**| Include files under sub directories too. This is always true when searching with phrase. | [optional] [default to false] |

### Return type

[**ListArchiveFiles200Response**](ListArchiveFiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="updateArchiveBucket"></a>
# **updateArchiveBucket**
> AddArchiveBucket200Response updateArchiveBucket(id, updateArchiveBucketRequest)

Update an Archive Bucket

Update an Archive Bucket

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ArchivesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ArchivesApi apiInstance = new ArchivesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    UpdateArchiveBucketRequest updateArchiveBucketRequest = new UpdateArchiveBucketRequest(); // UpdateArchiveBucketRequest | 
    try {
      AddArchiveBucket200Response result = apiInstance.updateArchiveBucket(id, updateArchiveBucketRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#updateArchiveBucket");
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
| **updateArchiveBucketRequest** | [**UpdateArchiveBucketRequest**](UpdateArchiveBucketRequest.md)|  | [optional] |

### Return type

[**AddArchiveBucket200Response**](AddArchiveBucket200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

