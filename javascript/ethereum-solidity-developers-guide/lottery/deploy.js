const HDWalletProdvider = require('truffle-hdwallet-provider');
const Web3 = require('web3');
const { interface, bytecode } = require('./compile');

const provider = new HDWalletProdvider(
  'head code remember wait grant ill urban spike emerge lecture plunge swap',
  'https://rinkeby.infura.io/v3/a236891c775d4dc9857ede1fca8b15af'
);

const web3 = new Web3(provider);

const deploy = async () => {
  const accounts = await web3.eth.getAccounts();

  console.log('Attempting to deploy from account: ', accounts[0]);

  const result = await new web3.eth.Contract(JSON.parse(interface))
    .deploy({ data: bytecode })
    .send({ gas: '1000000', from: accounts[0] });
  
  console.log('contract interface: ', interface);
  console.log('Contract deployed to: ', result.options.address);
};

deploy();