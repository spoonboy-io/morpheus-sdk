

# OptionTypeListCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name | 
**description** | **String** | Description |  [optional]
**labels** | **List&lt;String&gt;** | Array of label strings, can be used for filtering. |  [optional]
**type** | [**TypeEnum**](#TypeEnum) | Option List Type eg. &#x60;rest&#x60;, &#x60;api&#x60;, &#x60;ldap&#x60; or &#x60;manual&#x60;. |  [optional]
**sourceUrl** | **String** | Source URL, the http(s) URL to request data from. Required when type is rest. |  [optional]
**visibility** | [**VisibilityEnum**](#VisibilityEnum) | Visibility |  [optional]
**sourceMethod** | [**SourceMethodEnum**](#SourceMethodEnum) | Source Method, the HTTP method. |  [optional]
**apiType** | **String** | Api Type, The code of the api option list to use, eg. clouds, environments, groups, instances, instance-wiki, networks, servicePlans, resourcePools, securityGroups, servers, server-wiki. Required when type is api. |  [optional]
**ignoreSSLErrors** | **Boolean** | Ignore SSL Errors. |  [optional]
**realTime** | **Boolean** | Real Time. |  [optional]
**credential** | [**OptionTypeListCreateCredential**](OptionTypeListCreateCredential.md) |  |  [optional]
**serviceUsername** | **String** | Username for authenticating with Basic Auth when type is rest or ldap username. |  [optional]
**servicePassword** | **String** | Password for authenticating with Basic Auth when type is rest or ldap password. |  [optional]
**initialDataset** | **String** | Initial Dataset. Create an initial JSON or CSV dataset to be used as the collection for this option list. It should be a list containing objects with properties &#39;name&#39;, and &#39;value&#39;. Required when type is manual. |  [optional]
**translationScript** | **String** | Translation Script. Create a js script to translate the result data object into an Array containing objects with properties &#39;name&#39; and &#39;value&#39;. The input data is provided as data and the result should be put on the global variable results. |  [optional]
**requestScript** | **String** | Request Script. Create a js script to prepare the request. Return a data object as the body for a post, and return an array containing properties &#39;name&#39; and &#39;value&#39; for a get. The input data is provided as data and the result should be put on the global variable results. |  [optional]
**config** | [**OptionTypeListCreateConfig**](OptionTypeListCreateConfig.md) |  |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
REST | &quot;rest&quot;
API | &quot;api&quot;
LDAP | &quot;ldap&quot;
MANUAL | &quot;manual&quot;



## Enum: VisibilityEnum

Name | Value
---- | -----
PRIVATE | &quot;private&quot;
PUBLIC | &quot;public&quot;



## Enum: SourceMethodEnum

Name | Value
---- | -----
GET | &quot;GET&quot;
POST | &quot;POST&quot;



