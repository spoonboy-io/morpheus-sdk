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

.PARAMETER Id
No description available.
.PARAMETER Name
No description available.
.PARAMETER Visibility
No description available.
.OUTPUTS

ClusterPermissionsResourcePool<PSCustomObject>
#>

function Initialize-ClusterPermissionsResourcePool {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Visibility}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ClusterPermissionsResourcePool' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "visibility" = ${Visibility}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ClusterPermissionsResourcePool<PSCustomObject>

.DESCRIPTION

Convert from JSON to ClusterPermissionsResourcePool<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ClusterPermissionsResourcePool<PSCustomObject>
#>
function ConvertFrom-JsonToClusterPermissionsResourcePool {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ClusterPermissionsResourcePool' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ClusterPermissionsResourcePool
        $AllProperties = ("id", "name", "visibility")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "id"))) { #optional property not found
            $Id = $null
        } else {
            $Id = $JsonParameters.PSobject.Properties["id"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "visibility"))) { #optional property not found
            $Visibility = $null
        } else {
            $Visibility = $JsonParameters.PSobject.Properties["visibility"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "visibility" = ${Visibility}
        }

        return $PSO
    }

}

