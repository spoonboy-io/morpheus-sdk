# MorpheusApi.ClusterLayoutCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Cluster layout name | 
**description** | **String** | Cluster layout description | [optional] 
**labels** | **[String]** |  | [optional] 
**computeVersion** | **String** | Version of the cluster layout | 
**creatable** | **Boolean** | Can be used to enable / disable the creatability of the cluster layout. | [optional] [default to true]
**hasAutoScale** | **Boolean** | Can be used to enable / disable the horizontal scaling. | [optional] [default to false]
**installContainerRuntime** | **Boolean** | Install Docker (container runtime). | [optional] [default to false]
**memoryRequirement** | **Number** | Memory requirement in bytes | [optional] 
**groupType** | [**ClusterLayoutCreateGroupType**](ClusterLayoutCreateGroupType.md) |  | 
**provisionType** | [**ApiServicePlansServicePlanProvisionType**](ApiServicePlansServicePlanProvisionType.md) |  | 
**optionTypes** | [**[ApiBlueprintsIdUpdatePermissionsResourcePermissionSites]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of cluster layout option types | [optional] 
**taskSets** | [**[ApiBlueprintsIdUpdatePermissionsResourcePermissionSites]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of cluster layout task sets | [optional] 
**environmentVariables** | [**[ClusterLayoutCreateEnvironmentVariables]**](ClusterLayoutCreateEnvironmentVariables.md) | Array of cluster layout env variables | [optional] 
**masters** | [**[ClusterLayoutCreateMasters]**](ClusterLayoutCreateMasters.md) | Array of cluster layout master nodes | [optional] 
**workers** | [**[ClusterLayoutCreateMasters]**](ClusterLayoutCreateMasters.md) | Array of cluster layout worker nodes | [optional] 


