import socket
s = socket.socket(socket.AF_VSOCK, socket.SOCK_STREAM)
s.connect((16, 5000))    # CID 换成 nitro-cli 给的
s.sendall(b"GET /hello HTTP/1.1\r\nHost: x\r\nConnection: close\r\n\r\n")
data = b""
while chunk := s.recv(4096):
    data += chunk
print(data.decode())
