module.exports = {
  build: {
    "index.html": "index.html",
    "app.js": [
      "javascripts/app.js"
    ],
    "app.css": [
      "stylesheets/app.css"
    ],
    "images/": "images/"
  },
  rpc: {
    host: "localhost",
    port: 8545
  },
  network: {
    "morden": {
	  network_id: 2
    },
   	"development": {
	  network_id: "default"
    },
    "gethdev": {
      network_id: 12345,
	  host: "localhost",
	  port: 8545,
	  from: "0x429d61dc95cac25a24feffcf7db98f76d6ab3796"
    }
  }
};
