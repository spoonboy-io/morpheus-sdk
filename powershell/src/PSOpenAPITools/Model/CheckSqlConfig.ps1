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

.PARAMETER DbHost
Hostname or IP address of the database
.PARAMETER DbPort
Database Port (defaults to default port of DB type selected)
.PARAMETER DbUser
Database username
.PARAMETER DbPassword
Database password, (all check data is encrypted inside the database)
.PARAMETER DbPasswordHash
Database password hash
.PARAMETER DbName
Database name you would like to connect to
.PARAMETER DbQuery
Query to test
.PARAMETER CheckOperator
Can be set to `lt` (less than), `gt` (greater than), `equal` (Equal to) for comparison
.PARAMETER CheckResult
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
No description available.
.PARAMETER SshHost
No description available.
.PARAMETER SshPort
No description available.
.PARAMETER SshUser
No description available.
.PARAMETER SshPassword
Password for user, if not using key based authentication
.OUTPUTS

CheckSqlConfig<PSCustomObject>
#>

function Initialize-CheckSqlConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DbHost},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DbPort},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DbUser},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DbPassword},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DbPasswordHash},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DbName},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DbQuery},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("lt", "gt", "equal")]
        [String]
        ${CheckOperator},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CheckResult},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CheckUser},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${TextCheckOn},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CheckPassword},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${WebTextMatch},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CheckPasswordHash},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${TunnelOn},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SshHost},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${SshPort},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SshUser},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SshPassword}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CheckSqlConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $DbHost) {
            throw "invalid value for 'DbHost', 'DbHost' cannot be null."
        }

        if ($null -eq $DbPort) {
            throw "invalid value for 'DbPort', 'DbPort' cannot be null."
        }

        if ($null -eq $DbUser) {
            throw "invalid value for 'DbUser', 'DbUser' cannot be null."
        }

        if ($null -eq $DbPassword) {
            throw "invalid value for 'DbPassword', 'DbPassword' cannot be null."
        }

        if ($null -eq $DbName) {
            throw "invalid value for 'DbName', 'DbName' cannot be null."
        }

        if ($null -eq $DbQuery) {
            throw "invalid value for 'DbQuery', 'DbQuery' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "dbHost" = ${DbHost}
            "dbPort" = ${DbPort}
            "dbUser" = ${DbUser}
            "dbPassword" = ${DbPassword}
            "dbPasswordHash" = ${DbPasswordHash}
            "dbName" = ${DbName}
            "dbQuery" = ${DbQuery}
            "checkOperator" = ${CheckOperator}
            "checkResult" = ${CheckResult}
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

Convert from JSON to CheckSqlConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to CheckSqlConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CheckSqlConfig<PSCustomObject>
#>
function ConvertFrom-JsonToCheckSqlConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CheckSqlConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CheckSqlConfig
        $AllProperties = ("dbHost", "dbPort", "dbUser", "dbPassword", "dbPasswordHash", "dbName", "dbQuery", "checkOperator", "checkResult", "checkUser", "textCheckOn", "checkPassword", "webTextMatch", "checkPasswordHash", "tunnelOn", "sshHost", "sshPort", "sshUser", "sshPassword")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'dbHost' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dbHost"))) {
            throw "Error! JSON cannot be serialized due to the required property 'dbHost' missing."
        } else {
            $DbHost = $JsonParameters.PSobject.Properties["dbHost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dbPort"))) {
            throw "Error! JSON cannot be serialized due to the required property 'dbPort' missing."
        } else {
            $DbPort = $JsonParameters.PSobject.Properties["dbPort"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dbUser"))) {
            throw "Error! JSON cannot be serialized due to the required property 'dbUser' missing."
        } else {
            $DbUser = $JsonParameters.PSobject.Properties["dbUser"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dbPassword"))) {
            throw "Error! JSON cannot be serialized due to the required property 'dbPassword' missing."
        } else {
            $DbPassword = $JsonParameters.PSobject.Properties["dbPassword"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dbName"))) {
            throw "Error! JSON cannot be serialized due to the required property 'dbName' missing."
        } else {
            $DbName = $JsonParameters.PSobject.Properties["dbName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dbQuery"))) {
            throw "Error! JSON cannot be serialized due to the required property 'dbQuery' missing."
        } else {
            $DbQuery = $JsonParameters.PSobject.Properties["dbQuery"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dbPasswordHash"))) { #optional property not found
            $DbPasswordHash = $null
        } else {
            $DbPasswordHash = $JsonParameters.PSobject.Properties["dbPasswordHash"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "checkOperator"))) { #optional property not found
            $CheckOperator = $null
        } else {
            $CheckOperator = $JsonParameters.PSobject.Properties["checkOperator"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "checkResult"))) { #optional property not found
            $CheckResult = $null
        } else {
            $CheckResult = $JsonParameters.PSobject.Properties["checkResult"].value
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
            "dbHost" = ${DbHost}
            "dbPort" = ${DbPort}
            "dbUser" = ${DbUser}
            "dbPassword" = ${DbPassword}
            "dbPasswordHash" = ${DbPasswordHash}
            "dbName" = ${DbName}
            "dbQuery" = ${DbQuery}
            "checkOperator" = ${CheckOperator}
            "checkResult" = ${CheckResult}
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

