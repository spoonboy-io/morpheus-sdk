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
import BillingZonesInnerComputeServers from './BillingZonesInnerComputeServers';
import BillingZonesInnerInstances from './BillingZonesInnerInstances';
import BillingZonesInnerLoadBalancers from './BillingZonesInnerLoadBalancers';
import BillingZonesInnerSnapshots from './BillingZonesInnerSnapshots';
import BillingZonesInnerVirtualImages from './BillingZonesInnerVirtualImages';

/**
 * The BillingZonesInner model module.
 * @module model/BillingZonesInner
 * @version 6.1.1
 */
class BillingZonesInner {
    /**
     * Constructs a new <code>BillingZonesInner</code>.
     * @alias module:model/BillingZonesInner
     */
    constructor() { 
        
        BillingZonesInner.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingZonesInner</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingZonesInner} obj Optional instance to populate.
     * @return {module:model/BillingZonesInner} The populated <code>BillingZonesInner</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingZonesInner();

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
                obj['computeServers'] = BillingZonesInnerComputeServers.constructFromObject(data['computeServers']);
            }
            if (data.hasOwnProperty('instances')) {
                obj['instances'] = BillingZonesInnerInstances.constructFromObject(data['instances']);
            }
            if (data.hasOwnProperty('discoveredServers')) {
                obj['discoveredServers'] = BillingZonesInnerComputeServers.constructFromObject(data['discoveredServers']);
            }
            if (data.hasOwnProperty('loadBalancers')) {
                obj['loadBalancers'] = BillingZonesInnerLoadBalancers.constructFromObject(data['loadBalancers']);
            }
            if (data.hasOwnProperty('virtualImages')) {
                obj['virtualImages'] = BillingZonesInnerVirtualImages.constructFromObject(data['virtualImages']);
            }
            if (data.hasOwnProperty('snapshots')) {
                obj['snapshots'] = BillingZonesInnerSnapshots.constructFromObject(data['snapshots']);
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

    /**
     * Validates the JSON data with respect to <code>BillingZonesInner</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>BillingZonesInner</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['zoneName'] && !(typeof data['zoneName'] === 'string' || data['zoneName'] instanceof String)) {
            throw new Error("Expected the field `zoneName` to be a primitive type in the JSON string but got " + data['zoneName']);
        }
        // ensure the json data is a string
        if (data['zoneUUID'] && !(typeof data['zoneUUID'] === 'string' || data['zoneUUID'] instanceof String)) {
            throw new Error("Expected the field `zoneUUID` to be a primitive type in the JSON string but got " + data['zoneUUID']);
        }
        // ensure the json data is a string
        if (data['zoneCode'] && !(typeof data['zoneCode'] === 'string' || data['zoneCode'] instanceof String)) {
            throw new Error("Expected the field `zoneCode` to be a primitive type in the JSON string but got " + data['zoneCode']);
        }
        // ensure the json data is a string
        if (data['priceUnit'] && !(typeof data['priceUnit'] === 'string' || data['priceUnit'] instanceof String)) {
            throw new Error("Expected the field `priceUnit` to be a primitive type in the JSON string but got " + data['priceUnit']);
        }
        // validate the optional field `computeServers`
        if (data['computeServers']) { // data not null
          BillingZonesInnerComputeServers.validateJSON(data['computeServers']);
        }
        // validate the optional field `instances`
        if (data['instances']) { // data not null
          BillingZonesInnerInstances.validateJSON(data['instances']);
        }
        // validate the optional field `discoveredServers`
        if (data['discoveredServers']) { // data not null
          BillingZonesInnerComputeServers.validateJSON(data['discoveredServers']);
        }
        // validate the optional field `loadBalancers`
        if (data['loadBalancers']) { // data not null
          BillingZonesInnerLoadBalancers.validateJSON(data['loadBalancers']);
        }
        // validate the optional field `virtualImages`
        if (data['virtualImages']) { // data not null
          BillingZonesInnerVirtualImages.validateJSON(data['virtualImages']);
        }
        // validate the optional field `snapshots`
        if (data['snapshots']) { // data not null
          BillingZonesInnerSnapshots.validateJSON(data['snapshots']);
        }

        return true;
    }


}



/**
 * @member {String} zoneName
 */
BillingZonesInner.prototype['zoneName'] = undefined;

/**
 * @member {Number} zoneId
 */
BillingZonesInner.prototype['zoneId'] = undefined;

/**
 * @member {String} zoneUUID
 */
BillingZonesInner.prototype['zoneUUID'] = undefined;

/**
 * @member {String} zoneCode
 */
BillingZonesInner.prototype['zoneCode'] = undefined;

/**
 * @member {Date} startDate
 */
BillingZonesInner.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
BillingZonesInner.prototype['endDate'] = undefined;

/**
 * @member {String} priceUnit
 */
BillingZonesInner.prototype['priceUnit'] = undefined;

/**
 * @member {module:model/BillingZonesInnerComputeServers} computeServers
 */
BillingZonesInner.prototype['computeServers'] = undefined;

/**
 * @member {module:model/BillingZonesInnerInstances} instances
 */
BillingZonesInner.prototype['instances'] = undefined;

/**
 * @member {module:model/BillingZonesInnerComputeServers} discoveredServers
 */
BillingZonesInner.prototype['discoveredServers'] = undefined;

/**
 * @member {module:model/BillingZonesInnerLoadBalancers} loadBalancers
 */
BillingZonesInner.prototype['loadBalancers'] = undefined;

/**
 * @member {module:model/BillingZonesInnerVirtualImages} virtualImages
 */
BillingZonesInner.prototype['virtualImages'] = undefined;

/**
 * @member {module:model/BillingZonesInnerSnapshots} snapshots
 */
BillingZonesInner.prototype['snapshots'] = undefined;

/**
 * @member {Number} price
 */
BillingZonesInner.prototype['price'] = undefined;

/**
 * @member {Number} cost
 */
BillingZonesInner.prototype['cost'] = undefined;






export default BillingZonesInner;

