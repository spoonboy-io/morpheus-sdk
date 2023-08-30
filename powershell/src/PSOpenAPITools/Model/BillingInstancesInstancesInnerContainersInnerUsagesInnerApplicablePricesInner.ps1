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

.PARAMETER StartDate
No description available.
.PARAMETER EndDate
No description available.
.PARAMETER NumUnits
No description available.
.PARAMETER Cost
No description available.
.PARAMETER Price
No description available.
.PARAMETER Currency
No description available.
.PARAMETER Prices
No description available.
.OUTPUTS

BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner<PSCustomObject>
#>

function Initialize-BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${StartDate},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${EndDate},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${NumUnits},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Cost},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Price},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Currency},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Prices}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "numUnits" = ${NumUnits}
            "cost" = ${Cost}
            "price" = ${Price}
            "currency" = ${Currency}
            "prices" = ${Prices}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner<PSCustomObject>
#>
function ConvertFrom-JsonToBillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner
        $AllProperties = ("startDate", "endDate", "numUnits", "cost", "price", "currency", "prices")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "startDate"))) { #optional property not found
            $StartDate = $null
        } else {
            $StartDate = $JsonParameters.PSobject.Properties["startDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "endDate"))) { #optional property not found
            $EndDate = $null
        } else {
            $EndDate = $JsonParameters.PSobject.Properties["endDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "numUnits"))) { #optional property not found
            $NumUnits = $null
        } else {
            $NumUnits = $JsonParameters.PSobject.Properties["numUnits"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "currency"))) { #optional property not found
            $Currency = $null
        } else {
            $Currency = $JsonParameters.PSobject.Properties["currency"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "prices"))) { #optional property not found
            $Prices = $null
        } else {
            $Prices = $JsonParameters.PSobject.Properties["prices"].value
        }

        $PSO = [PSCustomObject]@{
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "numUnits" = ${NumUnits}
            "cost" = ${Cost}
            "price" = ${Price}
            "currency" = ${Currency}
            "prices" = ${Prices}
        }

        return $PSO
    }

}

