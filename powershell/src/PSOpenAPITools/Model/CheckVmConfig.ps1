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

.PARAMETER ContainerName
No description available.
.PARAMETER ExternalId
No description available.
.PARAMETER CheckUser
No description available.
.PARAMETER TextCheckOn
No description available.
.PARAMETER CheckPassword
No description available.
.PARAMETER WebTextMatch
No description available.
.PARAMETER CheckPasswordHash
No description available.
.PARAMETER TunnelOn
Set to on to turn on tunneling
.PARAMETER SshHost
Hostname or IP address of the proxy host
.PARAMETER SshPort
Port for SSH on the proxy host, defaults to 22
.PARAMETER SshUser
SSH user on the proxy host to login as
.PARAMETER SshPassword
Password for user, if not using key based authentication
.OUTPUTS

CheckVmConfig<PSCustomObject>
#>

function Initialize-CheckVmConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ContainerName},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalId},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CheckUser},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${TextCheckOn},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CheckPassword},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${WebTextMatch},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CheckPasswordHash},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("on", "off")]
        [String]
        ${TunnelOn} = "off",
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SshHost},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${SshPort},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SshUser},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SshPassword}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CheckVmConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $ContainerName) {
            throw "invalid value for 'ContainerName', 'ContainerName' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "containerName" = ${ContainerName}
            "externalId" = ${ExternalId}
            "checkUser" = ${CheckUser}
            "textCheckOn" = ${TextCheckOn}
            "checkPassword" = ${CheckPassword}
            "webTextMatch" = ${WebTextMatch}
            "checkPasswordHash" = ${CheckPasswordHash}
            "tunnelOn" = ${TunnelOn}
            "sshHost" = ${SshHost}
            "sshPort" = ${SshPort}
            "sshUser" = ${SshUser}
            "sshPassword" = ${SshPassword}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CheckVmConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to CheckVmConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CheckVmConfig<PSCustomObject>
#>
function ConvertFrom-JsonToCheckVmConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CheckVmConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CheckVmConfig
        $AllProperties = ("containerName", "externalId", "checkUser", "textCheckOn", "checkPassword", "webTextMatch", "checkPasswordHash", "tunnelOn", "sshHost", "sshPort", "sshUser", "sshPassword")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'containerName' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "containerName"))) {
            throw "Error! JSON cannot be serialized due to the required property 'containerName' missing."
        } else {
            $ContainerName = $JsonParameters.PSobject.Properties["containerName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalId"))) { #optional property not found
            $ExternalId = $null
        } else {
            $ExternalId = $JsonParameters.PSobject.Properties["externalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "checkUser"))) { #optional property not found
            $CheckUser = $null
        } else {
            $CheckUser = $JsonParameters.PSobject.Properties["checkUser"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "textCheckOn"))) { #optional property not found
            $TextCheckOn = $null
        } else {
            $TextCheckOn = $JsonParameters.PSobject.Properties["textCheckOn"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "checkPassword"))) { #optional property not found
            $CheckPassword = $null
        } else {
            $CheckPassword = $JsonParameters.PSobject.Properties["checkPassword"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "webTextMatch"))) { #optional property not found
            $WebTextMatch = $null
        } else {
            $WebTextMatch = $JsonParameters.PSobject.Properties["webTextMatch"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "checkPasswordHash"))) { #optional property not found
            $CheckPasswordHash = $null
        } else {
            $CheckPasswordHash = $JsonParameters.PSobject.Properties["checkPasswordHash"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tunnelOn"))) { #optional property not found
            $TunnelOn = $null
        } else {
            $TunnelOn = $JsonParameters.PSobject.Properties["tunnelOn"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sshUser"))) { #optional property not found
            $SshUser = $null
        } else {
            $SshUser = $JsonParameters.PSobject.Properties["sshUser"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sshPassword"))) { #optional property not found
            $SshPassword = $null
        } else {
            $SshPassword = $JsonParameters.PSobject.Properties["sshPassword"].value
        }

        $PSO = [PSCustomObject]@{
            "containerName" = ${ContainerName}
            "externalId" = ${ExternalId}
            "checkUser" = ${CheckUser}
            "textCheckOn" = ${TextCheckOn}
            "checkPassword" = ${CheckPassword}
            "webTextMatch" = ${WebTextMatch}
            "checkPasswordHash" = ${CheckPasswordHash}
            "tunnelOn" = ${TunnelOn}
            "sshHost" = ${SshHost}
            "sshPort" = ${SshPort}
            "sshUser" = ${SshUser}
            "sshPassword" = ${SshPassword}
        }

        return $PSO
    }

}

