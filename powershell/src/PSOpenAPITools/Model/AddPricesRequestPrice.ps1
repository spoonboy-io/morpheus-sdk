#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

<#
.SYNOPSIS

No summary available.

.DESCRIPTION

No description available.

.PARAMETER Name
Price name
.PARAMETER Code
Price code, must be unique
.PARAMETER Account
No description available.
.PARAMETER PriceType
Restricts query to only load only prices with specified priceType. * `fixed` - Everything * `compute` - Memory + CPU * `memory` - Memory * `cores` - Cores * `storage` - Storage * `datastore` - Datastore * `platform` - Platform * `software` - Software * `load_balancer` - Load Balancer * `load_balancer_virtual_server` - Load Balancer Virtual Server 
.PARAMETER PriceUnit
The unit of pricing
.PARAMETER IncurCharges
Indicates when to incur charge
.PARAMETER Currency
ISO Currency code
.PARAMETER Cost
Cost
.PARAMETER MarkupType
Price adjustment type
.PARAMETER Markup
Amount for `fixed` price adjustment type
.PARAMETER MarkupPercent
Percent for `percent` price adjustment type
.PARAMETER CustomPrice
Custom price for `custom` price adjustment type
.PARAMETER Platform
Platform.  Required for `platform` price type
.PARAMETER Software
Software.  Required for software price type
.PARAMETER VolumeType
No description available.
.PARAMETER Datastore
No description available.
.PARAMETER CrossCloudApply
Apply price across clouds, optional true/false flag for datastore price type
.OUTPUTS

AddPricesRequestPrice<PSCustomObject>
#>

function Initialize-AddPricesRequestPrice {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Code},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Account},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("fixed", "compute", "memory", "cores", "storage", "datastore", "platform", "software", "load_balancer", "load_balancer_virtual_server")]
        [String]
        ${PriceType},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("minute", "hour", "day", "month", "year", "two year", "three year", "four year", "five year")]
        [String]
        ${PriceUnit},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("running", "stopped", "always")]
        [String]
        ${IncurCharges},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Currency},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [Decimal]
        ${Cost},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("fixed", "percent", "custom")]
        [String]
        ${MarkupType},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Markup},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${MarkupPercent},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CustomPrice},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Platform},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Software},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${VolumeType},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Datastore},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${CrossCloudApply}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddPricesRequestPrice' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }

        if ($null -eq $Code) {
            throw "invalid value for 'Code', 'Code' cannot be null."
        }

        if ($null -eq $PriceType) {
            throw "invalid value for 'PriceType', 'PriceType' cannot be null."
        }

        if ($null -eq $PriceUnit) {
            throw "invalid value for 'PriceUnit', 'PriceUnit' cannot be null."
        }

        if ($null -eq $IncurCharges) {
            throw "invalid value for 'IncurCharges', 'IncurCharges' cannot be null."
        }

        if ($null -eq $Currency) {
            throw "invalid value for 'Currency', 'Currency' cannot be null."
        }

        if ($null -eq $Cost) {
            throw "invalid value for 'Cost', 'Cost' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "code" = ${Code}
            "account" = ${Account}
            "priceType" = ${PriceType}
            "priceUnit" = ${PriceUnit}
            "incurCharges" = ${IncurCharges}
            "currency" = ${Currency}
            "cost" = ${Cost}
            "markupType" = ${MarkupType}
            "markup" = ${Markup}
            "markupPercent" = ${MarkupPercent}
            "customPrice" = ${CustomPrice}
            "platform" = ${Platform}
            "software" = ${Software}
            "volumeType" = ${VolumeType}
            "datastore" = ${Datastore}
            "crossCloudApply" = ${CrossCloudApply}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddPricesRequestPrice<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddPricesRequestPrice<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddPricesRequestPrice<PSCustomObject>
#>
function ConvertFrom-JsonToAddPricesRequestPrice {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddPricesRequestPrice' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddPricesRequestPrice
        $AllProperties = ("name", "code", "account", "priceType", "priceUnit", "incurCharges", "currency", "cost", "markupType", "markup", "markupPercent", "customPrice", "platform", "software", "volumeType", "datastore", "crossCloudApply")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'name' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property 'name' missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "code"))) {
            throw "Error! JSON cannot be serialized due to the required property 'code' missing."
        } else {
            $Code = $JsonParameters.PSobject.Properties["code"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "priceType"))) {
            throw "Error! JSON cannot be serialized due to the required property 'priceType' missing."
        } else {
            $PriceType = $JsonParameters.PSobject.Properties["priceType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "priceUnit"))) {
            throw "Error! JSON cannot be serialized due to the required property 'priceUnit' missing."
        } else {
            $PriceUnit = $JsonParameters.PSobject.Properties["priceUnit"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "incurCharges"))) {
            throw "Error! JSON cannot be serialized due to the required property 'incurCharges' missing."
        } else {
            $IncurCharges = $JsonParameters.PSobject.Properties["incurCharges"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "currency"))) {
            throw "Error! JSON cannot be serialized due to the required property 'currency' missing."
        } else {
            $Currency = $JsonParameters.PSobject.Properties["currency"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cost"))) {
            throw "Error! JSON cannot be serialized due to the required property 'cost' missing."
        } else {
            $Cost = $JsonParameters.PSobject.Properties["cost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "account"))) { #optional property not found
            $Account = $null
        } else {
            $Account = $JsonParameters.PSobject.Properties["account"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "markupType"))) { #optional property not found
            $MarkupType = $null
        } else {
            $MarkupType = $JsonParameters.PSobject.Properties["markupType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "markup"))) { #optional property not found
            $Markup = $null
        } else {
            $Markup = $JsonParameters.PSobject.Properties["markup"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "markupPercent"))) { #optional property not found
            $MarkupPercent = $null
        } else {
            $MarkupPercent = $JsonParameters.PSobject.Properties["markupPercent"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "customPrice"))) { #optional property not found
            $CustomPrice = $null
        } else {
            $CustomPrice = $JsonParameters.PSobject.Properties["customPrice"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "platform"))) { #optional property not found
            $Platform = $null
        } else {
            $Platform = $JsonParameters.PSobject.Properties["platform"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "software"))) { #optional property not found
            $Software = $null
        } else {
            $Software = $JsonParameters.PSobject.Properties["software"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "volumeType"))) { #optional property not found
            $VolumeType = $null
        } else {
            $VolumeType = $JsonParameters.PSobject.Properties["volumeType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "datastore"))) { #optional property not found
            $Datastore = $null
        } else {
            $Datastore = $JsonParameters.PSobject.Properties["datastore"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "crossCloudApply"))) { #optional property not found
            $CrossCloudApply = $null
        } else {
            $CrossCloudApply = $JsonParameters.PSobject.Properties["crossCloudApply"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "code" = ${Code}
            "account" = ${Account}
            "priceType" = ${PriceType}
            "priceUnit" = ${PriceUnit}
            "incurCharges" = ${IncurCharges}
            "currency" = ${Currency}
            "cost" = ${Cost}
            "markupType" = ${MarkupType}
            "markup" = ${Markup}
            "markupPercent" = ${MarkupPercent}
            "customPrice" = ${CustomPrice}
            "platform" = ${Platform}
            "software" = ${Software}
            "volumeType" = ${VolumeType}
            "datastore" = ${Datastore}
            "crossCloudApply" = ${CrossCloudApply}
        }

        return $PSO
    }

}

