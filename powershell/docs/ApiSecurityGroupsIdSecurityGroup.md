# ApiSecurityGroupsIdSecurityGroup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name for your security group | [optional] 
**Description** | **String** | Optional description field | [optional] 
**Active** | **Boolean** | Set to &#x60;false&#x60; to disable a security group. | [optional] 
**TenantPermissions** | [**ApiSecurityGroupsSecurityGroupTenantPermissions[]**](ApiSecurityGroupsSecurityGroupTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | [**ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions**](ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiSecurityGroupsIdSecurityGroup = Initialize-PSOpenAPIToolsApiSecurityGroupsIdSecurityGroup  -Name null `
 -Description null `
 -Active null `
 -TenantPermissions null `
 -ResourcePermissions null
```

- Convert the resource to JSON
```powershell
$ApiSecurityGroupsIdSecurityGroup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

