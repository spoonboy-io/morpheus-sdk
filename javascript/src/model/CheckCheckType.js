/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The CheckCheckType model module.
 * @module model/CheckCheckType
 * @version 6.1.1
 */
class CheckCheckType {
    /**
     * Constructs a new <code>CheckCheckType</code>.
     * @alias module:model/CheckCheckType
     */
    constructor() { 
        
        CheckCheckType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CheckCheckType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CheckCheckType} obj Optional instance to populate.
     * @return {module:model/CheckCheckType} The populated <code>CheckCheckType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CheckCheckType();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('metricName')) {
                obj['metricName'] = ApiClient.convertToType(data['metricName'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>CheckCheckType</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>CheckCheckType</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['code'] && !(typeof data['code'] === 'string' || data['code'] instanceof String)) {
            throw new Error("Expected the field `code` to be a primitive type in the JSON string but got " + data['code']);
        }
        // ensure the json data is a string
        if (data['name'] && !(typeof data['name'] === 'string' || data['name'] instanceof String)) {
            throw new Error("Expected the field `name` to be a primitive type in the JSON string but got " + data['name']);
        }
        // ensure the json data is a string
        if (data['metricName'] && !(typeof data['metricName'] === 'string' || data['metricName'] instanceof String)) {
            throw new Error("Expected the field `metricName` to be a primitive type in the JSON string but got " + data['metricName']);
        }

        return true;
    }


}



/**
 * @member {Number} id
 */
CheckCheckType.prototype['id'] = undefined;

/**
 * @member {String} code
 */
CheckCheckType.prototype['code'] = undefined;

/**
 * @member {String} name
 */
CheckCheckType.prototype['name'] = undefined;

/**
 * @member {String} metricName
 */
CheckCheckType.prototype['metricName'] = undefined;






export default CheckCheckType;
