# ServiceCatalogApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCatalogCart**](ServiceCatalogApi.md#addCatalogCart) | **POST** /api/catalog/checkout | Checkout Catalog Cart
[**addCatalogCartItem**](ServiceCatalogApi.md#addCatalogCartItem) | **PUT** /api/catalog/cart/items | Add Catalog Item to Cart
[**addCatalogOrder**](ServiceCatalogApi.md#addCatalogOrder) | **POST** /api/catalog/orders | Place Catalog Order
[**deleteCatalogCart**](ServiceCatalogApi.md#deleteCatalogCart) | **DELETE** /api/catalog/cart | Clear Catalog Cart
[**deleteCatalogCartItem**](ServiceCatalogApi.md#deleteCatalogCartItem) | **DELETE** /api/catalog/cart/items/{id} | Remove a Catalog Item From Cart
[**deleteCatalogItem**](ServiceCatalogApi.md#deleteCatalogItem) | **DELETE** /api/catalog/items/{id} | Delete a Catalog Inventory Item
[**getCatalogItem**](ServiceCatalogApi.md#getCatalogItem) | **GET** /api/catalog/items/{id} | Get a Specific Catalog Inventory Item
[**getCatalogType**](ServiceCatalogApi.md#getCatalogType) | **GET** /api/catalog/types/{id} | Get a Specific Catalog Type
[**listCatalogCart**](ServiceCatalogApi.md#listCatalogCart) | **GET** /api/catalog/cart | Get Catalog Cart
[**listCatalogItems**](ServiceCatalogApi.md#listCatalogItems) | **GET** /api/catalog/items | List Catalog Inventory Items
[**listCatalogTypes**](ServiceCatalogApi.md#listCatalogTypes) | **GET** /api/catalog/types | List Catalog Types


<a name="addCatalogCart"></a>
# **addCatalogCart**
> Object addCatalogCart()

Checkout Catalog Cart

Use this command to checkout, finalizing your cart and placing an order. This converts each item in the cart to an inventory item, changing the status from IN_CART to ORDERED and potentially starts the provisioning process for each item.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ServiceCatalogApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ServiceCatalogApi apiInstance = new ServiceCatalogApi(defaultClient);
    try {
      Object result = apiInstance.addCatalogCart();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ServiceCatalogApi#addCatalogCart");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

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

<a name="addCatalogCartItem"></a>
# **addCatalogCartItem**
> Object addCatalogCartItem(validate, catalogCartItemCreate)

Add Catalog Item to Cart

Use this command to add an item to your service catalog cart.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ServiceCatalogApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ServiceCatalogApi apiInstance = new ServiceCatalogApi(defaultClient);
    Boolean validate = false; // Boolean | Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory
    CatalogCartItemCreate catalogCartItemCreate = {"item":{"type":{"name":"example"},"config":{"appName":"My App"}}}; // CatalogCartItemCreate | 
    try {
      Object result = apiInstance.addCatalogCartItem(validate, catalogCartItemCreate);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ServiceCatalogApi#addCatalogCartItem");
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
 **validate** | **Boolean**| Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory | [optional] [default to false]
 **catalogCartItemCreate** | [**CatalogCartItemCreate**](CatalogCartItemCreate.md)|  | [optional]

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

<a name="addCatalogOrder"></a>
# **addCatalogOrder**
> Object addCatalogOrder(validate, inlineObject227)

Place Catalog Order

This will place an order for the specified items, adding items to the inventory right away, without using the cart.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ServiceCatalogApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ServiceCatalogApi apiInstance = new ServiceCatalogApi(defaultClient);
    Boolean validate = false; // Boolean | Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory
    InlineObject227 inlineObject227 = new InlineObject227(); // InlineObject227 | 
    try {
      Object result = apiInstance.addCatalogOrder(validate, inlineObject227);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ServiceCatalogApi#addCatalogOrder");
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
 **validate** | **Boolean**| Validate Only. Use true to validate the configuration without actually placing the order or adding items to your inventory | [optional] [default to false]
 **inlineObject227** | [**InlineObject227**](InlineObject227.md)|  | [optional]

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

<a name="deleteCatalogCart"></a>
# **deleteCatalogCart**
> Model200Success deleteCatalogCart()

Clear Catalog Cart

Use this command to empty your cart, deleting all the items in it.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ServiceCatalogApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ServiceCatalogApi apiInstance = new ServiceCatalogApi(defaultClient);
    try {
      Model200Success result = apiInstance.deleteCatalogCart();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ServiceCatalogApi#deleteCatalogCart");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

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

<a name="deleteCatalogCartItem"></a>
# **deleteCatalogCartItem**
> Model200Success deleteCatalogCartItem(id)

Remove a Catalog Item From Cart

Will remove a catalog item that is currently in the cart.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ServiceCatalogApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ServiceCatalogApi apiInstance = new ServiceCatalogApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteCatalogCartItem(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ServiceCatalogApi#deleteCatalogCartItem");
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

<a name="deleteCatalogItem"></a>
# **deleteCatalogItem**
> Model200Success deleteCatalogItem(id, preserveVolumes, keepBackups, releaseFloatingIps, releaseEIPs, removeInstances, force)

Delete a Catalog Inventory Item

Will delete a catalog inventory item, which by default will deprovision any associated any instances and servers.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ServiceCatalogApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ServiceCatalogApi apiInstance = new ServiceCatalogApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    String preserveVolumes = "\"off\""; // String | Preserve Volumes
    String keepBackups = "\"off\""; // String | Preserve copy of backups
    String releaseFloatingIps = "\"on\""; // String | Release Floating IPs
    String releaseEIPs = "\"on\""; // String | Alias for releaseFloatingIps
    String removeInstances = "\"on\""; // String | Remove Instances. Only applies to type `blueprint` (Apps)
    String force = "\"off\""; // String | Force Delete
    try {
      Model200Success result = apiInstance.deleteCatalogItem(id, preserveVolumes, keepBackups, releaseFloatingIps, releaseEIPs, removeInstances, force);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ServiceCatalogApi#deleteCatalogItem");
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
 **preserveVolumes** | **String**| Preserve Volumes | [optional] [default to &quot;off&quot;]
 **keepBackups** | **String**| Preserve copy of backups | [optional] [default to &quot;off&quot;]
 **releaseFloatingIps** | **String**| Release Floating IPs | [optional] [default to &quot;on&quot;]
 **releaseEIPs** | **String**| Alias for releaseFloatingIps | [optional] [default to &quot;on&quot;]
 **removeInstances** | **String**| Remove Instances. Only applies to type &#x60;blueprint&#x60; (Apps) | [optional] [default to &quot;on&quot;]
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

<a name="getCatalogItem"></a>
# **getCatalogItem**
> InlineResponse200140 getCatalogItem(id)

Get a Specific Catalog Inventory Item

This endpoint retrieves a specific catalog inventory item.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ServiceCatalogApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ServiceCatalogApi apiInstance = new ServiceCatalogApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200140 result = apiInstance.getCatalogItem(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ServiceCatalogApi#getCatalogItem");
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

[**InlineResponse200140**](InlineResponse200140.md)

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

<a name="getCatalogType"></a>
# **getCatalogType**
> Object getCatalogType(id)

Get a Specific Catalog Type

This endpoint retrieves a specific catalog item type. This also returns an array of associated optionTypes that are used to configure the catalog item.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ServiceCatalogApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ServiceCatalogApi apiInstance = new ServiceCatalogApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getCatalogType(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ServiceCatalogApi#getCatalogType");
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

<a name="listCatalogCart"></a>
# **listCatalogCart**
> InlineResponse200139 listCatalogCart()

Get Catalog Cart

This endpoint retrieves the current catalog cart and all the items in it.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ServiceCatalogApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ServiceCatalogApi apiInstance = new ServiceCatalogApi(defaultClient);
    try {
      InlineResponse200139 result = apiInstance.listCatalogCart();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ServiceCatalogApi#listCatalogCart");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200139**](InlineResponse200139.md)

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

<a name="listCatalogItems"></a>
# **listCatalogItems**
> Object listCatalogItems(max, offset, sort, direction, phrase, name)

List Catalog Inventory Items

This endpoint retrieves a list of the catalog inventory items.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ServiceCatalogApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ServiceCatalogApi apiInstance = new ServiceCatalogApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    try {
      Object result = apiInstance.listCatalogItems(max, offset, sort, direction, phrase, name);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ServiceCatalogApi#listCatalogItems");
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
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]

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

<a name="listCatalogTypes"></a>
# **listCatalogTypes**
> Object listCatalogTypes(max, offset, sort, direction, phrase, name, featured)

List Catalog Types

This endpoint retrieves the types available for ordering.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ServiceCatalogApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ServiceCatalogApi apiInstance = new ServiceCatalogApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    Boolean featured = false; // Boolean | Filter by featured
    try {
      Object result = apiInstance.listCatalogTypes(max, offset, sort, direction, phrase, name, featured);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ServiceCatalogApi#listCatalogTypes");
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
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **featured** | **Boolean**| Filter by featured | [optional]

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

