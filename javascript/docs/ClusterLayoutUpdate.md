# MorpheusApi.ClusterLayoutUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Cluster layout name | [optional] 
**description** | **String** | Cluster layout description | [optional] 
**labels** | **[String]** | Array of label strings, can be used for filtering. | [optional] 
**computeVersion** | **String** | Version of the cluster layout | [optional] 
**creatable** | **Boolean** | Can be used to enable / disable the creatability of the cluster layout. | [optional] [default to true]
**hasAutoScale** | **Boolean** | Can be used to enable / disable the horizontal scaling. | [optional] [default to false]
**installContainerRuntime** | **Boolean** | Install Docker (container runtime). | [optional] [default to false]
**memoryRequirement** | **Number** | Memory requirement in bytes | [optional] 
**groupType** | [**ClusterLayoutCreateGroupType**](ClusterLayoutCreateGroupType.md) |  | [optional] 
**provisionType** | [**ApiServicePlansServicePlanProvisionType**](ApiServicePlansServicePlanProvisionType.md) |  | [optional] 
**optionTypes** | [**[ApiBlueprintsIdUpdatePermissionsResourcePermissionSites]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of cluster layout option types | [optional] 
**taskSets** | [**[ApiBlueprintsIdUpdatePermissionsResourcePermissionSites]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of cluster layout task sets | [optional] 
**environmentVariables** | [**[ClusterLayoutCreateEnvironmentVariables]**](ClusterLayoutCreateEnvironmentVariables.md) | Array of cluster layout env variables | [optional] 
**masters** | [**[ClusterLayoutCreateMasters]**](ClusterLayoutCreateMasters.md) | Array of cluster layout master nodes | [optional] 
**workers** | [**[ClusterLayoutCreateMasters]**](ClusterLayoutCreateMasters.md) | Array of cluster layout worker nodes | [optional] 


