from web3 import Web3, HTTPProvider
from time import sleep
from getpass import getpass
import json

dev = True

# if we are in dev mode connect to my poa network
if dev == True:
	w3 = Web3(HTTPProvider('http://127.0.0.1:8501'))
else:
	w3 = Web3(HTTPProvider('http://127.0.0.1:8545'))

abiPath = input('enter path to abi file\t')
bytePath = input('enter path to bytecdoe file\t')

with open(abiPath, 'r') as fh:
	abi = json.load(fh)

with open(bytePath, 'r') as fh:
	bytecode = fh.read()

w3.personal.unlockAccount(w3.eth.accounts[0], getpass('pw\t'), 0)
contract = w3.eth.contract(abi=abi, bytecode=bytecode.strip('\n'))
tx_hash = contract.deploy(transaction={'from': w3.eth.accounts[0], 'gas': 4100000})
choice = input('press any key to continue')
tx_receipt = w3.eth.getTransactionReceipt(tx_hash)
contractAddress = tx_receipt['contractAddress']
print(contractAddress)
