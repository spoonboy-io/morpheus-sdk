# ContainerType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int32** |  | [optional] 
**Account** | [**ContainerTypeAccount**](ContainerTypeAccount.md) |  | [optional] 
**Name** | **String** |  | [optional] 
**Labels** | **String[]** |  | [optional] 
**ShortName** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**ContainerVersion** | **String** |  | [optional] 
**ProvisionType** | [**ContainerTypeProvisionType**](ContainerTypeProvisionType.md) |  | [optional] 
**VirtualImage** | [**ContainerTypeAccount**](ContainerTypeAccount.md) |  | [optional] 
**Category** | **String** |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**ContainerPorts** | [**ContainerTypeContainerPortsInner[]**](ContainerTypeContainerPortsInner.md) |  | [optional] 
**ContainerScripts** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**ContainerTemplates** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**EnvironmentVariables** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ContainerType = Initialize-PSOpenAPIToolsContainerType  -Id null `
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
$ContainerType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

