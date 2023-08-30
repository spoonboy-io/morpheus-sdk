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

.PARAMETER Success
No description available.
.PARAMETER InstanceStats
No description available.
.PARAMETER Provisioning
No description available.
.PARAMETER Monitoring
No description available.
.PARAMETER Backups
No description available.
.PARAMETER Activity
No description available.
.PARAMETER LogStats
No description available.
.OUTPUTS

Dashboard<PSCustomObject>
#>

function Initialize-Dashboard {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Success},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${InstanceStats},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Provisioning},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Monitoring},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Backups},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Activity},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${LogStats}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => Dashboard' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "success" = ${Success}
            "instanceStats" = ${InstanceStats}
            "provisioning" = ${Provisioning}
            "monitoring" = ${Monitoring}
            "backups" = ${Backups}
            "activity" = ${Activity}
            "logStats" = ${LogStats}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to Dashboard<PSCustomObject>

.DESCRIPTION

Convert from JSON to Dashboard<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

Dashboard<PSCustomObject>
#>
function ConvertFrom-JsonToDashboard {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => Dashboard' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in Dashboard
        $AllProperties = ("success", "instanceStats", "provisioning", "monitoring", "backups", "activity", "logStats")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "success"))) { #optional property not found
            $Success = $null
        } else {
            $Success = $JsonParameters.PSobject.Properties["success"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceStats"))) { #optional property not found
            $InstanceStats = $null
        } else {
            $InstanceStats = $JsonParameters.PSobject.Properties["instanceStats"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "provisioning"))) { #optional property not found
            $Provisioning = $null
        } else {
            $Provisioning = $JsonParameters.PSobject.Properties["provisioning"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "monitoring"))) { #optional property not found
            $Monitoring = $null
        } else {
            $Monitoring = $JsonParameters.PSobject.Properties["monitoring"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "backups"))) { #optional property not found
            $Backups = $null
        } else {
            $Backups = $JsonParameters.PSobject.Properties["backups"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "activity"))) { #optional property not found
            $Activity = $null
        } else {
            $Activity = $JsonParameters.PSobject.Properties["activity"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "logStats"))) { #optional property not found
            $LogStats = $null
        } else {
            $LogStats = $JsonParameters.PSobject.Properties["logStats"].value
        }

        $PSO = [PSCustomObject]@{
            "success" = ${Success}
            "instanceStats" = ${InstanceStats}
            "provisioning" = ${Provisioning}
            "monitoring" = ${Monitoring}
            "backups" = ${Backups}
            "activity" = ${Activity}
            "logStats" = ${LogStats}
        }

        return $PSO
    }

}

