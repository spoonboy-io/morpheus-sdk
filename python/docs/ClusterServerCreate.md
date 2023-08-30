# ClusterServerCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Key for specific host type configuration  The config parameter is for configuration options that are specific to each Provision Type. The Provision Types api can be used to see which options are available.  | 
**name** | **str** | Name to be used for host(s) created in the cluster | 
**plan** | [**ClusterServerCreatePlan**](ClusterServerCreatePlan.md) |  | 
**server_type** | [**ClusterServerCreateServerType**](ClusterServerCreateServerType.md) |  | [optional] 
**volumes** | [**[ClusterServerCreateVolumesInner]**](ClusterServerCreateVolumesInner.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of Objects | [optional] 
**network_interfaces** | [**[ClusterServerCreateNetworkInterfacesInner]**](ClusterServerCreateNetworkInterfacesInner.md) | The networkInterfaces parameter is for network configuration.  The Options API /api/options/zoneNetworkOptions can be used to see which options are available.  It should be passed as an array of Objects with the following attributes  | [optional] 
**security_groups** | **[str]** | Key for security group configuration. | [optional] 
**visibility** | **str** | Visibility for server host | [optional]  if omitted the server will use the default value of "private"
**user_group** | [**ClusterServerCreateUserGroup**](ClusterServerCreateUserGroup.md) |  | [optional] 
**network_domain** | **str, none_type** | Network domain | [optional] 
**hostname** | **str, none_type** | Hostname for server host | [optional] 
**node_count** | **int** | Number of workers or hosts | [optional] 
**tags** | [**[UpdateHostManagedRequestServerTagsInner]**](UpdateHostManagedRequestServerTagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**labels** | **[str]** | Array of strings (keywords). This will set labels on the server and also on the cluster as well by default. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


