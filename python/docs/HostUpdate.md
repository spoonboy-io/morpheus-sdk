# HostUpdate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Unique name scoped to your account for the server. | [optional] 
**description** | **str** | Optional description field. | [optional] 
**ssh_username** | **str** | SSH Username | [optional] 
**ssh_password** | **str** | SSH Password | [optional] 
**power_schedule_type** | **int** | Power schedule ID. | [optional] 
**labels** | **[str]** |  | [optional] 
**tags** | [**[UpdateHostManagedRequestServerTagsInner]**](UpdateHostManagedRequestServerTagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**add_tags** | [**[UpdateHostManagedRequestServerTagsInner]**](UpdateHostManagedRequestServerTagsInner.md) | Add or update value of Metadata tags, Array of objects having a name and value. | [optional] 
**remove_tags** | [**[InstanceUpdateInstanceRemoveTagsInner]**](InstanceUpdateInstanceRemoveTagsInner.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. | [optional] 
**guest_console_type** | **str** | The Type of guest console this server provides such as disabled, vnc, rdp, ssh | [optional] 
**guest_console_username** | **str** | The optional guest console username if you don&#39;t want to use the user defaults | [optional] 
**guest_console_password** | **str** | The optional guest console password if not using the accessing users creds | [optional] 
**guest_console_port** | **str** | The port the guest console is being accessed from | [optional] 
**guest_console_preferred** | **bool** | Can turn off guest console preferences on server in favor of hypervisor console | [optional]  if omitted the server will use the default value of True
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


