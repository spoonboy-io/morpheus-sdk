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

.PARAMETER NetworkDhcpServer
No description available.
.OUTPUTS

CreateNetworkDhcpServerRequest<PSCustomObject>
#>

function Initialize-CreateNetworkDhcpServerRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${NetworkDhcpServer}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CreateNetworkDhcpServerRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "networkDhcpServer" = ${NetworkDhcpServer}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CreateNetworkDhcpServerRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to CreateNetworkDhcpServerRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CreateNetworkDhcpServerRequest<PSCustomObject>
#>
function ConvertFrom-JsonToCreateNetworkDhcpServerRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CreateNetworkDhcpServerRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CreateNetworkDhcpServerRequest
        $AllProperties = ("networkDhcpServer")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkDhcpServer"))) { #optional property not found
            $NetworkDhcpServer = $null
        } else {
            $NetworkDhcpServer = $JsonParameters.PSobject.Properties["networkDhcpServer"].value
        }

        $PSO = [PSCustomObject]@{
            "networkDhcpServer" = ${NetworkDhcpServer}
        }

        return $PSO
    }

}
