"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import openapi_client
from openapi_client.model.license_current_usage import LicenseCurrentUsage
from openapi_client.model.license_license import LicenseLicense
globals()['LicenseCurrentUsage'] = LicenseCurrentUsage
globals()['LicenseLicense'] = LicenseLicense
from openapi_client.model.license import License


class TestLicense(unittest.TestCase):
    """License unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testLicense(self):
        """Test License"""
        # FIXME: construct object with mandatory attributes with example values
        # model = License()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
