#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.1.1
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
Terraform definition in JSON for `configType` `json`
.PARAMETER Tf
Terraform definition for `configType` `tf`
.PARAMETER Git
No description available.
.PARAMETER TfvarSecret
tfvar secret from Morpheus Cypher
.OUTPUTS

BlueprintTerraformCreateTerraform<PSCustomObject>
#>

function Initialize-BlueprintTerraformCreateTerraform {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("tf", "spec", "git", "json")]
        [String]
        ${ConfigType},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Json},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Tf},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Git},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${TfvarSecret}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BlueprintTerraformCreateTerraform' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $ConfigType) {
            throw "invalid value for 'ConfigType', 'ConfigType' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "configType" = ${ConfigType}
            "json" = ${Json}
            "tf" = ${Tf}
            "git" = ${Git}
            "tfvarSecret" = ${TfvarSecret}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BlueprintTerraformCreateTerraform<PSCustomObject>

.DESCRIPTION

Convert from JSON to BlueprintTerraformCreateTerraform<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BlueprintTerraformCreateTerraform<PSCustomObject>
#>
function ConvertFrom-JsonToBlueprintTerraformCreateTerraform {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BlueprintTerraformCreateTerraform' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BlueprintTerraformCreateTerraform
        $AllProperties = ("configType", "json", "tf", "git", "tfvarSecret")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tf"))) { #optional property not found
            $Tf = $null
        } else {
            $Tf = $JsonParameters.PSobject.Properties["tf"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "git"))) { #optional property not found
            $Git = $null
        } else {
            $Git = $JsonParameters.PSobject.Properties["git"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tfvarSecret"))) { #optional property not found
            $TfvarSecret = $null
        } else {
            $TfvarSecret = $JsonParameters.PSobject.Properties["tfvarSecret"].value
        }

        $PSO = [PSCustomObject]@{
            "configType" = ${ConfigType}
            "json" = ${Json}
            "tf" = ${Tf}
            "git" = ${Git}
            "tfvarSecret" = ${TfvarSecret}
        }

        return $PSO
    }

}

