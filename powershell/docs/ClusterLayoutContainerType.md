# ClusterLayoutContainerType
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
**ProvisionType** | [**InlineResponse20094Network**](InlineResponse20094Network.md) |  | [optional] 
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
$ClusterLayoutContainerType = Initialize-PSOpenAPIToolsClusterLayoutContainerType  -Id null `
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
$ClusterLayoutContainerType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

