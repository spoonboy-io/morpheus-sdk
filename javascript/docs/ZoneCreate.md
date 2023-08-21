# MorpheusApi.ZoneCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A unique name scoped to your account for the cloud | 
**description** | **String** | Optional description field if you want to put more info there | [optional] 
**code** | **String** | Optional code for use with policies | [optional] 
**location** | **String** | Optional location for your cloud | [optional] 
**visibility** | **String** | private or public | [optional] [default to &#39;private&#39;]
**zoneType** | [**AnyOfobjectobject**](AnyOfobjectobject.md) |  | 
**groupId** | **Number** | Specifies which Server group this cloud should be assigned to | 
**accountId** | **Number** | Specifies which Tenant this cloud should be assigned to | [optional] 
**enabled** | **Boolean** | Can be used to disable the cloud | [optional] [default to true]
**autoRecoverPowerState** | **Boolean** | Automatically Power on VMs | [optional] [default to false]
**scalePriority** | **Number** | Scale Priority | [optional] [default to 1]
**linkedAccountId** | **Number** | Linked Account ID (enter commercial ID to get costing for AWS Govcloud) | [optional] 
**config** | **Object** | Map containing zone configuration settings. See the section on specific zone types for details. | [optional] 
**securityMode** | **String** | host firewall. &#x60;off&#x60; or &#x60;internal&#x60;. a.k.a. \&quot;local firewall\&quot; | [optional] [default to &#39;off&#39;]
**credential** | [**ZoneCreateCredential**](ZoneCreateCredential.md) |  | [optional] 



## Enum: VisibilityEnum


* `private` (value: `"private"`)

* `public` (value: `"public"`)




