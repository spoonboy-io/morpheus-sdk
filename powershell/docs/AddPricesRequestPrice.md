# AddPricesRequestPrice
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Price name | 
**Code** | **String** | Price code, must be unique | 
**Account** | [**AddPricesRequestPriceAccount**](AddPricesRequestPriceAccount.md) |  | [optional] 
**PriceType** | **String** | Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software * &#x60;load_balancer&#x60; - Load Balancer * &#x60;load_balancer_virtual_server&#x60; - Load Balancer Virtual Server  | 
**PriceUnit** | **String** | The unit of pricing | 
**IncurCharges** | **String** | Indicates when to incur charge | 
**Currency** | **String** | ISO Currency code | 
**Cost** | **Decimal** | Cost | 
**MarkupType** | **String** | Price adjustment type | [optional] 
**Markup** | **Decimal** | Amount for &#x60;fixed&#x60; price adjustment type | [optional] 
**MarkupPercent** | **Decimal** | Percent for &#x60;percent&#x60; price adjustment type | [optional] 
**CustomPrice** | **Decimal** | Custom price for &#x60;custom&#x60; price adjustment type | [optional] 
**Platform** | **String** | Platform.  Required for &#x60;platform&#x60; price type | [optional] 
**Software** | **String** | Software.  Required for software price type | [optional] 
**VolumeType** | [**AddPricesRequestPriceVolumeType**](AddPricesRequestPriceVolumeType.md) |  | [optional] 
**Datastore** | [**AddPricesRequestPriceDatastore**](AddPricesRequestPriceDatastore.md) |  | [optional] 
**CrossCloudApply** | **Boolean** | Apply price across clouds, optional true/false flag for datastore price type | [optional] 

## Examples

- Prepare the resource
```powershell
$AddPricesRequestPrice = Initialize-PSOpenAPIToolsAddPricesRequestPrice  -Name null `
 -Code null `
 -Account null `
 -PriceType null `
 -PriceUnit null `
 -IncurCharges null `
 -Currency USD `
 -Cost 10.5 `
 -MarkupType null `
 -Markup 2.5 `
 -MarkupPercent 13.5 `
 -CustomPrice 12.25 `
 -Platform linux `
 -Software null `
 -VolumeType null `
 -Datastore null `
 -CrossCloudApply null
```

- Convert the resource to JSON
```powershell
$AddPricesRequestPrice | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

