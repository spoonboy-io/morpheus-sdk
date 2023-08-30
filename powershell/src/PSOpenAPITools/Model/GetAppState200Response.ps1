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

.PARAMETER Workloads
No description available.
.PARAMETER IacDrift
No description available.
.PARAMETER PlanResources
No description available.
.PARAMETER Specs
No description available.
.PARAMETER StateData
No description available.
.PARAMETER PlanData
No description available.
.PARAMETER VarInput
No description available.
.PARAMETER Output
No description available.
.PARAMETER Success
No description available.
.OUTPUTS

GetAppState200Response<PSCustomObject>
#>

function Initialize-GetAppState200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Workloads},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${IacDrift},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${PlanResources},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Specs},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${StateData},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${PlanData},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${VarInput},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Output},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Success}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetAppState200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "workloads" = ${Workloads}
            "iacDrift" = ${IacDrift}
            "planResources" = ${PlanResources}
            "specs" = ${Specs}
            "stateData" = ${StateData}
            "planData" = ${PlanData}
            "input" = ${VarInput}
            "output" = ${Output}
            "success" = ${Success}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetAppState200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetAppState200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetAppState200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetAppState200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetAppState200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetAppState200Response
        $AllProperties = ("workloads", "iacDrift", "planResources", "specs", "stateData", "planData", "input", "output", "success")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "workloads"))) { #optional property not found
            $Workloads = $null
        } else {
            $Workloads = $JsonParameters.PSobject.Properties["workloads"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "iacDrift"))) { #optional property not found
            $IacDrift = $null
        } else {
            $IacDrift = $JsonParameters.PSobject.Properties["iacDrift"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "planResources"))) { #optional property not found
            $PlanResources = $null
        } else {
            $PlanResources = $JsonParameters.PSobject.Properties["planResources"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "specs"))) { #optional property not found
            $Specs = $null
        } else {
            $Specs = $JsonParameters.PSobject.Properties["specs"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "stateData"))) { #optional property not found
            $StateData = $null
        } else {
            $StateData = $JsonParameters.PSobject.Properties["stateData"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "planData"))) { #optional property not found
            $PlanData = $null
        } else {
            $PlanData = $JsonParameters.PSobject.Properties["planData"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "input"))) { #optional property not found
            $VarInput = $null
        } else {
            $VarInput = $JsonParameters.PSobject.Properties["input"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "output"))) { #optional property not found
            $Output = $null
        } else {
            $Output = $JsonParameters.PSobject.Properties["output"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "success"))) { #optional property not found
            $Success = $null
        } else {
            $Success = $JsonParameters.PSobject.Properties["success"].value
        }

        $PSO = [PSCustomObject]@{
            "workloads" = ${Workloads}
            "iacDrift" = ${IacDrift}
            "planResources" = ${PlanResources}
            "specs" = ${Specs}
            "stateData" = ${StateData}
            "planData" = ${PlanData}
            "input" = ${VarInput}
            "output" = ${Output}
            "success" = ${Success}
        }

        return $PSO
    }

}

