

# ClusterLayoutCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Cluster layout name | 
**description** | **String** | Cluster layout description |  [optional]
**labels** | **List&lt;String&gt;** |  |  [optional]
**computeVersion** | **String** | Version of the cluster layout | 
**creatable** | **Boolean** | Can be used to enable / disable the creatability of the cluster layout. |  [optional]
**hasAutoScale** | **Boolean** | Can be used to enable / disable the horizontal scaling. |  [optional]
**installContainerRuntime** | **Boolean** | Install Docker (container runtime). |  [optional]
**memoryRequirement** | **Long** | Memory requirement in bytes |  [optional]
**groupType** | [**ClusterLayoutCreateGroupType**](ClusterLayoutCreateGroupType.md) |  | 
**provisionType** | [**ApiServicePlansServicePlanProvisionType**](ApiServicePlansServicePlanProvisionType.md) |  | 
**optionTypes** | [**List&lt;ApiBlueprintsIdUpdatePermissionsResourcePermissionSites&gt;**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of cluster layout option types |  [optional]
**taskSets** | [**List&lt;ApiBlueprintsIdUpdatePermissionsResourcePermissionSites&gt;**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of cluster layout task sets |  [optional]
**environmentVariables** | [**List&lt;ClusterLayoutCreateEnvironmentVariables&gt;**](ClusterLayoutCreateEnvironmentVariables.md) | Array of cluster layout env variables |  [optional]
**masters** | [**List&lt;ClusterLayoutCreateMasters&gt;**](ClusterLayoutCreateMasters.md) | Array of cluster layout master nodes |  [optional]
**workers** | [**List&lt;ClusterLayoutCreateMasters&gt;**](ClusterLayoutCreateMasters.md) | Array of cluster layout worker nodes |  [optional]



