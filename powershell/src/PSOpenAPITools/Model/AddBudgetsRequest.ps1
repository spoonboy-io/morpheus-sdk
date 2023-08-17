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

.PARAMETER Budget
No description available.
.OUTPUTS

AddBudgetsRequest<PSCustomObject>
#>

function Initialize-AddBudgetsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Budget}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddBudgetsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Budget) {
            throw "invalid value for 'Budget', 'Budget' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "budget" = ${Budget}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddBudgetsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddBudgetsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddBudgetsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddBudgetsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddBudgetsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddBudgetsRequest
        $AllProperties = ("budget")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'budget' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "budget"))) {
            throw "Error! JSON cannot be serialized due to the required property 'budget' missing."
        } else {
            $Budget = $JsonParameters.PSobject.Properties["budget"].value
        }

        $PSO = [PSCustomObject]@{
            "budget" = ${Budget}
        }

        return $PSO
    }

}

