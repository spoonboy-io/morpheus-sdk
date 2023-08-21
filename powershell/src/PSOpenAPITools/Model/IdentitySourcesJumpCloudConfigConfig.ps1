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

.PARAMETER OrganizationId
No description available.
.PARAMETER BindingUsername
No description available.
.PARAMETER BindingPassword
No description available.
.PARAMETER RequiredRole
No description available.
.PARAMETER BindingPasswordHash
No description available.
.OUTPUTS

IdentitySourcesJumpCloudConfigConfig<PSCustomObject>
#>

function Initialize-IdentitySourcesJumpCloudConfigConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${OrganizationId},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${BindingUsername},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${BindingPassword},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RequiredRole},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${BindingPasswordHash}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => IdentitySourcesJumpCloudConfigConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "organizationId" = ${OrganizationId}
            "bindingUsername" = ${BindingUsername}
            "bindingPassword" = ${BindingPassword}
            "requiredRole" = ${RequiredRole}
            "bindingPasswordHash" = ${BindingPasswordHash}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to IdentitySourcesJumpCloudConfigConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to IdentitySourcesJumpCloudConfigConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

IdentitySourcesJumpCloudConfigConfig<PSCustomObject>
#>
function ConvertFrom-JsonToIdentitySourcesJumpCloudConfigConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => IdentitySourcesJumpCloudConfigConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in IdentitySourcesJumpCloudConfigConfig
        $AllProperties = ("organizationId", "bindingUsername", "bindingPassword", "requiredRole", "bindingPasswordHash")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "organizationId"))) { #optional property not found
            $OrganizationId = $null
        } else {
            $OrganizationId = $JsonParameters.PSobject.Properties["organizationId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "bindingUsername"))) { #optional property not found
            $BindingUsername = $null
        } else {
            $BindingUsername = $JsonParameters.PSobject.Properties["bindingUsername"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "bindingPassword"))) { #optional property not found
            $BindingPassword = $null
        } else {
            $BindingPassword = $JsonParameters.PSobject.Properties["bindingPassword"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "requiredRole"))) { #optional property not found
            $RequiredRole = $null
        } else {
            $RequiredRole = $JsonParameters.PSobject.Properties["requiredRole"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "bindingPasswordHash"))) { #optional property not found
            $BindingPasswordHash = $null
        } else {
            $BindingPasswordHash = $JsonParameters.PSobject.Properties["bindingPasswordHash"].value
        }

        $PSO = [PSCustomObject]@{
            "organizationId" = ${OrganizationId}
            "bindingUsername" = ${BindingUsername}
            "bindingPassword" = ${BindingPassword}
            "requiredRole" = ${RequiredRole}
            "bindingPasswordHash" = ${BindingPasswordHash}
        }

        return $PSO
    }

}
