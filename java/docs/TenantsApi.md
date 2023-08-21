# TenantsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addTenant**](TenantsApi.md#addTenant) | **POST** /api/accounts | Create a Tenant
[**addUserTenant**](TenantsApi.md#addUserTenant) | **POST** /api/accounts/{accountId}/users | Create a User For a Tenant
[**createTenantSubtenantGroup**](TenantsApi.md#createTenantSubtenantGroup) | **POST** /api/accounts/{accountId}/groups | Create a Group for Subtenant
[**getTenant**](TenantsApi.md#getTenant) | **GET** /api/accounts/{id} | Get tenant
[**getTenantSubtenantGroup**](TenantsApi.md#getTenantSubtenantGroup) | **GET** /api/accounts/{accountId}/groups/{id} | Get a Specific Group for Subtenant
[**listTenantSubtenantGroups**](TenantsApi.md#listTenantSubtenantGroups) | **GET** /api/accounts/{accountId}/groups | Get Subtenant Groups
[**listTenants**](TenantsApi.md#listTenants) | **GET** /api/accounts | List All Tenants
[**listTenantsAvailableRoles**](TenantsApi.md#listTenantsAvailableRoles) | **GET** /api/accounts/available-roles | List available roles for a tenant
[**removeTenant**](TenantsApi.md#removeTenant) | **DELETE** /api/accounts/{id} | Delete a Specific Tenant
[**removeTenantSubtenantGroup**](TenantsApi.md#removeTenantSubtenantGroup) | **DELETE** /api/accounts/{accountId}/groups/{id} | Delete a Group for Subtenant
[**updateTenant**](TenantsApi.md#updateTenant) | **PUT** /api/accounts/{id} | Update tenant
[**updateTenantSubtenantGroup**](TenantsApi.md#updateTenantSubtenantGroup) | **PUT** /api/accounts/{accountId}/groups/{id} | Updating a Group for Subtenant
[**updateTenantSubtenantGroupZones**](TenantsApi.md#updateTenantSubtenantGroupZones) | **PUT** /api/accounts/{accountId}/groups/{id}/update-zones | Updating Group Zones for Subtenant


<a name="addTenant"></a>
# **addTenant**
> Object addTenant(inlineObject238)

Create a Tenant

Create a new tenant. This new account will be a sub-tenant with the master tenant as its parent.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    InlineObject238 inlineObject238 = new InlineObject238(); // InlineObject238 | 
    try {
      Object result = apiInstance.addTenant(inlineObject238);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#addTenant");
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
 **inlineObject238** | [**InlineObject238**](InlineObject238.md)|  | [optional]

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
**200** | Tenant Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="addUserTenant"></a>
# **addUserTenant**
> Object addUserTenant(accountId, inlineObject243)

Create a User For a Tenant

Create a User For a Tenant.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    Long accountId = 7L; // Long | The ID of the subtenant account
    InlineObject243 inlineObject243 = new InlineObject243(); // InlineObject243 | 
    try {
      Object result = apiInstance.addUserTenant(accountId, inlineObject243);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#addUserTenant");
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
 **accountId** | **Long**| The ID of the subtenant account |
 **inlineObject243** | [**InlineObject243**](InlineObject243.md)|  | [optional]

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
**200** | User Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="createTenantSubtenantGroup"></a>
# **createTenantSubtenantGroup**
> InlineResponse200152 createTenantSubtenantGroup(accountId, inlineObject240)

Create a Group for Subtenant

Create a Group for Subtenant.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    Long accountId = 7L; // Long | The ID of the subtenant account
    InlineObject240 inlineObject240 = new InlineObject240(); // InlineObject240 | 
    try {
      InlineResponse200152 result = apiInstance.createTenantSubtenantGroup(accountId, inlineObject240);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#createTenantSubtenantGroup");
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
 **accountId** | **Long**| The ID of the subtenant account |
 **inlineObject240** | [**InlineObject240**](InlineObject240.md)|  | [optional]

### Return type

[**InlineResponse200152**](InlineResponse200152.md)

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

<a name="getTenant"></a>
# **getTenant**
> InlineResponse200150 getTenant(id)

Get tenant

Get details about a tenant

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200150 result = apiInstance.getTenant(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#getTenant");
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

[**InlineResponse200150**](InlineResponse200150.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tenant Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getTenantSubtenantGroup"></a>
# **getTenantSubtenantGroup**
> InlineResponse200153 getTenantSubtenantGroup(accountId, id)

Get a Specific Group for Subtenant

This endpoint retrieves a specific group.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    Long accountId = 7L; // Long | The ID of the subtenant account
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200153 result = apiInstance.getTenantSubtenantGroup(accountId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#getTenantSubtenantGroup");
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
 **accountId** | **Long**| The ID of the subtenant account |
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse200153**](InlineResponse200153.md)

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

<a name="listTenantSubtenantGroups"></a>
# **listTenantSubtenantGroups**
> Object listTenantSubtenantGroups(accountId, phrase, name, lastUpdated)

Get Subtenant Groups

Groups belonging to a subtenant can be managed by the master account.  This endpoint retrieves all groups and a list of zones associated with the group by id. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    Long accountId = 7L; // Long | The ID of the subtenant account
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    try {
      Object result = apiInstance.listTenantSubtenantGroups(accountId, phrase, name, lastUpdated);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#listTenantSubtenantGroups");
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
 **accountId** | **Long**| The ID of the subtenant account |
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
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

<a name="listTenants"></a>
# **listTenants**
> Object listTenants(max, offset, sort, direction, phrase, name, lastUpdated)

List All Tenants

Get a list of tenants. A tenant is also referred to as an account.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    try {
      Object result = apiInstance.listTenants(max, offset, sort, direction, phrase, name, lastUpdated);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#listTenants");
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
**200** | Array of Tenants |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listTenantsAvailableRoles"></a>
# **listTenantsAvailableRoles**
> TenantsAvailableRoles listTenantsAvailableRoles()

List available roles for a tenant

Get a list of available roles that can be assigned as the default base role for a sub tenant account.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    try {
      TenantsAvailableRoles result = apiInstance.listTenantsAvailableRoles();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#listTenantsAvailableRoles");
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

[**TenantsAvailableRoles**](TenantsAvailableRoles.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Array of roles |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="removeTenant"></a>
# **removeTenant**
> Model200Success removeTenant(id, removeResources)

Delete a Specific Tenant

Delete an existing tenant. This action is not reversible and will result in the removal of all data pertaining to this tenant as well as potentially any provisioned assets depending on the value of &#x60;removeResources&#x60;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Boolean removeResources = false; // Boolean | Remove Resources. This will delete all the managed resources in the tenant.
    try {
      Model200Success result = apiInstance.removeTenant(id, removeResources);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#removeTenant");
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
 **removeResources** | **Boolean**| Remove Resources. This will delete all the managed resources in the tenant. | [optional] [default to false]

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
**200** | Tenant Deletion Object |  -  |
**400** | Tenant still has managed resources |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="removeTenantSubtenantGroup"></a>
# **removeTenantSubtenantGroup**
> Model200Success removeTenantSubtenantGroup(accountId, id)

Delete a Group for Subtenant

If a group has zones or servers still tied to it, a delete action will fail.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    Long accountId = 7L; // Long | The ID of the subtenant account
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.removeTenantSubtenantGroup(accountId, id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#removeTenantSubtenantGroup");
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
 **accountId** | **Long**| The ID of the subtenant account |
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

<a name="updateTenant"></a>
# **updateTenant**
> InlineResponse200151 updateTenant(id, inlineObject239)

Update tenant

Update an existing tenant.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject239 inlineObject239 = new InlineObject239(); // InlineObject239 | 
    try {
      InlineResponse200151 result = apiInstance.updateTenant(id, inlineObject239);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#updateTenant");
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
 **inlineObject239** | [**InlineObject239**](InlineObject239.md)|  | [optional]

### Return type

[**InlineResponse200151**](InlineResponse200151.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tenant Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateTenantSubtenantGroup"></a>
# **updateTenantSubtenantGroup**
> InlineResponse200152 updateTenantSubtenantGroup(accountId, id, inlineObject241)

Updating a Group for Subtenant

Updating a Group for Subtenant.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    Long accountId = 7L; // Long | The ID of the subtenant account
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject241 inlineObject241 = new InlineObject241(); // InlineObject241 | 
    try {
      InlineResponse200152 result = apiInstance.updateTenantSubtenantGroup(accountId, id, inlineObject241);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#updateTenantSubtenantGroup");
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
 **accountId** | **Long**| The ID of the subtenant account |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject241** | [**InlineObject241**](InlineObject241.md)|  | [optional]

### Return type

[**InlineResponse200152**](InlineResponse200152.md)

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

<a name="updateTenantSubtenantGroupZones"></a>
# **updateTenantSubtenantGroupZones**
> Model200Success updateTenantSubtenantGroupZones(accountId, id, inlineObject242)

Updating Group Zones for Subtenant

This will update the zones that are assigned to the group. Any zones that are not passed in the zones parameter will be removed from the group.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.TenantsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    TenantsApi apiInstance = new TenantsApi(defaultClient);
    Long accountId = 7L; // Long | The ID of the subtenant account
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject242 inlineObject242 = new InlineObject242(); // InlineObject242 | 
    try {
      Model200Success result = apiInstance.updateTenantSubtenantGroupZones(accountId, id, inlineObject242);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling TenantsApi#updateTenantSubtenantGroupZones");
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
 **accountId** | **Long**| The ID of the subtenant account |
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject242** | [**InlineObject242**](InlineObject242.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

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

