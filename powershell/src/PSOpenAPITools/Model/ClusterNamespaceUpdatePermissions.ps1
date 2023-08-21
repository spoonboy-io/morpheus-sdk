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

.PARAMETER ResourcePermissions
No description available.
.OUTPUTS

ClusterNamespaceUpdatePermissions<PSCustomObject>
#>

function Initialize-ClusterNamespaceUpdatePermissions {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ResourcePermissions}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ClusterNamespaceUpdatePermissions' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "resourcePermissions" = ${ResourcePermissions}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ClusterNamespaceUpdatePermissions<PSCustomObject>

.DESCRIPTION

Convert from JSON to ClusterNamespaceUpdatePermissions<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ClusterNamespaceUpdatePermissions<PSCustomObject>
#>
function ConvertFrom-JsonToClusterNamespaceUpdatePermissions {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ClusterNamespaceUpdatePermissions' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ClusterNamespaceUpdatePermissions
        $AllProperties = ("resourcePermissions")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "resourcePermissions"))) { #optional property not found
            $ResourcePermissions = $null
        } else {
            $ResourcePermissions = $JsonParameters.PSobject.Properties["resourcePermissions"].value
        }

        $PSO = [PSCustomObject]@{
            "resourcePermissions" = ${ResourcePermissions}
        }

        return $PSO
    }

}
