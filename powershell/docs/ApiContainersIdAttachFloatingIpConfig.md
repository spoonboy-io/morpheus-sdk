# ApiContainersIdAttachFloatingIpConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsExternalNetworkId** | **String** | The Floating IP identifier in the format: &quot;&quot;ip-ID&quot;&quot; or &quot;&quot;pool-ID&quot;&quot;.  The Options API /api/options/openStack/openstackFloatingIpOptions?containerId&#x3D;{{containerId}} can be used to see which options are available.  | 
**FloatingIpBandwidth** | **Decimal** | Bandwidth (Mbit/s) Only the following cloud types support this parameter: Huawei, OpenTelekom  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiContainersIdAttachFloatingIpConfig = Initialize-PSOpenAPIToolsApiContainersIdAttachFloatingIpConfig  -OsExternalNetworkId ip-42 `
 -FloatingIpBandwidth 1024
```

- Convert the resource to JSON
```powershell
$ApiContainersIdAttachFloatingIpConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

