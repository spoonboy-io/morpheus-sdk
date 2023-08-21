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

.PARAMETER Id
No description available.
.PARAMETER Name
No description available.
.PARAMETER Code
No description available.
.PARAMETER PriceUnit
No description available.
.OUTPUTS

GuidanceVmwareSizingPlanBeforeActionPriceSets<PSCustomObject>
#>

function Initialize-GuidanceVmwareSizingPlanBeforeActionPriceSets {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Code},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${PriceUnit}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GuidanceVmwareSizingPlanBeforeActionPriceSets' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "code" = ${Code}
            "priceUnit" = ${PriceUnit}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GuidanceVmwareSizingPlanBeforeActionPriceSets<PSCustomObject>

.DESCRIPTION

Convert from JSON to GuidanceVmwareSizingPlanBeforeActionPriceSets<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GuidanceVmwareSizingPlanBeforeActionPriceSets<PSCustomObject>
#>
function ConvertFrom-JsonToGuidanceVmwareSizingPlanBeforeActionPriceSets {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GuidanceVmwareSizingPlanBeforeActionPriceSets' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GuidanceVmwareSizingPlanBeforeActionPriceSets
        $AllProperties = ("id", "name", "code", "priceUnit")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "id"))) { #optional property not found
            $Id = $null
        } else {
            $Id = $JsonParameters.PSobject.Properties["id"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "code"))) { #optional property not found
            $Code = $null
        } else {
            $Code = $JsonParameters.PSobject.Properties["code"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "priceUnit"))) { #optional property not found
            $PriceUnit = $null
        } else {
            $PriceUnit = $JsonParameters.PSobject.Properties["priceUnit"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "code" = ${Code}
            "priceUnit" = ${PriceUnit}
        }

        return $PSO
    }

}

