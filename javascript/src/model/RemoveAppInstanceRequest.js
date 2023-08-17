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

/**
 * The RemoveAppInstanceRequest model module.
 * @module model/RemoveAppInstanceRequest
 * @version 6.1.1
 */
class RemoveAppInstanceRequest {
    /**
     * Constructs a new <code>RemoveAppInstanceRequest</code>.
     * @alias module:model/RemoveAppInstanceRequest
     * @param instanceId {Number} The ID of the instance being removed
     */
    constructor(instanceId) { 
        
        RemoveAppInstanceRequest.initialize(this, instanceId);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, instanceId) { 
        obj['instanceId'] = instanceId;
    }

    /**
     * Constructs a <code>RemoveAppInstanceRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/RemoveAppInstanceRequest} obj Optional instance to populate.
     * @return {module:model/RemoveAppInstanceRequest} The populated <code>RemoveAppInstanceRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RemoveAppInstanceRequest();

            if (data.hasOwnProperty('instanceId')) {
                obj['instanceId'] = ApiClient.convertToType(data['instanceId'], 'Number');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>RemoveAppInstanceRequest</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>RemoveAppInstanceRequest</code>.
     */
    static validateJSON(data) {
        // check to make sure all required properties are present in the JSON string
        for (const property of RemoveAppInstanceRequest.RequiredProperties) {
            if (!data[property]) {
                throw new Error("The required field `" + property + "` is not found in the JSON data: " + JSON.stringify(data));
            }
        }

        return true;
    }


}

RemoveAppInstanceRequest.RequiredProperties = ["instanceId"];

/**
 * The ID of the instance being removed
 * @member {Number} instanceId
 */
RemoveAppInstanceRequest.prototype['instanceId'] = undefined;






export default RemoveAppInstanceRequest;

