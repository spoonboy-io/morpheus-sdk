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
import ActivityActivityInnerUser from './ActivityActivityInnerUser';
import CheckCheckType from './CheckCheckType';
import CheckGroupInstance from './CheckGroupInstance';
import UpdateBlueprintPermissionsRequestResourcePermissionSitesInner from './UpdateBlueprintPermissionsRequestResourcePermissionSitesInner';

/**
 * The CheckGroup model module.
 * @module model/CheckGroup
 * @version 6.1.1
 */
class CheckGroup {
    /**
     * Constructs a new <code>CheckGroup</code>.
     * @alias module:model/CheckGroup
     */
    constructor() { 
        
        CheckGroup.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CheckGroup</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CheckGroup} obj Optional instance to populate.
     * @return {module:model/CheckGroup} The populated <code>CheckGroup</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CheckGroup();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('instance')) {
                obj['instance'] = CheckGroupInstance.constructFromObject(data['instance']);
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('inUptime')) {
                obj['inUptime'] = ApiClient.convertToType(data['inUptime'], 'Boolean');
            }
            if (data.hasOwnProperty('lastCheckStatus')) {
                obj['lastCheckStatus'] = ApiClient.convertToType(data['lastCheckStatus'], 'String');
            }
            if (data.hasOwnProperty('lastWarningDate')) {
                obj['lastWarningDate'] = ApiClient.convertToType(data['lastWarningDate'], 'Date');
            }
            if (data.hasOwnProperty('lastErrorDate')) {
                obj['lastErrorDate'] = ApiClient.convertToType(data['lastErrorDate'], 'Date');
            }
            if (data.hasOwnProperty('lastSuccessDate')) {
                obj['lastSuccessDate'] = ApiClient.convertToType(data['lastSuccessDate'], 'Date');
            }
            if (data.hasOwnProperty('lastRunDate')) {
                obj['lastRunDate'] = ApiClient.convertToType(data['lastRunDate'], 'Date');
            }
            if (data.hasOwnProperty('lastError')) {
                obj['lastError'] = ApiClient.convertToType(data['lastError'], 'String');
            }
            if (data.hasOwnProperty('outageTime')) {
                obj['outageTime'] = ApiClient.convertToType(data['outageTime'], 'Number');
            }
            if (data.hasOwnProperty('lastTimer')) {
                obj['lastTimer'] = ApiClient.convertToType(data['lastTimer'], 'Number');
            }
            if (data.hasOwnProperty('health')) {
                obj['health'] = ApiClient.convertToType(data['health'], 'Number');
            }
            if (data.hasOwnProperty('history')) {
                obj['history'] = ApiClient.convertToType(data['history'], 'String');
            }
            if (data.hasOwnProperty('minHappy')) {
                obj['minHappy'] = ApiClient.convertToType(data['minHappy'], 'Number');
            }
            if (data.hasOwnProperty('lastMetric')) {
                obj['lastMetric'] = ApiClient.convertToType(data['lastMetric'], 'String');
            }
            if (data.hasOwnProperty('severity')) {
                obj['severity'] = ApiClient.convertToType(data['severity'], 'String');
            }
            if (data.hasOwnProperty('createIncident')) {
                obj['createIncident'] = ApiClient.convertToType(data['createIncident'], 'Boolean');
            }
            if (data.hasOwnProperty('muted')) {
                obj['muted'] = ApiClient.convertToType(data['muted'], 'Boolean');
            }
            if (data.hasOwnProperty('createdBy')) {
                obj['createdBy'] = ActivityActivityInnerUser.constructFromObject(data['createdBy']);
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('availability')) {
                obj['availability'] = ApiClient.convertToType(data['availability'], 'Number');
            }
            if (data.hasOwnProperty('checkType')) {
                obj['checkType'] = CheckCheckType.constructFromObject(data['checkType']);
            }
            if (data.hasOwnProperty('checks')) {
                obj['checks'] = ApiClient.convertToType(data['checks'], ['Number']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>CheckGroup</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>CheckGroup</code>.
     */
    static validateJSON(data) {
        // validate the optional field `account`
        if (data['account']) { // data not null
          UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.validateJSON(data['account']);
        }
        // validate the optional field `instance`
        if (data['instance']) { // data not null
          CheckGroupInstance.validateJSON(data['instance']);
        }
        // ensure the json data is a string
        if (data['name'] && !(typeof data['name'] === 'string' || data['name'] instanceof String)) {
            throw new Error("Expected the field `name` to be a primitive type in the JSON string but got " + data['name']);
        }
        // ensure the json data is a string
        if (data['description'] && !(typeof data['description'] === 'string' || data['description'] instanceof String)) {
            throw new Error("Expected the field `description` to be a primitive type in the JSON string but got " + data['description']);
        }
        // ensure the json data is a string
        if (data['lastCheckStatus'] && !(typeof data['lastCheckStatus'] === 'string' || data['lastCheckStatus'] instanceof String)) {
            throw new Error("Expected the field `lastCheckStatus` to be a primitive type in the JSON string but got " + data['lastCheckStatus']);
        }
        // ensure the json data is a string
        if (data['lastError'] && !(typeof data['lastError'] === 'string' || data['lastError'] instanceof String)) {
            throw new Error("Expected the field `lastError` to be a primitive type in the JSON string but got " + data['lastError']);
        }
        // ensure the json data is a string
        if (data['history'] && !(typeof data['history'] === 'string' || data['history'] instanceof String)) {
            throw new Error("Expected the field `history` to be a primitive type in the JSON string but got " + data['history']);
        }
        // ensure the json data is a string
        if (data['lastMetric'] && !(typeof data['lastMetric'] === 'string' || data['lastMetric'] instanceof String)) {
            throw new Error("Expected the field `lastMetric` to be a primitive type in the JSON string but got " + data['lastMetric']);
        }
        // ensure the json data is a string
        if (data['severity'] && !(typeof data['severity'] === 'string' || data['severity'] instanceof String)) {
            throw new Error("Expected the field `severity` to be a primitive type in the JSON string but got " + data['severity']);
        }
        // validate the optional field `createdBy`
        if (data['createdBy']) { // data not null
          ActivityActivityInnerUser.validateJSON(data['createdBy']);
        }
        // validate the optional field `checkType`
        if (data['checkType']) { // data not null
          CheckCheckType.validateJSON(data['checkType']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['checks'])) {
            throw new Error("Expected the field `checks` to be an array in the JSON data but got " + data['checks']);
        }

        return true;
    }


}



/**
 * @member {Number} id
 */
CheckGroup.prototype['id'] = undefined;

/**
 * @member {module:model/UpdateBlueprintPermissionsRequestResourcePermissionSitesInner} account
 */
CheckGroup.prototype['account'] = undefined;

/**
 * @member {module:model/CheckGroupInstance} instance
 */
CheckGroup.prototype['instance'] = undefined;

/**
 * @member {String} name
 */
CheckGroup.prototype['name'] = undefined;

/**
 * @member {String} description
 */
CheckGroup.prototype['description'] = undefined;

/**
 * @member {Boolean} inUptime
 */
CheckGroup.prototype['inUptime'] = undefined;

/**
 * @member {String} lastCheckStatus
 */
CheckGroup.prototype['lastCheckStatus'] = undefined;

/**
 * @member {Date} lastWarningDate
 */
CheckGroup.prototype['lastWarningDate'] = undefined;

/**
 * @member {Date} lastErrorDate
 */
CheckGroup.prototype['lastErrorDate'] = undefined;

/**
 * @member {Date} lastSuccessDate
 */
CheckGroup.prototype['lastSuccessDate'] = undefined;

/**
 * @member {Date} lastRunDate
 */
CheckGroup.prototype['lastRunDate'] = undefined;

/**
 * @member {String} lastError
 */
CheckGroup.prototype['lastError'] = undefined;

/**
 * @member {Number} outageTime
 */
CheckGroup.prototype['outageTime'] = undefined;

/**
 * @member {Number} lastTimer
 */
CheckGroup.prototype['lastTimer'] = undefined;

/**
 * @member {Number} health
 */
CheckGroup.prototype['health'] = undefined;

/**
 * @member {String} history
 */
CheckGroup.prototype['history'] = undefined;

/**
 * @member {Number} minHappy
 */
CheckGroup.prototype['minHappy'] = undefined;

/**
 * @member {String} lastMetric
 */
CheckGroup.prototype['lastMetric'] = undefined;

/**
 * @member {String} severity
 */
CheckGroup.prototype['severity'] = undefined;

/**
 * @member {Boolean} createIncident
 */
CheckGroup.prototype['createIncident'] = undefined;

/**
 * @member {Boolean} muted
 */
CheckGroup.prototype['muted'] = undefined;

/**
 * @member {module:model/ActivityActivityInnerUser} createdBy
 */
CheckGroup.prototype['createdBy'] = undefined;

/**
 * @member {Date} dateCreated
 */
CheckGroup.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
CheckGroup.prototype['lastUpdated'] = undefined;

/**
 * @member {Number} availability
 */
CheckGroup.prototype['availability'] = undefined;

/**
 * @member {module:model/CheckCheckType} checkType
 */
CheckGroup.prototype['checkType'] = undefined;

/**
 * @member {Array.<Number>} checks
 */
CheckGroup.prototype['checks'] = undefined;






export default CheckGroup;

