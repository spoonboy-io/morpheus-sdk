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
import CheckWebCheckType from './CheckWebCheckType';
import CheckWebConfig from './CheckWebConfig';

/**
 * The CheckWeb model module.
 * @module model/CheckWeb
 * @version 6.2.1
 */
class CheckWeb {
    /**
     * Constructs a new <code>CheckWeb</code>.
     * Web check type allows you to perform a standard web request and validate the response came back successfully.  Additionally, you can check for matching text within the result. 
     * @alias module:model/CheckWeb
     */
    constructor() { 
        
        CheckWeb.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CheckWeb</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CheckWeb} obj Optional instance to populate.
     * @return {module:model/CheckWeb} The populated <code>CheckWeb</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CheckWeb();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('checkType')) {
                obj['checkType'] = CheckWebCheckType.constructFromObject(data['checkType']);
            }
            if (data.hasOwnProperty('checkInterval')) {
                obj['checkInterval'] = ApiClient.convertToType(data['checkInterval'], 'Number');
            }
            if (data.hasOwnProperty('inUptime')) {
                obj['inUptime'] = ApiClient.convertToType(data['inUptime'], 'Boolean');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('severity')) {
                obj['severity'] = ApiClient.convertToType(data['severity'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], CheckWebConfig);
            }
        }
        return obj;
    }


}

/**
 * Unique name scoped to your account for the check
 * @member {String} name
 */
CheckWeb.prototype['name'] = undefined;

/**
 * Optional description field
 * @member {String} description
 */
CheckWeb.prototype['description'] = undefined;

/**
 * @member {module:model/CheckWebCheckType} checkType
 */
CheckWeb.prototype['checkType'] = undefined;

/**
 * Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan)
 * @member {Number} checkInterval
 * @default 300
 */
CheckWeb.prototype['checkInterval'] = 300;

/**
 * Used to determine if check should affect account wide availability calculations
 * @member {Boolean} inUptime
 * @default true
 */
CheckWeb.prototype['inUptime'] = true;

/**
 * Used to determine if check should be scheduled to execute
 * @member {Boolean} active
 * @default true
 */
CheckWeb.prototype['active'] = true;

/**
 * Severity level threshold for sending notifications.
 * @member {module:model/CheckWeb.SeverityEnum} severity
 * @default 'critical'
 */
CheckWeb.prototype['severity'] = 'critical';

/**
 * @member {module:model/CheckWebConfig} config
 */
CheckWeb.prototype['config'] = undefined;





/**
 * Allowed values for the <code>severity</code> property.
 * @enum {String}
 * @readonly
 */
CheckWeb['SeverityEnum'] = {

    /**
     * value: "info"
     * @const
     */
    "info": "info",

    /**
     * value: "warning"
     * @const
     */
    "warning": "warning",

    /**
     * value: "critical"
     * @const
     */
    "critical": "critical"
};



export default CheckWeb;
