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
 * The HealthLogs model module.
 * @module model/HealthLogs
 * @version 6.2.1
 */
class HealthLogs {
    /**
     * Constructs a new <code>HealthLogs</code>.
     * @alias module:model/HealthLogs
     */
    constructor() { 
        
        HealthLogs.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>HealthLogs</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/HealthLogs} obj Optional instance to populate.
     * @return {module:model/HealthLogs} The populated <code>HealthLogs</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new HealthLogs();

            if (data.hasOwnProperty('typeCode')) {
                obj['typeCode'] = ApiClient.convertToType(data['typeCode'], 'String');
            }
            if (data.hasOwnProperty('ts')) {
                obj['ts'] = ApiClient.convertToType(data['ts'], 'Date');
            }
            if (data.hasOwnProperty('level')) {
                obj['level'] = ApiClient.convertToType(data['level'], 'String');
            }
            if (data.hasOwnProperty('sourceType')) {
                obj['sourceType'] = ApiClient.convertToType(data['sourceType'], 'String');
            }
            if (data.hasOwnProperty('message')) {
                obj['message'] = ApiClient.convertToType(data['message'], 'String');
            }
            if (data.hasOwnProperty('hostname')) {
                obj['hostname'] = ApiClient.convertToType(data['hostname'], 'String');
            }
            if (data.hasOwnProperty('title')) {
                obj['title'] = ApiClient.convertToType(data['title'], 'String');
            }
            if (data.hasOwnProperty('logSignature')) {
                obj['logSignature'] = ApiClient.convertToType(data['logSignature'], 'String');
            }
            if (data.hasOwnProperty('objectId')) {
                obj['objectId'] = ApiClient.convertToType(data['objectId'], 'String');
            }
            if (data.hasOwnProperty('seq')) {
                obj['seq'] = ApiClient.convertToType(data['seq'], 'Number');
            }
            if (data.hasOwnProperty('_id')) {
                obj['_id'] = ApiClient.convertToType(data['_id'], 'String');
            }
            if (data.hasOwnProperty('signatureVerified')) {
                obj['signatureVerified'] = ApiClient.convertToType(data['signatureVerified'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {String} typeCode
 */
HealthLogs.prototype['typeCode'] = undefined;

/**
 * @member {Date} ts
 */
HealthLogs.prototype['ts'] = undefined;

/**
 * @member {String} level
 */
HealthLogs.prototype['level'] = undefined;

/**
 * @member {String} sourceType
 */
HealthLogs.prototype['sourceType'] = undefined;

/**
 * @member {String} message
 */
HealthLogs.prototype['message'] = undefined;

/**
 * @member {String} hostname
 */
HealthLogs.prototype['hostname'] = undefined;

/**
 * @member {String} title
 */
HealthLogs.prototype['title'] = undefined;

/**
 * @member {String} logSignature
 */
HealthLogs.prototype['logSignature'] = undefined;

/**
 * @member {String} objectId
 */
HealthLogs.prototype['objectId'] = undefined;

/**
 * @member {Number} seq
 */
HealthLogs.prototype['seq'] = undefined;

/**
 * @member {String} _id
 */
HealthLogs.prototype['_id'] = undefined;

/**
 * @member {Boolean} signatureVerified
 */
HealthLogs.prototype['signatureVerified'] = undefined;






export default HealthLogs;

