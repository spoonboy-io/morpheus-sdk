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
import InlineResponse20082LoadBalancerInstanceSslCert from './InlineResponse20082LoadBalancerInstanceSslCert';

/**
 * The SecurityGroupLocations model module.
 * @module model/SecurityGroupLocations
 * @version 6.2.1
 */
class SecurityGroupLocations {
    /**
     * Constructs a new <code>SecurityGroupLocations</code>.
     * @alias module:model/SecurityGroupLocations
     */
    constructor() { 
        
        SecurityGroupLocations.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>SecurityGroupLocations</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SecurityGroupLocations} obj Optional instance to populate.
     * @return {module:model/SecurityGroupLocations} The populated <code>SecurityGroupLocations</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new SecurityGroupLocations();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('iacId')) {
                obj['iacId'] = ApiClient.convertToType(data['iacId'], 'String');
            }
            if (data.hasOwnProperty('zone')) {
                obj['zone'] = InlineResponse20082LoadBalancerInstanceSslCert.constructFromObject(data['zone']);
            }
            if (data.hasOwnProperty('zonePool')) {
                obj['zonePool'] = InlineResponse20082LoadBalancerInstanceSslCert.constructFromObject(data['zonePool']);
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('priority')) {
                obj['priority'] = ApiClient.convertToType(data['priority'], 'String');
            }
            if (data.hasOwnProperty('groupLayer')) {
                obj['groupLayer'] = ApiClient.convertToType(data['groupLayer'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
SecurityGroupLocations.prototype['id'] = undefined;

/**
 * @member {String} name
 */
SecurityGroupLocations.prototype['name'] = undefined;

/**
 * @member {String} description
 */
SecurityGroupLocations.prototype['description'] = undefined;

/**
 * @member {String} externalId
 */
SecurityGroupLocations.prototype['externalId'] = undefined;

/**
 * @member {String} iacId
 */
SecurityGroupLocations.prototype['iacId'] = undefined;

/**
 * @member {module:model/InlineResponse20082LoadBalancerInstanceSslCert} zone
 */
SecurityGroupLocations.prototype['zone'] = undefined;

/**
 * @member {module:model/InlineResponse20082LoadBalancerInstanceSslCert} zonePool
 */
SecurityGroupLocations.prototype['zonePool'] = undefined;

/**
 * @member {String} status
 */
SecurityGroupLocations.prototype['status'] = undefined;

/**
 * @member {String} priority
 */
SecurityGroupLocations.prototype['priority'] = undefined;

/**
 * @member {String} groupLayer
 */
SecurityGroupLocations.prototype['groupLayer'] = undefined;






export default SecurityGroupLocations;

