# MorpheusApi.CypherApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCypherKey**](CypherApi.md#addCypherKey) | **POST** /api/cypher/{cypherPath} | Write a Cypher
[**getCypherKey**](CypherApi.md#getCypherKey) | **GET** /api/cypher/{cypherPath} | Read or Create a Cypher Key
[**listCypherKeys**](CypherApi.md#listCypherKeys) | **GET** /api/cypher | List Cypher Keys
[**removeCypher**](CypherApi.md#removeCypher) | **DELETE** /api/cypher/{cypherPath} | Delete a Cypher



## addCypherKey

> Object addCypherKey(cypherPath, opts)

Write a Cypher

This endpoint will create or update a cypher key.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CypherApi();
let cypherPath = "cypherPath_example"; // String | The key includes a mount prefix separated by a /. For example, the key secret/foo uses the secret mount.  Available Mounts  <table>   <tr>     <th>Mount</th>     <th>Description</th>     <th>Example</th>   </tr>   <tr>     <td>password</td>     <td>Generates a secure password of specified character length in the key pattern (or 15) with symbols, numbers, upper case, and lower case letters (i.e. password/15/mypass generates a 15 character password).</td>     <td>password/15/mypass</td>   </tr>   <tr>     <td>tfvars</td>     <td>This is a module to store a tfvars file for terraform.</td>     <td>tfvars/mytfvar</td>   </tr>   <tr>     <td>secret</td>     <td>This is the standard secret module that stores a key/value in encrypted form. Capable of storing entire JSON object or a String.</td>     <td>secret/foo</td>   </tr>   <tr>     <td>uuid</td>     <td>Returns a new UUID by key name when requested and stores the generated UUID by key name for a given lease timeout period.</td>     <td>uuid/autoMac1</td>   </tr>   <tr>     <td>key</td>     <td>Generates a Base 64 encoded AES Key of specified bit length in the key pattern (i.e. key/128/mykey generates a 128-bit key)</td>     <td>key/128/mykey</td>   </tr> </table> 
let opts = {
  'ttl': 30d, // OneOfstringlong | Time to Live. The lease duration in seconds, or a human readable format eg. '15m', 8h, '7d'.  0 means no expiry.
  'value': {"foo":"bar", "zip":"zap"}, // String | The secret value to be stored. Only required for certain mounts. Some mounts generate their own value and do not require a value to be passed. eg. `uuid`, `key` and `password`.
  'type': "type_example", // String | The type of data being stored, `string` or `object`. The data type depends on the cypher mount being used. Most mounts use `string` as their data type, but `secret` uses `object` by default. You can store a string instead by passing `type=string`. This means the `data` value returned by the API will be a string instead of an object.
  'inlineObject66': new MorpheusApi.InlineObject66() // InlineObject66 | 
};
apiInstance.addCypherKey(cypherPath, opts, (error, data, response) => {
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
 **cypherPath** | **String**| The key includes a mount prefix separated by a /. For example, the key secret/foo uses the secret mount.  Available Mounts  &lt;table&gt;   &lt;tr&gt;     &lt;th&gt;Mount&lt;/th&gt;     &lt;th&gt;Description&lt;/th&gt;     &lt;th&gt;Example&lt;/th&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;password&lt;/td&gt;     &lt;td&gt;Generates a secure password of specified character length in the key pattern (or 15) with symbols, numbers, upper case, and lower case letters (i.e. password/15/mypass generates a 15 character password).&lt;/td&gt;     &lt;td&gt;password/15/mypass&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;tfvars&lt;/td&gt;     &lt;td&gt;This is a module to store a tfvars file for terraform.&lt;/td&gt;     &lt;td&gt;tfvars/mytfvar&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;secret&lt;/td&gt;     &lt;td&gt;This is the standard secret module that stores a key/value in encrypted form. Capable of storing entire JSON object or a String.&lt;/td&gt;     &lt;td&gt;secret/foo&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;uuid&lt;/td&gt;     &lt;td&gt;Returns a new UUID by key name when requested and stores the generated UUID by key name for a given lease timeout period.&lt;/td&gt;     &lt;td&gt;uuid/autoMac1&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;key&lt;/td&gt;     &lt;td&gt;Generates a Base 64 encoded AES Key of specified bit length in the key pattern (i.e. key/128/mykey generates a 128-bit key)&lt;/td&gt;     &lt;td&gt;key/128/mykey&lt;/td&gt;   &lt;/tr&gt; &lt;/table&gt;  | 
 **ttl** | [**OneOfstringlong**](.md)| Time to Live. The lease duration in seconds, or a human readable format eg. &#39;15m&#39;, 8h, &#39;7d&#39;.  0 means no expiry. | [optional] 
 **value** | **String**| The secret value to be stored. Only required for certain mounts. Some mounts generate their own value and do not require a value to be passed. eg. &#x60;uuid&#x60;, &#x60;key&#x60; and &#x60;password&#x60;. | [optional] 
 **type** | **String**| The type of data being stored, &#x60;string&#x60; or &#x60;object&#x60;. The data type depends on the cypher mount being used. Most mounts use &#x60;string&#x60; as their data type, but &#x60;secret&#x60; uses &#x60;object&#x60; by default. You can store a string instead by passing &#x60;type&#x3D;string&#x60;. This means the &#x60;data&#x60; value returned by the API will be a string instead of an object. | [optional] 
 **inlineObject66** | [**InlineObject66**](InlineObject66.md)|  | [optional] 

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## getCypherKey

> Model200Success getCypherKey(cypherPath, opts)

Read or Create a Cypher Key

This endpoint retrieves a specific cypher key. The value of the key is decrypted and returned as data. It may be a String or an object with many {\&quot;key\&quot;:\&quot;value\&quot;} pairs.  The type depends on the cypher mount&#39;s capabilities and what type of data was written to the key.  For example the &#x60;secret/&#x60; mount allows either a string or an object, while the &#x60;password/&#x60; mount will always store and return a string. This endpoint can also create a key. This only applies to mount types &#x60;uuid&#x60;, &#x60;key&#x60;, &#x60;password&#x60;.  Refer to the &#x60;POST&#x60; endpoint for more information. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"
// Configure API key authorization: cypherAuth-XCToken
let cypherAuth-XCToken = defaultClient.authentications['cypherAuth-XCToken'];
cypherAuth-XCToken.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//cypherAuth-XCToken.apiKeyPrefix = 'Token';
// Configure API key authorization: cypherAuth-XMLease
let cypherAuth-XMLease = defaultClient.authentications['cypherAuth-XMLease'];
cypherAuth-XMLease.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//cypherAuth-XMLease.apiKeyPrefix = 'Token';
// Configure API key authorization: cypherAuth-XVToken
let cypherAuth-XVToken = defaultClient.authentications['cypherAuth-XVToken'];
cypherAuth-XVToken.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//cypherAuth-XVToken.apiKeyPrefix = 'Token';

let apiInstance = new MorpheusApi.CypherApi();
let cypherPath = password/15/mypass; // String | The cypher key including the mount prefix.
let opts = {
  'leaseToken': 5bd808dc-82bb-4974-b699-5f1cdd3cc503, // String | An execution lease token.
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc // String | Sort direction, use 'desc' to reverse sort
};
apiInstance.getCypherKey(cypherPath, opts, (error, data, response) => {
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
 **cypherPath** | **String**| The cypher key including the mount prefix. | 
 **leaseToken** | **String**| An execution lease token. | [optional] 
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cypherAuth-XCToken](../README.md#cypherAuth-XCToken), [cypherAuth-XMLease](../README.md#cypherAuth-XMLease), [cypherAuth-XVToken](../README.md#cypherAuth-XVToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listCypherKeys

> Object listCypherKeys(opts)

List Cypher Keys

This endpoint retrieves all cypher keys associated with the account, or user.  This method can be used to list keys as well, by passing the query parameter list&#x3D;true.

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"
// Configure API key authorization: cypherAuth-XCToken
let cypherAuth-XCToken = defaultClient.authentications['cypherAuth-XCToken'];
cypherAuth-XCToken.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//cypherAuth-XCToken.apiKeyPrefix = 'Token';
// Configure API key authorization: cypherAuth-XMLease
let cypherAuth-XMLease = defaultClient.authentications['cypherAuth-XMLease'];
cypherAuth-XMLease.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//cypherAuth-XMLease.apiKeyPrefix = 'Token';
// Configure API key authorization: cypherAuth-XVToken
let cypherAuth-XVToken = defaultClient.authentications['cypherAuth-XVToken'];
cypherAuth-XVToken.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//cypherAuth-XVToken.apiKeyPrefix = 'Token';

let apiInstance = new MorpheusApi.CypherApi();
let opts = {
  'leaseToken': 5bd808dc-82bb-4974-b699-5f1cdd3cc503, // String | An execution lease token.
  'list': false, // Boolean | This endpoint is available via the http method LIST. The GET method can be used to list keys as well, by passing the query parameter list=true.
  'phrase': "phrase_example", // String | Search phrase for partial matches on name or description
  'max': 25, // Number | Maximum number of records to return, -1 can be used to fetch all records
  'offset': 0, // Number | Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // String | Sort order, the name of the property to sort by
  'direction': asc // String | Sort direction, use 'desc' to reverse sort
};
apiInstance.listCypherKeys(opts, (error, data, response) => {
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
 **leaseToken** | **String**| An execution lease token. | [optional] 
 **list** | **Boolean**| This endpoint is available via the http method LIST. The GET method can be used to list keys as well, by passing the query parameter list&#x3D;true. | [optional] [default to false]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional] 
 **max** | **Number**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25]
 **offset** | **Number**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0]
 **sort** | **String**| Sort order, the name of the property to sort by | [optional] [default to &#39;name&#39;]
 **direction** | **String**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] [default to &#39;asc&#39;]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth), [cypherAuth-XCToken](../README.md#cypherAuth-XCToken), [cypherAuth-XMLease](../README.md#cypherAuth-XMLease), [cypherAuth-XVToken](../README.md#cypherAuth-XVToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## removeCypher

> Model200Success removeCypher(cypherPath)

Delete a Cypher

Will delete a cypher from the system and make it no longer usable. 

### Example

```javascript
import MorpheusApi from 'morpheus_api';
let defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new MorpheusApi.CypherApi();
let cypherPath = password/15/mypass; // String | The cypher key including the mount prefix.
apiInstance.removeCypher(cypherPath, (error, data, response) => {
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
 **cypherPath** | **String**| The cypher key including the mount prefix. | 

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

