import { h, createContext } from '/js/web_modules/preact.js';
import htm from '/js/web_modules/htm.js';
import { ethers } from '/js/web_modules/ethers/ethers.min.js';
import { clearLocalStorage } from '../utils/helpers.js';
import {
  KEY_ACCESS_TOKEN,
  KEY_CHAT_DISPLAYED,
  KEY_USERNAME,
  KEY_NONCE,
  KEY_SIGNATURE,
  KEY_WALLET_PUBLIC_ADDRESS,
} from '../utils/constants.js';
const html = htm.bind(h);

const moderatorFlag = html`
  <img src="/img/moderator-nobackground.svg" class="moderator-flag" />
`;

const Context = createContext();

export const ConnectWallet = (props) => {
  const {
    nonce,
    setNonce,
    signature,
    setSignature,
    walletPublicAddress,
    setWalletPublicAddress,
    error,
    setError,
  } = props;


  const signMessage = async ({ setError, message }) => {
    try {
      console.log({ message });
      if (!window.ethereum)
        throw new Error('No crypto wallet found. Please install it.');

      await window.ethereum.send('eth_requestAccounts');
      const provider = new ethers.providers.Web3Provider(window.ethereum);
      const signer = provider.getSigner();
      const signature = await signer.signMessage(message);
      const address = await signer.getAddress();

      return {
        message,
        signature,
        address,
      };
    } catch (err) {
      setError(err.message);
    }
  };

  const connectWallet = async (e) => {
    e.preventDefault();
    setError();
    const sig = await signMessage({
      setError,
      message: nonce,
    });
    if (sig) {
      setSignature(sig.signature);
      setWalletPublicAddress(sig.address);
      setNonce(sig.message);
      console.log('sig', sig);
    } else {
      setError('No signature');
    }
  };

  return html`
  <${Context.Provider} value=${props}>
  <div className="chat-menu p-2 relative shadow-lg">
    <button
      className="btn btn-secondary submit-button focus:ring focus:outline-none w-full bg-indigo-500 p-2 text-white"
      onClick=${() => {clearLocalStorage(KEY_ACCESS_TOKEN); clearLocalStorage(KEY_USERNAME); clearLocalStorage(KEY_NONCE); clearLocalStorage(KEY_SIGNATURE); clearLocalStorage(KEY_WALLET_PUBLIC_ADDRESS); location.reload(); document.cookie = 'Authorization=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';}}
      >
      Reconnect other wallet
    </button>
  </div>
  </${Context.Provider}>`;
};