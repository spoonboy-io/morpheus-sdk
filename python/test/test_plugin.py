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
from openapi_client.model.app_state_specs_inner_template import AppStateSpecsInnerTemplate
from openapi_client.model.option_type import OptionType
globals()['AppStateSpecsInnerTemplate'] = AppStateSpecsInnerTemplate
globals()['OptionType'] = OptionType
from openapi_client.model.plugin import Plugin


class TestPlugin(unittest.TestCase):
    """Plugin unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testPlugin(self):
        """Test Plugin"""
        # FIXME: construct object with mandatory attributes with example values
        # model = Plugin()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
