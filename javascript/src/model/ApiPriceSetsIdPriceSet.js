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
import ApiPriceSetsPriceSetZone from './ApiPriceSetsPriceSetZone';
import ApiPriceSetsPriceSetZonePool from './ApiPriceSetsPriceSetZonePool';

/**
 * The ApiPriceSetsIdPriceSet model module.
 * @module model/ApiPriceSetsIdPriceSet
 * @version 6.2.1
 */
class ApiPriceSetsIdPriceSet {
    /**
     * Constructs a new <code>ApiPriceSetsIdPriceSet</code>.
     * @alias module:model/ApiPriceSetsIdPriceSet
     */
    constructor() { 
        
        ApiPriceSetsIdPriceSet.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ApiPriceSetsIdPriceSet</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiPriceSetsIdPriceSet} obj Optional instance to populate.
     * @return {module:model/ApiPriceSetsIdPriceSet} The populated <code>ApiPriceSetsIdPriceSet</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiPriceSetsIdPriceSet();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('regionCode')) {
                obj['regionCode'] = ApiClient.convertToType(data['regionCode'], 'String');
            }
            if (data.hasOwnProperty('zone')) {
                obj['zone'] = ApiPriceSetsPriceSetZone.constructFromObject(data['zone']);
            }
            if (data.hasOwnProperty('zonePool')) {
                obj['zonePool'] = ApiPriceSetsPriceSetZonePool.constructFromObject(data['zonePool']);
            }
            if (data.hasOwnProperty('priceUnit')) {
                obj['priceUnit'] = ApiClient.convertToType(data['priceUnit'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('prices')) {
                obj['prices'] = ApiClient.convertToType(data['prices'], ['Number']);
            }
        }
        return obj;
    }


}

/**
 * Price set name
 * @member {String} name
 */
ApiPriceSetsIdPriceSet.prototype['name'] = undefined;

/**
 * Price set code. Must be unique.
 * @member {String} code
 */
ApiPriceSetsIdPriceSet.prototype['code'] = undefined;

/**
 * Price set region code
 * @member {String} regionCode
 */
ApiPriceSetsIdPriceSet.prototype['regionCode'] = undefined;

/**
 * @member {module:model/ApiPriceSetsPriceSetZone} zone
 */
ApiPriceSetsIdPriceSet.prototype['zone'] = undefined;

/**
 * @member {module:model/ApiPriceSetsPriceSetZonePool} zonePool
 */
ApiPriceSetsIdPriceSet.prototype['zonePool'] = undefined;

/**
 * Price Unit
 * @member {module:model/ApiPriceSetsIdPriceSet.PriceUnitEnum} priceUnit
 */
ApiPriceSetsIdPriceSet.prototype['priceUnit'] = undefined;

/**
 * Price set type
 * @member {module:model/ApiPriceSetsIdPriceSet.TypeEnum} type
 */
ApiPriceSetsIdPriceSet.prototype['type'] = undefined;

/**
 * @member {Array.<Number>} prices
 */
ApiPriceSetsIdPriceSet.prototype['prices'] = undefined;





/**
 * Allowed values for the <code>priceUnit</code> property.
 * @enum {String}
 * @readonly
 */
ApiPriceSetsIdPriceSet['PriceUnitEnum'] = {

    /**
     * value: "minute"
     * @const
     */
    "minute": "minute",

    /**
     * value: "hour"
     * @const
     */
    "hour": "hour",

    /**
     * value: "day"
     * @const
     */
    "day": "day",

    /**
     * value: "month"
     * @const
     */
    "month": "month",

    /**
     * value: "year"
     * @const
     */
    "year": "year",

    /**
     * value: "two year"
     * @const
     */
    "two year": "two year",

    /**
     * value: "three year"
     * @const
     */
    "three year": "three year",

    /**
     * value: "four year"
     * @const
     */
    "four year": "four year",

    /**
     * value: "five year"
     * @const
     */
    "five year": "five year"
};


/**
 * Allowed values for the <code>type</code> property.
 * @enum {String}
 * @readonly
 */
ApiPriceSetsIdPriceSet['TypeEnum'] = {

    /**
     * value: "fixed"
     * @const
     */
    "fixed": "fixed",

    /**
     * value: "compute_plus_storage"
     * @const
     */
    "compute_plus_storage": "compute_plus_storage",

    /**
     * value: "component"
     * @const
     */
    "component": "component",

    /**
     * value: "load_balancer"
     * @const
     */
    "load_balancer": "load_balancer",

    /**
     * value: "snapshot"
     * @const
     */
    "snapshot": "snapshot",

    /**
     * value: "virtual_image"
     * @const
     */
    "virtual_image": "virtual_image",

    /**
     * value: "software_or_service"
     * @const
     */
    "software_or_service": "software_or_service"
};



export default ApiPriceSetsIdPriceSet;

