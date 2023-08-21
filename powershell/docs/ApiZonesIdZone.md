# ApiZonesIdZone
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name scoped to your account for the cloud | 
**Description** | **String** | Optional description field if you want to put more info there | [optional] 
**Code** | **String** | Optional code for use with policies | [optional] 
**Location** | **String** | Optional location for your cloud | [optional] 
**Visibility** | **String** | private or public | [optional] [default to "private"]
**ZoneType** | **String** | Map containing code or id of the cloud type | [default to "standard"]
**GroupId** | **Int64** | Specifies which Server group this cloud should be assigned to | 
**AccountId** | **Int64** | Specifies which Tenant this cloud should be assigned to | [optional] 
**Enabled** | **Boolean** | Can be used to disable the cloud | [optional] [default to $true]
**AutoRecoverPowerState** | **Boolean** | Automatically Power on VMs | [optional] [default to $false]
**ScalePriority** | **Int64** | Scale Priority | [optional] [default to 1]
**LinkedAccountId** | **Int64** | Linked Account ID (enter commercial ID to get costing for AWS Govcloud) | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) | Map containing zone configuration settings. See the section on specific zone types for details. | [optional] 
**SecurityMode** | **String** | host firewall. &#x60;off&#x60; or &#x60;internal&#x60;. a.k.a. &quot;&quot;local firewall&quot;&quot; | [optional] [default to "off"]
**DefaultCloudLogos** | **Boolean** | Can be used to clear any custom logo and darkLogo, reverting to the defaults for the cloud type | [optional] 
**Credential** | [**SystemCollectionsHashtable**](.md) | Map containing Credential ID. &#x60;local&#x60; means use the values set in the local cloud config instead of associating a credential. | 

## Examples

- Prepare the resource
```powershell
$ApiZonesIdZone = Initialize-PSOpenAPIToolsApiZonesIdZone  -Name My Cloud `
 -Description null `
 -Code mycloud `
 -Location US East `
 -Visibility null `
 -ZoneType null `
 -GroupId 3 `
 -AccountId 1 `
 -Enabled null `
 -AutoRecoverPowerState null `
 -ScalePriority null `
 -LinkedAccountId null `
 -Config null `
 -SecurityMode null `
 -DefaultCloudLogos null `
 -Credential {&quot;id&quot;:1}
```

- Convert the resource to JSON
```powershell
$ApiZonesIdZone | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

