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
.PARAMETER Uuid
No description available.
.PARAMETER ExternalId
No description available.
.PARAMETER InternalId
No description available.
.PARAMETER ExternalUniqueId
No description available.
.PARAMETER Name
No description available.
.PARAMETER ExternalName
No description available.
.PARAMETER Hostname
No description available.
.PARAMETER AccountId
No description available.
.PARAMETER Account
No description available.
.PARAMETER Owner
No description available.
.PARAMETER Zone
No description available.
.PARAMETER Plan
No description available.
.PARAMETER ComputeServerType
No description available.
.PARAMETER Visibility
No description available.
.PARAMETER Description
No description available.
.PARAMETER ZoneId
No description available.
.PARAMETER SiteId
No description available.
.PARAMETER ResourcePoolId
No description available.
.PARAMETER FolderId
No description available.
.PARAMETER SshHost
No description available.
.PARAMETER SshPort
No description available.
.PARAMETER ExternalIp
No description available.
.PARAMETER InternalIp
No description available.
.PARAMETER VolumeId
No description available.
.PARAMETER Platform
No description available.
.PARAMETER PlatformVersion
No description available.
.PARAMETER SshUsername
No description available.
.PARAMETER SshPassword
No description available.
.PARAMETER SshPasswordHash
No description available.
.PARAMETER OsDevice
No description available.
.PARAMETER OsType
No description available.
.PARAMETER DataDevice
No description available.
.PARAMETER LvmEnabled
No description available.
.PARAMETER ApiKey
No description available.
.PARAMETER SoftwareRaid
No description available.
.PARAMETER DateCreated
No description available.
.PARAMETER LastUpdated
No description available.
.PARAMETER Stats
No description available.
.PARAMETER Status
No description available.
.PARAMETER StatusMessage
No description available.
.PARAMETER ErrorMessage
No description available.
.PARAMETER StatusDate
No description available.
.PARAMETER StatusPercent
No description available.
.PARAMETER StatusEta
No description available.
.PARAMETER PowerState
No description available.
.PARAMETER AgentInstalled
No description available.
.PARAMETER LastAgentUpdate
No description available.
.PARAMETER AgentVersion
No description available.
.PARAMETER MaxCores
No description available.
.PARAMETER CoresPerSocket
No description available.
.PARAMETER MaxMemory
No description available.
.PARAMETER MaxStorage
No description available.
.PARAMETER MaxCpu
No description available.
.PARAMETER HourlyPrice
No description available.
.PARAMETER SourceImage
No description available.
.PARAMETER ServerOs
No description available.
.PARAMETER Volumes
No description available.
.PARAMETER Controllers
No description available.
.PARAMETER Interfaces
No description available.
.PARAMETER Labels
No description available.
.PARAMETER Tags
No description available.
.PARAMETER Enabled
No description available.
.PARAMETER TagCompliant
No description available.
.PARAMETER Containers
No description available.
.PARAMETER GuestConsolePreferred
No description available.
.PARAMETER GuestConsoleType
No description available.
.PARAMETER GuestConsoleUsername
No description available.
.PARAMETER GuestConsolePassword
No description available.
.PARAMETER GuestConsolePasswordHash
No description available.
.PARAMETER GuestConsolePort
No description available.
.OUTPUTS

ClusterWorkers<PSCustomObject>
#>

function Initialize-ClusterWorkers {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Uuid},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalId},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InternalId},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalUniqueId},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalName},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Hostname},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${AccountId},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Account},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Owner},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Zone},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Plan},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ComputeServerType},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Visibility},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ZoneId},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${SiteId},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ResourcePoolId},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${FolderId},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SshHost},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${SshPort},
        [Parameter(Position = 22, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalIp},
        [Parameter(Position = 23, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InternalIp},
        [Parameter(Position = 24, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${VolumeId},
        [Parameter(Position = 25, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Platform},
        [Parameter(Position = 26, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${PlatformVersion},
        [Parameter(Position = 27, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SshUsername},
        [Parameter(Position = 28, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SshPassword},
        [Parameter(Position = 29, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SshPasswordHash},
        [Parameter(Position = 30, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${OsDevice},
        [Parameter(Position = 31, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${OsType},
        [Parameter(Position = 32, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DataDevice},
        [Parameter(Position = 33, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${LvmEnabled},
        [Parameter(Position = 34, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ApiKey},
        [Parameter(Position = 35, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${SoftwareRaid},
        [Parameter(Position = 36, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 37, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated},
        [Parameter(Position = 38, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Stats},
        [Parameter(Position = 39, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Status},
        [Parameter(Position = 40, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${StatusMessage},
        [Parameter(Position = 41, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ErrorMessage},
        [Parameter(Position = 42, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${StatusDate},
        [Parameter(Position = 43, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${StatusPercent},
        [Parameter(Position = 44, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${StatusEta},
        [Parameter(Position = 45, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${PowerState},
        [Parameter(Position = 46, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${AgentInstalled},
        [Parameter(Position = 47, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastAgentUpdate},
        [Parameter(Position = 48, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${AgentVersion},
        [Parameter(Position = 49, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${MaxCores},
        [Parameter(Position = 50, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CoresPerSocket},
        [Parameter(Position = 51, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${MaxMemory},
        [Parameter(Position = 52, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${MaxStorage},
        [Parameter(Position = 53, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${MaxCpu},
        [Parameter(Position = 54, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${HourlyPrice},
        [Parameter(Position = 55, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${SourceImage},
        [Parameter(Position = 56, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServerOs},
        [Parameter(Position = 57, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Volumes},
        [Parameter(Position = 58, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Controllers},
        [Parameter(Position = 59, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Interfaces},
        [Parameter(Position = 60, ValueFromPipelineByPropertyName = $true)]
        [String[]]
        ${Labels},
        [Parameter(Position = 61, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Tags},
        [Parameter(Position = 62, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled},
        [Parameter(Position = 63, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${TagCompliant},
        [Parameter(Position = 64, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${Containers},
        [Parameter(Position = 65, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${GuestConsolePreferred},
        [Parameter(Position = 66, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GuestConsoleType},
        [Parameter(Position = 67, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GuestConsoleUsername},
        [Parameter(Position = 68, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GuestConsolePassword},
        [Parameter(Position = 69, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GuestConsolePasswordHash},
        [Parameter(Position = 70, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GuestConsolePort}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ClusterWorkers' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "uuid" = ${Uuid}
            "externalId" = ${ExternalId}
            "internalId" = ${InternalId}
            "externalUniqueId" = ${ExternalUniqueId}
            "name" = ${Name}
            "externalName" = ${ExternalName}
            "hostname" = ${Hostname}
            "accountId" = ${AccountId}
            "account" = ${Account}
            "owner" = ${Owner}
            "zone" = ${Zone}
            "plan" = ${Plan}
            "computeServerType" = ${ComputeServerType}
            "visibility" = ${Visibility}
            "description" = ${Description}
            "zoneId" = ${ZoneId}
            "siteId" = ${SiteId}
            "resourcePoolId" = ${ResourcePoolId}
            "folderId" = ${FolderId}
            "sshHost" = ${SshHost}
            "sshPort" = ${SshPort}
            "externalIp" = ${ExternalIp}
            "internalIp" = ${InternalIp}
            "volumeId" = ${VolumeId}
            "platform" = ${Platform}
            "platformVersion" = ${PlatformVersion}
            "sshUsername" = ${SshUsername}
            "sshPassword" = ${SshPassword}
            "sshPasswordHash" = ${SshPasswordHash}
            "osDevice" = ${OsDevice}
            "osType" = ${OsType}
            "dataDevice" = ${DataDevice}
            "lvmEnabled" = ${LvmEnabled}
            "apiKey" = ${ApiKey}
            "softwareRaid" = ${SoftwareRaid}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "stats" = ${Stats}
            "status" = ${Status}
            "statusMessage" = ${StatusMessage}
            "errorMessage" = ${ErrorMessage}
            "statusDate" = ${StatusDate}
            "statusPercent" = ${StatusPercent}
            "statusEta" = ${StatusEta}
            "powerState" = ${PowerState}
            "agentInstalled" = ${AgentInstalled}
            "lastAgentUpdate" = ${LastAgentUpdate}
            "agentVersion" = ${AgentVersion}
            "maxCores" = ${MaxCores}
            "coresPerSocket" = ${CoresPerSocket}
            "maxMemory" = ${MaxMemory}
            "maxStorage" = ${MaxStorage}
            "maxCpu" = ${MaxCpu}
            "hourlyPrice" = ${HourlyPrice}
            "sourceImage" = ${SourceImage}
            "serverOs" = ${ServerOs}
            "volumes" = ${Volumes}
            "controllers" = ${Controllers}
            "interfaces" = ${Interfaces}
            "labels" = ${Labels}
            "tags" = ${Tags}
            "enabled" = ${Enabled}
            "tagCompliant" = ${TagCompliant}
            "containers" = ${Containers}
            "guestConsolePreferred" = ${GuestConsolePreferred}
            "guestConsoleType" = ${GuestConsoleType}
            "guestConsoleUsername" = ${GuestConsoleUsername}
            "guestConsolePassword" = ${GuestConsolePassword}
            "guestConsolePasswordHash" = ${GuestConsolePasswordHash}
            "guestConsolePort" = ${GuestConsolePort}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ClusterWorkers<PSCustomObject>

.DESCRIPTION

Convert from JSON to ClusterWorkers<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ClusterWorkers<PSCustomObject>
#>
function ConvertFrom-JsonToClusterWorkers {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ClusterWorkers' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ClusterWorkers
        $AllProperties = ("id", "uuid", "externalId", "internalId", "externalUniqueId", "name", "externalName", "hostname", "accountId", "account", "owner", "zone", "plan", "computeServerType", "visibility", "description", "zoneId", "siteId", "resourcePoolId", "folderId", "sshHost", "sshPort", "externalIp", "internalIp", "volumeId", "platform", "platformVersion", "sshUsername", "sshPassword", "sshPasswordHash", "osDevice", "osType", "dataDevice", "lvmEnabled", "apiKey", "softwareRaid", "dateCreated", "lastUpdated", "stats", "status", "statusMessage", "errorMessage", "statusDate", "statusPercent", "statusEta", "powerState", "agentInstalled", "lastAgentUpdate", "agentVersion", "maxCores", "coresPerSocket", "maxMemory", "maxStorage", "maxCpu", "hourlyPrice", "sourceImage", "serverOs", "volumes", "controllers", "interfaces", "labels", "tags", "enabled", "tagCompliant", "containers", "guestConsolePreferred", "guestConsoleType", "guestConsoleUsername", "guestConsolePassword", "guestConsolePasswordHash", "guestConsolePort")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "uuid"))) { #optional property not found
            $Uuid = $null
        } else {
            $Uuid = $JsonParameters.PSobject.Properties["uuid"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalId"))) { #optional property not found
            $ExternalId = $null
        } else {
            $ExternalId = $JsonParameters.PSobject.Properties["externalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "internalId"))) { #optional property not found
            $InternalId = $null
        } else {
            $InternalId = $JsonParameters.PSobject.Properties["internalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalUniqueId"))) { #optional property not found
            $ExternalUniqueId = $null
        } else {
            $ExternalUniqueId = $JsonParameters.PSobject.Properties["externalUniqueId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalName"))) { #optional property not found
            $ExternalName = $null
        } else {
            $ExternalName = $JsonParameters.PSobject.Properties["externalName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hostname"))) { #optional property not found
            $Hostname = $null
        } else {
            $Hostname = $JsonParameters.PSobject.Properties["hostname"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "accountId"))) { #optional property not found
            $AccountId = $null
        } else {
            $AccountId = $JsonParameters.PSobject.Properties["accountId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "account"))) { #optional property not found
            $Account = $null
        } else {
            $Account = $JsonParameters.PSobject.Properties["account"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "owner"))) { #optional property not found
            $Owner = $null
        } else {
            $Owner = $JsonParameters.PSobject.Properties["owner"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "zone"))) { #optional property not found
            $Zone = $null
        } else {
            $Zone = $JsonParameters.PSobject.Properties["zone"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "plan"))) { #optional property not found
            $Plan = $null
        } else {
            $Plan = $JsonParameters.PSobject.Properties["plan"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "computeServerType"))) { #optional property not found
            $ComputeServerType = $null
        } else {
            $ComputeServerType = $JsonParameters.PSobject.Properties["computeServerType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "visibility"))) { #optional property not found
            $Visibility = $null
        } else {
            $Visibility = $JsonParameters.PSobject.Properties["visibility"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "zoneId"))) { #optional property not found
            $ZoneId = $null
        } else {
            $ZoneId = $JsonParameters.PSobject.Properties["zoneId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "siteId"))) { #optional property not found
            $SiteId = $null
        } else {
            $SiteId = $JsonParameters.PSobject.Properties["siteId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "resourcePoolId"))) { #optional property not found
            $ResourcePoolId = $null
        } else {
            $ResourcePoolId = $JsonParameters.PSobject.Properties["resourcePoolId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "folderId"))) { #optional property not found
            $FolderId = $null
        } else {
            $FolderId = $JsonParameters.PSobject.Properties["folderId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sshHost"))) { #optional property not found
            $SshHost = $null
        } else {
            $SshHost = $JsonParameters.PSobject.Properties["sshHost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sshPort"))) { #optional property not found
            $SshPort = $null
        } else {
            $SshPort = $JsonParameters.PSobject.Properties["sshPort"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalIp"))) { #optional property not found
            $ExternalIp = $null
        } else {
            $ExternalIp = $JsonParameters.PSobject.Properties["externalIp"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "internalIp"))) { #optional property not found
            $InternalIp = $null
        } else {
            $InternalIp = $JsonParameters.PSobject.Properties["internalIp"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "volumeId"))) { #optional property not found
            $VolumeId = $null
        } else {
            $VolumeId = $JsonParameters.PSobject.Properties["volumeId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "platform"))) { #optional property not found
            $Platform = $null
        } else {
            $Platform = $JsonParameters.PSobject.Properties["platform"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "platformVersion"))) { #optional property not found
            $PlatformVersion = $null
        } else {
            $PlatformVersion = $JsonParameters.PSobject.Properties["platformVersion"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sshUsername"))) { #optional property not found
            $SshUsername = $null
        } else {
            $SshUsername = $JsonParameters.PSobject.Properties["sshUsername"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sshPassword"))) { #optional property not found
            $SshPassword = $null
        } else {
            $SshPassword = $JsonParameters.PSobject.Properties["sshPassword"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sshPasswordHash"))) { #optional property not found
            $SshPasswordHash = $null
        } else {
            $SshPasswordHash = $JsonParameters.PSobject.Properties["sshPasswordHash"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "osDevice"))) { #optional property not found
            $OsDevice = $null
        } else {
            $OsDevice = $JsonParameters.PSobject.Properties["osDevice"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "osType"))) { #optional property not found
            $OsType = $null
        } else {
            $OsType = $JsonParameters.PSobject.Properties["osType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dataDevice"))) { #optional property not found
            $DataDevice = $null
        } else {
            $DataDevice = $JsonParameters.PSobject.Properties["dataDevice"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lvmEnabled"))) { #optional property not found
            $LvmEnabled = $null
        } else {
            $LvmEnabled = $JsonParameters.PSobject.Properties["lvmEnabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "apiKey"))) { #optional property not found
            $ApiKey = $null
        } else {
            $ApiKey = $JsonParameters.PSobject.Properties["apiKey"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "softwareRaid"))) { #optional property not found
            $SoftwareRaid = $null
        } else {
            $SoftwareRaid = $JsonParameters.PSobject.Properties["softwareRaid"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dateCreated"))) { #optional property not found
            $DateCreated = $null
        } else {
            $DateCreated = $JsonParameters.PSobject.Properties["dateCreated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastUpdated"))) { #optional property not found
            $LastUpdated = $null
        } else {
            $LastUpdated = $JsonParameters.PSobject.Properties["lastUpdated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "stats"))) { #optional property not found
            $Stats = $null
        } else {
            $Stats = $JsonParameters.PSobject.Properties["stats"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "status"))) { #optional property not found
            $Status = $null
        } else {
            $Status = $JsonParameters.PSobject.Properties["status"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "statusMessage"))) { #optional property not found
            $StatusMessage = $null
        } else {
            $StatusMessage = $JsonParameters.PSobject.Properties["statusMessage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "errorMessage"))) { #optional property not found
            $ErrorMessage = $null
        } else {
            $ErrorMessage = $JsonParameters.PSobject.Properties["errorMessage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "statusDate"))) { #optional property not found
            $StatusDate = $null
        } else {
            $StatusDate = $JsonParameters.PSobject.Properties["statusDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "statusPercent"))) { #optional property not found
            $StatusPercent = $null
        } else {
            $StatusPercent = $JsonParameters.PSobject.Properties["statusPercent"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "statusEta"))) { #optional property not found
            $StatusEta = $null
        } else {
            $StatusEta = $JsonParameters.PSobject.Properties["statusEta"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "powerState"))) { #optional property not found
            $PowerState = $null
        } else {
            $PowerState = $JsonParameters.PSobject.Properties["powerState"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "agentInstalled"))) { #optional property not found
            $AgentInstalled = $null
        } else {
            $AgentInstalled = $JsonParameters.PSobject.Properties["agentInstalled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastAgentUpdate"))) { #optional property not found
            $LastAgentUpdate = $null
        } else {
            $LastAgentUpdate = $JsonParameters.PSobject.Properties["lastAgentUpdate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "agentVersion"))) { #optional property not found
            $AgentVersion = $null
        } else {
            $AgentVersion = $JsonParameters.PSobject.Properties["agentVersion"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxCores"))) { #optional property not found
            $MaxCores = $null
        } else {
            $MaxCores = $JsonParameters.PSobject.Properties["maxCores"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "coresPerSocket"))) { #optional property not found
            $CoresPerSocket = $null
        } else {
            $CoresPerSocket = $JsonParameters.PSobject.Properties["coresPerSocket"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxMemory"))) { #optional property not found
            $MaxMemory = $null
        } else {
            $MaxMemory = $JsonParameters.PSobject.Properties["maxMemory"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxStorage"))) { #optional property not found
            $MaxStorage = $null
        } else {
            $MaxStorage = $JsonParameters.PSobject.Properties["maxStorage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxCpu"))) { #optional property not found
            $MaxCpu = $null
        } else {
            $MaxCpu = $JsonParameters.PSobject.Properties["maxCpu"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hourlyPrice"))) { #optional property not found
            $HourlyPrice = $null
        } else {
            $HourlyPrice = $JsonParameters.PSobject.Properties["hourlyPrice"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sourceImage"))) { #optional property not found
            $SourceImage = $null
        } else {
            $SourceImage = $JsonParameters.PSobject.Properties["sourceImage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serverOs"))) { #optional property not found
            $ServerOs = $null
        } else {
            $ServerOs = $JsonParameters.PSobject.Properties["serverOs"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "volumes"))) { #optional property not found
            $Volumes = $null
        } else {
            $Volumes = $JsonParameters.PSobject.Properties["volumes"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "controllers"))) { #optional property not found
            $Controllers = $null
        } else {
            $Controllers = $JsonParameters.PSobject.Properties["controllers"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "interfaces"))) { #optional property not found
            $Interfaces = $null
        } else {
            $Interfaces = $JsonParameters.PSobject.Properties["interfaces"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "labels"))) { #optional property not found
            $Labels = $null
        } else {
            $Labels = $JsonParameters.PSobject.Properties["labels"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tags"))) { #optional property not found
            $Tags = $null
        } else {
            $Tags = $JsonParameters.PSobject.Properties["tags"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tagCompliant"))) { #optional property not found
            $TagCompliant = $null
        } else {
            $TagCompliant = $JsonParameters.PSobject.Properties["tagCompliant"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "containers"))) { #optional property not found
            $Containers = $null
        } else {
            $Containers = $JsonParameters.PSobject.Properties["containers"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "guestConsolePreferred"))) { #optional property not found
            $GuestConsolePreferred = $null
        } else {
            $GuestConsolePreferred = $JsonParameters.PSobject.Properties["guestConsolePreferred"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "guestConsoleType"))) { #optional property not found
            $GuestConsoleType = $null
        } else {
            $GuestConsoleType = $JsonParameters.PSobject.Properties["guestConsoleType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "guestConsoleUsername"))) { #optional property not found
            $GuestConsoleUsername = $null
        } else {
            $GuestConsoleUsername = $JsonParameters.PSobject.Properties["guestConsoleUsername"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "guestConsolePassword"))) { #optional property not found
            $GuestConsolePassword = $null
        } else {
            $GuestConsolePassword = $JsonParameters.PSobject.Properties["guestConsolePassword"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "guestConsolePasswordHash"))) { #optional property not found
            $GuestConsolePasswordHash = $null
        } else {
            $GuestConsolePasswordHash = $JsonParameters.PSobject.Properties["guestConsolePasswordHash"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "guestConsolePort"))) { #optional property not found
            $GuestConsolePort = $null
        } else {
            $GuestConsolePort = $JsonParameters.PSobject.Properties["guestConsolePort"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "uuid" = ${Uuid}
            "externalId" = ${ExternalId}
            "internalId" = ${InternalId}
            "externalUniqueId" = ${ExternalUniqueId}
            "name" = ${Name}
            "externalName" = ${ExternalName}
            "hostname" = ${Hostname}
            "accountId" = ${AccountId}
            "account" = ${Account}
            "owner" = ${Owner}
            "zone" = ${Zone}
            "plan" = ${Plan}
            "computeServerType" = ${ComputeServerType}
            "visibility" = ${Visibility}
            "description" = ${Description}
            "zoneId" = ${ZoneId}
            "siteId" = ${SiteId}
            "resourcePoolId" = ${ResourcePoolId}
            "folderId" = ${FolderId}
            "sshHost" = ${SshHost}
            "sshPort" = ${SshPort}
            "externalIp" = ${ExternalIp}
            "internalIp" = ${InternalIp}
            "volumeId" = ${VolumeId}
            "platform" = ${Platform}
            "platformVersion" = ${PlatformVersion}
            "sshUsername" = ${SshUsername}
            "sshPassword" = ${SshPassword}
            "sshPasswordHash" = ${SshPasswordHash}
            "osDevice" = ${OsDevice}
            "osType" = ${OsType}
            "dataDevice" = ${DataDevice}
            "lvmEnabled" = ${LvmEnabled}
            "apiKey" = ${ApiKey}
            "softwareRaid" = ${SoftwareRaid}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "stats" = ${Stats}
            "status" = ${Status}
            "statusMessage" = ${StatusMessage}
            "errorMessage" = ${ErrorMessage}
            "statusDate" = ${StatusDate}
            "statusPercent" = ${StatusPercent}
            "statusEta" = ${StatusEta}
            "powerState" = ${PowerState}
            "agentInstalled" = ${AgentInstalled}
            "lastAgentUpdate" = ${LastAgentUpdate}
            "agentVersion" = ${AgentVersion}
            "maxCores" = ${MaxCores}
            "coresPerSocket" = ${CoresPerSocket}
            "maxMemory" = ${MaxMemory}
            "maxStorage" = ${MaxStorage}
            "maxCpu" = ${MaxCpu}
            "hourlyPrice" = ${HourlyPrice}
            "sourceImage" = ${SourceImage}
            "serverOs" = ${ServerOs}
            "volumes" = ${Volumes}
            "controllers" = ${Controllers}
            "interfaces" = ${Interfaces}
            "labels" = ${Labels}
            "tags" = ${Tags}
            "enabled" = ${Enabled}
            "tagCompliant" = ${TagCompliant}
            "containers" = ${Containers}
            "guestConsolePreferred" = ${GuestConsolePreferred}
            "guestConsoleType" = ${GuestConsoleType}
            "guestConsoleUsername" = ${GuestConsoleUsername}
            "guestConsolePassword" = ${GuestConsolePassword}
            "guestConsolePasswordHash" = ${GuestConsolePasswordHash}
            "guestConsolePort" = ${GuestConsolePort}
        }

        return $PSO
    }

}
