# ClusterLayoutSpecTemplatesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Account** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Labels** | **String[]** |  | [optional] 
**Code** | **String** |  | [optional] 
**Type** | [**GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork**](GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork.md) |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**ExternalType** | **String** |  | [optional] 
**DeploymentId** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**File** | [**ClusterLayoutSpecTemplatesInnerFile**](ClusterLayoutSpecTemplatesInnerFile.md) |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**CreatedBy** | **String** |  | [optional] 
**UpdatedBy** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterLayoutSpecTemplatesInner = Initialize-PSOpenAPIToolsClusterLayoutSpecTemplatesInner  -Id null `
 -Account null `
 -Name null `
 -Labels null `
 -Code null `
 -Type null `
 -ExternalId null `
 -ExternalType null `
 -DeploymentId null `
 -Status null `
 -File null `
 -Config null `
 -CreatedBy null `
 -UpdatedBy null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$ClusterLayoutSpecTemplatesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

