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
from openapi_client.model.app_state import AppState
from openapi_client.model.app_state_input import AppStateInput
from openapi_client.model.app_state_output import AppStateOutput
from openapi_client.model.app_state_specs_inner import AppStateSpecsInner
from openapi_client.model.app_state_workloads_inner import AppStateWorkloadsInner
from openapi_client.model.model200_success import Model200Success
globals()['AppState'] = AppState
globals()['AppStateInput'] = AppStateInput
globals()['AppStateOutput'] = AppStateOutput
globals()['AppStateSpecsInner'] = AppStateSpecsInner
globals()['AppStateWorkloadsInner'] = AppStateWorkloadsInner
globals()['Model200Success'] = Model200Success
from openapi_client.model.get_app_state200_response import GetAppState200Response


class TestGetAppState200Response(unittest.TestCase):
    """GetAppState200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testGetAppState200Response(self):
        """Test GetAppState200Response"""
        # FIXME: construct object with mandatory attributes with example values
        # model = GetAppState200Response()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
