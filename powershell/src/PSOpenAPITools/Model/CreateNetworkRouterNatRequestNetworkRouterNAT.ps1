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

For a full list of available NAT options, see natOptionTypes in the specific Network Router Type

.PARAMETER Name
NAT name
.OUTPUTS

CreateNetworkRouterNatRequestNetworkRouterNAT<PSCustomObject>
#>

function Initialize-CreateNetworkRouterNatRequestNetworkRouterNAT {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Name}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CreateNetworkRouterNatRequestNetworkRouterNAT' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CreateNetworkRouterNatRequestNetworkRouterNAT<PSCustomObject>

.DESCRIPTION

Convert from JSON to CreateNetworkRouterNatRequestNetworkRouterNAT<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CreateNetworkRouterNatRequestNetworkRouterNAT<PSCustomObject>
#>
function ConvertFrom-JsonToCreateNetworkRouterNatRequestNetworkRouterNAT {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CreateNetworkRouterNatRequestNetworkRouterNAT' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CreateNetworkRouterNatRequestNetworkRouterNAT
        $AllProperties = ("name")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'name' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property 'name' missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
        }

        return $PSO
    }

}

