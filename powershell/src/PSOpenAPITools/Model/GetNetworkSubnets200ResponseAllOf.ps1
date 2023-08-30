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

.PARAMETER Subnets
No description available.
.OUTPUTS

GetNetworkSubnets200ResponseAllOf<PSCustomObject>
#>

function Initialize-GetNetworkSubnets200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Subnets}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetNetworkSubnets200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "subnets" = ${Subnets}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetNetworkSubnets200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetNetworkSubnets200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetNetworkSubnets200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToGetNetworkSubnets200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetNetworkSubnets200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetNetworkSubnets200ResponseAllOf
        $AllProperties = ("subnets")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "subnets"))) { #optional property not found
            $Subnets = $null
        } else {
            $Subnets = $JsonParameters.PSobject.Properties["subnets"].value
        }

        $PSO = [PSCustomObject]@{
            "subnets" = ${Subnets}
        }

        return $PSO
    }

}
