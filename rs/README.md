# Vsock Communication Sample

A hello-world example for Nitro Enclaves vsock server and client communication.

## Prerequisites

1. The `vsock-sample` is a Nitro Enclaves application that can be run either as a server
or a client. The client sends the “Hello, world!” message to the server. The server receives
the message and prints it to the standard output.

__Note__: You can change the behavior by replacing the code marked with *TODO*. 

2. For convenience, a `Makefile` is provided in order to automate the building of the required
enclave image files. The steps **3** - **4** below may be run either manually or through the
`make build` command.

3. Because the code is written in Rust, you need to install `cargo`
(Rust’s build system and package manager) and add the `$ARCH-unknown-linux-musl`
target. You can easily do that using `rustup`, a command line tool
for managing Rust versions and associated tools.

__Note__: The `$ARCH-unknown-linux-musl` target is needed in order to
generate a fully static binary. Currently, the supported values for `$ARCH` (the parent instance's
architecture) are `x86_64` or `aarch64`.

```
ARCH=$(uname -m)
curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh -s -- -y
rustup target install $ARCH-unknown-linux-musl
```

4. Compile the code using `cargo`. You can find the resulting binary
at `vsock-sample/target/$ARCH-unknown-linux-musl/release/vsock-sample`.

```
cargo build --target=$ARCH-unknown-linux-musl --release
cp target/$ARCH-unknown-linux-musl/release/vsock-sample .
```
## How to use the enclave as server and the parent instance as client

1. Build the Enclave Image File (EIF) using the provided Makefile by running `make server`.
Alternately, use the provided `Dockerfile.server`. We chose to start from the alpine distro
to keep the enclave image as small as possible, but you can also use other distros (e.g. Ubuntu).
Then copy the `vsock-sample` binary and use it to start the server inside the enclave.

```
FROM alpine:latest
COPY vsock-sample .
CMD ["./vsock-sample", "server", "--port", "5005"]
```
__Note__: You can use other port numbers as well.

```
docker build -t vsock-sample-server -f Dockerfile.server .
nitro-cli build-enclave --docker-uri vsock-sample-server --output-file vsock_sample_server.eif
```

2. Ensure the `nitro_enclaves` driver is inserted. Configure the vCPUs and the memory pool and run the
enclave using the built EIF.

Note: The `nitro-cli-config` script from https://github.com/aws/aws-nitro-enclaves-cli may be used here.

```
// Use 2 vCPUs and 256 MiB memory
nitro-cli-config -t 2 -m 256
nitro-cli run-enclave --eif-path vsock_sample_server.eif --cpu-count 2 --memory 256 --debug-mode
```

3. Connect to the enclave console using `nitro-cli`.

```
nitro-cli console --enclave-id $ENCLAVE_ID
```

4. Run the client in a terminal from the parent instance.

```
./vsock-sample client --cid $ENCLAVE_CID --port 5005
```

5. The enclave's console output should now look like this:

```
[    0.127079] Freeing unused kernel memory: 476K
[    0.127631] nsm: loading out-of-tree module taints kernel.
[    0.128055] nsm: module verification failed: signature and/or required key missing - tainting kernel
[    0.154010] random: vsock-sample: uninitialized urandom read (16 bytes read)
Hello, world!
```

## How to use the parent instance as server and the enclave as client

1. Build the Enclave Image File (EIF) using the provided Makefile by running `make client`.
Alternately, use the provided `Dockerfile.client`. We chose to start from the alpine distro
to keep the enclave image as small as possible , but you can also use other distros (e.g. Ubuntu).
Then copy the `vsock-sample` binary and use it to start the client inside the enclave.

```
FROM alpine:latest
COPY vsock-sample .
CMD ["./vsock-sample", "client", "--cid", "3",  "--port", "5005"]
```

__Notes__:
* 3 is the CID of the parent instance.
* You can use other port numbers as well.

```
docker build -t vsock-sample-client -f Dockerfile.client .
nitro-cli build-enclave --docker-uri vsock-sample-client --output-file vsock_sample_client.eif
```

2. Run the server inside a terminal of the parent instance.

```
./vsock-sample server --port 5005
```

3. Ensure the `nitro_enclaves` driver is inserted. Configure the vCPUs and the memory pool and run the
enclave using the built EIF.

Note: The `nitro-cli-config` script from https://github.com/aws/aws-nitro-enclaves-cli may be used here.

```
// Use 2 vCPUs and 256 MiB memory
nitro-cli-config -t 2 -m 256
nitro-cli run-enclave --eif-path vsock_sample_client.eif --cpu-count 2 --memory 256 --debug-mode
```

4. The server's output should now look like this:
```
Hello, world!
```

Now you can replace the client / server code with your own code.
