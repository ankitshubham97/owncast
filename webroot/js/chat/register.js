import { URL_CHAT_REGISTRATION } from '../utils/constants.js';

export async function registerChat(username, nonce, signature, walletPublicAddress, nftContractAddress, nftId) {
  const options = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ displayName: username, nonce, signature, walletPublicAddress, nftContractAddress, nftId }),
  };

  try {
    const response = await fetch(URL_CHAT_REGISTRATION, options);
    const result = await response.json();
    return result;
  } catch (e) {
    console.error(e);
  }
}
