# CatalogItemTypeInstanceScribePortsInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **Int64** | Port number. | [optional] 
**Name** | **String** | A name for the port. | [optional] 
**Lb** | **String** | The load balancer protocol. HTTP, HTTPS, or TCP. | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogItemTypeInstanceScribePortsInner = Initialize-PSOpenAPIToolsCatalogItemTypeInstanceScribePortsInner  -Port 8080 `
 -Name null `
 -Lb null
```

- Convert the resource to JSON
```powershell
$CatalogItemTypeInstanceScribePortsInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

