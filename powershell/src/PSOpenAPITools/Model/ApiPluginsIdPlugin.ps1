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

.PARAMETER Enabled
Can be used to disable the plugin
.PARAMETER Config
Configuration object that contains settings for the applicable option types.
.OUTPUTS

ApiPluginsIdPlugin<PSCustomObject>
#>

function Initialize-ApiPluginsIdPlugin {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Config}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiPluginsIdPlugin' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "enabled" = ${Enabled}
            "config" = ${Config}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiPluginsIdPlugin<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiPluginsIdPlugin<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiPluginsIdPlugin<PSCustomObject>
#>
function ConvertFrom-JsonToApiPluginsIdPlugin {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiPluginsIdPlugin' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiPluginsIdPlugin
        $AllProperties = ("enabled", "config")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "config"))) { #optional property not found
            $Config = $null
        } else {
            $Config = $JsonParameters.PSobject.Properties["config"].value
        }

        $PSO = [PSCustomObject]@{
            "enabled" = ${Enabled}
            "config" = ${Config}
        }

        return $PSO
    }

}
