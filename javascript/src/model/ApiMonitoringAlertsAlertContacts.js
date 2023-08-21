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
 * The ApiMonitoringAlertsAlertContacts model module.
 * @module model/ApiMonitoringAlertsAlertContacts
 * @version 6.2.1
 */
class ApiMonitoringAlertsAlertContacts {
    /**
     * Constructs a new <code>ApiMonitoringAlertsAlertContacts</code>.
     * @alias module:model/ApiMonitoringAlertsAlertContacts
     */
    constructor() { 
        
        ApiMonitoringAlertsAlertContacts.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiMonitoringAlertsAlertContacts</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiMonitoringAlertsAlertContacts} obj Optional instance to populate.
     * @return {module:model/ApiMonitoringAlertsAlertContacts} The populated <code>ApiMonitoringAlertsAlertContacts</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiMonitoringAlertsAlertContacts();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('method')) {
                obj['method'] = ApiClient.convertToType(data['method'], 'String');
            }
            if (data.hasOwnProperty('notify')) {
                obj['notify'] = ApiClient.convertToType(data['notify'], 'Boolean');
            }
            if (data.hasOwnProperty('close')) {
                obj['close'] = ApiClient.convertToType(data['close'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ApiMonitoringAlertsAlertContacts.prototype['id'] = undefined;

/**
 * @member {String} name
 */
ApiMonitoringAlertsAlertContacts.prototype['name'] = undefined;

/**
 * @member {String} method
 */
ApiMonitoringAlertsAlertContacts.prototype['method'] = undefined;

/**
 * @member {Boolean} notify
 */
ApiMonitoringAlertsAlertContacts.prototype['notify'] = undefined;

/**
 * @member {Boolean} close
 */
ApiMonitoringAlertsAlertContacts.prototype['close'] = undefined;






export default ApiMonitoringAlertsAlertContacts;

