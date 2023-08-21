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

.PARAMETER Type
Credential Type Code
.PARAMETER Name
A unique name scoped to your account for the credential
.PARAMETER Description
Optional Description
.PARAMETER Enabled
Credential enabled
.PARAMETER Integration
No description available.
.PARAMETER Username
Username
.PARAMETER Password
Password
.OUTPUTS

CredentialUsernamePasswordConfig<PSCustomObject>
#>

function Initialize-CredentialUsernamePasswordConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("username-password")]
        [String]
        ${Type},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled} = $true,
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Integration},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Username},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Password}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CredentialUsernamePasswordConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if (!$Type) {
            throw "invalid value for 'Type', 'Type' cannot be null."
        }

        if (!$Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }

        if (!$Username) {
            throw "invalid value for 'Username', 'Username' cannot be null."
        }

        if (!$Password) {
            throw "invalid value for 'Password', 'Password' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "type" = ${Type}
            "name" = ${Name}
            "description" = ${Description}
            "enabled" = ${Enabled}
            "integration" = ${Integration}
            "username" = ${Username}
            "password" = ${Password}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CredentialUsernamePasswordConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to CredentialUsernamePasswordConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CredentialUsernamePasswordConfig<PSCustomObject>
#>
function ConvertFrom-JsonToCredentialUsernamePasswordConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CredentialUsernamePasswordConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CredentialUsernamePasswordConfig
        $AllProperties = ("type", "name", "description", "enabled", "integration", "username", "password")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property `type` missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) {
            throw "Error! JSON cannot be serialized due to the required property `type` missing."
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property `name` missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "username"))) {
            throw "Error! JSON cannot be serialized due to the required property `username` missing."
        } else {
            $Username = $JsonParameters.PSobject.Properties["username"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "password"))) {
            throw "Error! JSON cannot be serialized due to the required property `password` missing."
        } else {
            $Password = $JsonParameters.PSobject.Properties["password"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "integration"))) { #optional property not found
            $Integration = $null
        } else {
            $Integration = $JsonParameters.PSobject.Properties["integration"].value
        }

        $PSO = [PSCustomObject]@{
            "type" = ${Type}
            "name" = ${Name}
            "description" = ${Description}
            "enabled" = ${Enabled}
            "integration" = ${Integration}
            "username" = ${Username}
            "password" = ${Password}
        }

        return $PSO
    }

}

