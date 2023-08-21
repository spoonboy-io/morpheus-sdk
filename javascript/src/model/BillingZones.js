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
import BillingComputeServers from './BillingComputeServers';
import BillingInstances from './BillingInstances';
import BillingLoadBalancers from './BillingLoadBalancers';
import BillingSnapshots from './BillingSnapshots';
import BillingVirtualImages from './BillingVirtualImages';

/**
 * The BillingZones model module.
 * @module model/BillingZones
 * @version 6.2.1
 */
class BillingZones {
    /**
     * Constructs a new <code>BillingZones</code>.
     * @alias module:model/BillingZones
     */
    constructor() { 
        
        BillingZones.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingZones</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingZones} obj Optional instance to populate.
     * @return {module:model/BillingZones} The populated <code>BillingZones</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingZones();

            if (data.hasOwnProperty('zoneName')) {
                obj['zoneName'] = ApiClient.convertToType(data['zoneName'], 'String');
            }
            if (data.hasOwnProperty('zoneId')) {
                obj['zoneId'] = ApiClient.convertToType(data['zoneId'], 'Number');
            }
            if (data.hasOwnProperty('zoneUUID')) {
                obj['zoneUUID'] = ApiClient.convertToType(data['zoneUUID'], 'String');
            }
            if (data.hasOwnProperty('zoneCode')) {
                obj['zoneCode'] = ApiClient.convertToType(data['zoneCode'], 'String');
            }
            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('endDate')) {
                obj['endDate'] = ApiClient.convertToType(data['endDate'], 'Date');
            }
            if (data.hasOwnProperty('priceUnit')) {
                obj['priceUnit'] = ApiClient.convertToType(data['priceUnit'], 'String');
            }
            if (data.hasOwnProperty('computeServers')) {
                obj['computeServers'] = BillingComputeServers.constructFromObject(data['computeServers']);
            }
            if (data.hasOwnProperty('instances')) {
                obj['instances'] = BillingInstances.constructFromObject(data['instances']);
            }
            if (data.hasOwnProperty('discoveredServers')) {
                obj['discoveredServers'] = BillingComputeServers.constructFromObject(data['discoveredServers']);
            }
            if (data.hasOwnProperty('loadBalancers')) {
                obj['loadBalancers'] = BillingLoadBalancers.constructFromObject(data['loadBalancers']);
            }
            if (data.hasOwnProperty('virtualImages')) {
                obj['virtualImages'] = BillingVirtualImages.constructFromObject(data['virtualImages']);
            }
            if (data.hasOwnProperty('snapshots')) {
                obj['snapshots'] = BillingSnapshots.constructFromObject(data['snapshots']);
            }
            if (data.hasOwnProperty('price')) {
                obj['price'] = ApiClient.convertToType(data['price'], 'Number');
            }
            if (data.hasOwnProperty('cost')) {
                obj['cost'] = ApiClient.convertToType(data['cost'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {String} zoneName
 */
BillingZones.prototype['zoneName'] = undefined;

/**
 * @member {Number} zoneId
 */
BillingZones.prototype['zoneId'] = undefined;

/**
 * @member {String} zoneUUID
 */
BillingZones.prototype['zoneUUID'] = undefined;

/**
 * @member {String} zoneCode
 */
BillingZones.prototype['zoneCode'] = undefined;

/**
 * @member {Date} startDate
 */
BillingZones.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
BillingZones.prototype['endDate'] = undefined;

/**
 * @member {String} priceUnit
 */
BillingZones.prototype['priceUnit'] = undefined;

/**
 * @member {module:model/BillingComputeServers} computeServers
 */
BillingZones.prototype['computeServers'] = undefined;

/**
 * @member {module:model/BillingInstances} instances
 */
BillingZones.prototype['instances'] = undefined;

/**
 * @member {module:model/BillingComputeServers} discoveredServers
 */
BillingZones.prototype['discoveredServers'] = undefined;

/**
 * @member {module:model/BillingLoadBalancers} loadBalancers
 */
BillingZones.prototype['loadBalancers'] = undefined;

/**
 * @member {module:model/BillingVirtualImages} virtualImages
 */
BillingZones.prototype['virtualImages'] = undefined;

/**
 * @member {module:model/BillingSnapshots} snapshots
 */
BillingZones.prototype['snapshots'] = undefined;

/**
 * @member {Number} price
 */
BillingZones.prototype['price'] = undefined;

/**
 * @member {Number} cost
 */
BillingZones.prototype['cost'] = undefined;






export default BillingZones;

