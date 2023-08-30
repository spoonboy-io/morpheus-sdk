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

.PARAMETER Deployment
No description available.
.OUTPUTS

GetDeployment200Response<PSCustomObject>
#>

function Initialize-GetDeployment200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Deployment}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetDeployment200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "deployment" = ${Deployment}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetDeployment200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetDeployment200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetDeployment200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetDeployment200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetDeployment200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetDeployment200Response
        $AllProperties = ("deployment")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "deployment"))) { #optional property not found
            $Deployment = $null
        } else {
            $Deployment = $JsonParameters.PSobject.Properties["deployment"].value
        }

        $PSO = [PSCustomObject]@{
            "deployment" = ${Deployment}
        }

        return $PSO
    }

}

