# ArchivesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addArchiveBucket**](ArchivesApi.md#addArchiveBucket) | **POST** /api/archives/buckets | Create an Archive Bucket
[**addArchiveFile**](ArchivesApi.md#addArchiveFile) | **POST** /api/archives/buckets/{bucket}/files/{filepath} | Upload Archive File
[**addArchiveFileLink**](ArchivesApi.md#addArchiveFileLink) | **POST** /api/archives/files/{id}/links | Create an Archive File Link
[**deleteArchiveBucket**](ArchivesApi.md#deleteArchiveBucket) | **DELETE** /api/archives/buckets/{id} | Delete an Archive Bucket
[**deleteArchiveFile**](ArchivesApi.md#deleteArchiveFile) | **DELETE** /api/archives/files/{id} | Delete Archive File
[**deleteArchiveFileLink**](ArchivesApi.md#deleteArchiveFileLink) | **DELETE** /api/archives/files/{id}/links/{linkId} | Delete an Archive File Link
[**getArchiveBucket**](ArchivesApi.md#getArchiveBucket) | **GET** /api/archives/buckets/{id} | Get a Specific Archive Bucket
[**getArchiveFile**](ArchivesApi.md#getArchiveFile) | **GET** /api/archives/download/{bucket}/{filepath} | Download an Archive File
[**getArchiveFileDetail**](ArchivesApi.md#getArchiveFileDetail) | **GET** /api/archives/files/{id} | Get Archive File Details
[**getArchiveFileLinks**](ArchivesApi.md#getArchiveFileLinks) | **GET** /api/archives/files/{id}/links | Get Archive File Links
[**getArchivePublicFile**](ArchivesApi.md#getArchivePublicFile) | **GET** /api/public-archives/download/{bucket}/{filepath} | Download a Public Archive File
[**getArchivePublicFileLink**](ArchivesApi.md#getArchivePublicFileLink) | **GET** /api/public-archives/link | Download an Archive File Link
[**listArchiveBuckets**](ArchivesApi.md#listArchiveBuckets) | **GET** /api/archives/buckets | Get All Archive Buckets
[**listArchiveFiles**](ArchivesApi.md#listArchiveFiles) | **GET** /api/archives/buckets/{bucket}/files/{filepath} | Get All Archive Files
[**updateArchiveBucket**](ArchivesApi.md#updateArchiveBucket) | **PUT** /api/archives/buckets/{id} | Update an Archive Bucket


<a name="addArchiveBucket"></a>
# **addArchiveBucket**
> Object addArchiveBucket(inlineObject7)

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
    InlineObject7 inlineObject7 = new InlineObject7(); // InlineObject7 | 
    try {
      Object result = apiInstance.addArchiveBucket(inlineObject7);
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject7** | [**InlineObject7**](InlineObject7.md)|  | [optional]

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

<a name="addArchiveFile"></a>
# **addArchiveFile**
> Object addArchiveFile(bucket, filepath, filename, file)

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
    String filepath = "\"/\""; // String | The path to to search for files under. Default is the root directory /.
    String filename = "/path/to/file"; // String | Specify a filename for archive file. The base filename of the uploaded file is the default.
    File file = new File("/path/to/file"); // File | 
    try {
      Object result = apiInstance.addArchiveFile(bucket, filepath, filename, file);
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **String**| Bucket name |
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]
 **filename** | **String**| Specify a filename for archive file. The base filename of the uploaded file is the default. | [optional]
 **file** | **File**|  | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="addArchiveFileLink"></a>
# **addArchiveFileLink**
> Model200Success addArchiveFileLink(id, expireSeconds)

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
    Long expireSeconds = 0l; // Long | Time to live in seconds. 0 means do not expire.
    try {
      Model200Success result = apiInstance.addArchiveFileLink(id, expireSeconds);
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **expireSeconds** | **Long**| Time to live in seconds. 0 means do not expire. | [optional] [default to 0l]

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

<a name="deleteArchiveBucket"></a>
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

<a name="deleteArchiveFile"></a>
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

<a name="deleteArchiveFileLink"></a>
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **linkId** | **Long**| The ID of the archive file link. |

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

<a name="getArchiveBucket"></a>
# **getArchiveBucket**
> InlineResponse2004 getArchiveBucket(id)

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
      InlineResponse2004 result = apiInstance.getArchiveBucket(id);
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

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

<a name="getArchiveFile"></a>
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
    String filepath = "\"/\""; // String | The path to to search for files under. Default is the root directory /.
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **String**| Bucket name |
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]

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
**200** | Returns the contents of the specified file as an attachment with Content-Type dictated by the file |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getArchiveFileDetail"></a>
# **getArchiveFileDetail**
> InlineResponse2005 getArchiveFileDetail(id)

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
      InlineResponse2005 result = apiInstance.getArchiveFileDetail(id);
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

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

<a name="getArchiveFileLinks"></a>
# **getArchiveFileLinks**
> Object getArchiveFileLinks(id)

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
      Object result = apiInstance.getArchiveFileLinks(id);
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

<a name="getArchivePublicFile"></a>
# **getArchivePublicFile**
> getArchivePublicFile(bucket, filepath)

Download a Public Archive File

Files in an archive bucket that has Public URL enabled can be downloaded via this endpoint without any authentication, anonymously.

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
    String filepath = "\"/\""; // String | The path to to search for files under. Default is the root directory /.
    try {
      apiInstance.getArchivePublicFile(bucket, filepath);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#getArchivePublicFile");
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
 **bucket** | **String**| Bucket name |
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]

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
**200** | Returns the contents of the specified file as an attachment with Content-Type dicated by the file |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getArchivePublicFileLink"></a>
# **getArchivePublicFileLink**
> getArchivePublicFileLink(s)

Download an Archive File Link

Download an archive file link.

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
    String s = "45a214fce9a546b9"; // String | The secret access token for the archive file being downloaded.
    try {
      apiInstance.getArchivePublicFileLink(s);
    } catch (ApiException e) {
      System.err.println("Exception when calling ArchivesApi#getArchivePublicFileLink");
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
 **s** | **String**| The secret access token for the archive file being downloaded. |

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
**200** | Returns the contents of the specified file as an attachment with Content-Type dicated by the file |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listArchiveBuckets"></a>
# **listArchiveBuckets**
> Object listArchiveBuckets(name, phrase)

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
      Object result = apiInstance.listArchiveBuckets(name, phrase);
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]

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

<a name="listArchiveFiles"></a>
# **listArchiveFiles**
> Object listArchiveFiles(bucket, filepath, name, phrase, fullTree)

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
    String filepath = "\"/\""; // String | The path to to search for files under. Default is the root directory /.
    String name = "test-config.yaml"; // String | If specified will return an exact match on file name
    String phrase = "test-%.yaml"; // String | Search phrase for partial matches on file name, wildcard may be specified as %, eg. example-% This also searches for files under sub directories too. 
    Boolean fullTree = false; // Boolean | Include files under sub directories too. This is always true when searching with phrase.
    try {
      Object result = apiInstance.listArchiveFiles(bucket, filepath, name, phrase, fullTree);
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **String**| Bucket name |
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]
 **name** | **String**| If specified will return an exact match on file name | [optional]
 **phrase** | **String**| Search phrase for partial matches on file name, wildcard may be specified as %, eg. example-% This also searches for files under sub directories too.  | [optional]
 **fullTree** | **Boolean**| Include files under sub directories too. This is always true when searching with phrase. | [optional] [default to false]

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

<a name="updateArchiveBucket"></a>
# **updateArchiveBucket**
> Object updateArchiveBucket(id, inlineObject8)

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
    InlineObject8 inlineObject8 = new InlineObject8(); // InlineObject8 | 
    try {
      Object result = apiInstance.updateArchiveBucket(id, inlineObject8);
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject8** | [**InlineObject8**](InlineObject8.md)|  | [optional]

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

