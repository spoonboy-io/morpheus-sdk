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

.PARAMETER NetworkRouter
No description available.
.OUTPUTS

GetNetworkRouter200Response<PSCustomObject>
#>

function Initialize-GetNetworkRouter200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${NetworkRouter}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetNetworkRouter200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "networkRouter" = ${NetworkRouter}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetNetworkRouter200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetNetworkRouter200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetNetworkRouter200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetNetworkRouter200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetNetworkRouter200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetNetworkRouter200Response
        $AllProperties = ("networkRouter")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkRouter"))) { #optional property not found
            $NetworkRouter = $null
        } else {
            $NetworkRouter = $JsonParameters.PSobject.Properties["networkRouter"].value
        }

        $PSO = [PSCustomObject]@{
            "networkRouter" = ${NetworkRouter}
        }

        return $PSO
    }

}

