# GuidanceVmwareSizingZone
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**ZoneType** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceVmwareSizingZone = Initialize-PSOpenAPIToolsGuidanceVmwareSizingZone  -Id null `
 -Name null `
 -ZoneType null
```

- Convert the resource to JSON
```powershell
$GuidanceVmwareSizingZone | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

