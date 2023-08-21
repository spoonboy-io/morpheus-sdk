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
import OneOfobjectobject from './OneOfobjectobject';
import OneOfobjectstring from './OneOfobjectstring';
import VirtualImageCreateStorageProvider from './VirtualImageCreateStorageProvider';
import VirtualImageCreateTags from './VirtualImageCreateTags';
import VirtualImageUpdateRemoveTags from './VirtualImageUpdateRemoveTags';

/**
 * The VirtualImageUpdate model module.
 * @module model/VirtualImageUpdate
 * @version 6.2.1
 */
class VirtualImageUpdate {
    /**
     * Constructs a new <code>VirtualImageUpdate</code>.
     * @alias module:model/VirtualImageUpdate
     */
    constructor() { 
        
        VirtualImageUpdate.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>VirtualImageUpdate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/VirtualImageUpdate} obj Optional instance to populate.
     * @return {module:model/VirtualImageUpdate} The populated <code>VirtualImageUpdate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new VirtualImageUpdate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('imageType')) {
                obj['imageType'] = ApiClient.convertToType(data['imageType'], 'String');
            }
            if (data.hasOwnProperty('storageProvider')) {
                obj['storageProvider'] = VirtualImageCreateStorageProvider.constructFromObject(data['storageProvider']);
            }
            if (data.hasOwnProperty('isCloudInit')) {
                obj['isCloudInit'] = ApiClient.convertToType(data['isCloudInit'], 'Boolean');
            }
            if (data.hasOwnProperty('userData')) {
                obj['userData'] = ApiClient.convertToType(data['userData'], 'String');
            }
            if (data.hasOwnProperty('installAgent')) {
                obj['installAgent'] = ApiClient.convertToType(data['installAgent'], 'Boolean');
            }
            if (data.hasOwnProperty('sshUsername')) {
                obj['sshUsername'] = ApiClient.convertToType(data['sshUsername'], 'String');
            }
            if (data.hasOwnProperty('sshPassword')) {
                obj['sshPassword'] = ApiClient.convertToType(data['sshPassword'], 'String');
            }
            if (data.hasOwnProperty('sshKey')) {
                obj['sshKey'] = ApiClient.convertToType(data['sshKey'], 'String');
            }
            if (data.hasOwnProperty('osType')) {
                obj['osType'] = ApiClient.convertToType(data['osType'], OneOfobjectstring);
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('accounts')) {
                obj['accounts'] = ApiClient.convertToType(data['accounts'], ['Number']);
            }
            if (data.hasOwnProperty('isAutoJoinDomain')) {
                obj['isAutoJoinDomain'] = ApiClient.convertToType(data['isAutoJoinDomain'], 'Boolean');
            }
            if (data.hasOwnProperty('virtioSupported')) {
                obj['virtioSupported'] = ApiClient.convertToType(data['virtioSupported'], 'Boolean');
            }
            if (data.hasOwnProperty('vmToolsInstalled')) {
                obj['vmToolsInstalled'] = ApiClient.convertToType(data['vmToolsInstalled'], 'Boolean');
            }
            if (data.hasOwnProperty('isForceCustomization')) {
                obj['isForceCustomization'] = ApiClient.convertToType(data['isForceCustomization'], 'Boolean');
            }
            if (data.hasOwnProperty('trialVersion')) {
                obj['trialVersion'] = ApiClient.convertToType(data['trialVersion'], 'Boolean');
            }
            if (data.hasOwnProperty('isSysprep')) {
                obj['isSysprep'] = ApiClient.convertToType(data['isSysprep'], 'Boolean');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], OneOfobjectobject);
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], [VirtualImageCreateTags]);
            }
            if (data.hasOwnProperty('addTags')) {
                obj['addTags'] = ApiClient.convertToType(data['addTags'], [VirtualImageCreateTags]);
            }
            if (data.hasOwnProperty('removeTags')) {
                obj['removeTags'] = ApiClient.convertToType(data['removeTags'], [VirtualImageUpdateRemoveTags]);
            }
        }
        return obj;
    }


}

/**
 * A name for the virtual image
 * @member {String} name
 */
VirtualImageUpdate.prototype['name'] = undefined;

/**
 * Array of label strings, can be used for filtering.
 * @member {Array.<String>} labels
 */
VirtualImageUpdate.prototype['labels'] = undefined;

/**
 * Code of image type. eg. vmware, ami, etc.
 * @member {String} imageType
 */
VirtualImageUpdate.prototype['imageType'] = undefined;

/**
 * @member {module:model/VirtualImageCreateStorageProvider} storageProvider
 */
VirtualImageUpdate.prototype['storageProvider'] = undefined;

/**
 * Cloud Init Enabled?
 * @member {Boolean} isCloudInit
 * @default false
 */
VirtualImageUpdate.prototype['isCloudInit'] = false;

/**
 * Cloud-Init User Data, a bash script
 * @member {String} userData
 */
VirtualImageUpdate.prototype['userData'] = undefined;

/**
 * Install Agent?
 * @member {Boolean} installAgent
 * @default false
 */
VirtualImageUpdate.prototype['installAgent'] = false;

/**
 * SSH Username
 * @member {String} sshUsername
 */
VirtualImageUpdate.prototype['sshUsername'] = undefined;

/**
 * SSH Password
 * @member {String} sshPassword
 */
VirtualImageUpdate.prototype['sshPassword'] = undefined;

/**
 * SSH Key
 * @member {String} sshKey
 */
VirtualImageUpdate.prototype['sshKey'] = undefined;

/**
 * A Map containing the id of the OS Type. This can also be passed as a string (code or name) instead.
 * @member {module:model/OneOfobjectstring} osType
 */
VirtualImageUpdate.prototype['osType'] = undefined;

/**
 * private or public
 * @member {String} visibility
 * @default 'private'
 */
VirtualImageUpdate.prototype['visibility'] = 'private';

/**
 * @member {Array.<Number>} accounts
 */
VirtualImageUpdate.prototype['accounts'] = undefined;

/**
 * Auto Join Domain?
 * @member {Boolean} isAutoJoinDomain
 * @default false
 */
VirtualImageUpdate.prototype['isAutoJoinDomain'] = false;

/**
 * VirtIO Drivers Loaded?
 * @member {Boolean} virtioSupported
 * @default true
 */
VirtualImageUpdate.prototype['virtioSupported'] = true;

/**
 * VM Tools Installed?
 * @member {Boolean} vmToolsInstalled
 * @default true
 */
VirtualImageUpdate.prototype['vmToolsInstalled'] = true;

/**
 * Force Guest Customization?
 * @member {Boolean} isForceCustomization
 * @default false
 */
VirtualImageUpdate.prototype['isForceCustomization'] = false;

/**
 * Trial Version
 * @member {Boolean} trialVersion
 * @default false
 */
VirtualImageUpdate.prototype['trialVersion'] = false;

/**
 * Sysprep Enabled?
 * @member {Boolean} isSysprep
 * @default false
 */
VirtualImageUpdate.prototype['isSysprep'] = false;

/**
 * Map of configuration properties, varies by image type.
 * @member {module:model/OneOfobjectobject} config
 */
VirtualImageUpdate.prototype['config'] = undefined;

/**
 * Metadata tags, Array of objects having a name and value, this adds or updates the specified tags and removes any tags not specified.
 * @member {Array.<module:model/VirtualImageCreateTags>} tags
 */
VirtualImageUpdate.prototype['tags'] = undefined;

/**
 * Add or update value of Metadata tags, Array of objects having a name and value.
 * @member {Array.<module:model/VirtualImageCreateTags>} addTags
 */
VirtualImageUpdate.prototype['addTags'] = undefined;

/**
 * Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed.
 * @member {Array.<module:model/VirtualImageUpdateRemoveTags>} removeTags
 */
VirtualImageUpdate.prototype['removeTags'] = undefined;






export default VirtualImageUpdate;
