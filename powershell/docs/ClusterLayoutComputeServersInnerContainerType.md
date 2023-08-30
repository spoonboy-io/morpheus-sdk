# ClusterLayoutComputeServersInnerContainerType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Account** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Labels** | **String[]** |  | [optional] 
**ShortName** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**ContainerVersion** | **String** |  | [optional] 
**ProvisionType** | [**GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork**](GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork.md) |  | [optional] 
**VirtualImage** | **String** |  | [optional] 
**Category** | **String** |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**ContainerPorts** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**ContainerScripts** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**ContainerTemplates** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**EnvironmentVariables** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterLayoutComputeServersInnerContainerType = Initialize-PSOpenAPIToolsClusterLayoutComputeServersInnerContainerType  -Id null `
 -Account null `
 -Name null `
 -Labels null `
 -ShortName null `
 -Code null `
 -ContainerVersion null `
 -ProvisionType null `
 -VirtualImage null `
 -Category null `
 -Config null `
 -ContainerPorts null `
 -ContainerScripts null `
 -ContainerTemplates null `
 -EnvironmentVariables null
```

- Convert the resource to JSON
```powershell
$ClusterLayoutComputeServersInnerContainerType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

