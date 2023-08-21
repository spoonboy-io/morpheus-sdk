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
import InlineResponse20082LoadBalancerInstanceSslCert from './InlineResponse20082LoadBalancerInstanceSslCert';
import InstanceTypesInstanceTypeLayouts from './InstanceTypesInstanceTypeLayouts';

/**
 * The InstanceTypes model module.
 * @module model/InstanceTypes
 * @version 6.2.1
 */
class InstanceTypes {
    /**
     * Constructs a new <code>InstanceTypes</code>.
     * @alias module:model/InstanceTypes
     */
    constructor() { 
        
        InstanceTypes.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InstanceTypes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceTypes} obj Optional instance to populate.
     * @return {module:model/InstanceTypes} The populated <code>InstanceTypes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceTypes();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = InlineResponse20082LoadBalancerInstanceSslCert.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('provisionTypeCode')) {
                obj['provisionTypeCode'] = ApiClient.convertToType(data['provisionTypeCode'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('environmentPrefix')) {
                obj['environmentPrefix'] = ApiClient.convertToType(data['environmentPrefix'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('featured')) {
                obj['featured'] = ApiClient.convertToType(data['featured'], 'Boolean');
            }
            if (data.hasOwnProperty('versions')) {
                obj['versions'] = ApiClient.convertToType(data['versions'], ['String']);
            }
            if (data.hasOwnProperty('instanceTypeLayouts')) {
                obj['instanceTypeLayouts'] = ApiClient.convertToType(data['instanceTypeLayouts'], [InstanceTypesInstanceTypeLayouts]);
            }
            if (data.hasOwnProperty('imagePath')) {
                obj['imagePath'] = ApiClient.convertToType(data['imagePath'], 'String');
            }
            if (data.hasOwnProperty('darkImagePath')) {
                obj['darkImagePath'] = ApiClient.convertToType(data['darkImagePath'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
InstanceTypes.prototype['id'] = undefined;

/**
 * @member {module:model/InlineResponse20082LoadBalancerInstanceSslCert} account
 */
InstanceTypes.prototype['account'] = undefined;

/**
 * @member {String} name
 */
InstanceTypes.prototype['name'] = undefined;

/**
 * @member {Array.<String>} labels
 */
InstanceTypes.prototype['labels'] = undefined;

/**
 * @member {String} code
 */
InstanceTypes.prototype['code'] = undefined;

/**
 * @member {String} description
 */
InstanceTypes.prototype['description'] = undefined;

/**
 * @member {String} provisionTypeCode
 */
InstanceTypes.prototype['provisionTypeCode'] = undefined;

/**
 * @member {String} category
 */
InstanceTypes.prototype['category'] = undefined;

/**
 * @member {Boolean} active
 */
InstanceTypes.prototype['active'] = undefined;

/**
 * @member {String} environmentPrefix
 */
InstanceTypes.prototype['environmentPrefix'] = undefined;

/**
 * @member {String} visibility
 */
InstanceTypes.prototype['visibility'] = undefined;

/**
 * @member {Boolean} featured
 */
InstanceTypes.prototype['featured'] = undefined;

/**
 * @member {Array.<String>} versions
 */
InstanceTypes.prototype['versions'] = undefined;

/**
 * @member {Array.<module:model/InstanceTypesInstanceTypeLayouts>} instanceTypeLayouts
 */
InstanceTypes.prototype['instanceTypeLayouts'] = undefined;

/**
 * Logo image URL
 * @member {String} imagePath
 */
InstanceTypes.prototype['imagePath'] = undefined;

/**
 * Dark logo image URL
 * @member {String} darkImagePath
 */
InstanceTypes.prototype['darkImagePath'] = undefined;






export default InstanceTypes;

