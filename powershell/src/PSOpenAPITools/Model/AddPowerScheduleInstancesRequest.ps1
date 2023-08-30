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

.PARAMETER Instances
No description available.
.OUTPUTS

AddPowerScheduleInstancesRequest<PSCustomObject>
#>

function Initialize-AddPowerScheduleInstancesRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${Instances}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddPowerScheduleInstancesRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Instances) {
            throw "invalid value for 'Instances', 'Instances' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "instances" = ${Instances}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddPowerScheduleInstancesRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddPowerScheduleInstancesRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddPowerScheduleInstancesRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddPowerScheduleInstancesRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddPowerScheduleInstancesRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddPowerScheduleInstancesRequest
        $AllProperties = ("instances")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'instances' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instances"))) {
            throw "Error! JSON cannot be serialized due to the required property 'instances' missing."
        } else {
            $Instances = $JsonParameters.PSobject.Properties["instances"].value
        }

        $PSO = [PSCustomObject]@{
            "instances" = ${Instances}
        }

        return $PSO
    }

}
