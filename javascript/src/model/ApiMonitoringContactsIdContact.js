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
 * The ApiMonitoringContactsIdContact model module.
 * @module model/ApiMonitoringContactsIdContact
 * @version 6.2.1
 */
class ApiMonitoringContactsIdContact {
    /**
     * Constructs a new <code>ApiMonitoringContactsIdContact</code>.
     * Payload for updating a monitoring contact
     * @alias module:model/ApiMonitoringContactsIdContact
     */
    constructor() { 
        
        ApiMonitoringContactsIdContact.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiMonitoringContactsIdContact</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiMonitoringContactsIdContact} obj Optional instance to populate.
     * @return {module:model/ApiMonitoringContactsIdContact} The populated <code>ApiMonitoringContactsIdContact</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiMonitoringContactsIdContact();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('emailAddress')) {
                obj['emailAddress'] = ApiClient.convertToType(data['emailAddress'], 'String');
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
 * Unique name scoped to your account for the contact
 * @member {String} name
 */
ApiMonitoringContactsIdContact.prototype['name'] = undefined;

/**
 * Email notification address
 * @member {String} emailAddress
 */
ApiMonitoringContactsIdContact.prototype['emailAddress'] = undefined;

/**
 * SMS notification address
 * @member {String} smsAddress
 */
ApiMonitoringContactsIdContact.prototype['smsAddress'] = undefined;

/**
 * Slack Hook
 * @member {String} slackHook
 */
ApiMonitoringContactsIdContact.prototype['slackHook'] = undefined;






export default ApiMonitoringContactsIdContact;

