import socket
import hashlib
import argparse

HOST = "0.0.0.0"  # The server's hostname or IP address
PORT = 13564  # The port used by the server

operations = {
    "write": 0b00100100,
    "fetch": 0b00100010,
    "update": 0b00100100, # same as write
    "delete": 0b00100110,
}

owner = ""

def str_to_hex(s: str):
    r = ""

    for c in s:
        r += hex(c)[2:]

    return r

def byte_array_to_hex(byte_array):
    hex_string = ''.join(['{:02x}'.format(b) for b in byte_array])
    return hex_string

def construct_packet(operation: int, owner_token: str, content: bytes):
    b = bytearray()
    cb = bytearray(content)

    print(f"Owner: {owner_token}, len: {len(owner_token)}")

    if len(owner_token) != 64:
        print("Error: owner token must be 64 bytes long")
        return None

    if len(content) < 2048:
        rem = 2048 - len(content)
        i = 0
        while i < rem:
            cb.append(0)
            i += 1

    b.extend(owner_token)
    b.extend(cb)

    r = [
        operation
    ]
    r.extend(b)

    return bytes(r)

def write_data(connection: socket.socket, owner_token: str, content: str) -> bytes:
    pkt = construct_packet(operations["write"], owner_token, bytearray(content.encode("ISO-8859-1")))

    connection.sendall(pkt)

    response = connection.recv(2048)

    response_code = response[0]
    response_status = response[1]
    response_cluster = response[2]
    response_data = response[3:]

    print(f"response code: {response_code}, response valid: {response_code == (operations['write'] + 1)}, ok: {response_status == 1}")

    return response_cluster, response_data

def read_data(connection: socket.socket, owner_token: str, content: str) -> bytes:
    print(len(bytes.fromhex(content).decode("ISO-8859-1")))
    pkt = construct_packet(operations["fetch"], owner_token, bytes.fromhex(content))

    connection.sendall(pkt)

    response = connection.recv(2048)

    response_code = response[0]
    response_status = response[1]
    response_data = response[2:]

    print(f"response code: {response_code}, response valid: {response_code == (operations['fetch'] + 1)}, ok: {response_status == 1}")

    return response_data

def delete_data(connection: socket.socket, owner_token: str, content: str) -> bytes:
    print(len(bytes.fromhex(content).decode("ISO-8859-1")))
    pkt = construct_packet(operations["delete"], owner_token, bytes.fromhex(content))

    connection.sendall(pkt)

    response = connection.recv(2048)

    response_code = response[0]
    response_status = response[1]
    response_data = response[2:]

    print(f"response code: {response_code}, response valid: {response_code == (operations['delete'] + 1)}, ok: {response_status == 1}")

    return response_data


def main():
    parser = argparse.ArgumentParser(description="A simple command-line program.")
    
    # Define command-line arguments
    parser.add_argument('-o', '--owner',  help='Owner username')
    parser.add_argument('-d', '--data',   help='Optional data to write')
    parser.add_argument('-r', '--record', help='The record has you want to perform an operation on (read | delete)')
    parser.add_argument('-m', '--method', help='Nomad method to be used (fetch | write | delete)')

    # Parse the command-line arguments
    args = parser.parse_args()

    # Access the values
    method = args.method or 'write'
    record = args.record
    content = args.data or ''
    owner = hashlib.sha512(args.owner.encode('ISO-8859-1')).digest() if args.owner != None else bytes(chr(0) * 64, "ISO-8859-1")

    # Your program logic goes here
    print("Method:", method)
    print("Record:", record)
    print("Content", content)
    print("Owner:", owner)

    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect((HOST, PORT))


    if method == "write":
        c, r = write_data(s, owner, content)
        print(c, r.hex(), len(r))
    elif method == "fetch":
        r = read_data(s, owner, record)
        print(r)
    elif method == "delete":
        r = delete_data(s, owner, record)
        print(r)

if __name__ == "__main__":
    main()