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

.PARAMETER NetworkRoute
No description available.
.OUTPUTS

GetNetworkRouterRoute200Response<PSCustomObject>
#>

function Initialize-GetNetworkRouterRoute200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${NetworkRoute}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetNetworkRouterRoute200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "networkRoute" = ${NetworkRoute}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetNetworkRouterRoute200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetNetworkRouterRoute200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetNetworkRouterRoute200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetNetworkRouterRoute200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetNetworkRouterRoute200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetNetworkRouterRoute200Response
        $AllProperties = ("networkRoute")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkRoute"))) { #optional property not found
            $NetworkRoute = $null
        } else {
            $NetworkRoute = $JsonParameters.PSobject.Properties["networkRoute"].value
        }

        $PSO = [PSCustomObject]@{
            "networkRoute" = ${NetworkRoute}
        }

        return $PSO
    }

}

