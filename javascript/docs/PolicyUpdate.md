# MorpheusApi.PolicyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the policy | [optional] 
**description** | **String** | A description for the policy | [optional] 
**config** | [**AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject**](AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject.md) | A map of config values. The expected values vary by policy type. See &#x60;Retrieves all Policy Types&#x60; endpoint for &#x60;fieldName&#x60;(s) of required options. | [optional] 
**enabled** | **Boolean** | Set to false to disable | [optional] [default to true]
**refType** | **String** | Scope object type | [optional] 
**refId** | **Number** | Scope object ID (&#x60;group&#x60;,&#x60;cloud&#x60;,&#x60;user&#x60;, etc) | [optional] 
**accounts** | **[Number]** | Array of tenants to scope the policy to | [optional] 
**eachUser** | **Boolean** | Apply individually to each user in role.  Only when &#x60;refType&#x60; equals &#x60;Role&#x60; | [optional] 



## Enum: RefTypeEnum


* `ComputeSite` (value: `"ComputeSite"`)

* `ComputeZone` (value: `"ComputeZone"`)

* `User` (value: `"User"`)

* `Role` (value: `"Role"`)

* `Network` (value: `"Network"`)

* `Plan` (value: `"Plan"`)




