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
.PARAMETER AccountId
No description available.
.PARAMETER Instance
No description available.
.PARAMETER ContainerType
No description available.
.PARAMETER ContainerTypeSet
No description available.
.PARAMETER Server
No description available.
.PARAMETER Cloud
No description available.
.PARAMETER Name
No description available.
.PARAMETER Ip
No description available.
.PARAMETER InternalIp
No description available.
.PARAMETER InternalHostname
No description available.
.PARAMETER ExternalHostname
No description available.
.PARAMETER ExternalDomain
No description available.
.PARAMETER ExternalFqdn
No description available.
.PARAMETER Ports
No description available.
.PARAMETER Plan
No description available.
.PARAMETER DateCreated
No description available.
.PARAMETER LastUpdated
No description available.
.PARAMETER StatsEnabled
No description available.
.PARAMETER Status
No description available.
.PARAMETER UserStatus
No description available.
.PARAMETER EnvironmentPrefix
No description available.
.PARAMETER Stats
No description available.
.PARAMETER RuntimeInfo
No description available.
.PARAMETER ContainerVersion
No description available.
.PARAMETER RepositoryImage
No description available.
.PARAMETER PlanCategory
No description available.
.PARAMETER Hostname
No description available.
.PARAMETER DomainName
No description available.
.PARAMETER VolumeCreated
No description available.
.PARAMETER ContainerCreated
No description available.
.PARAMETER MaxStorage
No description available.
.PARAMETER MaxMemory
No description available.
.PARAMETER MaxCores
No description available.
.PARAMETER MaxCpu
No description available.
.PARAMETER AvailableActions
No description available.
.PARAMETER ConfigGroup
No description available.
.PARAMETER ConfigId
No description available.
.PARAMETER ConfigRole
No description available.
.PARAMETER HourlyCost
No description available.
.PARAMETER HourlyPrice
No description available.
.OUTPUTS

Container<PSCustomObject>
#>

function Initialize-Container {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int32]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Uuid},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int32]]
        ${AccountId},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Instance},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ContainerType},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ContainerTypeSet},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Server},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Cloud},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Ip},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InternalIp},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InternalHostname},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalHostname},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalDomain},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalFqdn},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Ports},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Plan},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${StatsEnabled},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Status},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${UserStatus},
        [Parameter(Position = 22, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${EnvironmentPrefix},
        [Parameter(Position = 23, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Stats},
        [Parameter(Position = 24, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${RuntimeInfo},
        [Parameter(Position = 25, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ContainerVersion},
        [Parameter(Position = 26, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RepositoryImage},
        [Parameter(Position = 27, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${PlanCategory},
        [Parameter(Position = 28, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Hostname},
        [Parameter(Position = 29, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DomainName},
        [Parameter(Position = 30, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${VolumeCreated},
        [Parameter(Position = 31, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${ContainerCreated},
        [Parameter(Position = 32, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int32]]
        ${MaxStorage},
        [Parameter(Position = 33, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int32]]
        ${MaxMemory},
        [Parameter(Position = 34, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int32]]
        ${MaxCores},
        [Parameter(Position = 35, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int32]]
        ${MaxCpu},
        [Parameter(Position = 36, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${AvailableActions},
        [Parameter(Position = 37, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ConfigGroup},
        [Parameter(Position = 38, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ConfigId},
        [Parameter(Position = 39, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ConfigRole},
        [Parameter(Position = 40, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Double]]
        ${HourlyCost},
        [Parameter(Position = 41, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Double]]
        ${HourlyPrice}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => Container' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "uuid" = ${Uuid}
            "accountId" = ${AccountId}
            "instance" = ${Instance}
            "containerType" = ${ContainerType}
            "containerTypeSet" = ${ContainerTypeSet}
            "server" = ${Server}
            "cloud" = ${Cloud}
            "name" = ${Name}
            "ip" = ${Ip}
            "internalIp" = ${InternalIp}
            "internalHostname" = ${InternalHostname}
            "externalHostname" = ${ExternalHostname}
            "externalDomain" = ${ExternalDomain}
            "externalFqdn" = ${ExternalFqdn}
            "ports" = ${Ports}
            "plan" = ${Plan}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "statsEnabled" = ${StatsEnabled}
            "status" = ${Status}
            "userStatus" = ${UserStatus}
            "environmentPrefix" = ${EnvironmentPrefix}
            "stats" = ${Stats}
            "runtimeInfo" = ${RuntimeInfo}
            "containerVersion" = ${ContainerVersion}
            "repositoryImage" = ${RepositoryImage}
            "planCategory" = ${PlanCategory}
            "hostname" = ${Hostname}
            "domainName" = ${DomainName}
            "volumeCreated" = ${VolumeCreated}
            "containerCreated" = ${ContainerCreated}
            "maxStorage" = ${MaxStorage}
            "maxMemory" = ${MaxMemory}
            "maxCores" = ${MaxCores}
            "maxCpu" = ${MaxCpu}
            "availableActions" = ${AvailableActions}
            "configGroup" = ${ConfigGroup}
            "configId" = ${ConfigId}
            "configRole" = ${ConfigRole}
            "hourlyCost" = ${HourlyCost}
            "hourlyPrice" = ${HourlyPrice}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to Container<PSCustomObject>

.DESCRIPTION

Convert from JSON to Container<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

Container<PSCustomObject>
#>
function ConvertFrom-JsonToContainer {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => Container' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in Container
        $AllProperties = ("id", "uuid", "accountId", "instance", "containerType", "containerTypeSet", "server", "cloud", "name", "ip", "internalIp", "internalHostname", "externalHostname", "externalDomain", "externalFqdn", "ports", "plan", "dateCreated", "lastUpdated", "statsEnabled", "status", "userStatus", "environmentPrefix", "stats", "runtimeInfo", "containerVersion", "repositoryImage", "planCategory", "hostname", "domainName", "volumeCreated", "containerCreated", "maxStorage", "maxMemory", "maxCores", "maxCpu", "availableActions", "configGroup", "configId", "configRole", "hourlyCost", "hourlyPrice")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "accountId"))) { #optional property not found
            $AccountId = $null
        } else {
            $AccountId = $JsonParameters.PSobject.Properties["accountId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instance"))) { #optional property not found
            $Instance = $null
        } else {
            $Instance = $JsonParameters.PSobject.Properties["instance"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "containerType"))) { #optional property not found
            $ContainerType = $null
        } else {
            $ContainerType = $JsonParameters.PSobject.Properties["containerType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "containerTypeSet"))) { #optional property not found
            $ContainerTypeSet = $null
        } else {
            $ContainerTypeSet = $JsonParameters.PSobject.Properties["containerTypeSet"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "server"))) { #optional property not found
            $Server = $null
        } else {
            $Server = $JsonParameters.PSobject.Properties["server"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cloud"))) { #optional property not found
            $Cloud = $null
        } else {
            $Cloud = $JsonParameters.PSobject.Properties["cloud"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ip"))) { #optional property not found
            $Ip = $null
        } else {
            $Ip = $JsonParameters.PSobject.Properties["ip"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "internalIp"))) { #optional property not found
            $InternalIp = $null
        } else {
            $InternalIp = $JsonParameters.PSobject.Properties["internalIp"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "internalHostname"))) { #optional property not found
            $InternalHostname = $null
        } else {
            $InternalHostname = $JsonParameters.PSobject.Properties["internalHostname"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalHostname"))) { #optional property not found
            $ExternalHostname = $null
        } else {
            $ExternalHostname = $JsonParameters.PSobject.Properties["externalHostname"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalDomain"))) { #optional property not found
            $ExternalDomain = $null
        } else {
            $ExternalDomain = $JsonParameters.PSobject.Properties["externalDomain"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalFqdn"))) { #optional property not found
            $ExternalFqdn = $null
        } else {
            $ExternalFqdn = $JsonParameters.PSobject.Properties["externalFqdn"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ports"))) { #optional property not found
            $Ports = $null
        } else {
            $Ports = $JsonParameters.PSobject.Properties["ports"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "plan"))) { #optional property not found
            $Plan = $null
        } else {
            $Plan = $JsonParameters.PSobject.Properties["plan"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "statsEnabled"))) { #optional property not found
            $StatsEnabled = $null
        } else {
            $StatsEnabled = $JsonParameters.PSobject.Properties["statsEnabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "status"))) { #optional property not found
            $Status = $null
        } else {
            $Status = $JsonParameters.PSobject.Properties["status"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "userStatus"))) { #optional property not found
            $UserStatus = $null
        } else {
            $UserStatus = $JsonParameters.PSobject.Properties["userStatus"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "environmentPrefix"))) { #optional property not found
            $EnvironmentPrefix = $null
        } else {
            $EnvironmentPrefix = $JsonParameters.PSobject.Properties["environmentPrefix"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "stats"))) { #optional property not found
            $Stats = $null
        } else {
            $Stats = $JsonParameters.PSobject.Properties["stats"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "runtimeInfo"))) { #optional property not found
            $RuntimeInfo = $null
        } else {
            $RuntimeInfo = $JsonParameters.PSobject.Properties["runtimeInfo"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "containerVersion"))) { #optional property not found
            $ContainerVersion = $null
        } else {
            $ContainerVersion = $JsonParameters.PSobject.Properties["containerVersion"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "repositoryImage"))) { #optional property not found
            $RepositoryImage = $null
        } else {
            $RepositoryImage = $JsonParameters.PSobject.Properties["repositoryImage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "planCategory"))) { #optional property not found
            $PlanCategory = $null
        } else {
            $PlanCategory = $JsonParameters.PSobject.Properties["planCategory"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hostname"))) { #optional property not found
            $Hostname = $null
        } else {
            $Hostname = $JsonParameters.PSobject.Properties["hostname"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "domainName"))) { #optional property not found
            $DomainName = $null
        } else {
            $DomainName = $JsonParameters.PSobject.Properties["domainName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "volumeCreated"))) { #optional property not found
            $VolumeCreated = $null
        } else {
            $VolumeCreated = $JsonParameters.PSobject.Properties["volumeCreated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "containerCreated"))) { #optional property not found
            $ContainerCreated = $null
        } else {
            $ContainerCreated = $JsonParameters.PSobject.Properties["containerCreated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxStorage"))) { #optional property not found
            $MaxStorage = $null
        } else {
            $MaxStorage = $JsonParameters.PSobject.Properties["maxStorage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxMemory"))) { #optional property not found
            $MaxMemory = $null
        } else {
            $MaxMemory = $JsonParameters.PSobject.Properties["maxMemory"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxCores"))) { #optional property not found
            $MaxCores = $null
        } else {
            $MaxCores = $JsonParameters.PSobject.Properties["maxCores"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxCpu"))) { #optional property not found
            $MaxCpu = $null
        } else {
            $MaxCpu = $JsonParameters.PSobject.Properties["maxCpu"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "availableActions"))) { #optional property not found
            $AvailableActions = $null
        } else {
            $AvailableActions = $JsonParameters.PSobject.Properties["availableActions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "configGroup"))) { #optional property not found
            $ConfigGroup = $null
        } else {
            $ConfigGroup = $JsonParameters.PSobject.Properties["configGroup"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "configId"))) { #optional property not found
            $ConfigId = $null
        } else {
            $ConfigId = $JsonParameters.PSobject.Properties["configId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "configRole"))) { #optional property not found
            $ConfigRole = $null
        } else {
            $ConfigRole = $JsonParameters.PSobject.Properties["configRole"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hourlyCost"))) { #optional property not found
            $HourlyCost = $null
        } else {
            $HourlyCost = $JsonParameters.PSobject.Properties["hourlyCost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hourlyPrice"))) { #optional property not found
            $HourlyPrice = $null
        } else {
            $HourlyPrice = $JsonParameters.PSobject.Properties["hourlyPrice"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "uuid" = ${Uuid}
            "accountId" = ${AccountId}
            "instance" = ${Instance}
            "containerType" = ${ContainerType}
            "containerTypeSet" = ${ContainerTypeSet}
            "server" = ${Server}
            "cloud" = ${Cloud}
            "name" = ${Name}
            "ip" = ${Ip}
            "internalIp" = ${InternalIp}
            "internalHostname" = ${InternalHostname}
            "externalHostname" = ${ExternalHostname}
            "externalDomain" = ${ExternalDomain}
            "externalFqdn" = ${ExternalFqdn}
            "ports" = ${Ports}
            "plan" = ${Plan}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "statsEnabled" = ${StatsEnabled}
            "status" = ${Status}
            "userStatus" = ${UserStatus}
            "environmentPrefix" = ${EnvironmentPrefix}
            "stats" = ${Stats}
            "runtimeInfo" = ${RuntimeInfo}
            "containerVersion" = ${ContainerVersion}
            "repositoryImage" = ${RepositoryImage}
            "planCategory" = ${PlanCategory}
            "hostname" = ${Hostname}
            "domainName" = ${DomainName}
            "volumeCreated" = ${VolumeCreated}
            "containerCreated" = ${ContainerCreated}
            "maxStorage" = ${MaxStorage}
            "maxMemory" = ${MaxMemory}
            "maxCores" = ${MaxCores}
            "maxCpu" = ${MaxCpu}
            "availableActions" = ${AvailableActions}
            "configGroup" = ${ConfigGroup}
            "configId" = ${ConfigId}
            "configRole" = ${ConfigRole}
            "hourlyCost" = ${HourlyCost}
            "hourlyPrice" = ${HourlyPrice}
        }

        return $PSO
    }

}

