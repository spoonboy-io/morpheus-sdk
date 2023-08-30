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

.PARAMETER NetworkPoolServer
No description available.
.OUTPUTS

CreateNetworkPoolServer200ResponseAllOf<PSCustomObject>
#>

function Initialize-CreateNetworkPoolServer200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${NetworkPoolServer}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CreateNetworkPoolServer200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "networkPoolServer" = ${NetworkPoolServer}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CreateNetworkPoolServer200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to CreateNetworkPoolServer200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CreateNetworkPoolServer200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToCreateNetworkPoolServer200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CreateNetworkPoolServer200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CreateNetworkPoolServer200ResponseAllOf
        $AllProperties = ("networkPoolServer")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkPoolServer"))) { #optional property not found
            $NetworkPoolServer = $null
        } else {
            $NetworkPoolServer = $JsonParameters.PSobject.Properties["networkPoolServer"].value
        }

        $PSO = [PSCustomObject]@{
            "networkPoolServer" = ${NetworkPoolServer}
        }

        return $PSO
    }

}

