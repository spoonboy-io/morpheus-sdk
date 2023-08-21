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

.PARAMETER Config
JSON encoded list of parameters that varies by instance type.
.OUTPUTS

ApiDeploysIdAppDeploy<PSCustomObject>
#>

function Initialize-ApiDeploysIdAppDeploy {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Config}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiDeploysIdAppDeploy' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "config" = ${Config}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiDeploysIdAppDeploy<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiDeploysIdAppDeploy<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiDeploysIdAppDeploy<PSCustomObject>
#>
function ConvertFrom-JsonToApiDeploysIdAppDeploy {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiDeploysIdAppDeploy' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiDeploysIdAppDeploy
        $AllProperties = ("config")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "config"))) { #optional property not found
            $Config = $null
        } else {
            $Config = $JsonParameters.PSobject.Properties["config"].value
        }

        $PSO = [PSCustomObject]@{
            "config" = ${Config}
        }

        return $PSO
    }

}

