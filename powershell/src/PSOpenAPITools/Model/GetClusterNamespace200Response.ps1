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

.PARAMETER Namespace
No description available.
.OUTPUTS

GetClusterNamespace200Response<PSCustomObject>
#>

function Initialize-GetClusterNamespace200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Namespace}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetClusterNamespace200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "namespace" = ${Namespace}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetClusterNamespace200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetClusterNamespace200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetClusterNamespace200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetClusterNamespace200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetClusterNamespace200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetClusterNamespace200Response
        $AllProperties = ("namespace")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "namespace"))) { #optional property not found
            $Namespace = $null
        } else {
            $Namespace = $JsonParameters.PSobject.Properties["namespace"].value
        }

        $PSO = [PSCustomObject]@{
            "namespace" = ${Namespace}
        }

        return $PSO
    }

}

