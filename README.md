# Shitkey: Simple and Secure File Encryption

Shitkey is a command-line tool for securely encrypting and decrypting files using a password. It is written in Go and uses modern cryptographic primitives to ensure the confidentiality and integrity of your data.

- **Simple Interface**: A straightforward command-line interface for encrypting and decrypting files.
- **Strong Cryptography**: Shitkey uses AES-256 in GCM mode for encryption.
- **Password-Based Key Derivation**: It uses `scrypt` to derive a strong cryptographic key from your password, making it resistant to brute-force attacks.
- **Salt and Nonce Generation**: Each encryption operation uses a unique, cryptographically-secure random salt and nonce to protect against common attacks.
- **Cross-Platform**: The tool is written in Go, which makes it easy to compile and run on various operating systems.

## Installation

### From Source

1. Clone the repository:

    ```sh
    git clone https://github.com/nthnn/shitkey.git
    cd shitkey
    ```

2. Build the executable:

    ```sh
    go build -o shitkey
    ```

3. (Optional) Move the executable to your PATH:

    ```sh
    sudo mv shitkey /usr/local/bin/
    ```

### Debian/Ubuntu

You can build a `.deb` package using the provided `build.sh` script.

1. Clone the repository and navigate to the directory:

    ```sh
    git clone https://github.com/nthnn/shitkey.git
    cd shitkey
    ```

2. Run the build script:

    ```sh
    ./build.sh
    ```

3. Install the generated `.deb` package:

    ```sh
    sudo dpkg -i shitkey_1.4_amd64.deb
    ```

## Usage

### Encrypt a file

To encrypt a file, use the `encrypt` command followed by the filename. The encrypted file will have a `.sk` extension.

```sh
shitkey encrypt mysecretfile.txt
```

### Decrypt a file

To decrypt a file, use the `decrypt` command followed by the filename. The input file must have a `.sk` extension.

```
shitkey decrypt mysecretfile.txt.sk
```

### View version
To display the version of Shitkey, use the version command.

```
shitkey version
```

## How it works

- **Password Input**: The user is prompted to enter a password, which is read securely without being echoed to the terminal.
- **Salt Generation**: A cryptographically-secure random salt is generated for each encryption process. The salt has a size of 32 bytes.
- **Key Derivation**: The scrypt algorithm is used to derive a strong encryption key from the user's password and the generated salt. The key size is 32 bytes.
- **Encryption**: The derived key is used to create an AES-256 cipher block. An AEAD (Authenticated Encryption with Associated Data) scheme, GCM (Galois/Counter Mode), is used for encryption to ensure both confidentiality and integrity. A unique nonce (12 bytes in size) is also generated for the encryption process.
- **File Output**: The encrypted data is written to a new file with the .sk extension. This new file contains the salt, the nonce, and the ciphertext in that specific order.

## License

Shitkey is free software licensed under the GNU General Public License v3.0 or any later version. You should have received a copy of the GNU General Public License with this program. If not, see https://www.gnu.org/licenses/.
