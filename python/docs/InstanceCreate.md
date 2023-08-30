# InstanceCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**instance** | [**InstanceCreateInstance**](InstanceCreateInstance.md) |  | 
**config** | [**CatalogItemTypeInstanceScribeConfig**](CatalogItemTypeInstanceScribeConfig.md) |  | 
**zone_id** | **int** | The Cloud ID to provision the instance onto. | [optional] 
**evars** | [**[UpdateHostManagedRequestServerTagsInner]**](UpdateHostManagedRequestServerTagsInner.md) | Environment Variables, an array of objects that have name and value. | [optional] 
**copies** | **int** | Number of copies to provision. | [optional]  if omitted the server will use the default value of 1
**layout_size** | **int** | Apply a multiply factor of containers/vms within the instance. | [optional]  if omitted the server will use the default value of 1
**service_plan_options** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Map of custom options depending on selected service plan. | [optional] 
**security_groups** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. | [optional] 
**volumes** | [**[InstanceCreateVolume]**](InstanceCreateVolume.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of | [optional] 
**network_interfaces** | [**[InstanceCreateNetwork]**](InstanceCreateNetwork.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  | [optional] 
**labels** | **[str]** | Array of strings (keywords). | [optional] 
**tags** | [**[UpdateHostManagedRequestServerTagsInner]**](UpdateHostManagedRequestServerTagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**metadata** | [**[UpdateHostManagedRequestServerTagsInner]**](UpdateHostManagedRequestServerTagsInner.md) | Alias for &#x60;tags&#x60;. | [optional] 
**ports** | [**[InstanceCreatePortsInner]**](InstanceCreatePortsInner.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  | [optional] 
**task_set_id** | **int** | The Workflow ID to execute. | [optional] 
**task_set_name** | **str** | The Workflow Name to execute. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


