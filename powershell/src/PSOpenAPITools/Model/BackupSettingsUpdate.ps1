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

.PARAMETER BackupsEnabled
Use this to enable / disable scheduled backups
.PARAMETER RetentionCount
Maximum number of successful backups to retain
.PARAMETER CreateBackups
Use this to enable / disable create backups
.PARAMETER BackupAppliance
When enabled, a Backup will be created to backup the Morpheus appliance database
.PARAMETER UpdateExisting
Use this to update existing backups with new settings
.PARAMETER DefaultSchedule
No description available.
.PARAMETER ClearDefaultSchedule
Use this to clear existing default backup schedule
.PARAMETER DefaultStorageBucket
No description available.
.PARAMETER ClearDefaultStorageBucket
Use this to clear default store bucket
.OUTPUTS

BackupSettingsUpdate<PSCustomObject>
#>

function Initialize-BackupSettingsUpdate {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${BackupsEnabled},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${RetentionCount},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${CreateBackups},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${BackupAppliance},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${UpdateExisting},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${DefaultSchedule},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${ClearDefaultSchedule},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${DefaultStorageBucket},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${ClearDefaultStorageBucket}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BackupSettingsUpdate' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "backupsEnabled" = ${BackupsEnabled}
            "retentionCount" = ${RetentionCount}
            "createBackups" = ${CreateBackups}
            "backupAppliance" = ${BackupAppliance}
            "updateExisting" = ${UpdateExisting}
            "defaultSchedule" = ${DefaultSchedule}
            "clearDefaultSchedule" = ${ClearDefaultSchedule}
            "defaultStorageBucket" = ${DefaultStorageBucket}
            "clearDefaultStorageBucket" = ${ClearDefaultStorageBucket}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BackupSettingsUpdate<PSCustomObject>

.DESCRIPTION

Convert from JSON to BackupSettingsUpdate<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BackupSettingsUpdate<PSCustomObject>
#>
function ConvertFrom-JsonToBackupSettingsUpdate {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BackupSettingsUpdate' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BackupSettingsUpdate
        $AllProperties = ("backupsEnabled", "retentionCount", "createBackups", "backupAppliance", "updateExisting", "defaultSchedule", "clearDefaultSchedule", "defaultStorageBucket", "clearDefaultStorageBucket")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "backupsEnabled"))) { #optional property not found
            $BackupsEnabled = $null
        } else {
            $BackupsEnabled = $JsonParameters.PSobject.Properties["backupsEnabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "retentionCount"))) { #optional property not found
            $RetentionCount = $null
        } else {
            $RetentionCount = $JsonParameters.PSobject.Properties["retentionCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createBackups"))) { #optional property not found
            $CreateBackups = $null
        } else {
            $CreateBackups = $JsonParameters.PSobject.Properties["createBackups"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "backupAppliance"))) { #optional property not found
            $BackupAppliance = $null
        } else {
            $BackupAppliance = $JsonParameters.PSobject.Properties["backupAppliance"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "updateExisting"))) { #optional property not found
            $UpdateExisting = $null
        } else {
            $UpdateExisting = $JsonParameters.PSobject.Properties["updateExisting"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "defaultSchedule"))) { #optional property not found
            $DefaultSchedule = $null
        } else {
            $DefaultSchedule = $JsonParameters.PSobject.Properties["defaultSchedule"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "clearDefaultSchedule"))) { #optional property not found
            $ClearDefaultSchedule = $null
        } else {
            $ClearDefaultSchedule = $JsonParameters.PSobject.Properties["clearDefaultSchedule"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "defaultStorageBucket"))) { #optional property not found
            $DefaultStorageBucket = $null
        } else {
            $DefaultStorageBucket = $JsonParameters.PSobject.Properties["defaultStorageBucket"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "clearDefaultStorageBucket"))) { #optional property not found
            $ClearDefaultStorageBucket = $null
        } else {
            $ClearDefaultStorageBucket = $JsonParameters.PSobject.Properties["clearDefaultStorageBucket"].value
        }

        $PSO = [PSCustomObject]@{
            "backupsEnabled" = ${BackupsEnabled}
            "retentionCount" = ${RetentionCount}
            "createBackups" = ${CreateBackups}
            "backupAppliance" = ${BackupAppliance}
            "updateExisting" = ${UpdateExisting}
            "defaultSchedule" = ${DefaultSchedule}
            "clearDefaultSchedule" = ${ClearDefaultSchedule}
            "defaultStorageBucket" = ${DefaultStorageBucket}
            "clearDefaultStorageBucket" = ${ClearDefaultStorageBucket}
        }

        return $PSO
    }

}

