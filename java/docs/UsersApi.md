# UsersApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addUser**](UsersApi.md#addUser) | **POST** /api/users | Create a New User
[**deleteUser**](UsersApi.md#deleteUser) | **DELETE** /api/users/{id} | Delete a User
[**deleteUserSettingsAccessToken**](UsersApi.md#deleteUserSettingsAccessToken) | **PUT** /api/user-settings/clear-access-token | Revoke API Access Token
[**deleteUserSettingsAvatar**](UsersApi.md#deleteUserSettingsAvatar) | **DELETE** /api/user-settings/avatar | Delete Avatar
[**deleteUserSettingsDesktopBackground**](UsersApi.md#deleteUserSettingsDesktopBackground) | **DELETE** /api/user-settings/desktop-background | Delete Desktop Background
[**getUser**](UsersApi.md#getUser) | **GET** /api/users/{id} | Get a Specific User
[**getUserPermissions**](UsersApi.md#getUserPermissions) | **GET** /api/users/{id}/permissions | Get a Specific User Permissions
[**getUserSettingsApiClients**](UsersApi.md#getUserSettingsApiClients) | **GET** /api/user-settings/api-clients | Get Available API Clients
[**listUserSettings**](UsersApi.md#listUserSettings) | **GET** /api/user-settings | User Settings
[**listUsers**](UsersApi.md#listUsers) | **GET** /api/users | List All Users
[**listUsersAvailableRoles**](UsersApi.md#listUsersAvailableRoles) | **GET** /api/users/available-roles | List available roles for a user
[**updateUserSettings**](UsersApi.md#updateUserSettings) | **PUT** /api/user-settings | Update User Settings
[**updateUserSettingsAccessToken**](UsersApi.md#updateUserSettingsAccessToken) | **PUT** /api/user-settings/regenerate-access-token | Regenerate API Access Token
[**updateUserSettingsAvatar**](UsersApi.md#updateUserSettingsAvatar) | **POST** /api/user-settings/avatar | Update Avatar
[**updateUserSettingsDesktopBackground**](UsersApi.md#updateUserSettingsDesktopBackground) | **POST** /api/user-settings/desktop-background | Update Desktop Background
[**updateUsers**](UsersApi.md#updateUsers) | **PUT** /api/users/{id} | Update user
[**whoami**](UsersApi.md#whoami) | **GET** /api/whoami | Retrieves information about current user roles and permissions


<a name="addUser"></a>
# **addUser**
> Object addUser(accountId, inlineObject255)

Create a New User

Create a new user.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long accountId = 56L; // Long | Tenant ID, create user in a sub tenant account instead of your own.
    InlineObject255 inlineObject255 = new InlineObject255(); // InlineObject255 | 
    try {
      Object result = apiInstance.addUser(accountId, inlineObject255);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#addUser");
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
 **accountId** | **Long**| Tenant ID, create user in a sub tenant account instead of your own. | [optional]
 **inlineObject255** | [**InlineObject255**](InlineObject255.md)|  | [optional]

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

<a name="deleteUser"></a>
# **deleteUser**
> Model200Success deleteUser(id)

Delete a User

Delete an existing user.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteUser(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#deleteUser");
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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteUserSettingsAccessToken"></a>
# **deleteUserSettingsAccessToken**
> Model200Success deleteUserSettingsAccessToken(userId, clientId)

Revoke API Access Token

This endpoint revokes your API access token for the specified client.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long userId = 56L; // Long | ID of User (Only available to the master tenant.)  Default is the current user
    String clientId = "morph-cli"; // String | Client ID.  See `Get Available API Clients` for a list of valid `clientId` values.
    try {
      Model200Success result = apiInstance.deleteUserSettingsAccessToken(userId, clientId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#deleteUserSettingsAccessToken");
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
 **userId** | **Long**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **clientId** | **String**| Client ID.  See &#x60;Get Available API Clients&#x60; for a list of valid &#x60;clientId&#x60; values. | [optional]

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
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteUserSettingsAvatar"></a>
# **deleteUserSettingsAvatar**
> Model200Success deleteUserSettingsAvatar(userId)

Delete Avatar

Delete your avatar image.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long userId = 56L; // Long | ID of User (Only available to the master tenant.)  Default is the current user
    try {
      Model200Success result = apiInstance.deleteUserSettingsAvatar(userId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#deleteUserSettingsAvatar");
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
 **userId** | **Long**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]

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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteUserSettingsDesktopBackground"></a>
# **deleteUserSettingsDesktopBackground**
> Model200Success deleteUserSettingsDesktopBackground(userId)

Delete Desktop Background

Delete your desktop background image.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long userId = 56L; // Long | ID of User (Only available to the master tenant.)  Default is the current user
    try {
      Model200Success result = apiInstance.deleteUserSettingsDesktopBackground(userId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#deleteUserSettingsDesktopBackground");
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
 **userId** | **Long**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]

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
**200** | Success Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getUser"></a>
# **getUser**
> InlineResponse200158 getUser(id, includeAccess)

Get a Specific User

This endpoint will retrieve a specific user by id if the user has permission to access the user.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    Boolean includeAccess = true; // Boolean | Include `access` information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s).
    try {
      InlineResponse200158 result = apiInstance.getUser(id, includeAccess);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#getUser");
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
 **includeAccess** | **Boolean**| Include &#x60;access&#x60; information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s). | [optional]

### Return type

[**InlineResponse200158**](InlineResponse200158.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | User Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getUserPermissions"></a>
# **getUserPermissions**
> InlineResponse200159 getUserPermissions(id)

Get a Specific User Permissions

This will list all the permissions for a specific user.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse200159 result = apiInstance.getUserPermissions(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#getUserPermissions");
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

[**InlineResponse200159**](InlineResponse200159.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | User Object |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getUserSettingsApiClients"></a>
# **getUserSettingsApiClients**
> InlineResponse200157 getUserSettingsApiClients(userId)

Get Available API Clients

This endpoint retrieves a list of available API clients.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long userId = 56L; // Long | ID of User (Only available to the master tenant.)  Default is the current user
    try {
      InlineResponse200157 result = apiInstance.getUserSettingsApiClients(userId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#getUserSettingsApiClients");
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
 **userId** | **Long**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]

### Return type

[**InlineResponse200157**](InlineResponse200157.md)

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

<a name="listUserSettings"></a>
# **listUserSettings**
> UserSettings listUserSettings(userId)

User Settings

Provides API for managing your own user settings and api access tokens.  This endpoint retrieves your user settings and API access token information. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long userId = 56L; // Long | ID of User (Only available to the master tenant.)  Default is the current user
    try {
      UserSettings result = apiInstance.listUserSettings(userId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#listUserSettings");
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
 **userId** | **Long**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]

### Return type

[**UserSettings**](UserSettings.md)

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

<a name="listUsers"></a>
# **listUsers**
> Object listUsers(max, offset, sort, direction, phrase, username, email, roleId, lastUpdated, accountId, global)

List All Users

This endpoint retrieves all users in the current user&#39;s tenant account. Master tenant users with permission to manage subtenants can use &#x60;global&#x3D;true&#x60; to find users across all tenants. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String sort = "\"name\""; // String | Sort order, the name of the property to sort by
    String direction = "asc"; // String | Sort direction, use 'desc' to reverse sort
    String phrase = "phrase_example"; // String | Search phrase for partial matches on username or email
    String username = "administrator"; // String | Username
    String email = "mytest@email.com"; // String | E-Mail Address
    Long roleId = 13L; // Long | Filter by Role ID (supports multiple values)
    OffsetDateTime lastUpdated = OffsetDateTime.now(); // OffsetDateTime | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
    Long accountId = 3L; // Long | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
    Boolean global = false; // Boolean | Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users.
    try {
      Object result = apiInstance.listUsers(max, offset, sort, direction, phrase, username, email, roleId, lastUpdated, accountId, global);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#listUsers");
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
 **phrase** | **String**| Search phrase for partial matches on username or email | [optional]
 **username** | **String**| Username | [optional]
 **email** | **String**| E-Mail Address | [optional]
 **roleId** | **Long**| Filter by Role ID (supports multiple values) | [optional]
 **lastUpdated** | **OffsetDateTime**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional]
 **accountId** | **Long**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional]
 **global** | **Boolean**| Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. | [optional] [default to false]

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
**200** | Array of Users |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listUsersAvailableRoles"></a>
# **listUsersAvailableRoles**
> UsersAvailableRoles listUsersAvailableRoles(accountId)

List available roles for a user

Get a list of roles that can be assigned to a user.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long accountId = 56L; // Long | Tenant ID, find roles available to users of a sub tenant account.
    try {
      UsersAvailableRoles result = apiInstance.listUsersAvailableRoles(accountId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#listUsersAvailableRoles");
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
 **accountId** | **Long**| Tenant ID, find roles available to users of a sub tenant account. | [optional]

### Return type

[**UsersAvailableRoles**](UsersAvailableRoles.md)

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

<a name="updateUserSettings"></a>
# **updateUserSettings**
> Model200Success updateUserSettings(userId, inlineObject252)

Update User Settings

This endpoint allows updating user settings. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long userId = 56L; // Long | ID of User (Only available to the master tenant.)  Default is the current user
    InlineObject252 inlineObject252 = new InlineObject252(); // InlineObject252 | 
    try {
      Model200Success result = apiInstance.updateUserSettings(userId, inlineObject252);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#updateUserSettings");
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
 **userId** | **Long**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **inlineObject252** | [**InlineObject252**](InlineObject252.md)|  | [optional]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateUserSettingsAccessToken"></a>
# **updateUserSettingsAccessToken**
> UserSettingsRegenerateAccessToken updateUserSettingsAccessToken(userId, clientId)

Regenerate API Access Token

This endpoint regenerates your API access token for the specified client. If a current token exists, it is revoked and a new token is returned.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long userId = 56L; // Long | ID of User (Only available to the master tenant.)  Default is the current user
    String clientId = "morph-cli"; // String | Client ID.  See `Get Available API Clients` for a list of valid `clientId` values.
    try {
      UserSettingsRegenerateAccessToken result = apiInstance.updateUserSettingsAccessToken(userId, clientId);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#updateUserSettingsAccessToken");
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
 **userId** | **Long**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **clientId** | **String**| Client ID.  See &#x60;Get Available API Clients&#x60; for a list of valid &#x60;clientId&#x60; values. | [optional]

### Return type

[**UserSettingsRegenerateAccessToken**](UserSettingsRegenerateAccessToken.md)

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

<a name="updateUserSettingsAvatar"></a>
# **updateUserSettingsAvatar**
> Model200Success updateUserSettingsAvatar(userId, userAvatar)

Update Avatar

This endpoint updates your avatar image. Expects multipart form data as the request format, not JSON.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long userId = 56L; // Long | ID of User (Only available to the master tenant.)  Default is the current user
    String userAvatar = "userAvatar_example"; // String | 
    try {
      Model200Success result = apiInstance.updateUserSettingsAvatar(userId, userAvatar);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#updateUserSettingsAvatar");
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
 **userId** | **Long**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **userAvatar** | **String**|  | [optional]

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

<a name="updateUserSettingsDesktopBackground"></a>
# **updateUserSettingsDesktopBackground**
> Model200Success updateUserSettingsDesktopBackground(userId, userDesktopBackground)

Update Desktop Background

This endpoint updates your desktop background image that is used in the Virtual Desktop persona. Expects multipart form data as the request format, not JSON.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long userId = 56L; // Long | ID of User (Only available to the master tenant.)  Default is the current user
    String userDesktopBackground = "userDesktopBackground_example"; // String | 
    try {
      Model200Success result = apiInstance.updateUserSettingsDesktopBackground(userId, userDesktopBackground);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#updateUserSettingsDesktopBackground");
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
 **userId** | **Long**| ID of User (Only available to the master tenant.)  Default is the current user | [optional]
 **userDesktopBackground** | **String**|  | [optional]

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

<a name="updateUsers"></a>
# **updateUsers**
> Object updateUsers(id, inlineObject256)

Update user

Update an existing user.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject256 inlineObject256 = new InlineObject256(); // InlineObject256 | 
    try {
      Object result = apiInstance.updateUsers(id, inlineObject256);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#updateUsers");
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
 **inlineObject256** | [**InlineObject256**](InlineObject256.md)|  | [optional]

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

<a name="whoami"></a>
# **whoami**
> InlineResponse200167 whoami()

Retrieves information about current user roles and permissions

Provides API to retrieve information about yourself, including your roles and permissions.  The appliance build version is also returned. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.UsersApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    UsersApi apiInstance = new UsersApi(defaultClient);
    try {
      InlineResponse200167 result = apiInstance.whoami();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling UsersApi#whoami");
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

[**InlineResponse200167**](InlineResponse200167.md)

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

