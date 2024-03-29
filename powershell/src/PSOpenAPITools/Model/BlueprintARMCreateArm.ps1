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

.PARAMETER ConfigType
Configuration Type
.PARAMETER Json
ARM Template in JSON
.PARAMETER Yaml
ARM Template in YAML
.PARAMETER Git
No description available.
.PARAMETER OsType
OS Type
.PARAMETER InstallAgent
No description available.
.PARAMETER CloudInitEnabled
No description available.
.OUTPUTS

BlueprintARMCreateArm<PSCustomObject>
#>

function Initialize-BlueprintARMCreateArm {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("json", "yaml", "git")]
        [String]
        ${ConfigType},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Json},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Yaml},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Git},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("linux", "windows")]
        [String]
        ${OsType},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${InstallAgent},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${CloudInitEnabled}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BlueprintARMCreateArm' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $ConfigType) {
            throw "invalid value for 'ConfigType', 'ConfigType' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "configType" = ${ConfigType}
            "json" = ${Json}
            "yaml" = ${Yaml}
            "git" = ${Git}
            "osType" = ${OsType}
            "installAgent" = ${InstallAgent}
            "cloudInitEnabled" = ${CloudInitEnabled}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BlueprintARMCreateArm<PSCustomObject>

.DESCRIPTION

Convert from JSON to BlueprintARMCreateArm<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BlueprintARMCreateArm<PSCustomObject>
#>
function ConvertFrom-JsonToBlueprintARMCreateArm {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BlueprintARMCreateArm' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BlueprintARMCreateArm
        $AllProperties = ("configType", "json", "yaml", "git", "osType", "installAgent", "cloudInitEnabled")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'configType' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "configType"))) {
            throw "Error! JSON cannot be serialized due to the required property 'configType' missing."
        } else {
            $ConfigType = $JsonParameters.PSobject.Properties["configType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "json"))) { #optional property not found
            $Json = $null
        } else {
            $Json = $JsonParameters.PSobject.Properties["json"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "yaml"))) { #optional property not found
            $Yaml = $null
        } else {
            $Yaml = $JsonParameters.PSobject.Properties["yaml"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "git"))) { #optional property not found
            $Git = $null
        } else {
            $Git = $JsonParameters.PSobject.Properties["git"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "osType"))) { #optional property not found
            $OsType = $null
        } else {
            $OsType = $JsonParameters.PSobject.Properties["osType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "installAgent"))) { #optional property not found
            $InstallAgent = $null
        } else {
            $InstallAgent = $JsonParameters.PSobject.Properties["installAgent"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cloudInitEnabled"))) { #optional property not found
            $CloudInitEnabled = $null
        } else {
            $CloudInitEnabled = $JsonParameters.PSobject.Properties["cloudInitEnabled"].value
        }

        $PSO = [PSCustomObject]@{
            "configType" = ${ConfigType}
            "json" = ${Json}
            "yaml" = ${Yaml}
            "git" = ${Git}
            "osType" = ${OsType}
            "installAgent" = ${InstallAgent}
            "cloudInitEnabled" = ${CloudInitEnabled}
        }

        return $PSO
    }

}

