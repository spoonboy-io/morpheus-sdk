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

/**
 * The NetworkPoolServerType model module.
 * @module model/NetworkPoolServerType
 * @version 6.2.1
 */
class NetworkPoolServerType {
    /**
     * Constructs a new <code>NetworkPoolServerType</code>.
     * Network Pool Server Type
     * @alias module:model/NetworkPoolServerType
     */
    constructor() { 
        
        NetworkPoolServerType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>NetworkPoolServerType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkPoolServerType} obj Optional instance to populate.
     * @return {module:model/NetworkPoolServerType} The populated <code>NetworkPoolServerType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkPoolServerType();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
NetworkPoolServerType.prototype['id'] = undefined;

/**
 * @member {String} name
 */
NetworkPoolServerType.prototype['name'] = undefined;

/**
 * @member {String} code
 */
NetworkPoolServerType.prototype['code'] = undefined;






export default NetworkPoolServerType;

