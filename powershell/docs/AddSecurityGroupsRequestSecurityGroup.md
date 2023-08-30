# AddSecurityGroupsRequestSecurityGroup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name for your security group | 
**Description** | **String** | Optional description field | [optional] 
**ZoneId** | **Int64** | Scoped Cloud ID | 
**Active** | **Boolean** | Set to &#x60;false&#x60; to disable a security group. | [optional] 
**CustomOptions** | [**AddSecurityGroupsRequestSecurityGroupCustomOptions**](AddSecurityGroupsRequestSecurityGroupCustomOptions.md) |  | [optional] 
**TenantPermissions** | [**AddSecurityGroupsRequestSecurityGroupTenantPermissionsInner[]**](AddSecurityGroupsRequestSecurityGroupTenantPermissionsInner.md) |  | [optional] 
**ResourcePermissions** | [**UpdateCloudDatastoresRequestDatastoreResourcePermissions**](UpdateCloudDatastoresRequestDatastoreResourcePermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddSecurityGroupsRequestSecurityGroup = Initialize-PSOpenAPIToolsAddSecurityGroupsRequestSecurityGroup  -Name null `
 -Description null `
 -ZoneId 3 `
 -Active null `
 -CustomOptions null `
 -TenantPermissions null `
 -ResourcePermissions null
```

- Convert the resource to JSON
```powershell
$AddSecurityGroupsRequestSecurityGroup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

