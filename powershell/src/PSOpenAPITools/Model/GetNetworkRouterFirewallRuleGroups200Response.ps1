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

.PARAMETER RuleGroups
No description available.
.OUTPUTS

GetNetworkRouterFirewallRuleGroups200Response<PSCustomObject>
#>

function Initialize-GetNetworkRouterFirewallRuleGroups200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${RuleGroups}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetNetworkRouterFirewallRuleGroups200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "ruleGroups" = ${RuleGroups}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetNetworkRouterFirewallRuleGroups200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetNetworkRouterFirewallRuleGroups200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetNetworkRouterFirewallRuleGroups200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetNetworkRouterFirewallRuleGroups200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetNetworkRouterFirewallRuleGroups200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetNetworkRouterFirewallRuleGroups200Response
        $AllProperties = ("ruleGroups")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ruleGroups"))) { #optional property not found
            $RuleGroups = $null
        } else {
            $RuleGroups = $JsonParameters.PSobject.Properties["ruleGroups"].value
        }

        $PSO = [PSCustomObject]@{
            "ruleGroups" = ${RuleGroups}
        }

        return $PSO
    }

}

