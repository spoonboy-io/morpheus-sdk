# ApiPricesPriceVolumeType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Volume type ID, required for &#x60;storage&#x60; price type. The endpoint /api/prices/volume-types provides a list of available volume type options.  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiPricesPriceVolumeType = Initialize-PSOpenAPIToolsApiPricesPriceVolumeType  -Id null
```

- Convert the resource to JSON
```powershell
$ApiPricesPriceVolumeType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

