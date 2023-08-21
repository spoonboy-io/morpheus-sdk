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

Object containing workflow configuration parameters

.PARAMETER CustomOptions
Object containing any custom option type configuration parameters
.OUTPUTS

ApiServersIdWorkflowTaskSet<PSCustomObject>
#>

function Initialize-ApiServersIdWorkflowTaskSet {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${CustomOptions}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiServersIdWorkflowTaskSet' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "customOptions" = ${CustomOptions}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiServersIdWorkflowTaskSet<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiServersIdWorkflowTaskSet<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiServersIdWorkflowTaskSet<PSCustomObject>
#>
function ConvertFrom-JsonToApiServersIdWorkflowTaskSet {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiServersIdWorkflowTaskSet' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiServersIdWorkflowTaskSet
        $AllProperties = ("customOptions")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "customOptions"))) { #optional property not found
            $CustomOptions = $null
        } else {
            $CustomOptions = $JsonParameters.PSobject.Properties["customOptions"].value
        }

        $PSO = [PSCustomObject]@{
            "customOptions" = ${CustomOptions}
        }

        return $PSO
    }

}
