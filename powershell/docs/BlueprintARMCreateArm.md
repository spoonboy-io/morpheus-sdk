# BlueprintARMCreateArm
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **String** | Configuration Type | 
**Json** | **String** | ARM Template in JSON | [optional] 
**Yaml** | **String** | ARM Template in YAML | [optional] 
**Git** | [**BlueprintARMCreateArmGit**](BlueprintARMCreateArmGit.md) |  | [optional] 
**OsType** | **String** | OS Type | [optional] 
**InstallAgent** | [**BlueprintARMCreateArmInstallAgent**](BlueprintARMCreateArmInstallAgent.md) |  | [optional] 
**CloudInitEnabled** | [**BlueprintARMCreateArmCloudInitEnabled**](BlueprintARMCreateArmCloudInitEnabled.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintARMCreateArm = Initialize-PSOpenAPIToolsBlueprintARMCreateArm  -ConfigType null `
 -Json null `
 -Yaml null `
 -Git null `
 -OsType null `
 -InstallAgent null `
 -CloudInitEnabled null
```

- Convert the resource to JSON
```powershell
$BlueprintARMCreateArm | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

