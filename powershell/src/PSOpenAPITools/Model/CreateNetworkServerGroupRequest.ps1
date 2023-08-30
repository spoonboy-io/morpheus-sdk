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

.PARAMETER Group
No description available.
.OUTPUTS

CreateNetworkServerGroupRequest<PSCustomObject>
#>

function Initialize-CreateNetworkServerGroupRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Group}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CreateNetworkServerGroupRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "group" = ${Group}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CreateNetworkServerGroupRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to CreateNetworkServerGroupRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CreateNetworkServerGroupRequest<PSCustomObject>
#>
function ConvertFrom-JsonToCreateNetworkServerGroupRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CreateNetworkServerGroupRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CreateNetworkServerGroupRequest
        $AllProperties = ("group")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "group"))) { #optional property not found
            $Group = $null
        } else {
            $Group = $JsonParameters.PSobject.Properties["group"].value
        }

        $PSO = [PSCustomObject]@{
            "group" = ${Group}
        }

        return $PSO
    }

}

