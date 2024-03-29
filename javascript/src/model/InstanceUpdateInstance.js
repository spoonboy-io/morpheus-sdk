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
import ApiServersIdMakeManagedServerTags from './ApiServersIdMakeManagedServerTags';
import InstanceUpdateInstanceRemoveTags from './InstanceUpdateInstanceRemoveTags';
import InstanceUpdateInstanceSite from './InstanceUpdateInstanceSite';

/**
 * The InstanceUpdateInstance model module.
 * @module model/InstanceUpdateInstance
 * @version 6.2.1
 */
class InstanceUpdateInstance {
    /**
     * Constructs a new <code>InstanceUpdateInstance</code>.
     * @alias module:model/InstanceUpdateInstance
     */
    constructor() { 
        
        InstanceUpdateInstance.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InstanceUpdateInstance</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstanceUpdateInstance} obj Optional instance to populate.
     * @return {module:model/InstanceUpdateInstance} The populated <code>InstanceUpdateInstance</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstanceUpdateInstance();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('instanceContext')) {
                obj['instanceContext'] = ApiClient.convertToType(data['instanceContext'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], [ApiServersIdMakeManagedServerTags]);
            }
            if (data.hasOwnProperty('addTags')) {
                obj['addTags'] = ApiClient.convertToType(data['addTags'], [ApiServersIdMakeManagedServerTags]);
            }
            if (data.hasOwnProperty('removeTags')) {
                obj['removeTags'] = ApiClient.convertToType(data['removeTags'], [InstanceUpdateInstanceRemoveTags]);
            }
            if (data.hasOwnProperty('powerScheduleType')) {
                obj['powerScheduleType'] = ApiClient.convertToType(data['powerScheduleType'], 'Number');
            }
            if (data.hasOwnProperty('site')) {
                obj['site'] = InstanceUpdateInstanceSite.constructFromObject(data['site']);
            }
            if (data.hasOwnProperty('ownerId')) {
                obj['ownerId'] = ApiClient.convertToType(data['ownerId'], 'Number');
            }
            if (data.hasOwnProperty('displayName')) {
                obj['displayName'] = ApiClient.convertToType(data['displayName'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Unique name scoped to your account for the instance.
 * @member {String} name
 */
InstanceUpdateInstance.prototype['name'] = undefined;

/**
 * Optional description field.
 * @member {String} description
 */
InstanceUpdateInstance.prototype['description'] = undefined;

/**
 * Environment
 * @member {String} instanceContext
 */
InstanceUpdateInstance.prototype['instanceContext'] = undefined;

/**
 * Array of strings (keywords).
 * @member {Array.<String>} labels
 */
InstanceUpdateInstance.prototype['labels'] = undefined;

/**
 * Metadata tags, Array of objects having a name and value.
 * @member {Array.<module:model/ApiServersIdMakeManagedServerTags>} tags
 */
InstanceUpdateInstance.prototype['tags'] = undefined;

/**
 * Add or update value of Metadata tags, Array of objects having a name and value.
 * @member {Array.<module:model/ApiServersIdMakeManagedServerTags>} addTags
 */
InstanceUpdateInstance.prototype['addTags'] = undefined;

/**
 * Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed.
 * @member {Array.<module:model/InstanceUpdateInstanceRemoveTags>} removeTags
 */
InstanceUpdateInstance.prototype['removeTags'] = undefined;

/**
 * Power schedule ID.
 * @member {Number} powerScheduleType
 */
InstanceUpdateInstance.prototype['powerScheduleType'] = undefined;

/**
 * @member {module:model/InstanceUpdateInstanceSite} site
 */
InstanceUpdateInstance.prototype['site'] = undefined;

/**
 * User ID, can be used to change instance owner.
 * @member {Number} ownerId
 */
InstanceUpdateInstance.prototype['ownerId'] = undefined;

/**
 * Name used in the UI for display
 * @member {String} displayName
 */
InstanceUpdateInstance.prototype['displayName'] = undefined;






export default InstanceUpdateInstance;

