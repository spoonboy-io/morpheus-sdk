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
import Budgets from './Budgets';
import Model200Success from './Model200Success';

/**
 * The AddBudgets200Response model module.
 * @module model/AddBudgets200Response
 * @version 6.1.1
 */
class AddBudgets200Response {
    /**
     * Constructs a new <code>AddBudgets200Response</code>.
     * @alias module:model/AddBudgets200Response
     * @implements module:model/Model200Success
     */
    constructor() { 
        Model200Success.initialize(this);
        AddBudgets200Response.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>AddBudgets200Response</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AddBudgets200Response} obj Optional instance to populate.
     * @return {module:model/AddBudgets200Response} The populated <code>AddBudgets200Response</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AddBudgets200Response();
            Model200Success.constructFromObject(data, obj);

            if (data.hasOwnProperty('success')) {
                obj['success'] = ApiClient.convertToType(data['success'], 'Boolean');
            }
            if (data.hasOwnProperty('budget')) {
                obj['budget'] = Budgets.constructFromObject(data['budget']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>AddBudgets200Response</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>AddBudgets200Response</code>.
     */
    static validateJSON(data) {
        // validate the optional field `budget`
        if (data['budget']) { // data not null
          Budgets.validateJSON(data['budget']);
        }

        return true;
    }


}



/**
 * @member {Boolean} success
 */
AddBudgets200Response.prototype['success'] = undefined;

/**
 * @member {module:model/Budgets} budget
 */
AddBudgets200Response.prototype['budget'] = undefined;


// Implement Model200Success interface:
/**
 * @member {Boolean} success
 */
Model200Success.prototype['success'] = undefined;




export default AddBudgets200Response;

