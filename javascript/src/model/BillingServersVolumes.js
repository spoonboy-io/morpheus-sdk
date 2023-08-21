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
 * The BillingServersVolumes model module.
 * @module model/BillingServersVolumes
 * @version 6.2.1
 */
class BillingServersVolumes {
    /**
     * Constructs a new <code>BillingServersVolumes</code>.
     * @alias module:model/BillingServersVolumes
     */
    constructor() { 
        
        BillingServersVolumes.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingServersVolumes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingServersVolumes} obj Optional instance to populate.
     * @return {module:model/BillingServersVolumes} The populated <code>BillingServersVolumes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingServersVolumes();

            if (data.hasOwnProperty('size')) {
                obj['size'] = ApiClient.convertToType(data['size'], 'Number');
            }
            if (data.hasOwnProperty('typeCode')) {
                obj['typeCode'] = ApiClient.convertToType(data['typeCode'], 'String');
            }
            if (data.hasOwnProperty('datastore')) {
                obj['datastore'] = ApiClient.convertToType(data['datastore'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} size
 */
BillingServersVolumes.prototype['size'] = undefined;

/**
 * @member {String} typeCode
 */
BillingServersVolumes.prototype['typeCode'] = undefined;

/**
 * @member {String} datastore
 */
BillingServersVolumes.prototype['datastore'] = undefined;






export default BillingServersVolumes;

