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

.PARAMETER Price
No description available.
.PARAMETER Cost
No description available.
.PARAMETER StartDate
No description available.
.PARAMETER EndDate
No description available.
.PARAMETER Zones
No description available.
.OUTPUTS

BillingZones<PSCustomObject>
#>

function Initialize-BillingZones {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Price},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Cost},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${StartDate},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${EndDate},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Zones}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BillingZones' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "price" = ${Price}
            "cost" = ${Cost}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "zones" = ${Zones}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BillingZones<PSCustomObject>

.DESCRIPTION

Convert from JSON to BillingZones<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BillingZones<PSCustomObject>
#>
function ConvertFrom-JsonToBillingZones {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BillingZones' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BillingZones
        $AllProperties = ("price", "cost", "startDate", "endDate", "zones")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "price"))) { #optional property not found
            $Price = $null
        } else {
            $Price = $JsonParameters.PSobject.Properties["price"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cost"))) { #optional property not found
            $Cost = $null
        } else {
            $Cost = $JsonParameters.PSobject.Properties["cost"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "zones"))) { #optional property not found
            $Zones = $null
        } else {
            $Zones = $JsonParameters.PSobject.Properties["zones"].value
        }

        $PSO = [PSCustomObject]@{
            "price" = ${Price}
            "cost" = ${Cost}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "zones" = ${Zones}
        }

        return $PSO
    }

}
