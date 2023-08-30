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

.PARAMETER LocationType
No description available.
.PARAMETER Name
A name for the backup
.PARAMETER ServerId
The ID of the server to backup
.PARAMETER BackupType
The backup type code, options vary by the type of cloud and source
.PARAMETER JobAction
Create a new backup job, clone an existing job or add the new backup to an existing job
.PARAMETER JobId
The ID of the job to clone or add to. Only applies to jobAction `clone` and `addTo`.
.PARAMETER JobName
Name for new job. Only applies to jobAction `new` and `clone`.
.PARAMETER JobSchedule
The ID of the execute schedule for new job. See Execute Schedules. Only applies to jobAction `new` and `clone`.
.PARAMETER RetentionCount
Retention Count for new job. By default the backup settings value will be used. Only applies to jobAction `new` and `clone`.
.OUTPUTS

BackupTypeServer<PSCustomObject>
#>

function Initialize-BackupTypeServer {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("instance")]
        [String]
        ${LocationType},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ServerId},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("alibabaSnapshot", "amazonSnapshot", "avamarVMWareBackup", "azureSnapshot", "bluemixSnapshot", "commvaultFileBackup", "commvaultOpenstackBackup", "commvaultVMWareBackup", "digitaloceanSnapshot", "directoryBackup", "esxiSnapshot", "fileBackup", "fusionSnapshot", "googleSnapshot", "huaweiSnapshot", "hypervSnapshot", "kvm", "lvmImage", "lvmMigration", "lvmSnapshot", "MongoDB", "morpheusAppliance", "morpheusContainerBackup", "morpheusStorageBackup", "morpheusVmBackup", "MySQL", "nutanixSnapshot", "openstackSnapshot", "opentelekomSnapshot", "oracleCloudSnapshot", "Postgres", "powervcSnapshot", "rubrikVMWareBackup", "scvmmSnapshot", "softlayerSnapshot", "SqlServer", "tarDirectoryBackup", "upCloudSnapshot", "vcdSnapshot", "veeamHypervBackup", "veeamScvmmBackup", "veeamVcdBackup", "veeamVMWareBackup", "virtustreamSnapshot", "vmwareSnapshot", "winMigration", "xenSnapshot")]
        [String]
        ${BackupType},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("new", "clone", "addTo")]
        [String]
        ${JobAction},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${JobId},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${JobName},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${JobSchedule},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${RetentionCount}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BackupTypeServer' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $LocationType) {
            throw "invalid value for 'LocationType', 'LocationType' cannot be null."
        }

        if ($null -eq $Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }

        if ($null -eq $BackupType) {
            throw "invalid value for 'BackupType', 'BackupType' cannot be null."
        }

        if ($null -eq $JobAction) {
            throw "invalid value for 'JobAction', 'JobAction' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "locationType" = ${LocationType}
            "name" = ${Name}
            "serverId" = ${ServerId}
            "backupType" = ${BackupType}
            "jobAction" = ${JobAction}
            "jobId" = ${JobId}
            "jobName" = ${JobName}
            "jobSchedule" = ${JobSchedule}
            "retentionCount" = ${RetentionCount}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BackupTypeServer<PSCustomObject>

.DESCRIPTION

Convert from JSON to BackupTypeServer<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BackupTypeServer<PSCustomObject>
#>
function ConvertFrom-JsonToBackupTypeServer {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BackupTypeServer' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BackupTypeServer
        $AllProperties = ("locationType", "name", "serverId", "backupType", "jobAction", "jobId", "jobName", "jobSchedule", "retentionCount")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'locationType' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "locationType"))) {
            throw "Error! JSON cannot be serialized due to the required property 'locationType' missing."
        } else {
            $LocationType = $JsonParameters.PSobject.Properties["locationType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property 'name' missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "backupType"))) {
            throw "Error! JSON cannot be serialized due to the required property 'backupType' missing."
        } else {
            $BackupType = $JsonParameters.PSobject.Properties["backupType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "jobAction"))) {
            throw "Error! JSON cannot be serialized due to the required property 'jobAction' missing."
        } else {
            $JobAction = $JsonParameters.PSobject.Properties["jobAction"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serverId"))) { #optional property not found
            $ServerId = $null
        } else {
            $ServerId = $JsonParameters.PSobject.Properties["serverId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "jobId"))) { #optional property not found
            $JobId = $null
        } else {
            $JobId = $JsonParameters.PSobject.Properties["jobId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "jobName"))) { #optional property not found
            $JobName = $null
        } else {
            $JobName = $JsonParameters.PSobject.Properties["jobName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "jobSchedule"))) { #optional property not found
            $JobSchedule = $null
        } else {
            $JobSchedule = $JsonParameters.PSobject.Properties["jobSchedule"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "retentionCount"))) { #optional property not found
            $RetentionCount = $null
        } else {
            $RetentionCount = $JsonParameters.PSobject.Properties["retentionCount"].value
        }

        $PSO = [PSCustomObject]@{
            "locationType" = ${LocationType}
            "name" = ${Name}
            "serverId" = ${ServerId}
            "backupType" = ${BackupType}
            "jobAction" = ${JobAction}
            "jobId" = ${JobId}
            "jobName" = ${JobName}
            "jobSchedule" = ${JobSchedule}
            "retentionCount" = ${RetentionCount}
        }

        return $PSO
    }

}

