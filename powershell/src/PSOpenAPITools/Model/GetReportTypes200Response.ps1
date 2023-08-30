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

.PARAMETER ReportTypes
No description available.
.OUTPUTS

GetReportTypes200Response<PSCustomObject>
#>

function Initialize-GetReportTypes200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${ReportTypes}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetReportTypes200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "reportTypes" = ${ReportTypes}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetReportTypes200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetReportTypes200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetReportTypes200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetReportTypes200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetReportTypes200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetReportTypes200Response
        $AllProperties = ("reportTypes")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "reportTypes"))) { #optional property not found
            $ReportTypes = $null
        } else {
            $ReportTypes = $JsonParameters.PSobject.Properties["reportTypes"].value
        }

        $PSO = [PSCustomObject]@{
            "reportTypes" = ${ReportTypes}
        }

        return $PSO
    }

}
