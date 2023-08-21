# ApiVdiPoolsIdVdiPoolConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name of instance | 
**Group** | [**OneOfstringlong**](OneOfstringlong.md) |  | 
**Cloud** | [**OneOfstringlong**](OneOfstringlong.md) |  | 
**Type** | [**OneOfstringlong**](OneOfstringlong.md) |  | 
**Layout** | [**OneOfstringlong**](OneOfstringlong.md) |  | 
**Plan** | [**OneOfstringlong**](OneOfstringlong.md) |  | 

## Examples

- Prepare the resource
```powershell
$ApiVdiPoolsIdVdiPoolConfig = Initialize-PSOpenAPIToolsApiVdiPoolsIdVdiPoolConfig  -Name null `
 -Group null `
 -Cloud null `
 -Type null `
 -Layout null `
 -Plan null
```

- Convert the resource to JSON
```powershell
$ApiVdiPoolsIdVdiPoolConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

