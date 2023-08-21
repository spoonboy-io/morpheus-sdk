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
 * The NetworkRoutersCreateZone model module.
 * @module model/NetworkRoutersCreateZone
 * @version 6.2.1
 */
class NetworkRoutersCreateZone {
    /**
     * Constructs a new <code>NetworkRoutersCreateZone</code>.
     * Required when type does not support a network server
     * @alias module:model/NetworkRoutersCreateZone
     * @param id {Number} Cloud ID
     */
    constructor(id) { 
        
        NetworkRoutersCreateZone.initialize(this, id);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id) { 
        obj['id'] = id;
    }

    /**
     * Constructs a <code>NetworkRoutersCreateZone</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/NetworkRoutersCreateZone} obj Optional instance to populate.
     * @return {module:model/NetworkRoutersCreateZone} The populated <code>NetworkRoutersCreateZone</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new NetworkRoutersCreateZone();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * Cloud ID
 * @member {Number} id
 */
NetworkRoutersCreateZone.prototype['id'] = undefined;






export default NetworkRoutersCreateZone;

