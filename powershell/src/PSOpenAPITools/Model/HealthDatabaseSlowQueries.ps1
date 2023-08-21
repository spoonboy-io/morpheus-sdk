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

.PARAMETER Count
No description available.
.PARAMETER Time
No description available.
.PARAMETER Query
No description available.
.OUTPUTS

HealthDatabaseSlowQueries<PSCustomObject>
#>

function Initialize-HealthDatabaseSlowQueries {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Count},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Time},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Query}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => HealthDatabaseSlowQueries' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "count" = ${Count}
            "time" = ${Time}
            "query" = ${Query}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to HealthDatabaseSlowQueries<PSCustomObject>

.DESCRIPTION

Convert from JSON to HealthDatabaseSlowQueries<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

HealthDatabaseSlowQueries<PSCustomObject>
#>
function ConvertFrom-JsonToHealthDatabaseSlowQueries {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => HealthDatabaseSlowQueries' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in HealthDatabaseSlowQueries
        $AllProperties = ("count", "time", "query")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "count"))) { #optional property not found
            $Count = $null
        } else {
            $Count = $JsonParameters.PSobject.Properties["count"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "time"))) { #optional property not found
            $Time = $null
        } else {
            $Time = $JsonParameters.PSobject.Properties["time"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "query"))) { #optional property not found
            $Query = $null
        } else {
            $Query = $JsonParameters.PSobject.Properties["query"].value
        }

        $PSO = [PSCustomObject]@{
            "count" = ${Count}
            "time" = ${Time}
            "query" = ${Query}
        }

        return $PSO
    }

}

