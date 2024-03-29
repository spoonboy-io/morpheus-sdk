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

.PARAMETER Role
No description available.
.PARAMETER FeaturePermissions
No description available.
.PARAMETER GlobalSiteAccess
No description available.
.PARAMETER Sites
No description available.
.PARAMETER GlobalZoneAccess
No description available.
.PARAMETER Zones
No description available.
.PARAMETER GlobalInstanceTypeAccess
No description available.
.PARAMETER InstanceTypePermissions
No description available.
.PARAMETER GlobalAppTemplateAccess
No description available.
.PARAMETER AppTemplatePermissions
No description available.
.PARAMETER GlobalCatalogItemTypeAccess
No description available.
.PARAMETER CatalogItemTypePermissions
No description available.
.PARAMETER GlobalPersonaAccess
No description available.
.PARAMETER PersonaPermissions
No description available.
.PARAMETER GlobalVdiPoolAccess
No description available.
.PARAMETER VdiPoolPermissions
No description available.
.PARAMETER GlobalReportTypeAccess
No description available.
.PARAMETER ReportTypePermissions
No description available.
.PARAMETER GlobalTaskAccess
No description available.
.PARAMETER TaskPermissions
No description available.
.PARAMETER GlobalTaskSetAccess
No description available.
.PARAMETER TaskSetPermissions
No description available.
.PARAMETER Success
No description available.
.OUTPUTS

AddRoles200Response<PSCustomObject>
#>

function Initialize-AddRoles200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Role},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${FeaturePermissions},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GlobalSiteAccess},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Sites},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GlobalZoneAccess},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Zones},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GlobalInstanceTypeAccess},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${InstanceTypePermissions},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GlobalAppTemplateAccess},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${AppTemplatePermissions},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GlobalCatalogItemTypeAccess},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${CatalogItemTypePermissions},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GlobalPersonaAccess},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${PersonaPermissions},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GlobalVdiPoolAccess},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${VdiPoolPermissions},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GlobalReportTypeAccess},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${ReportTypePermissions},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GlobalTaskAccess},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${TaskPermissions},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GlobalTaskSetAccess},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${TaskSetPermissions},
        [Parameter(Position = 22, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Success}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddRoles200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "role" = ${Role}
            "featurePermissions" = ${FeaturePermissions}
            "globalSiteAccess" = ${GlobalSiteAccess}
            "sites" = ${Sites}
            "globalZoneAccess" = ${GlobalZoneAccess}
            "zones" = ${Zones}
            "globalInstanceTypeAccess" = ${GlobalInstanceTypeAccess}
            "instanceTypePermissions" = ${InstanceTypePermissions}
            "globalAppTemplateAccess" = ${GlobalAppTemplateAccess}
            "appTemplatePermissions" = ${AppTemplatePermissions}
            "globalCatalogItemTypeAccess" = ${GlobalCatalogItemTypeAccess}
            "catalogItemTypePermissions" = ${CatalogItemTypePermissions}
            "globalPersonaAccess" = ${GlobalPersonaAccess}
            "personaPermissions" = ${PersonaPermissions}
            "globalVdiPoolAccess" = ${GlobalVdiPoolAccess}
            "vdiPoolPermissions" = ${VdiPoolPermissions}
            "globalReportTypeAccess" = ${GlobalReportTypeAccess}
            "reportTypePermissions" = ${ReportTypePermissions}
            "globalTaskAccess" = ${GlobalTaskAccess}
            "taskPermissions" = ${TaskPermissions}
            "globalTaskSetAccess" = ${GlobalTaskSetAccess}
            "taskSetPermissions" = ${TaskSetPermissions}
            "success" = ${Success}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddRoles200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddRoles200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddRoles200Response<PSCustomObject>
#>
function ConvertFrom-JsonToAddRoles200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddRoles200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddRoles200Response
        $AllProperties = ("role", "featurePermissions", "globalSiteAccess", "sites", "globalZoneAccess", "zones", "globalInstanceTypeAccess", "instanceTypePermissions", "globalAppTemplateAccess", "appTemplatePermissions", "globalCatalogItemTypeAccess", "catalogItemTypePermissions", "globalPersonaAccess", "personaPermissions", "globalVdiPoolAccess", "vdiPoolPermissions", "globalReportTypeAccess", "reportTypePermissions", "globalTaskAccess", "taskPermissions", "globalTaskSetAccess", "taskSetPermissions", "success")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "role"))) { #optional property not found
            $Role = $null
        } else {
            $Role = $JsonParameters.PSobject.Properties["role"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "featurePermissions"))) { #optional property not found
            $FeaturePermissions = $null
        } else {
            $FeaturePermissions = $JsonParameters.PSobject.Properties["featurePermissions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalSiteAccess"))) { #optional property not found
            $GlobalSiteAccess = $null
        } else {
            $GlobalSiteAccess = $JsonParameters.PSobject.Properties["globalSiteAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sites"))) { #optional property not found
            $Sites = $null
        } else {
            $Sites = $JsonParameters.PSobject.Properties["sites"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalZoneAccess"))) { #optional property not found
            $GlobalZoneAccess = $null
        } else {
            $GlobalZoneAccess = $JsonParameters.PSobject.Properties["globalZoneAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "zones"))) { #optional property not found
            $Zones = $null
        } else {
            $Zones = $JsonParameters.PSobject.Properties["zones"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalInstanceTypeAccess"))) { #optional property not found
            $GlobalInstanceTypeAccess = $null
        } else {
            $GlobalInstanceTypeAccess = $JsonParameters.PSobject.Properties["globalInstanceTypeAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceTypePermissions"))) { #optional property not found
            $InstanceTypePermissions = $null
        } else {
            $InstanceTypePermissions = $JsonParameters.PSobject.Properties["instanceTypePermissions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalAppTemplateAccess"))) { #optional property not found
            $GlobalAppTemplateAccess = $null
        } else {
            $GlobalAppTemplateAccess = $JsonParameters.PSobject.Properties["globalAppTemplateAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "appTemplatePermissions"))) { #optional property not found
            $AppTemplatePermissions = $null
        } else {
            $AppTemplatePermissions = $JsonParameters.PSobject.Properties["appTemplatePermissions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalCatalogItemTypeAccess"))) { #optional property not found
            $GlobalCatalogItemTypeAccess = $null
        } else {
            $GlobalCatalogItemTypeAccess = $JsonParameters.PSobject.Properties["globalCatalogItemTypeAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "catalogItemTypePermissions"))) { #optional property not found
            $CatalogItemTypePermissions = $null
        } else {
            $CatalogItemTypePermissions = $JsonParameters.PSobject.Properties["catalogItemTypePermissions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalPersonaAccess"))) { #optional property not found
            $GlobalPersonaAccess = $null
        } else {
            $GlobalPersonaAccess = $JsonParameters.PSobject.Properties["globalPersonaAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "personaPermissions"))) { #optional property not found
            $PersonaPermissions = $null
        } else {
            $PersonaPermissions = $JsonParameters.PSobject.Properties["personaPermissions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalVdiPoolAccess"))) { #optional property not found
            $GlobalVdiPoolAccess = $null
        } else {
            $GlobalVdiPoolAccess = $JsonParameters.PSobject.Properties["globalVdiPoolAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "vdiPoolPermissions"))) { #optional property not found
            $VdiPoolPermissions = $null
        } else {
            $VdiPoolPermissions = $JsonParameters.PSobject.Properties["vdiPoolPermissions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalReportTypeAccess"))) { #optional property not found
            $GlobalReportTypeAccess = $null
        } else {
            $GlobalReportTypeAccess = $JsonParameters.PSobject.Properties["globalReportTypeAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "reportTypePermissions"))) { #optional property not found
            $ReportTypePermissions = $null
        } else {
            $ReportTypePermissions = $JsonParameters.PSobject.Properties["reportTypePermissions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalTaskAccess"))) { #optional property not found
            $GlobalTaskAccess = $null
        } else {
            $GlobalTaskAccess = $JsonParameters.PSobject.Properties["globalTaskAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "taskPermissions"))) { #optional property not found
            $TaskPermissions = $null
        } else {
            $TaskPermissions = $JsonParameters.PSobject.Properties["taskPermissions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalTaskSetAccess"))) { #optional property not found
            $GlobalTaskSetAccess = $null
        } else {
            $GlobalTaskSetAccess = $JsonParameters.PSobject.Properties["globalTaskSetAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "taskSetPermissions"))) { #optional property not found
            $TaskSetPermissions = $null
        } else {
            $TaskSetPermissions = $JsonParameters.PSobject.Properties["taskSetPermissions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "success"))) { #optional property not found
            $Success = $null
        } else {
            $Success = $JsonParameters.PSobject.Properties["success"].value
        }

        $PSO = [PSCustomObject]@{
            "role" = ${Role}
            "featurePermissions" = ${FeaturePermissions}
            "globalSiteAccess" = ${GlobalSiteAccess}
            "sites" = ${Sites}
            "globalZoneAccess" = ${GlobalZoneAccess}
            "zones" = ${Zones}
            "globalInstanceTypeAccess" = ${GlobalInstanceTypeAccess}
            "instanceTypePermissions" = ${InstanceTypePermissions}
            "globalAppTemplateAccess" = ${GlobalAppTemplateAccess}
            "appTemplatePermissions" = ${AppTemplatePermissions}
            "globalCatalogItemTypeAccess" = ${GlobalCatalogItemTypeAccess}
            "catalogItemTypePermissions" = ${CatalogItemTypePermissions}
            "globalPersonaAccess" = ${GlobalPersonaAccess}
            "personaPermissions" = ${PersonaPermissions}
            "globalVdiPoolAccess" = ${GlobalVdiPoolAccess}
            "vdiPoolPermissions" = ${VdiPoolPermissions}
            "globalReportTypeAccess" = ${GlobalReportTypeAccess}
            "reportTypePermissions" = ${ReportTypePermissions}
            "globalTaskAccess" = ${GlobalTaskAccess}
            "taskPermissions" = ${TaskPermissions}
            "globalTaskSetAccess" = ${GlobalTaskSetAccess}
            "taskSetPermissions" = ${TaskSetPermissions}
            "success" = ${Success}
        }

        return $PSO
    }

}

