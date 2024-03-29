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
import BudgetStats from './BudgetStats';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';

/**
 * The Budget model module.
 * @module model/Budget
 * @version 6.2.1
 */
class Budget {
    /**
     * Constructs a new <code>Budget</code>.
     * @alias module:model/Budget
     */
    constructor() { 
        
        Budget.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Budget</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Budget} obj Optional instance to populate.
     * @return {module:model/Budget} The populated <code>Budget</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Budget();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = InlineResponse20040AppDeployInstance.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('refScope')) {
                obj['refScope'] = ApiClient.convertToType(data['refScope'], 'String');
            }
            if (data.hasOwnProperty('refType')) {
                obj['refType'] = ApiClient.convertToType(data['refType'], 'String');
            }
            if (data.hasOwnProperty('refId')) {
                obj['refId'] = ApiClient.convertToType(data['refId'], 'Number');
            }
            if (data.hasOwnProperty('refName')) {
                obj['refName'] = ApiClient.convertToType(data['refName'], 'String');
            }
            if (data.hasOwnProperty('period')) {
                obj['period'] = ApiClient.convertToType(data['period'], 'String');
            }
            if (data.hasOwnProperty('year')) {
                obj['year'] = ApiClient.convertToType(data['year'], 'String');
            }
            if (data.hasOwnProperty('resourceType')) {
                obj['resourceType'] = ApiClient.convertToType(data['resourceType'], 'String');
            }
            if (data.hasOwnProperty('timezone')) {
                obj['timezone'] = ApiClient.convertToType(data['timezone'], 'String');
            }
            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('endDate')) {
                obj['endDate'] = ApiClient.convertToType(data['endDate'], 'Date');
            }
            if (data.hasOwnProperty('interval')) {
                obj['interval'] = ApiClient.convertToType(data['interval'], 'String');
            }
            if (data.hasOwnProperty('costs')) {
                obj['costs'] = ApiClient.convertToType(data['costs'], ['Number']);
            }
            if (data.hasOwnProperty('isFiscal')) {
                obj['isFiscal'] = ApiClient.convertToType(data['isFiscal'], 'Boolean');
            }
            if (data.hasOwnProperty('averageCost')) {
                obj['averageCost'] = ApiClient.convertToType(data['averageCost'], 'Number');
            }
            if (data.hasOwnProperty('totalCost')) {
                obj['totalCost'] = ApiClient.convertToType(data['totalCost'], 'Number');
            }
            if (data.hasOwnProperty('currency')) {
                obj['currency'] = ApiClient.convertToType(data['currency'], 'String');
            }
            if (data.hasOwnProperty('rollover')) {
                obj['rollover'] = ApiClient.convertToType(data['rollover'], 'Boolean');
            }
            if (data.hasOwnProperty('warningLimit')) {
                obj['warningLimit'] = ApiClient.convertToType(data['warningLimit'], 'String');
            }
            if (data.hasOwnProperty('overLimit')) {
                obj['overLimit'] = ApiClient.convertToType(data['overLimit'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('internalId')) {
                obj['internalId'] = ApiClient.convertToType(data['internalId'], 'String');
            }
            if (data.hasOwnProperty('createdById')) {
                obj['createdById'] = ApiClient.convertToType(data['createdById'], 'Number');
            }
            if (data.hasOwnProperty('createdByName')) {
                obj['createdByName'] = ApiClient.convertToType(data['createdByName'], 'String');
            }
            if (data.hasOwnProperty('updatedById')) {
                obj['updatedById'] = ApiClient.convertToType(data['updatedById'], 'String');
            }
            if (data.hasOwnProperty('updatedByName')) {
                obj['updatedByName'] = ApiClient.convertToType(data['updatedByName'], 'String');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('stats')) {
                obj['stats'] = BudgetStats.constructFromObject(data['stats']);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
Budget.prototype['id'] = undefined;

/**
 * @member {String} name
 */
Budget.prototype['name'] = undefined;

/**
 * @member {String} description
 */
Budget.prototype['description'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} account
 */
Budget.prototype['account'] = undefined;

/**
 * @member {Boolean} enabled
 */
Budget.prototype['enabled'] = undefined;

/**
 * @member {String} refScope
 */
Budget.prototype['refScope'] = undefined;

/**
 * @member {String} refType
 */
Budget.prototype['refType'] = undefined;

/**
 * @member {Number} refId
 */
Budget.prototype['refId'] = undefined;

/**
 * @member {String} refName
 */
Budget.prototype['refName'] = undefined;

/**
 * @member {String} period
 */
Budget.prototype['period'] = undefined;

/**
 * @member {String} year
 */
Budget.prototype['year'] = undefined;

/**
 * @member {String} resourceType
 */
Budget.prototype['resourceType'] = undefined;

/**
 * @member {String} timezone
 */
Budget.prototype['timezone'] = undefined;

/**
 * @member {Date} startDate
 */
Budget.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
Budget.prototype['endDate'] = undefined;

/**
 * @member {String} interval
 */
Budget.prototype['interval'] = undefined;

/**
 * @member {Array.<Number>} costs
 */
Budget.prototype['costs'] = undefined;

/**
 * @member {Boolean} isFiscal
 */
Budget.prototype['isFiscal'] = undefined;

/**
 * @member {Number} averageCost
 */
Budget.prototype['averageCost'] = undefined;

/**
 * @member {Number} totalCost
 */
Budget.prototype['totalCost'] = undefined;

/**
 * @member {String} currency
 */
Budget.prototype['currency'] = undefined;

/**
 * @member {Boolean} rollover
 */
Budget.prototype['rollover'] = undefined;

/**
 * @member {String} warningLimit
 */
Budget.prototype['warningLimit'] = undefined;

/**
 * @member {String} overLimit
 */
Budget.prototype['overLimit'] = undefined;

/**
 * @member {String} externalId
 */
Budget.prototype['externalId'] = undefined;

/**
 * @member {String} internalId
 */
Budget.prototype['internalId'] = undefined;

/**
 * @member {Number} createdById
 */
Budget.prototype['createdById'] = undefined;

/**
 * @member {String} createdByName
 */
Budget.prototype['createdByName'] = undefined;

/**
 * @member {String} updatedById
 */
Budget.prototype['updatedById'] = undefined;

/**
 * @member {String} updatedByName
 */
Budget.prototype['updatedByName'] = undefined;

/**
 * @member {Date} dateCreated
 */
Budget.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
Budget.prototype['lastUpdated'] = undefined;

/**
 * @member {module:model/BudgetStats} stats
 */
Budget.prototype['stats'] = undefined;






export default Budget;

