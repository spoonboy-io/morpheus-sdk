# CatalogItemTypeInstanceScribe


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group** | [**CatalogItemTypeInstanceScribeGroup**](CatalogItemTypeInstanceScribeGroup.md) |  | 
**cloud** | [**CatalogItemTypeInstanceScribeCloud**](CatalogItemTypeInstanceScribeCloud.md) |  | 
**type** | **str** | The type of instance by code we want to fetch. | 
**name** | **str** | Name of the instance to be created. | 
**config** | [**CatalogItemTypeInstanceScribeConfig**](CatalogItemTypeInstanceScribeConfig.md) |  | 
**volumes** | [**[InstanceCreateVolume]**](InstanceCreateVolume.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of | 
**layout** | [**CatalogItemTypeInstanceScribeLayout**](CatalogItemTypeInstanceScribeLayout.md) |  | 
**plan** | [**CatalogItemTypeInstanceScribePlan**](CatalogItemTypeInstanceScribePlan.md) |  | 
**host_name** | **str** | Hostname of the instance to be created.  Can be the same as the instance name. | [optional] 
**environment** | **str** | Environment code | [optional] 
**version** | **str** | Version of the layout to create. | [optional] 
**evars** | [**[UpdateHostManagedRequestServerTagsInner]**](UpdateHostManagedRequestServerTagsInner.md) | Environment Variables, an array of objects that have name and value. | [optional] 
**service_plan_options** | [**ServicePlanOptions**](ServicePlanOptions.md) |  | [optional] 
**security_groups** | [**[CatalogItemTypeInstanceScribeSecurityGroupsInner], none_type**](CatalogItemTypeInstanceScribeSecurityGroupsInner.md) | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. | [optional] 
**network_interfaces** | [**[InstanceCreateNetwork]**](InstanceCreateNetwork.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  | [optional] 
**labels** | **[str]** | Array of strings (keywords). | [optional] 
**tags** | [**[UpdateHostManagedRequestServerTagsInner]**](UpdateHostManagedRequestServerTagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**metadata** | [**[UpdateHostManagedRequestServerTagsInner]**](UpdateHostManagedRequestServerTagsInner.md) | Alias for &#x60;tags&#x60;. | [optional] 
**ports** | [**[CatalogItemTypeInstanceScribePortsInner]**](CatalogItemTypeInstanceScribePortsInner.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  | [optional] 
**task_set_id** | **int** | The Workflow ID to execute. | [optional] 
**task_set_name** | **str** | The Workflow Name to execute. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


