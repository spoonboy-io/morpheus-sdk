

# ApiZonesIdZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A unique name scoped to your account for the cloud | 
**description** | **String** | Optional description field if you want to put more info there |  [optional]
**code** | **String** | Optional code for use with policies |  [optional]
**location** | **String** | Optional location for your cloud |  [optional]
**visibility** | [**VisibilityEnum**](#VisibilityEnum) | private or public |  [optional]
**zoneType** | **String** | Map containing code or id of the cloud type | 
**groupId** | **Long** | Specifies which Server group this cloud should be assigned to | 
**accountId** | **Long** | Specifies which Tenant this cloud should be assigned to |  [optional]
**enabled** | **Boolean** | Can be used to disable the cloud |  [optional]
**autoRecoverPowerState** | **Boolean** | Automatically Power on VMs |  [optional]
**scalePriority** | **Long** | Scale Priority |  [optional]
**linkedAccountId** | **Long** | Linked Account ID (enter commercial ID to get costing for AWS Govcloud) |  [optional]
**config** | **Object** | Map containing zone configuration settings. See the section on specific zone types for details. |  [optional]
**securityMode** | **String** | host firewall. &#x60;off&#x60; or &#x60;internal&#x60;. a.k.a. \&quot;local firewall\&quot; |  [optional]
**defaultCloudLogos** | **Boolean** | Can be used to clear any custom logo and darkLogo, reverting to the defaults for the cloud type |  [optional]
**credential** | **Object** | Map containing Credential ID. &#x60;local&#x60; means use the values set in the local cloud config instead of associating a credential. | 



## Enum: VisibilityEnum

Name | Value
---- | -----
PRIVATE | &quot;private&quot;
PUBLIC | &quot;public&quot;



