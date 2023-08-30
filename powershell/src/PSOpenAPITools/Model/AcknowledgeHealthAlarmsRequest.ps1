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

.PARAMETER Alarm
No description available.
.OUTPUTS

AcknowledgeHealthAlarmsRequest<PSCustomObject>
#>

function Initialize-AcknowledgeHealthAlarmsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Alarm}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AcknowledgeHealthAlarmsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Alarm) {
            throw "invalid value for 'Alarm', 'Alarm' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "alarm" = ${Alarm}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AcknowledgeHealthAlarmsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AcknowledgeHealthAlarmsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AcknowledgeHealthAlarmsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAcknowledgeHealthAlarmsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AcknowledgeHealthAlarmsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AcknowledgeHealthAlarmsRequest
        $AllProperties = ("alarm")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'alarm' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "alarm"))) {
            throw "Error! JSON cannot be serialized due to the required property 'alarm' missing."
        } else {
            $Alarm = $JsonParameters.PSobject.Properties["alarm"].value
        }

        $PSO = [PSCustomObject]@{
            "alarm" = ${Alarm}
        }

        return $PSO
    }

}

