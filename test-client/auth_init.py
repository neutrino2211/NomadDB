import socket
import struct

HOST = "127.0.0.1"  # The server's hostname or IP address
PORT = 47595  # The port used by the server

operations = {
    "add": 0,
    "fetch": 1,
    "update": 2,
    "delete": 3,

    "auth": 0x10,
    "db_auth": 0x20,

    "end": 255
}

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:

    s.connect((HOST, PORT))

    while True:
        d = input("> ")
        op, data = d.split(" ")

        if op in operations.keys():
            print(op, data)
            s.sendall(bytes([operations[op],*bytes(data, "utf-8")]))
            data = s.recv(2048)

            print(f"Received {data!r}")