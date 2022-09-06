var Vote = artifacts.require("Vote");
module.exports = function (deployer) {
    deployer.deploy(Vote, ["1", "2", "3"]);
    // Additional contracts can be deployed here
};
