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
import App from './App';

/**
 * The GetApp200Response model module.
 * @module model/GetApp200Response
 * @version 6.1.1
 */
class GetApp200Response {
    /**
     * Constructs a new <code>GetApp200Response</code>.
     * @alias module:model/GetApp200Response
     */
    constructor() { 
        
        GetApp200Response.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>GetApp200Response</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/GetApp200Response} obj Optional instance to populate.
     * @return {module:model/GetApp200Response} The populated <code>GetApp200Response</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new GetApp200Response();

            if (data.hasOwnProperty('app')) {
                obj['app'] = App.constructFromObject(data['app']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>GetApp200Response</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>GetApp200Response</code>.
     */
    static validateJSON(data) {
        // validate the optional field `app`
        if (data['app']) { // data not null
          App.validateJSON(data['app']);
        }

        return true;
    }


}



/**
 * @member {module:model/App} app
 */
GetApp200Response.prototype['app'] = undefined;






export default GetApp200Response;

