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

AddSecurityGroupRulesRequest<PSCustomObject>
#>

function Initialize-AddSecurityGroupRulesRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Rule}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddSecurityGroupRulesRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Rule) {
            throw "invalid value for 'Rule', 'Rule' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "rule" = ${Rule}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddSecurityGroupRulesRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddSecurityGroupRulesRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddSecurityGroupRulesRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddSecurityGroupRulesRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddSecurityGroupRulesRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddSecurityGroupRulesRequest
        $AllProperties = ("rule")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'rule' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "rule"))) {
            throw "Error! JSON cannot be serialized due to the required property 'rule' missing."
        } else {
            $Rule = $JsonParameters.PSobject.Properties["rule"].value
        }

        $PSO = [PSCustomObject]@{
            "rule" = ${Rule}
        }

        return $PSO
    }

}

