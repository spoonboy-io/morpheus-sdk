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
import ApiBlueprintsIdUpdatePermissionsResourcePermissionSites from './ApiBlueprintsIdUpdatePermissionsResourcePermissionSites';
import NetworkServerGroupMember from './NetworkServerGroupMember';
import Permissions from './Permissions';
import Tag from './Tag';

/**
 * The InlineResponse200117Group model module.
 * @module model/InlineResponse200117Group
 * @version 6.2.1
 */
class InlineResponse200117Group {
    /**
     * Constructs a new <code>InlineResponse200117Group</code>.
     * @alias module:model/InlineResponse200117Group
     */
    constructor() { 
        
        InlineResponse200117Group.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse200117Group</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse200117Group} obj Optional instance to populate.
     * @return {module:model/InlineResponse200117Group} The populated <code>InlineResponse200117Group</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse200117Group();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('internalId')) {
                obj['internalId'] = ApiClient.convertToType(data['internalId'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['owner']);
            }
            if (data.hasOwnProperty('networkServer')) {
                obj['networkServer'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['networkServer']);
            }
            if (data.hasOwnProperty('permissions')) {
                obj['permissions'] = Permissions.constructFromObject(data['permissions']);
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], [Tag]);
            }
            if (data.hasOwnProperty('members')) {
                obj['members'] = ApiClient.convertToType(data['members'], [NetworkServerGroupMember]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
InlineResponse200117Group.prototype['id'] = undefined;

/**
 * @member {String} name
 */
InlineResponse200117Group.prototype['name'] = undefined;

/**
 * @member {String} description
 */
InlineResponse200117Group.prototype['description'] = undefined;

/**
 * @member {String} internalId
 */
InlineResponse200117Group.prototype['internalId'] = undefined;

/**
 * @member {String} externalId
 */
InlineResponse200117Group.prototype['externalId'] = undefined;

/**
 * @member {String} visibility
 */
InlineResponse200117Group.prototype['visibility'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} account
 */
InlineResponse200117Group.prototype['account'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} owner
 */
InlineResponse200117Group.prototype['owner'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} networkServer
 */
InlineResponse200117Group.prototype['networkServer'] = undefined;

/**
 * @member {module:model/Permissions} permissions
 */
InlineResponse200117Group.prototype['permissions'] = undefined;

/**
 * @member {Array.<module:model/Tag>} tags
 */
InlineResponse200117Group.prototype['tags'] = undefined;

/**
 * @member {Array.<module:model/NetworkServerGroupMember>} members
 */
InlineResponse200117Group.prototype['members'] = undefined;






export default InlineResponse200117Group;
