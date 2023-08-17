#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.1.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

<#
.SYNOPSIS

No summary available.

.DESCRIPTION

No description available.

.PARAMETER InstanceId
The ID of the instance being removed
.OUTPUTS

RemoveAppInstanceRequest<PSCustomObject>
#>

function Initialize-RemoveAppInstanceRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [Int64]
        ${InstanceId}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => RemoveAppInstanceRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $InstanceId) {
            throw "invalid value for 'InstanceId', 'InstanceId' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "instanceId" = ${InstanceId}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to RemoveAppInstanceRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to RemoveAppInstanceRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

RemoveAppInstanceRequest<PSCustomObject>
#>
function ConvertFrom-JsonToRemoveAppInstanceRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => RemoveAppInstanceRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in RemoveAppInstanceRequest
        $AllProperties = ("instanceId")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'instanceId' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceId"))) {
            throw "Error! JSON cannot be serialized due to the required property 'instanceId' missing."
        } else {
            $InstanceId = $JsonParameters.PSobject.Properties["instanceId"].value
        }

        $PSO = [PSCustomObject]@{
            "instanceId" = ${InstanceId}
        }

        return $PSO
    }

}
