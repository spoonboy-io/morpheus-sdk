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

.PARAMETER NetworkDomain
No description available.
.OUTPUTS

CreateNetworkDomainRequest<PSCustomObject>
#>

function Initialize-CreateNetworkDomainRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${NetworkDomain}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CreateNetworkDomainRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "networkDomain" = ${NetworkDomain}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CreateNetworkDomainRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to CreateNetworkDomainRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CreateNetworkDomainRequest<PSCustomObject>
#>
function ConvertFrom-JsonToCreateNetworkDomainRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CreateNetworkDomainRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CreateNetworkDomainRequest
        $AllProperties = ("networkDomain")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkDomain"))) { #optional property not found
            $NetworkDomain = $null
        } else {
            $NetworkDomain = $JsonParameters.PSobject.Properties["networkDomain"].value
        }

        $PSO = [PSCustomObject]@{
            "networkDomain" = ${NetworkDomain}
        }

        return $PSO
    }

}
