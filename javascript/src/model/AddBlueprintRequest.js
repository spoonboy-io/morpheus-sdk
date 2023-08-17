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
import BlueprintARMCreate from './BlueprintARMCreate';
import BlueprintARMCreateArm from './BlueprintARMCreateArm';
import BlueprintCFTCreate from './BlueprintCFTCreate';
import BlueprintCFTCreateCloudFormation from './BlueprintCFTCreateCloudFormation';
import BlueprintHelmCreate from './BlueprintHelmCreate';
import BlueprintHelmCreateHelm from './BlueprintHelmCreateHelm';
import BlueprintKubernetesCreate from './BlueprintKubernetesCreate';
import BlueprintKubernetesCreateKubernetes from './BlueprintKubernetesCreateKubernetes';
import BlueprintMorpheusCreate from './BlueprintMorpheusCreate';
import BlueprintTerraformCreate from './BlueprintTerraformCreate';
import BlueprintTerraformCreateConfig from './BlueprintTerraformCreateConfig';
import BlueprintTerraformCreateTerraform from './BlueprintTerraformCreateTerraform';

/**
 * The AddBlueprintRequest model module.
 * @module model/AddBlueprintRequest
 * @version 6.1.1
 */
class AddBlueprintRequest {
    /**
     * Constructs a new <code>AddBlueprintRequest</code>.
     * @alias module:model/AddBlueprintRequest
     * @param {(module:model/BlueprintARMCreate|module:model/BlueprintCFTCreate|module:model/BlueprintHelmCreate|module:model/BlueprintKubernetesCreate|module:model/BlueprintMorpheusCreate|module:model/BlueprintTerraformCreate)} instance The actual instance to initialize AddBlueprintRequest.
     */
    constructor(instance = null) {
        if (instance === null) {
            this.actualInstance = null;
            return;
        }
        var match = 0;
        var errorMessages = [];
        try {
            if (typeof instance === "BlueprintARMCreate") {
                this.actualInstance = instance;
            } else {
                // plain JS object
                // validate the object
                BlueprintARMCreate.validateJSON(instance); // throw an exception if no match
                // create BlueprintARMCreate from JS object
                this.actualInstance = BlueprintARMCreate.constructFromObject(instance);
            }
            match++;
        } catch(err) {
            // json data failed to deserialize into BlueprintARMCreate
            errorMessages.push("Failed to construct BlueprintARMCreate: " + err)
        }

        try {
            if (typeof instance === "BlueprintCFTCreate") {
                this.actualInstance = instance;
            } else {
                // plain JS object
                // validate the object
                BlueprintCFTCreate.validateJSON(instance); // throw an exception if no match
                // create BlueprintCFTCreate from JS object
                this.actualInstance = BlueprintCFTCreate.constructFromObject(instance);
            }
            match++;
        } catch(err) {
            // json data failed to deserialize into BlueprintCFTCreate
            errorMessages.push("Failed to construct BlueprintCFTCreate: " + err)
        }

        try {
            if (typeof instance === "BlueprintHelmCreate") {
                this.actualInstance = instance;
            } else {
                // plain JS object
                // validate the object
                BlueprintHelmCreate.validateJSON(instance); // throw an exception if no match
                // create BlueprintHelmCreate from JS object
                this.actualInstance = BlueprintHelmCreate.constructFromObject(instance);
            }
            match++;
        } catch(err) {
            // json data failed to deserialize into BlueprintHelmCreate
            errorMessages.push("Failed to construct BlueprintHelmCreate: " + err)
        }

        try {
            if (typeof instance === "BlueprintKubernetesCreate") {
                this.actualInstance = instance;
            } else {
                // plain JS object
                // validate the object
                BlueprintKubernetesCreate.validateJSON(instance); // throw an exception if no match
                // create BlueprintKubernetesCreate from JS object
                this.actualInstance = BlueprintKubernetesCreate.constructFromObject(instance);
            }
            match++;
        } catch(err) {
            // json data failed to deserialize into BlueprintKubernetesCreate
            errorMessages.push("Failed to construct BlueprintKubernetesCreate: " + err)
        }

        try {
            if (typeof instance === "BlueprintMorpheusCreate") {
                this.actualInstance = instance;
            } else {
                // plain JS object
                // validate the object
                BlueprintMorpheusCreate.validateJSON(instance); // throw an exception if no match
                // create BlueprintMorpheusCreate from JS object
                this.actualInstance = BlueprintMorpheusCreate.constructFromObject(instance);
            }
            match++;
        } catch(err) {
            // json data failed to deserialize into BlueprintMorpheusCreate
            errorMessages.push("Failed to construct BlueprintMorpheusCreate: " + err)
        }

        try {
            if (typeof instance === "BlueprintTerraformCreate") {
                this.actualInstance = instance;
            } else {
                // plain JS object
                // validate the object
                BlueprintTerraformCreate.validateJSON(instance); // throw an exception if no match
                // create BlueprintTerraformCreate from JS object
                this.actualInstance = BlueprintTerraformCreate.constructFromObject(instance);
            }
            match++;
        } catch(err) {
            // json data failed to deserialize into BlueprintTerraformCreate
            errorMessages.push("Failed to construct BlueprintTerraformCreate: " + err)
        }

        if (match > 1) {
            throw new Error("Multiple matches found constructing `AddBlueprintRequest` with oneOf schemas BlueprintARMCreate, BlueprintCFTCreate, BlueprintHelmCreate, BlueprintKubernetesCreate, BlueprintMorpheusCreate, BlueprintTerraformCreate. Input: " + JSON.stringify(instance));
        } else if (match === 0) {
            this.actualInstance = null; // clear the actual instance in case there are multiple matches
            throw new Error("No match found constructing `AddBlueprintRequest` with oneOf schemas BlueprintARMCreate, BlueprintCFTCreate, BlueprintHelmCreate, BlueprintKubernetesCreate, BlueprintMorpheusCreate, BlueprintTerraformCreate. Details: " +
                            errorMessages.join(", "));
        } else { // only 1 match
            // the input is valid
        }
    }

    /**
     * Constructs a <code>AddBlueprintRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AddBlueprintRequest} obj Optional instance to populate.
     * @return {module:model/AddBlueprintRequest} The populated <code>AddBlueprintRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        return new AddBlueprintRequest(data);
    }

    /**
     * Gets the actual instance, which can be <code>BlueprintARMCreate</code>, <code>BlueprintCFTCreate</code>, <code>BlueprintHelmCreate</code>, <code>BlueprintKubernetesCreate</code>, <code>BlueprintMorpheusCreate</code>, <code>BlueprintTerraformCreate</code>.
     * @return {(module:model/BlueprintARMCreate|module:model/BlueprintCFTCreate|module:model/BlueprintHelmCreate|module:model/BlueprintKubernetesCreate|module:model/BlueprintMorpheusCreate|module:model/BlueprintTerraformCreate)} The actual instance.
     */
    getActualInstance() {
        return this.actualInstance;
    }

    /**
     * Sets the actual instance, which can be <code>BlueprintARMCreate</code>, <code>BlueprintCFTCreate</code>, <code>BlueprintHelmCreate</code>, <code>BlueprintKubernetesCreate</code>, <code>BlueprintMorpheusCreate</code>, <code>BlueprintTerraformCreate</code>.
     * @param {(module:model/BlueprintARMCreate|module:model/BlueprintCFTCreate|module:model/BlueprintHelmCreate|module:model/BlueprintKubernetesCreate|module:model/BlueprintMorpheusCreate|module:model/BlueprintTerraformCreate)} obj The actual instance.
     */
    setActualInstance(obj) {
       this.actualInstance = AddBlueprintRequest.constructFromObject(obj).getActualInstance();
    }

    /**
     * Returns the JSON representation of the actual instance.
     * @return {string}
     */
    toJSON = function(){
        return this.getActualInstance();
    }

    /**
     * Create an instance of AddBlueprintRequest from a JSON string.
     * @param {string} json_string JSON string.
     * @return {module:model/AddBlueprintRequest} An instance of AddBlueprintRequest.
     */
    static fromJSON = function(json_string){
        return AddBlueprintRequest.constructFromObject(JSON.parse(json_string));
    }
}

/**
 * A name for the blueprint
 * @member {String} name
 */
AddBlueprintRequest.prototype['name'] = undefined;

/**
 * Path to display image. Defaults to an internal Morpheus image.
 * @member {String} image
 */
AddBlueprintRequest.prototype['image'] = undefined;

/**
 * Blueprint Type
 * @member {module:model/AddBlueprintRequest.TypeEnum} type
 */
AddBlueprintRequest.prototype['type'] = undefined;

/**
 * Array of label strings, can be used for filtering.
 * @member {Array.<String>} labels
 */
AddBlueprintRequest.prototype['labels'] = undefined;

/**
 * @member {module:model/BlueprintARMCreateArm} arm
 */
AddBlueprintRequest.prototype['arm'] = undefined;

/**
 * @member {module:model/BlueprintCFTCreateCloudFormation} cloudFormation
 */
AddBlueprintRequest.prototype['cloudFormation'] = undefined;

/**
 * @member {module:model/BlueprintHelmCreateHelm} helm
 */
AddBlueprintRequest.prototype['helm'] = undefined;

/**
 * @member {module:model/BlueprintKubernetesCreateKubernetes} kubernetes
 */
AddBlueprintRequest.prototype['kubernetes'] = undefined;

/**
 * @member {module:model/BlueprintTerraformCreateConfig} config
 */
AddBlueprintRequest.prototype['config'] = undefined;

/**
 * Tier definitions - Create in UI to view a baseline for object
 * @member {Object} tiers
 */
AddBlueprintRequest.prototype['tiers'] = undefined;

/**
 * @member {module:model/BlueprintTerraformCreateTerraform} terraform
 */
AddBlueprintRequest.prototype['terraform'] = undefined;


AddBlueprintRequest.OneOf = ["BlueprintARMCreate", "BlueprintCFTCreate", "BlueprintHelmCreate", "BlueprintKubernetesCreate", "BlueprintMorpheusCreate", "BlueprintTerraformCreate"];

export default AddBlueprintRequest;
