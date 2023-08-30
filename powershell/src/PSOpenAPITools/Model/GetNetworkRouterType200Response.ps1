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

.PARAMETER NetworkRouterType
No description available.
.OUTPUTS

GetNetworkRouterType200Response<PSCustomObject>
#>

function Initialize-GetNetworkRouterType200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${NetworkRouterType}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetNetworkRouterType200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "networkRouterType" = ${NetworkRouterType}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetNetworkRouterType200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetNetworkRouterType200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetNetworkRouterType200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetNetworkRouterType200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetNetworkRouterType200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetNetworkRouterType200Response
        $AllProperties = ("networkRouterType")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkRouterType"))) { #optional property not found
            $NetworkRouterType = $null
        } else {
            $NetworkRouterType = $JsonParameters.PSobject.Properties["networkRouterType"].value
        }

        $PSO = [PSCustomObject]@{
            "networkRouterType" = ${NetworkRouterType}
        }

        return $PSO
    }

}
