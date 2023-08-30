# AddNodeType200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerType** | [**ContainerType**](ContainerType.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddNodeType200Response = Initialize-PSOpenAPIToolsAddNodeType200Response  -ContainerType null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddNodeType200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

