/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import ApiRolesRoleAppTemplates from './ApiRolesRoleAppTemplates';
import ApiRolesRoleCatalogItemTypes from './ApiRolesRoleCatalogItemTypes';
import ApiRolesRoleInstanceTypes from './ApiRolesRoleInstanceTypes';
import ApiRolesRolePermissions from './ApiRolesRolePermissions';
import ApiRolesRolePersonas from './ApiRolesRolePersonas';
import ApiRolesRoleReportTypes from './ApiRolesRoleReportTypes';
import ApiRolesRoleSites from './ApiRolesRoleSites';
import ApiRolesRoleTaskSets from './ApiRolesRoleTaskSets';
import ApiRolesRoleTasks from './ApiRolesRoleTasks';
import ApiRolesRoleVdiPools from './ApiRolesRoleVdiPools';
import ApiRolesRoleZones from './ApiRolesRoleZones';

/**
 * The ApiRolesRole model module.
 * @module model/ApiRolesRole
 * @version 6.2.1
 */
class ApiRolesRole {
    /**
     * Constructs a new <code>ApiRolesRole</code>.
     * Payload for creating a new role
     * @alias module:model/ApiRolesRole
     * @param authority {String} Authority (Name)
     */
    constructor(authority) { 
        
        ApiRolesRole.initialize(this, authority);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, authority) { 
        obj['authority'] = authority;
    }

    /**
     * Constructs a <code>ApiRolesRole</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiRolesRole} obj Optional instance to populate.
     * @return {module:model/ApiRolesRole} The populated <code>ApiRolesRole</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiRolesRole();

            if (data.hasOwnProperty('authority')) {
                obj['authority'] = ApiClient.convertToType(data['authority'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('roleType')) {
                obj['roleType'] = ApiClient.convertToType(data['roleType'], 'String');
            }
            if (data.hasOwnProperty('baseRoleId')) {
                obj['baseRoleId'] = ApiClient.convertToType(data['baseRoleId'], 'Number');
            }
            if (data.hasOwnProperty('multitenant')) {
                obj['multitenant'] = ApiClient.convertToType(data['multitenant'], 'Boolean');
            }
            if (data.hasOwnProperty('multitenantLocked')) {
                obj['multitenantLocked'] = ApiClient.convertToType(data['multitenantLocked'], 'Boolean');
            }
            if (data.hasOwnProperty('defaultPersona')) {
                obj['defaultPersona'] = ApiClient.convertToType(data['defaultPersona'], 'String');
            }
            if (data.hasOwnProperty('permissions')) {
                obj['permissions'] = ApiClient.convertToType(data['permissions'], [ApiRolesRolePermissions]);
            }
            if (data.hasOwnProperty('globalSiteAccess')) {
                obj['globalSiteAccess'] = ApiClient.convertToType(data['globalSiteAccess'], 'String');
            }
            if (data.hasOwnProperty('sites')) {
                obj['sites'] = ApiClient.convertToType(data['sites'], [ApiRolesRoleSites]);
            }
            if (data.hasOwnProperty('globalZoneAccess')) {
                obj['globalZoneAccess'] = ApiClient.convertToType(data['globalZoneAccess'], 'String');
            }
            if (data.hasOwnProperty('zones')) {
                obj['zones'] = ApiClient.convertToType(data['zones'], [ApiRolesRoleZones]);
            }
            if (data.hasOwnProperty('globalInstanceTypeAccess')) {
                obj['globalInstanceTypeAccess'] = ApiClient.convertToType(data['globalInstanceTypeAccess'], 'String');
            }
            if (data.hasOwnProperty('instanceTypes')) {
                obj['instanceTypes'] = ApiClient.convertToType(data['instanceTypes'], [ApiRolesRoleInstanceTypes]);
            }
            if (data.hasOwnProperty('globalAppTemplateAccess')) {
                obj['globalAppTemplateAccess'] = ApiClient.convertToType(data['globalAppTemplateAccess'], 'String');
            }
            if (data.hasOwnProperty('appTemplates')) {
                obj['appTemplates'] = ApiClient.convertToType(data['appTemplates'], [ApiRolesRoleAppTemplates]);
            }
            if (data.hasOwnProperty('globalCatalogItemTypeAccess')) {
                obj['globalCatalogItemTypeAccess'] = ApiClient.convertToType(data['globalCatalogItemTypeAccess'], 'String');
            }
            if (data.hasOwnProperty('catalogItemTypes')) {
                obj['catalogItemTypes'] = ApiClient.convertToType(data['catalogItemTypes'], [ApiRolesRoleCatalogItemTypes]);
            }
            if (data.hasOwnProperty('globalPersonaAccess')) {
                obj['globalPersonaAccess'] = ApiClient.convertToType(data['globalPersonaAccess'], 'String');
            }
            if (data.hasOwnProperty('personas')) {
                obj['personas'] = ApiClient.convertToType(data['personas'], [ApiRolesRolePersonas]);
            }
            if (data.hasOwnProperty('globalVdiPoolAccess')) {
                obj['globalVdiPoolAccess'] = ApiClient.convertToType(data['globalVdiPoolAccess'], 'String');
            }
            if (data.hasOwnProperty('vdiPools')) {
                obj['vdiPools'] = ApiClient.convertToType(data['vdiPools'], [ApiRolesRoleVdiPools]);
            }
            if (data.hasOwnProperty('globalReportTypeAccess')) {
                obj['globalReportTypeAccess'] = ApiClient.convertToType(data['globalReportTypeAccess'], 'String');
            }
            if (data.hasOwnProperty('reportTypes')) {
                obj['reportTypes'] = ApiClient.convertToType(data['reportTypes'], [ApiRolesRoleReportTypes]);
            }
            if (data.hasOwnProperty('globalTaskAccess')) {
                obj['globalTaskAccess'] = ApiClient.convertToType(data['globalTaskAccess'], 'String');
            }
            if (data.hasOwnProperty('tasks')) {
                obj['tasks'] = ApiClient.convertToType(data['tasks'], [ApiRolesRoleTasks]);
            }
            if (data.hasOwnProperty('globalTaskSetAccess')) {
                obj['globalTaskSetAccess'] = ApiClient.convertToType(data['globalTaskSetAccess'], 'String');
            }
            if (data.hasOwnProperty('taskSets')) {
                obj['taskSets'] = ApiClient.convertToType(data['taskSets'], [ApiRolesRoleTaskSets]);
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = ApiClient.convertToType(data['owner'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * Authority (Name)
 * @member {String} authority
 */
ApiRolesRole.prototype['authority'] = undefined;

/**
 * Description
 * @member {String} description
 */
ApiRolesRole.prototype['description'] = undefined;

/**
 * Role type
 * @member {module:model/ApiRolesRole.RoleTypeEnum} roleType
 * @default 'user'
 */
ApiRolesRole.prototype['roleType'] = 'user';

/**
 * Base Role ID. Create the new role with the same permissions and access levels that the specified base role has. If this is not passed, the role is create without any permissions.
 * @member {Number} baseRoleId
 */
ApiRolesRole.prototype['baseRoleId'] = undefined;

/**
 * Multitenant roles are copied to all tenant accounts and kept in sync until a sub-tenant user modifies their copy of the role. *Only available to master tenant*
 * @member {Boolean} multitenant
 * @default false
 */
ApiRolesRole.prototype['multitenant'] = false;

/**
 * Multitenant Locked, prevents sub-tenant users from modifying their copy of multienant roles. *Only available to master tenant*
 * @member {Boolean} multitenantLocked
 * @default false
 */
ApiRolesRole.prototype['multitenantLocked'] = false;

/**
 * @member {module:model/ApiRolesRole.DefaultPersonaEnum} defaultPersona
 */
ApiRolesRole.prototype['defaultPersona'] = undefined;

/**
 * Set the access level for the specified permissions.
 * @member {Array.<module:model/ApiRolesRolePermissions>} permissions
 */
ApiRolesRole.prototype['permissions'] = undefined;

/**
 * Set the default access level for for groups (sites). Only applies to user roles.
 * @member {module:model/ApiRolesRole.GlobalSiteAccessEnum} globalSiteAccess
 */
ApiRolesRole.prototype['globalSiteAccess'] = undefined;

/**
 * Set the access level for the specified groups (sites). Only applies to user roles.
 * @member {Array.<module:model/ApiRolesRoleSites>} sites
 */
ApiRolesRole.prototype['sites'] = undefined;

/**
 * Set the default access level for for clouds (zones). Only applies to base account (tenant) roles.
 * @member {module:model/ApiRolesRole.GlobalZoneAccessEnum} globalZoneAccess
 */
ApiRolesRole.prototype['globalZoneAccess'] = undefined;

/**
 * Set the access level for the specified clouds (zones). Only applies to base account (tenant) roles.
 * @member {Array.<module:model/ApiRolesRoleZones>} zones
 */
ApiRolesRole.prototype['zones'] = undefined;

/**
 * Set the default access level for for instance types
 * @member {module:model/ApiRolesRole.GlobalInstanceTypeAccessEnum} globalInstanceTypeAccess
 */
ApiRolesRole.prototype['globalInstanceTypeAccess'] = undefined;

/**
 * Set the access level for the specified instance types
 * @member {Array.<module:model/ApiRolesRoleInstanceTypes>} instanceTypes
 */
ApiRolesRole.prototype['instanceTypes'] = undefined;

/**
 * Set the default access level for blueprints
 * @member {module:model/ApiRolesRole.GlobalAppTemplateAccessEnum} globalAppTemplateAccess
 */
ApiRolesRole.prototype['globalAppTemplateAccess'] = undefined;

/**
 * Set the access level for the specified blueprints (appTemplates)
 * @member {Array.<module:model/ApiRolesRoleAppTemplates>} appTemplates
 */
ApiRolesRole.prototype['appTemplates'] = undefined;

/**
 * Set the default access level for catalog item types
 * @member {module:model/ApiRolesRole.GlobalCatalogItemTypeAccessEnum} globalCatalogItemTypeAccess
 */
ApiRolesRole.prototype['globalCatalogItemTypeAccess'] = undefined;

/**
 * Set the access level for the specified catalog item types
 * @member {Array.<module:model/ApiRolesRoleCatalogItemTypes>} catalogItemTypes
 */
ApiRolesRole.prototype['catalogItemTypes'] = undefined;

/**
 * Set the default access level for personas
 * @member {module:model/ApiRolesRole.GlobalPersonaAccessEnum} globalPersonaAccess
 */
ApiRolesRole.prototype['globalPersonaAccess'] = undefined;

/**
 * Set the access level for the specified personas
 * @member {Array.<module:model/ApiRolesRolePersonas>} personas
 */
ApiRolesRole.prototype['personas'] = undefined;

/**
 * Set the default access level for VDI pools
 * @member {module:model/ApiRolesRole.GlobalVdiPoolAccessEnum} globalVdiPoolAccess
 */
ApiRolesRole.prototype['globalVdiPoolAccess'] = undefined;

/**
 * Set the access level for the specified VDI pools
 * @member {Array.<module:model/ApiRolesRoleVdiPools>} vdiPools
 */
ApiRolesRole.prototype['vdiPools'] = undefined;

/**
 * Set the default access level for report types
 * @member {module:model/ApiRolesRole.GlobalReportTypeAccessEnum} globalReportTypeAccess
 */
ApiRolesRole.prototype['globalReportTypeAccess'] = undefined;

/**
 * Set the access level for the specified report types
 * @member {Array.<module:model/ApiRolesRoleReportTypes>} reportTypes
 */
ApiRolesRole.prototype['reportTypes'] = undefined;

/**
 * Set the default access level for tasks
 * @member {module:model/ApiRolesRole.GlobalTaskAccessEnum} globalTaskAccess
 */
ApiRolesRole.prototype['globalTaskAccess'] = undefined;

/**
 * Set the access level for the specified tasks
 * @member {Array.<module:model/ApiRolesRoleTasks>} tasks
 */
ApiRolesRole.prototype['tasks'] = undefined;

/**
 * Set the default access level for workflows (taskSets)
 * @member {module:model/ApiRolesRole.GlobalTaskSetAccessEnum} globalTaskSetAccess
 */
ApiRolesRole.prototype['globalTaskSetAccess'] = undefined;

/**
 * Set the access level for the specified workflows (taskSets)
 * @member {Array.<module:model/ApiRolesRoleTaskSets>} taskSets
 */
ApiRolesRole.prototype['taskSets'] = undefined;

/**
 * Set the role owner (tenant) by ID. *Only available to master tenant*
 * @member {Number} owner
 */
ApiRolesRole.prototype['owner'] = undefined;





/**
 * Allowed values for the <code>roleType</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRole['RoleTypeEnum'] = {

    /**
     * value: "user"
     * @const
     */
    "user": "user",

    /**
     * value: "account"
     * @const
     */
    "account": "account"
};


/**
 * Allowed values for the <code>defaultPersona</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRole['DefaultPersonaEnum'] = {

    /**
     * value: "standard"
     * @const
     */
    "standard": "standard",

    /**
     * value: "serviceCatalog"
     * @const
     */
    "serviceCatalog": "serviceCatalog",

    /**
     * value: "vdi"
     * @const
     */
    "vdi": "vdi"
};


/**
 * Allowed values for the <code>globalSiteAccess</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRole['GlobalSiteAccessEnum'] = {

    /**
     * value: "default"
     * @const
     */
    "default": "default",

    /**
     * value: "full"
     * @const
     */
    "full": "full",

    /**
     * value: "read"
     * @const
     */
    "read": "read",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};


/**
 * Allowed values for the <code>globalZoneAccess</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRole['GlobalZoneAccessEnum'] = {

    /**
     * value: "default"
     * @const
     */
    "default": "default",

    /**
     * value: "full"
     * @const
     */
    "full": "full",

    /**
     * value: "read"
     * @const
     */
    "read": "read",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};


/**
 * Allowed values for the <code>globalInstanceTypeAccess</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRole['GlobalInstanceTypeAccessEnum'] = {

    /**
     * value: "full"
     * @const
     */
    "full": "full",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};


/**
 * Allowed values for the <code>globalAppTemplateAccess</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRole['GlobalAppTemplateAccessEnum'] = {

    /**
     * value: "full"
     * @const
     */
    "full": "full",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};


/**
 * Allowed values for the <code>globalCatalogItemTypeAccess</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRole['GlobalCatalogItemTypeAccessEnum'] = {

    /**
     * value: "full"
     * @const
     */
    "full": "full",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};


/**
 * Allowed values for the <code>globalPersonaAccess</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRole['GlobalPersonaAccessEnum'] = {

    /**
     * value: "full"
     * @const
     */
    "full": "full",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};


/**
 * Allowed values for the <code>globalVdiPoolAccess</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRole['GlobalVdiPoolAccessEnum'] = {

    /**
     * value: "full"
     * @const
     */
    "full": "full",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};


/**
 * Allowed values for the <code>globalReportTypeAccess</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRole['GlobalReportTypeAccessEnum'] = {

    /**
     * value: "full"
     * @const
     */
    "full": "full",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};


/**
 * Allowed values for the <code>globalTaskAccess</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRole['GlobalTaskAccessEnum'] = {

    /**
     * value: "full"
     * @const
     */
    "full": "full",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};


/**
 * Allowed values for the <code>globalTaskSetAccess</code> property.
 * @enum {String}
 * @readonly
 */
ApiRolesRole['GlobalTaskSetAccessEnum'] = {

    /**
     * value: "full"
     * @const
     */
    "full": "full",

    /**
     * value: "none"
     * @const
     */
    "none": "none"
};



export default ApiRolesRole;
