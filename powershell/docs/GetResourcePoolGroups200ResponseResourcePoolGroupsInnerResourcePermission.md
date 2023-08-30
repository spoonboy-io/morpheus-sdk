# GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | **Boolean** | Pass &#x60;true&#x60; to allow access all groups | [optional] [default to $true]
**Sites** | [**GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermissionSitesInner[]**](GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermissionSitesInner.md) | Array of groups that are allowed access | [optional] 

## Examples

- Prepare the resource
```powershell
$GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission = Initialize-PSOpenAPIToolsGetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission  -All null `
 -Sites null
```

- Convert the resource to JSON
```powershell
$GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

