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
.PARAMETER Labels
Array of label strings, can be used for filtering.
.PARAMETER Kubernetes
No description available.
.PARAMETER Config
No description available.
.OUTPUTS

BlueprintKubernetesCreate<PSCustomObject>
#>

function Initialize-BlueprintKubernetesCreate {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Image},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("kubernetes")]
        [String]
        ${Type},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String[]]
        ${Labels},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Kubernetes},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Config}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BlueprintKubernetesCreate' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if (!$Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }

        if (!$Type) {
            throw "invalid value for 'Type', 'Type' cannot be null."
        }

        if (!$Kubernetes) {
            throw "invalid value for 'Kubernetes', 'Kubernetes' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "image" = ${Image}
            "type" = ${Type}
            "labels" = ${Labels}
            "kubernetes" = ${Kubernetes}
            "config" = ${Config}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BlueprintKubernetesCreate<PSCustomObject>

.DESCRIPTION

Convert from JSON to BlueprintKubernetesCreate<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BlueprintKubernetesCreate<PSCustomObject>
#>
function ConvertFrom-JsonToBlueprintKubernetesCreate {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BlueprintKubernetesCreate' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BlueprintKubernetesCreate
        $AllProperties = ("name", "image", "type", "labels", "kubernetes", "config")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property `name` missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property `name` missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) {
            throw "Error! JSON cannot be serialized due to the required property `type` missing."
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "kubernetes"))) {
            throw "Error! JSON cannot be serialized due to the required property `kubernetes` missing."
        } else {
            $Kubernetes = $JsonParameters.PSobject.Properties["kubernetes"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "image"))) { #optional property not found
            $Image = $null
        } else {
            $Image = $JsonParameters.PSobject.Properties["image"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "labels"))) { #optional property not found
            $Labels = $null
        } else {
            $Labels = $JsonParameters.PSobject.Properties["labels"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "config"))) { #optional property not found
            $Config = $null
        } else {
            $Config = $JsonParameters.PSobject.Properties["config"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "image" = ${Image}
            "type" = ${Type}
            "labels" = ${Labels}
            "kubernetes" = ${Kubernetes}
            "config" = ${Config}
        }

        return $PSO
    }

}

