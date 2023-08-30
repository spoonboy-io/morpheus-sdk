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

.PARAMETER TemplateParameter
No description available.
.OUTPUTS

GetPrepareApplyInstance200ResponseAllOfData<PSCustomObject>
#>

function Initialize-GetPrepareApplyInstance200ResponseAllOfData {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${TemplateParameter}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetPrepareApplyInstance200ResponseAllOfData' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "templateParameter" = ${TemplateParameter}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetPrepareApplyInstance200ResponseAllOfData<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetPrepareApplyInstance200ResponseAllOfData<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetPrepareApplyInstance200ResponseAllOfData<PSCustomObject>
#>
function ConvertFrom-JsonToGetPrepareApplyInstance200ResponseAllOfData {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetPrepareApplyInstance200ResponseAllOfData' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetPrepareApplyInstance200ResponseAllOfData
        $AllProperties = ("templateParameter")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "templateParameter"))) { #optional property not found
            $TemplateParameter = $null
        } else {
            $TemplateParameter = $JsonParameters.PSobject.Properties["templateParameter"].value
        }

        $PSO = [PSCustomObject]@{
            "templateParameter" = ${TemplateParameter}
        }

        return $PSO
    }

}

