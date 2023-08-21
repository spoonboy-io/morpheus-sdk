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

.PARAMETER ZoneId
The ID of the Zone (Cloud)
.PARAMETER CustomOptions
No description available.
.OUTPUTS

ApiSecurityGroupsIdLocationsSecurityGroupLocation<PSCustomObject>
#>

function Initialize-ApiSecurityGroupsIdLocationsSecurityGroupLocation {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [Int64]
        ${ZoneId},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${CustomOptions}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiSecurityGroupsIdLocationsSecurityGroupLocation' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if (!$ZoneId) {
            throw "invalid value for 'ZoneId', 'ZoneId' cannot be null."
        }

        if (!$CustomOptions) {
            throw "invalid value for 'CustomOptions', 'CustomOptions' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "zoneId" = ${ZoneId}
            "customOptions" = ${CustomOptions}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiSecurityGroupsIdLocationsSecurityGroupLocation<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiSecurityGroupsIdLocationsSecurityGroupLocation<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiSecurityGroupsIdLocationsSecurityGroupLocation<PSCustomObject>
#>
function ConvertFrom-JsonToApiSecurityGroupsIdLocationsSecurityGroupLocation {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiSecurityGroupsIdLocationsSecurityGroupLocation' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiSecurityGroupsIdLocationsSecurityGroupLocation
        $AllProperties = ("zoneId", "customOptions")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property `zoneId` missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "zoneId"))) {
            throw "Error! JSON cannot be serialized due to the required property `zoneId` missing."
        } else {
            $ZoneId = $JsonParameters.PSobject.Properties["zoneId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "customOptions"))) {
            throw "Error! JSON cannot be serialized due to the required property `customOptions` missing."
        } else {
            $CustomOptions = $JsonParameters.PSobject.Properties["customOptions"].value
        }

        $PSO = [PSCustomObject]@{
            "zoneId" = ${ZoneId}
            "customOptions" = ${CustomOptions}
        }

        return $PSO
    }

}
