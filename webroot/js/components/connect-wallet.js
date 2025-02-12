import { h, createContext } from '/js/web_modules/preact.js';
import htm from '/js/web_modules/htm.js';
import { ethers } from '/js/web_modules/ethers/ethers.min.js';

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
      className="btn btn-primary submit-button focus:ring focus:outline-none w-full"
      onClick=${connectWallet}
      >
      Sign message
    </button>
  </div>
  </${Context.Provider}>`;
};

