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

.PARAMETER Script
A script or command to be executed
.OUTPUTS

ExecuteExecutionRequestRequest<PSCustomObject>
#>

function Initialize-ExecuteExecutionRequestRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Script}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ExecuteExecutionRequestRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Script) {
            throw "invalid value for 'Script', 'Script' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "script" = ${Script}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ExecuteExecutionRequestRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to ExecuteExecutionRequestRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ExecuteExecutionRequestRequest<PSCustomObject>
#>
function ConvertFrom-JsonToExecuteExecutionRequestRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ExecuteExecutionRequestRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ExecuteExecutionRequestRequest
        $AllProperties = ("script")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'script' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "script"))) {
            throw "Error! JSON cannot be serialized due to the required property 'script' missing."
        } else {
            $Script = $JsonParameters.PSobject.Properties["script"].value
        }

        $PSO = [PSCustomObject]@{
            "script" = ${Script}
        }

        return $PSO
    }

}
