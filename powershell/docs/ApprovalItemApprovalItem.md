# ApprovalItemApprovalItem
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**ExternalName** | **String** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**ApprovedBy** | **String** |  | [optional] 
**DeniedBy** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**ErrorMessage** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**DateApproved** | **System.DateTime** |  | [optional] 
**DateDenied** | **System.DateTime** |  | [optional] 
**Approval** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional] 
**Reference** | [**ApprovalItemApprovalItemReference**](ApprovalItemApprovalItemReference.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApprovalItemApprovalItem = Initialize-PSOpenAPIToolsApprovalItemApprovalItem  -Id null `
 -Name null `
 -ExternalId null `
 -ExternalName null `
 -InternalId null `
 -ApprovedBy null `
 -DeniedBy null `
 -Status null `
 -ErrorMessage null `
 -DateCreated null `
 -LastUpdated null `
 -DateApproved null `
 -DateDenied null `
 -Approval null `
 -Reference null
```

- Convert the resource to JSON
```powershell
$ApprovalItemApprovalItem | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

