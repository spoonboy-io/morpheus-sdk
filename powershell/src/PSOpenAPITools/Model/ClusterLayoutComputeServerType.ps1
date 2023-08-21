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
.PARAMETER Code
No description available.
.PARAMETER Name
No description available.
.PARAMETER Managed
No description available.
.PARAMETER ExternalDelete
No description available.
.OUTPUTS

ClusterLayoutComputeServerType<PSCustomObject>
#>

function Initialize-ClusterLayoutComputeServerType {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Code},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Managed},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${ExternalDelete}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ClusterLayoutComputeServerType' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "code" = ${Code}
            "name" = ${Name}
            "managed" = ${Managed}
            "externalDelete" = ${ExternalDelete}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ClusterLayoutComputeServerType<PSCustomObject>

.DESCRIPTION

Convert from JSON to ClusterLayoutComputeServerType<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ClusterLayoutComputeServerType<PSCustomObject>
#>
function ConvertFrom-JsonToClusterLayoutComputeServerType {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ClusterLayoutComputeServerType' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ClusterLayoutComputeServerType
        $AllProperties = ("id", "code", "name", "managed", "externalDelete")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "code"))) { #optional property not found
            $Code = $null
        } else {
            $Code = $JsonParameters.PSobject.Properties["code"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "managed"))) { #optional property not found
            $Managed = $null
        } else {
            $Managed = $JsonParameters.PSobject.Properties["managed"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalDelete"))) { #optional property not found
            $ExternalDelete = $null
        } else {
            $ExternalDelete = $JsonParameters.PSobject.Properties["externalDelete"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "code" = ${Code}
            "name" = ${Name}
            "managed" = ${Managed}
            "externalDelete" = ${ExternalDelete}
        }

        return $PSO
    }

}
