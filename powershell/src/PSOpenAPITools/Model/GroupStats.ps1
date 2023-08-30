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

.PARAMETER InstanceCounts
No description available.
.PARAMETER ServerCounts
No description available.
.OUTPUTS

GroupStats<PSCustomObject>
#>

function Initialize-GroupStats {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${InstanceCounts},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ServerCounts}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GroupStats' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "instanceCounts" = ${InstanceCounts}
            "serverCounts" = ${ServerCounts}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GroupStats<PSCustomObject>

.DESCRIPTION

Convert from JSON to GroupStats<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GroupStats<PSCustomObject>
#>
function ConvertFrom-JsonToGroupStats {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GroupStats' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GroupStats
        $AllProperties = ("instanceCounts", "serverCounts")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceCounts"))) { #optional property not found
            $InstanceCounts = $null
        } else {
            $InstanceCounts = $JsonParameters.PSobject.Properties["instanceCounts"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serverCounts"))) { #optional property not found
            $ServerCounts = $null
        } else {
            $ServerCounts = $JsonParameters.PSobject.Properties["serverCounts"].value
        }

        $PSO = [PSCustomObject]@{
            "instanceCounts" = ${InstanceCounts}
            "serverCounts" = ${ServerCounts}
        }

        return $PSO
    }

}
