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
import GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions from './GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions';

/**
 * The GuidanceAzureReservationsConfigServicesAzureVms model module.
 * @module model/GuidanceAzureReservationsConfigServicesAzureVms
 * @version 6.2.1
 */
class GuidanceAzureReservationsConfigServicesAzureVms {
    /**
     * Constructs a new <code>GuidanceAzureReservationsConfigServicesAzureVms</code>.
     * @alias module:model/GuidanceAzureReservationsConfigServicesAzureVms
     */
    constructor() { 
        
        GuidanceAzureReservationsConfigServicesAzureVms.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>GuidanceAzureReservationsConfigServicesAzureVms</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/GuidanceAzureReservationsConfigServicesAzureVms} obj Optional instance to populate.
     * @return {module:model/GuidanceAzureReservationsConfigServicesAzureVms} The populated <code>GuidanceAzureReservationsConfigServicesAzureVms</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new GuidanceAzureReservationsConfigServicesAzureVms();

            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('paymentOptions')) {
                obj['paymentOptions'] = ApiClient.convertToType(data['paymentOptions'], {'String': GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions});
            }
        }
        return obj;
    }


}

/**
 * @member {String} code
 */
GuidanceAzureReservationsConfigServicesAzureVms.prototype['code'] = undefined;

/**
 * @member {String} name
 */
GuidanceAzureReservationsConfigServicesAzureVms.prototype['name'] = undefined;

/**
 * @member {Object.<String, module:model/GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions>} paymentOptions
 */
GuidanceAzureReservationsConfigServicesAzureVms.prototype['paymentOptions'] = undefined;






export default GuidanceAzureReservationsConfigServicesAzureVms;

