# ProvisioningLicensesApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addProvisioningLicense**](ProvisioningLicensesApi.md#addProvisioningLicense) | **POST** /api/provisioning-licenses | Create a License
[**getProvisioningLicense**](ProvisioningLicensesApi.md#getProvisioningLicense) | **GET** /api/provisioning-licenses/{id} | Get a Specific License
[**getProvisioningLicenseReservations**](ProvisioningLicensesApi.md#getProvisioningLicenseReservations) | **GET** /api/provisioning-licenses/{id}/reservations | Get Reservations for Specific License
[**listProvisioningLicenses**](ProvisioningLicensesApi.md#listProvisioningLicenses) | **GET** /api/provisioning-licenses | Get All Licenses
[**removeProvisioningLicense**](ProvisioningLicensesApi.md#removeProvisioningLicense) | **DELETE** /api/provisioning-licenses/{id} | Delete a License
[**updateProvisioningLicense**](ProvisioningLicensesApi.md#updateProvisioningLicense) | **PUT** /api/provisioning-licenses/{id} | Update a License


<a name="addProvisioningLicense"></a>
# **addProvisioningLicense**
> Model200Success addProvisioningLicense(inlineObject202)

Create a License

Use this command to create a new license.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProvisioningLicensesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ProvisioningLicensesApi apiInstance = new ProvisioningLicensesApi(defaultClient);
    InlineObject202 inlineObject202 = new InlineObject202(); // InlineObject202 | 
    try {
      Model200Success result = apiInstance.addProvisioningLicense(inlineObject202);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProvisioningLicensesApi#addProvisioningLicense");
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
 **inlineObject202** | [**InlineObject202**](InlineObject202.md)|  | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getProvisioningLicense"></a>
# **getProvisioningLicense**
> InlineResponse200126 getProvisioningLicense(id)

Get a Specific License

This endpoint retrieves a specific license.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProvisioningLicensesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ProvisioningLicensesApi apiInstance = new ProvisioningLicensesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200126 result = apiInstance.getProvisioningLicense(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProvisioningLicensesApi#getProvisioningLicense");
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

[**InlineResponse200126**](InlineResponse200126.md)

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

<a name="getProvisioningLicenseReservations"></a>
# **getProvisioningLicenseReservations**
> InlineResponse200127 getProvisioningLicenseReservations(id)

Get Reservations for Specific License

This endpoint retrieves all reservations for a specific license. Each time a license is applied to a new server, a reservation is created, reducing the available copies for the license.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProvisioningLicensesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ProvisioningLicensesApi apiInstance = new ProvisioningLicensesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200127 result = apiInstance.getProvisioningLicenseReservations(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProvisioningLicensesApi#getProvisioningLicenseReservations");
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

[**InlineResponse200127**](InlineResponse200127.md)

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

<a name="listProvisioningLicenses"></a>
# **listProvisioningLicenses**
> Object listProvisioningLicenses(max, offset, sort, direction, phrase, name, licenseType, licenseVersion, orgName, fullName)

Get All Licenses

This endpoint retrieves all licenses.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProvisioningLicensesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ProvisioningLicensesApi apiInstance = new ProvisioningLicensesApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String licenseType = "win"; // String | If specified will return an exact match on licenseType code
    String licenseVersion = "2012 R2 Standard"; // String | If specified will return an exact match on licenseVersion
    String orgName = "Acme Motors"; // String | If specified will return an exact match on orgName
    String fullName = "Bugs Bunny"; // String | If specified will return an exact match on fullName
    try {
      Object result = apiInstance.listProvisioningLicenses(max, offset, sort, direction, phrase, name, licenseType, licenseVersion, orgName, fullName);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProvisioningLicensesApi#listProvisioningLicenses");
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
 **licenseType** | **String**| If specified will return an exact match on licenseType code | [optional]
 **licenseVersion** | **String**| If specified will return an exact match on licenseVersion | [optional]
 **orgName** | **String**| If specified will return an exact match on orgName | [optional]
 **fullName** | **String**| If specified will return an exact match on fullName | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="removeProvisioningLicense"></a>
# **removeProvisioningLicense**
> Model200Success removeProvisioningLicense(id)

Delete a License

Will delete a license.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProvisioningLicensesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ProvisioningLicensesApi apiInstance = new ProvisioningLicensesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.removeProvisioningLicense(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProvisioningLicensesApi#removeProvisioningLicense");
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

<a name="updateProvisioningLicense"></a>
# **updateProvisioningLicense**
> Model200Success updateProvisioningLicense(id, inlineObject203)

Update a License

Use this command to update an existing license.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.ProvisioningLicensesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    ProvisioningLicensesApi apiInstance = new ProvisioningLicensesApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject203 inlineObject203 = new InlineObject203(); // InlineObject203 | 
    try {
      Model200Success result = apiInstance.updateProvisioningLicense(id, inlineObject203);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling ProvisioningLicensesApi#updateProvisioningLicense");
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
 **inlineObject203** | [**InlineObject203**](InlineObject203.md)|  | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

