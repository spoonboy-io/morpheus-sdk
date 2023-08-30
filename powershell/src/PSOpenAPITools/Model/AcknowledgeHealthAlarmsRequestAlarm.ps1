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

.PARAMETER Acknowledged
Pass `true` to ackowledge an alarm, or pass `false` to unacknowledge it.
.PARAMETER Ids
Array of Alarm ID(s)to be updated.
.PARAMETER All
Pass `true` to update all alarms instead of passing ids. This will update any active alarm that is not already acknowledged. 
.OUTPUTS

AcknowledgeHealthAlarmsRequestAlarm<PSCustomObject>
#>

function Initialize-AcknowledgeHealthAlarmsRequestAlarm {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [Boolean]
        ${Acknowledged},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${Ids},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${All} = $false
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AcknowledgeHealthAlarmsRequestAlarm' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Acknowledged) {
            throw "invalid value for 'Acknowledged', 'Acknowledged' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "acknowledged" = ${Acknowledged}
            "ids" = ${Ids}
            "all" = ${All}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AcknowledgeHealthAlarmsRequestAlarm<PSCustomObject>

.DESCRIPTION

Convert from JSON to AcknowledgeHealthAlarmsRequestAlarm<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AcknowledgeHealthAlarmsRequestAlarm<PSCustomObject>
#>
function ConvertFrom-JsonToAcknowledgeHealthAlarmsRequestAlarm {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AcknowledgeHealthAlarmsRequestAlarm' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AcknowledgeHealthAlarmsRequestAlarm
        $AllProperties = ("acknowledged", "ids", "all")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'acknowledged' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "acknowledged"))) {
            throw "Error! JSON cannot be serialized due to the required property 'acknowledged' missing."
        } else {
            $Acknowledged = $JsonParameters.PSobject.Properties["acknowledged"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ids"))) { #optional property not found
            $Ids = $null
        } else {
            $Ids = $JsonParameters.PSobject.Properties["ids"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "all"))) { #optional property not found
            $All = $null
        } else {
            $All = $JsonParameters.PSobject.Properties["all"].value
        }

        $PSO = [PSCustomObject]@{
            "acknowledged" = ${Acknowledged}
            "ids" = ${Ids}
            "all" = ${All}
        }

        return $PSO
    }

}

