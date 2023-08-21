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

/**
 * The IntegrationInventory model module.
 * @module model/IntegrationInventory
 * @version 6.2.1
 */
class IntegrationInventory {
    /**
     * Constructs a new <code>IntegrationInventory</code>.
     * @alias module:model/IntegrationInventory
     */
    constructor() { 
        
        IntegrationInventory.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>IntegrationInventory</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/IntegrationInventory} obj Optional instance to populate.
     * @return {module:model/IntegrationInventory} The populated <code>IntegrationInventory</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new IntegrationInventory();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = InlineResponse20040AppDeployInstance.constructFromObject(data['owner']);
            }
            if (data.hasOwnProperty('tenants')) {
                obj['tenants'] = ApiClient.convertToType(data['tenants'], [InlineResponse20040AppDeployInstance]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
IntegrationInventory.prototype['id'] = undefined;

/**
 * @member {String} name
 */
IntegrationInventory.prototype['name'] = undefined;

/**
 * @member {String} description
 */
IntegrationInventory.prototype['description'] = undefined;

/**
 * @member {String} externalId
 */
IntegrationInventory.prototype['externalId'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} owner
 */
IntegrationInventory.prototype['owner'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse20040AppDeployInstance>} tenants
 */
IntegrationInventory.prototype['tenants'] = undefined;






export default IntegrationInventory;

