# openapi_client.CypherApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_cypher_key**](CypherApi.md#add_cypher_key) | **POST** /api/cypher/{cypherPath} | Write a Cypher
[**get_cypher_key**](CypherApi.md#get_cypher_key) | **GET** /api/cypher/{cypherPath} | Read or Create a Cypher Key
[**list_cypher_keys**](CypherApi.md#list_cypher_keys) | **GET** /api/cypher | List Cypher Keys
[**remove_cypher**](CypherApi.md#remove_cypher) | **DELETE** /api/cypher/{cypherPath} | Delete a Cypher


# **add_cypher_key**
> AddCypherKey200Response add_cypher_key(cypher_path)

Write a Cypher

This endpoint will create or update a cypher key.

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import cypher_api
from openapi_client.model.add_cypher_key200_response import AddCypherKey200Response
from openapi_client.model.add_cypher_key_request import AddCypherKeyRequest
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cypher_api.CypherApi(api_client)
    cypher_path = "cypherPath_example" # str | The key includes a mount prefix separated by a /. For example, the key secret/foo uses the secret mount.  Available Mounts  <table>   <tr>     <th>Mount</th>     <th>Description</th>     <th>Example</th>   </tr>   <tr>     <td>password</td>     <td>Generates a secure password of specified character length in the key pattern (or 15) with symbols, numbers, upper case, and lower case letters (i.e. password/15/mypass generates a 15 character password).</td>     <td>password/15/mypass</td>   </tr>   <tr>     <td>tfvars</td>     <td>This is a module to store a tfvars file for terraform.</td>     <td>tfvars/mytfvar</td>   </tr>   <tr>     <td>secret</td>     <td>This is the standard secret module that stores a key/value in encrypted form. Capable of storing entire JSON object or a String.</td>     <td>secret/foo</td>   </tr>   <tr>     <td>uuid</td>     <td>Returns a new UUID by key name when requested and stores the generated UUID by key name for a given lease timeout period.</td>     <td>uuid/autoMac1</td>   </tr>   <tr>     <td>key</td>     <td>Generates a Base 64 encoded AES Key of specified bit length in the key pattern (i.e. key/128/mykey generates a 128-bit key)</td>     <td>key/128/mykey</td>   </tr> </table> 
    ttl = AddCypherKeyTtlParameter(None) # AddCypherKeyTtlParameter | Time to Live. The lease duration in seconds, or a human readable format eg. '15m', 8h, '7d'.  0 means no expiry. (optional)
    value = "{"foo":"bar", "zip":"zap"}" # str | The secret value to be stored. Only required for certain mounts. Some mounts generate their own value and do not require a value to be passed. eg. `uuid`, `key` and `password`. (optional)
    type = "type_example" # str | The type of data being stored, `string` or `object`. The data type depends on the cypher mount being used. Most mounts use `string` as their data type, but `secret` uses `object` by default. You can store a string instead by passing `type=string`. This means the `data` value returned by the API will be a string instead of an object. (optional)
    add_cypher_key_request = AddCypherKeyRequest(
        ttl=AddCypherKeyRequestTtl(None),
        value="value_example",
    ) # AddCypherKeyRequest |  (optional)

    # example passing only required values which don't have defaults set
    try:
        # Write a Cypher
        api_response = api_instance.add_cypher_key(cypher_path)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CypherApi->add_cypher_key: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Write a Cypher
        api_response = api_instance.add_cypher_key(cypher_path, ttl=ttl, value=value, type=type, add_cypher_key_request=add_cypher_key_request)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CypherApi->add_cypher_key: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cypher_path** | **str**| The key includes a mount prefix separated by a /. For example, the key secret/foo uses the secret mount.  Available Mounts  &lt;table&gt;   &lt;tr&gt;     &lt;th&gt;Mount&lt;/th&gt;     &lt;th&gt;Description&lt;/th&gt;     &lt;th&gt;Example&lt;/th&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;password&lt;/td&gt;     &lt;td&gt;Generates a secure password of specified character length in the key pattern (or 15) with symbols, numbers, upper case, and lower case letters (i.e. password/15/mypass generates a 15 character password).&lt;/td&gt;     &lt;td&gt;password/15/mypass&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;tfvars&lt;/td&gt;     &lt;td&gt;This is a module to store a tfvars file for terraform.&lt;/td&gt;     &lt;td&gt;tfvars/mytfvar&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;secret&lt;/td&gt;     &lt;td&gt;This is the standard secret module that stores a key/value in encrypted form. Capable of storing entire JSON object or a String.&lt;/td&gt;     &lt;td&gt;secret/foo&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;uuid&lt;/td&gt;     &lt;td&gt;Returns a new UUID by key name when requested and stores the generated UUID by key name for a given lease timeout period.&lt;/td&gt;     &lt;td&gt;uuid/autoMac1&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;key&lt;/td&gt;     &lt;td&gt;Generates a Base 64 encoded AES Key of specified bit length in the key pattern (i.e. key/128/mykey generates a 128-bit key)&lt;/td&gt;     &lt;td&gt;key/128/mykey&lt;/td&gt;   &lt;/tr&gt; &lt;/table&gt;  |
 **ttl** | **AddCypherKeyTtlParameter**| Time to Live. The lease duration in seconds, or a human readable format eg. &#39;15m&#39;, 8h, &#39;7d&#39;.  0 means no expiry. | [optional]
 **value** | **str**| The secret value to be stored. Only required for certain mounts. Some mounts generate their own value and do not require a value to be passed. eg. &#x60;uuid&#x60;, &#x60;key&#x60; and &#x60;password&#x60;. | [optional]
 **type** | **str**| The type of data being stored, &#x60;string&#x60; or &#x60;object&#x60;. The data type depends on the cypher mount being used. Most mounts use &#x60;string&#x60; as their data type, but &#x60;secret&#x60; uses &#x60;object&#x60; by default. You can store a string instead by passing &#x60;type&#x3D;string&#x60;. This means the &#x60;data&#x60; value returned by the API will be a string instead of an object. | [optional]
 **add_cypher_key_request** | [**AddCypherKeyRequest**](AddCypherKeyRequest.md)|  | [optional]

### Return type

[**AddCypherKey200Response**](AddCypherKey200Response.md)

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cypher_key**
> GetCypherKey200Response get_cypher_key(cypher_path)

Read or Create a Cypher Key

This endpoint retrieves a specific cypher key. The value of the key is decrypted and returned as data. It may be a String or an object with many {\"key\":\"value\"} pairs.  The type depends on the cypher mount's capabilities and what type of data was written to the key.  For example the `secret/` mount allows either a string or an object, while the `password/` mount will always store and return a string. This endpoint can also create a key. This only applies to mount types `uuid`, `key`, `password`.  Refer to the `POST` endpoint for more information. 

### Example

* Bearer Authentication (bearerAuth):
* Api Key Authentication (cypherAuth-XCToken):
* Api Key Authentication (cypherAuth-XMLease):
* Api Key Authentication (cypherAuth-XVToken):

```python
import time
import openapi_client
from openapi_client.api import cypher_api
from openapi_client.model.get_cypher_key200_response import GetCypherKey200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Configure API key authorization: cypherAuth-XCToken
configuration.api_key['cypherAuth-XCToken'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['cypherAuth-XCToken'] = 'Bearer'

# Configure API key authorization: cypherAuth-XMLease
configuration.api_key['cypherAuth-XMLease'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['cypherAuth-XMLease'] = 'Bearer'

# Configure API key authorization: cypherAuth-XVToken
configuration.api_key['cypherAuth-XVToken'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['cypherAuth-XVToken'] = 'Bearer'

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cypher_api.CypherApi(api_client)
    cypher_path = "password/15/mypass" # str | The cypher key including the mount prefix.
    lease_token = "5bd808dc-82bb-4974-b699-5f1cdd3cc503" # str | An execution lease token. (optional)
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"

    # example passing only required values which don't have defaults set
    try:
        # Read or Create a Cypher Key
        api_response = api_instance.get_cypher_key(cypher_path)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CypherApi->get_cypher_key: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Read or Create a Cypher Key
        api_response = api_instance.get_cypher_key(cypher_path, lease_token=lease_token, sort=sort, direction=direction)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CypherApi->get_cypher_key: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cypher_path** | **str**| The cypher key including the mount prefix. |
 **lease_token** | **str**| An execution lease token. | [optional]
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"

### Return type

[**GetCypherKey200Response**](GetCypherKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cypherAuth-XCToken](../README.md#cypherAuth-XCToken), [cypherAuth-XMLease](../README.md#cypherAuth-XMLease), [cypherAuth-XVToken](../README.md#cypherAuth-XVToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_cypher_keys**
> ListCypherKeys200Response list_cypher_keys()

List Cypher Keys

This endpoint retrieves all cypher keys associated with the account, or user.  This method can be used to list keys as well, by passing the query parameter list=true.

### Example

* Bearer Authentication (bearerAuth):
* Api Key Authentication (cypherAuth-XCToken):
* Api Key Authentication (cypherAuth-XMLease):
* Api Key Authentication (cypherAuth-XVToken):

```python
import time
import openapi_client
from openapi_client.api import cypher_api
from openapi_client.model.list_cypher_keys200_response import ListCypherKeys200Response
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Configure API key authorization: cypherAuth-XCToken
configuration.api_key['cypherAuth-XCToken'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['cypherAuth-XCToken'] = 'Bearer'

# Configure API key authorization: cypherAuth-XMLease
configuration.api_key['cypherAuth-XMLease'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['cypherAuth-XMLease'] = 'Bearer'

# Configure API key authorization: cypherAuth-XVToken
configuration.api_key['cypherAuth-XVToken'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['cypherAuth-XVToken'] = 'Bearer'

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cypher_api.CypherApi(api_client)
    lease_token = "5bd808dc-82bb-4974-b699-5f1cdd3cc503" # str | An execution lease token. (optional)
    list = False # bool | This endpoint is available via the http method LIST. The GET method can be used to list keys as well, by passing the query parameter list=true. (optional) if omitted the server will use the default value of False
    phrase = "phrase_example" # str | Search phrase for partial matches on name or description (optional)
    max = 25 # int | Maximum number of records to return, -1 can be used to fetch all records (optional) if omitted the server will use the default value of 25
    offset = 0 # int | Offset records, the number of records to skip, for paginating requests (optional) if omitted the server will use the default value of 0
    sort = "name" # str | Sort order, the name of the property to sort by (optional) if omitted the server will use the default value of "name"
    direction = "asc" # str | Sort direction, use 'desc' to reverse sort (optional) if omitted the server will use the default value of "asc"

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # List Cypher Keys
        api_response = api_instance.list_cypher_keys(lease_token=lease_token, list=list, phrase=phrase, max=max, offset=offset, sort=sort, direction=direction)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CypherApi->list_cypher_keys: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lease_token** | **str**| An execution lease token. | [optional]
 **list** | **bool**| This endpoint is available via the http method LIST. The GET method can be used to list keys as well, by passing the query parameter list&#x3D;true. | [optional] if omitted the server will use the default value of False
 **phrase** | **str**| Search phrase for partial matches on name or description | [optional]
 **max** | **int**| Maximum number of records to return, -1 can be used to fetch all records | [optional] if omitted the server will use the default value of 25
 **offset** | **int**| Offset records, the number of records to skip, for paginating requests | [optional] if omitted the server will use the default value of 0
 **sort** | **str**| Sort order, the name of the property to sort by | [optional] if omitted the server will use the default value of "name"
 **direction** | **str**| Sort direction, use &#39;desc&#39; to reverse sort | [optional] if omitted the server will use the default value of "asc"

### Return type

[**ListCypherKeys200Response**](ListCypherKeys200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cypherAuth-XCToken](../README.md#cypherAuth-XCToken), [cypherAuth-XMLease](../README.md#cypherAuth-XMLease), [cypherAuth-XVToken](../README.md#cypherAuth-XVToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_cypher**
> Model200Success remove_cypher(cypher_path)

Delete a Cypher

Will delete a cypher from the system and make it no longer usable. 

### Example

* Bearer Authentication (bearerAuth):

```python
import time
import openapi_client
from openapi_client.api import cypher_api
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.default_error import DefaultError
from pprint import pprint
# Defining the host is optional and defaults to https://CHANGEME
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://CHANGEME"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cypher_api.CypherApi(api_client)
    cypher_path = "password/15/mypass" # str | The cypher key including the mount prefix.

    # example passing only required values which don't have defaults set
    try:
        # Delete a Cypher
        api_response = api_instance.remove_cypher(cypher_path)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling CypherApi->remove_cypher: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cypher_path** | **str**| The cypher key including the mount prefix. |

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

