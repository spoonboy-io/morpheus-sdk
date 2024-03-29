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

AddBackupJobsRequest<PSCustomObject>
#>

function Initialize-AddBackupJobsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Job}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddBackupJobsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Job) {
            throw "invalid value for 'Job', 'Job' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "job" = ${Job}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddBackupJobsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddBackupJobsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddBackupJobsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddBackupJobsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddBackupJobsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddBackupJobsRequest
        $AllProperties = ("job")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'job' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "job"))) {
            throw "Error! JSON cannot be serialized due to the required property 'job' missing."
        } else {
            $Job = $JsonParameters.PSobject.Properties["job"].value
        }

        $PSO = [PSCustomObject]@{
            "job" = ${Job}
        }

        return $PSO
    }

}

