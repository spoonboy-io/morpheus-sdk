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
    instance = new MorpheusApi.ContainersApi();
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

  describe('ContainersApi', function() {
    describe('cloneImageContainerAction', function() {
      it('should call cloneImageContainerAction successfully', function(done) {
        //uncomment below and update the code to test cloneImageContainerAction
        //instance.cloneImageContainerAction(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('containersAttachFloatingIp', function() {
      it('should call containersAttachFloatingIp successfully', function(done) {
        //uncomment below and update the code to test containersAttachFloatingIp
        //instance.containersAttachFloatingIp(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('containersDetachFloatingIp', function() {
      it('should call containersDetachFloatingIp successfully', function(done) {
        //uncomment below and update the code to test containersDetachFloatingIp
        //instance.containersDetachFloatingIp(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('ejectContainerAction', function() {
      it('should call ejectContainerAction successfully', function(done) {
        //uncomment below and update the code to test ejectContainerAction
        //instance.ejectContainerAction(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('executeContainerAction', function() {
      it('should call executeContainerAction successfully', function(done) {
        //uncomment below and update the code to test executeContainerAction
        //instance.executeContainerAction(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getContainer', function() {
      it('should call getContainer successfully', function(done) {
        //uncomment below and update the code to test getContainer
        //instance.getContainer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getContainerActions', function() {
      it('should call getContainerActions successfully', function(done) {
        //uncomment below and update the code to test getContainerActions
        //instance.getContainerActions(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('importContainerAction', function() {
      it('should call importContainerAction successfully', function(done) {
        //uncomment below and update the code to test importContainerAction
        //instance.importContainerAction(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('restartContainerAction', function() {
      it('should call restartContainerAction successfully', function(done) {
        //uncomment below and update the code to test restartContainerAction
        //instance.restartContainerAction(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('startContainerAction', function() {
      it('should call startContainerAction successfully', function(done) {
        //uncomment below and update the code to test startContainerAction
        //instance.startContainerAction(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('stopContainerAction', function() {
      it('should call stopContainerAction successfully', function(done) {
        //uncomment below and update the code to test stopContainerAction
        //instance.stopContainerAction(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('suspendContainerAction', function() {
      it('should call suspendContainerAction successfully', function(done) {
        //uncomment below and update the code to test suspendContainerAction
        //instance.suspendContainerAction(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
