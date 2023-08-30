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

.PARAMETER Name
Virtual Desktop name
.PARAMETER Description
Virtual Desktop description
.PARAMETER Owner
Owner (User) ID
.PARAMETER MinIdle
Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby. 
.PARAMETER InitialPoolSize
The initial size of the pool to be allocated on creation
.PARAMETER MaxIdle
Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances 
.PARAMETER MaxPoolSize
Max limit on number of allocations and instances within the pool. 
.PARAMETER AllocationTimeoutMinutes
Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence. 
.PARAMETER PersistentUser
Persistent Desktop Pool
.PARAMETER Recyclable
Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD)
.PARAMETER AllowCopy
Allow copy from desktop
.PARAMETER AllowPrinter
Allow local printers from Desktop
.PARAMETER AllowFileshare
Allow File Share
.PARAMETER AllowHypervisorConsole
Allow hypervisor console
.PARAMETER AutoCreateLocalUserOnReservation
Auto create local user upon reservation
.PARAMETER Enabled
Can be used to enable or disable the VDI pool
.PARAMETER IconPath
The relative location of an icon image
.PARAMETER Apps
Array of VDI App IDs
.PARAMETER Gateway
VDI Gateway ID
.PARAMETER InstanceConfig
Instance Config JSON. Passing as a string will preserve property order.  See `config` object for required values.
.PARAMETER Config
No description available.
.PARAMETER GuestConsoleJumpHost
Guest Console Jump Host
.PARAMETER GuestConsoleJumpPort
Guest Console Jump Port
.PARAMETER GuestConsoleJumpUsername
Guest Console Jump Username
.PARAMETER GuestConsoleJumpPassword
Guest Console Jump Password
.PARAMETER GuestConsoleJumpKeypair
Guest Console Jump Key Pair. see `Key Pair`
.OUTPUTS

AddVDIPoolsRequestVdiPoolOneOf1<PSCustomObject>
#>

function Initialize-AddVDIPoolsRequestVdiPoolOneOf1 {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Owner},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${MinIdle},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${InitialPoolSize},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${MaxIdle},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [Decimal]
        ${MaxPoolSize},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${AllocationTimeoutMinutes},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${PersistentUser} = $false,
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Recyclable} = $false,
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${AllowCopy} = $false,
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${AllowPrinter} = $false,
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${AllowFileshare} = $false,
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${AllowHypervisorConsole} = $false,
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${AutoCreateLocalUserOnReservation} = $false,
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled} = $true,
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${IconPath},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${Apps},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Gateway},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InstanceConfig},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Config},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GuestConsoleJumpHost},
        [Parameter(Position = 22, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${GuestConsoleJumpPort},
        [Parameter(Position = 23, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GuestConsoleJumpUsername},
        [Parameter(Position = 24, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GuestConsoleJumpPassword},
        [Parameter(Position = 25, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${GuestConsoleJumpKeypair}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddVDIPoolsRequestVdiPoolOneOf1' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }

        if ($null -eq $MaxPoolSize) {
            throw "invalid value for 'MaxPoolSize', 'MaxPoolSize' cannot be null."
        }

        if ($null -eq $Config) {
            throw "invalid value for 'Config', 'Config' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "owner" = ${Owner}
            "minIdle" = ${MinIdle}
            "initialPoolSize" = ${InitialPoolSize}
            "maxIdle" = ${MaxIdle}
            "maxPoolSize" = ${MaxPoolSize}
            "allocationTimeoutMinutes" = ${AllocationTimeoutMinutes}
            "persistentUser" = ${PersistentUser}
            "recyclable" = ${Recyclable}
            "allowCopy" = ${AllowCopy}
            "allowPrinter" = ${AllowPrinter}
            "allowFileshare" = ${AllowFileshare}
            "allowHypervisorConsole" = ${AllowHypervisorConsole}
            "autoCreateLocalUserOnReservation" = ${AutoCreateLocalUserOnReservation}
            "enabled" = ${Enabled}
            "iconPath" = ${IconPath}
            "apps" = ${Apps}
            "gateway" = ${Gateway}
            "instanceConfig" = ${InstanceConfig}
            "config" = ${Config}
            "guestConsoleJumpHost" = ${GuestConsoleJumpHost}
            "guestConsoleJumpPort" = ${GuestConsoleJumpPort}
            "guestConsoleJumpUsername" = ${GuestConsoleJumpUsername}
            "guestConsoleJumpPassword" = ${GuestConsoleJumpPassword}
            "guestConsoleJumpKeypair" = ${GuestConsoleJumpKeypair}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddVDIPoolsRequestVdiPoolOneOf1<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddVDIPoolsRequestVdiPoolOneOf1<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddVDIPoolsRequestVdiPoolOneOf1<PSCustomObject>
#>
function ConvertFrom-JsonToAddVDIPoolsRequestVdiPoolOneOf1 {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddVDIPoolsRequestVdiPoolOneOf1' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddVDIPoolsRequestVdiPoolOneOf1
        $AllProperties = ("name", "description", "owner", "minIdle", "initialPoolSize", "maxIdle", "maxPoolSize", "allocationTimeoutMinutes", "persistentUser", "recyclable", "allowCopy", "allowPrinter", "allowFileshare", "allowHypervisorConsole", "autoCreateLocalUserOnReservation", "enabled", "iconPath", "apps", "gateway", "instanceConfig", "config", "guestConsoleJumpHost", "guestConsoleJumpPort", "guestConsoleJumpUsername", "guestConsoleJumpPassword", "guestConsoleJumpKeypair")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'name' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property 'name' missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxPoolSize"))) {
            throw "Error! JSON cannot be serialized due to the required property 'maxPoolSize' missing."
        } else {
            $MaxPoolSize = $JsonParameters.PSobject.Properties["maxPoolSize"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "config"))) {
            throw "Error! JSON cannot be serialized due to the required property 'config' missing."
        } else {
            $Config = $JsonParameters.PSobject.Properties["config"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "owner"))) { #optional property not found
            $Owner = $null
        } else {
            $Owner = $JsonParameters.PSobject.Properties["owner"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "minIdle"))) { #optional property not found
            $MinIdle = $null
        } else {
            $MinIdle = $JsonParameters.PSobject.Properties["minIdle"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "initialPoolSize"))) { #optional property not found
            $InitialPoolSize = $null
        } else {
            $InitialPoolSize = $JsonParameters.PSobject.Properties["initialPoolSize"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxIdle"))) { #optional property not found
            $MaxIdle = $null
        } else {
            $MaxIdle = $JsonParameters.PSobject.Properties["maxIdle"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "allocationTimeoutMinutes"))) { #optional property not found
            $AllocationTimeoutMinutes = $null
        } else {
            $AllocationTimeoutMinutes = $JsonParameters.PSobject.Properties["allocationTimeoutMinutes"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "persistentUser"))) { #optional property not found
            $PersistentUser = $null
        } else {
            $PersistentUser = $JsonParameters.PSobject.Properties["persistentUser"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "recyclable"))) { #optional property not found
            $Recyclable = $null
        } else {
            $Recyclable = $JsonParameters.PSobject.Properties["recyclable"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "allowCopy"))) { #optional property not found
            $AllowCopy = $null
        } else {
            $AllowCopy = $JsonParameters.PSobject.Properties["allowCopy"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "allowPrinter"))) { #optional property not found
            $AllowPrinter = $null
        } else {
            $AllowPrinter = $JsonParameters.PSobject.Properties["allowPrinter"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "allowFileshare"))) { #optional property not found
            $AllowFileshare = $null
        } else {
            $AllowFileshare = $JsonParameters.PSobject.Properties["allowFileshare"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "allowHypervisorConsole"))) { #optional property not found
            $AllowHypervisorConsole = $null
        } else {
            $AllowHypervisorConsole = $JsonParameters.PSobject.Properties["allowHypervisorConsole"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "autoCreateLocalUserOnReservation"))) { #optional property not found
            $AutoCreateLocalUserOnReservation = $null
        } else {
            $AutoCreateLocalUserOnReservation = $JsonParameters.PSobject.Properties["autoCreateLocalUserOnReservation"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "iconPath"))) { #optional property not found
            $IconPath = $null
        } else {
            $IconPath = $JsonParameters.PSobject.Properties["iconPath"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "apps"))) { #optional property not found
            $Apps = $null
        } else {
            $Apps = $JsonParameters.PSobject.Properties["apps"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "gateway"))) { #optional property not found
            $Gateway = $null
        } else {
            $Gateway = $JsonParameters.PSobject.Properties["gateway"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceConfig"))) { #optional property not found
            $InstanceConfig = $null
        } else {
            $InstanceConfig = $JsonParameters.PSobject.Properties["instanceConfig"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "guestConsoleJumpHost"))) { #optional property not found
            $GuestConsoleJumpHost = $null
        } else {
            $GuestConsoleJumpHost = $JsonParameters.PSobject.Properties["guestConsoleJumpHost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "guestConsoleJumpPort"))) { #optional property not found
            $GuestConsoleJumpPort = $null
        } else {
            $GuestConsoleJumpPort = $JsonParameters.PSobject.Properties["guestConsoleJumpPort"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "guestConsoleJumpUsername"))) { #optional property not found
            $GuestConsoleJumpUsername = $null
        } else {
            $GuestConsoleJumpUsername = $JsonParameters.PSobject.Properties["guestConsoleJumpUsername"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "guestConsoleJumpPassword"))) { #optional property not found
            $GuestConsoleJumpPassword = $null
        } else {
            $GuestConsoleJumpPassword = $JsonParameters.PSobject.Properties["guestConsoleJumpPassword"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "guestConsoleJumpKeypair"))) { #optional property not found
            $GuestConsoleJumpKeypair = $null
        } else {
            $GuestConsoleJumpKeypair = $JsonParameters.PSobject.Properties["guestConsoleJumpKeypair"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "owner" = ${Owner}
            "minIdle" = ${MinIdle}
            "initialPoolSize" = ${InitialPoolSize}
            "maxIdle" = ${MaxIdle}
            "maxPoolSize" = ${MaxPoolSize}
            "allocationTimeoutMinutes" = ${AllocationTimeoutMinutes}
            "persistentUser" = ${PersistentUser}
            "recyclable" = ${Recyclable}
            "allowCopy" = ${AllowCopy}
            "allowPrinter" = ${AllowPrinter}
            "allowFileshare" = ${AllowFileshare}
            "allowHypervisorConsole" = ${AllowHypervisorConsole}
            "autoCreateLocalUserOnReservation" = ${AutoCreateLocalUserOnReservation}
            "enabled" = ${Enabled}
            "iconPath" = ${IconPath}
            "apps" = ${Apps}
            "gateway" = ${Gateway}
            "instanceConfig" = ${InstanceConfig}
            "config" = ${Config}
            "guestConsoleJumpHost" = ${GuestConsoleJumpHost}
            "guestConsoleJumpPort" = ${GuestConsoleJumpPort}
            "guestConsoleJumpUsername" = ${GuestConsoleJumpUsername}
            "guestConsoleJumpPassword" = ${GuestConsoleJumpPassword}
            "guestConsoleJumpKeypair" = ${GuestConsoleJumpKeypair}
        }

        return $PSO
    }

}
