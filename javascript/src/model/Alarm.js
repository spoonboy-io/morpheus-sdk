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
 * The Alarm model module.
 * @module model/Alarm
 * @version 6.2.1
 */
class Alarm {
    /**
     * Constructs a new <code>Alarm</code>.
     * @alias module:model/Alarm
     */
    constructor() { 
        
        Alarm.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Alarm</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Alarm} obj Optional instance to populate.
     * @return {module:model/Alarm} The populated <code>Alarm</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Alarm();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
Alarm.prototype['id'] = undefined;

/**
 * @member {String} name
 */
Alarm.prototype['name'] = undefined;

/**
 * @member {Date} dateCreated
 */
Alarm.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
Alarm.prototype['lastUpdated'] = undefined;






export default Alarm;

