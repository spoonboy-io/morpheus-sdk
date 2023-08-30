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
.PARAMETER MetricName
No description available.
.OUTPUTS

CheckCheckType<PSCustomObject>
#>

function Initialize-CheckCheckType {
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
        [String]
        ${MetricName}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CheckCheckType' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "code" = ${Code}
            "name" = ${Name}
            "metricName" = ${MetricName}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CheckCheckType<PSCustomObject>

.DESCRIPTION

Convert from JSON to CheckCheckType<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CheckCheckType<PSCustomObject>
#>
function ConvertFrom-JsonToCheckCheckType {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CheckCheckType' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CheckCheckType
        $AllProperties = ("id", "code", "name", "metricName")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "metricName"))) { #optional property not found
            $MetricName = $null
        } else {
            $MetricName = $JsonParameters.PSobject.Properties["metricName"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "code" = ${Code}
            "name" = ${Name}
            "metricName" = ${MetricName}
        }

        return $PSO
    }

}
