# BlueprintsApi

All URIs are relative to *https://CHANGEME*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addBlueprint**](BlueprintsApi.md#addBlueprint) | **POST** /api/blueprints | Create a Blueprint |
| [**deleteBlueprint**](BlueprintsApi.md#deleteBlueprint) | **DELETE** /api/blueprints/{id} | Delete a Blueprint |
| [**getBlueprint**](BlueprintsApi.md#getBlueprint) | **GET** /api/blueprints/{id} | Get a Specific Blueprint |
| [**listBlueprints**](BlueprintsApi.md#listBlueprints) | **GET** /api/blueprints | Get All Blueprints |
| [**updateBlueprint**](BlueprintsApi.md#updateBlueprint) | **PUT** /api/blueprints/{id} | Updating a Blueprint |
| [**updateBlueprintImage**](BlueprintsApi.md#updateBlueprintImage) | **POST** /api/blueprints/{id}/image | Update Blueprint Image |
| [**updateBlueprintPermissions**](BlueprintsApi.md#updateBlueprintPermissions) | **PUT** /api/blueprints/{id}/update-permissions | Update Blueprint Permissions |


<a id="addBlueprint"></a>
# **addBlueprint**
> AddBlueprint200Response addBlueprint(addBlueprintRequest)

Create a Blueprint

Create a Blueprint

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    AddBlueprintRequest addBlueprintRequest = new AddBlueprintRequest(); // AddBlueprintRequest | 
    try {
      AddBlueprint200Response result = apiInstance.addBlueprint(addBlueprintRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#addBlueprint");
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
| **addBlueprintRequest** | [**AddBlueprintRequest**](AddBlueprintRequest.md)|  | [optional] |

### Return type

[**AddBlueprint200Response**](AddBlueprint200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="deleteBlueprint"></a>
# **deleteBlueprint**
> Model200Success deleteBlueprint(id)

Delete a Blueprint

Delete a Blueprint

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteBlueprint(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#deleteBlueprint");
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

<a id="getBlueprint"></a>
# **getBlueprint**
> GetBlueprint200Response getBlueprint(id)

Get a Specific Blueprint

This endpoint retrieves a specific blueprint.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      GetBlueprint200Response result = apiInstance.getBlueprint(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#getBlueprint");
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

[**GetBlueprint200Response**](GetBlueprint200Response.md)

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

<a id="listBlueprints"></a>
# **listBlueprints**
> ListBlueprints200Response listBlueprints(max, offset, name, phrase, labels, allLabels)

Get All Blueprints

This endpoint retrieves all blueprints.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    Long max = 25L; // Long | Maximum number of records to return
    Long offset = 0L; // Long | Offset records, the number of records to skip, for paginating requests
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String labels = "labels_example"; // String | Filter by label(s), matches records that contain any of the specified labels
    String allLabels = "allLabels_example"; // String | Filter by label(s), matches records that contain all of the specified labels
    try {
      ListBlueprints200Response result = apiInstance.listBlueprints(max, offset, name, phrase, labels, allLabels);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#listBlueprints");
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
| **max** | **Long**| Maximum number of records to return | [optional] [default to 25] |
| **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0] |
| **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] |
| **phrase** | **String**| Search phrase for partial matches on name or description | [optional] |
| **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] |
| **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] |

### Return type

[**ListBlueprints200Response**](ListBlueprints200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Response |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="updateBlueprint"></a>
# **updateBlueprint**
> GetBlueprint200Response updateBlueprint(id, addBlueprintRequest)

Updating a Blueprint

Update a Blueprint. This overwrites the entire config, so the entire blueprint config should be passed.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    AddBlueprintRequest addBlueprintRequest = new AddBlueprintRequest(); // AddBlueprintRequest | 
    try {
      GetBlueprint200Response result = apiInstance.updateBlueprint(id, addBlueprintRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#updateBlueprint");
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
| **addBlueprintRequest** | [**AddBlueprintRequest**](AddBlueprintRequest.md)|  | [optional] |

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md)

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

<a id="updateBlueprintImage"></a>
# **updateBlueprintImage**
> GetBlueprint200Response updateBlueprintImage(id, templateImage)

Update Blueprint Image

Update Blueprint Image

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    File templateImage = new File("/path/to/file"); // File | 
    try {
      GetBlueprint200Response result = apiInstance.updateBlueprintImage(id, templateImage);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#updateBlueprintImage");
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
| **templateImage** | **File**|  | [optional] |

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md)

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

<a id="updateBlueprintPermissions"></a>
# **updateBlueprintPermissions**
> GetBlueprint200Response updateBlueprintPermissions(id, updateBlueprintPermissionsRequest)

Update Blueprint Permissions

Update Blueprint Permissions

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    UpdateBlueprintPermissionsRequest updateBlueprintPermissionsRequest = new UpdateBlueprintPermissionsRequest(); // UpdateBlueprintPermissionsRequest | 
    try {
      GetBlueprint200Response result = apiInstance.updateBlueprintPermissions(id, updateBlueprintPermissionsRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#updateBlueprintPermissions");
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
| **updateBlueprintPermissionsRequest** | [**UpdateBlueprintPermissionsRequest**](UpdateBlueprintPermissionsRequest.md)|  | [optional] |

### Return type

[**GetBlueprint200Response**](GetBlueprint200Response.md)

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

