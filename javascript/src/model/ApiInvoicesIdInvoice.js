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
 * The ApiInvoicesIdInvoice model module.
 * @module model/ApiInvoicesIdInvoice
 * @version 6.2.1
 */
class ApiInvoicesIdInvoice {
    /**
     * Constructs a new <code>ApiInvoicesIdInvoice</code>.
     * @alias module:model/ApiInvoicesIdInvoice
     */
    constructor() { 
        
        ApiInvoicesIdInvoice.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiInvoicesIdInvoice</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiInvoicesIdInvoice} obj Optional instance to populate.
     * @return {module:model/ApiInvoicesIdInvoice} The populated <code>ApiInvoicesIdInvoice</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiInvoicesIdInvoice();

            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], [Object]);
            }
            if (data.hasOwnProperty('addTags')) {
                obj['addTags'] = ApiClient.convertToType(data['addTags'], [Object]);
            }
            if (data.hasOwnProperty('removeTags')) {
                obj['removeTags'] = ApiClient.convertToType(data['removeTags'], [Object]);
            }
        }
        return obj;
    }


}

/**
 * This adds or updates the specified Metadata tags and removes any tags not specified. Array of objects having a name and value. 
 * @member {Array.<Object>} tags
 */
ApiInvoicesIdInvoice.prototype['tags'] = undefined;

/**
 * Add or update value of Metadata tags. Array of objects having a name and value. 
 * @member {Array.<Object>} addTags
 */
ApiInvoicesIdInvoice.prototype['addTags'] = undefined;

/**
 * This removes the specified Metadata tags matching name and optionally value (if provided). Array of objects having a name and value. 
 * @member {Array.<Object>} removeTags
 */
ApiInvoicesIdInvoice.prototype['removeTags'] = undefined;






export default ApiInvoicesIdInvoice;

