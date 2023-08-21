# DeploysApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addInstanceDeploy**](DeploysApi.md#addInstanceDeploy) | **POST** /api/instances/{id}/deploys | Deploy to an Instance
[**deletedeploy**](DeploysApi.md#deletedeploy) | **DELETE** /api/deploys/{id} | Delete a Deploy
[**getInstanceDeploys**](DeploysApi.md#getInstanceDeploys) | **GET** /api/instances/{id}/deploys | Get all Deploys for an Instance
[**listDeploys**](DeploysApi.md#listDeploys) | **GET** /api/deploys | Get all Deploys
[**runDeploy**](DeploysApi.md#runDeploy) | **POST** /api/deploys/{id}/deploy | Run a Deploy
[**updateDeploy**](DeploysApi.md#updateDeploy) | **PUT** /api/deploys/{id} | Update a Deploy


<a name="addInstanceDeploy"></a>
# **addInstanceDeploy**
> InlineResponse20040 addInstanceDeploy(id, inlineObject92)

Deploy to an Instance

This endpoint will deploy the specified deployment version to specified instance. The version to deploy can be identified with deploymentId and version or with versionId alone.  By default, the deployment is executed right away. To prevent this so that it can be run manually later on. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploysApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploysApi apiInstance = new DeploysApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject92 inlineObject92 = new InlineObject92(); // InlineObject92 | 
    try {
      InlineResponse20040 result = apiInstance.addInstanceDeploy(id, inlineObject92);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploysApi#addInstanceDeploy");
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
 **inlineObject92** | [**InlineObject92**](InlineObject92.md)|  | [optional]

### Return type

[**InlineResponse20040**](InlineResponse20040.md)

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

<a name="deletedeploy"></a>
# **deletedeploy**
> Model200Success deletedeploy(id)

Delete a Deploy

This endpoint will delete an archived instance deploy.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploysApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploysApi apiInstance = new DeploysApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deletedeploy(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploysApi#deletedeploy");
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

<a name="getInstanceDeploys"></a>
# **getInstanceDeploys**
> Object getInstanceDeploys(id, max, offset, phrase, name, deploymentId, instanceName, instanceId, version, versionId, createdById, deployType, dateCreated, lastUpdated, deployDate, status)

Get all Deploys for an Instance

This endpoint retrieves all deploys for a specific instance.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploysApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploysApi apiInstance = new DeploysApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    Long deploymentId = 5L; // Long | Filter by deployment id
    String instanceName = "instance1"; // String | Filter by instance name
    Long instanceId = 94L; // Long | The Instance ID for Filtering
    Long version = 5L; // Long | Filter by version number (userVersion)
    Long versionId = 5L; // Long | Filter by deployment version id
    Long createdById = 63L; // Long | Filter by owner (user) id
    String deployType = "file"; // String | Filter by type (deployType), file, git, fetch
    String dateCreated = "2019-01-01"; // String | Filter by dateCreated, the created timestamp is more recent or equal to the date specified
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    String deployDate = "2020-01-01"; // String | Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified
    String status = "deploying"; // String | Filter by status
    try {
      Object result = apiInstance.getInstanceDeploys(id, max, offset, phrase, name, deploymentId, instanceName, instanceId, version, versionId, createdById, deployType, dateCreated, lastUpdated, deployDate, status);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploysApi#getInstanceDeploys");
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
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **deploymentId** | **Long**| Filter by deployment id | [optional]
 **instanceName** | **String**| Filter by instance name | [optional]
 **instanceId** | **Long**| The Instance ID for Filtering | [optional]
 **version** | **Long**| Filter by version number (userVersion) | [optional]
 **versionId** | **Long**| Filter by deployment version id | [optional]
 **createdById** | **Long**| Filter by owner (user) id | [optional]
 **deployType** | **String**| Filter by type (deployType), file, git, fetch | [optional] [enum: file, git, fetch]
 **dateCreated** | **String**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional]
 **lastUpdated** | **OffsetDateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deployDate** | **String**| Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified | [optional]
 **status** | **String**| Filter by status | [optional] [enum: staged, deploying, commited, archived]

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

<a name="listDeploys"></a>
# **listDeploys**
> Object listDeploys(max, offset, phrase, name, deploymentId, instanceName, instanceId, version, versionId, createdById, deployType, dateCreated, lastUpdated, deployDate, status)

Get all Deploys

This endpoint retrieves all deploys.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploysApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploysApi apiInstance = new DeploysApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    Long deploymentId = 5L; // Long | Filter by deployment id
    String instanceName = "instance1"; // String | Filter by instance name
    Long instanceId = 94L; // Long | The Instance ID for Filtering
    Long version = 5L; // Long | Filter by version number (userVersion)
    Long versionId = 5L; // Long | Filter by deployment version id
    Long createdById = 63L; // Long | Filter by owner (user) id
    String deployType = "file"; // String | Filter by type (deployType), file, git, fetch
    String dateCreated = "2019-01-01"; // String | Filter by dateCreated, the created timestamp is more recent or equal to the date specified
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    String deployDate = "2020-01-01"; // String | Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified
    String status = "deploying"; // String | Filter by status
    try {
      Object result = apiInstance.listDeploys(max, offset, phrase, name, deploymentId, instanceName, instanceId, version, versionId, createdById, deployType, dateCreated, lastUpdated, deployDate, status);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploysApi#listDeploys");
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
 **deploymentId** | **Long**| Filter by deployment id | [optional]
 **instanceName** | **String**| Filter by instance name | [optional]
 **instanceId** | **Long**| The Instance ID for Filtering | [optional]
 **version** | **Long**| Filter by version number (userVersion) | [optional]
 **versionId** | **Long**| Filter by deployment version id | [optional]
 **createdById** | **Long**| Filter by owner (user) id | [optional]
 **deployType** | **String**| Filter by type (deployType), file, git, fetch | [optional] [enum: file, git, fetch]
 **dateCreated** | **String**| Filter by dateCreated, the created timestamp is more recent or equal to the date specified | [optional]
 **lastUpdated** | **OffsetDateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **deployDate** | **String**| Filter by deployDate, deployment completion timestamp is more recent or equal to the date specified | [optional]
 **status** | **String**| Filter by status | [optional] [enum: staged, deploying, commited, archived]

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

<a name="runDeploy"></a>
# **runDeploy**
> InlineResponse20040 runDeploy(id, inlineObject73)

Run a Deploy

This endpoint will run an existing instance deploy. This is for running a new staged deploy or to rollback to previous version by re-running a deploy that is archived.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploysApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploysApi apiInstance = new DeploysApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject73 inlineObject73 = new InlineObject73(); // InlineObject73 | 
    try {
      InlineResponse20040 result = apiInstance.runDeploy(id, inlineObject73);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploysApi#runDeploy");
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
 **inlineObject73** | [**InlineObject73**](InlineObject73.md)|  | [optional]

### Return type

[**InlineResponse20040**](InlineResponse20040.md)

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

<a name="updateDeploy"></a>
# **updateDeploy**
> InlineResponse20040 updateDeploy(id, inlineObject72)

Update a Deploy

This endpoint will update an existing deploy. This is typically only needed to change settings on a deploy that is staged, before it is run.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DeploysApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    DeploysApi apiInstance = new DeploysApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject72 inlineObject72 = new InlineObject72(); // InlineObject72 | 
    try {
      InlineResponse20040 result = apiInstance.updateDeploy(id, inlineObject72);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploysApi#updateDeploy");
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
 **inlineObject72** | [**InlineObject72**](InlineObject72.md)|  | [optional]

### Return type

[**InlineResponse20040**](InlineResponse20040.md)

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

