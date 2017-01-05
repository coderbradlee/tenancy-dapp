var Web3 = require("web3");
var SolidityEvent = require("web3/lib/web3/event.js");

(function() {
  // Planned for future features, logging, etc.
  function Provider(provider) {
    this.provider = provider;
  }

  Provider.prototype.send = function() {
    this.provider.send.apply(this.provider, arguments);
  };

  Provider.prototype.sendAsync = function() {
    this.provider.sendAsync.apply(this.provider, arguments);
  };

  var BigNumber = (new Web3()).toBigNumber(0).constructor;

  var Utils = {
    is_object: function(val) {
      return typeof val == "object" && !Array.isArray(val);
    },
    is_big_number: function(val) {
      if (typeof val != "object") return false;

      // Instanceof won't work because we have multiple versions of Web3.
      try {
        new BigNumber(val);
        return true;
      } catch (e) {
        return false;
      }
    },
    merge: function() {
      var merged = {};
      var args = Array.prototype.slice.call(arguments);

      for (var i = 0; i < args.length; i++) {
        var object = args[i];
        var keys = Object.keys(object);
        for (var j = 0; j < keys.length; j++) {
          var key = keys[j];
          var value = object[key];
          merged[key] = value;
        }
      }

      return merged;
    },
    promisifyFunction: function(fn, C) {
      var self = this;
      return function() {
        var instance = this;

        var args = Array.prototype.slice.call(arguments);
        var tx_params = {};
        var last_arg = args[args.length - 1];

        // It's only tx_params if it's an object and not a BigNumber.
        if (Utils.is_object(last_arg) && !Utils.is_big_number(last_arg)) {
          tx_params = args.pop();
        }

        tx_params = Utils.merge(C.class_defaults, tx_params);

        return new Promise(function(accept, reject) {
          var callback = function(error, result) {
            if (error != null) {
              reject(error);
            } else {
              accept(result);
            }
          };
          args.push(tx_params, callback);
          fn.apply(instance.contract, args);
        });
      };
    },
    synchronizeFunction: function(fn, instance, C) {
      var self = this;
      return function() {
        var args = Array.prototype.slice.call(arguments);
        var tx_params = {};
        var last_arg = args[args.length - 1];

        // It's only tx_params if it's an object and not a BigNumber.
        if (Utils.is_object(last_arg) && !Utils.is_big_number(last_arg)) {
          tx_params = args.pop();
        }

        tx_params = Utils.merge(C.class_defaults, tx_params);

        return new Promise(function(accept, reject) {

          var decodeLogs = function(logs) {
            return logs.map(function(log) {
              var logABI = C.events[log.topics[0]];

              if (logABI == null) {
                return null;
              }

              var decoder = new SolidityEvent(null, logABI, instance.address);
              return decoder.decode(log);
            }).filter(function(log) {
              return log != null;
            });
          };

          var callback = function(error, tx) {
            if (error != null) {
              reject(error);
              return;
            }

            var timeout = C.synchronization_timeout || 240000;
            var start = new Date().getTime();

            var make_attempt = function() {
              C.web3.eth.getTransactionReceipt(tx, function(err, receipt) {
                if (err) return reject(err);

                if (receipt != null) {
                  // If they've opted into next gen, return more information.
                  if (C.next_gen == true) {
                    return accept({
                      tx: tx,
                      receipt: receipt,
                      logs: decodeLogs(receipt.logs)
                    });
                  } else {
                    return accept(tx);
                  }
                }

                if (timeout > 0 && new Date().getTime() - start > timeout) {
                  return reject(new Error("Transaction " + tx + " wasn't processed in " + (timeout / 1000) + " seconds!"));
                }

                setTimeout(make_attempt, 1000);
              });
            };

            make_attempt();
          };

          args.push(tx_params, callback);
          fn.apply(self, args);
        });
      };
    }
  };

  function instantiate(instance, contract) {
    instance.contract = contract;
    var constructor = instance.constructor;

    // Provision our functions.
    for (var i = 0; i < instance.abi.length; i++) {
      var item = instance.abi[i];
      if (item.type == "function") {
        if (item.constant == true) {
          instance[item.name] = Utils.promisifyFunction(contract[item.name], constructor);
        } else {
          instance[item.name] = Utils.synchronizeFunction(contract[item.name], instance, constructor);
        }

        instance[item.name].call = Utils.promisifyFunction(contract[item.name].call, constructor);
        instance[item.name].sendTransaction = Utils.promisifyFunction(contract[item.name].sendTransaction, constructor);
        instance[item.name].request = contract[item.name].request;
        instance[item.name].estimateGas = Utils.promisifyFunction(contract[item.name].estimateGas, constructor);
      }

      if (item.type == "event") {
        instance[item.name] = contract[item.name];
      }
    }

    instance.allEvents = contract.allEvents;
    instance.address = contract.address;
    instance.transactionHash = contract.transactionHash;
  };

  // Use inheritance to create a clone of this contract,
  // and copy over contract's static functions.
  function mutate(fn) {
    var temp = function Clone() { return fn.apply(this, arguments); };

    Object.keys(fn).forEach(function(key) {
      temp[key] = fn[key];
    });

    temp.prototype = Object.create(fn.prototype);
    bootstrap(temp);
    return temp;
  };

  function bootstrap(fn) {
    fn.web3 = new Web3();
    fn.class_defaults  = fn.prototype.defaults || {};

    // Set the network iniitally to make default data available and re-use code.
    // Then remove the saved network id so the network will be auto-detected on first use.
    fn.setNetwork("default");
    fn.network_id = null;
    return fn;
  };

  // Accepts a contract object created with web3.eth.contract.
  // Optionally, if called without `new`, accepts a network_id and will
  // create a new version of the contract abstraction with that network_id set.
  function Contract() {
    if (this instanceof Contract) {
      instantiate(this, arguments[0]);
    } else {
      var C = mutate(Contract);
      var network_id = arguments.length > 0 ? arguments[0] : "default";
      C.setNetwork(network_id);
      return C;
    }
  };

  Contract.currentProvider = null;

  Contract.setProvider = function(provider) {
    var wrapped = new Provider(provider);
    this.web3.setProvider(wrapped);
    this.currentProvider = provider;
  };

  Contract.new = function() {
    if (this.currentProvider == null) {
      throw new Error("Tenancy error: Please call setProvider() first before calling new().");
    }

    var args = Array.prototype.slice.call(arguments);

    if (!this.unlinked_binary) {
      throw new Error("Tenancy error: contract binary not set. Can't deploy new instance.");
    }

    var regex = /__[^_]+_+/g;
    var unlinked_libraries = this.binary.match(regex);

    if (unlinked_libraries != null) {
      unlinked_libraries = unlinked_libraries.map(function(name) {
        // Remove underscores
        return name.replace(/_/g, "");
      }).sort().filter(function(name, index, arr) {
        // Remove duplicates
        if (index + 1 >= arr.length) {
          return true;
        }

        return name != arr[index + 1];
      }).join(", ");

      throw new Error("Tenancy contains unresolved libraries. You must deploy and link the following libraries before you can deploy a new version of Tenancy: " + unlinked_libraries);
    }

    var self = this;

    return new Promise(function(accept, reject) {
      var contract_class = self.web3.eth.contract(self.abi);
      var tx_params = {};
      var last_arg = args[args.length - 1];

      // It's only tx_params if it's an object and not a BigNumber.
      if (Utils.is_object(last_arg) && !Utils.is_big_number(last_arg)) {
        tx_params = args.pop();
      }

      tx_params = Utils.merge(self.class_defaults, tx_params);

      if (tx_params.data == null) {
        tx_params.data = self.binary;
      }

      // web3 0.9.0 and above calls new twice this callback twice.
      // Why, I have no idea...
      var intermediary = function(err, web3_instance) {
        if (err != null) {
          reject(err);
          return;
        }

        if (err == null && web3_instance != null && web3_instance.address != null) {
          accept(new self(web3_instance));
        }
      };

      args.push(tx_params, intermediary);
      contract_class.new.apply(contract_class, args);
    });
  };

  Contract.at = function(address) {
    if (address == null || typeof address != "string" || address.length != 42) {
      throw new Error("Invalid address passed to Tenancy.at(): " + address);
    }

    var contract_class = this.web3.eth.contract(this.abi);
    var contract = contract_class.at(address);

    return new this(contract);
  };

  Contract.deployed = function() {
    if (!this.address) {
      throw new Error("Cannot find deployed address: Tenancy not deployed or address not set.");
    }

    return this.at(this.address);
  };

  Contract.defaults = function(class_defaults) {
    if (this.class_defaults == null) {
      this.class_defaults = {};
    }

    if (class_defaults == null) {
      class_defaults = {};
    }

    var self = this;
    Object.keys(class_defaults).forEach(function(key) {
      var value = class_defaults[key];
      self.class_defaults[key] = value;
    });

    return this.class_defaults;
  };

  Contract.extend = function() {
    var args = Array.prototype.slice.call(arguments);

    for (var i = 0; i < arguments.length; i++) {
      var object = arguments[i];
      var keys = Object.keys(object);
      for (var j = 0; j < keys.length; j++) {
        var key = keys[j];
        var value = object[key];
        this.prototype[key] = value;
      }
    }
  };

  Contract.all_networks = {
  "default": {
    "abi": [
      {
        "constant": false,
        "inputs": [],
        "name": "terminate",
        "outputs": [],
        "payable": true,
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [],
        "name": "property",
        "outputs": [
          {
            "name": "",
            "type": "address"
          }
        ],
        "payable": false,
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "prospectiveTenant",
            "type": "address"
          }
        ],
        "name": "acceptNegotiationOwner",
        "outputs": [],
        "payable": false,
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [],
        "name": "withdraw",
        "outputs": [],
        "payable": false,
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [],
        "name": "acceptNegotiationTenant",
        "outputs": [],
        "payable": true,
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [],
        "name": "owner",
        "outputs": [
          {
            "name": "",
            "type": "address"
          }
        ],
        "payable": false,
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "prospectiveTenant",
            "type": "address"
          }
        ],
        "name": "rejectNegotiation",
        "outputs": [],
        "payable": false,
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [],
        "name": "tenant",
        "outputs": [
          {
            "name": "",
            "type": "address"
          }
        ],
        "payable": false,
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "rent",
            "type": "uint256"
          },
          {
            "name": "security",
            "type": "uint256"
          },
          {
            "name": "start",
            "type": "uint256"
          },
          {
            "name": "end",
            "type": "uint256"
          }
        ],
        "name": "updateCondition",
        "outputs": [],
        "payable": false,
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [],
        "name": "state",
        "outputs": [
          {
            "name": "",
            "type": "uint8"
          }
        ],
        "payable": false,
        "type": "function"
      },
      {
        "constant": true,
        "inputs": [],
        "name": "condition",
        "outputs": [
          {
            "name": "rent",
            "type": "uint256"
          },
          {
            "name": "security",
            "type": "uint256"
          },
          {
            "name": "startTime",
            "type": "uint256"
          },
          {
            "name": "endTime",
            "type": "uint256"
          }
        ],
        "payable": false,
        "type": "function"
      },
      {
        "constant": false,
        "inputs": [
          {
            "name": "person",
            "type": "address"
          },
          {
            "name": "rent",
            "type": "uint256"
          },
          {
            "name": "security",
            "type": "uint256"
          },
          {
            "name": "start",
            "type": "uint256"
          },
          {
            "name": "end",
            "type": "uint256"
          }
        ],
        "name": "negotiate",
        "outputs": [],
        "payable": false,
        "type": "function"
      },
      {
        "inputs": [
          {
            "name": "registry",
            "type": "address"
          },
          {
            "name": "person",
            "type": "address"
          },
          {
            "name": "property",
            "type": "address"
          },
          {
            "name": "rent",
            "type": "uint256"
          },
          {
            "name": "security",
            "type": "uint256"
          },
          {
            "name": "start",
            "type": "uint256"
          },
          {
            "name": "end",
            "type": "uint256"
          }
        ],
        "payable": false,
        "type": "constructor"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "prospectiveTenant",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "owner",
            "type": "address"
          }
        ],
        "name": "Negotiate",
        "type": "event"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "prospectiveTenant",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "owner",
            "type": "address"
          }
        ],
        "name": "Withdraw",
        "type": "event"
      },
      {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "owner",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "prospectiveTenant",
            "type": "address"
          }
        ],
        "name": "RejectNegotiation",
        "type": "event"
      }
    ],
    "unlinked_binary": "0x6060604052346100005760405160e080610ced83398101604090815281516020830151918301516060840151608085015160a086015160c09096015193959293919290915b86600160a060020a0316633faa78b6868660006000604051602001526040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018084600160a060020a0316600160a060020a031681526020018381526020018281526020019350505050602060405180830381600087803b156100005760325a03f115610000575050604051511590506101405760018054600160a060020a031916600160a060020a038816179055604080516080810182528581526020810185905290810183905260600181905260038490556004839055600582905560068190556007805460ff19169055610145565b610000565b5b505050505050505b610b908061015d6000396000f3006060604052361561009e5763ffffffff60e060020a6000350416630c08bf8881146100a3578063176fd3f0146100ad578063199a620a146100d65780633ccfd60b146100f15780637ed02af9146101005780638da5cb5b1461010a5780639bcc912314610133578063adf077911461014e578063b2e2c1c914610177578063c19d93fb14610192578063c5031331146101c0578063d5c69585146101f3575b610000565b6100ab61021a565b005b34610000576100ba6103cd565b60408051600160a060020a039092168252519081900360200190f35b34610000576100ab600160a060020a03600435166103dc565b005b34610000576100ab610575565b005b6100ab610622565b005b34610000576100ba6107ae565b60408051600160a060020a039092168252519081900360200190f35b34610000576100ab600160a060020a03600435166107bd565b005b34610000576100ba61087b565b60408051600160a060020a039092168252519081900360200190f35b34610000576100ab60043560243560443560643561088a565b005b346100005761019f610957565b6040518082600281116100005760ff16815260200191505060405180910390f35b34610000576101cd610960565b604080519485526020850193909352838301919091526060830152519081900360800190f35b34610000576100ab600160a060020a036004351660243560443560643560843561096f565b005b6001546040805160006020918201819052825160e060020a630e2562d9028152925190938493600160a060020a0390911692630e2562d99260048084019382900301818787803b156100005760325a03f1156100005750505060405180519050600160a060020a031633600160a060020a0316148061030857506002546040805160006020918201819052825160e060020a630e2562d90281529251600160a060020a0390941693630e2562d99360048082019493918390030190829087803b156100005760325a03f1156100005750505060405180519050600160a060020a031633600160a060020a0316145b151561031357610000565b60065442111561009e576002546040805160006020918201819052825160e060020a630e2562d90281529251600160a060020a0390941693630e2562d99360048082019493918390030190829087803b156100005760325a03f1156100005750506040518051600160a060020a0316915083830380156108fc02916000818181858888f1935050505015156103a757610000565b50506007805460ff19166002179055600b54600a8104906103c8565b610000565b5b5050565b600054600160a060020a031681565b6001546040805160006020918201819052825160e060020a630e2562d902815292519093600160a060020a031692630e2562d992600480830193919282900301818787803b156100005760325a03f1156100005750506040515133600160a060020a03908116911614905061045057610000565b600260075460ff166002811161000057141561046b57610000565b6001546040805160006020918201819052825160e060020a630e2562d90281529251600160a060020a0390941693630e2562d99360048082019493918390030190829087803b156100005760325a03f1156100005750506040515133600160a060020a0390811691161490506104e057610000565b50600160a060020a03811660009081526008602052604090206005810154151561050957610000565b805460098054600160a060020a031916600160a060020a039092169190911790556001810154600a556002810154600b556003810154600c556004810154600d556005810154600e556006810154600f805460ff909216151560ff199092169190911790555b5b5b5050565b600160a060020a03331660009081526008602052604081206005810154151561059d57610000565b600260075460ff16600281116100005714156105b857610000565b600160a060020a033381166000818152600860209081526040808320600581019390935560015481519485529094169083015282519094507f34d58c18c6c1df2c698ccac556acea92941ca7b99d2fccf9e3ac1852d0dec36f929181900390910190a15b5b5b5050565b600080600260075460ff166002811161000057141561064057610000565b6009546040805160006020918201819052825160e060020a630e2562d90281529251600160a060020a0390941693630e2562d99360048082019493918390030190829087803b156100005760325a03f1156100005750506040515133600160a060020a0390811691161490506106b557610000565b600b54600a54013410156106c857610000565b5050600a546001546040805160006020918201819052825160e060020a630e2562d90281529251606460028702049594600160a060020a031693630e2562d9936004808301949193928390030190829087803b156100005760325a03f1156100005750506040518051600160a060020a0316915083830380156108fc02916000818181858888f19350505050151561075f57610000565b6007805460ff1916600117905560095460028054600160a060020a03909216600160a060020a0319909216919091179055600a54600355600b54600455600c54600555600d546006555b5b5050565b600154600160a060020a031681565b6001546040805160006020918201819052825160e060020a630e2562d902815292519093600160a060020a031692630e2562d992600480830193919282900301818787803b156100005760325a03f1156100005750506040515133600160a060020a03908116911614905061083157610000565b600260075460ff166002811161000057141561084c57610000565b50600160a060020a038116600090815260086020526040902060068101805460ff191660011790555b5b5b5050565b600254600160a060020a031681565b6001546040805160006020918201819052825160e060020a630e2562d90281529251600160a060020a0390941693630e2562d99360048082019493918390030190829087803b156100005760325a03f1156100005750506040515133600160a060020a0390811691161490506108ff57610000565b600260075460ff166002811161000057141561091a57610000565b604080516080810182528581526020810185905290810183905260600181905260038490556004839055600582905560068190555b5b5b50505050565b60075460ff1681565b60035460045460055460065484565b6000600260075460ff166002811161000057141561098c57610000565b33600160a060020a031686600160a060020a0316630e2562d96000604051602001526040518163ffffffff1660e060020a028152600401809050602060405180830381600087803b156100005760325a03f11561000057505060405151600160a060020a0316919091149050610a0157610000565b600160075460ff1660028111610000571415610a8d576002546040805160006020918201819052825160e060020a630e2562d90281529251600160a060020a03338116951693630e2562d9936004808301949193928390030190829087803b156100005760325a03f11561000057505060405151600160a060020a0316919091149050610a8d57610000565b5b50600160a060020a033381166000818152600860208181526040808420805482516080810184528c81528085018c90528084018b905260600189905260018083018d9055600283018c9055600383018b9055600483018a905560068301805460ff1916905560058301819055958790529383528b8716600160a060020a031994851617938416938716939093178355925483519485529094169383019390935280517f5bfa20921b9a30cfbdde3c55ae53d357aa9fa3ec5ddb8e0700767a8e604ad8949281900390910190a15b5b5050505050505600a165627a7a72305820fab4dc89a5712888217e30bd6fa36562e050e8673a663060d380720e397299080029",
    "events": {
      "0xd0d0546c1578257629800deb0b45ef2b35d76057132b69ac2be057062fc8e3b8": {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "prospectiveTenant",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "owner",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "amount",
            "type": "uint256"
          }
        ],
        "name": "Negotiate",
        "type": "event"
      },
      "0x20760a4e390a305962dd840b66c77f19509f4aaec0ae5ee1e2e0b9ab892fd04e": {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "owner",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "prospectiveTenant",
            "type": "address"
          }
        ],
        "name": "RejectNegotiation",
        "type": "event"
      },
      "0x5bfa20921b9a30cfbdde3c55ae53d357aa9fa3ec5ddb8e0700767a8e604ad894": {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "prospectiveTenant",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "owner",
            "type": "address"
          }
        ],
        "name": "Negotiate",
        "type": "event"
      },
      "0x34d58c18c6c1df2c698ccac556acea92941ca7b99d2fccf9e3ac1852d0dec36f": {
        "anonymous": false,
        "inputs": [
          {
            "indexed": false,
            "name": "prospectiveTenant",
            "type": "address"
          },
          {
            "indexed": false,
            "name": "owner",
            "type": "address"
          }
        ],
        "name": "Withdraw",
        "type": "event"
      }
    },
    "updated_at": 1483576443915
  }
};

  Contract.checkNetwork = function(callback) {
    var self = this;

    if (this.network_id != null) {
      return callback();
    }

    this.web3.version.network(function(err, result) {
      if (err) return callback(err);

      var network_id = result.toString();

      // If we have the main network,
      if (network_id == "1") {
        var possible_ids = ["1", "live", "default"];

        for (var i = 0; i < possible_ids.length; i++) {
          var id = possible_ids[i];
          if (Contract.all_networks[id] != null) {
            network_id = id;
            break;
          }
        }
      }

      if (self.all_networks[network_id] == null) {
        return callback(new Error(self.name + " error: Can't find artifacts for network id '" + network_id + "'"));
      }

      self.setNetwork(network_id);
      callback();
    })
  };

  Contract.setNetwork = function(network_id) {
    var network = this.all_networks[network_id] || {};

    this.abi             = this.prototype.abi             = network.abi;
    this.unlinked_binary = this.prototype.unlinked_binary = network.unlinked_binary;
    this.address         = this.prototype.address         = network.address;
    this.updated_at      = this.prototype.updated_at      = network.updated_at;
    this.links           = this.prototype.links           = network.links || {};
    this.events          = this.prototype.events          = network.events || {};

    this.network_id = network_id;
  };

  Contract.networks = function() {
    return Object.keys(this.all_networks);
  };

  Contract.link = function(name, address) {
    if (typeof name == "function") {
      var contract = name;

      if (contract.address == null) {
        throw new Error("Cannot link contract without an address.");
      }

      Contract.link(contract.contract_name, contract.address);

      // Merge events so this contract knows about library's events
      Object.keys(contract.events).forEach(function(topic) {
        Contract.events[topic] = contract.events[topic];
      });

      return;
    }

    if (typeof name == "object") {
      var obj = name;
      Object.keys(obj).forEach(function(name) {
        var a = obj[name];
        Contract.link(name, a);
      });
      return;
    }

    Contract.links[name] = address;
  };

  Contract.contract_name   = Contract.prototype.contract_name   = "Tenancy";
  Contract.generated_with  = Contract.prototype.generated_with  = "3.2.0";

  // Allow people to opt-in to breaking changes now.
  Contract.next_gen = false;

  var properties = {
    binary: function() {
      var binary = Contract.unlinked_binary;

      Object.keys(Contract.links).forEach(function(library_name) {
        var library_address = Contract.links[library_name];
        var regex = new RegExp("__" + library_name + "_*", "g");

        binary = binary.replace(regex, library_address.replace("0x", ""));
      });

      return binary;
    }
  };

  Object.keys(properties).forEach(function(key) {
    var getter = properties[key];

    var definition = {};
    definition.enumerable = true;
    definition.configurable = false;
    definition.get = getter;

    Object.defineProperty(Contract, key, definition);
    Object.defineProperty(Contract.prototype, key, definition);
  });

  bootstrap(Contract);

  if (typeof module != "undefined" && typeof module.exports != "undefined") {
    module.exports = Contract;
  } else {
    // There will only be one version of this contract in the browser,
    // and we can use that.
    window.Tenancy = Contract;
  }
})();
