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
 * The Contact model module.
 * @module model/Contact
 * @version 6.2.1
 */
class Contact {
    /**
     * Constructs a new <code>Contact</code>.
     * @alias module:model/Contact
     */
    constructor() { 
        
        Contact.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Contact</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Contact} obj Optional instance to populate.
     * @return {module:model/Contact} The populated <code>Contact</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Contact();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('emailAddress')) {
                obj['emailAddress'] = ApiClient.convertToType(data['emailAddress'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('smsAddress')) {
                obj['smsAddress'] = ApiClient.convertToType(data['smsAddress'], 'String');
            }
            if (data.hasOwnProperty('slackHook')) {
                obj['slackHook'] = ApiClient.convertToType(data['slackHook'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
Contact.prototype['id'] = undefined;

/**
 * @member {String} emailAddress
 */
Contact.prototype['emailAddress'] = undefined;

/**
 * @member {String} name
 */
Contact.prototype['name'] = undefined;

/**
 * @member {String} smsAddress
 */
Contact.prototype['smsAddress'] = undefined;

/**
 * @member {String} slackHook
 */
Contact.prototype['slackHook'] = undefined;






export default Contact;

