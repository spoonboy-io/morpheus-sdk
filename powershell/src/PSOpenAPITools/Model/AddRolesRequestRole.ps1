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

Payload for creating a new role

.PARAMETER Authority
Authority (Name)
.PARAMETER Description
Description
.PARAMETER RoleType
Role type
.PARAMETER BaseRoleId
Base Role ID. Create the new role with the same permissions and access levels that the specified base role has. If this is not passed, the role is create without any permissions.
.PARAMETER Multitenant
Multitenant roles are copied to all tenant accounts and kept in sync until a sub-tenant user modifies their copy of the role. *Only available to master tenant*
.PARAMETER MultitenantLocked
Multitenant Locked, prevents sub-tenant users from modifying their copy of multienant roles. *Only available to master tenant*
.PARAMETER DefaultPersona
No description available.
.PARAMETER Permissions
Set the access level for the specified permissions.
.PARAMETER GlobalSiteAccess
Set the default access level for for groups (sites). Only applies to user roles.
.PARAMETER Sites
Set the access level for the specified groups (sites). Only applies to user roles.
.PARAMETER GlobalZoneAccess
Set the default access level for for clouds (zones). Only applies to base account (tenant) roles.
.PARAMETER Zones
Set the access level for the specified clouds (zones). Only applies to base account (tenant) roles.
.PARAMETER GlobalInstanceTypeAccess
Set the default access level for for instance types
.PARAMETER InstanceTypes
Set the access level for the specified instance types
.PARAMETER GlobalAppTemplateAccess
Set the default access level for blueprints
.PARAMETER AppTemplates
Set the access level for the specified blueprints (appTemplates)
.PARAMETER GlobalCatalogItemTypeAccess
Set the default access level for catalog item types
.PARAMETER CatalogItemTypes
Set the access level for the specified catalog item types
.PARAMETER GlobalPersonaAccess
Set the default access level for personas
.PARAMETER Personas
Set the access level for the specified personas
.PARAMETER GlobalVdiPoolAccess
Set the default access level for VDI pools
.PARAMETER VdiPools
Set the access level for the specified VDI pools
.PARAMETER GlobalReportTypeAccess
Set the default access level for report types
.PARAMETER ReportTypes
Set the access level for the specified report types
.PARAMETER GlobalTaskAccess
Set the default access level for tasks
.PARAMETER Tasks
Set the access level for the specified tasks
.PARAMETER GlobalTaskSetAccess
Set the default access level for workflows (taskSets)
.PARAMETER TaskSets
Set the access level for the specified workflows (taskSets)
.PARAMETER Owner
Set the role owner (tenant) by ID. *Only available to master tenant*
.OUTPUTS

AddRolesRequestRole<PSCustomObject>
#>

function Initialize-AddRolesRequestRole {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Authority},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("user", "account")]
        [String]
        ${RoleType} = "user",
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${BaseRoleId},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Multitenant} = $false,
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${MultitenantLocked} = $false,
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("standard", "serviceCatalog", "vdi")]
        [String]
        ${DefaultPersona},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Permissions},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("default", "full", "read", "none")]
        [String]
        ${GlobalSiteAccess},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Sites},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("default", "full", "read", "none")]
        [String]
        ${GlobalZoneAccess},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Zones},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("full", "none")]
        [String]
        ${GlobalInstanceTypeAccess},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${InstanceTypes},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("full", "none")]
        [String]
        ${GlobalAppTemplateAccess},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${AppTemplates},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("full", "none")]
        [String]
        ${GlobalCatalogItemTypeAccess},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${CatalogItemTypes},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("full", "none")]
        [String]
        ${GlobalPersonaAccess},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Personas},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("full", "none")]
        [String]
        ${GlobalVdiPoolAccess},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${VdiPools},
        [Parameter(Position = 22, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("full", "none")]
        [String]
        ${GlobalReportTypeAccess},
        [Parameter(Position = 23, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${ReportTypes},
        [Parameter(Position = 24, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("full", "none")]
        [String]
        ${GlobalTaskAccess},
        [Parameter(Position = 25, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Tasks},
        [Parameter(Position = 26, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("full", "none")]
        [String]
        ${GlobalTaskSetAccess},
        [Parameter(Position = 27, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${TaskSets},
        [Parameter(Position = 28, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Owner}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddRolesRequestRole' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Authority) {
            throw "invalid value for 'Authority', 'Authority' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "authority" = ${Authority}
            "description" = ${Description}
            "roleType" = ${RoleType}
            "baseRoleId" = ${BaseRoleId}
            "multitenant" = ${Multitenant}
            "multitenantLocked" = ${MultitenantLocked}
            "defaultPersona" = ${DefaultPersona}
            "permissions" = ${Permissions}
            "globalSiteAccess" = ${GlobalSiteAccess}
            "sites" = ${Sites}
            "globalZoneAccess" = ${GlobalZoneAccess}
            "zones" = ${Zones}
            "globalInstanceTypeAccess" = ${GlobalInstanceTypeAccess}
            "instanceTypes" = ${InstanceTypes}
            "globalAppTemplateAccess" = ${GlobalAppTemplateAccess}
            "appTemplates" = ${AppTemplates}
            "globalCatalogItemTypeAccess" = ${GlobalCatalogItemTypeAccess}
            "catalogItemTypes" = ${CatalogItemTypes}
            "globalPersonaAccess" = ${GlobalPersonaAccess}
            "personas" = ${Personas}
            "globalVdiPoolAccess" = ${GlobalVdiPoolAccess}
            "vdiPools" = ${VdiPools}
            "globalReportTypeAccess" = ${GlobalReportTypeAccess}
            "reportTypes" = ${ReportTypes}
            "globalTaskAccess" = ${GlobalTaskAccess}
            "tasks" = ${Tasks}
            "globalTaskSetAccess" = ${GlobalTaskSetAccess}
            "taskSets" = ${TaskSets}
            "owner" = ${Owner}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddRolesRequestRole<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddRolesRequestRole<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddRolesRequestRole<PSCustomObject>
#>
function ConvertFrom-JsonToAddRolesRequestRole {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddRolesRequestRole' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddRolesRequestRole
        $AllProperties = ("authority", "description", "roleType", "baseRoleId", "multitenant", "multitenantLocked", "defaultPersona", "permissions", "globalSiteAccess", "sites", "globalZoneAccess", "zones", "globalInstanceTypeAccess", "instanceTypes", "globalAppTemplateAccess", "appTemplates", "globalCatalogItemTypeAccess", "catalogItemTypes", "globalPersonaAccess", "personas", "globalVdiPoolAccess", "vdiPools", "globalReportTypeAccess", "reportTypes", "globalTaskAccess", "tasks", "globalTaskSetAccess", "taskSets", "owner")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'authority' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "authority"))) {
            throw "Error! JSON cannot be serialized due to the required property 'authority' missing."
        } else {
            $Authority = $JsonParameters.PSobject.Properties["authority"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "roleType"))) { #optional property not found
            $RoleType = $null
        } else {
            $RoleType = $JsonParameters.PSobject.Properties["roleType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "baseRoleId"))) { #optional property not found
            $BaseRoleId = $null
        } else {
            $BaseRoleId = $JsonParameters.PSobject.Properties["baseRoleId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "multitenant"))) { #optional property not found
            $Multitenant = $null
        } else {
            $Multitenant = $JsonParameters.PSobject.Properties["multitenant"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "multitenantLocked"))) { #optional property not found
            $MultitenantLocked = $null
        } else {
            $MultitenantLocked = $JsonParameters.PSobject.Properties["multitenantLocked"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "defaultPersona"))) { #optional property not found
            $DefaultPersona = $null
        } else {
            $DefaultPersona = $JsonParameters.PSobject.Properties["defaultPersona"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "permissions"))) { #optional property not found
            $Permissions = $null
        } else {
            $Permissions = $JsonParameters.PSobject.Properties["permissions"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceTypes"))) { #optional property not found
            $InstanceTypes = $null
        } else {
            $InstanceTypes = $JsonParameters.PSobject.Properties["instanceTypes"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalAppTemplateAccess"))) { #optional property not found
            $GlobalAppTemplateAccess = $null
        } else {
            $GlobalAppTemplateAccess = $JsonParameters.PSobject.Properties["globalAppTemplateAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "appTemplates"))) { #optional property not found
            $AppTemplates = $null
        } else {
            $AppTemplates = $JsonParameters.PSobject.Properties["appTemplates"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalCatalogItemTypeAccess"))) { #optional property not found
            $GlobalCatalogItemTypeAccess = $null
        } else {
            $GlobalCatalogItemTypeAccess = $JsonParameters.PSobject.Properties["globalCatalogItemTypeAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "catalogItemTypes"))) { #optional property not found
            $CatalogItemTypes = $null
        } else {
            $CatalogItemTypes = $JsonParameters.PSobject.Properties["catalogItemTypes"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalPersonaAccess"))) { #optional property not found
            $GlobalPersonaAccess = $null
        } else {
            $GlobalPersonaAccess = $JsonParameters.PSobject.Properties["globalPersonaAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "personas"))) { #optional property not found
            $Personas = $null
        } else {
            $Personas = $JsonParameters.PSobject.Properties["personas"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalVdiPoolAccess"))) { #optional property not found
            $GlobalVdiPoolAccess = $null
        } else {
            $GlobalVdiPoolAccess = $JsonParameters.PSobject.Properties["globalVdiPoolAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "vdiPools"))) { #optional property not found
            $VdiPools = $null
        } else {
            $VdiPools = $JsonParameters.PSobject.Properties["vdiPools"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalReportTypeAccess"))) { #optional property not found
            $GlobalReportTypeAccess = $null
        } else {
            $GlobalReportTypeAccess = $JsonParameters.PSobject.Properties["globalReportTypeAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "reportTypes"))) { #optional property not found
            $ReportTypes = $null
        } else {
            $ReportTypes = $JsonParameters.PSobject.Properties["reportTypes"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalTaskAccess"))) { #optional property not found
            $GlobalTaskAccess = $null
        } else {
            $GlobalTaskAccess = $JsonParameters.PSobject.Properties["globalTaskAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tasks"))) { #optional property not found
            $Tasks = $null
        } else {
            $Tasks = $JsonParameters.PSobject.Properties["tasks"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "globalTaskSetAccess"))) { #optional property not found
            $GlobalTaskSetAccess = $null
        } else {
            $GlobalTaskSetAccess = $JsonParameters.PSobject.Properties["globalTaskSetAccess"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "taskSets"))) { #optional property not found
            $TaskSets = $null
        } else {
            $TaskSets = $JsonParameters.PSobject.Properties["taskSets"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "owner"))) { #optional property not found
            $Owner = $null
        } else {
            $Owner = $JsonParameters.PSobject.Properties["owner"].value
        }

        $PSO = [PSCustomObject]@{
            "authority" = ${Authority}
            "description" = ${Description}
            "roleType" = ${RoleType}
            "baseRoleId" = ${BaseRoleId}
            "multitenant" = ${Multitenant}
            "multitenantLocked" = ${MultitenantLocked}
            "defaultPersona" = ${DefaultPersona}
            "permissions" = ${Permissions}
            "globalSiteAccess" = ${GlobalSiteAccess}
            "sites" = ${Sites}
            "globalZoneAccess" = ${GlobalZoneAccess}
            "zones" = ${Zones}
            "globalInstanceTypeAccess" = ${GlobalInstanceTypeAccess}
            "instanceTypes" = ${InstanceTypes}
            "globalAppTemplateAccess" = ${GlobalAppTemplateAccess}
            "appTemplates" = ${AppTemplates}
            "globalCatalogItemTypeAccess" = ${GlobalCatalogItemTypeAccess}
            "catalogItemTypes" = ${CatalogItemTypes}
            "globalPersonaAccess" = ${GlobalPersonaAccess}
            "personas" = ${Personas}
            "globalVdiPoolAccess" = ${GlobalVdiPoolAccess}
            "vdiPools" = ${VdiPools}
            "globalReportTypeAccess" = ${GlobalReportTypeAccess}
            "reportTypes" = ${ReportTypes}
            "globalTaskAccess" = ${GlobalTaskAccess}
            "tasks" = ${Tasks}
            "globalTaskSetAccess" = ${GlobalTaskSetAccess}
            "taskSets" = ${TaskSets}
            "owner" = ${Owner}
        }

        return $PSO
    }

}

