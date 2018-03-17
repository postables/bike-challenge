pragma solidity 0.4.21;

import "Interfaces/AdministrationInterface.sol";
import "Interfaces/ERC20Interface.sol";
import "Modules/Administration.sol";

contract Factory is Administration {


	bytes constant private bikeTokenContract = hex"6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a0319918216811790925560018054909116909117905560408051908101604052600981527f42696b65546f6b656e000000000000000000000000000000000000000000000060208201526002908051610081929160200190610106565b5060408051908101604052600281527f4254000000000000000000000000000000000000000000000000000000000000602082015260039080516100c9929160200190610106565b506005805460ff191660121790556a108b2a2c280290940000006004819055600160a060020a0333166000908152600660205260409020556101a1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014757805160ff1916838001178555610174565b82800160010185558215610174579182015b82811115610174578251825591602001919060010190610159565b50610180929150610184565b5090565b61019e91905b80821115610180576000815560010161018a565b90565b610a65806101b06000396000f3006060604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100df578063095ea7b31461016957806318160ddd1461019f57806323b872dd146101c457806327e235e3146101ec578063313ce5671461020b5780635c65816514610234578063704b6c021461025957806370a08231146102785780638da5cb5b1461029757806395d89b41146102c6578063a6f9dae1146102d9578063a9059cbb146102f8578063dd62ed3e1461031a578063f851a4401461033f575b600080fd5b34156100ea57600080fd5b6100f2610352565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561012e578082015183820152602001610116565b50505050905090810190601f16801561015b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017457600080fd5b61018b600160a060020a03600435166024356103f0565b604051901515815260200160405180910390f35b34156101aa57600080fd5b6101b2610494565b60405190815260200160405180910390f35b34156101cf57600080fd5b61018b600160a060020a036004358116906024351660443561049a565b34156101f757600080fd5b6101b2600160a060020a036004351661063e565b341561021657600080fd5b61021e610650565b60405160ff909116815260200160405180910390f35b341561023f57600080fd5b6101b2600160a060020a0360043581169060243516610659565b341561026457600080fd5b61018b600160a060020a0360043516610676565b341561028357600080fd5b6101b2600160a060020a03600435166106ff565b34156102a257600080fd5b6102aa61071a565b604051600160a060020a03909116815260200160405180910390f35b34156102d157600080fd5b6100f2610729565b34156102e457600080fd5b61018b600160a060020a0360043516610794565b341561030357600080fd5b61018b600160a060020a036004351660243561081d565b341561032557600080fd5b6101b2600160a060020a03600435811690602435166108f4565b341561034a57600080fd5b6102aa61091f565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103e85780601f106103bd576101008083540402835291602001916103e8565b820191906000526020600020905b8154815290600101906020018083116103cb57829003601f168201915b505050505081565b600160a060020a033381166000908152600760209081526040808320938616835292905290812054610428908363ffffffff61092e16565b600160a060020a03338116600081815260076020908152604080832094891680845294909152908190209390935590917f6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e9085905190815260200160405180910390a350600192915050565b60045490565b60006104a7848484610947565b15156104b257600080fd5b600160a060020a0380851660009081526007602090815260408083203390941683529290522054829010156104e657600080fd5b600160a060020a0380851660009081526007602090815260408083203390941683529290529081205461051f908463ffffffff610a2416565b101561052a57600080fd5b600160a060020a0380851660009081526007602090815260408083203390941683529290522054610561908363ffffffff610a2416565b600160a060020a0380861660008181526007602090815260408083203390951683529381528382209490945590815260069092529020546105a8908363ffffffff610a2416565b600160a060020a0380861660009081526006602052604080822093909355908516815220546105dd908363ffffffff61092e16565b600160a060020a03808516600081815260066020526040908190209390935591908616907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b60066020526000908152604090205481565b60055460ff1681565b600760209081526000928352604080842090915290825290205481565b6000805433600160a060020a0390811691161461069257600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c82604051600160a060020a03909116815260200160405180910390a1506001919050565b600160a060020a031660009081526006602052604090205490565b600054600160a060020a031681565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103e85780601f106103bd576101008083540402835291602001916103e8565b6000805433600160a060020a039081169116146107b057600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc82604051600160a060020a03909116815260200160405180910390a1506001919050565b600061082a338484610947565b151561083557600080fd5b600160a060020a03331660009081526006602052604090205461085e908363ffffffff610a2416565b600160a060020a033381166000908152600660205260408082209390935590851681522054610893908363ffffffff61092e16565b600160a060020a0380851660008181526006602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260076020908152604080832093909416825291909152205490565b600154600160a060020a031681565b60008282018381101561094057600080fd5b9392505050565b6000600160a060020a038416158015906109695750600160a060020a03831615155b80156109755750600082115b151561098057600080fd5b600160a060020a0384166000908152600660205260408120546109a9908463ffffffff610a2416565b10156109b457600080fd5b600160a060020a0383166000908152600660205260408120546109dd908463ffffffff61092e16565b116109e757600080fd5b600160a060020a038316600090815260066020526040902054610a10818463ffffffff61092e16565b11610a1a57600080fd5b5060019392505050565b600082821115610a3357600080fd5b509003905600a165627a7a7230582082a389e55dfae8434b69adc9f31041d11d0bf7fbc751746e370d12edce5409b00029";
	bytes constant private btEthContract = hex"606060405260008054600160a060020a033316600160a060020a031991821681179092556001805490911690911790556109d08061003e6000396000f3006060604052600436106100d75763ffffffff60e060020a600035041663158ef93e8114610113578063214013ca1461013a578063238dafe01461015957806342e94c901461016c578063524592a11461019d57806361348708146101b3578063630b4989146101c6578063704b6c02146101d957806382f5decb146101f857806388a244ee1461020b5780638da5cb5b1461021e578063a6f9dae11461024d578063bf5839031461026c578063c63df5ef1461027f578063e3b172d714610292578063f851a440146102a5578063fddf0fc0146102b8575b600654610100900460ff1615156100ed57600080fd5b60065460ff1615156100fe57600080fd5b6101066102cb565b151561011157600080fd5b005b341561011e57600080fd5b6101266104be565b604051901515815260200160405180910390f35b341561014557600080fd5b610126600160a060020a03600435166104cc565b341561016457600080fd5b61012661065f565b341561017757600080fd5b61018b600160a060020a0360043516610668565b60405190815260200160405180910390f35b34156101a857600080fd5b61018b60043561067a565b34156101be57600080fd5b6101266106a9565b34156101d157600080fd5b61018b610729565b34156101e457600080fd5b610126600160a060020a036004351661072f565b341561020357600080fd5b61018b6107b8565b341561021657600080fd5b61018b6107c2565b341561022957600080fd5b6102316107c8565b604051600160a060020a03909116815260200160405180910390f35b341561025857600080fd5b610126600160a060020a03600435166107d7565b341561027757600080fd5b61018b610860565b341561028a57600080fd5b610126610866565b341561029d57600080fd5b6101266108e7565b34156102b057600080fd5b610231610923565b34156102c357600080fd5b61018b610932565b600080348190116102db57600080fd5b600254600090116102eb57600080fd5b61032a67016345785d8a000061031e651cf33b72a60061031234600a63ffffffff61093816565b9063ffffffff61096616565b9063ffffffff61093816565b90506000811161033657fe5b600160a060020a03331660009081526007602052604090205461035f908263ffffffff61097d16565b600160a060020a03331660009081526007602052604090205560025461038b908263ffffffff61098f16565b6002556003546103a1908263ffffffff61097d16565b6003556004546103b890600163ffffffff61097d16565b60045560025415156103fb576006805460ff191690557f3267a72caa59ef19fda30fe15c5dbd5ba0e4d3c0ced49756eb70da19f27d0e4760405160405180910390a15b33600160a060020a03167f57d61f3ccd4ccd25ec5d234d6049553a586fac134c85c98d0b0d9d5724f4e43e8260405190815260200160405180910390a2600654620100009004600160a060020a031663a9059cbb338360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561049457600080fd5b5af115156104a157600080fd5b5050506040518051905015156104b657600080fd5b600191505090565b600654610100900460ff1681565b6000805433600160a060020a039081169116146104e857600080fd5b600654610100900460ff16156104fd57600080fd5b6006805475ffffffffffffffffffffffffffffffffffffffff0000191662010000600160a060020a0385811682029290921792839055651cf33b72a6009204166370a082313060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561058157600080fd5b5af1151561058e57600080fd5b505050604051805190501115156105a457600080fd5b600654620100009004600160a060020a03166370a082313060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156105fa57600080fd5b5af1151561060757600080fd5b5050506040518051600255506006805460ff1961ff0019909116610100171660011790557f68f3e71400a107502ed4901f333a22781427bea3f4c32385e9c80cb46468ec2660405160405180910390a1506001919050565b60065460ff1681565b60076020526000908152604090205481565b60006106a367016345785d8a000061031e651cf33b72a60061031286600a63ffffffff61093816565b92915050565b6000805433600160a060020a039081169116146106c557600080fd5b60065460ff1615156106d657600080fd5b600654610100900460ff1615156106ec57600080fd5b6006805460ff191690557f3267a72caa59ef19fda30fe15c5dbd5ba0e4d3c0ced49756eb70da19f27d0e4760405160405180910390a15060015b90565b60045481565b6000805433600160a060020a0390811691161461074b57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c82604051600160a060020a03909116815260200160405180910390a1506001919050565b651cf33b72a60081565b60035481565b600054600160a060020a031681565b6000805433600160a060020a039081169116146107f357600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc82604051600160a060020a03909116815260200160405180910390a1506001919050565b60025481565b6000805433600160a060020a0390811691161461088257600080fd5b60065460ff161561089257600080fd5b600654610100900460ff1615156108a857600080fd5b6006805460ff191660011790557fec8f62fea11b2d0b268c0f08bc273cef8aeeb2347809eaf5ce4a1a04bb19c99160405160405180910390a150600190565b60065460009060ff610100909104161515600114801561090e575060065460ff1615156001145b1561091b57506001610726565b506000610726565b600154600160a060020a031681565b60055481565b6000828202831580610954575082848281151561095157fe5b04145b151561095f57600080fd5b9392505050565b600080828481151561097457fe5b04949350505050565b60008282018381101561095f57600080fd5b60008282111561099e57600080fd5b509003905600a165627a7a7230582080dcf037f7c8f12a4ffbcf8fe824f7d77caafd4876d7c4b64e4837764bc2fc170029";
	bytes constant private bikeRentalContract = hex"60606040526003805460ff1916600190811761010060b060020a0319169091556000805433600160a060020a0316600160a060020a03199182168117835583549091161790915561124a90819061005690396000f3006060604052600436106100cc5763ffffffff60e060020a6000350416630436156981146100d15780635c526e19146101045780636ba85eb414610155578063704b6c02146101bc57806375f1473f146101db57806389e1476f146101f15780638da5cb5b1461021357806391cca3db146102425780639616734a14610255578063a56a91c51461026e578063a6f9dae114610284578063c4540dc1146102a3578063cb24450d146102c5578063d7d172fe146102e1578063f4f3bc90146102f7578063f851a4401461031c575b600080fd5b34156100dc57600080fd5b6100f0600160a060020a036004351661032f565b604051901515815260200160405180910390f35b341561010f57600080fd5b61011a6004356103af565b6040518085815260200184815260200183600381111561013657fe5b60ff168152911515602083015250604090810193509150505180910390f35b341561016057600080fd5b610174600160a060020a03600435166103dc565b604051600160a060020a03909716875260208701959095526040808701949094526060860192909252608085015260a084015290151560c083015260e0909101905180910390f35b34156101c757600080fd5b6100f0600160a060020a0360043516610427565b34156101e657600080fd5b6100f06004356104b0565b34156101fc57600080fd5b6100f0600160a060020a0360043516602435610513565b341561021e57600080fd5b610226610778565b604051600160a060020a03909116815260200160405180910390f35b341561024d57600080fd5b6100f0610787565b341561026057600080fd5b6100f0600435602435610790565b341561027957600080fd5b6100f06004356108f2565b341561028f57600080fd5b6100f0600160a060020a03600435166109c9565b34156102ae57600080fd5b6100f0600160a060020a0360043516602435610a52565b34156102d057600080fd5b6100f0600435602435604435610ade565b34156102ec57600080fd5b6100f0600435610e55565b341561030257600080fd5b61030a610e72565b60405190815260200160405180910390f35b341561032757600080fd5b610226610e78565b6000805433600160a060020a0390811691161461034b57600080fd5b6003805475ffffffffffffffffffffffffffffffffffffffff0000191662010000600160a060020a038516021790557fb29c3cdc5585c95d761d74073070493d10b570687a64a745578a4efec2d58a2560405160405180910390a15060015b919050565b60046020526000908152604090208054600182015460029092015490919060ff8082169161010090041684565b60056020819052600091825260409091208054600182015460028301546003840154600485015495850154600690950154600160a060020a0390941695929491939092919060ff1687565b6000805433600160a060020a0390811691161461044357600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c82604051600160a060020a03909116815260200160405180910390a1506001919050565b60008060008381526004602052604090206002015460ff1660038111156104d357fe5b14806104fe5750600260008381526004602052604090206002015460ff1660038111156104fc57fe5b145b1561050b575060016103aa565b5060006103aa565b60008054819033600160a060020a0390811691161461053157600080fd5b600160a060020a0384166000908152600560205260409020600601548490849060ff161515600114801561057f5750600160a060020a03821660009081526005602052604090206001015481145b80156105aa5750600160008281526004602052604090206002015460ff1660038111156105a857fe5b145b15156105b557600080fd5b600160a060020a038616600090815260056020819052604090912001546105e59062278d0063ffffffff610e8716565b4210156105f157600080fd5b600160a060020a038616600090815260056020526040902060038101546002909101546106239163ffffffff610ea016565b600160a060020a0387166000908152600560208181526040808420805473ffffffffffffffffffffffffffffffffffffffff19168155600181018590556002808201869055600380830187905560048084018890559583018790556006909201805460ff199081169091558c87529490935293819020909101805490921690921790559093507f15c8a138c30e2363cd45884a114dc17ee4f798ba3e652372ab5928a43a18be25908790879051600160a060020a03909216825260208201526040908101905180910390a1600354620100009004600160a060020a031663a9059cbb338560405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561074a57600080fd5b5af1151561075757600080fd5b50505060405180519050151561076c57600080fd5b50600195945050505050565b600054600160a060020a031681565b60035460ff1681565b6000805433600160a060020a039081169116146107ac57600080fd5b826000808281526004602052604090206002015460ff1660038111156107ce57fe5b1480156107f25750600081815260046020526040902060020154610100900460ff16155b15156107fd57600080fd5b608060405190810160405280858152602001848152602001600360019054906101000a900460ff16600381111561083057fe5b8152600160209182015260008681526004909152604090208151815560208201518160010155604082015160028201805460ff1916600183600381111561087357fe5b02179055506060820151600291820180549115156101000261ff0019909216919091179055546108ab9150600163ffffffff610e8716565b6002557fea1ff50cba48e6d14a4f84fe71a137438579dcb2f65a99aab7a6e14918531416848460405191825260208201526040908101905180910390a15060019392505050565b600160a060020a0333908116600090815260056020526040812060060154909190839060ff16151560011480156109435750600160a060020a03821660009081526005602052604090206001015481145b801561096e5750600160008281526004602052604090206002015460ff16600381111561096c57fe5b145b151561097957600080fd5b600160a060020a0333166000908152600560208190526040909120015442116109b9576109a584610eb5565b15156109b057600080fd5b600192506109c2565b6109a584611068565b5050919050565b6000805433600160a060020a039081169116146109e557600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc82604051600160a060020a03909116815260200160405180910390a1506001919050565b600160a060020a03821660009081526005602052604081206006015460ff1615156001148015610a9c5750600160a060020a03831660009081526005602052604090206001015482145b8015610ac75750600160008381526004602052604090206002015460ff166003811115610ac557fe5b145b15610ad457506001610ad8565b5060005b92915050565b60008080858160008281526004602052604090206002015460ff166003811115610b0457fe5b148015610b2d575060008181526004602052604090206002015460ff6101009091041615156001145b1515610b3857600080fd5b33600160a060020a03811660009081526005602052604090206006015460ff1615610b6257600080fd5b600088815260046020526040902060010154610b9790600390610b8b908a63ffffffff6111f716565b9063ffffffff6111f716565b8614610ba257600080fd5b600088815260046020526040902060010154610bc4908863ffffffff6111f716565b9350610bd9876201518063ffffffff6111f716565b9250610beb428463ffffffff610e8716565b60008981526004602052604090819020600201805460ff1916600117905590935060e090519081016040908152600160a060020a03331680835260208084018c90528284018a905260608401889052608084018b905260a08401879052600160c0850152600091825260059052208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03919091161781556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151600691909101805460ff1916911515919091179055507fabbb103bc3d5dd5c36c3a0dbcb3bcef8158419c90cbc424f1ad059b3e6bd24083389896040518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390a1600354620100009004600160a060020a03166323b872dd33308960405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610d9257600080fd5b5af11515610d9f57600080fd5b505050604051805190501515610db457600080fd5b600354600054600160a060020a03620100009092048216916323b872dd913391168760405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610e2557600080fd5b5af11515610e3257600080fd5b505050604051805190501515610e4757600080fd5b506001979650505050505050565b600090815260046020526040902060020154610100900460ff1690565b60025481565b600154600160a060020a031681565b600082820183811015610e9957600080fd5b9392505050565b600082821115610eaf57600080fd5b50900390565b600160a060020a033316600090815260056020526040812060038101546002909101548291610eea919063ffffffff610ea016565b600160a060020a0333166000908152600560208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916815560018082018690556002808301879055600380840188905560048085018990559684018890556006909301805460ff1990811690915583548c89529690955292909520909101805495965061010090930460ff169492939290911691908490811115610f8857fe5b02179055507fd5aa72d703f442b594a668fb96e436d34a8defaa6d4318df92f0c652caa8c58f33846000604051600160a060020a039093168352602083019190915215156040808301919091526060909101905180910390a1600354620100009004600160a060020a031663a9059cbb338360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561103d57600080fd5b5af1151561104a57600080fd5b50505060405180519050151561105f57600080fd5b50600192915050565b600160a060020a03331660009081526005602052604081206003810154600290910154829161109d919063ffffffff610ea016565b600160a060020a0333166000908152600560208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916815560018082018690556002808301879055600380840188905560048085018990559684018890556006909301805460ff1990811690915583548c89529690955292909520909101805495965061010090930460ff16949293929091169190849081111561113b57fe5b02179055507fd5aa72d703f442b594a668fb96e436d34a8defaa6d4318df92f0c652caa8c58f33846001604051600160a060020a039093168352602083019190915215156040808301919091526060909101905180910390a1600354600054600160a060020a036201000090920482169163a9059cbb91168360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561103d57600080fd5b6000828202831580611213575082848281151561121057fe5b04145b1515610e9957600080fd00a165627a7a723058209e9746140e4005f9f5b51bcf8cf3c8c83d577b96c36101788585c729236656310029";

	address private unitTester;

	event ContractCreated(address _contractAddress, string _contractName);

	function setUnitTesterContract(address _contractAddress) external onlyOwner returns (bool) {
		unitTester = _contractAddress;
		return true;
	}

	function sendTokens(address _tokenContract) external onlyOwner returns (bool) {
		ERC20Interface ercI = ERC20Interface(_tokenContract);
		require(ercI.transfer(unitTester, ercI.balanceOf(address(this))));
		return true;
	}

	// ok
	function deployTokenContract() external returns (bool) {
		bytes memory _code = bikeTokenContract;
		address a;
		assembly {
			a := create(0, add(_code, 0x20), mload(_code))
		}
		AdministrationInterface(a).changeOwner(unitTester);
		emit ContractCreated(a, "token contract");
		assert(a != address(0));
	}

	// ok
	function deploySaleContract() external returns (bool) {
		bytes memory _code = btEthContract;
		address a;
		assembly {
			a := create(0, add(_code, 0x20), mload(_code))
		}
		AdministrationInterface(a).changeOwner(unitTester);
		emit ContractCreated(a, "token sale contract");
		assert(a != address(0));
	}

	// ok
	function deployRentalContract() external returns (bool) {
		bytes memory _code = bikeRentalContract;
		address a;
		assembly {
			a := create(0, add(_code, 0x20), mload(_code))
		}
		AdministrationInterface(a).changeOwner(unitTester);
		emit ContractCreated(a, "bike rental contract");
		assert(a != address(0));
	}


}