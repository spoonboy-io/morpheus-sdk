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

.PARAMETER NetworkDhcpRelay
No description available.
.OUTPUTS

CreateNetworkDhcpRelayRequest<PSCustomObject>
#>

function Initialize-CreateNetworkDhcpRelayRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${NetworkDhcpRelay}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CreateNetworkDhcpRelayRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "networkDhcpRelay" = ${NetworkDhcpRelay}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CreateNetworkDhcpRelayRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to CreateNetworkDhcpRelayRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CreateNetworkDhcpRelayRequest<PSCustomObject>
#>
function ConvertFrom-JsonToCreateNetworkDhcpRelayRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CreateNetworkDhcpRelayRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CreateNetworkDhcpRelayRequest
        $AllProperties = ("networkDhcpRelay")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkDhcpRelay"))) { #optional property not found
            $NetworkDhcpRelay = $null
        } else {
            $NetworkDhcpRelay = $JsonParameters.PSobject.Properties["networkDhcpRelay"].value
        }

        $PSO = [PSCustomObject]@{
            "networkDhcpRelay" = ${NetworkDhcpRelay}
        }

        return $PSO
    }

}

