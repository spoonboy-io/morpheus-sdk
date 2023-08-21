# SecurityGroupsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addSecurityGroupLocations**](SecurityGroupsApi.md#addSecurityGroupLocations) | **POST** /api/security-groups/{id}/locations | Creates a Security Group Location
[**addSecurityGroupRules**](SecurityGroupsApi.md#addSecurityGroupRules) | **POST** /api/security-groups/{id}/rules | Creates a Security Group Rule
[**addSecurityGroups**](SecurityGroupsApi.md#addSecurityGroups) | **POST** /api/security-groups | Creates a Security Group
[**getSecurityGroupRules**](SecurityGroupsApi.md#getSecurityGroupRules) | **GET** /api/security-groups/{id}/rules/{sgId} | Retrieves a Specific Security Group Rule
[**getSecurityGroups**](SecurityGroupsApi.md#getSecurityGroups) | **GET** /api/security-groups/{id} | Retrieves a Specific Security Group
[**listSecurityGroupRules**](SecurityGroupsApi.md#listSecurityGroupRules) | **GET** /api/security-groups/{id}/rules | Retrieves all Security Group Rules
[**listSecurityGroups**](SecurityGroupsApi.md#listSecurityGroups) | **GET** /api/security-groups | Retrieves all Security Groups
[**removeSecurityGroupLocations**](SecurityGroupsApi.md#removeSecurityGroupLocations) | **DELETE** /api/security-groups/{id}/locations/{locationId} | Deletes a Security Group Location
[**removeSecurityGroupRules**](SecurityGroupsApi.md#removeSecurityGroupRules) | **DELETE** /api/security-groups/{id}/rules/{sgId} | Deletes a Security Group Rule
[**removeSecurityGroups**](SecurityGroupsApi.md#removeSecurityGroups) | **DELETE** /api/security-groups/{id} | Deletes a Security Group
[**updateSecurityGroupRules**](SecurityGroupsApi.md#updateSecurityGroupRules) | **PUT** /api/security-groups/{id}/rules/{sgId} | Updates a Security Group Rule
[**updateSecurityGroups**](SecurityGroupsApi.md#updateSecurityGroups) | **PUT** /api/security-groups/{id} | Updating a Security Group


<a name="addSecurityGroupLocations"></a>
# **addSecurityGroupLocations**
> Object addSecurityGroupLocations(id, inlineObject215)

Creates a Security Group Location

Creates a security group location. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityGroupsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityGroupsApi apiInstance = new SecurityGroupsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject215 inlineObject215 = new InlineObject215(); // InlineObject215 | 
    try {
      Object result = apiInstance.addSecurityGroupLocations(id, inlineObject215);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityGroupsApi#addSecurityGroupLocations");
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
 **inlineObject215** | [**InlineObject215**](InlineObject215.md)|  | [optional]

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

<a name="addSecurityGroupRules"></a>
# **addSecurityGroupRules**
> Object addSecurityGroupRules(id, inlineObject216)

Creates a Security Group Rule

Creates a security group rule on specified security group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityGroupsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityGroupsApi apiInstance = new SecurityGroupsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject216 inlineObject216 = new InlineObject216(); // InlineObject216 | 
    try {
      Object result = apiInstance.addSecurityGroupRules(id, inlineObject216);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityGroupsApi#addSecurityGroupRules");
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
 **inlineObject216** | [**InlineObject216**](InlineObject216.md)|  | [optional]

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

<a name="addSecurityGroups"></a>
# **addSecurityGroups**
> InlineResponse200133 addSecurityGroups(inlineObject213)

Creates a Security Group

Creates a security group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityGroupsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityGroupsApi apiInstance = new SecurityGroupsApi(defaultClient);
    InlineObject213 inlineObject213 = new InlineObject213(); // InlineObject213 | 
    try {
      InlineResponse200133 result = apiInstance.addSecurityGroups(inlineObject213);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityGroupsApi#addSecurityGroups");
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
 **inlineObject213** | [**InlineObject213**](InlineObject213.md)|  | [optional]

### Return type

[**InlineResponse200133**](InlineResponse200133.md)

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

<a name="getSecurityGroupRules"></a>
# **getSecurityGroupRules**
> InlineResponse200135 getSecurityGroupRules(id, sgId)

Retrieves a Specific Security Group Rule

Retrieves a specific security group rule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityGroupsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityGroupsApi apiInstance = new SecurityGroupsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal sgId = new BigDecimal(78); // BigDecimal | Morpheus ID of the security group rule being referenced
    try {
      InlineResponse200135 result = apiInstance.getSecurityGroupRules(id, sgId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityGroupsApi#getSecurityGroupRules");
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
 **sgId** | **BigDecimal**| Morpheus ID of the security group rule being referenced |

### Return type

[**InlineResponse200135**](InlineResponse200135.md)

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

<a name="getSecurityGroups"></a>
# **getSecurityGroups**
> InlineResponse200134 getSecurityGroups(id)

Retrieves a Specific Security Group

Retrieves a specific security group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityGroupsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityGroupsApi apiInstance = new SecurityGroupsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200134 result = apiInstance.getSecurityGroups(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityGroupsApi#getSecurityGroups");
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

[**InlineResponse200134**](InlineResponse200134.md)

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

<a name="listSecurityGroupRules"></a>
# **listSecurityGroupRules**
> Object listSecurityGroupRules(id, max, offset, sort, direction, phrase, name)

Retrieves all Security Group Rules

Retrieves all security group rules for specified security group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityGroupsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityGroupsApi apiInstance = new SecurityGroupsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    try {
      Object result = apiInstance.listSecurityGroupRules(id, max, offset, sort, direction, phrase, name);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityGroupsApi#listSecurityGroupRules");
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

<a name="listSecurityGroups"></a>
# **listSecurityGroups**
> Object listSecurityGroups(max, offset, sort, direction, phrase, name)

Retrieves all Security Groups

This endpoint retrieves all security groups and their JSON encoded configuration attributes. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityGroupsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityGroupsApi apiInstance = new SecurityGroupsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    try {
      Object result = apiInstance.listSecurityGroups(max, offset, sort, direction, phrase, name);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityGroupsApi#listSecurityGroups");
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

<a name="removeSecurityGroupLocations"></a>
# **removeSecurityGroupLocations**
> Model200Success removeSecurityGroupLocations(id, locationId)

Deletes a Security Group Location

Deletes a security group location. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityGroupsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityGroupsApi apiInstance = new SecurityGroupsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal locationId = new BigDecimal(78); // BigDecimal | The ID of the location
    try {
      Model200Success result = apiInstance.removeSecurityGroupLocations(id, locationId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityGroupsApi#removeSecurityGroupLocations");
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
 **locationId** | **BigDecimal**| The ID of the location |

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

<a name="removeSecurityGroupRules"></a>
# **removeSecurityGroupRules**
> Model200Success removeSecurityGroupRules(id, sgId)

Deletes a Security Group Rule

Deletes a security group rule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityGroupsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityGroupsApi apiInstance = new SecurityGroupsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal sgId = new BigDecimal(78); // BigDecimal | Morpheus ID of the security group rule being referenced
    try {
      Model200Success result = apiInstance.removeSecurityGroupRules(id, sgId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityGroupsApi#removeSecurityGroupRules");
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
 **sgId** | **BigDecimal**| Morpheus ID of the security group rule being referenced |

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

<a name="removeSecurityGroups"></a>
# **removeSecurityGroups**
> Model200Success removeSecurityGroups(id)

Deletes a Security Group

Deletes a specified security group. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityGroupsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityGroupsApi apiInstance = new SecurityGroupsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.removeSecurityGroups(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityGroupsApi#removeSecurityGroups");
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

<a name="updateSecurityGroupRules"></a>
# **updateSecurityGroupRules**
> Object updateSecurityGroupRules(id, sgId, inlineObject217)

Updates a Security Group Rule

Updates a security group rule. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityGroupsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityGroupsApi apiInstance = new SecurityGroupsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    BigDecimal sgId = new BigDecimal(78); // BigDecimal | Morpheus ID of the security group rule being referenced
    InlineObject217 inlineObject217 = new InlineObject217(); // InlineObject217 | 
    try {
      Object result = apiInstance.updateSecurityGroupRules(id, sgId, inlineObject217);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityGroupsApi#updateSecurityGroupRules");
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
 **sgId** | **BigDecimal**| Morpheus ID of the security group rule being referenced |
 **inlineObject217** | [**InlineObject217**](InlineObject217.md)|  | [optional]

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

<a name="updateSecurityGroups"></a>
# **updateSecurityGroups**
> InlineResponse200133 updateSecurityGroups(id, inlineObject214)

Updating a Security Group

Updating a Security Group 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SecurityGroupsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    SecurityGroupsApi apiInstance = new SecurityGroupsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject214 inlineObject214 = new InlineObject214(); // InlineObject214 | 
    try {
      InlineResponse200133 result = apiInstance.updateSecurityGroups(id, inlineObject214);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SecurityGroupsApi#updateSecurityGroups");
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
 **inlineObject214** | [**InlineObject214**](InlineObject214.md)|  | [optional]

### Return type

[**InlineResponse200133**](InlineResponse200133.md)

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

