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
import AnyOfobjectobjectobject from './AnyOfobjectobjectobject';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';
import InlineResponse20082LoadBalancerInstanceSslCert from './InlineResponse20082LoadBalancerInstanceSslCert';
import ResourcePermissions from './ResourcePermissions';

/**
 * The ZoneResourcePool model module.
 * @module model/ZoneResourcePool
 * @version 6.2.1
 */
class ZoneResourcePool {
    /**
     * Constructs a new <code>ZoneResourcePool</code>.
     * @alias module:model/ZoneResourcePool
     */
    constructor() { 
        
        ZoneResourcePool.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ZoneResourcePool</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ZoneResourcePool} obj Optional instance to populate.
     * @return {module:model/ZoneResourcePool} The populated <code>ZoneResourcePool</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ZoneResourcePool();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
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
            if (data.hasOwnProperty('regionCode')) {
                obj['regionCode'] = ApiClient.convertToType(data['regionCode'], 'String');
            }
            if (data.hasOwnProperty('iacId')) {
                obj['iacId'] = ApiClient.convertToType(data['iacId'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('readOnly')) {
                obj['readOnly'] = ApiClient.convertToType(data['readOnly'], 'Boolean');
            }
            if (data.hasOwnProperty('defaultPool')) {
                obj['defaultPool'] = ApiClient.convertToType(data['defaultPool'], 'Boolean');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('inventory')) {
                obj['inventory'] = ApiClient.convertToType(data['inventory'], 'Boolean');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], AnyOfobjectobjectobject);
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('displayName')) {
                obj['displayName'] = ApiClient.convertToType(data['displayName'], 'String');
            }
            if (data.hasOwnProperty('tenants')) {
                obj['tenants'] = ApiClient.convertToType(data['tenants'], [InlineResponse20040AppDeployInstance]);
            }
            if (data.hasOwnProperty('resourcePermission')) {
                obj['resourcePermission'] = ResourcePermissions.constructFromObject(data['resourcePermission']);
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
ZoneResourcePool.prototype['id'] = undefined;

/**
 * @member {String} description
 */
ZoneResourcePool.prototype['description'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} zone
 */
ZoneResourcePool.prototype['zone'] = undefined;

/**
 * @member {module:model/InlineResponse20082LoadBalancerInstanceSslCert} parent
 */
ZoneResourcePool.prototype['parent'] = undefined;

/**
 * @member {String} type
 */
ZoneResourcePool.prototype['type'] = undefined;

/**
 * @member {String} externalId
 */
ZoneResourcePool.prototype['externalId'] = undefined;

/**
 * @member {String} regionCode
 */
ZoneResourcePool.prototype['regionCode'] = undefined;

/**
 * @member {String} iacId
 */
ZoneResourcePool.prototype['iacId'] = undefined;

/**
 * @member {String} visibility
 */
ZoneResourcePool.prototype['visibility'] = undefined;

/**
 * @member {Boolean} readOnly
 */
ZoneResourcePool.prototype['readOnly'] = undefined;

/**
 * @member {Boolean} defaultPool
 */
ZoneResourcePool.prototype['defaultPool'] = undefined;

/**
 * @member {Boolean} active
 */
ZoneResourcePool.prototype['active'] = undefined;

/**
 * @member {String} status
 */
ZoneResourcePool.prototype['status'] = undefined;

/**
 * @member {Boolean} inventory
 */
ZoneResourcePool.prototype['inventory'] = undefined;

/**
 * @member {module:model/AnyOfobjectobjectobject} config
 */
ZoneResourcePool.prototype['config'] = undefined;

/**
 * @member {String} name
 */
ZoneResourcePool.prototype['name'] = undefined;

/**
 * @member {String} displayName
 */
ZoneResourcePool.prototype['displayName'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse20040AppDeployInstance>} tenants
 */
ZoneResourcePool.prototype['tenants'] = undefined;

/**
 * @member {module:model/ResourcePermissions} resourcePermission
 */
ZoneResourcePool.prototype['resourcePermission'] = undefined;

/**
 * @member {Number} depth
 */
ZoneResourcePool.prototype['depth'] = undefined;






export default ZoneResourcePool;
