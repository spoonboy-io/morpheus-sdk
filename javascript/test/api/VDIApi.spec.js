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
    instance = new MorpheusApi.VDIApi();
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

  describe('VDIApi', function() {
    describe('addVDIApps', function() {
      it('should call addVDIApps successfully', function(done) {
        //uncomment below and update the code to test addVDIApps
        //instance.addVDIApps(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('addVDIGateways', function() {
      it('should call addVDIGateways successfully', function(done) {
        //uncomment below and update the code to test addVDIGateways
        //instance.addVDIGateways(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('addVDIPools', function() {
      it('should call addVDIPools successfully', function(done) {
        //uncomment below and update the code to test addVDIPools
        //instance.addVDIPools(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('addVdiAllocation', function() {
      it('should call addVdiAllocation successfully', function(done) {
        //uncomment below and update the code to test addVdiAllocation
        //instance.addVdiAllocation(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getVDIAllocations', function() {
      it('should call getVDIAllocations successfully', function(done) {
        //uncomment below and update the code to test getVDIAllocations
        //instance.getVDIAllocations(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getVDIApps', function() {
      it('should call getVDIApps successfully', function(done) {
        //uncomment below and update the code to test getVDIApps
        //instance.getVDIApps(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getVDIGateways', function() {
      it('should call getVDIGateways successfully', function(done) {
        //uncomment below and update the code to test getVDIGateways
        //instance.getVDIGateways(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getVDIPools', function() {
      it('should call getVDIPools successfully', function(done) {
        //uncomment below and update the code to test getVDIPools
        //instance.getVDIPools(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getVdi', function() {
      it('should call getVdi successfully', function(done) {
        //uncomment below and update the code to test getVdi
        //instance.getVdi(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listVDIAllocations', function() {
      it('should call listVDIAllocations successfully', function(done) {
        //uncomment below and update the code to test listVDIAllocations
        //instance.listVDIAllocations(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listVDIApps', function() {
      it('should call listVDIApps successfully', function(done) {
        //uncomment below and update the code to test listVDIApps
        //instance.listVDIApps(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listVDIGateways', function() {
      it('should call listVDIGateways successfully', function(done) {
        //uncomment below and update the code to test listVDIGateways
        //instance.listVDIGateways(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listVDIPools', function() {
      it('should call listVDIPools successfully', function(done) {
        //uncomment below and update the code to test listVDIPools
        //instance.listVDIPools(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listVdi', function() {
      it('should call listVdi successfully', function(done) {
        //uncomment below and update the code to test listVdi
        //instance.listVdi(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeVDIApps', function() {
      it('should call removeVDIApps successfully', function(done) {
        //uncomment below and update the code to test removeVDIApps
        //instance.removeVDIApps(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeVDIGateways', function() {
      it('should call removeVDIGateways successfully', function(done) {
        //uncomment below and update the code to test removeVDIGateways
        //instance.removeVDIGateways(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeVDIPools', function() {
      it('should call removeVDIPools successfully', function(done) {
        //uncomment below and update the code to test removeVDIPools
        //instance.removeVDIPools(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateVDIApps', function() {
      it('should call updateVDIApps successfully', function(done) {
        //uncomment below and update the code to test updateVDIApps
        //instance.updateVDIApps(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateVDIGateways', function() {
      it('should call updateVDIGateways successfully', function(done) {
        //uncomment below and update the code to test updateVDIGateways
        //instance.updateVDIGateways(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateVDIPools', function() {
      it('should call updateVDIPools successfully', function(done) {
        //uncomment below and update the code to test updateVDIPools
        //instance.updateVDIPools(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));