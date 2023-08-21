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
A name for the blueprint
.PARAMETER Type
Blueprint Type
.PARAMETER Config
No description available.
.PARAMETER Visibility
Private or Public Access
.PARAMETER ResourcePermission
Resource Permission Block
.PARAMETER Owner
Owner
.PARAMETER Tenant
Tenant
.OUTPUTS

BlueprintMorpheusCreateSuccess<PSCustomObject>
#>

function Initialize-BlueprintMorpheusCreateSuccess {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("morpheus")]
        [String]
        ${Type},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Config},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("private", "public")]
        [String]
        ${Visibility} = "private",
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ResourcePermission},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Owner},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Tenant}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BlueprintMorpheusCreateSuccess' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "type" = ${Type}
            "config" = ${Config}
            "visibility" = ${Visibility}
            "resourcePermission" = ${ResourcePermission}
            "owner" = ${Owner}
            "tenant" = ${Tenant}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BlueprintMorpheusCreateSuccess<PSCustomObject>

.DESCRIPTION

Convert from JSON to BlueprintMorpheusCreateSuccess<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BlueprintMorpheusCreateSuccess<PSCustomObject>
#>
function ConvertFrom-JsonToBlueprintMorpheusCreateSuccess {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BlueprintMorpheusCreateSuccess' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BlueprintMorpheusCreateSuccess
        $AllProperties = ("name", "type", "config", "visibility", "resourcePermission", "owner", "tenant")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) { #optional property not found
            $Type = $null
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "config"))) { #optional property not found
            $Config = $null
        } else {
            $Config = $JsonParameters.PSobject.Properties["config"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "visibility"))) { #optional property not found
            $Visibility = $null
        } else {
            $Visibility = $JsonParameters.PSobject.Properties["visibility"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "resourcePermission"))) { #optional property not found
            $ResourcePermission = $null
        } else {
            $ResourcePermission = $JsonParameters.PSobject.Properties["resourcePermission"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "owner"))) { #optional property not found
            $Owner = $null
        } else {
            $Owner = $JsonParameters.PSobject.Properties["owner"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tenant"))) { #optional property not found
            $Tenant = $null
        } else {
            $Tenant = $JsonParameters.PSobject.Properties["tenant"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "type" = ${Type}
            "config" = ${Config}
            "visibility" = ${Visibility}
            "resourcePermission" = ${ResourcePermission}
            "owner" = ${Owner}
            "tenant" = ${Tenant}
        }

        return $PSO
    }

}

