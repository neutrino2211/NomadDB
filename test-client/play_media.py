import pyglet
import io
import struct
import socket
import threading

HOST = "127.0.0.1"  # The server's hostname or IP address
PORT = 16475  # The port used by the server

operations = {
    "add": 0,
    "fetch": 1,
    "update": 2,
    "delete": 3,

    "end": 255
}


addresses = []

with open("addresses") as a:
    addresses = a.readlines()

def read_chunk(conn, address: str):
    address = address.strip()
    address = struct.pack("s2048s", bytes([operations["fetch"]]), bytes(address, "ISO-8859-1"))
    conn.sendall(address)
    data = conn.recv(2048)
    return data

def buffer_file(_, addresses):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect((HOST, PORT))
        i = 0
        for address in addresses:
            data = read_chunk(s, address)
            i += 1
            memory_file.write(data)
            print("wrote", len(data), i)

        s.close()
        print("DONE")

print(len(addresses))
memory_file = io.BytesIO()


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