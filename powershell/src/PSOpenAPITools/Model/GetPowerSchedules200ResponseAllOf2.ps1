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

.PARAMETER Schedule
No description available.
.OUTPUTS

GetPowerSchedules200ResponseAllOf2<PSCustomObject>
#>

function Initialize-GetPowerSchedules200ResponseAllOf2 {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Schedule}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetPowerSchedules200ResponseAllOf2' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "schedule" = ${Schedule}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetPowerSchedules200ResponseAllOf2<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetPowerSchedules200ResponseAllOf2<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetPowerSchedules200ResponseAllOf2<PSCustomObject>
#>
function ConvertFrom-JsonToGetPowerSchedules200ResponseAllOf2 {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetPowerSchedules200ResponseAllOf2' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetPowerSchedules200ResponseAllOf2
        $AllProperties = ("schedule")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "schedule"))) { #optional property not found
            $Schedule = $null
        } else {
            $Schedule = $JsonParameters.PSobject.Properties["schedule"].value
        }

        $PSO = [PSCustomObject]@{
            "schedule" = ${Schedule}
        }

        return $PSO
    }

}
