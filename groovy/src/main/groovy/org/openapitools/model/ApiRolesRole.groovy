package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiRolesRoleAppTemplates;
import org.openapitools.model.ApiRolesRoleCatalogItemTypes;
import org.openapitools.model.ApiRolesRoleInstanceTypes;
import org.openapitools.model.ApiRolesRolePermissions;
import org.openapitools.model.ApiRolesRolePersonas;
import org.openapitools.model.ApiRolesRoleReportTypes;
import org.openapitools.model.ApiRolesRoleSites;
import org.openapitools.model.ApiRolesRoleTaskSets;
import org.openapitools.model.ApiRolesRoleTasks;
import org.openapitools.model.ApiRolesRoleVdiPools;
import org.openapitools.model.ApiRolesRoleZones;

@Canonical
class ApiRolesRole {
    /* Authority (Name) */
    String authority
    /* Description */
    String description
    /* Role type */
    String roleType = RoleTypeEnum.USER
    /* Base Role ID. Create the new role with the same permissions and access levels that the specified base role has. If this is not passed, the role is create without any permissions. */
    Long baseRoleId
    /* Multitenant roles are copied to all tenant accounts and kept in sync until a sub-tenant user modifies their copy of the role. *Only available to master tenant* */
    Boolean multitenant = false
    /* Multitenant Locked, prevents sub-tenant users from modifying their copy of multienant roles. *Only available to master tenant* */
    Boolean multitenantLocked = false
    
    String defaultPersona
    /* Set the access level for the specified permissions. */
    List<ApiRolesRolePermissions> permissions = new ArrayList<ApiRolesRolePermissions>()
    /* Set the default access level for for groups (sites). Only applies to user roles. */
    String globalSiteAccess
    /* Set the access level for the specified groups (sites). Only applies to user roles. */
    List<ApiRolesRoleSites> sites = new ArrayList<ApiRolesRoleSites>()
    /* Set the default access level for for clouds (zones). Only applies to base account (tenant) roles. */
    String globalZoneAccess
    /* Set the access level for the specified clouds (zones). Only applies to base account (tenant) roles. */
    List<ApiRolesRoleZones> zones = new ArrayList<ApiRolesRoleZones>()
    /* Set the default access level for for instance types */
    String globalInstanceTypeAccess
    /* Set the access level for the specified instance types */
    List<ApiRolesRoleInstanceTypes> instanceTypes = new ArrayList<ApiRolesRoleInstanceTypes>()
    /* Set the default access level for blueprints */
    String globalAppTemplateAccess
    /* Set the access level for the specified blueprints (appTemplates) */
    List<ApiRolesRoleAppTemplates> appTemplates = new ArrayList<ApiRolesRoleAppTemplates>()
    /* Set the default access level for catalog item types */
    String globalCatalogItemTypeAccess
    /* Set the access level for the specified catalog item types */
    List<ApiRolesRoleCatalogItemTypes> catalogItemTypes = new ArrayList<ApiRolesRoleCatalogItemTypes>()
    /* Set the default access level for personas */
    String globalPersonaAccess
    /* Set the access level for the specified personas */
    List<ApiRolesRolePersonas> personas = new ArrayList<ApiRolesRolePersonas>()
    /* Set the default access level for VDI pools */
    String globalVdiPoolAccess
    /* Set the access level for the specified VDI pools */
    List<ApiRolesRoleVdiPools> vdiPools = new ArrayList<ApiRolesRoleVdiPools>()
    /* Set the default access level for report types */
    String globalReportTypeAccess
    /* Set the access level for the specified report types */
    List<ApiRolesRoleReportTypes> reportTypes = new ArrayList<ApiRolesRoleReportTypes>()
    /* Set the default access level for tasks */
    String globalTaskAccess
    /* Set the access level for the specified tasks */
    List<ApiRolesRoleTasks> tasks = new ArrayList<ApiRolesRoleTasks>()
    /* Set the default access level for workflows (taskSets) */
    String globalTaskSetAccess
    /* Set the access level for the specified workflows (taskSets) */
    List<ApiRolesRoleTaskSets> taskSets = new ArrayList<ApiRolesRoleTaskSets>()
    /* Set the role owner (tenant) by ID. *Only available to master tenant* */
    Long owner
}
