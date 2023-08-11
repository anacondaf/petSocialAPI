FROM docker.elastic.co/beats/metricbeat:8.8.1
COPY ./elastic/certs /usr/share/metricbeat/certs
COPY ./metricbeat.yml /usr/share/metricbeat/metricbeat.yml
USER root
RUN chown root /usr/share/metricbeat/metricbeat.yml