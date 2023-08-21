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
 * The InstanceConnectionInfo model module.
 * @module model/InstanceConnectionInfo
 * @version 6.2.1
 */
class InstanceConnectionInfo {
    /**
     * Constructs a new <code>InstanceConnectionInfo</code>.
     * @alias module:model/InstanceConnectionInfo
     */
    constructor() { 
        
        InstanceConnectionInfo.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InstanceConnectionInfo</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceConnectionInfo} obj Optional instance to populate.
     * @return {module:model/InstanceConnectionInfo} The populated <code>InstanceConnectionInfo</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceConnectionInfo();

            if (data.hasOwnProperty('ip')) {
                obj['ip'] = ApiClient.convertToType(data['ip'], 'String');
            }
            if (data.hasOwnProperty('port')) {
                obj['port'] = ApiClient.convertToType(data['port'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {String} ip
 */
InstanceConnectionInfo.prototype['ip'] = undefined;

/**
 * @member {Number} port
 */
InstanceConnectionInfo.prototype['port'] = undefined;






export default InstanceConnectionInfo;

