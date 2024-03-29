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

.PARAMETER ServerType
No description available.
.OUTPUTS

GetHostType200Response<PSCustomObject>
#>

function Initialize-GetHostType200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ServerType}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetHostType200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "serverType" = ${ServerType}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetHostType200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetHostType200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetHostType200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetHostType200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetHostType200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetHostType200Response
        $AllProperties = ("serverType")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serverType"))) { #optional property not found
            $ServerType = $null
        } else {
            $ServerType = $JsonParameters.PSobject.Properties["serverType"].value
        }

        $PSO = [PSCustomObject]@{
            "serverType" = ${ServerType}
        }

        return $PSO
    }

}

