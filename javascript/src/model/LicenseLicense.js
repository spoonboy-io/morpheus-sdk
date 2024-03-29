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
import LicenseLicenseFeatures from './LicenseLicenseFeatures';

/**
 * The LicenseLicense model module.
 * @module model/LicenseLicense
 * @version 6.2.1
 */
class LicenseLicense {
    /**
     * Constructs a new <code>LicenseLicense</code>.
     * @alias module:model/LicenseLicense
     */
    constructor() { 
        
        LicenseLicense.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>LicenseLicense</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/LicenseLicense} obj Optional instance to populate.
     * @return {module:model/LicenseLicense} The populated <code>LicenseLicense</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new LicenseLicense();

            if (data.hasOwnProperty('productTier')) {
                obj['productTier'] = ApiClient.convertToType(data['productTier'], 'String');
            }
            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('endDate')) {
                obj['endDate'] = ApiClient.convertToType(data['endDate'], 'Date');
            }
            if (data.hasOwnProperty('maxInstances')) {
                obj['maxInstances'] = ApiClient.convertToType(data['maxInstances'], 'Number');
            }
            if (data.hasOwnProperty('maxMemory')) {
                obj['maxMemory'] = ApiClient.convertToType(data['maxMemory'], 'Number');
            }
            if (data.hasOwnProperty('maxStorage')) {
                obj['maxStorage'] = ApiClient.convertToType(data['maxStorage'], 'Number');
            }
            if (data.hasOwnProperty('hardLimit')) {
                obj['hardLimit'] = ApiClient.convertToType(data['hardLimit'], 'Boolean');
            }
            if (data.hasOwnProperty('freeTrial')) {
                obj['freeTrial'] = ApiClient.convertToType(data['freeTrial'], 'Boolean');
            }
            if (data.hasOwnProperty('multiTenant')) {
                obj['multiTenant'] = ApiClient.convertToType(data['multiTenant'], 'Boolean');
            }
            if (data.hasOwnProperty('whitelabel')) {
                obj['whitelabel'] = ApiClient.convertToType(data['whitelabel'], 'Boolean');
            }
            if (data.hasOwnProperty('reportStatus')) {
                obj['reportStatus'] = ApiClient.convertToType(data['reportStatus'], 'Boolean');
            }
            if (data.hasOwnProperty('supportLevel')) {
                obj['supportLevel'] = ApiClient.convertToType(data['supportLevel'], 'String');
            }
            if (data.hasOwnProperty('accountName')) {
                obj['accountName'] = ApiClient.convertToType(data['accountName'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
            if (data.hasOwnProperty('amazonProductCodes')) {
                obj['amazonProductCodes'] = ApiClient.convertToType(data['amazonProductCodes'], 'String');
            }
            if (data.hasOwnProperty('features')) {
                obj['features'] = LicenseLicenseFeatures.constructFromObject(data['features']);
            }
            if (data.hasOwnProperty('zoneTypes')) {
                obj['zoneTypes'] = ApiClient.convertToType(data['zoneTypes'], 'String');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {String} productTier
 */
LicenseLicense.prototype['productTier'] = undefined;

/**
 * @member {Date} startDate
 */
LicenseLicense.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
LicenseLicense.prototype['endDate'] = undefined;

/**
 * @member {Number} maxInstances
 */
LicenseLicense.prototype['maxInstances'] = undefined;

/**
 * @member {Number} maxMemory
 */
LicenseLicense.prototype['maxMemory'] = undefined;

/**
 * @member {Number} maxStorage
 */
LicenseLicense.prototype['maxStorage'] = undefined;

/**
 * @member {Boolean} hardLimit
 */
LicenseLicense.prototype['hardLimit'] = undefined;

/**
 * @member {Boolean} freeTrial
 */
LicenseLicense.prototype['freeTrial'] = undefined;

/**
 * @member {Boolean} multiTenant
 */
LicenseLicense.prototype['multiTenant'] = undefined;

/**
 * @member {Boolean} whitelabel
 */
LicenseLicense.prototype['whitelabel'] = undefined;

/**
 * @member {Boolean} reportStatus
 */
LicenseLicense.prototype['reportStatus'] = undefined;

/**
 * @member {String} supportLevel
 */
LicenseLicense.prototype['supportLevel'] = undefined;

/**
 * @member {String} accountName
 */
LicenseLicense.prototype['accountName'] = undefined;

/**
 * @member {Object} config
 */
LicenseLicense.prototype['config'] = undefined;

/**
 * @member {String} amazonProductCodes
 */
LicenseLicense.prototype['amazonProductCodes'] = undefined;

/**
 * @member {module:model/LicenseLicenseFeatures} features
 */
LicenseLicense.prototype['features'] = undefined;

/**
 * @member {String} zoneTypes
 */
LicenseLicense.prototype['zoneTypes'] = undefined;

/**
 * @member {Date} lastUpdated
 */
LicenseLicense.prototype['lastUpdated'] = undefined;

/**
 * @member {Date} dateCreated
 */
LicenseLicense.prototype['dateCreated'] = undefined;






export default LicenseLicense;

