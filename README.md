# Word of Wisdom TCP Server

This is a test task for Server Engineer implemented in Go.

## Description

The task involves designing and implementing a TCP server named "Word of Wisdom". The server should be capable of handling Proof of Work (PoW) challenges to protect against DDOS attacks. Upon successful PoW verification, the server will respond with a quote from the "Word of Wisdom" book or any other collection of quotes.

The main algorithm is based on the difficulty of finding the SHA1 hash of a string with a specified difficulty of n zeros at the beginning of the string. Additionally, a time limit is incorporated into the "issue" to control key leakage. The difficulty parameter can be changed in the middleware settings.

Issues
This algorithm is excessively complex; it is impossible to adjust the difficulty in such a way that the number of zeros is greater than 2, otherwise the difficulty becomes excessive.

This problem can be solved by adding a simpler hash check, for example, checking the total number of zeros in the hash rather than just at the beginning. To add controllable difficulty, a sequence of hashes can be introduced, meaning checking the hash of a hash with the addition of a nonce. This parameter can be specified in the middleware settings.

## Features

TCP server implementation in Go.  
Protection against DDOS attacks using Proof of Work.  
Responds with a quote after PoW verification.  
Dockerfiles provided for both server and client.  
Proof of Work (PoW) Algorithm  
The PoW algorithm used in this implementation is based on SHA1 hashing.  

## Docker

The server build is based on dockerfile traefik

```bash
make build

# then like this
docker build -t traefik-gateway ./traefik
```

## Demo

```bash
make demo

# or

docker compouse up --build
```

## Usage

Clone the repository.  
Build the Docker images for server.  
Run the server container.  
Go to: http://quotes-127.0.0.1.nip.io:8081/

## Installation

To run the Word of Wisdom TCP server and client, Docker must be installed on your system.
