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

.PARAMETER Instances
No description available.
.OUTPUTS

GetPowerSchedules200ResponseAllOf<PSCustomObject>
#>

function Initialize-GetPowerSchedules200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Instances}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetPowerSchedules200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "instances" = ${Instances}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetPowerSchedules200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetPowerSchedules200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetPowerSchedules200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToGetPowerSchedules200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetPowerSchedules200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetPowerSchedules200ResponseAllOf
        $AllProperties = ("instances")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instances"))) { #optional property not found
            $Instances = $null
        } else {
            $Instances = $JsonParameters.PSobject.Properties["instances"].value
        }

        $PSO = [PSCustomObject]@{
            "instances" = ${Instances}
        }

        return $PSO
    }

}

