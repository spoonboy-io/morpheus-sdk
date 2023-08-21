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
import InlineResponse20094Network from './InlineResponse20094Network';

/**
 * The SecurityPackage model module.
 * @module model/SecurityPackage
 * @version 6.2.1
 */
class SecurityPackage {
    /**
     * Constructs a new <code>SecurityPackage</code>.
     * @alias module:model/SecurityPackage
     */
    constructor() { 
        
        SecurityPackage.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>SecurityPackage</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/SecurityPackage} obj Optional instance to populate.
     * @return {module:model/SecurityPackage} The populated <code>SecurityPackage</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new SecurityPackage();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = InlineResponse20094Network.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('url')) {
                obj['url'] = ApiClient.convertToType(data['url'], 'String');
            }
            if (data.hasOwnProperty('uuid')) {
                obj['uuid'] = ApiClient.convertToType(data['uuid'], 'String');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
SecurityPackage.prototype['id'] = undefined;

/**
 * @member {String} name
 */
SecurityPackage.prototype['name'] = undefined;

/**
 * @member {Array.<String>} labels
 */
SecurityPackage.prototype['labels'] = undefined;

/**
 * @member {String} description
 */
SecurityPackage.prototype['description'] = undefined;

/**
 * @member {module:model/InlineResponse20094Network} type
 */
SecurityPackage.prototype['type'] = undefined;

/**
 * @member {Boolean} enabled
 */
SecurityPackage.prototype['enabled'] = undefined;

/**
 * @member {String} url
 */
SecurityPackage.prototype['url'] = undefined;

/**
 * @member {String} uuid
 */
SecurityPackage.prototype['uuid'] = undefined;

/**
 * @member {Date} dateCreated
 */
SecurityPackage.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
SecurityPackage.prototype['lastUpdated'] = undefined;

/**
 * @member {Object} config
 */
SecurityPackage.prototype['config'] = undefined;






export default SecurityPackage;

