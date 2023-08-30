# UpdateCloudsRequestZone


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A unique name scoped to your account for the cloud | 
**group_id** | **int** | Specifies which Server group this cloud should be assigned to | 
**zone_type** | **str** | Map containing code or id of the cloud type | defaults to "standard"
**credential** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Map containing Credential ID. &#x60;local&#x60; means use the values set in the local cloud config instead of associating a credential. | defaults to {"type":"local"}
**description** | **str** | Optional description field if you want to put more info there | [optional] 
**code** | **str** | Optional code for use with policies | [optional] 
**location** | **str** | Optional location for your cloud | [optional] 
**visibility** | **str** | private or public | [optional]  if omitted the server will use the default value of "private"
**account_id** | **int** | Specifies which Tenant this cloud should be assigned to | [optional] 
**enabled** | **bool** | Can be used to disable the cloud | [optional]  if omitted the server will use the default value of True
**auto_recover_power_state** | **bool** | Automatically Power on VMs | [optional]  if omitted the server will use the default value of False
**scale_priority** | **int** | Scale Priority | [optional]  if omitted the server will use the default value of 1
**linked_account_id** | **int** | Linked Account ID (enter commercial ID to get costing for AWS Govcloud) | [optional] 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Map containing zone configuration settings. See the section on specific zone types for details. | [optional] 
**security_mode** | **str** | host firewall. &#x60;off&#x60; or &#x60;internal&#x60;. a.k.a. \&quot;local firewall\&quot; | [optional]  if omitted the server will use the default value of "off"
**default_cloud_logos** | **bool** | Can be used to clear any custom logo and darkLogo, reverting to the defaults for the cloud type | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


