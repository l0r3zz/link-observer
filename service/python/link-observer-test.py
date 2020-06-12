# Basic unit test for link-observer.py
# NOTE: there is a symbolic link hack of
# ln -s link_observer.py link-observer.py
# TODO: change the name of the service to link_observer

import unittest
from link_observer import process_request


class TestSum(unittest.TestCase):
    def test_200_true(self):
        """
        Test that it can probe the always 200 site
        """
        url = "https://httpstat.us/200"
        result = process_request(url)
        self.assertEqual(result, 200)

    def test_503_true(self):
        """
        Test that it can probe the always 503 site
        """
        url = "https://httpstat.us/503"
        result = process_request(url)
        self.assertEqual(result, 503)

    def test_incorrect_result(self):
        """
        Test that it returns correct incorrect result
        """
        url = "https://httpstat.us/503"
        result = process_request(url)
        self.assertTrue(result != 200)

if __name__ == '__main__':
    unittest.main()
