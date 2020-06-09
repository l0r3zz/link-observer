FROM python:3
ADD link-observer.py /
RUN pip install prometheus_client
RUN pip install requests
RUN pip install git+git://github.com/broadinstitute/cromwell-tools.git@v0.3.1
#ENTRYPOINT [ "python" ]
CMD [ "python", "./link-observer.py" ]
