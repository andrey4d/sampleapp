#!/usr/bin/env python3

import json
import os
from http.server import HTTPServer, BaseHTTPRequestHandler
from socketserver import ThreadingMixIn
import signal
import socket
from urllib import request


class ThreadingSimpleServer(ThreadingMixIn,HTTPServer):
    pass


def get_backend_info(url):
    response = request.urlopen(url)
    if response.code == 200:
        data = response.read()
        return json.loads(data.decode('utf-8'))
    return {}


class SundboxServer(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path in ["/", "/info"]:
            self.header("text/html")
            self.print_index(get_backend_info(f"{os.getenv('BACKEND_URL')}/info"))
        elif self.path == "/health":
            self.header("text/html")
            try:
                response = request.urlopen(f"{os.getenv('BACKEND_URL')}/health")
                if response.code == 200:
                    data = json.loads(response.read().decode('utf-8'))
                    self.html_tag_print(data.get('health'))
            except Exception :
                self.html_tag_print("unhealthy")
        elif self.path == "/about":
            self.header("text/html")
            self.print_about(get_backend_info(f"{os.getenv('BACKEND_URL')}/about"))
        else:
            self.send_error(404)

    def print_index(self, rest_data: dict):
        self.html_tag_print('<html><head><meta charset="utf-8">')
        self.style()
        self.html_tag_print('<title>Простой HTTP-сервер.</title></head>')
        self.html_tag_print('<body>')
        self.html_tag_print(f'<h2>FRONT_SERVER HOSTNAME: {socket.gethostname()}</h2>')
        self.html_tag_print(f'<h2>BACK__SERVER HOSTNAME: {rest_data.get("backend_hostname")}</h2>')
        self.print_index_all_env(rest_data['environment'])
        self.footer()
        self.html_tag_print('</body></html>')

    def print_about(self, rest_data: dict):
        self.html_tag_print('<html><head><meta charset="utf-8">')
        self.style()
        self.html_tag_print('<title>Простой HTTP-сервер.</title></head>')
        self.html_tag_print('<body>')
        self.print_index_all_env(rest_data)
        self.footer()
        self.html_tag_print('</body></html>')

    def print_index_all_env(self, rest_data: dict):
        self.html_tag_print('<table border="1"> \
                            <caption>HOST VARIABLES</caption> \
                            <tr><th>Name</th><th>Value</th></tr>')
        for name in rest_data.keys():
            value = rest_data.get(name)
            self.html_tag_print(f"<tr><td>{name}</td><td>{value}</td></tr>")
        self.html_tag_print('</table>')

    def html_tag_print(self, html_tag: str):
        self.wfile.write(f'{html_tag}'.encode())

    def style(self):
        self.html_tag_print('<style>footer \
                            { text-align: center; padding: 3px; background-color: DarkSalmon; color: white;}\
                            </style>')

    def header(self, content_type: str):
        self.send_response(200)
        self.send_header("Content-type", content_type)
        self.end_headers()

    def footer(self):
        self.html_tag_print('<footer> \
                            <p>Сервисная команда по Контейнеризации 2023.</p> \
                            <p><a href="mailto:andrey4d.dev@gmail.com">andrey4d.dev@gmail.com</a></p> \
                            </footer>')


def signal_handler(signal, frame):
    print('Stop HTTP')
    exit()


def main(handler_class=SundboxServer, addr="0.0.0.0", port=80):
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

