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
import PriceSetPrices from './PriceSetPrices';

/**
 * The PriceSet model module.
 * @module model/PriceSet
 * @version 6.2.1
 */
class PriceSet {
    /**
     * Constructs a new <code>PriceSet</code>.
     * @alias module:model/PriceSet
     */
    constructor() { 
        
        PriceSet.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>PriceSet</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/PriceSet} obj Optional instance to populate.
     * @return {module:model/PriceSet} The populated <code>PriceSet</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new PriceSet();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('priceUnit')) {
                obj['priceUnit'] = ApiClient.convertToType(data['priceUnit'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('regionCode')) {
                obj['regionCode'] = ApiClient.convertToType(data['regionCode'], 'String');
            }
            if (data.hasOwnProperty('systemCreated')) {
                obj['systemCreated'] = ApiClient.convertToType(data['systemCreated'], 'Boolean');
            }
            if (data.hasOwnProperty('zone')) {
                obj['zone'] = ApiClient.convertToType(data['zone'], 'String');
            }
            if (data.hasOwnProperty('zonePool')) {
                obj['zonePool'] = ApiClient.convertToType(data['zonePool'], 'String');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = ApiClient.convertToType(data['account'], 'String');
            }
            if (data.hasOwnProperty('prices')) {
                obj['prices'] = ApiClient.convertToType(data['prices'], [PriceSetPrices]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
PriceSet.prototype['id'] = undefined;

/**
 * @member {String} name
 */
PriceSet.prototype['name'] = undefined;

/**
 * @member {String} code
 */
PriceSet.prototype['code'] = undefined;

/**
 * @member {Boolean} active
 */
PriceSet.prototype['active'] = undefined;

/**
 * @member {String} priceUnit
 */
PriceSet.prototype['priceUnit'] = undefined;

/**
 * @member {String} type
 */
PriceSet.prototype['type'] = undefined;

/**
 * @member {String} regionCode
 */
PriceSet.prototype['regionCode'] = undefined;

/**
 * @member {Boolean} systemCreated
 */
PriceSet.prototype['systemCreated'] = undefined;

/**
 * @member {String} zone
 */
PriceSet.prototype['zone'] = undefined;

/**
 * @member {String} zonePool
 */
PriceSet.prototype['zonePool'] = undefined;

/**
 * @member {String} account
 */
PriceSet.prototype['account'] = undefined;

/**
 * @member {Array.<module:model/PriceSetPrices>} prices
 */
PriceSet.prototype['prices'] = undefined;






export default PriceSet;

