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
import MetaObject from './MetaObject';
import UsagesActivity from './UsagesActivity';

/**
 * The Usages model module.
 * @module model/Usages
 * @version 6.2.1
 */
class Usages {
    /**
     * Constructs a new <code>Usages</code>.
     * @alias module:model/Usages
     */
    constructor() { 
        
        Usages.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Usages</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Usages} obj Optional instance to populate.
     * @return {module:model/Usages} The populated <code>Usages</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Usages();

            if (data.hasOwnProperty('activity')) {
                obj['activity'] = ApiClient.convertToType(data['activity'], [UsagesActivity]);
            }
            if (data.hasOwnProperty('meta')) {
                obj['meta'] = ApiClient.convertToType(data['meta'], MetaObject);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/UsagesActivity>} activity
 */
Usages.prototype['activity'] = undefined;

/**
 * @member {module:model/MetaObject} meta
 */
Usages.prototype['meta'] = undefined;






export default Usages;
