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

.PARAMETER Object
No description available.
.OUTPUTS

GetIntegrationObjects200Response<PSCustomObject>
#>

function Initialize-GetIntegrationObjects200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Object}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetIntegrationObjects200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "object" = ${Object}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetIntegrationObjects200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetIntegrationObjects200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetIntegrationObjects200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetIntegrationObjects200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetIntegrationObjects200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetIntegrationObjects200Response
        $AllProperties = ("object")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "object"))) { #optional property not found
            $Object = $null
        } else {
            $Object = $JsonParameters.PSobject.Properties["object"].value
        }

        $PSO = [PSCustomObject]@{
            "object" = ${Object}
        }

        return $PSO
    }

}

