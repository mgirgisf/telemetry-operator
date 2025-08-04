#!/usr/bin/env python3
#
# Copyright 2022 Red Hat Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may
# not use this file except in compliance with the License. You may obtain
# a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.

# CloudKitty-API has its own health check endpoint and does not need this.
#
# This script performs a series of checks on the CloudKitty processor's dependencies,
#  including the database, Keystone, and the Prometheus collector.
#
# Configuration directory (defaults to /etc/cloudkitty/cloudkitty.conf)

#!/usr/bin/env python

import http.server
import signal
import socket
import sys
import time
import threading
import requests
import sqlalchemy
from oslo_config import cfg
from oslo_db.sqlalchemy import enginefacade


SERVER_PORT = 8080
CONF = cfg.CONF
SERVICE_NAME = 'cloudkitty-processor'


class HTTPServerV6(http.server.HTTPServer):
    address_family = socket.AF_INET6


class HealthCheckError(Exception):
    def __init__(self, message):
        super(HealthCheckError, self).__init__(message)



class HeartbeatServer(http.server.BaseHTTPRequestHandler):

    def perform_checks():

        print("Starting health checks")
        results = {}

        # Database Connectivity check
        try:
            db_connection_str = CONF.database.connection
            engine = enginefacade.get_engine(connection=db_connection_str)
            with engine.connect():
                results['database'] = 'OK'
                print("Database connectivity check passed.")
        except (cfg.ConfigFilesNotFoundError, sqlalchemy.exc.OperationalError) as e:
            results['database'] = 'FAIL'
            print(f"ERROR: Database connectivity check failed: {e}")
            return False, results

        # Keystone Endpoint Reachability
        try:
            keystone_uri = CONF.keystone_authtoken.auth_url
            response = requests.get(keystone_uri, timeout=5)
            response.raise_for_status()
            if 'keystone' in response.headers.get('Server', '').lower():
                results['keystone_endpoint'] = 'OK'
                print("Keystone endpoint reachable and responsive.")
            else:
                results['keystone_endpoint'] = 'WARN'
                print(f"Keystone endpoint reachable, but not a valid service: {keystone_uri}")
        except requests.exceptions.RequestException as e:
            results['keystone_endpoint'] = 'FAIL'
            print(f"ERROR: Keystone endpoint check failed: {e}")
            return False, results

        # Prometheus Collector Endpoint
        try:
            prometheus_url = CONF.collector_prometheus.prometheus_url
            insecure = CONF.collector_prometheus.insecure
            cafile = CONF.collector_prometheus.cafile
            verify_ssl = cafile if cafile and not insecure else not insecure

            response = requests.get(prometheus_url, timeout=5, verify=verify_ssl)
            response.raise_for_status()
            results['collector_endpoint'] = 'OK'
            print("Prometheus collector endpoint reachable.")
        except requests.exceptions.RequestException as e:
            results['collector_endpoint'] = 'FAIL'
            print(f"ERROR: Prometheus collector check failed: {e}")
            return False, results

        print("All  checks passed.")
        return True, results

    def do_GET(self):
        print("Received health check request.")
        try:
            # Perform all checks and respond based on the result
            success, _ = self.perform_checks()
            if success:
                self.send_response(200)
                self.send_header("Content-type", "text/html")
                self.end_headers()
                self.wfile.write('<html><body>OK</body></html>'.encode('utf-8'))
            else:
                self.send_error(500, "Health check failed", "One or more dependencies are unhealthy.")
        except Exception as e:
            self.send_error(500, "Internal Server Error", str(e))


def get_stopper(server):
    def stopper(signal_number=None, frame=None):
        print("Stopping server.")
        server.shutdown()
        server.server_close()
        print("Server stopped.")
        sys.exit(0)
    return stopper


def main():
    # Register all necessary configuration options
    cfg.CONF.register_group(cfg.OptGroup(name='database', title='Database Options'))
    cfg.CONF.register_opt(cfg.StrOpt('connection', default='sqlite:///:memory:'), group='database')

    cfg.CONF.register_group(cfg.OptGroup(name='keystone_authtoken', title='Keystone Auth Token Options'))
    cfg.CONF.register_opt(cfg.StrOpt('auth_url', default='http://localhost:5000/v3'), group='keystone_authtoken')

    cfg.CONF.register_group(cfg.OptGroup(name='collector_prometheus', title='Prometheus Collector Options'))
    cfg.CONF.register_opt(cfg.StrOpt('prometheus_url', default='http://metric-storage-prometheus.openstack.svc:9090'), group='collector_prometheus')
    cfg.CONF.register_opt(cfg.BoolOpt('insecure', default=False), group='collector_prometheus')
    cfg.CONF.register_opt(cfg.StrOpt('cafile', default=None), group='collector_prometheus')


    # Load the configuration file.
    try:
        cfg.CONF(sys.argv[1:], default_config_files=['/etc/cloudkitty/cloudkitty.conf'])
    except cfg.ConfigFilesNotFoundError as e:
        print(f"Health check failed: {e}", file=sys.stderr)
        sys.exit(1)


    # Start the HTTP server
    hostname = socket.gethostname()
    ipv6_address = socket.getaddrinfo(hostname, None, socket.AF_INET6)
    if ipv6_address:
        webServer = HTTPServerV6(("::", SERVER_PORT), HeartbeatServer)
    else:
        webServer = http.server.HTTPServer(("0.0.0.0", SERVER_PORT), HeartbeatServer)

    stop = get_stopper(webServer)
    thread = threading.Thread(target=webServer.serve_forever)
    thread.daemon = True
    thread.start()
    print(f"CloudKitty Processor Healthcheck Server started http://{hostname}:{SERVER_PORT}")
    signal.signal(signal.SIGTERM, stop)

    try:
        while True:
            time.sleep(60)
    except KeyboardInterrupt:
        pass
    finally:
        stop()


if __name__ == "__main__":
    main()