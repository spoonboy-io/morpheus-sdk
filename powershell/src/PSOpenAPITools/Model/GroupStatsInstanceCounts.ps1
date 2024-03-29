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

.PARAMETER All
No description available.
.OUTPUTS

GroupStatsInstanceCounts<PSCustomObject>
#>

function Initialize-GroupStatsInstanceCounts {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${All}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GroupStatsInstanceCounts' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "all" = ${All}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GroupStatsInstanceCounts<PSCustomObject>

.DESCRIPTION

Convert from JSON to GroupStatsInstanceCounts<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GroupStatsInstanceCounts<PSCustomObject>
#>
function ConvertFrom-JsonToGroupStatsInstanceCounts {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GroupStatsInstanceCounts' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GroupStatsInstanceCounts
        $AllProperties = ("all")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "all"))) { #optional property not found
            $All = $null
        } else {
            $All = $JsonParameters.PSobject.Properties["all"].value
        }

        $PSO = [PSCustomObject]@{
            "all" = ${All}
        }

        return $PSO
    }

}

