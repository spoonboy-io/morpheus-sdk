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
import ZoneTypeOptionTypes from './ZoneTypeOptionTypes';
import ZoneTypeProvisionType from './ZoneTypeProvisionType';

/**
 * The ZoneTypeServerTypes model module.
 * @module model/ZoneTypeServerTypes
 * @version 6.2.1
 */
class ZoneTypeServerTypes {
    /**
     * Constructs a new <code>ZoneTypeServerTypes</code>.
     * @alias module:model/ZoneTypeServerTypes
     */
    constructor() { 
        
        ZoneTypeServerTypes.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ZoneTypeServerTypes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ZoneTypeServerTypes} obj Optional instance to populate.
     * @return {module:model/ZoneTypeServerTypes} The populated <code>ZoneTypeServerTypes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ZoneTypeServerTypes();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('nodeType')) {
                obj['nodeType'] = ApiClient.convertToType(data['nodeType'], 'String');
            }
            if (data.hasOwnProperty('platform')) {
                obj['platform'] = ApiClient.convertToType(data['platform'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('selectable')) {
                obj['selectable'] = ApiClient.convertToType(data['selectable'], 'Boolean');
            }
            if (data.hasOwnProperty('externalDelete')) {
                obj['externalDelete'] = ApiClient.convertToType(data['externalDelete'], 'Boolean');
            }
            if (data.hasOwnProperty('managed')) {
                obj['managed'] = ApiClient.convertToType(data['managed'], 'Boolean');
            }
            if (data.hasOwnProperty('controlPower')) {
                obj['controlPower'] = ApiClient.convertToType(data['controlPower'], 'Boolean');
            }
            if (data.hasOwnProperty('controlSuspend')) {
                obj['controlSuspend'] = ApiClient.convertToType(data['controlSuspend'], 'Boolean');
            }
            if (data.hasOwnProperty('creatable')) {
                obj['creatable'] = ApiClient.convertToType(data['creatable'], 'Boolean');
            }
            if (data.hasOwnProperty('hasAgent')) {
                obj['hasAgent'] = ApiClient.convertToType(data['hasAgent'], 'Boolean');
            }
            if (data.hasOwnProperty('vmHypervisor')) {
                obj['vmHypervisor'] = ApiClient.convertToType(data['vmHypervisor'], 'Boolean');
            }
            if (data.hasOwnProperty('containerHypervisor')) {
                obj['containerHypervisor'] = ApiClient.convertToType(data['containerHypervisor'], 'Boolean');
            }
            if (data.hasOwnProperty('bareMetalHost')) {
                obj['bareMetalHost'] = ApiClient.convertToType(data['bareMetalHost'], 'Boolean');
            }
            if (data.hasOwnProperty('guestVm')) {
                obj['guestVm'] = ApiClient.convertToType(data['guestVm'], 'Boolean');
            }
            if (data.hasOwnProperty('hasAutomation')) {
                obj['hasAutomation'] = ApiClient.convertToType(data['hasAutomation'], 'Boolean');
            }
            if (data.hasOwnProperty('provisionType')) {
                obj['provisionType'] = ZoneTypeProvisionType.constructFromObject(data['provisionType']);
            }
            if (data.hasOwnProperty('optionTypes')) {
                obj['optionTypes'] = ApiClient.convertToType(data['optionTypes'], [ZoneTypeOptionTypes]);
            }
            if (data.hasOwnProperty('displayOrder')) {
                obj['displayOrder'] = ApiClient.convertToType(data['displayOrder'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ZoneTypeServerTypes.prototype['id'] = undefined;

/**
 * @member {String} code
 */
ZoneTypeServerTypes.prototype['code'] = undefined;

/**
 * @member {String} name
 */
ZoneTypeServerTypes.prototype['name'] = undefined;

/**
 * @member {String} description
 */
ZoneTypeServerTypes.prototype['description'] = undefined;

/**
 * @member {String} nodeType
 */
ZoneTypeServerTypes.prototype['nodeType'] = undefined;

/**
 * @member {String} platform
 */
ZoneTypeServerTypes.prototype['platform'] = undefined;

/**
 * @member {Boolean} enabled
 */
ZoneTypeServerTypes.prototype['enabled'] = undefined;

/**
 * @member {Boolean} selectable
 */
ZoneTypeServerTypes.prototype['selectable'] = undefined;

/**
 * @member {Boolean} externalDelete
 */
ZoneTypeServerTypes.prototype['externalDelete'] = undefined;

/**
 * @member {Boolean} managed
 */
ZoneTypeServerTypes.prototype['managed'] = undefined;

/**
 * @member {Boolean} controlPower
 */
ZoneTypeServerTypes.prototype['controlPower'] = undefined;

/**
 * @member {Boolean} controlSuspend
 */
ZoneTypeServerTypes.prototype['controlSuspend'] = undefined;

/**
 * @member {Boolean} creatable
 */
ZoneTypeServerTypes.prototype['creatable'] = undefined;

/**
 * @member {Boolean} hasAgent
 */
ZoneTypeServerTypes.prototype['hasAgent'] = undefined;

/**
 * @member {Boolean} vmHypervisor
 */
ZoneTypeServerTypes.prototype['vmHypervisor'] = undefined;

/**
 * @member {Boolean} containerHypervisor
 */
ZoneTypeServerTypes.prototype['containerHypervisor'] = undefined;

/**
 * @member {Boolean} bareMetalHost
 */
ZoneTypeServerTypes.prototype['bareMetalHost'] = undefined;

/**
 * @member {Boolean} guestVm
 */
ZoneTypeServerTypes.prototype['guestVm'] = undefined;

/**
 * @member {Boolean} hasAutomation
 */
ZoneTypeServerTypes.prototype['hasAutomation'] = undefined;

/**
 * @member {module:model/ZoneTypeProvisionType} provisionType
 */
ZoneTypeServerTypes.prototype['provisionType'] = undefined;

/**
 * @member {Array.<module:model/ZoneTypeOptionTypes>} optionTypes
 */
ZoneTypeServerTypes.prototype['optionTypes'] = undefined;

/**
 * @member {Number} displayOrder
 */
ZoneTypeServerTypes.prototype['displayOrder'] = undefined;






export default ZoneTypeServerTypes;

