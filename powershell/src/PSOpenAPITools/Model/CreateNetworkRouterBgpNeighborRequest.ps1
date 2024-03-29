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

.PARAMETER NetworkRouterBgpNeighbor
For a full list of available options, see bgpOptionTypes in the specific Network Router Type
.OUTPUTS

CreateNetworkRouterBgpNeighborRequest<PSCustomObject>
#>

function Initialize-CreateNetworkRouterBgpNeighborRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${NetworkRouterBgpNeighbor}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CreateNetworkRouterBgpNeighborRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "networkRouterBgpNeighbor" = ${NetworkRouterBgpNeighbor}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CreateNetworkRouterBgpNeighborRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to CreateNetworkRouterBgpNeighborRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CreateNetworkRouterBgpNeighborRequest<PSCustomObject>
#>
function ConvertFrom-JsonToCreateNetworkRouterBgpNeighborRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CreateNetworkRouterBgpNeighborRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CreateNetworkRouterBgpNeighborRequest
        $AllProperties = ("networkRouterBgpNeighbor")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkRouterBgpNeighbor"))) { #optional property not found
            $NetworkRouterBgpNeighbor = $null
        } else {
            $NetworkRouterBgpNeighbor = $JsonParameters.PSobject.Properties["networkRouterBgpNeighbor"].value
        }

        $PSO = [PSCustomObject]@{
            "networkRouterBgpNeighbor" = ${NetworkRouterBgpNeighbor}
        }

        return $PSO
    }

}

