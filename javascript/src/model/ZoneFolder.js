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
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';
import InlineResponse20082LoadBalancerInstanceSslCert from './InlineResponse20082LoadBalancerInstanceSslCert';
import ResourcePermissions from './ResourcePermissions';
import ZoneDatastoreTenants from './ZoneDatastoreTenants';

/**
 * The ZoneFolder model module.
 * @module model/ZoneFolder
 * @version 6.2.1
 */
class ZoneFolder {
    /**
     * Constructs a new <code>ZoneFolder</code>.
     * @alias module:model/ZoneFolder
     */
    constructor() { 
        
        ZoneFolder.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ZoneFolder</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ZoneFolder} obj Optional instance to populate.
     * @return {module:model/ZoneFolder} The populated <code>ZoneFolder</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ZoneFolder();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('zone')) {
                obj['zone'] = InlineResponse20040AppDeployInstance.constructFromObject(data['zone']);
            }
            if (data.hasOwnProperty('parent')) {
                obj['parent'] = InlineResponse20082LoadBalancerInstanceSslCert.constructFromObject(data['parent']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('readOnly')) {
                obj['readOnly'] = ApiClient.convertToType(data['readOnly'], 'Boolean');
            }
            if (data.hasOwnProperty('defaultFolder')) {
                obj['defaultFolder'] = ApiClient.convertToType(data['defaultFolder'], 'Boolean');
            }
            if (data.hasOwnProperty('defaultStore')) {
                obj['defaultStore'] = ApiClient.convertToType(data['defaultStore'], 'Boolean');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('tenants')) {
                obj['tenants'] = ApiClient.convertToType(data['tenants'], [ZoneDatastoreTenants]);
            }
            if (data.hasOwnProperty('resourcePermissions')) {
                obj['resourcePermissions'] = ResourcePermissions.constructFromObject(data['resourcePermissions']);
            }
            if (data.hasOwnProperty('depth')) {
                obj['depth'] = ApiClient.convertToType(data['depth'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ZoneFolder.prototype['id'] = undefined;

/**
 * @member {String} name
 */
ZoneFolder.prototype['name'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} zone
 */
ZoneFolder.prototype['zone'] = undefined;

/**
 * @member {module:model/InlineResponse20082LoadBalancerInstanceSslCert} parent
 */
ZoneFolder.prototype['parent'] = undefined;

/**
 * @member {String} type
 */
ZoneFolder.prototype['type'] = undefined;

/**
 * @member {String} externalId
 */
ZoneFolder.prototype['externalId'] = undefined;

/**
 * @member {String} visibility
 */
ZoneFolder.prototype['visibility'] = undefined;

/**
 * @member {Boolean} readOnly
 */
ZoneFolder.prototype['readOnly'] = undefined;

/**
 * @member {Boolean} defaultFolder
 */
ZoneFolder.prototype['defaultFolder'] = undefined;

/**
 * @member {Boolean} defaultStore
 */
ZoneFolder.prototype['defaultStore'] = undefined;

/**
 * @member {Boolean} active
 */
ZoneFolder.prototype['active'] = undefined;

/**
 * @member {Array.<module:model/ZoneDatastoreTenants>} tenants
 */
ZoneFolder.prototype['tenants'] = undefined;

/**
 * @member {module:model/ResourcePermissions} resourcePermissions
 */
ZoneFolder.prototype['resourcePermissions'] = undefined;

/**
 * @member {Number} depth
 */
ZoneFolder.prototype['depth'] = undefined;






export default ZoneFolder;
