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

.PARAMETER SyslogRule
No description available.
.OUTPUTS

AddLogSettingsSyslogRulesRequest<PSCustomObject>
#>

function Initialize-AddLogSettingsSyslogRulesRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${SyslogRule}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddLogSettingsSyslogRulesRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $SyslogRule) {
            throw "invalid value for 'SyslogRule', 'SyslogRule' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "syslogRule" = ${SyslogRule}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddLogSettingsSyslogRulesRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddLogSettingsSyslogRulesRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddLogSettingsSyslogRulesRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddLogSettingsSyslogRulesRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddLogSettingsSyslogRulesRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddLogSettingsSyslogRulesRequest
        $AllProperties = ("syslogRule")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'syslogRule' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "syslogRule"))) {
            throw "Error! JSON cannot be serialized due to the required property 'syslogRule' missing."
        } else {
            $SyslogRule = $JsonParameters.PSobject.Properties["syslogRule"].value
        }

        $PSO = [PSCustomObject]@{
            "syslogRule" = ${SyslogRule}
        }

        return $PSO
    }

}
