# AddNodeTypeRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerType** | [**ContainerTypeCreate**](ContainerTypeCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddNodeTypeRequest = Initialize-PSOpenAPIToolsAddNodeTypeRequest  -ContainerType null
```

- Convert the resource to JSON
```powershell
$AddNodeTypeRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

