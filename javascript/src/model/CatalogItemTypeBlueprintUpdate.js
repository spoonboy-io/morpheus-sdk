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
import CatalogItemTypeBlueprintCreateBlueprint from './CatalogItemTypeBlueprintCreateBlueprint';

/**
 * The CatalogItemTypeBlueprintUpdate model module.
 * @module model/CatalogItemTypeBlueprintUpdate
 * @version 6.2.1
 */
class CatalogItemTypeBlueprintUpdate {
    /**
     * Constructs a new <code>CatalogItemTypeBlueprintUpdate</code>.
     * @alias module:model/CatalogItemTypeBlueprintUpdate
     */
    constructor() { 
        
        CatalogItemTypeBlueprintUpdate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>CatalogItemTypeBlueprintUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/CatalogItemTypeBlueprintUpdate} obj Optional instance to populate.
     * @return {module:model/CatalogItemTypeBlueprintUpdate} The populated <code>CatalogItemTypeBlueprintUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new CatalogItemTypeBlueprintUpdate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('code')) {
                obj['code'] = ApiClient.convertToType(data['code'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('layoutCode')) {
                obj['layoutCode'] = ApiClient.convertToType(data['layoutCode'], 'String');
            }
            if (data.hasOwnProperty('iconPath')) {
                obj['iconPath'] = ApiClient.convertToType(data['iconPath'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('featured')) {
                obj['featured'] = ApiClient.convertToType(data['featured'], 'Boolean');
            }
            if (data.hasOwnProperty('allowQuantity')) {
                obj['allowQuantity'] = ApiClient.convertToType(data['allowQuantity'], 'Boolean');
            }
            if (data.hasOwnProperty('blueprint')) {
                obj['blueprint'] = CatalogItemTypeBlueprintCreateBlueprint.constructFromObject(data['blueprint']);
            }
            if (data.hasOwnProperty('appSpec')) {
                obj['appSpec'] = ApiClient.convertToType(data['appSpec'], 'String');
            }
            if (data.hasOwnProperty('optionTypes')) {
                obj['optionTypes'] = ApiClient.convertToType(data['optionTypes'], ['Number']);
            }
        }
        return obj;
    }


}

/**
 * Catalog Item Type name
 * @member {String} name
 */
CatalogItemTypeBlueprintUpdate.prototype['name'] = undefined;

/**
 * Useful shortcode for provisioning naming schemes and export reference.
 * @member {String} code
 */
CatalogItemTypeBlueprintUpdate.prototype['code'] = undefined;

/**
 * Catalog Item Type category
 * @member {String} category
 */
CatalogItemTypeBlueprintUpdate.prototype['category'] = undefined;

/**
 * Catalog Item Type description
 * @member {String} description
 */
CatalogItemTypeBlueprintUpdate.prototype['description'] = undefined;

/**
 * Array of label strings, can be used for filtering.
 * @member {Array.<String>} labels
 */
CatalogItemTypeBlueprintUpdate.prototype['labels'] = undefined;

/**
 * Type, `instance`, `blueprint` or `workflow`. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context.
 * @member {module:model/CatalogItemTypeBlueprintUpdate.TypeEnum} type
 */
CatalogItemTypeBlueprintUpdate.prototype['type'] = undefined;

/**
 * Visibility - Set to public to allow all tenants
 * @member {String} visibility
 * @default 'private'
 */
CatalogItemTypeBlueprintUpdate.prototype['visibility'] = 'private';

/**
 * Identifier primarily used for Plugin Catalog Item Types
 * @member {String} layoutCode
 */
CatalogItemTypeBlueprintUpdate.prototype['layoutCode'] = undefined;

/**
 * Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png.
 * @member {String} iconPath
 */
CatalogItemTypeBlueprintUpdate.prototype['iconPath'] = undefined;

/**
 * Can be used to enable / disable the catalog item type.
 * @member {Boolean} enabled
 * @default true
 */
CatalogItemTypeBlueprintUpdate.prototype['enabled'] = true;

/**
 * Can be used to feature the catalog item type.
 * @member {Boolean} featured
 * @default false
 */
CatalogItemTypeBlueprintUpdate.prototype['featured'] = false;

/**
 * Can users order more than one of this item at a time.
 * @member {Boolean} allowQuantity
 * @default false
 */
CatalogItemTypeBlueprintUpdate.prototype['allowQuantity'] = false;

/**
 * @member {module:model/CatalogItemTypeBlueprintCreateBlueprint} blueprint
 */
CatalogItemTypeBlueprintUpdate.prototype['blueprint'] = undefined;

/**
 * The appSpec for blueprint type catalog items is a string in the Scribe YAML format with fields
 * @member {String} appSpec
 */
CatalogItemTypeBlueprintUpdate.prototype['appSpec'] = undefined;

/**
 * Array of option type IDs, see Inputs. Only applies to type instance and blueprint.
 * @member {Array.<Number>} optionTypes
 */
CatalogItemTypeBlueprintUpdate.prototype['optionTypes'] = undefined;





/**
 * Allowed values for the <code>type</code> property.
 * @enum {String}
 * @readonly
 */
CatalogItemTypeBlueprintUpdate['TypeEnum'] = {

    /**
     * value: "blueprint"
     * @const
     */
    "blueprint": "blueprint"
};



export default CatalogItemTypeBlueprintUpdate;

