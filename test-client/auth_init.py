import socket
import struct

HOST = "127.0.0.1"  # The server's hostname or IP address
PORT = 37806  # The port used by the server

operations = {
    "add": 0,
    "fetch": 1,
    "update": 2,
    "delete": 3,

    "auth": 0x10,
    "db_auth": 0x20,

    "end": 255
}
token = "0" * 63 + "3"
data = "0"* 2047 + "1"

b = bytearray()
b.extend(map(ord, token))
b.extend(map(ord, data))

auth = [
    0b00100100,
]

auth.extend(b)
auth.append(0b00110011)


with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:

    s.connect((HOST, PORT))

    while True:
        d = input("> ")
        op, data = d.split(" ")

        if op in operations.keys():
            print(op, data)
            print(auth)
            s.sendall(bytes(auth))
            data = s.recv(2048)

            print(f"Received {data!r}")