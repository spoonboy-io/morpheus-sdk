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
import ApiMonitoringContactsContact from './ApiMonitoringContactsContact';

/**
 * The InlineObject60 model module.
 * @module model/InlineObject60
 * @version 6.2.1
 */
class InlineObject60 {
    /**
     * Constructs a new <code>InlineObject60</code>.
     * @alias module:model/InlineObject60
     * @param contact {module:model/ApiMonitoringContactsContact} 
     */
    constructor(contact) { 
        
        InlineObject60.initialize(this, contact);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, contact) { 
        obj['contact'] = contact;
    }

    /**
     * Constructs a <code>InlineObject60</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineObject60} obj Optional instance to populate.
     * @return {module:model/InlineObject60} The populated <code>InlineObject60</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineObject60();

            if (data.hasOwnProperty('contact')) {
                obj['contact'] = ApiMonitoringContactsContact.constructFromObject(data['contact']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ApiMonitoringContactsContact} contact
 */
InlineObject60.prototype['contact'] = undefined;






export default InlineObject60;

