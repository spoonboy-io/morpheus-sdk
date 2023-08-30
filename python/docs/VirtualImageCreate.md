# VirtualImageCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A name for the virtual image | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**image_type** | **str** | Code of image type. eg. vmware, ami, etc. | [optional] 
**storage_provider** | [**VirtualImageCreateStorageProvider**](VirtualImageCreateStorageProvider.md) |  | [optional] 
**is_cloud_init** | **bool** | Cloud Init Enabled? | [optional]  if omitted the server will use the default value of False
**user_data** | **str, none_type** | Cloud-Init User Data, a bash script | [optional] 
**install_agent** | **bool** | Install Agent? | [optional]  if omitted the server will use the default value of False
**ssh_username** | **str, none_type** | SSH Username | [optional] 
**ssh_password** | **str, none_type** | SSH Password | [optional] 
**ssh_key** | **str, none_type** | SSH Key | [optional] 
**os_type** | [**VirtualImageCreateOsType**](VirtualImageCreateOsType.md) |  | [optional] 
**visibility** | **str** | private or public | [optional]  if omitted the server will use the default value of "private"
**accounts** | **[int]** |  | [optional] 
**is_auto_join_domain** | **bool** | Auto Join Domain? | [optional]  if omitted the server will use the default value of False
**virtio_supported** | **bool** | VirtIO Drivers Loaded? | [optional]  if omitted the server will use the default value of True
**vm_tools_installed** | **bool** | VM Tools Installed? | [optional]  if omitted the server will use the default value of True
**is_force_customization** | **bool** | Force Guest Customization? | [optional]  if omitted the server will use the default value of False
**trial_version** | **bool** | Trial Version | [optional]  if omitted the server will use the default value of False
**is_sysprep** | **bool** | Sysprep Enabled? | [optional]  if omitted the server will use the default value of False
**config** | [**VirtualImageCreateConfig**](VirtualImageCreateConfig.md) |  | [optional] 
**tags** | [**[VirtualImageCreateTagsInner]**](VirtualImageCreateTagsInner.md) | Metadata tags, Array of objects having a name and value | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


