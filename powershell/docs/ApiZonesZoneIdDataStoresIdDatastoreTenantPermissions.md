# ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | **Int64[]** | Array of tenant account ids that are allowed access | [optional] 
**DefaultTarget** | **Int64[]** | Array of tenant account ids which should use the data store as the Default | [optional] 
**DefaultStore** | **Int64[]** | Array of tenant account ids which should use the data store as the Image Target | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions = Initialize-PSOpenAPIToolsApiZonesZoneIdDataStoresIdDatastoreTenantPermissions  -Accounts [1,3] `
 -DefaultTarget [1,3] `
 -DefaultStore [1,3]
```

- Convert the resource to JSON
```powershell
$ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

