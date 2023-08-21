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
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';
import InstanceTypeLayout from './InstanceTypeLayout';
import OptionType from './OptionType';

/**
 * The InstanceType model module.
 * @module model/InstanceType
 * @version 6.2.1
 */
class InstanceType {
    /**
     * Constructs a new <code>InstanceType</code>.
     * @alias module:model/InstanceType
     */
    constructor() { 
        
        InstanceType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InstanceType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceType} obj Optional instance to populate.
     * @return {module:model/InstanceType} The populated <code>InstanceType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceType();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = InlineResponse20040AppDeployInstance.constructFromObject(data['account']);
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
            if (data.hasOwnProperty('hasProvisioningStep')) {
                obj['hasProvisioningStep'] = ApiClient.convertToType(data['hasProvisioningStep'], 'Boolean');
            }
            if (data.hasOwnProperty('hasDeployment')) {
                obj['hasDeployment'] = ApiClient.convertToType(data['hasDeployment'], 'Boolean');
            }
            if (data.hasOwnProperty('hasConfig')) {
                obj['hasConfig'] = ApiClient.convertToType(data['hasConfig'], 'Boolean');
            }
            if (data.hasOwnProperty('hasSettings')) {
                obj['hasSettings'] = ApiClient.convertToType(data['hasSettings'], 'Boolean');
            }
            if (data.hasOwnProperty('hasAutoScale')) {
                obj['hasAutoScale'] = ApiClient.convertToType(data['hasAutoScale'], 'Boolean');
            }
            if (data.hasOwnProperty('proxyType')) {
                obj['proxyType'] = ApiClient.convertToType(data['proxyType'], 'String');
            }
            if (data.hasOwnProperty('proxyPort')) {
                obj['proxyPort'] = ApiClient.convertToType(data['proxyPort'], 'String');
            }
            if (data.hasOwnProperty('proxyProtocol')) {
                obj['proxyProtocol'] = ApiClient.convertToType(data['proxyProtocol'], 'String');
            }
            if (data.hasOwnProperty('environmentPrefix')) {
                obj['environmentPrefix'] = ApiClient.convertToType(data['environmentPrefix'], 'String');
            }
            if (data.hasOwnProperty('backupType')) {
                obj['backupType'] = ApiClient.convertToType(data['backupType'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
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
                obj['instanceTypeLayouts'] = ApiClient.convertToType(data['instanceTypeLayouts'], [InstanceTypeLayout]);
            }
            if (data.hasOwnProperty('optionTypes')) {
                obj['optionTypes'] = ApiClient.convertToType(data['optionTypes'], [OptionType]);
            }
            if (data.hasOwnProperty('environmentVariables')) {
                obj['environmentVariables'] = ApiClient.convertToType(data['environmentVariables'], [Object]);
            }
            if (data.hasOwnProperty('priceSets')) {
                obj['priceSets'] = ApiClient.convertToType(data['priceSets'], [Object]);
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
InstanceType.prototype['id'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} account
 */
InstanceType.prototype['account'] = undefined;

/**
 * @member {String} name
 */
InstanceType.prototype['name'] = undefined;

/**
 * @member {Array.<String>} labels
 */
InstanceType.prototype['labels'] = undefined;

/**
 * @member {String} code
 */
InstanceType.prototype['code'] = undefined;

/**
 * @member {String} description
 */
InstanceType.prototype['description'] = undefined;

/**
 * @member {String} provisionTypeCode
 */
InstanceType.prototype['provisionTypeCode'] = undefined;

/**
 * @member {String} category
 */
InstanceType.prototype['category'] = undefined;

/**
 * @member {Boolean} active
 */
InstanceType.prototype['active'] = undefined;

/**
 * @member {Boolean} hasProvisioningStep
 */
InstanceType.prototype['hasProvisioningStep'] = undefined;

/**
 * @member {Boolean} hasDeployment
 */
InstanceType.prototype['hasDeployment'] = undefined;

/**
 * @member {Boolean} hasConfig
 */
InstanceType.prototype['hasConfig'] = undefined;

/**
 * @member {Boolean} hasSettings
 */
InstanceType.prototype['hasSettings'] = undefined;

/**
 * @member {Boolean} hasAutoScale
 */
InstanceType.prototype['hasAutoScale'] = undefined;

/**
 * @member {String} proxyType
 */
InstanceType.prototype['proxyType'] = undefined;

/**
 * @member {String} proxyPort
 */
InstanceType.prototype['proxyPort'] = undefined;

/**
 * @member {String} proxyProtocol
 */
InstanceType.prototype['proxyProtocol'] = undefined;

/**
 * @member {String} environmentPrefix
 */
InstanceType.prototype['environmentPrefix'] = undefined;

/**
 * @member {String} backupType
 */
InstanceType.prototype['backupType'] = undefined;

/**
 * @member {Object} config
 */
InstanceType.prototype['config'] = undefined;

/**
 * @member {String} visibility
 */
InstanceType.prototype['visibility'] = undefined;

/**
 * @member {Boolean} featured
 */
InstanceType.prototype['featured'] = undefined;

/**
 * @member {Array.<String>} versions
 */
InstanceType.prototype['versions'] = undefined;

/**
 * @member {Array.<module:model/InstanceTypeLayout>} instanceTypeLayouts
 */
InstanceType.prototype['instanceTypeLayouts'] = undefined;

/**
 * @member {Array.<module:model/OptionType>} optionTypes
 */
InstanceType.prototype['optionTypes'] = undefined;

/**
 * @member {Array.<Object>} environmentVariables
 */
InstanceType.prototype['environmentVariables'] = undefined;

/**
 * @member {Array.<Object>} priceSets
 */
InstanceType.prototype['priceSets'] = undefined;

/**
 * Logo image URL
 * @member {String} imagePath
 */
InstanceType.prototype['imagePath'] = undefined;

/**
 * Dark logo image URL
 * @member {String} darkImagePath
 */
InstanceType.prototype['darkImagePath'] = undefined;






export default InstanceType;
