# OptionTypeListUpdate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name | [optional] 
**description** | **str, none_type** | Description | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**type** | **str** | Option List Type eg. &#x60;rest&#x60;, &#x60;api&#x60;, &#x60;ldap&#x60; or &#x60;manual&#x60;. | [optional]  if omitted the server will use the default value of "rest"
**source_url** | **str** | Source URL, the http(s) URL to request data from. Required when type is rest. | [optional] 
**visibility** | **str** | Visibility | [optional]  if omitted the server will use the default value of "private"
**source_method** | **str** | Source Method, the HTTP method. | [optional]  if omitted the server will use the default value of "GET"
**api_type** | **str, none_type** | Api Type, The code of the api option list to use, eg. clouds, environments, groups, instances, instance-wiki, networks, servicePlans, resourcePools, securityGroups, servers, server-wiki. Required when type is api. | [optional] 
**ignore_ssl_errors** | **bool** | Ignore SSL Errors. | [optional]  if omitted the server will use the default value of False
**real_time** | **bool** | Real Time. | [optional]  if omitted the server will use the default value of False
**credential** | [**OptionTypeListCreateCredential**](OptionTypeListCreateCredential.md) |  | [optional] 
**service_username** | **str, none_type** | Username for authenticating with Basic Auth when type is rest or ldap username. | [optional] 
**service_password** | **str, none_type** | Password for authenticating with Basic Auth when type is rest or ldap password. | [optional] 
**initial_dataset** | **str, none_type** | Initial Dataset. Create an initial JSON or CSV dataset to be used as the collection for this option list. It should be a list containing objects with properties &#39;name&#39;, and &#39;value&#39;. Required when type is manual. | [optional] 
**translation_script** | **str, none_type** | Translation Script. Create a js script to translate the result data object into an Array containing objects with properties &#39;name&#39; and &#39;value&#39;. The input data is provided as data and the result should be put on the global variable results. | [optional] 
**request_script** | **str, none_type** | Request Script. Create a js script to prepare the request. Return a data object as the body for a post, and return an array containing properties &#39;name&#39; and &#39;value&#39; for a get. The input data is provided as data and the result should be put on the global variable results. | [optional] 
**config** | [**OptionTypeListCreateConfig**](OptionTypeListCreateConfig.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


