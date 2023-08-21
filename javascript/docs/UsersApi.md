# MorpheusApi.UsersApi

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



## addUser

> Object addUser(opts)

Create a New User

Create a new user.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let opts = {
  'accountId': 789, // Number | Tenant ID, create user in a sub tenant account instead of your own.
  'inlineObject255': new MorpheusApi.InlineObject255() // InlineObject255 | 
};
apiInstance.addUser(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Number**| Tenant ID, create user in a sub tenant account instead of your own. | [optional] 
 **inlineObject255** | [**InlineObject255**](InlineObject255.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteUser

> Model200Success deleteUser(id)

Delete a User

Delete an existing user.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.deleteUser(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteUserSettingsAccessToken

> Model200Success deleteUserSettingsAccessToken(opts)

Revoke API Access Token

This endpoint revokes your API access token for the specified client.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let opts = {
  'userId': 789, // Number | ID of User (Only available to the master tenant.)  Default is the current user
  'clientId': morph-cli // String | Client ID.  See `Get Available API Clients` for a list of valid `clientId` values.
};
apiInstance.deleteUserSettingsAccessToken(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **Number**| ID of User (Only available to the master tenant.)  Default is the current user | [optional] 
 **clientId** | **String**| Client ID.  See &#x60;Get Available API Clients&#x60; for a list of valid &#x60;clientId&#x60; values. | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteUserSettingsAvatar

> Model200Success deleteUserSettingsAvatar(opts)

Delete Avatar

Delete your avatar image.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let opts = {
  'userId': 789 // Number | ID of User (Only available to the master tenant.)  Default is the current user
};
apiInstance.deleteUserSettingsAvatar(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **Number**| ID of User (Only available to the master tenant.)  Default is the current user | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## deleteUserSettingsDesktopBackground

> Model200Success deleteUserSettingsDesktopBackground(opts)

Delete Desktop Background

Delete your desktop background image.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let opts = {
  'userId': 789 // Number | ID of User (Only available to the master tenant.)  Default is the current user
};
apiInstance.deleteUserSettingsDesktopBackground(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **Number**| ID of User (Only available to the master tenant.)  Default is the current user | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getUser

> InlineResponse200158 getUser(id, opts)

Get a Specific User

This endpoint will retrieve a specific user by id if the user has permission to access the user.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'includeAccess': true // Boolean | Include `access` information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s).
};
apiInstance.getUser(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **includeAccess** | **Boolean**| Include &#x60;access&#x60; information in the response. This is the permissions, clouds, instanceTypes, etc. that the user is authorized for based on their assigned Role(s). | [optional] 

### Return type

[**InlineResponse200158**](InlineResponse200158.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getUserPermissions

> InlineResponse200159 getUserPermissions(id)

Get a Specific User Permissions

This will list all the permissions for a specific user.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
apiInstance.getUserPermissions(id, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 

### Return type

[**InlineResponse200159**](InlineResponse200159.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getUserSettingsApiClients

> InlineResponse200157 getUserSettingsApiClients(opts)

Get Available API Clients

This endpoint retrieves a list of available API clients.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let opts = {
  'userId': 789 // Number | ID of User (Only available to the master tenant.)  Default is the current user
};
apiInstance.getUserSettingsApiClients(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **Number**| ID of User (Only available to the master tenant.)  Default is the current user | [optional] 

### Return type

[**InlineResponse200157**](InlineResponse200157.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listUserSettings

> UserSettings listUserSettings(opts)

User Settings

Provides API for managing your own user settings and api access tokens.  This endpoint retrieves your user settings and API access token information. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let opts = {
  'userId': 789 // Number | ID of User (Only available to the master tenant.)  Default is the current user
};
apiInstance.listUserSettings(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **Number**| ID of User (Only available to the master tenant.)  Default is the current user | [optional] 

### Return type

[**UserSettings**](UserSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listUsers

> Object listUsers(opts)

List All Users

This endpoint retrieves all users in the current user&#39;s tenant account. Master tenant users with permission to manage subtenants can use &#x60;global&#x3D;true&#x60; to find users across all tenants. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let opts = {
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc, // String | Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // String | Search phrase for partial matches on username or email
  'username': administrator, // String | Username
  'email': mytest@email.com, // String | E-Mail Address
  'roleId': 13, // Number | Filter by Role ID (supports multiple values)
  'lastUpdated': 2019-03-06T17:52:29+0000, // Date | Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601)
  'accountId': 3, // Number | Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users.
  'global': false // Boolean | Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users.
};
apiInstance.listUsers(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]
 **phrase** | **String**| Search phrase for partial matches on username or email | [optional] 
 **username** | **String**| Username | [optional] 
 **email** | **String**| E-Mail Address | [optional] 
 **roleId** | **Number**| Filter by Role ID (supports multiple values) | [optional] 
 **lastUpdated** | **Date**| Date filter, restricts query to only load resources updated more recently than the date specified (ISO 8601) | [optional] 
 **accountId** | **Number**| Filter by Tenant ID. This is only available to master tenant users with permission to manage tenants and users. | [optional] 
 **global** | **Boolean**| Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. | [optional] [default to false]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listUsersAvailableRoles

> UsersAvailableRoles listUsersAvailableRoles(opts)

List available roles for a user

Get a list of roles that can be assigned to a user.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let opts = {
  'accountId': 789 // Number | Tenant ID, find roles available to users of a sub tenant account.
};
apiInstance.listUsersAvailableRoles(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Number**| Tenant ID, find roles available to users of a sub tenant account. | [optional] 

### Return type

[**UsersAvailableRoles**](UsersAvailableRoles.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateUserSettings

> Model200Success updateUserSettings(opts)

Update User Settings

This endpoint allows updating user settings. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let opts = {
  'userId': 789, // Number | ID of User (Only available to the master tenant.)  Default is the current user
  'inlineObject252': new MorpheusApi.InlineObject252() // InlineObject252 | 
};
apiInstance.updateUserSettings(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **Number**| ID of User (Only available to the master tenant.)  Default is the current user | [optional] 
 **inlineObject252** | [**InlineObject252**](InlineObject252.md)|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json


## updateUserSettingsAccessToken

> UserSettingsRegenerateAccessToken updateUserSettingsAccessToken(opts)

Regenerate API Access Token

This endpoint regenerates your API access token for the specified client. If a current token exists, it is revoked and a new token is returned.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let opts = {
  'userId': 789, // Number | ID of User (Only available to the master tenant.)  Default is the current user
  'clientId': morph-cli // String | Client ID.  See `Get Available API Clients` for a list of valid `clientId` values.
};
apiInstance.updateUserSettingsAccessToken(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **Number**| ID of User (Only available to the master tenant.)  Default is the current user | [optional] 
 **clientId** | **String**| Client ID.  See &#x60;Get Available API Clients&#x60; for a list of valid &#x60;clientId&#x60; values. | [optional] 

### Return type

[**UserSettingsRegenerateAccessToken**](UserSettingsRegenerateAccessToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## updateUserSettingsAvatar

> Model200Success updateUserSettingsAvatar(opts)

Update Avatar

This endpoint updates your avatar image. Expects multipart form data as the request format, not JSON.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let opts = {
  'userId': 789, // Number | ID of User (Only available to the master tenant.)  Default is the current user
  'userAvatar': "userAvatar_example" // String | 
};
apiInstance.updateUserSettingsAvatar(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **Number**| ID of User (Only available to the master tenant.)  Default is the current user | [optional] 
 **userAvatar** | **String**|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## updateUserSettingsDesktopBackground

> Model200Success updateUserSettingsDesktopBackground(opts)

Update Desktop Background

This endpoint updates your desktop background image that is used in the Virtual Desktop persona. Expects multipart form data as the request format, not JSON.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let opts = {
  'userId': 789, // Number | ID of User (Only available to the master tenant.)  Default is the current user
  'userDesktopBackground': "userDesktopBackground_example" // String | 
};
apiInstance.updateUserSettingsDesktopBackground(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **Number**| ID of User (Only available to the master tenant.)  Default is the current user | [optional] 
 **userDesktopBackground** | **String**|  | [optional] 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## updateUsers

> Object updateUsers(id, opts)

Update user

Update an existing user.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
let id = 1; // Number | Morpheus ID of the Object being referenced
let opts = {
  'inlineObject256': new MorpheusApi.InlineObject256() // InlineObject256 | 
};
apiInstance.updateUsers(id, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**| Morpheus ID of the Object being referenced | 
 **inlineObject256** | [**InlineObject256**](InlineObject256.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## whoami

> InlineResponse200167 whoami()

Retrieves information about current user roles and permissions

Provides API to retrieve information about yourself, including your roles and permissions.  The appliance build version is also returned. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.UsersApi();
apiInstance.whoami((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
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

