# DeploymentsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addDeploymentFile**](DeploymentsApi.md#addDeploymentFile) | **POST** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | Upload a Deployment File
[**addDeploymentVersion**](DeploymentsApi.md#addDeploymentVersion) | **POST** /api/deployments/{deploymentId}/versions | Create a new Deployment Version
[**addDeployments**](DeploymentsApi.md#addDeployments) | **POST** /api/deployments | Create a new Deployment
[**deleteDeployment**](DeploymentsApi.md#deleteDeployment) | **DELETE** /api/deployments/{deploymentId} | Delete a Deployment
[**deleteDeploymentFile**](DeploymentsApi.md#deleteDeploymentFile) | **DELETE** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | Delete a Deployment File
[**deleteDeploymentVersion**](DeploymentsApi.md#deleteDeploymentVersion) | **DELETE** /api/deployments/{deploymentId}/versions/{id} | Delete a Deployment Version
[**getDeployment**](DeploymentsApi.md#getDeployment) | **GET** /api/deployments/{deploymentId} | Get a Specific Deployment
[**getDeploymentVersion**](DeploymentsApi.md#getDeploymentVersion) | **GET** /api/deployments/{deploymentId}/versions/{id} | Get a Specific Deployment Version
[**listDeploymentFiles**](DeploymentsApi.md#listDeploymentFiles) | **GET** /api/deployments/{deploymentId}/versions/{id}/files{filepath} | List Deployment Files
[**listDeploymentVersions**](DeploymentsApi.md#listDeploymentVersions) | **GET** /api/deployments/{deploymentId}/versions | Get All Versions For a Deployment
[**listDeployments**](DeploymentsApi.md#listDeployments) | **GET** /api/deployments | Get All Deployments
[**updateDeployment**](DeploymentsApi.md#updateDeployment) | **PUT** /api/deployments/{deploymentId} | Updating a Deployment
[**updateDeploymentVersion**](DeploymentsApi.md#updateDeploymentVersion) | **PUT** /api/deployments/{deploymentId}/versions/{id} | Updating a Deployment Version


<a name="addDeploymentFile"></a>
# **addDeploymentFile**
> Model200Success addDeploymentFile(deploymentId, id, filepath, file)

Upload a Deployment File

This endpoint will upload a file for a specific deployment version. This will overwrite the file if one with the same name exists already.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    Long deploymentId = 4L; // Long | Deployment ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String filepath = "\"/\""; // String | The path to to search for files under. Default is the root directory /.
    File file = new File("/path/to/file"); // File | 
    try {
      Model200Success result = apiInstance.addDeploymentFile(deploymentId, id, filepath, file);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#addDeploymentFile");
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
 **deploymentId** | **Long**| Deployment ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]
 **file** | **File**|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

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

<a name="addDeploymentVersion"></a>
# **addDeploymentVersion**
> Object addDeploymentVersion(deploymentId, inlineObject69)

Create a new Deployment Version

This endpoint will create a new deployment version that is ready to have files uploaded to it. The default type is file, which has files directly uploaded via Morpheus. Alternatively, the type git or fetch can be used to just point to a repository or remote url.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    Long deploymentId = 4L; // Long | Deployment ID
    InlineObject69 inlineObject69 = new InlineObject69(); // InlineObject69 | 
    try {
      Object result = apiInstance.addDeploymentVersion(deploymentId, inlineObject69);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#addDeploymentVersion");
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
 **deploymentId** | **Long**| Deployment ID |
 **inlineObject69** | [**InlineObject69**](InlineObject69.md)|  | [optional]

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

<a name="addDeployments"></a>
# **addDeployments**
> Object addDeployments(inlineObject67)

Create a new Deployment

This endpoint will create a new deployment that is ready to have versions added to it.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    InlineObject67 inlineObject67 = new InlineObject67(); // InlineObject67 | 
    try {
      Object result = apiInstance.addDeployments(inlineObject67);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#addDeployments");
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
 **inlineObject67** | [**InlineObject67**](InlineObject67.md)|  | [optional]

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

<a name="deleteDeployment"></a>
# **deleteDeployment**
> Model200Success deleteDeployment(deploymentId)

Delete a Deployment

This endpoint will delete an existing deployment.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    Long deploymentId = 4L; // Long | Deployment ID
    try {
      Model200Success result = apiInstance.deleteDeployment(deploymentId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#deleteDeployment");
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
 **deploymentId** | **Long**| Deployment ID |

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

<a name="deleteDeploymentFile"></a>
# **deleteDeploymentFile**
> Model200Success deleteDeploymentFile(deploymentId, id, filepath, force)

Delete a Deployment File

This endpoint will delete an existing deployment file. To recursively delete a directory and all of its contents, the force parameter must be specified.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    Long deploymentId = 4L; // Long | Deployment ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String filepath = "\"/\""; // String | The path to to search for files under. Default is the root directory /.
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.deleteDeploymentFile(deploymentId, id, filepath, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#deleteDeploymentFile");
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
 **deploymentId** | **Long**| Deployment ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]
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

<a name="deleteDeploymentVersion"></a>
# **deleteDeploymentVersion**
> Model200Success deleteDeploymentVersion(deploymentId, id)

Delete a Deployment Version

This endpoint will delete an existing deployment version.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    Long deploymentId = 4L; // Long | Deployment ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteDeploymentVersion(deploymentId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#deleteDeploymentVersion");
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
 **deploymentId** | **Long**| Deployment ID |
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

<a name="getDeployment"></a>
# **getDeployment**
> InlineResponse20038 getDeployment(deploymentId, maxVersions)

Get a Specific Deployment

This endpoint retrieves a specific deployment. By default the 5 most recent versions are returned, more can be returned by specifying the maxVersions parameter.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    Long deploymentId = 4L; // Long | Deployment ID
    Long maxVersions = 6L; // Long | Max number of recent versions to return.
    try {
      InlineResponse20038 result = apiInstance.getDeployment(deploymentId, maxVersions);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#getDeployment");
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
 **deploymentId** | **Long**| Deployment ID |
 **maxVersions** | **Long**| Max number of recent versions to return. | [optional]

### Return type

[**InlineResponse20038**](InlineResponse20038.md)

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

<a name="getDeploymentVersion"></a>
# **getDeploymentVersion**
> InlineResponse20039 getDeploymentVersion(deploymentId, id)

Get a Specific Deployment Version

This endpoint retrieves a specific deployment version.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    Long deploymentId = 4L; // Long | Deployment ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20039 result = apiInstance.getDeploymentVersion(deploymentId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#getDeploymentVersion");
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
 **deploymentId** | **Long**| Deployment ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20039**](InlineResponse20039.md)

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

<a name="listDeploymentFiles"></a>
# **listDeploymentFiles**
> Object listDeploymentFiles(deploymentId, id, filepath, max, offset, phrase, version)

List Deployment Files

This endpoint returns a list of files for a specific deployment version. This only applies to deploy type file. Files are sorted alphabetically, with directories appearing at the beginning of the list.  The filepath parameter can be specified to search for specific files or directories. To list files under a directory, use a trailing / in the filepath parameter.  To list a specific file, provide it&#39;s full path. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    Long deploymentId = 4L; // Long | Deployment ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String filepath = "\"/\""; // String | The path to to search for files under. Default is the root directory /.
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    Long version = 5L; // Long | Filter by version number (userVersion)
    try {
      Object result = apiInstance.listDeploymentFiles(deploymentId, id, filepath, max, offset, phrase, version);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#listDeploymentFiles");
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
 **deploymentId** | **Long**| Deployment ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **filepath** | **String**| The path to to search for files under. Default is the root directory /. | [default to &quot;/&quot;]
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **version** | **Long**| Filter by version number (userVersion) | [optional]

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

<a name="listDeploymentVersions"></a>
# **listDeploymentVersions**
> Object listDeploymentVersions(deploymentId, max, offset, phrase, version, type, dateCreated, lastUpdated)

Get All Versions For a Deployment

This endpoint returns a paginated list of versions for a specific deployment.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    Long deploymentId = 4L; // Long | Deployment ID
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    Long version = 5L; // Long | Filter by version number (userVersion)
    String type = "file"; // String | Filter by type (deployType), file, git, fetch
    String dateCreated = "2019-01-01"; // String | Filter by dateCreated, the created timestamp is more recent or equal to the date specified
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    try {
      Object result = apiInstance.listDeploymentVersions(deploymentId, max, offset, phrase, version, type, dateCreated, lastUpdated);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#listDeploymentVersions");
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
 **deploymentId** | **Long**| Deployment ID |
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **version** | **Long**| Filter by version number (userVersion) | [optional]
 **type** | **String**| Filter by type (deployType), file, git, fetch | [optional] [enum: file, git, fetch]
 **dateCreated** | **String**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional]
 **lastUpdated** | **OffsetDateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]

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

<a name="listDeployments"></a>
# **listDeployments**
> Object listDeployments(max, offset, phrase, name, description, dateCreated, lastUpdated)

Get All Deployments

This endpoint returns a paginated list of deployments.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String description = "The desription of my cool object"; // String | Filter by description, wildcard may be specified as %. eg. `example-%`
    String dateCreated = "2019-01-01"; // String | Filter by dateCreated, the created timestamp is more recent or equal to the date specified
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    try {
      Object result = apiInstance.listDeployments(max, offset, phrase, name, description, dateCreated, lastUpdated);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#listDeployments");
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
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **description** | **String**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]
 **dateCreated** | **String**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional]
 **lastUpdated** | **OffsetDateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]

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

<a name="updateDeployment"></a>
# **updateDeployment**
> Object updateDeployment(deploymentId, inlineObject68)

Updating a Deployment

This endpoint will update an existing deployment.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    Long deploymentId = 4L; // Long | Deployment ID
    InlineObject68 inlineObject68 = new InlineObject68(); // InlineObject68 | 
    try {
      Object result = apiInstance.updateDeployment(deploymentId, inlineObject68);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#updateDeployment");
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
 **deploymentId** | **Long**| Deployment ID |
 **inlineObject68** | [**InlineObject68**](InlineObject68.md)|  | [optional]

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

<a name="updateDeploymentVersion"></a>
# **updateDeploymentVersion**
> Object updateDeploymentVersion(deploymentId, id, inlineObject70)

Updating a Deployment Version

This endpoint will update an existing deployment version.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    Long deploymentId = 4L; // Long | Deployment ID
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject70 inlineObject70 = new InlineObject70(); // InlineObject70 | 
    try {
      Object result = apiInstance.updateDeploymentVersion(deploymentId, id, inlineObject70);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#updateDeploymentVersion");
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
 **deploymentId** | **Long**| Deployment ID |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject70** | [**InlineObject70**](InlineObject70.md)|  | [optional]

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

