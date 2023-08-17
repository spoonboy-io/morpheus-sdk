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

.PARAMETER Image
No description available.
.PARAMETER Name
No description available.
.PARAMETER AutoValidate
No description available.
.PARAMETER Terraform
No description available.
.PARAMETER Type
No description available.
.PARAMETER Config
No description available.
.PARAMETER BlueprintName
No description available.
.PARAMETER Description
No description available.
.PARAMETER TemplateId
No description available.
.PARAMETER BlueprintId
No description available.
.PARAMETER Group
No description available.
.OUTPUTS

AppPrepareApplyData<PSCustomObject>
#>

function Initialize-AppPrepareApplyData {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Image},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${AutoValidate},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Terraform},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Type},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Config},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${BlueprintName},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${TemplateId},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${BlueprintId},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Group}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AppPrepareApplyData' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "image" = ${Image}
            "name" = ${Name}
            "autoValidate" = ${AutoValidate}
            "terraform" = ${Terraform}
            "type" = ${Type}
            "config" = ${Config}
            "blueprintName" = ${BlueprintName}
            "description" = ${Description}
            "templateId" = ${TemplateId}
            "blueprintId" = ${BlueprintId}
            "group" = ${Group}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AppPrepareApplyData<PSCustomObject>

.DESCRIPTION

Convert from JSON to AppPrepareApplyData<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AppPrepareApplyData<PSCustomObject>
#>
function ConvertFrom-JsonToAppPrepareApplyData {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AppPrepareApplyData' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AppPrepareApplyData
        $AllProperties = ("image", "name", "autoValidate", "terraform", "type", "config", "blueprintName", "description", "templateId", "blueprintId", "group")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "image"))) { #optional property not found
            $Image = $null
        } else {
            $Image = $JsonParameters.PSobject.Properties["image"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "autoValidate"))) { #optional property not found
            $AutoValidate = $null
        } else {
            $AutoValidate = $JsonParameters.PSobject.Properties["autoValidate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "terraform"))) { #optional property not found
            $Terraform = $null
        } else {
            $Terraform = $JsonParameters.PSobject.Properties["terraform"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "blueprintName"))) { #optional property not found
            $BlueprintName = $null
        } else {
            $BlueprintName = $JsonParameters.PSobject.Properties["blueprintName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "templateId"))) { #optional property not found
            $TemplateId = $null
        } else {
            $TemplateId = $JsonParameters.PSobject.Properties["templateId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "blueprintId"))) { #optional property not found
            $BlueprintId = $null
        } else {
            $BlueprintId = $JsonParameters.PSobject.Properties["blueprintId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "group"))) { #optional property not found
            $Group = $null
        } else {
            $Group = $JsonParameters.PSobject.Properties["group"].value
        }

        $PSO = [PSCustomObject]@{
            "image" = ${Image}
            "name" = ${Name}
            "autoValidate" = ${AutoValidate}
            "terraform" = ${Terraform}
            "type" = ${Type}
            "config" = ${Config}
            "blueprintName" = ${BlueprintName}
            "description" = ${Description}
            "templateId" = ${TemplateId}
            "blueprintId" = ${BlueprintId}
            "group" = ${Group}
        }

        return $PSO
    }

}

