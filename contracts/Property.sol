pragma solidity ^0.4.4;

contract Property {
    address constant government = 0x429d61dc95cac25a24feffcf7db98f76d6ab3796;
    bool valid = false;

    string latitude;
    string longitude;

    address owner;
    uint rent;
    uint security;

    address tenant;
    uint startTime;
    uint endTime;

    event Registered(address owner, address government);
    event Validated(address government, address owner);

    function Property(string _latitude, string _longitude, uint _rent, uint _security) {
        owner = msg.sender;
        latitude = _latitude;
        longitude = _longitude;
        rent = _rent;
        security = _security;

        Registered(owner, government);
    }

    modifier onlyGovernment() {
    	if (msg.sender != government) throw;
    	_;
    }

    function validate() onlyGovernment {
        if (msg.sender != government) throw;
        valid = true;
        Validated(government, owner);
    }

    struct EjariRule {
        uint incrementPercentage;
		uint maxRent;
    }

    EjariRule ejariRule;

    function setupEjariRule(uint _incrementPercentage, uint _maxRent) onlyGovernment {
        ejariRule.incrementPercentage = _incrementPercentage;
        ejariRule.maxRent = _maxRent;
    }

    struct Offer {
        address tenant;
        uint startTime;
        uint endTime;
    }

    Offer tenantOffer;
    bool acceptedOffer = false;

    event Interested(address tenant, address owner);

    function tenantInterested(uint startTime, uint endTime) {
        if (acceptedOffer) throw;

        tenantOffer = Offer(msg.sender, startTime, endTime);
        Interested(msg.sender, owner);
    }

    modifier onlyOwner() {
    	if (msg.sender != owner) throw;
    	_;
    }

    event Accepted(address owner, address tenant);
    function acceptOffer() onlyOwner {
        acceptedOffer = true;
        Accepted(owner, tenantOffer.tenant);
    }

    modifier onlyAcceptedTenant() {
    	if (acceptedOffer && tenantOffer.tenant != msg.sender) throw;
    	_;
    }

    event Payment(address tenant, address owner);
    function pay() payable onlyAcceptedTenant {
        if (msg.value < rent + security) throw;

        if (!owner.send(rent)) throw;

        tenant = tenantOffer.tenant;
        startTime = tenantOffer.startTime;
        endTime = tenantOffer.endTime;

        Payment(tenant, owner);
    }


    // EXTENTION OF TENANCY
    // TO BE DONE

}
