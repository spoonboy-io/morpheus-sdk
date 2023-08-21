# IntegrationsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addIntegrationSnowObjects**](IntegrationsApi.md#addIntegrationSnowObjects) | **POST** /api/integrations/{id}/objects | Creates an Exposed ServiceNow Catalog Item
[**addIntegrations**](IntegrationsApi.md#addIntegrations) | **POST** /api/integrations | Creates an Integration
[**getIntegrationInventory**](IntegrationsApi.md#getIntegrationInventory) | **GET** /api/integrations/{id}/inventory/{inventoryId} | Get a Specific Integration Inventory
[**getIntegrationObjects**](IntegrationsApi.md#getIntegrationObjects) | **GET** /api/integrations/{id}/objects/{objectId} | Get a Specific ServiceNow Integration Object
[**getIntegrationTypeOptionTypes**](IntegrationsApi.md#getIntegrationTypeOptionTypes) | **GET** /api/integration-types/{id}/option-types | Retrieves a Option Types for a Specific Integration Type
[**getIntegrationTypes**](IntegrationsApi.md#getIntegrationTypes) | **GET** /api/integration-types/{id} | Retrieves a Specific Integration Type
[**getIntegrations**](IntegrationsApi.md#getIntegrations) | **GET** /api/integrations/{id} | Retrieves a Specific Integration
[**listIntegrationInventory**](IntegrationsApi.md#listIntegrationInventory) | **GET** /api/integrations/{id}/inventory | Get All Integration Inventory
[**listIntegrationObjects**](IntegrationsApi.md#listIntegrationObjects) | **GET** /api/integrations/{id}/objects | Get ServiceNow Integration Objects
[**listIntegrationTypes**](IntegrationsApi.md#listIntegrationTypes) | **GET** /api/integration-types | Retrieves all Integration Types
[**listIntegrations**](IntegrationsApi.md#listIntegrations) | **GET** /api/integrations | Retrieves all Integrations
[**refreshIntegrations**](IntegrationsApi.md#refreshIntegrations) | **POST** /api/integrations/{id}/refresh | Refresh an Integration
[**removeIntegrationObjects**](IntegrationsApi.md#removeIntegrationObjects) | **DELETE** /api/integrations/{id}/objects/{objectId} | Deletes a ServiceNow Integration object
[**removeIntegrations**](IntegrationsApi.md#removeIntegrations) | **DELETE** /api/integrations/{id} | Deletes an Integration
[**updateIntegrationInventory**](IntegrationsApi.md#updateIntegrationInventory) | **PUT** /api/integrations/{id}/inventory/{inventoryId} | Updating an Integration Inventory
[**updateIntegrations**](IntegrationsApi.md#updateIntegrations) | **PUT** /api/integrations/{id} | Updates an Integration


<a name="addIntegrationSnowObjects"></a>
# **addIntegrationSnowObjects**
> Object addIntegrationSnowObjects(id, inlineObject100)

Creates an Exposed ServiceNow Catalog Item

This endpoint creates an Exposed Catalog Item. This is an integration object of type &#x60;catalog&#x60; that references a &#x60;Catalog Item Type.&#x60; 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject100 inlineObject100 = new InlineObject100(); // InlineObject100 | 
    try {
      Object result = apiInstance.addIntegrationSnowObjects(id, inlineObject100);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#addIntegrationSnowObjects");
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
 **inlineObject100** | [**InlineObject100**](InlineObject100.md)|  | [optional]

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

<a name="addIntegrations"></a>
# **addIntegrations**
> Object addIntegrations(UNKNOWN_BASE_TYPE)

Creates an Integration

Creates an integration. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE = new UNKNOWN_BASE_TYPE(); // UNKNOWN_BASE_TYPE | 
    try {
      Object result = apiInstance.addIntegrations(UNKNOWN_BASE_TYPE);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#addIntegrations");
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional]

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

<a name="getIntegrationInventory"></a>
# **getIntegrationInventory**
> InlineResponse20065 getIntegrationInventory(id, inventoryId)

Get a Specific Integration Inventory

This endpoint retrieves a specific integration inventory. Only certain types of integrations support this operation, such as Ansible Tower. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 9L; // Long | Morpheus ID of the Integration
    Long inventoryId = 1L; // Long | Morpheus ID of the Integration Inventory
    try {
      InlineResponse20065 result = apiInstance.getIntegrationInventory(id, inventoryId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#getIntegrationInventory");
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
 **id** | **Long**| Morpheus ID of the Integration |
 **inventoryId** | **Long**| Morpheus ID of the Integration Inventory |

### Return type

[**InlineResponse20065**](InlineResponse20065.md)

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

<a name="getIntegrationObjects"></a>
# **getIntegrationObjects**
> InlineResponse20064 getIntegrationObjects(id, objectId)

Get a Specific ServiceNow Integration Object

This endpoint retrieves a specific ServiceNow integration object.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Long objectId = 122L; // Long | Morpheus ID of the Object being created or referenced
    try {
      InlineResponse20064 result = apiInstance.getIntegrationObjects(id, objectId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#getIntegrationObjects");
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
 **objectId** | **Long**| Morpheus ID of the Object being created or referenced |

### Return type

[**InlineResponse20064**](InlineResponse20064.md)

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

<a name="getIntegrationTypeOptionTypes"></a>
# **getIntegrationTypeOptionTypes**
> InlineResponse20062 getIntegrationTypeOptionTypes(id)

Retrieves a Option Types for a Specific Integration Type

This endpoint will retrieve the list of option types for a specific integration type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20062 result = apiInstance.getIntegrationTypeOptionTypes(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#getIntegrationTypeOptionTypes");
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

[**InlineResponse20062**](InlineResponse20062.md)

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

<a name="getIntegrationTypes"></a>
# **getIntegrationTypes**
> InlineResponse20061 getIntegrationTypes(id, optiontypes)

Retrieves a Specific Integration Type

Retrieves a specific integration type. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Boolean optiontypes = false; // Boolean | Pass `true` to include optionTypes in the response for each integration type
    try {
      InlineResponse20061 result = apiInstance.getIntegrationTypes(id, optiontypes);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#getIntegrationTypes");
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
 **optiontypes** | **Boolean**| Pass &#x60;true&#x60; to include optionTypes in the response for each integration type | [optional] [default to false]

### Return type

[**InlineResponse20061**](InlineResponse20061.md)

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

<a name="getIntegrations"></a>
# **getIntegrations**
> Object getIntegrations(id)

Retrieves a Specific Integration

Retrieves a specific integration. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.getIntegrations(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#getIntegrations");
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

<a name="listIntegrationInventory"></a>
# **listIntegrationInventory**
> Object listIntegrationInventory(id, max, offset, sort, direction, phrase, name)

Get All Integration Inventory

This endpoint retrieves a list of inventory for a specific integration. Only certain types of integrations support this operation, such as Ansible Tower. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 9L; // Long | Morpheus ID of the Integration
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    try {
      Object result = apiInstance.listIntegrationInventory(id, max, offset, sort, direction, phrase, name);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#listIntegrationInventory");
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
 **id** | **Long**| Morpheus ID of the Integration |
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

<a name="listIntegrationObjects"></a>
# **listIntegrationObjects**
> InlineResponse20063 listIntegrationObjects(id, max, offset, sort, direction, phrase, name, value, refId)

Get ServiceNow Integration Objects

This endpoint retrieves a list of exposed objects for a specific ServiceNow integration. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String value = "value_example"; // String | The type of data being stored, string or object. The data type depends on the cypher mount being used. Most mounts use string as their data type, but secret uses object by default.  You can store a string instead by passing `type=string`. This means the data value returned by the API will be a string instead of an object. 
    Long refId = 3L; // Long | If specified will return an exact match on refId
    try {
      InlineResponse20063 result = apiInstance.listIntegrationObjects(id, max, offset, sort, direction, phrase, name, value, refId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#listIntegrationObjects");
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
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &quot;name&quot;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to asc] [enum: asc, desc]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **value** | **String**| The type of data being stored, string or object. The data type depends on the cypher mount being used. Most mounts use string as their data type, but secret uses object by default.  You can store a string instead by passing &#x60;type&#x3D;string&#x60;. This means the data value returned by the API will be a string instead of an object.  | [optional] [enum: string, object]
 **refId** | **Long**| If specified will return an exact match on refId | [optional]

### Return type

[**InlineResponse20063**](InlineResponse20063.md)

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

<a name="listIntegrationTypes"></a>
# **listIntegrationTypes**
> Object listIntegrationTypes(max, offset, sort, direction, phrase, name, code, optiontypes, description, category, creatable, enabled, hasCMDB, hasCMDBDiscovery, hasCM, hasDNS, hasApprovals, isPlugin)

Retrieves all Integration Types

An Integration Type is specific third party software that the appliance can be configured to integrate with, such as Ansible, Chef, Git, ServiceNOW, etc. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String code = "azr"; // String | If specified will return an exact match on code
    Boolean optiontypes = false; // Boolean | Pass `true` to include optionTypes in the response for each integration type
    String description = "The desription of my cool object"; // String | Filter by description, wildcard may be specified as %. eg. `example-%`
    String category = "category_example"; // String | If specified will return an exact match on category
    Boolean creatable = true; // Boolean | Filter by creatable
    Boolean enabled = false; // Boolean | Filter by enabled
    Boolean hasCMDB = true; // Boolean | Filter by hasCMDB
    Boolean hasCMDBDiscovery = true; // Boolean | Filter by hasCMDBDiscovery
    Boolean hasCM = true; // Boolean | Filter by hasCM
    Boolean hasDNS = true; // Boolean | Filter by hasDNS
    Boolean hasApprovals = true; // Boolean | Filter by hasApprovals
    Boolean isPlugin = true; // Boolean | Filter by isPlugin
    try {
      Object result = apiInstance.listIntegrationTypes(max, offset, sort, direction, phrase, name, code, optiontypes, description, category, creatable, enabled, hasCMDB, hasCMDBDiscovery, hasCM, hasDNS, hasApprovals, isPlugin);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#listIntegrationTypes");
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
 **code** | **String**| If specified will return an exact match on code | [optional]
 **optiontypes** | **Boolean**| Pass &#x60;true&#x60; to include optionTypes in the response for each integration type | [optional] [default to false]
 **description** | **String**| Filter by description, wildcard may be specified as %. eg. &#x60;example-%&#x60; | [optional]
 **category** | **String**| If specified will return an exact match on category | [optional]
 **creatable** | **Boolean**| Filter by creatable | [optional]
 **enabled** | **Boolean**| Filter by enabled | [optional]
 **hasCMDB** | **Boolean**| Filter by hasCMDB | [optional]
 **hasCMDBDiscovery** | **Boolean**| Filter by hasCMDBDiscovery | [optional]
 **hasCM** | **Boolean**| Filter by hasCM | [optional]
 **hasDNS** | **Boolean**| Filter by hasDNS | [optional]
 **hasApprovals** | **Boolean**| Filter by hasApprovals | [optional]
 **isPlugin** | **Boolean**| Filter by isPlugin | [optional]

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

<a name="listIntegrations"></a>
# **listIntegrations**
> AnyOfobjectmeta listIntegrations(max, offset, sort, direction, phrase, name, id, url, host, port, username, version, windowsVersion, status, objects)

Retrieves all Integrations

Retrieves all integrations. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    Long id = 7L; // Long | Morpheus ID of the Object being created or referenced
    String url = "https://example.com/testimage.ovf"; // String | Download the file from a remote url. This can be used instead of uploading a local file.
    String host = "host_example"; // String | Filter by integration Host
    String port = "port_example"; // String | Filter by integration Port
    String username = "administrator"; // String | Username
    Long version = 5L; // Long | Filter by version number (userVersion)
    String windowsVersion = "windowsVersion_example"; // String | Filter by integration Windows Version
    String status = "running"; // String | The instance status for filtering.
    Boolean objects = false; // Boolean | Include `objects=true` to return the Integration Objects for each integration.  Available in api version 5.2.8/5.3.2. 
    try {
      AnyOfobjectmeta result = apiInstance.listIntegrations(max, offset, sort, direction, phrase, name, id, url, host, port, username, version, windowsVersion, status, objects);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#listIntegrations");
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
 **id** | **Long**| Morpheus ID of the Object being created or referenced | [optional]
 **url** | **String**| Download the file from a remote url. This can be used instead of uploading a local file. | [optional]
 **host** | **String**| Filter by integration Host | [optional]
 **port** | **String**| Filter by integration Port | [optional]
 **username** | **String**| Username | [optional]
 **version** | **Long**| Filter by version number (userVersion) | [optional]
 **windowsVersion** | **String**| Filter by integration Windows Version | [optional]
 **status** | **String**| The instance status for filtering. | [optional]
 **objects** | **Boolean**| Include &#x60;objects&#x3D;true&#x60; to return the Integration Objects for each integration.  Available in api version 5.2.8/5.3.2.  | [optional] [default to false]

### Return type

[**AnyOfobjectmeta**](AnyOfobjectmeta.md)

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

<a name="refreshIntegrations"></a>
# **refreshIntegrations**
> Object refreshIntegrations(id)

Refresh an Integration

This endpoint will refresh an existing Integration. Only some types support this and will actually perform an action as a result. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Object result = apiInstance.refreshIntegrations(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#refreshIntegrations");
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

<a name="removeIntegrationObjects"></a>
# **removeIntegrationObjects**
> Model200Success removeIntegrationObjects(id, objectId)

Deletes a ServiceNow Integration object

Deletes a specified ServiceNow integration object. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Long objectId = 122L; // Long | Morpheus ID of the Object being created or referenced
    try {
      Model200Success result = apiInstance.removeIntegrationObjects(id, objectId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#removeIntegrationObjects");
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
 **objectId** | **Long**| Morpheus ID of the Object being created or referenced |

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

<a name="removeIntegrations"></a>
# **removeIntegrations**
> Model200Success removeIntegrations(id)

Deletes an Integration

Deletes a specified integration. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.removeIntegrations(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#removeIntegrations");
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

<a name="updateIntegrationInventory"></a>
# **updateIntegrationInventory**
> Object updateIntegrationInventory(id, inventoryId, inlineObject101)

Updating an Integration Inventory

This endpoint provides updating of integration inventory

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 9L; // Long | Morpheus ID of the Integration
    Long inventoryId = 1L; // Long | Morpheus ID of the Integration Inventory
    InlineObject101 inlineObject101 = new InlineObject101(); // InlineObject101 | 
    try {
      Object result = apiInstance.updateIntegrationInventory(id, inventoryId, inlineObject101);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#updateIntegrationInventory");
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
 **id** | **Long**| Morpheus ID of the Integration |
 **inventoryId** | **Long**| Morpheus ID of the Integration Inventory |
 **inlineObject101** | [**InlineObject101**](InlineObject101.md)|  | [optional]

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

<a name="updateIntegrations"></a>
# **updateIntegrations**
> Object updateIntegrations(id, UNKNOWN_BASE_TYPE)

Updates an Integration

Updates an integration. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.IntegrationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    IntegrationsApi apiInstance = new IntegrationsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE = new UNKNOWN_BASE_TYPE(); // UNKNOWN_BASE_TYPE | 
    try {
      Object result = apiInstance.updateIntegrations(id, UNKNOWN_BASE_TYPE);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling IntegrationsApi#updateIntegrations");
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
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional]

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

