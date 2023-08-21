# ApiGroupsIdUpdateZonesGroup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zones** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) | An array of all the zones assigned to this group. | 

## Examples

- Prepare the resource
```powershell
$ApiGroupsIdUpdateZonesGroup = Initialize-PSOpenAPIToolsApiGroupsIdUpdateZonesGroup  -Zones [{&quot;id&quot;:1},{&quot;id&quot;:5}]
```

- Convert the resource to JSON
```powershell
$ApiGroupsIdUpdateZonesGroup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

