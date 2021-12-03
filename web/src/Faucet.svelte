<script>
  import { onMount } from 'svelte';
  import { getAddress } from '@ethersproject/address';
  import { formatEther } from '@ethersproject/units';
  import { setDefaults as setToast, toast } from 'bulma-toast';

  $: address = null
  $: network = null
  $: disabled = false
  $: faucetInfo = {
    account: '0x0000000000000000000000000000000000000000',
    network: 'Testnet',
    payout: 1,
    rpcAvailable: true,
  }
  const testnetConfig = {
    chainId: '0x3ac',
    chainName: 'PulseChain Testnet',
    nativeCurrency: {
      name: 'Pulse',
      symbol: 'PLS',
      decimals: 18,
    },
    rpcUrls: [
      'https://rpc.testnet.pulsechain.com/',
    ],
    blockExplorerUrls: [
      'https://scan.pulsechain.com/',
    ],
    iconUrls: [],
  }
  const updateChainId = async () => {
    if (!ethereum) {
      return
    }
    const chainId = await ethereum.request({
      method: 'eth_chainId',
    })
    network = chainId
  }
  const requestAccounts = async () => {
    if (!ethereum) {
      return
    }
    const accounts = await ethereum.request({
      method: 'eth_requestAccounts',
    })
    address = accounts[0]
  }
  if (window.ethereum) {
    network = ethereum.chainId
    ethereum.once('connect', requestAccounts)
    ethereum.on('connect', updateChainId)
    ethereum.on('chainChanged', updateChainId)
    ethereum.on('accountsChanged', (accounts) => {
      address = accounts[0]
    })
  }

  onMount(async () => {
    const res = await fetch('/api/info');
    try {
      faucetInfo = await res.json();
    } catch (err) {}
    if (!faucetInfo.account || faucetInfo.account === '0x0000000000000000000000000000000000000000') {
      network = testnetConfig.chainId
      disabled = true
      return
    }
    updateChainId()
    requestAccounts()
    faucetInfo.network = capitalize(faucetInfo.network);
    faucetInfo.payout = parseInt(formatEther(faucetInfo.payout));
  });

  setToast({
    duration: 10000,
    position: 'bottom-center',
    dismissible: true,
    pauseOnHover: true,
    closeOnClick: false,
    animate: { in: 'fadeIn', out: 'fadeOut' },
  });

  async function addTestnetToMetamask () {
    try {
      await ethereum.request({
        method: 'wallet_switchEthereumChain',
        params: [{ chainId: testnetConfig.chainId }],
      })
    } catch (switchError) {
      if (switchError.code === 4902) {
        try {
          await ethereum.request({
            method: 'wallet_addEthereumChain',
            params: [testnetConfig],
          })
        } catch (addError) {
          console.log('unable to add a new chain', addError)
          return
        }
      }
      console.log('unable to complete request', switchError)
    }
  }

  async function handleRequest() {
    try {
      address = getAddress(address);
    } catch (error) {
      toast({ message: error.reason, type: 'is-warning' });
      return;
    }

    let formData = new FormData();
    formData.append('address', address);
    const res = await fetch('/api/claim', {
      method: 'POST',
      body: formData,
    })
    const { ok } = res
    let type = ok ? 'is-success' : 'is-warning'
    let message = await res.text()
    toast({ message, type });
  }

  function capitalize(str) {
    const lower = str.toLowerCase();
    return str.charAt(0).toUpperCase() + lower.slice(1);
  }
</script>

<svelte:head>
  <title>tPLS {faucetInfo.network} Faucet</title>
</svelte:head>

<main>
  <section class="hero is-info is-fullheight">
    <div class="hero-head">
      <nav class="navbar">
        <div class="container">
          <div class="navbar-brand">
            <a class="navbar-item" href=".">
              <span class="icon">
                <i class="fa fa-bath" />
              </span>
              <span><b>tPLS Faucet</b></span>
            </a>
          </div>
          <div id="navbarMenu" class="navbar-menu">
            <div class="navbar-end">
              <span class="navbar-item">
                <a
                  class="button is-white is-outlined"
                  href="https://gitlab.com/pulsechaincom/pls-faucet"
                >
                  <span class="icon">
                    <i class="fa fa-github" />
                  </span>
                  <span>View Source</span>
                </a>
              </span>
            </div>
          </div>
        </div>
      </nav>
    </div>

    <div class="hero-body">
      <div class="container has-text-centered">
        <div class="column is-6 is-offset-3">
          <h1 class="title">
            Receive {faucetInfo.payout} tPLS per request
          </h1>
          <h2 class="subtitle">
            Serving from {faucetInfo.account}
          </h2>
          {#if disabled}
            <div class="box">
              <p>unable to contact network. please try again later</p>
            </div>
          {:else if network === testnetConfig.chainId}
          <div class="box">
            <div class="field is-grouped">
              <p class="control is-expanded">
                <input
                  bind:value={address}
                  class="input is-rounded"
                  type="text"
                  placeholder="Enter your address"
                />
              </p>
              <p class="control">
                <button
                  on:click={handleRequest}
                  class="button is-primary is-rounded"
                >
                  Request
                </button>
              </p>
            </div>
          </div>
          {:else}
          <button
            on:click={addTestnetToMetamask}
            class="button is-primary is-rounded">Switch to / Add PulseChain Testnet to Metamask</button>
          {/if}
        </div>
      </div>
    </div>
  </section>
</main>

<style>
  .hero.is-info {
    background: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
      url('/saturn-hexagon-vortex.jpeg') no-repeat center center fixed;
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
  }
  .hero .subtitle {
    padding: 3rem 0;
    line-height: 1.5;
  }
  .box {
    border-radius: 19px;
  }
</style>
