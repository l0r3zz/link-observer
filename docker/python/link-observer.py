from prometheus_client import start_http_server, Summary
import random
import time
import requests

urls = ['https://httpstat.us/503','https://httpstat.us/200']

# Create a metric to track time spent and requests made.
REQUEST_SUMMARY = Summary('sample_external_url_response_ms', 'Time spent processing request',['url'])

# Decorate function with metric.
#@REQUEST_SUMMARY.time()
def process_request(url):
    with REQUEST_SUMMARY.labels(url).time():
        response = requests.get(url)
        return response.status_code

if __name__ == '__main__':
    # Start up the server to expose the metrics.
    start_http_server(8001)
    # Generate some requests.
    while True:
        for url in urls:
            process_request(url)
