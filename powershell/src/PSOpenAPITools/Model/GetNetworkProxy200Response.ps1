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

.PARAMETER NetworkProxy
No description available.
.OUTPUTS

GetNetworkProxy200Response<PSCustomObject>
#>

function Initialize-GetNetworkProxy200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${NetworkProxy}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetNetworkProxy200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "networkProxy" = ${NetworkProxy}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetNetworkProxy200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetNetworkProxy200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetNetworkProxy200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetNetworkProxy200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetNetworkProxy200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetNetworkProxy200Response
        $AllProperties = ("networkProxy")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkProxy"))) { #optional property not found
            $NetworkProxy = $null
        } else {
            $NetworkProxy = $JsonParameters.PSobject.Properties["networkProxy"].value
        }

        $PSO = [PSCustomObject]@{
            "networkProxy" = ${NetworkProxy}
        }

        return $PSO
    }

}

