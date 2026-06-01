#!/usr/bin/env python3
"""
HTTP-over-vsock server — runs inside an AWS Nitro Enclave.
Listens on vsock (CID_ANY, 5000).
"""
import socket
import socketserver
from http.server import BaseHTTPRequestHandler

VSOCK_PORT = 5000
VMADDR_CID_ANY = getattr(socket, "VMADDR_CID_ANY", 0xFFFFFFFF)


class VsockHTTPServer(socketserver.ThreadingMixIn, socketserver.TCPServer):
    address_family = socket.AF_VSOCK
    daemon_threads = True


class Handler(BaseHTTPRequestHandler):
    def address_string(self):
        return f"cid={self.client_address[0]}"

    def do_GET(self):
        body = f"hello from enclave! path={self.path}\n".encode()
        self._reply(200, body)

    def do_POST(self):
        n = int(self.headers.get("Content-Length", 0))
        data = self.rfile.read(n) if n else b""
        body = f"received {len(data)} bytes: {data!r}\n".encode()
        self._reply(200, body)

    def _reply(self, code, body):
        self.send_response(code)
        self.send_header("Content-Type", "text/plain")
        self.send_header("Content-Length", str(len(body)))
        self.end_headers()
        self.wfile.write(body)


if __name__ == "__main__":
    addr = (VMADDR_CID_ANY, VSOCK_PORT)
    print(f"listening on vsock {addr}")
    with VsockHTTPServer(addr, Handler) as srv:
        srv.serve_forever()
