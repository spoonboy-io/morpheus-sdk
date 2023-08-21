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
import VirtualImageConfig from './VirtualImageConfig';
import VirtualImageOsType from './VirtualImageOsType';

/**
 * The VirtualImage model module.
 * @module model/VirtualImage
 * @version 6.2.1
 */
class VirtualImage {
    /**
     * Constructs a new <code>VirtualImage</code>.
     * @alias module:model/VirtualImage
     */
    constructor() { 
        
        VirtualImage.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>VirtualImage</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/VirtualImage} obj Optional instance to populate.
     * @return {module:model/VirtualImage} The populated <code>VirtualImage</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new VirtualImage();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('ownerId')) {
                obj['ownerId'] = ApiClient.convertToType(data['ownerId'], 'Number');
            }
            if (data.hasOwnProperty('tenant')) {
                obj['tenant'] = InlineResponse20040AppDeployInstance.constructFromObject(data['tenant']);
            }
            if (data.hasOwnProperty('imageType')) {
                obj['imageType'] = ApiClient.convertToType(data['imageType'], 'String');
            }
            if (data.hasOwnProperty('userUploaded')) {
                obj['userUploaded'] = ApiClient.convertToType(data['userUploaded'], 'Boolean');
            }
            if (data.hasOwnProperty('userDefined')) {
                obj['userDefined'] = ApiClient.convertToType(data['userDefined'], 'Boolean');
            }
            if (data.hasOwnProperty('systemImage')) {
                obj['systemImage'] = ApiClient.convertToType(data['systemImage'], 'Boolean');
            }
            if (data.hasOwnProperty('isCloudInit')) {
                obj['isCloudInit'] = ApiClient.convertToType(data['isCloudInit'], 'Boolean');
            }
            if (data.hasOwnProperty('sshUsername')) {
                obj['sshUsername'] = ApiClient.convertToType(data['sshUsername'], 'String');
            }
            if (data.hasOwnProperty('sshPassword')) {
                obj['sshPassword'] = ApiClient.convertToType(data['sshPassword'], 'String');
            }
            if (data.hasOwnProperty('sshPasswordHash')) {
                obj['sshPasswordHash'] = ApiClient.convertToType(data['sshPasswordHash'], 'String');
            }
            if (data.hasOwnProperty('sshKey')) {
                obj['sshKey'] = ApiClient.convertToType(data['sshKey'], 'String');
            }
            if (data.hasOwnProperty('osType')) {
                obj['osType'] = VirtualImageOsType.constructFromObject(data['osType']);
            }
            if (data.hasOwnProperty('minRam')) {
                obj['minRam'] = ApiClient.convertToType(data['minRam'], 'Number');
            }
            if (data.hasOwnProperty('minRamGB')) {
                obj['minRamGB'] = ApiClient.convertToType(data['minRamGB'], 'Number');
            }
            if (data.hasOwnProperty('minDisk')) {
                obj['minDisk'] = ApiClient.convertToType(data['minDisk'], 'String');
            }
            if (data.hasOwnProperty('minDiskGB')) {
                obj['minDiskGB'] = ApiClient.convertToType(data['minDiskGB'], 'String');
            }
            if (data.hasOwnProperty('rawSize')) {
                obj['rawSize'] = ApiClient.convertToType(data['rawSize'], 'Number');
            }
            if (data.hasOwnProperty('rawSizeGB')) {
                obj['rawSizeGB'] = ApiClient.convertToType(data['rawSizeGB'], 'Number');
            }
            if (data.hasOwnProperty('trialVersion')) {
                obj['trialVersion'] = ApiClient.convertToType(data['trialVersion'], 'Boolean');
            }
            if (data.hasOwnProperty('virtioSupported')) {
                obj['virtioSupported'] = ApiClient.convertToType(data['virtioSupported'], 'Boolean');
            }
            if (data.hasOwnProperty('uefi')) {
                obj['uefi'] = ApiClient.convertToType(data['uefi'], 'String');
            }
            if (data.hasOwnProperty('isAutoJoinDomain')) {
                obj['isAutoJoinDomain'] = ApiClient.convertToType(data['isAutoJoinDomain'], 'Boolean');
            }
            if (data.hasOwnProperty('vmToolsInstalled')) {
                obj['vmToolsInstalled'] = ApiClient.convertToType(data['vmToolsInstalled'], 'Boolean');
            }
            if (data.hasOwnProperty('installAgent')) {
                obj['installAgent'] = ApiClient.convertToType(data['installAgent'], 'Boolean');
            }
            if (data.hasOwnProperty('isForceCustomization')) {
                obj['isForceCustomization'] = ApiClient.convertToType(data['isForceCustomization'], 'Boolean');
            }
            if (data.hasOwnProperty('isSysprep')) {
                obj['isSysprep'] = ApiClient.convertToType(data['isSysprep'], 'Boolean');
            }
            if (data.hasOwnProperty('fipsEnabled')) {
                obj['fipsEnabled'] = ApiClient.convertToType(data['fipsEnabled'], 'Boolean');
            }
            if (data.hasOwnProperty('userData')) {
                obj['userData'] = ApiClient.convertToType(data['userData'], 'String');
            }
            if (data.hasOwnProperty('consoleKeymap')) {
                obj['consoleKeymap'] = ApiClient.convertToType(data['consoleKeymap'], 'String');
            }
            if (data.hasOwnProperty('storageProvider')) {
                obj['storageProvider'] = ApiClient.convertToType(data['storageProvider'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('accounts')) {
                obj['accounts'] = ApiClient.convertToType(data['accounts'], [InlineResponse20040AppDeployInstance]);
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = VirtualImageConfig.constructFromObject(data['config']);
            }
            if (data.hasOwnProperty('volumes')) {
                obj['volumes'] = ApiClient.convertToType(data['volumes'], [Object]);
            }
            if (data.hasOwnProperty('storageControllers')) {
                obj['storageControllers'] = ApiClient.convertToType(data['storageControllers'], [Object]);
            }
            if (data.hasOwnProperty('networkInterfaces')) {
                obj['networkInterfaces'] = ApiClient.convertToType(data['networkInterfaces'], [Object]);
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], [Object]);
            }
            if (data.hasOwnProperty('locations')) {
                obj['locations'] = ApiClient.convertToType(data['locations'], [Object]);
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
VirtualImage.prototype['id'] = undefined;

/**
 * @member {String} name
 */
VirtualImage.prototype['name'] = undefined;

/**
 * @member {String} description
 */
VirtualImage.prototype['description'] = undefined;

/**
 * @member {Array.<String>} labels
 */
VirtualImage.prototype['labels'] = undefined;

/**
 * @member {Number} ownerId
 */
VirtualImage.prototype['ownerId'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} tenant
 */
VirtualImage.prototype['tenant'] = undefined;

/**
 * @member {String} imageType
 */
VirtualImage.prototype['imageType'] = undefined;

/**
 * @member {Boolean} userUploaded
 */
VirtualImage.prototype['userUploaded'] = undefined;

/**
 * @member {Boolean} userDefined
 */
VirtualImage.prototype['userDefined'] = undefined;

/**
 * @member {Boolean} systemImage
 */
VirtualImage.prototype['systemImage'] = undefined;

/**
 * @member {Boolean} isCloudInit
 */
VirtualImage.prototype['isCloudInit'] = undefined;

/**
 * @member {String} sshUsername
 */
VirtualImage.prototype['sshUsername'] = undefined;

/**
 * @member {String} sshPassword
 */
VirtualImage.prototype['sshPassword'] = undefined;

/**
 * @member {String} sshPasswordHash
 */
VirtualImage.prototype['sshPasswordHash'] = undefined;

/**
 * @member {String} sshKey
 */
VirtualImage.prototype['sshKey'] = undefined;

/**
 * @member {module:model/VirtualImageOsType} osType
 */
VirtualImage.prototype['osType'] = undefined;

/**
 * @member {Number} minRam
 */
VirtualImage.prototype['minRam'] = undefined;

/**
 * @member {Number} minRamGB
 */
VirtualImage.prototype['minRamGB'] = undefined;

/**
 * @member {String} minDisk
 */
VirtualImage.prototype['minDisk'] = undefined;

/**
 * @member {String} minDiskGB
 */
VirtualImage.prototype['minDiskGB'] = undefined;

/**
 * @member {Number} rawSize
 */
VirtualImage.prototype['rawSize'] = undefined;

/**
 * @member {Number} rawSizeGB
 */
VirtualImage.prototype['rawSizeGB'] = undefined;

/**
 * @member {Boolean} trialVersion
 */
VirtualImage.prototype['trialVersion'] = undefined;

/**
 * @member {Boolean} virtioSupported
 */
VirtualImage.prototype['virtioSupported'] = undefined;

/**
 * @member {String} uefi
 */
VirtualImage.prototype['uefi'] = undefined;

/**
 * @member {Boolean} isAutoJoinDomain
 */
VirtualImage.prototype['isAutoJoinDomain'] = undefined;

/**
 * @member {Boolean} vmToolsInstalled
 */
VirtualImage.prototype['vmToolsInstalled'] = undefined;

/**
 * @member {Boolean} installAgent
 */
VirtualImage.prototype['installAgent'] = undefined;

/**
 * @member {Boolean} isForceCustomization
 */
VirtualImage.prototype['isForceCustomization'] = undefined;

/**
 * @member {Boolean} isSysprep
 */
VirtualImage.prototype['isSysprep'] = undefined;

/**
 * @member {Boolean} fipsEnabled
 */
VirtualImage.prototype['fipsEnabled'] = undefined;

/**
 * @member {String} userData
 */
VirtualImage.prototype['userData'] = undefined;

/**
 * @member {String} consoleKeymap
 */
VirtualImage.prototype['consoleKeymap'] = undefined;

/**
 * @member {String} storageProvider
 */
VirtualImage.prototype['storageProvider'] = undefined;

/**
 * @member {String} externalId
 */
VirtualImage.prototype['externalId'] = undefined;

/**
 * @member {String} visibility
 */
VirtualImage.prototype['visibility'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse20040AppDeployInstance>} accounts
 */
VirtualImage.prototype['accounts'] = undefined;

/**
 * @member {module:model/VirtualImageConfig} config
 */
VirtualImage.prototype['config'] = undefined;

/**
 * @member {Array.<Object>} volumes
 */
VirtualImage.prototype['volumes'] = undefined;

/**
 * @member {Array.<Object>} storageControllers
 */
VirtualImage.prototype['storageControllers'] = undefined;

/**
 * @member {Array.<Object>} networkInterfaces
 */
VirtualImage.prototype['networkInterfaces'] = undefined;

/**
 * @member {Array.<Object>} tags
 */
VirtualImage.prototype['tags'] = undefined;

/**
 * @member {Array.<Object>} locations
 */
VirtualImage.prototype['locations'] = undefined;

/**
 * @member {Date} dateCreated
 */
VirtualImage.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
VirtualImage.prototype['lastUpdated'] = undefined;

/**
 * @member {String} status
 */
VirtualImage.prototype['status'] = undefined;






export default VirtualImage;

