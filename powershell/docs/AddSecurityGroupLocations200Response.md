# AddSecurityGroupLocations200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupLocation** | [**SecurityGroupLocation**](SecurityGroupLocation.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddSecurityGroupLocations200Response = Initialize-PSOpenAPIToolsAddSecurityGroupLocations200Response  -SecurityGroupLocation null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddSecurityGroupLocations200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

