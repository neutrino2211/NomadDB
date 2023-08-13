from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.primitives.asymmetric import rsa

private_key = rsa.generate_private_key(
    public_exponent=65537,
    key_size=2048
)

private_key_pass = b"your-password"

encrypted_pem_private_key = private_key.private_bytes(
    encoding=serialization.Encoding.PEM,
    format=serialization.PrivateFormat.PKCS8,
    encryption_algorithm=serialization.BestAvailableEncryption(private_key_pass)
)

print(encrypted_pem_private_key.splitlines()[0])
# b'-----BEGIN ENCRYPTED PRIVATE KEY-----'

pem_public_key = private_key.public_key().public_numbers()

print(pem_public_key)