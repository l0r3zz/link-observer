from prometheus_client import start_http_server, Summary, Gauge
import random
import time
import requests

urls = ['https://httpstat.us/503','https://httpstat.us/200']

# Create a metric to track time spent and requests made.
MONITOR_PORT= 8001
REQUEST_SUMMARY = Summary('lo_external_url_response_ms', 'Time spent processing request',['url'])
RESPONSE_GAUGE = Gauge('lo_external_url_up', 'up = 1 all else 0', ['url'])

def process_request(url):
    with REQUEST_SUMMARY.labels(url).time():
        response = requests.get(url)
        return response.status_code

if __name__ == '__main__':
    # Start up the server to expose the metrics.
    start_http_server(MONITOR_PORT)
    # Generate some requests.
    while True:
        time.sleep(1)
        for url in urls:
            #RESPONSE_GAUGE_SET = RESPONSE_GAUGE.labels(url)
            resp = process_request(url)
            if resp == 200:
                RESPONSE_GAUGE.labels(url).set(1)
            else:
                RESPONSE_GAUGE.labels(url).set(0)
