# link-observer
Repo to demonstrate modern DevOps/SRE techniques.
This took a little longer than anticipated. In my normal job, the Prometheus and Graphana infrastructure is prebuilt, and we instrument with **statsd** and not the Prometheus client library directly. This currently runs on my local Ubuntu Linux workstation. The service is containerized, but currently ithe Prometheus and Grafana infrastructure is not, I have a GKE cluster built and have Prometheus deployed but have not finished that part. Considering that I'm overdue, I wanted to submit this MVP and continue working. Once I get all of the infrastructure built out, I will create a golang and rust version of the service.

Note that you can add additional URLs to the *urls* array and they will be monitored as well


## Screenshots

### Link Observability Dashboard
<img width="964" alt="Link observability Dashboard" src="https://github.com/l0r3zz/link-observer/blob/master/screenshots/Link%20Observability%20Dashboard%202020-06-11%2011-15-08.png">

### Local Prometheus Targets
<img width="964" alt="Local Prometheus Server Targets" src="https://github.com/l0r3zz/link-observer/blob/master/screenshots/Local%20Prometheus%20Targets%202020-06-11%2009-27-52.png">

## Some informative links:
1.  How To Install and Configure Prometheus On a Linux Server (https://devopscube.com/install-configure-prometheus-linux/).
2.  How to set up Prometheus Monitoring on Kubernetes Clusters (https://devopscube.com/setup-prometheus-monitoring-on-kubernetes/).
3.  Prometheus Python Client Repo (https://github.com/prometheus/client_python).
4.  Grafana Getting Started (https://grafana.com/docs/grafana/latest/getting-started/getting-started/).
5.  How to instrument code: Custom metrics vs APM vs OpenTracing (https://sysdig.com/blog/how-to-instrument-code-custom-metrics-vs-apm-vs-opentracing/).
6.  Instrumenting Python with Prometheus (https://www.robustperception.io/instrumenting-python-with-prometheus).
7.  Prometheus Installation: https://prometheus.io/docs/prometheus/latest/installation/).
8.  Prometheus metrics / OpenMetrics code instrumentation.;https://sysdig.com/blog/prometheus-metrics/).
9.  Prometheus (https://prometheus.io/docs/instrumenting/clientlibs/).
10. Prometheus Getting Started (https://prometheus.io/docs/prometheus/latest/getting_started/).
11. Measuring Latency Using Prometheus (https://medium.com/@benpourian/measuring-latency-using-prometheus-d3b3fe1cac57).
12. Prometheus Up and Running (https://books.google.com/books/about/Prometheus_Up_and_Running.html?id=mdv2tQEACAAJ).
13. Monitoring with Prometheus {https://www.prometheusbook.com/).
