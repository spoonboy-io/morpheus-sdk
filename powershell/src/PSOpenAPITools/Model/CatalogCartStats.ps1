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
.PARAMETER Currency
No description available.
.PARAMETER Unit
No description available.
.OUTPUTS

CatalogCartStats<PSCustomObject>
#>

function Initialize-CatalogCartStats {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${Price},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Currency},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Unit}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CatalogCartStats' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "price" = ${Price}
            "currency" = ${Currency}
            "unit" = ${Unit}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CatalogCartStats<PSCustomObject>

.DESCRIPTION

Convert from JSON to CatalogCartStats<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CatalogCartStats<PSCustomObject>
#>
function ConvertFrom-JsonToCatalogCartStats {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CatalogCartStats' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CatalogCartStats
        $AllProperties = ("price", "currency", "unit")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "currency"))) { #optional property not found
            $Currency = $null
        } else {
            $Currency = $JsonParameters.PSobject.Properties["currency"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "unit"))) { #optional property not found
            $Unit = $null
        } else {
            $Unit = $JsonParameters.PSobject.Properties["unit"].value
        }

        $PSO = [PSCustomObject]@{
            "price" = ${Price}
            "currency" = ${Currency}
            "unit" = ${Unit}
        }

        return $PSO
    }

}

