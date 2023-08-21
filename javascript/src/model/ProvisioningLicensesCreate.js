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

/**
 * The ProvisioningLicensesCreate model module.
 * @module model/ProvisioningLicensesCreate
 * @version 6.2.1
 */
class ProvisioningLicensesCreate {
    /**
     * Constructs a new <code>ProvisioningLicensesCreate</code>.
     * @alias module:model/ProvisioningLicensesCreate
     * @param name {String} Name
     * @param licenseType {String} License Type - The license type code.
     * @param licenseKey {String} License Key - The license key, to be kept a secret.
     */
    constructor(name, licenseType, licenseKey) { 
        
        ProvisioningLicensesCreate.initialize(this, name, licenseType, licenseKey);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, licenseType, licenseKey) { 
        obj['name'] = name;
        obj['licenseType'] = licenseType;
        obj['licenseKey'] = licenseKey;
    }

    /**
     * Constructs a <code>ProvisioningLicensesCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ProvisioningLicensesCreate} obj Optional instance to populate.
     * @return {module:model/ProvisioningLicensesCreate} The populated <code>ProvisioningLicensesCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ProvisioningLicensesCreate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('licenseType')) {
                obj['licenseType'] = ApiClient.convertToType(data['licenseType'], 'String');
            }
            if (data.hasOwnProperty('licenseKey')) {
                obj['licenseKey'] = ApiClient.convertToType(data['licenseKey'], 'String');
            }
            if (data.hasOwnProperty('orgName')) {
                obj['orgName'] = ApiClient.convertToType(data['orgName'], 'String');
            }
            if (data.hasOwnProperty('fullName')) {
                obj['fullName'] = ApiClient.convertToType(data['fullName'], 'String');
            }
            if (data.hasOwnProperty('licenseVersion')) {
                obj['licenseVersion'] = ApiClient.convertToType(data['licenseVersion'], 'String');
            }
            if (data.hasOwnProperty('copies')) {
                obj['copies'] = ApiClient.convertToType(data['copies'], 'Number');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('virtualImages')) {
                obj['virtualImages'] = ApiClient.convertToType(data['virtualImages'], ['Number']);
            }
            if (data.hasOwnProperty('tenants')) {
                obj['tenants'] = ApiClient.convertToType(data['tenants'], ['Number']);
            }
        }
        return obj;
    }


}

/**
 * Name
 * @member {String} name
 */
ProvisioningLicensesCreate.prototype['name'] = undefined;

/**
 * License Type - The license type code.
 * @member {String} licenseType
 */
ProvisioningLicensesCreate.prototype['licenseType'] = undefined;

/**
 * License Key - The license key, to be kept a secret.
 * @member {String} licenseKey
 */
ProvisioningLicensesCreate.prototype['licenseKey'] = undefined;

/**
 * Org Name - The Organization Name (if applicable) related to the license key
 * @member {String} orgName
 */
ProvisioningLicensesCreate.prototype['orgName'] = undefined;

/**
 * Full Name - The Full Name (if applicable) related to the license key
 * @member {String} fullName
 */
ProvisioningLicensesCreate.prototype['fullName'] = undefined;

/**
 * License Version
 * @member {String} licenseVersion
 */
ProvisioningLicensesCreate.prototype['licenseVersion'] = undefined;

/**
 * Copies - The number of times the key can be used.
 * @member {Number} copies
 * @default 1
 */
ProvisioningLicensesCreate.prototype['copies'] = 1;

/**
 * Description
 * @member {String} description
 */
ProvisioningLicensesCreate.prototype['description'] = undefined;

/**
 * Virtual Images - Array of Virtual Image IDs to associate with license.
 * @member {Array.<Number>} virtualImages
 */
ProvisioningLicensesCreate.prototype['virtualImages'] = undefined;

/**
 * Tenants - Array of tenants that are allowed to use the key.
 * @member {Array.<Number>} tenants
 */
ProvisioningLicensesCreate.prototype['tenants'] = undefined;






export default ProvisioningLicensesCreate;

