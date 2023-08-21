# # OptionTypeListCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name |
**description** | **string** | Description | [optional]
**labels** | **string[]** | Array of label strings, can be used for filtering. | [optional]
**type** | **string** | Option List Type eg. &#x60;rest&#x60;, &#x60;api&#x60;, &#x60;ldap&#x60; or &#x60;manual&#x60;. | [optional] [default to 'rest']
**source_url** | **string** | Source URL, the http(s) URL to request data from. Required when type is rest. | [optional]
**visibility** | **string** | Visibility | [optional] [default to 'private']
**source_method** | **string** | Source Method, the HTTP method. | [optional] [default to 'GET']
**api_type** | **string** | Api Type, The code of the api option list to use, eg. clouds, environments, groups, instances, instance-wiki, networks, servicePlans, resourcePools, securityGroups, servers, server-wiki. Required when type is api. | [optional]
**ignore_ssl_errors** | **bool** | Ignore SSL Errors. | [optional] [default to false]
**real_time** | **bool** | Real Time. | [optional] [default to false]
**credential** | [**\OpenAPI\Client\Model\OptionTypeListCreateCredential**](OptionTypeListCreateCredential.md) |  | [optional]
**service_username** | **string** | Username for authenticating with Basic Auth when type is rest or ldap username. | [optional]
**service_password** | **string** | Password for authenticating with Basic Auth when type is rest or ldap password. | [optional]
**initial_dataset** | **string** | Initial Dataset. Create an initial JSON or CSV dataset to be used as the collection for this option list. It should be a list containing objects with properties &#39;name&#39;, and &#39;value&#39;. Required when type is manual. | [optional]
**translation_script** | **string** | Translation Script. Create a js script to translate the result data object into an Array containing objects with properties &#39;name&#39; and &#39;value&#39;. The input data is provided as data and the result should be put on the global variable results. | [optional]
**request_script** | **string** | Request Script. Create a js script to prepare the request. Return a data object as the body for a post, and return an array containing properties &#39;name&#39; and &#39;value&#39; for a get. The input data is provided as data and the result should be put on the global variable results. | [optional]
**config** | [**\OpenAPI\Client\Model\OptionTypeListCreateConfig**](OptionTypeListCreateConfig.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
