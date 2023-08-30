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

.PARAMETER Job
No description available.
.OUTPUTS

GetJobs200Response<PSCustomObject>
#>

function Initialize-GetJobs200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Job}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetJobs200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "job" = ${Job}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetJobs200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetJobs200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetJobs200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetJobs200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetJobs200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetJobs200Response
        $AllProperties = ("job")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "job"))) { #optional property not found
            $Job = $null
        } else {
            $Job = $JsonParameters.PSobject.Properties["job"].value
        }

        $PSO = [PSCustomObject]@{
            "job" = ${Job}
        }

        return $PSO
    }

}
