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

.PARAMETER Price
No description available.
.PARAMETER Cost
No description available.
.PARAMETER Servers
No description available.
.PARAMETER Count
No description available.
.OUTPUTS

BillingZonesInnerComputeServers<PSCustomObject>
#>

function Initialize-BillingZonesInnerComputeServers {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Price},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Cost},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Servers},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Count}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BillingZonesInnerComputeServers' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "price" = ${Price}
            "cost" = ${Cost}
            "servers" = ${Servers}
            "count" = ${Count}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BillingZonesInnerComputeServers<PSCustomObject>

.DESCRIPTION

Convert from JSON to BillingZonesInnerComputeServers<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BillingZonesInnerComputeServers<PSCustomObject>
#>
function ConvertFrom-JsonToBillingZonesInnerComputeServers {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BillingZonesInnerComputeServers' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BillingZonesInnerComputeServers
        $AllProperties = ("price", "cost", "servers", "count")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "servers"))) { #optional property not found
            $Servers = $null
        } else {
            $Servers = $JsonParameters.PSobject.Properties["servers"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "count"))) { #optional property not found
            $Count = $null
        } else {
            $Count = $JsonParameters.PSobject.Properties["count"].value
        }

        $PSO = [PSCustomObject]@{
            "price" = ${Price}
            "cost" = ${Cost}
            "servers" = ${Servers}
            "count" = ${Count}
        }

        return $PSO
    }

}
