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
 * The BlueprintMorpheusCreate model module.
 * @module model/BlueprintMorpheusCreate
 * @version 6.2.1
 */
class BlueprintMorpheusCreate {
    /**
     * Constructs a new <code>BlueprintMorpheusCreate</code>.
     * @alias module:model/BlueprintMorpheusCreate
     * @param name {String} A name for the blueprint
     * @param type {module:model/BlueprintMorpheusCreate.TypeEnum} Blueprint Type
     * @param tiers {Object} Tier definitions - Create in UI to view a baseline for object
     */
    constructor(name, type, tiers) { 
        
        BlueprintMorpheusCreate.initialize(this, name, type, tiers);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name, type, tiers) { 
        obj['name'] = name;
        obj['type'] = type;
        obj['tiers'] = tiers;
    }

    /**
     * Constructs a <code>BlueprintMorpheusCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BlueprintMorpheusCreate} obj Optional instance to populate.
     * @return {module:model/BlueprintMorpheusCreate} The populated <code>BlueprintMorpheusCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BlueprintMorpheusCreate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('tiers')) {
                obj['tiers'] = ApiClient.convertToType(data['tiers'], Object);
            }
        }
        return obj;
    }


}

/**
 * A name for the blueprint
 * @member {String} name
 */
BlueprintMorpheusCreate.prototype['name'] = undefined;

/**
 * Blueprint Type
 * @member {module:model/BlueprintMorpheusCreate.TypeEnum} type
 */
BlueprintMorpheusCreate.prototype['type'] = undefined;

/**
 * Array of label strings, can be used for filtering.
 * @member {Array.<String>} labels
 */
BlueprintMorpheusCreate.prototype['labels'] = undefined;

/**
 * Tier definitions - Create in UI to view a baseline for object
 * @member {Object} tiers
 */
BlueprintMorpheusCreate.prototype['tiers'] = undefined;





/**
 * Allowed values for the <code>type</code> property.
 * @enum {String}
 * @readonly
 */
BlueprintMorpheusCreate['TypeEnum'] = {

    /**
     * value: "morpheus"
     * @const
     */
    "morpheus": "morpheus"
};



export default BlueprintMorpheusCreate;

