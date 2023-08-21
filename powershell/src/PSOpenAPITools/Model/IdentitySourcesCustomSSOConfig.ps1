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
.PARAMETER Name
No description available.
.PARAMETER Description
No description available.
.PARAMETER Code
No description available.
.PARAMETER Type
No description available.
.PARAMETER Active
No description available.
.PARAMETER Deleted
No description available.
.PARAMETER AutoSyncOnLogin
No description available.
.PARAMETER ExternalLogin
No description available.
.PARAMETER AllowCustomMappings
No description available.
.PARAMETER ManualRoleAssignment
No description available.
.PARAMETER Account
No description available.
.PARAMETER DefaultAccountRole
No description available.
.PARAMETER Config
No description available.
.PARAMETER RoleMappings
No description available.
.PARAMETER Subdomain
No description available.
.PARAMETER LoginURL
No description available.
.PARAMETER ProviderSettings
No description available.
.PARAMETER DateCreated
No description available.
.PARAMETER LastUpdated
No description available.
.OUTPUTS

IdentitySourcesCustomSSOConfig<PSCustomObject>
#>

function Initialize-IdentitySourcesCustomSSOConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Code},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Type},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Active},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Deleted},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${AutoSyncOnLogin},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${ExternalLogin},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${AllowCustomMappings},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${ManualRoleAssignment},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Account},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${DefaultAccountRole},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Config},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${RoleMappings},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Subdomain},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LoginURL},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ProviderSettings},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => IdentitySourcesCustomSSOConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "description" = ${Description}
            "code" = ${Code}
            "type" = ${Type}
            "active" = ${Active}
            "deleted" = ${Deleted}
            "autoSyncOnLogin" = ${AutoSyncOnLogin}
            "externalLogin" = ${ExternalLogin}
            "allowCustomMappings" = ${AllowCustomMappings}
            "manualRoleAssignment" = ${ManualRoleAssignment}
            "account" = ${Account}
            "defaultAccountRole" = ${DefaultAccountRole}
            "config" = ${Config}
            "roleMappings" = ${RoleMappings}
            "subdomain" = ${Subdomain}
            "loginURL" = ${LoginURL}
            "providerSettings" = ${ProviderSettings}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to IdentitySourcesCustomSSOConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to IdentitySourcesCustomSSOConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

IdentitySourcesCustomSSOConfig<PSCustomObject>
#>
function ConvertFrom-JsonToIdentitySourcesCustomSSOConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => IdentitySourcesCustomSSOConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in IdentitySourcesCustomSSOConfig
        $AllProperties = ("id", "name", "description", "code", "type", "active", "deleted", "autoSyncOnLogin", "externalLogin", "allowCustomMappings", "manualRoleAssignment", "account", "defaultAccountRole", "config", "roleMappings", "subdomain", "loginURL", "providerSettings", "dateCreated", "lastUpdated")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "code"))) { #optional property not found
            $Code = $null
        } else {
            $Code = $JsonParameters.PSobject.Properties["code"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) { #optional property not found
            $Type = $null
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "active"))) { #optional property not found
            $Active = $null
        } else {
            $Active = $JsonParameters.PSobject.Properties["active"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "deleted"))) { #optional property not found
            $Deleted = $null
        } else {
            $Deleted = $JsonParameters.PSobject.Properties["deleted"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "autoSyncOnLogin"))) { #optional property not found
            $AutoSyncOnLogin = $null
        } else {
            $AutoSyncOnLogin = $JsonParameters.PSobject.Properties["autoSyncOnLogin"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalLogin"))) { #optional property not found
            $ExternalLogin = $null
        } else {
            $ExternalLogin = $JsonParameters.PSobject.Properties["externalLogin"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "allowCustomMappings"))) { #optional property not found
            $AllowCustomMappings = $null
        } else {
            $AllowCustomMappings = $JsonParameters.PSobject.Properties["allowCustomMappings"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "manualRoleAssignment"))) { #optional property not found
            $ManualRoleAssignment = $null
        } else {
            $ManualRoleAssignment = $JsonParameters.PSobject.Properties["manualRoleAssignment"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "account"))) { #optional property not found
            $Account = $null
        } else {
            $Account = $JsonParameters.PSobject.Properties["account"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "defaultAccountRole"))) { #optional property not found
            $DefaultAccountRole = $null
        } else {
            $DefaultAccountRole = $JsonParameters.PSobject.Properties["defaultAccountRole"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "config"))) { #optional property not found
            $Config = $null
        } else {
            $Config = $JsonParameters.PSobject.Properties["config"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "roleMappings"))) { #optional property not found
            $RoleMappings = $null
        } else {
            $RoleMappings = $JsonParameters.PSobject.Properties["roleMappings"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "subdomain"))) { #optional property not found
            $Subdomain = $null
        } else {
            $Subdomain = $JsonParameters.PSobject.Properties["subdomain"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "loginURL"))) { #optional property not found
            $LoginURL = $null
        } else {
            $LoginURL = $JsonParameters.PSobject.Properties["loginURL"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "providerSettings"))) { #optional property not found
            $ProviderSettings = $null
        } else {
            $ProviderSettings = $JsonParameters.PSobject.Properties["providerSettings"].value
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

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "description" = ${Description}
            "code" = ${Code}
            "type" = ${Type}
            "active" = ${Active}
            "deleted" = ${Deleted}
            "autoSyncOnLogin" = ${AutoSyncOnLogin}
            "externalLogin" = ${ExternalLogin}
            "allowCustomMappings" = ${AllowCustomMappings}
            "manualRoleAssignment" = ${ManualRoleAssignment}
            "account" = ${Account}
            "defaultAccountRole" = ${DefaultAccountRole}
            "config" = ${Config}
            "roleMappings" = ${RoleMappings}
            "subdomain" = ${Subdomain}
            "loginURL" = ${LoginURL}
            "providerSettings" = ${ProviderSettings}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
        }

        return $PSO
    }

}

