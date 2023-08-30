# AddCheckGroupsRequestCheckGroup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Unique name scoped to your account for the check group | 
**Description** | **String** | Optional description field | [optional] 
**MinHappy** | **Int32** | This specifies the minimum number of checks within the group that must be happy to keep the group from becoming unhealthy. | [optional] [default to 1]
**InUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations | [optional] [default to $true]
**Severity** | **String** | Determines the maximum severity level this group can incur on an incident when failing | [optional] [default to "critical"]
**Active** | **Boolean** | Used to determine if check group is active | [optional] [default to $true]
**Checks** | **Int32[]** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddCheckGroupsRequestCheckGroup = Initialize-PSOpenAPIToolsAddCheckGroupsRequestCheckGroup  -Name My Check Group `
 -Description My cool description `
 -MinHappy null `
 -InUptime null `
 -Severity null `
 -Active null `
 -Checks null
```

- Convert the resource to JSON
```powershell
$AddCheckGroupsRequestCheckGroup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

