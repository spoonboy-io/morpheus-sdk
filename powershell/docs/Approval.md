# Approval
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**ExternalName** | **String** |  | [optional] 
**RequestType** | **String** |  | [optional] 
**Account** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Approver** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**AccountIntegration** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**ErrorMessage** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**RequestBy** | **String** |  | [optional] 
**ApprovalItems** | [**ApprovalItemApprovalItem[]**](ApprovalItemApprovalItem.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Approval = Initialize-PSOpenAPIToolsApproval  -Id null `
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
 -RequestBy null `
 -ApprovalItems null
```

- Convert the resource to JSON
```powershell
$Approval | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

