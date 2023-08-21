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
import DashboardLogStatsData from './DashboardLogStatsData';

/**
 * The DashboardLogStats model module.
 * @module model/DashboardLogStats
 * @version 6.2.1
 */
class DashboardLogStats {
    /**
     * Constructs a new <code>DashboardLogStats</code>.
     * @alias module:model/DashboardLogStats
     */
    constructor() { 
        
        DashboardLogStats.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>DashboardLogStats</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/DashboardLogStats} obj Optional instance to populate.
     * @return {module:model/DashboardLogStats} The populated <code>DashboardLogStats</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new DashboardLogStats();

            if (data.hasOwnProperty('success')) {
                obj['success'] = ApiClient.convertToType(data['success'], 'Boolean');
            }
            if (data.hasOwnProperty('data')) {
                obj['data'] = ApiClient.convertToType(data['data'], [DashboardLogStatsData]);
            }
            if (data.hasOwnProperty('startMs')) {
                obj['startMs'] = ApiClient.convertToType(data['startMs'], 'Number');
            }
            if (data.hasOwnProperty('earliest')) {
                obj['earliest'] = ApiClient.convertToType(data['earliest'], 'Number');
            }
            if (data.hasOwnProperty('endMs')) {
                obj['endMs'] = ApiClient.convertToType(data['endMs'], 'Number');
            }
            if (data.hasOwnProperty('interval')) {
                obj['interval'] = ApiClient.convertToType(data['interval'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} success
 */
DashboardLogStats.prototype['success'] = undefined;

/**
 * @member {Array.<module:model/DashboardLogStatsData>} data
 */
DashboardLogStats.prototype['data'] = undefined;

/**
 * @member {Number} startMs
 */
DashboardLogStats.prototype['startMs'] = undefined;

/**
 * @member {Number} earliest
 */
DashboardLogStats.prototype['earliest'] = undefined;

/**
 * @member {Number} endMs
 */
DashboardLogStats.prototype['endMs'] = undefined;

/**
 * @member {Number} interval
 */
DashboardLogStats.prototype['interval'] = undefined;






export default DashboardLogStats;

