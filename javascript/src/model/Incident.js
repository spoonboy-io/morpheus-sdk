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
import ApiBlueprintsIdUpdatePermissionsResourcePermissionSites from './ApiBlueprintsIdUpdatePermissionsResourcePermissionSites';
import Check from './Check';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';

/**
 * The Incident model module.
 * @module model/Incident
 * @version 6.2.1
 */
class Incident {
    /**
     * Constructs a new <code>Incident</code>.
     * @alias module:model/Incident
     */
    constructor() { 
        
        Incident.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Incident</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Incident} obj Optional instance to populate.
     * @return {module:model/Incident} The populated <code>Incident</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Incident();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('app')) {
                obj['app'] = ApiClient.convertToType(data['app'], 'String');
            }
            if (data.hasOwnProperty('autoClose')) {
                obj['autoClose'] = ApiClient.convertToType(data['autoClose'], 'Boolean');
            }
            if (data.hasOwnProperty('channelId')) {
                obj['channelId'] = ApiClient.convertToType(data['channelId'], 'String');
            }
            if (data.hasOwnProperty('checkGroups')) {
                obj['checkGroups'] = ApiClient.convertToType(data['checkGroups'], [InlineResponse20040AppDeployInstance]);
            }
            if (data.hasOwnProperty('checks')) {
                obj['checks'] = ApiClient.convertToType(data['checks'], [Check]);
            }
            if (data.hasOwnProperty('comment')) {
                obj['comment'] = ApiClient.convertToType(data['comment'], 'String');
            }
            if (data.hasOwnProperty('displayName')) {
                obj['displayName'] = ApiClient.convertToType(data['displayName'], 'String');
            }
            if (data.hasOwnProperty('duration')) {
                obj['duration'] = ApiClient.convertToType(data['duration'], 'String');
            }
            if (data.hasOwnProperty('endDate')) {
                obj['endDate'] = ApiClient.convertToType(data['endDate'], 'Date');
            }
            if (data.hasOwnProperty('inUptime')) {
                obj['inUptime'] = ApiClient.convertToType(data['inUptime'], 'Boolean');
            }
            if (data.hasOwnProperty('muted')) {
                obj['muted'] = ApiClient.convertToType(data['muted'], 'Boolean');
            }
            if (data.hasOwnProperty('lastCheckTime')) {
                obj['lastCheckTime'] = ApiClient.convertToType(data['lastCheckTime'], 'Date');
            }
            if (data.hasOwnProperty('lastError')) {
                obj['lastError'] = ApiClient.convertToType(data['lastError'], 'String');
            }
            if (data.hasOwnProperty('lastMessage')) {
                obj['lastMessage'] = ApiClient.convertToType(data['lastMessage'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('resolution')) {
                obj['resolution'] = ApiClient.convertToType(data['resolution'], 'String');
            }
            if (data.hasOwnProperty('severity')) {
                obj['severity'] = ApiClient.convertToType(data['severity'], 'String');
            }
            if (data.hasOwnProperty('severityId')) {
                obj['severityId'] = ApiClient.convertToType(data['severityId'], 'Number');
            }
            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
Incident.prototype['id'] = undefined;

/**
 * @member {module:model/ApiBlueprintsIdUpdatePermissionsResourcePermissionSites} account
 */
Incident.prototype['account'] = undefined;

/**
 * @member {String} app
 */
Incident.prototype['app'] = undefined;

/**
 * @member {Boolean} autoClose
 */
Incident.prototype['autoClose'] = undefined;

/**
 * @member {String} channelId
 */
Incident.prototype['channelId'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse20040AppDeployInstance>} checkGroups
 */
Incident.prototype['checkGroups'] = undefined;

/**
 * @member {Array.<module:model/Check>} checks
 */
Incident.prototype['checks'] = undefined;

/**
 * @member {String} comment
 */
Incident.prototype['comment'] = undefined;

/**
 * @member {String} displayName
 */
Incident.prototype['displayName'] = undefined;

/**
 * @member {String} duration
 */
Incident.prototype['duration'] = undefined;

/**
 * @member {Date} endDate
 */
Incident.prototype['endDate'] = undefined;

/**
 * @member {Boolean} inUptime
 */
Incident.prototype['inUptime'] = undefined;

/**
 * @member {Boolean} muted
 */
Incident.prototype['muted'] = undefined;

/**
 * @member {Date} lastCheckTime
 */
Incident.prototype['lastCheckTime'] = undefined;

/**
 * @member {String} lastError
 */
Incident.prototype['lastError'] = undefined;

/**
 * @member {String} lastMessage
 */
Incident.prototype['lastMessage'] = undefined;

/**
 * @member {String} name
 */
Incident.prototype['name'] = undefined;

/**
 * @member {String} resolution
 */
Incident.prototype['resolution'] = undefined;

/**
 * @member {String} severity
 */
Incident.prototype['severity'] = undefined;

/**
 * @member {Number} severityId
 */
Incident.prototype['severityId'] = undefined;

/**
 * @member {Date} startDate
 */
Incident.prototype['startDate'] = undefined;

/**
 * @member {String} status
 */
Incident.prototype['status'] = undefined;

/**
 * @member {String} visibility
 */
Incident.prototype['visibility'] = undefined;






export default Incident;

