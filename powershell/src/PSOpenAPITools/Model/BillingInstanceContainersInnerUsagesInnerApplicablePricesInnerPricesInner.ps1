#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.1.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

<#
.SYNOPSIS

No summary available.

.DESCRIPTION

No description available.

.PARAMETER Type
No description available.
.PARAMETER PricePerUnit
No description available.
.PARAMETER CostPerUnit
No description available.
.PARAMETER Cost
No description available.
.PARAMETER Price
No description available.
.PARAMETER Quantity
No description available.
.OUTPUTS

BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner<PSCustomObject>
#>

function Initialize-BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Type},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${PricePerUnit},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CostPerUnit},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Cost},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Price},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Quantity}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "type" = ${Type}
            "pricePerUnit" = ${PricePerUnit}
            "costPerUnit" = ${CostPerUnit}
            "cost" = ${Cost}
            "price" = ${Price}
            "quantity" = ${Quantity}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner<PSCustomObject>
#>
function ConvertFrom-JsonToBillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner
        $AllProperties = ("type", "pricePerUnit", "costPerUnit", "cost", "price", "quantity")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) { #optional property not found
            $Type = $null
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "pricePerUnit"))) { #optional property not found
            $PricePerUnit = $null
        } else {
            $PricePerUnit = $JsonParameters.PSobject.Properties["pricePerUnit"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "costPerUnit"))) { #optional property not found
            $CostPerUnit = $null
        } else {
            $CostPerUnit = $JsonParameters.PSobject.Properties["costPerUnit"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cost"))) { #optional property not found
            $Cost = $null
        } else {
            $Cost = $JsonParameters.PSobject.Properties["cost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "price"))) { #optional property not found
            $Price = $null
        } else {
            $Price = $JsonParameters.PSobject.Properties["price"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "quantity"))) { #optional property not found
            $Quantity = $null
        } else {
            $Quantity = $JsonParameters.PSobject.Properties["quantity"].value
        }

        $PSO = [PSCustomObject]@{
            "type" = ${Type}
            "pricePerUnit" = ${PricePerUnit}
            "costPerUnit" = ${CostPerUnit}
            "cost" = ${Cost}
            "price" = ${Price}
            "quantity" = ${Quantity}
        }

        return $PSO
    }

}

