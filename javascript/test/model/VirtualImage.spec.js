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

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.MorpheusApi);
  }
}(this, function(expect, MorpheusApi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new MorpheusApi.VirtualImage();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('VirtualImage', function() {
    it('should create an instance of VirtualImage', function() {
      // uncomment below and update the code to test VirtualImage
      //var instane = new MorpheusApi.VirtualImage();
      //expect(instance).to.be.a(MorpheusApi.VirtualImage);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property labels (base name: "labels")', function() {
      // uncomment below and update the code to test the property labels
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property ownerId (base name: "ownerId")', function() {
      // uncomment below and update the code to test the property ownerId
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property tenant (base name: "tenant")', function() {
      // uncomment below and update the code to test the property tenant
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property imageType (base name: "imageType")', function() {
      // uncomment below and update the code to test the property imageType
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property userUploaded (base name: "userUploaded")', function() {
      // uncomment below and update the code to test the property userUploaded
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property userDefined (base name: "userDefined")', function() {
      // uncomment below and update the code to test the property userDefined
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property systemImage (base name: "systemImage")', function() {
      // uncomment below and update the code to test the property systemImage
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property isCloudInit (base name: "isCloudInit")', function() {
      // uncomment below and update the code to test the property isCloudInit
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property sshUsername (base name: "sshUsername")', function() {
      // uncomment below and update the code to test the property sshUsername
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property sshPassword (base name: "sshPassword")', function() {
      // uncomment below and update the code to test the property sshPassword
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property sshPasswordHash (base name: "sshPasswordHash")', function() {
      // uncomment below and update the code to test the property sshPasswordHash
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property sshKey (base name: "sshKey")', function() {
      // uncomment below and update the code to test the property sshKey
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property osType (base name: "osType")', function() {
      // uncomment below and update the code to test the property osType
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property minRam (base name: "minRam")', function() {
      // uncomment below and update the code to test the property minRam
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property minRamGB (base name: "minRamGB")', function() {
      // uncomment below and update the code to test the property minRamGB
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property minDisk (base name: "minDisk")', function() {
      // uncomment below and update the code to test the property minDisk
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property minDiskGB (base name: "minDiskGB")', function() {
      // uncomment below and update the code to test the property minDiskGB
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property rawSize (base name: "rawSize")', function() {
      // uncomment below and update the code to test the property rawSize
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property rawSizeGB (base name: "rawSizeGB")', function() {
      // uncomment below and update the code to test the property rawSizeGB
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property trialVersion (base name: "trialVersion")', function() {
      // uncomment below and update the code to test the property trialVersion
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property virtioSupported (base name: "virtioSupported")', function() {
      // uncomment below and update the code to test the property virtioSupported
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property uefi (base name: "uefi")', function() {
      // uncomment below and update the code to test the property uefi
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property isAutoJoinDomain (base name: "isAutoJoinDomain")', function() {
      // uncomment below and update the code to test the property isAutoJoinDomain
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property vmToolsInstalled (base name: "vmToolsInstalled")', function() {
      // uncomment below and update the code to test the property vmToolsInstalled
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property installAgent (base name: "installAgent")', function() {
      // uncomment below and update the code to test the property installAgent
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property isForceCustomization (base name: "isForceCustomization")', function() {
      // uncomment below and update the code to test the property isForceCustomization
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property isSysprep (base name: "isSysprep")', function() {
      // uncomment below and update the code to test the property isSysprep
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property fipsEnabled (base name: "fipsEnabled")', function() {
      // uncomment below and update the code to test the property fipsEnabled
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property userData (base name: "userData")', function() {
      // uncomment below and update the code to test the property userData
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property consoleKeymap (base name: "consoleKeymap")', function() {
      // uncomment below and update the code to test the property consoleKeymap
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property storageProvider (base name: "storageProvider")', function() {
      // uncomment below and update the code to test the property storageProvider
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property visibility (base name: "visibility")', function() {
      // uncomment below and update the code to test the property visibility
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property accounts (base name: "accounts")', function() {
      // uncomment below and update the code to test the property accounts
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property volumes (base name: "volumes")', function() {
      // uncomment below and update the code to test the property volumes
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property storageControllers (base name: "storageControllers")', function() {
      // uncomment below and update the code to test the property storageControllers
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property networkInterfaces (base name: "networkInterfaces")', function() {
      // uncomment below and update the code to test the property networkInterfaces
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property tags (base name: "tags")', function() {
      // uncomment below and update the code to test the property tags
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property locations (base name: "locations")', function() {
      // uncomment below and update the code to test the property locations
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.VirtualImage();
      //expect(instance).to.be();
    });

  });

}));
