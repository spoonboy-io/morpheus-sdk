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

.PARAMETER Outputs
No description available.
.OUTPUTS

AppStateOutput<PSCustomObject>
#>

function Initialize-AppStateOutput {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Outputs}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AppStateOutput' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "outputs" = ${Outputs}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AppStateOutput<PSCustomObject>

.DESCRIPTION

Convert from JSON to AppStateOutput<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AppStateOutput<PSCustomObject>
#>
function ConvertFrom-JsonToAppStateOutput {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AppStateOutput' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AppStateOutput
        $AllProperties = ("outputs")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "outputs"))) { #optional property not found
            $Outputs = $null
        } else {
            $Outputs = $JsonParameters.PSobject.Properties["outputs"].value
        }

        $PSO = [PSCustomObject]@{
            "outputs" = ${Outputs}
        }

        return $PSO
    }

}

