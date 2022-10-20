

'use strict';


if(process.env.DISABLE_PROFILER) {
  console.log("Profiler disabled.")
}
else {
  console.log("Profiler enabled.")
  require('@google-cloud/profiler').start({
    serviceContext: {
      service: 'paymentservice',
      version: '1.0.0'
    }
  });
}


if(process.env.DISABLE_TRACING) {
  console.log("Tracing disabled.")
}
else {
  console.log("Tracing enabled.")
  require('@google-cloud/trace-agent').start();

}

if(process.env.DISABLE_DEBUGGER) {
  console.log("Debugger disabled.")
}
else {
  console.log("Debugger enabled.")
  require('@google-cloud/debug-agent').start({
    serviceContext: {
      service: 'paymentservice',
      version: 'VERSION'
    }
  });
}


const path = require('path');
const HipsterShopServer = require('./server');

const PORT = process.env['PORT'];
const PROTO_PATH = path.join(__dirname, '/proto/');

const server = new HipsterShopServer(PROTO_PATH, PORT);

server.listen();
