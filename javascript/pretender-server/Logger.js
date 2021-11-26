let winston = require('winston')
// const appRoot = require('app-root-path');

const tp = () => { new Date().toLocaleString(); };


const options = {
    file: {
      level: 'info',
      filename: `api.log`,
      json: true,
      colorize: true,
      timestamp: tp,
    },
    console: {
      level: 'debug',
      timestamp: tp,
      json: false,
      colorize: true,
    },
};

const logger = winston.createLogger({
    level: 'info',
    format: winston.format.json(),
    transports: [
      new winston.transports.File(options.file),
      new winston.transports.File({
            level: 'error',
            filename: `error.log`,
            name: 'error-file',
            timestamp: tp,
        })
    ],
    exitOnError: false,

})

module.exports = logger;