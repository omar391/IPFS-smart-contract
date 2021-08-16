# IPFS smart contract demo

### This repo demonstrate how to upload a File using IPFS embedded go node and storing information using smart contracts

# Features:
⭐ Based on "Standard Go Project Layout".
<br>⭐ Use embedded IPFS node, so its easier to deploy as standalone service.
<br>⭐ Simple way to install solc and abigen via make
<br>⭐ It upload a file into IPFS node and then store file CID as smart contract into the blockchain

# Steps:
1. clone the repo using: git clone
2. exec: make
3. setup local Ganache cli or UI
4. update the etc/dev.env file
5. exec: go run cmd/main.go
	
## APP output:
```
1.-- Running an embedded IPFS node on a tmp directory-- 
-- IPFS node is running successfully --

2.-- Adding test.pdf from the etc dir --
Added file to IPFS with CID: /ipfs/QmVDf2317mT4pPwg85G9XAvaRZLdFWNxoBqCQpw6bEmj4a

3.-- Deploying our CID storage smart contract into Ganache --
Deployed address: 0x75930c9E9492e76832cd163723c7ee3dd324427C --

4.-- Storing our CID into the blockchain --

5.-- Retrieving our CID from the blockchain --
We have successfully stored and retrieved the CID: /ipfs/QmVDf2317mT4pPwg85G9XAvaRZLdFWNxoBqCQpw6bEmj4a
```

## `Ganache` interface:
![image](https://user-images.githubusercontent.com/13569609/129557878-08645539-c7a7-4419-91a8-51b502498ded.png)

<br>
