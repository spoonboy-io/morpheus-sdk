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
.PARAMETER Tenants
No description available.
.OUTPUTS

ClusterDatastoresPermissions<PSCustomObject>
#>

function Initialize-ClusterDatastoresPermissions {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ResourcePermissions},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Tenants}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ClusterDatastoresPermissions' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "resourcePermissions" = ${ResourcePermissions}
            "tenants" = ${Tenants}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ClusterDatastoresPermissions<PSCustomObject>

.DESCRIPTION

Convert from JSON to ClusterDatastoresPermissions<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ClusterDatastoresPermissions<PSCustomObject>
#>
function ConvertFrom-JsonToClusterDatastoresPermissions {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ClusterDatastoresPermissions' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ClusterDatastoresPermissions
        $AllProperties = ("resourcePermissions", "tenants")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tenants"))) { #optional property not found
            $Tenants = $null
        } else {
            $Tenants = $JsonParameters.PSobject.Properties["tenants"].value
        }

        $PSO = [PSCustomObject]@{
            "resourcePermissions" = ${ResourcePermissions}
            "tenants" = ${Tenants}
        }

        return $PSO
    }

}
