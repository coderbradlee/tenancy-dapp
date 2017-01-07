pragma solidity ^0.4.4;

contract EjariRules {
    address owner;

    struct Rule {
        uint incrementPercentage;
        uint maxRent;
    }

    mapping (bytes32 => Rule) rules;

    function EjariRules() {
        owner = msg.sender;
    }

    function addEjariRule(string latitude, string longitude, uint incrementPercentage, uint maxRent) {
        if (msg.sender != owner) {
            throw;
        }

        rules[sha256(latitude, longitude)] = Rule(incrementPercentage, maxRent);
    }

    function isValid(string latitude, string longitude, uint oldRent, uint newRent) public returns (bool) {
        Rule rule = rules[sha256(latitude, longitude)];

        uint maxIncrementedRent = (oldRent * (100 + rule.incrementPercentage)) / 100;

        if (newRent > rule.maxRent) return false;
        if (newRent > maxIncrementedRent) return false;

        return true;
    }
}

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

    mapping(address => Rating) tenantRatings;

    struct Rating {
        uint totalRatings;
        uint numberOfRatings;
    }

    Rating public ownerRating;
    Rating public propertyRating;

    function rateTenant(uint rating) onlyOwner {
        Rating tenantRating = tenantRatings[tenant];
        tenantRating.totalRatings += rating;
        tenantRating.numberOfRatings++;
    }

    function rateOwner(uint rating) onlyTenant {
            ownerRating.totalRatings += rating;
            ownerRating.numberOfRatings++;
    }

    function rateProperty(uint rating) onlyTenant {
                propertyRating.totalRatings += rating;
                propertyRating.numberOfRatings++;
        }

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

    modifier onlyTenant() {
            if (msg.sender != tenant) throw;
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

        // reset values
        tenantOffer.tenant = 0;
        tenantOffer.startTime = 0;
        tenantOffer.endTime = 0;
        acceptedOffer = false;

        Payment(tenant, owner);
    }

    // owner will update the rent value and then the whole cycle of tenant offer can start
    function updateRent(uint _rent) onlyOwner {
        if (now < endTime) throw;

        rent = _rent;
    }

    // termination
    function terminate(uint deduction) payable onlyOwner {
        if (!(tenant.send(security - deduction) && owner.send(deduction))) throw;

        // reset values?
        tenantOffer.tenant = 0;
        tenantOffer.startTime = 0;
        tenantOffer.endTime = 0;
        acceptedOffer = false;
        tenant = 0;
        startTime = 0;
        endTime = 0;
    }

}
