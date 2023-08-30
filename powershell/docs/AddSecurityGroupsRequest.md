# AddSecurityGroupsRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroup** | [**AddSecurityGroupsRequestSecurityGroup**](AddSecurityGroupsRequestSecurityGroup.md) |  | 

## Examples

- Prepare the resource
```powershell
$AddSecurityGroupsRequest = Initialize-PSOpenAPIToolsAddSecurityGroupsRequest  -SecurityGroup null
```

- Convert the resource to JSON
```powershell
$AddSecurityGroupsRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

