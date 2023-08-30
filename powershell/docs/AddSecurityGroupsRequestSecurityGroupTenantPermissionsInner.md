# AddSecurityGroupsRequestSecurityGroupTenantPermissionsInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | **Int64[]** | Array of tenant account ids that are allowed access | [optional] 
**CanManageAccounts** | **Int64[]** | Array of tenant account ids that can manage | [optional] 

## Examples

- Prepare the resource
```powershell
$AddSecurityGroupsRequestSecurityGroupTenantPermissionsInner = Initialize-PSOpenAPIToolsAddSecurityGroupsRequestSecurityGroupTenantPermissionsInner  -Accounts [1,3] `
 -CanManageAccounts [1,3]
```

- Convert the resource to JSON
```powershell
$AddSecurityGroupsRequestSecurityGroupTenantPermissionsInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

