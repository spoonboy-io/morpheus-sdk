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
 * The ApiMonitoringIncidentsIncident model module.
 * @module model/ApiMonitoringIncidentsIncident
 * @version 6.2.1
 */
class ApiMonitoringIncidentsIncident {
    /**
     * Constructs a new <code>ApiMonitoringIncidentsIncident</code>.
     * Payload for creating a new incident
     * @alias module:model/ApiMonitoringIncidentsIncident
     * @param name {String} Set display name
     */
    constructor(name) { 
        
        ApiMonitoringIncidentsIncident.initialize(this, name);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name) { 
        obj['name'] = name;
    }

    /**
     * Constructs a <code>ApiMonitoringIncidentsIncident</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiMonitoringIncidentsIncident} obj Optional instance to populate.
     * @return {module:model/ApiMonitoringIncidentsIncident} The populated <code>ApiMonitoringIncidentsIncident</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiMonitoringIncidentsIncident();

            if (data.hasOwnProperty('resolution')) {
                obj['resolution'] = ApiClient.convertToType(data['resolution'], 'String');
            }
            if (data.hasOwnProperty('comment')) {
                obj['comment'] = ApiClient.convertToType(data['comment'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('severity')) {
                obj['severity'] = ApiClient.convertToType(data['severity'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('endDate')) {
                obj['endDate'] = ApiClient.convertToType(data['endDate'], 'Date');
            }
            if (data.hasOwnProperty('inUptime')) {
                obj['inUptime'] = ApiClient.convertToType(data['inUptime'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * Description of the resolution to this incident
 * @member {String} resolution
 */
ApiMonitoringIncidentsIncident.prototype['resolution'] = undefined;

/**
 * Comment on this incident, updates summary field
 * @member {String} comment
 */
ApiMonitoringIncidentsIncident.prototype['comment'] = undefined;

/**
 * Set status
 * @member {module:model/ApiMonitoringIncidentsIncident.StatusEnum} status
 */
ApiMonitoringIncidentsIncident.prototype['status'] = undefined;

/**
 * Set severity
 * @member {module:model/ApiMonitoringIncidentsIncident.SeverityEnum} severity
 */
ApiMonitoringIncidentsIncident.prototype['severity'] = undefined;

/**
 * Set display name
 * @member {String} name
 */
ApiMonitoringIncidentsIncident.prototype['name'] = undefined;

/**
 * Set start time
 * @member {Date} startDate
 */
ApiMonitoringIncidentsIncident.prototype['startDate'] = undefined;

/**
 * Set start time
 * @member {Date} endDate
 */
ApiMonitoringIncidentsIncident.prototype['endDate'] = undefined;

/**
 * Set 'In Availability'
 * @member {Boolean} inUptime
 */
ApiMonitoringIncidentsIncident.prototype['inUptime'] = undefined;





/**
 * Allowed values for the <code>status</code> property.
 * @enum {String}
 * @readonly
 */
ApiMonitoringIncidentsIncident['StatusEnum'] = {

    /**
     * value: "open"
     * @const
     */
    "open": "open",

    /**
     * value: "closed"
     * @const
     */
    "closed": "closed"
};


/**
 * Allowed values for the <code>severity</code> property.
 * @enum {String}
 * @readonly
 */
ApiMonitoringIncidentsIncident['SeverityEnum'] = {

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



export default ApiMonitoringIncidentsIncident;
