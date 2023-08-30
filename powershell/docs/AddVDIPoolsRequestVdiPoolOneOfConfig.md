# AddVDIPoolsRequestVdiPoolOneOfConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name of instance | 
**Group** | [**AddVDIPoolsRequestVdiPoolOneOfConfigGroup**](AddVDIPoolsRequestVdiPoolOneOfConfigGroup.md) |  | 
**Cloud** | [**AddVDIPoolsRequestVdiPoolOneOfConfigCloud**](AddVDIPoolsRequestVdiPoolOneOfConfigCloud.md) |  | 
**Type** | [**AddVDIPoolsRequestVdiPoolOneOfConfigType**](AddVDIPoolsRequestVdiPoolOneOfConfigType.md) |  | 
**Layout** | [**AddVDIPoolsRequestVdiPoolOneOfConfigLayout**](AddVDIPoolsRequestVdiPoolOneOfConfigLayout.md) |  | 
**Plan** | [**AddVDIPoolsRequestVdiPoolOneOfConfigPlan**](AddVDIPoolsRequestVdiPoolOneOfConfigPlan.md) |  | 

## Examples

- Prepare the resource
```powershell
$AddVDIPoolsRequestVdiPoolOneOfConfig = Initialize-PSOpenAPIToolsAddVDIPoolsRequestVdiPoolOneOfConfig  -Name null `
 -Group null `
 -Cloud null `
 -Type null `
 -Layout null `
 -Plan null
```

- Convert the resource to JSON
```powershell
$AddVDIPoolsRequestVdiPoolOneOfConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

