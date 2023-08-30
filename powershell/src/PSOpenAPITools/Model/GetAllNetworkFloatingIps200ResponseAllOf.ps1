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

.PARAMETER NetworkFloatingIps
No description available.
.OUTPUTS

GetAllNetworkFloatingIps200ResponseAllOf<PSCustomObject>
#>

function Initialize-GetAllNetworkFloatingIps200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${NetworkFloatingIps}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetAllNetworkFloatingIps200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "networkFloatingIps" = ${NetworkFloatingIps}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetAllNetworkFloatingIps200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetAllNetworkFloatingIps200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetAllNetworkFloatingIps200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToGetAllNetworkFloatingIps200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetAllNetworkFloatingIps200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetAllNetworkFloatingIps200ResponseAllOf
        $AllProperties = ("networkFloatingIps")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkFloatingIps"))) { #optional property not found
            $NetworkFloatingIps = $null
        } else {
            $NetworkFloatingIps = $JsonParameters.PSobject.Properties["networkFloatingIps"].value
        }

        $PSO = [PSCustomObject]@{
            "networkFloatingIps" = ${NetworkFloatingIps}
        }

        return $PSO
    }

}

