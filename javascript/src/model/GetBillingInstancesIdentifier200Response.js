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
import BillingInstance from './BillingInstance';
import Model200Success from './Model200Success';

/**
 * The GetBillingInstancesIdentifier200Response model module.
 * @module model/GetBillingInstancesIdentifier200Response
 * @version 6.1.1
 */
class GetBillingInstancesIdentifier200Response {
    /**
     * Constructs a new <code>GetBillingInstancesIdentifier200Response</code>.
     * @alias module:model/GetBillingInstancesIdentifier200Response
     * @implements module:model/Model200Success
     */
    constructor() { 
        Model200Success.initialize(this);
        GetBillingInstancesIdentifier200Response.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>GetBillingInstancesIdentifier200Response</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/GetBillingInstancesIdentifier200Response} obj Optional instance to populate.
     * @return {module:model/GetBillingInstancesIdentifier200Response} The populated <code>GetBillingInstancesIdentifier200Response</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new GetBillingInstancesIdentifier200Response();
            Model200Success.constructFromObject(data, obj);

            if (data.hasOwnProperty('success')) {
                obj['success'] = ApiClient.convertToType(data['success'], 'Boolean');
            }
            if (data.hasOwnProperty('billingInfo')) {
                obj['billingInfo'] = BillingInstance.constructFromObject(data['billingInfo']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>GetBillingInstancesIdentifier200Response</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>GetBillingInstancesIdentifier200Response</code>.
     */
    static validateJSON(data) {
        // validate the optional field `billingInfo`
        if (data['billingInfo']) { // data not null
          BillingInstance.validateJSON(data['billingInfo']);
        }

        return true;
    }


}



/**
 * @member {Boolean} success
 */
GetBillingInstancesIdentifier200Response.prototype['success'] = undefined;

/**
 * @member {module:model/BillingInstance} billingInfo
 */
GetBillingInstancesIdentifier200Response.prototype['billingInfo'] = undefined;


// Implement Model200Success interface:
/**
 * @member {Boolean} success
 */
Model200Success.prototype['success'] = undefined;




export default GetBillingInstancesIdentifier200Response;

