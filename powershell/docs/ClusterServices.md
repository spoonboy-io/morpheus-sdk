# ClusterServices
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Type** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**ExternalIp** | **String** |  | [optional] 
**InternalIp** | **String** |  | [optional] 
**ExternalPort** | **String** |  | [optional] 
**InternalPort** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterServices = Initialize-PSOpenAPIToolsClusterServices  -Id null `
 -Name null `
 -Type null `
 -Code null `
 -ExternalIp null `
 -InternalIp null `
 -ExternalPort null `
 -InternalPort null `
 -Status null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$ClusterServices | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

