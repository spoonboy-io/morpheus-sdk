# ClusterLayoutUpdate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Cluster layout name | [optional] 
**description** | **str, none_type** | Cluster layout description | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**compute_version** | **str** | Version of the cluster layout | [optional] 
**creatable** | **bool** | Can be used to enable / disable the creatability of the cluster layout. | [optional]  if omitted the server will use the default value of True
**has_auto_scale** | **bool** | Can be used to enable / disable the horizontal scaling. | [optional]  if omitted the server will use the default value of False
**install_container_runtime** | **bool** | Install Docker (container runtime). | [optional]  if omitted the server will use the default value of False
**memory_requirement** | **int** | Memory requirement in bytes | [optional] 
**group_type** | [**ClusterLayoutCreateGroupType**](ClusterLayoutCreateGroupType.md) |  | [optional] 
**provision_type** | [**AddServicePlansRequestServicePlanProvisionTypeInner**](AddServicePlansRequestServicePlanProvisionTypeInner.md) |  | [optional] 
**option_types** | [**[UpdateBlueprintPermissionsRequestResourcePermissionSitesInner]**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) | Array of cluster layout option types | [optional] 
**task_sets** | [**[UpdateBlueprintPermissionsRequestResourcePermissionSitesInner]**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) | Array of cluster layout task sets | [optional] 
**environment_variables** | [**[ClusterLayoutCreateEnvironmentVariablesInner]**](ClusterLayoutCreateEnvironmentVariablesInner.md) | Array of cluster layout env variables | [optional] 
**masters** | [**[ClusterLayoutCreateMastersInner]**](ClusterLayoutCreateMastersInner.md) | Array of cluster layout master nodes | [optional] 
**workers** | [**[ClusterLayoutCreateMastersInner]**](ClusterLayoutCreateMastersInner.md) | Array of cluster layout worker nodes | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


