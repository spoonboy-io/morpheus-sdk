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

.PARAMETER RuleGroup
No description available.
.OUTPUTS

CreateNetworkRouterFirewallRuleGroupRequest<PSCustomObject>
#>

function Initialize-CreateNetworkRouterFirewallRuleGroupRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${RuleGroup}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CreateNetworkRouterFirewallRuleGroupRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "ruleGroup" = ${RuleGroup}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CreateNetworkRouterFirewallRuleGroupRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to CreateNetworkRouterFirewallRuleGroupRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CreateNetworkRouterFirewallRuleGroupRequest<PSCustomObject>
#>
function ConvertFrom-JsonToCreateNetworkRouterFirewallRuleGroupRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CreateNetworkRouterFirewallRuleGroupRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CreateNetworkRouterFirewallRuleGroupRequest
        $AllProperties = ("ruleGroup")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ruleGroup"))) { #optional property not found
            $RuleGroup = $null
        } else {
            $RuleGroup = $JsonParameters.PSobject.Properties["ruleGroup"].value
        }

        $PSO = [PSCustomObject]@{
            "ruleGroup" = ${RuleGroup}
        }

        return $PSO
    }

}

