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
.PARAMETER Image
Path to display image. Defaults to an internal Morpheus image.
.PARAMETER Type
Blueprint Type
.PARAMETER Helm
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

BlueprintHelmCreateSuccess<PSCustomObject>
#>

function Initialize-BlueprintHelmCreateSuccess {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Image},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("helm")]
        [String]
        ${Type},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Helm},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("private", "public")]
        [String]
        ${Visibility} = "private",
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ResourcePermission},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Owner},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Tenant}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BlueprintHelmCreateSuccess' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "image" = ${Image}
            "type" = ${Type}
            "helm" = ${Helm}
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

Convert from JSON to BlueprintHelmCreateSuccess<PSCustomObject>

.DESCRIPTION

Convert from JSON to BlueprintHelmCreateSuccess<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BlueprintHelmCreateSuccess<PSCustomObject>
#>
function ConvertFrom-JsonToBlueprintHelmCreateSuccess {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BlueprintHelmCreateSuccess' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BlueprintHelmCreateSuccess
        $AllProperties = ("name", "image", "type", "helm", "visibility", "resourcePermission", "owner", "tenant")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "image"))) { #optional property not found
            $Image = $null
        } else {
            $Image = $JsonParameters.PSobject.Properties["image"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) { #optional property not found
            $Type = $null
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "helm"))) { #optional property not found
            $Helm = $null
        } else {
            $Helm = $JsonParameters.PSobject.Properties["helm"].value
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
            "image" = ${Image}
            "type" = ${Type}
            "helm" = ${Helm}
            "visibility" = ${Visibility}
            "resourcePermission" = ${ResourcePermission}
            "owner" = ${Owner}
            "tenant" = ${Tenant}
        }

        return $PSO
    }

}

