#!/usr/bin/env python3

import os
from http.server import HTTPServer, BaseHTTPRequestHandler
from socketserver import ThreadingMixIn
import signal
import socket
import json


class ThreadingSimpleServer(ThreadingMixIn, HTTPServer):
    pass


def env_to_dict() -> dict:
    out = dict()
    out['backend_hostname'] = socket.gethostname()
    out['environment'] = {}
    for name, value in os.environ.items():
        out['environment'][name] = value
    return out


def about() -> dict:
    out = dict()
    out['/'] = "Return Ok"
    out['/info'] = "Return OS environment"
    out['/health'] = "Return healthy"
    out['/about'] = "Return this info"
    return out


class SundboxServer(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/":
            self.header("application/json")
            self.wfile.write(json.dumps({'ok': "Ok"}).encode())
        elif self.path == "/info":
            self.header("application/json")
            self.wfile.write(json.dumps(env_to_dict()).encode())
        elif self.path == "/health":
            self.header("application/json")
            self.wfile.write(json.dumps({'health': "healthy"}).encode())
        elif self.path == "/about":
            self.header("application/json")
            self.wfile.write(json.dumps(about()).encode())
        else:
            self.send_error(404)

    def header(self, content_type: str):
        self.send_response(200)
        self.send_header("Content-type", content_type)
        self.end_headers()


def signal_handler(signal, frame):
    print('Stop HTTP')
    exit()


def main(handler_class=SundboxServer, addr="0.0.0.0", port=8070):
    server_address = (addr, port)
    signal.signal(signal.SIGINT, signal_handler)
    signal.signal(signal.SIGTERM, signal_handler)
    httpd = ThreadingSimpleServer(server_address, handler_class)

    try:
        print(f'Serving HTTP on {addr}:{port}')
        httpd.serve_forever()
    except:
        httpd.shutdown()
    httpd.server_close()


if __name__ == "__main__":
    main()
