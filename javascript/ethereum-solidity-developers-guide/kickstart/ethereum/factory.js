import web3 from './web3';
import CampaignFactory from './build/CampaignFactory.json';

const instance = new web3.eth.Contract(
  JSON.parse(CampaignFactory.interface),
  '0x8cC0a543E32ADeE454c3939cDB83281Af35A3c66'
);

export default instance;