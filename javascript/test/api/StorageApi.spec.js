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
    instance = new MorpheusApi.StorageApi();
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

  describe('StorageApi', function() {
    describe('addStorageBuckets', function() {
      it('should call addStorageBuckets successfully', function(done) {
        //uncomment below and update the code to test addStorageBuckets
        //instance.addStorageBuckets(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('addStorageServers', function() {
      it('should call addStorageServers successfully', function(done) {
        //uncomment below and update the code to test addStorageServers
        //instance.addStorageServers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('addStorageVolumes', function() {
      it('should call addStorageVolumes successfully', function(done) {
        //uncomment below and update the code to test addStorageVolumes
        //instance.addStorageVolumes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getStorageBuckets', function() {
      it('should call getStorageBuckets successfully', function(done) {
        //uncomment below and update the code to test getStorageBuckets
        //instance.getStorageBuckets(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getStorageServerTypes', function() {
      it('should call getStorageServerTypes successfully', function(done) {
        //uncomment below and update the code to test getStorageServerTypes
        //instance.getStorageServerTypes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getStorageServers', function() {
      it('should call getStorageServers successfully', function(done) {
        //uncomment below and update the code to test getStorageServers
        //instance.getStorageServers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getStorageVolumeTypes', function() {
      it('should call getStorageVolumeTypes successfully', function(done) {
        //uncomment below and update the code to test getStorageVolumeTypes
        //instance.getStorageVolumeTypes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getStorageVolumes', function() {
      it('should call getStorageVolumes successfully', function(done) {
        //uncomment below and update the code to test getStorageVolumes
        //instance.getStorageVolumes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listStorageBuckets', function() {
      it('should call listStorageBuckets successfully', function(done) {
        //uncomment below and update the code to test listStorageBuckets
        //instance.listStorageBuckets(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listStorageServerTypes', function() {
      it('should call listStorageServerTypes successfully', function(done) {
        //uncomment below and update the code to test listStorageServerTypes
        //instance.listStorageServerTypes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listStorageServers', function() {
      it('should call listStorageServers successfully', function(done) {
        //uncomment below and update the code to test listStorageServers
        //instance.listStorageServers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listStorageVolumeTypes', function() {
      it('should call listStorageVolumeTypes successfully', function(done) {
        //uncomment below and update the code to test listStorageVolumeTypes
        //instance.listStorageVolumeTypes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listStorageVolumes', function() {
      it('should call listStorageVolumes successfully', function(done) {
        //uncomment below and update the code to test listStorageVolumes
        //instance.listStorageVolumes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeStorageBuckets', function() {
      it('should call removeStorageBuckets successfully', function(done) {
        //uncomment below and update the code to test removeStorageBuckets
        //instance.removeStorageBuckets(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeStorageServers', function() {
      it('should call removeStorageServers successfully', function(done) {
        //uncomment below and update the code to test removeStorageServers
        //instance.removeStorageServers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeStorageVolumes', function() {
      it('should call removeStorageVolumes successfully', function(done) {
        //uncomment below and update the code to test removeStorageVolumes
        //instance.removeStorageVolumes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateStorageBuckets', function() {
      it('should call updateStorageBuckets successfully', function(done) {
        //uncomment below and update the code to test updateStorageBuckets
        //instance.updateStorageBuckets(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateStorageServers', function() {
      it('should call updateStorageServers successfully', function(done) {
        //uncomment below and update the code to test updateStorageServers
        //instance.updateStorageServers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateStorageVolumes', function() {
      it('should call updateStorageVolumes successfully', function(done) {
        //uncomment below and update the code to test updateStorageVolumes
        //instance.updateStorageVolumes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));