import pyglet
import io
import struct
import socket
import threading
import sys

HOST = "127.0.0.1"  # The server's hostname or IP address
PORT = 8090  # The port used by the server

operations = {
    "add": 0,
    "fetch": 1,
    "update": 2,
    "delete": 3,

    "end": 255
}


addresses = []

def get_file_chunks(file):
    chunks = []
    with open(file, "rb") as f:
        while True:
            data = f.read(2048)
            if len(data) == 0:
                return chunks
            
            chunks.append(data)

def send_chunk(conn, chunk):
    data = struct.pack("s2048s", bytes([operations["add"]]), chunk)
    conn.sendall(data)
    address = conn.recv(2048)
    addresses.append(address)

def read_chunk(conn, address):
    address = struct.pack("s2048s", bytes([operations["fetch"]]), address)
    conn.sendall(address)
    data = conn.recv(2048)
    return data

def buffer_file(_, addresses):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect((HOST, PORT))
        for address in addresses:
            data = read_chunk(s, address)
            print("wrote", len(data))
            memory_file.write(data)

        s.close()

memory_file = io.BytesIO()

vidPath = sys.argv[1] or '/home/tsowamainasara/Downloads/AnimePahe_Spy_x_Family_-_19_360p_SubsPlease.mp4'
with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:

    s.connect((HOST, PORT))
    
    for chunk in get_file_chunks(vidPath):
        send_chunk(s, chunk)
    s.close()



with open("addresses", "w") as addr:
    addr.write('\n'.join(x.decode('utf8') for x in addresses))

read_thread = threading.Thread(target=buffer_file, args=(0,addresses))
read_thread.start()
read_thread.join()
memory_file.seek(0)
window= pyglet.window.Window()
player = pyglet.media.Player()
source = pyglet.media.StreamingSource()
MediaLoad = pyglet.media.load("file.mp4", file=memory_file)

player.queue(MediaLoad)
player.play()


@window.event
def on_draw():
    if player.source and player.source.video_format:
        player.get_texture().blit(50,50)


window.set_size(640, 360)
pyglet.app.run()