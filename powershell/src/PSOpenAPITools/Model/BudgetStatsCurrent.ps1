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

.PARAMETER EstimatedCost
No description available.
.PARAMETER LastCost
No description available.
.OUTPUTS

BudgetStatsCurrent<PSCustomObject>
#>

function Initialize-BudgetStatsCurrent {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${EstimatedCost},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LastCost}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BudgetStatsCurrent' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "estimatedCost" = ${EstimatedCost}
            "lastCost" = ${LastCost}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BudgetStatsCurrent<PSCustomObject>

.DESCRIPTION

Convert from JSON to BudgetStatsCurrent<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BudgetStatsCurrent<PSCustomObject>
#>
function ConvertFrom-JsonToBudgetStatsCurrent {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BudgetStatsCurrent' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BudgetStatsCurrent
        $AllProperties = ("estimatedCost", "lastCost")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "estimatedCost"))) { #optional property not found
            $EstimatedCost = $null
        } else {
            $EstimatedCost = $JsonParameters.PSobject.Properties["estimatedCost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastCost"))) { #optional property not found
            $LastCost = $null
        } else {
            $LastCost = $JsonParameters.PSobject.Properties["lastCost"].value
        }

        $PSO = [PSCustomObject]@{
            "estimatedCost" = ${EstimatedCost}
            "lastCost" = ${LastCost}
        }

        return $PSO
    }

}
