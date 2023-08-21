# Approvals
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**ExternalName** | **String** |  | [optional] 
**RequestType** | **String** |  | [optional] 
**Account** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**Approver** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**AccountIntegration** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**ErrorMessage** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**RequestBy** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Approvals = Initialize-PSOpenAPIToolsApprovals  -Id null `
 -Name null `
 -InternalId null `
 -ExternalId null `
 -ExternalName null `
 -RequestType null `
 -Account null `
 -Approver null `
 -AccountIntegration null `
 -Status null `
 -ErrorMessage null `
 -DateCreated null `
 -LastUpdated null `
 -RequestBy null
```

- Convert the resource to JSON
```powershell
$Approvals | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

