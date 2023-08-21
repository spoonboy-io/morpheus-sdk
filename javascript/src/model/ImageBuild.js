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
import ImageBuildConfig from './ImageBuildConfig';
import ImageBuildLastResult from './ImageBuildLastResult';
import ImageBuildsBootScript from './ImageBuildsBootScript';
import ImageBuildsScripts from './ImageBuildsScripts';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';
import InlineResponse20079LoadBalancerMonitorLoadBalancerType from './InlineResponse20079LoadBalancerMonitorLoadBalancerType';

/**
 * The ImageBuild model module.
 * @module model/ImageBuild
 * @version 6.2.1
 */
class ImageBuild {
    /**
     * Constructs a new <code>ImageBuild</code>.
     * @alias module:model/ImageBuild
     */
    constructor() { 
        
        ImageBuild.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ImageBuild</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ImageBuild} obj Optional instance to populate.
     * @return {module:model/ImageBuild} The populated <code>ImageBuild</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ImageBuild();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = InlineResponse20040AppDeployInstance.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = InlineResponse20079LoadBalancerMonitorLoadBalancerType.constructFromObject(data['type']);
            }
            if (data.hasOwnProperty('site')) {
                obj['site'] = InlineResponse20040AppDeployInstance.constructFromObject(data['site']);
            }
            if (data.hasOwnProperty('zone')) {
                obj['zone'] = InlineResponse20040AppDeployInstance.constructFromObject(data['zone']);
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('bootScript')) {
                obj['bootScript'] = ImageBuildsBootScript.constructFromObject(data['bootScript']);
            }
            if (data.hasOwnProperty('bootCommand')) {
                obj['bootCommand'] = ApiClient.convertToType(data['bootCommand'], 'String');
            }
            if (data.hasOwnProperty('preseedScript')) {
                obj['preseedScript'] = ImageBuildsBootScript.constructFromObject(data['preseedScript']);
            }
            if (data.hasOwnProperty('scripts')) {
                obj['scripts'] = ApiClient.convertToType(data['scripts'], [ImageBuildsScripts]);
            }
            if (data.hasOwnProperty('sshUsername')) {
                obj['sshUsername'] = ApiClient.convertToType(data['sshUsername'], 'String');
            }
            if (data.hasOwnProperty('sshPassword')) {
                obj['sshPassword'] = ApiClient.convertToType(data['sshPassword'], 'String');
            }
            if (data.hasOwnProperty('storageProvider')) {
                obj['storageProvider'] = ApiClient.convertToType(data['storageProvider'], 'String');
            }
            if (data.hasOwnProperty('buildOutputName')) {
                obj['buildOutputName'] = ApiClient.convertToType(data['buildOutputName'], 'String');
            }
            if (data.hasOwnProperty('conversionFormats')) {
                obj['conversionFormats'] = ApiClient.convertToType(data['conversionFormats'], 'String');
            }
            if (data.hasOwnProperty('isCloudInit')) {
                obj['isCloudInit'] = ApiClient.convertToType(data['isCloudInit'], 'Boolean');
            }
            if (data.hasOwnProperty('vmToolsInstalled')) {
                obj['vmToolsInstalled'] = ApiClient.convertToType(data['vmToolsInstalled'], 'Boolean');
            }
            if (data.hasOwnProperty('keepResults')) {
                obj['keepResults'] = ApiClient.convertToType(data['keepResults'], 'Number');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ImageBuildConfig.constructFromObject(data['config']);
            }
            if (data.hasOwnProperty('lastResult')) {
                obj['lastResult'] = ImageBuildLastResult.constructFromObject(data['lastResult']);
            }
            if (data.hasOwnProperty('executionCount')) {
                obj['executionCount'] = ApiClient.convertToType(data['executionCount'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
ImageBuild.prototype['id'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} account
 */
ImageBuild.prototype['account'] = undefined;

/**
 * @member {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancerType} type
 */
ImageBuild.prototype['type'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} site
 */
ImageBuild.prototype['site'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} zone
 */
ImageBuild.prototype['zone'] = undefined;

/**
 * @member {String} name
 */
ImageBuild.prototype['name'] = undefined;

/**
 * @member {String} description
 */
ImageBuild.prototype['description'] = undefined;

/**
 * @member {module:model/ImageBuildsBootScript} bootScript
 */
ImageBuild.prototype['bootScript'] = undefined;

/**
 * @member {String} bootCommand
 */
ImageBuild.prototype['bootCommand'] = undefined;

/**
 * @member {module:model/ImageBuildsBootScript} preseedScript
 */
ImageBuild.prototype['preseedScript'] = undefined;

/**
 * @member {Array.<module:model/ImageBuildsScripts>} scripts
 */
ImageBuild.prototype['scripts'] = undefined;

/**
 * @member {String} sshUsername
 */
ImageBuild.prototype['sshUsername'] = undefined;

/**
 * @member {String} sshPassword
 */
ImageBuild.prototype['sshPassword'] = undefined;

/**
 * @member {String} storageProvider
 */
ImageBuild.prototype['storageProvider'] = undefined;

/**
 * @member {String} buildOutputName
 */
ImageBuild.prototype['buildOutputName'] = undefined;

/**
 * @member {String} conversionFormats
 */
ImageBuild.prototype['conversionFormats'] = undefined;

/**
 * @member {Boolean} isCloudInit
 */
ImageBuild.prototype['isCloudInit'] = undefined;

/**
 * @member {Boolean} vmToolsInstalled
 */
ImageBuild.prototype['vmToolsInstalled'] = undefined;

/**
 * @member {Number} keepResults
 */
ImageBuild.prototype['keepResults'] = undefined;

/**
 * @member {module:model/ImageBuildConfig} config
 */
ImageBuild.prototype['config'] = undefined;

/**
 * @member {module:model/ImageBuildLastResult} lastResult
 */
ImageBuild.prototype['lastResult'] = undefined;

/**
 * @member {Number} executionCount
 */
ImageBuild.prototype['executionCount'] = undefined;






export default ImageBuild;
