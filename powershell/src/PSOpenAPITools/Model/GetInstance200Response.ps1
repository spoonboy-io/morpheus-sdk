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

.PARAMETER Instance
No description available.
.OUTPUTS

GetInstance200Response<PSCustomObject>
#>

function Initialize-GetInstance200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Instance}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetInstance200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "instance" = ${Instance}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetInstance200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetInstance200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetInstance200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetInstance200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetInstance200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetInstance200Response
        $AllProperties = ("instance")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instance"))) { #optional property not found
            $Instance = $null
        } else {
            $Instance = $JsonParameters.PSobject.Properties["instance"].value
        }

        $PSO = [PSCustomObject]@{
            "instance" = ${Instance}
        }

        return $PSO
    }

}
