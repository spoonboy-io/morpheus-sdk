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

Payload for creating a new tenant

.PARAMETER Name
Name
.PARAMETER Description
Description
.PARAMETER Role
No description available.
.PARAMETER Subdomain
The subdomain. This will be part of the login URL and username for sub tenant users.
.PARAMETER Currency
No description available.
.OUTPUTS

AddTenantRequestAccount<PSCustomObject>
#>

function Initialize-AddTenantRequestAccount {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Role},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Subdomain},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Currency}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddTenantRequestAccount' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "role" = ${Role}
            "subdomain" = ${Subdomain}
            "currency" = ${Currency}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddTenantRequestAccount<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddTenantRequestAccount<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddTenantRequestAccount<PSCustomObject>
#>
function ConvertFrom-JsonToAddTenantRequestAccount {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddTenantRequestAccount' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddTenantRequestAccount
        $AllProperties = ("name", "description", "role", "subdomain", "currency")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "role"))) { #optional property not found
            $Role = $null
        } else {
            $Role = $JsonParameters.PSobject.Properties["role"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "subdomain"))) { #optional property not found
            $Subdomain = $null
        } else {
            $Subdomain = $JsonParameters.PSobject.Properties["subdomain"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "currency"))) { #optional property not found
            $Currency = $null
        } else {
            $Currency = $JsonParameters.PSobject.Properties["currency"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "role" = ${Role}
            "subdomain" = ${Subdomain}
            "currency" = ${Currency}
        }

        return $PSO
    }

}

