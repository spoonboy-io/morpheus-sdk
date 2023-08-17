# CatalogItemsApi

All URIs are relative to *https://CHANGEME*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addCatalogItemType**](CatalogItemsApi.md#addCatalogItemType) | **POST** /api/catalog-item-types | Create a Catalog Item Type |
| [**getCatalogItemType**](CatalogItemsApi.md#getCatalogItemType) | **GET** /api/catalog-item-types/{id} | Get a Specific Catalog Item Type |
| [**listCatalogItemTypes**](CatalogItemsApi.md#listCatalogItemTypes) | **GET** /api/catalog-item-types | Get All Catalog Item Types |
| [**removeCatalogItemType**](CatalogItemsApi.md#removeCatalogItemType) | **DELETE** /api/catalog-item-types/{id} | Delete a Catalog Item Type |
| [**updateCatalogItemType**](CatalogItemsApi.md#updateCatalogItemType) | **PUT** /api/catalog-item-types/{id} | Update a Catalog Item Type |
| [**updateCatalogItemTypeLogo**](CatalogItemsApi.md#updateCatalogItemTypeLogo) | **PUT** /api/catalog-item-types/{id}/update-logo | Update Logo For Catalog Item Type |


<a id="addCatalogItemType"></a>
# **addCatalogItemType**
> AddCatalogItemType200Response addCatalogItemType(addCatalogItemTypeRequest)

Create a Catalog Item Type

Use this command to create a catalog item type.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CatalogItemsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CatalogItemsApi apiInstance = new CatalogItemsApi(defaultClient);
    AddCatalogItemTypeRequest addCatalogItemTypeRequest = new AddCatalogItemTypeRequest(); // AddCatalogItemTypeRequest | 
    try {
      AddCatalogItemType200Response result = apiInstance.addCatalogItemType(addCatalogItemTypeRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CatalogItemsApi#addCatalogItemType");
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
| **addCatalogItemTypeRequest** | [**AddCatalogItemTypeRequest**](AddCatalogItemTypeRequest.md)|  | [optional] |

### Return type

[**AddCatalogItemType200Response**](AddCatalogItemType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="getCatalogItemType"></a>
# **getCatalogItemType**
> GetCatalogItemType200Response getCatalogItemType(id)

Get a Specific Catalog Item Type

This endpoint retrieves a specific category item type.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CatalogItemsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CatalogItemsApi apiInstance = new CatalogItemsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      GetCatalogItemType200Response result = apiInstance.getCatalogItemType(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CatalogItemsApi#getCatalogItemType");
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

[**GetCatalogItemType200Response**](GetCatalogItemType200Response.md)

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

<a id="listCatalogItemTypes"></a>
# **listCatalogItemTypes**
> ListCatalogItemTypes200Response listCatalogItemTypes(max, offset, sort, direction, phrase, name, description, enabled, featured, labels, allLabels, code)

Get All Catalog Item Types

This endpoint retrieves all catalog item types.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CatalogItemsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CatalogItemsApi apiInstance = new CatalogItemsApi(defaultClient);
    Long max = 25L; // Long | Maximum number of records to return
    Long offset = 0L; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "name"; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String description = "The desription of my cool object"; // String | Filter by description, wildcard may be specified as %. eg. `example-%`
    Boolean enabled = false; // Boolean | Filter by enabled
    Boolean featured = false; // Boolean | Filter by featured
    String labels = "labels_example"; // String | Filter by label(s), matches records that contain any of the specified labels
    String allLabels = "allLabels_example"; // String | Filter by label(s), matches records that contain all of the specified labels
    String code = "azr"; // String | If specified will return an exact match on code
    try {
      ListCatalogItemTypes200Response result = apiInstance.listCatalogItemTypes(max, offset, sort, direction, phrase, name, description, enabled, featured, labels, allLabels, code);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CatalogItemsApi#listCatalogItemTypes");
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
| **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to name] |
| **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc] |
| **phrase** | **String**| Search phrase for partial matches on name or description | [optional] |
| **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional] |
| **description** | **String**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional] |
| **enabled** | **Boolean**| Filter by enabled | [optional] |
| **featured** | **Boolean**| Filter by featured | [optional] |
| **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional] |
| **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional] |
| **code** | **String**| If specified will return an exact match on code | [optional] |

### Return type

[**ListCatalogItemTypes200Response**](ListCatalogItemTypes200Response.md)

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

<a id="removeCatalogItemType"></a>
# **removeCatalogItemType**
> Model200Success removeCatalogItemType(id)

Delete a Catalog Item Type

Will delete a catalog item type.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CatalogItemsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CatalogItemsApi apiInstance = new CatalogItemsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.removeCatalogItemType(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CatalogItemsApi#removeCatalogItemType");
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

<a id="updateCatalogItemType"></a>
# **updateCatalogItemType**
> UpdateCatalogItemType200Response updateCatalogItemType(id, updateCatalogItemTypeRequest)

Update a Catalog Item Type

Use this command to update an existing catalog item type.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CatalogItemsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CatalogItemsApi apiInstance = new CatalogItemsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    UpdateCatalogItemTypeRequest updateCatalogItemTypeRequest = new UpdateCatalogItemTypeRequest(); // UpdateCatalogItemTypeRequest | 
    try {
      UpdateCatalogItemType200Response result = apiInstance.updateCatalogItemType(id, updateCatalogItemTypeRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CatalogItemsApi#updateCatalogItemType");
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
| **updateCatalogItemTypeRequest** | [**UpdateCatalogItemTypeRequest**](UpdateCatalogItemTypeRequest.md)|  | [optional] |

### Return type

[**UpdateCatalogItemType200Response**](UpdateCatalogItemType200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful Request |  -  |
| **4XX** | Error Codes |  -  |
| **5XX** | Error Codes |  -  |

<a id="updateCatalogItemTypeLogo"></a>
# **updateCatalogItemTypeLogo**
> UpdateCatalogItemType200Response updateCatalogItemTypeLogo(id, catalogItemTypeLogo, catalogItemTypeDarkLogo)

Update Logo For Catalog Item Type

Use this command to update the logo and dark logo images for an existing catalog item type. This endpoint expects multipart form data as the request format, not JSON.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.CatalogItemsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    CatalogItemsApi apiInstance = new CatalogItemsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    File catalogItemTypeLogo = new File("/path/to/file"); // File | Logo File png,jpg,svg
    File catalogItemTypeDarkLogo = new File("/path/to/file"); // File | Dark Logo File png,jpg,svg
    try {
      UpdateCatalogItemType200Response result = apiInstance.updateCatalogItemTypeLogo(id, catalogItemTypeLogo, catalogItemTypeDarkLogo);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling CatalogItemsApi#updateCatalogItemTypeLogo");
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
| **catalogItemTypeLogo** | **File**| Logo File png,jpg,svg | [optional] |
| **catalogItemTypeDarkLogo** | **File**| Dark Logo File png,jpg,svg | [optional] |

### Return type

[**UpdateCatalogItemType200Response**](UpdateCatalogItemType200Response.md)

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

