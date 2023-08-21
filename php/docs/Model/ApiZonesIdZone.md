# # ApiZonesIdZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A unique name scoped to your account for the cloud |
**description** | **string** | Optional description field if you want to put more info there | [optional]
**code** | **string** | Optional code for use with policies | [optional]
**location** | **string** | Optional location for your cloud | [optional]
**visibility** | **string** | private or public | [optional] [default to 'private']
**zone_type** | **string** | Map containing code or id of the cloud type | [default to 'standard']
**group_id** | **int** | Specifies which Server group this cloud should be assigned to |
**account_id** | **int** | Specifies which Tenant this cloud should be assigned to | [optional]
**enabled** | **bool** | Can be used to disable the cloud | [optional] [default to true]
**auto_recover_power_state** | **bool** | Automatically Power on VMs | [optional] [default to false]
**scale_priority** | **int** | Scale Priority | [optional] [default to 1]
**linked_account_id** | **int** | Linked Account ID (enter commercial ID to get costing for AWS Govcloud) | [optional]
**config** | **object** | Map containing zone configuration settings. See the section on specific zone types for details. | [optional]
**security_mode** | **string** | host firewall. &#x60;off&#x60; or &#x60;internal&#x60;. a.k.a. \&quot;local firewall\&quot; | [optional] [default to 'off']
**default_cloud_logos** | **bool** | Can be used to clear any custom logo and darkLogo, reverting to the defaults for the cloud type | [optional]
**credential** | **object** | Map containing Credential ID. &#x60;local&#x60; means use the values set in the local cloud config instead of associating a credential. |

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
