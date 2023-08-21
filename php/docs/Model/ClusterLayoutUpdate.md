# # ClusterLayoutUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Cluster layout name | [optional]
**description** | **string** | Cluster layout description | [optional]
**labels** | **string[]** | Array of label strings, can be used for filtering. | [optional]
**compute_version** | **string** | Version of the cluster layout | [optional]
**creatable** | **bool** | Can be used to enable / disable the creatability of the cluster layout. | [optional] [default to true]
**has_auto_scale** | **bool** | Can be used to enable / disable the horizontal scaling. | [optional] [default to false]
**install_container_runtime** | **bool** | Install Docker (container runtime). | [optional] [default to false]
**memory_requirement** | **int** | Memory requirement in bytes | [optional]
**group_type** | [**\OpenAPI\Client\Model\ClusterLayoutCreateGroupType**](ClusterLayoutCreateGroupType.md) |  | [optional]
**provision_type** | [**\OpenAPI\Client\Model\ApiServicePlansServicePlanProvisionType**](ApiServicePlansServicePlanProvisionType.md) |  | [optional]
**option_types** | [**\OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites[]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of cluster layout option types | [optional]
**task_sets** | [**\OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites[]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of cluster layout task sets | [optional]
**environment_variables** | [**\OpenAPI\Client\Model\ClusterLayoutCreateEnvironmentVariables[]**](ClusterLayoutCreateEnvironmentVariables.md) | Array of cluster layout env variables | [optional]
**masters** | [**\OpenAPI\Client\Model\ClusterLayoutCreateMasters[]**](ClusterLayoutCreateMasters.md) | Array of cluster layout master nodes | [optional]
**workers** | [**\OpenAPI\Client\Model\ClusterLayoutCreateMasters[]**](ClusterLayoutCreateMasters.md) | Array of cluster layout worker nodes | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
