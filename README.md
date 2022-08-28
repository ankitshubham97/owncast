<br />
<p align="center">
  <a href="https://github.com/ankitshubham97/owncast/edit/unfold2022-1" alt="Fanment">
    <img src="https://i.postimg.cc/sXY4899J/Group-1.png" alt="Logo" width="200">
  </a>
</p>

<br/>

  <p align="center">
    <strong>A decentralized content sharing/streaming platform built for creators so that they could directly monetize their fans!</strong>
    <br />
    <a href="#"><strong>Project doc @ Devfolio »</strong></a>
    <br />
    <a href="https://youtu.be/9-PB62kO-RU">View Demo</a>
    ·
    <a href="https://github.com/owncast/owncast#readme">Owncast README</a>
  </p>
</p>

## Table of Contents

- [About the Project](#about-the-project)
- [Unfold 2022 Specific](#unfold-2022-specific)
- [Getting Started](#getting-started)
- [Technical concepts](#technical-concepts)

<!-- ABOUT THE PROJECT -->

## About The Project

<p align="center">
  <a href="https://i.postimg.cc/sXY4899J/Group-1.png">
    <img src="https://i.postimg.cc/sXY4899J/Group-1.png" width="70%">
  </a>
</p>

Fanment is a decentralized content sharing/streaming platform built for creators so that they could directly monetize their fans!

---

<!-- UNFOLD 2022 SPECIFIC -->

## Unfold 2022 Specific

Demo video: https://youtu.be/9-PB62kO-RU

This is a submission aimed for `Best NFT project on Polygon` Bounty track.

If the judges need any help to test this out live, I have written an elaborate steps on how to get this up and running at [Getting Started](#getting-started). If they need any help, I will be present at the Unfold 2022 venue and reachable at Telegram at `@ANKITSHUBHAM97`; feel free to ping me! (Since it involved interaction with Streamlabs which is hard to deploy as SAAS, I chose to focus on writing elaborate steps on how to get this up locally.)

<!-- GETTING STARTED -->

## Getting Started

Fanment is written on top of an open source and self-hosted live video streaming and chat server called <a href="https://github.com/owncast/owncast"> Owncast </a>. It is basically a NFT-gating layer on top of Owncast. There are primarily 2 modifications done on the top of Owncast:

1. Adding the capability to sign a nonce via a wallet.
2. Addign an auth-server which validates if the signature is valid and if the wallet signing the message contains the correct NFT. If yes, then generate an access token which the user can use to access the protected resource.

### Getting the services up

1. Clone the repo. `git clone git@github.com:ankitshubham97/owncast.git`
2. Checkout the `unfold2022-1` branch. `git checkout unfold2022-1`
3. `go run main.go` will run from source. It will get the Fanment core backend up and deploy a frontend at http://localhost:8080/ 
4. We have to get the Fanment auth-server up now. Open a new terminal, cd into `auth-server`.
5. If it is the first time, you will need to install the node packages, build and start the server. `npm i && npm run build && npm run start`
6. Auth server should be up and running at http://localhost:3000

### Download an OBS software to broadcast content to Fanment

Let's download <a href="https://streamlabs.com/">Streamlabs</a> now. Once downloaded and the app is running, we will need to point it to our Fanment instance. For this, let's follow how to set up rmtp url from <a href="https://youtu.be/9-PB62kO-RU?t=94"> this point of time in the demo</a>.

1. Open Streamlabs.
2. Go to `Settings` (usually ⚙️ emoji in the bottom left corner)
3. Go to `Stream`.
4. For `Stream Type`, choose `Custom Streaming Server`.
5. For `URL`, choose `rtmp://localhost:1935/live`.
6. For `Stream Key`, input `abc123`.
Here is a <a href="https://youtu.be/9-PB62kO-RU?t=110">reference screen</a> of how the settings would look like.

Put up any video sources that you want to broadcast (In the demo, I am just broadcasting a banner of `Unfold 2022`.) After you are done, hit `Go Live` button (usually in the bottom right)

### Interacting with stream via Fanment

Go to http://localhost:8080/. You should be seeing a screen <a href="https://youtu.be/9-PB62kO-RU?t=27">like this</a>.

You will need a wallet with the correct NFT in it to test the happy flow. Here is a test metamask wallet to test the happy flow:

1. Wallet address: 0x4ad53d31Cb104Cf5f7622f1AF8Ed09C3ca980523
1. Private key: dec5213b700bc944b06584aaf3d508f88a1ce0221b77067b7e7b95d7b88d2ae3
1. Seed phrase: loop lobster mechanic grit lecture video they expose marble photo now family

For the unhappy flow, you could just use any metamask wallet!

## Technical concepts

The technical problem is to validate if a wallet holds an NFT.

We use the concept of cryptography. We need a nonce to be signed by the private key of the wallet. This would generate a signature. We can then verify this signature using the nince and the public key of the wallet(i.e. wallet address).

The nonce is randomly generated. To sign this nonce, we use ethers.js library to open the metamask wallet and let the wallet sign this nonce. We send the nonce, signature and the public address to the auth server.

Auth server uses plain cryptography to verify if the signature was indeed created by the wallet signing the nonce. If verification is successful, we proceed to generate an access token. The access token is implemented as JWT. This access token will be passed back to the user so that he could piggy-back this with every request to fetch the content. I have implemented this as cookie; the access token gets passed with the request as cookie.

Now, when a user sends a request to access the content with access-token piggybacked as cookie, Fanment uses auth server to verify the validity of the access-token. If it is valid then only responds with the content payload; otherwise it sends error.
