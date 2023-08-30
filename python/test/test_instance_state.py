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
from openapi_client.model.app_state_output import AppStateOutput
from openapi_client.model.instance_state_input import InstanceStateInput
globals()['AppStateOutput'] = AppStateOutput
globals()['InstanceStateInput'] = InstanceStateInput
from openapi_client.model.instance_state import InstanceState


class TestInstanceState(unittest.TestCase):
    """InstanceState unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testInstanceState(self):
        """Test InstanceState"""
        # FIXME: construct object with mandatory attributes with example values
        # model = InstanceState()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
