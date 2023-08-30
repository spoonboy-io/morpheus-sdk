# AddSecurityGroupLocationsRequestSecurityGroupLocation
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | **Int64** | The ID of the Zone (Cloud) | 
**CustomOptions** | [**AddSecurityGroupsRequestSecurityGroupCustomOptions**](AddSecurityGroupsRequestSecurityGroupCustomOptions.md) |  | 

## Examples

- Prepare the resource
```powershell
$AddSecurityGroupLocationsRequestSecurityGroupLocation = Initialize-PSOpenAPIToolsAddSecurityGroupLocationsRequestSecurityGroupLocation  -ZoneId 1 `
 -CustomOptions null
```

- Convert the resource to JSON
```powershell
$AddSecurityGroupLocationsRequestSecurityGroupLocation | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

