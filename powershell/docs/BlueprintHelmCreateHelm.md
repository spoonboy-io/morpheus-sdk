# BlueprintHelmCreateHelm
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **String** | Configuration Type | 
**Git** | [**BlueprintHelmCreateHelmGit**](BlueprintHelmCreateHelmGit.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintHelmCreateHelm = Initialize-PSOpenAPIToolsBlueprintHelmCreateHelm  -ConfigType null `
 -Git null
```

- Convert the resource to JSON
```powershell
$BlueprintHelmCreateHelm | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

