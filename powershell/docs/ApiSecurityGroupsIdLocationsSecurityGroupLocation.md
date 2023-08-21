# ApiSecurityGroupsIdLocationsSecurityGroupLocation
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | **Int64** | The ID of the Zone (Cloud) | 
**CustomOptions** | [**AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions**](AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions.md) |  | 

## Examples

- Prepare the resource
```powershell
$ApiSecurityGroupsIdLocationsSecurityGroupLocation = Initialize-PSOpenAPIToolsApiSecurityGroupsIdLocationsSecurityGroupLocation  -ZoneId 1 `
 -CustomOptions null
```

- Convert the resource to JSON
```powershell
$ApiSecurityGroupsIdLocationsSecurityGroupLocation | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

