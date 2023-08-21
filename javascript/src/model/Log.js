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
import LogData from './LogData';
import LogSort from './LogSort';

/**
 * The Log model module.
 * @module model/Log
 * @version 6.2.1
 */
class Log {
    /**
     * Constructs a new <code>Log</code>.
     * @alias module:model/Log
     */
    constructor() { 
        
        Log.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Log</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Log} obj Optional instance to populate.
     * @return {module:model/Log} The populated <code>Log</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Log();

            if (data.hasOwnProperty('sort')) {
                obj['sort'] = LogSort.constructFromObject(data['sort']);
            }
            if (data.hasOwnProperty('offset')) {
                obj['offset'] = ApiClient.convertToType(data['offset'], 'Number');
            }
            if (data.hasOwnProperty('start')) {
                obj['start'] = ApiClient.convertToType(data['start'], 'Date');
            }
            if (data.hasOwnProperty('end')) {
                obj['end'] = ApiClient.convertToType(data['end'], 'Date');
            }
            if (data.hasOwnProperty('data')) {
                obj['data'] = ApiClient.convertToType(data['data'], [LogData]);
            }
            if (data.hasOwnProperty('max')) {
                obj['max'] = ApiClient.convertToType(data['max'], 'Number');
            }
            if (data.hasOwnProperty('grandTotal')) {
                obj['grandTotal'] = ApiClient.convertToType(data['grandTotal'], 'Number');
            }
            if (data.hasOwnProperty('total')) {
                obj['total'] = ApiClient.convertToType(data['total'], 'Number');
            }
            if (data.hasOwnProperty('success')) {
                obj['success'] = ApiClient.convertToType(data['success'], 'Boolean');
            }
            if (data.hasOwnProperty('count')) {
                obj['count'] = ApiClient.convertToType(data['count'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/LogSort} sort
 */
Log.prototype['sort'] = undefined;

/**
 * @member {Number} offset
 */
Log.prototype['offset'] = undefined;

/**
 * @member {Date} start
 */
Log.prototype['start'] = undefined;

/**
 * @member {Date} end
 */
Log.prototype['end'] = undefined;

/**
 * @member {Array.<module:model/LogData>} data
 */
Log.prototype['data'] = undefined;

/**
 * @member {Number} max
 */
Log.prototype['max'] = undefined;

/**
 * @member {Number} grandTotal
 */
Log.prototype['grandTotal'] = undefined;

/**
 * @member {Number} total
 */
Log.prototype['total'] = undefined;

/**
 * @member {Boolean} success
 */
Log.prototype['success'] = undefined;

/**
 * @member {Number} count
 */
Log.prototype['count'] = undefined;






export default Log;
