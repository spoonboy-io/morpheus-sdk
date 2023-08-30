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
from openapi_client.model.hub_register_object_hub import HubRegisterObjectHub
from openapi_client.model.setup_request_any_of import SetupRequestAnyOf
from openapi_client.model.setup_request_any_of1 import SetupRequestAnyOf1
globals()['HubRegisterObjectHub'] = HubRegisterObjectHub
globals()['SetupRequestAnyOf'] = SetupRequestAnyOf
globals()['SetupRequestAnyOf1'] = SetupRequestAnyOf1
from openapi_client.model.setup_request import SetupRequest


class TestSetupRequest(unittest.TestCase):
    """SetupRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testSetupRequest(self):
        """Test SetupRequest"""
        # FIXME: construct object with mandatory attributes with example values
        # model = SetupRequest()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
