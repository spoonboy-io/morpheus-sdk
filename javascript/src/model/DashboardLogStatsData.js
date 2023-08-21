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
 * The DashboardLogStatsData model module.
 * @module model/DashboardLogStatsData
 * @version 6.2.1
 */
class DashboardLogStatsData {
    /**
     * Constructs a new <code>DashboardLogStatsData</code>.
     * @alias module:model/DashboardLogStatsData
     */
    constructor() { 
        
        DashboardLogStatsData.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>DashboardLogStatsData</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DashboardLogStatsData} obj Optional instance to populate.
     * @return {module:model/DashboardLogStatsData} The populated <code>DashboardLogStatsData</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DashboardLogStatsData();

            if (data.hasOwnProperty('key')) {
                obj['key'] = ApiClient.convertToType(data['key'], 'String');
            }
            if (data.hasOwnProperty('values')) {
                obj['values'] = ApiClient.convertToType(data['values'], {'String': 'Number'});
            }
            if (data.hasOwnProperty('count')) {
                obj['count'] = ApiClient.convertToType(data['count'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {String} key
 */
DashboardLogStatsData.prototype['key'] = undefined;

/**
 * @member {Object.<String, Number>} values
 */
DashboardLogStatsData.prototype['values'] = undefined;

/**
 * @member {Number} count
 */
DashboardLogStatsData.prototype['count'] = undefined;






export default DashboardLogStatsData;
