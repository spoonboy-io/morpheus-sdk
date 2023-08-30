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

.PARAMETER Rule
No description available.
.OUTPUTS

GetNetworkRouterFirewallRule200Response<PSCustomObject>
#>

function Initialize-GetNetworkRouterFirewallRule200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Rule}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetNetworkRouterFirewallRule200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "rule" = ${Rule}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetNetworkRouterFirewallRule200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetNetworkRouterFirewallRule200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetNetworkRouterFirewallRule200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetNetworkRouterFirewallRule200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetNetworkRouterFirewallRule200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetNetworkRouterFirewallRule200Response
        $AllProperties = ("rule")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "rule"))) { #optional property not found
            $Rule = $null
        } else {
            $Rule = $JsonParameters.PSobject.Properties["rule"].value
        }

        $PSO = [PSCustomObject]@{
            "rule" = ${Rule}
        }

        return $PSO
    }

}

