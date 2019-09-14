import web3 from './web3';
import CampaignFactory from './build/CampaignFactory.json';

const instance = new web3.eth.Contract(
  JSON.parse(CampaignFactory.interface),
  '0xD921B701C2dA641F81d6D592C399B9087D2edFEf'
);

export default instance;