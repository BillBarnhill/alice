# TSS example

This example demonstrates a simple p2p application using our TSS library. Let's assume we have 3 nodes where their ranks are all 0. These 3 nodes will interact with each other by using `go-libp2p` library. After each process (DKG, signer, and reshare), the results will be written in files located in `config/`.

## Build
```sh
> make tss-example
```

## Usage
### DKG

First, we run 3 hosts on different terminals. These 3 nodes will try to connect to each other. Once it connects to a peer, it will send the peer message out. After the peer messages are fully transmitted, each node will try to get the result and write it to the respective config file.

On node A, 
```sh
> ./example -config config/id-1.yaml
```

On node B,
```sh
> ./example -config config/id-2.yaml
```

On node C,
```sh
> ./example -config config/id-3.yaml
```