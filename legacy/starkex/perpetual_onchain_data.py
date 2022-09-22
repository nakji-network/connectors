import argparse
import dataclasses
import json
import logging
import re
import sys
from collections import defaultdict
from typing import Dict, List, Mapping, Tuple

import web3
from eth_typing.encoding import HexStr
from web3.contract import Contract
from web3.exceptions import InvalidAddress
from web3.main import Web3

logging.basicConfig()
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)

ASSET_ID_UPPER_BOUND = 2**120
FUNDING_INDEX_LOWER_BOUND = -2**63
ONCHAIN_DATA_ASSET_ID_OFFSET = 64
DEFUALT_GET_LOGS_MAX_CHUNK_SIZE = 10 ** 4

MEMORY_PAGE_FACT_REGISTRY_START_BLOCK = 11813182
GPS_STATEMENT_VERIFIER_START_BLOCK = 12004790

USDC_RESOLUTION = 10 ** 6
USDC_CONTRACT_ADDRESS = '0x73dabaff2d1a1fd00cd11998d3cb8d3ae2d2fe8a'
CALL_PROXY_ABI = [
    {
        'constant': True,
        'inputs': [],
        'name': 'callProxyImplementation',
        'outputs': [{
            'internalType': 'address',
            'name': '_implementation',
            'type': 'address'
        }],
        'payable': False,
        'stateMutability': 'view',
        'type': 'function'
    },
    {
        'anonymous': False,
        'inputs': [
            {
                'indexed': True,
                'internalType': 'address',
                'name': 'implementation',
                'type': 'address'
            },
            {
                'indexed': False,
                'internalType': 'bytes',
                'name': 'initializer',
                'type': 'bytes'
            },
            {
                'indexed': False,
                'internalType': 'bool',
                'name': 'finalize',
                'type': 'bool'
            }
        ],
        'name': 'ImplementationAdded',
        'type': 'event'
    },
    {
        'anonymous': False,
        'inputs': [{
            'indexed': True,
            'internalType': 'address',
            'name': 'implementation',
            'type': 'address'
        }],
        'name': 'Upgraded',
        'type': 'event'
    }
]

CONTRACTS_JSON = {
    'StarkPerpetual': {
        'abi': [
            {
                'anonymous': False,
                'inputs': [
                    {
                        'indexed': False,
                        'internalType': 'bytes32',
                        'name': 'stateTransitionFact',
                        'type': 'bytes32'
                    }
                ],
                'name': 'LogStateTransitionFact',
                'type': 'event'
            },
            {
                'constant': True,
                'inputs': [],
                'name': 'getSystemAssetType',
                'outputs': [
                    {
                        'internalType': 'uint256',
                        'name': '',
                        'type': 'uint256'
                    }
                ],
                'payable': False,
                'stateMutability': 'view',
                'type': 'function'
            },
            {
                'constant': False,
                'inputs': [
                    {
                        'internalType': 'uint256',
                        'name': 'starkKey',
                        'type': 'uint256'
                    },
                    {
                        'internalType': 'uint256',
                        'name': 'assetType',
                        'type': 'uint256'
                    }
                ],
                'name': 'withdraw',
                'outputs': [],
                'payable': False,
                'stateMutability': 'nonpayable',
                'type': 'function'
            },
            {
                'inputs': [
                    {
                        'internalType': 'uint256',
                        'name': 'starkKey',
                        'type': 'uint256'
                    },
                    {
                        'internalType': 'uint256',
                        'name': 'vaultId',
                        'type': 'uint256'
                    },
                    {
                        'internalType': 'uint256',
                        'name': 'quantizedAmount',
                        'type': 'uint256'
                    }
                ],
                'name': 'escape',
                'outputs': [],
                'stateMutability': 'nonpayable',
                'type': 'function'
            }
        ],
        'address': '0xD54f502e184B6B739d7D27a6410a67dc462D69c8'
    },
    'MemoryPageFactRegistry': {
        'abi': [
            {
                'anonymous': False,
                'inputs': [
                    {
                        'indexed': False,
                        'internalType': 'bytes32',
                        'name': 'factHash',
                        'type': 'bytes32'
                    },
                    {
                        'indexed': False,
                        'internalType': 'uint256',
                        'name': 'memoryHash',
                        'type': 'uint256'
                    },
                    {
                        'indexed': False,
                        'internalType': 'uint256',
                        'name': 'prod',
                        'type': 'uint256'
                    }
                ],
                'name': 'LogMemoryPageFactContinuous',
                'type': 'event'
            },
            {
                'constant': False,
                'inputs': [
                    {
                        'internalType': 'uint256',
                        'name': 'startAddr',
                        'type': 'uint256'
                    },
                    {
                        'internalType': 'uint256[]',
                        'name': 'values',
                        'type': 'uint256[]'
                    },
                    {
                        'internalType': 'uint256',
                        'name': 'z',
                        'type': 'uint256'
                    },
                    {
                        'internalType': 'uint256',
                        'name': 'alpha',
                        'type': 'uint256'
                    },
                    {
                        'internalType': 'uint256',
                        'name': 'prime',
                        'type': 'uint256'
                    }
                ],
                'name': 'registerContinuousMemoryPage',
                'outputs': [
                    {
                        'internalType': 'bytes32',
                        'name': 'factHash',
                        'type': 'bytes32'
                    },
                    {
                        'internalType': 'uint256',
                        'name': 'memoryHash',
                        'type': 'uint256'
                    },
                    {
                        'internalType': 'uint256',
                        'name': 'prod',
                        'type': 'uint256'
                    }
                ],
                'payable': False,
                'stateMutability': 'nonpayable',
                'type': 'function'
            }
        ],
        'address': '0xEfbCcE4659db72eC6897F46783303708cf9ACef8'
    },
    'GpsStatementVerifier': {
        'abi': [
            {
                'anonymous': False,
                'inputs': [
                    {
                        'indexed': False,
                        'internalType': 'bytes32',
                        'name': 'factHash',
                        'type': 'bytes32'
                    },
                    {
                        'indexed': False,
                        'internalType': 'bytes32[]',
                        'name': 'pagesHashes',
                        'type': 'bytes32[]'
                    }
                ],
                'name': 'LogMemoryPagesHashes',
                'type': 'event'
            },
            {
                'constant': True,
                'inputs': [],
                'name': 'identify',
                'outputs': [
                    {
                        'internalType': 'string',
                        'name': '',
                        'type': 'string'
                    }
                ],
                'payable': False,
                'stateMutability': 'view',
                'type': 'function'
            }
        ],
        'address': '0xC8c212f11f6ACca77A7afeB7282dEBa5530eb46C'
    }
}


@dataclasses.dataclass
class FundingIndicesState:
    """
    Represents a collection of timestamped global funding indices for all assets.

    :param indices: Map of synthetic asset to its global funding index.
    :type indices: Mapping[int, int]
    :param timestamp: The timestamp for which the funding indices are valid for.
    :type timestamp: int
    """
    indices: Mapping[int, int]
    timestamp: int

    @classmethod
    def empty(cls) -> 'FundingIndicesState':
        return cls(indices={}, timestamp=0)


@dataclasses.dataclass
class PositionStateUpdate:
    """
    Holds inforamation about a single position update.
    Members:
     - position_id: the id of the position the update refers to.
     - public_key: the new public key of the position.
     - collateral_balance: the new amount of collateral_balance of the position.
     - funding_timestamp: The timestamp of the last funding that was applied to the position.
     - asset_balance_updates: information regarding the assets that changed in the position
     (contains asset_id and balance)
    """
    position_id: int
    public_key: int
    collateral_balance: int
    funding_timestamp: int
    asset_balance_updates: Dict[int, int]


@dataclasses.dataclass
class PerpetualOnChainData:
    """
    Dataclass for storing the onchain data.
    """
    funding_indices_mapping: Dict[int, FundingIndicesState]
    updates: List[PositionStateUpdate]


def load_contracts(
        web3: web3.Web3, contracts_names: List[str]) -> Dict[str, web3.contract.Contract]:
    """
    Given a list of contract names, returns a dict of contract names and contracts.
    """
    res = {}
    for contract_name in contracts_names:
        try:
            res[contract_name] = web3.eth.contract(
                address=CONTRACTS_JSON[contract_name]['address'],
                abi=CONTRACTS_JSON[contract_name]['abi'])
        except (KeyError, InvalidAddress) as ex:
            logger.error(
                f'Error in reading {contract_name} contract from CONTRACTS_JSON. '
                f'Please check that {contract_name} exists in CONTRACTS_JSON and that the address '
                'and abi are correct.')
            raise ex
    return res


def _get_contract_events(
        contract_event, from_block: int, to_block: int,
        get_logs_max_chunk_size: int) -> list:
    """
    Given a contract event and block numbers, retrieves a list of events in blocks
    [from_block, to_block).
    Splits the query in order to avoid Infura's maximal query limitation.
    See https://infura.io/docs/ethereum/json-rpc/eth_getLogs.
    """
    events = []
    assert from_block <= to_block, \
        f'from_block {from_block} must be smaller or equal to to_block {to_block}.'
    split_queries_block_nums = list(range(from_block, to_block, get_logs_max_chunk_size))
    split_queries = [
        (query_from_block, query_to_block)
        for query_from_block, query_to_block in zip(
            split_queries_block_nums, split_queries_block_nums[1:] + [to_block])]
    logger.info(
        f'Fetching {contract_event} for blocks [{from_block}, {to_block}]. '
        f'Number of queries to the web3 provider: {len(split_queries)}.')
    for query_from_block, query_to_block in split_queries:
        events.extend(list(
            contract_event.getLogs(fromBlock=query_from_block, toBlock=query_to_block)))
    return events


def _initialize_memory_page_map(
        memory_page_fact_registry_contract: Contract, from_block: int,
        to_block: int, get_logs_max_chunk_size: int) -> Dict[int, str]:
    """
    Returns a mapping between the memory pages' hashes and the Ethereum transaction's hash for the
    transactions in blocks [from_block, to_block).
    """
    memory_page_contract_event = \
        memory_page_fact_registry_contract.events.LogMemoryPageFactContinuous
    logger.info(f'Constructing memory pages dictionary for blocks [{from_block}, {to_block}].')
    memory_page_events = _get_contract_events(
        contract_event=memory_page_contract_event, from_block=from_block, to_block=to_block,
        get_logs_max_chunk_size=get_logs_max_chunk_size)
    return {
        event['args']['memoryHash']: event['transactionHash'].hex()
        for event in memory_page_events}


def _get_statement_verifier_memory_hashes_logs(
        statement_verifier_impl_contracts: List[Contract], from_block: int, to_block: int,
        get_logs_max_chunk_size: int) -> list:
    """
    Given a list of statement verifier implementation contracts and block numbers, returns the
    LogMemoryPagesHashes events for each contract for blocks [from_block, to_block).
    """
    statement_verifier_events = []
    for statement_verifier_impl_contract in statement_verifier_impl_contracts:
        # Asserts that the contract is a statement verifier implementation contract.
        assert 'GpsStatementVerifier' in \
            statement_verifier_impl_contract.functions.identify().call(), (
                f'Contract with address {statement_verifier_impl_contract.address} is not a '
                'statement verifier contract.')
        statement_verifier_contract_event = \
            statement_verifier_impl_contract.events.LogMemoryPagesHashes
        statement_verifier_events.extend(_get_contract_events(
            contract_event=statement_verifier_contract_event, from_block=from_block,
            to_block=to_block, get_logs_max_chunk_size=get_logs_max_chunk_size))
    return statement_verifier_events


def get_upgrade_events(proxy_contract) -> Dict[str, List[int]]:
    """
    Returns a dictionary with implementation address as keys,
    and the block numbers in which the upgrade event took place for those addresses as values.
    """
    upgrade_event_block_by_address = defaultdict(list)
    upgrade_to_events = proxy_contract.events.Upgraded().getLogs(fromBlock=0)
    upgrade_to_events = sorted(upgrade_to_events, key=lambda ev: ev.blockNumber)
    for upgrade_event in upgrade_to_events:
        impl_address = upgrade_event.args.implementation
        block = upgrade_event.blockNumber
        upgrade_event_block_by_address[impl_address].append(block)
    return dict(upgrade_event_block_by_address)


def get_add_events(proxy_contract) -> Dict[str, Dict[int, bytes]]:
    """
    Returns a dictionary with implementation address as keys,
        and dictionary of block numbers in which the addImplementation event took place as key,
            and then implementation initializer vector as value.
    """
    add_impl_event_info: dict = defaultdict(dict)
    add_impl_events = proxy_contract.events.ImplementationAdded().getLogs(fromBlock=0)
    add_impl_events = sorted(add_impl_events, key=lambda ev: ev.blockNumber)
    for add_impl_event in add_impl_events:
        impl_address = add_impl_event.args.implementation
        add_impl_event_info[
            impl_address][add_impl_event.blockNumber] = add_impl_event.args.initializer
    return dict(add_impl_event_info)


# TODO(Remo,25/04/2021): Add test for this.
def get_upgrade_detailed_event(proxy_contract) -> Dict[str, Dict[int, bytes]]:
    """
    Returns a dictionary with implementation address as keys,
        and dictionary of block numbers in which the upgradeTo event took place as key,
            and then implementation initializer vector as value.
    The Upgraded event does not contain initializer information,
    so this function merges information from ImplementationAdded & Upgraded events.
    """
    add_impl_events = get_add_events(proxy_contract)
    upgrade_events = get_upgrade_events(proxy_contract)
    detailed_events: Dict[str, Dict[int, bytes]] = defaultdict(dict)

    # We have two problems:
    # 1. We can upgrade to an eligible address anytime, as long as we use the correct initializer.
    # 2. We can jump to ANY eligible address, regardless of the order of addition.
    # 3. We also know that an address has only one current initializer, the last one set.
    # Therefore, We do the following:
    # For an upgrade event - i.e. an address-block combination, we search the addImpl event for that
    # address, the latest one that is not after the event, and extract the initializer from there.
    for address, upgrade_block_list in upgrade_events.items():
        add_block_list = sorted(list(add_impl_events[address].keys()), reverse=True)
        for upgrade_block in upgrade_block_list:
            block_with_data = None
            for potential_block in add_block_list:
                # We need the highest add block, that is smaller or equal to the upgrade block.
                if potential_block <= upgrade_block:
                    block_with_data = potential_block
                    break
            assert block_with_data, \
                f'Failed to find data for upgrade to address {address} in block {upgrade_block}'
            detailed_events[address][upgrade_block] = add_impl_events[address][block_with_data]
    return dict(detailed_events)


def extract_gps_verifiers(proxy_contract) -> List[Tuple[int, bytes]]:
    """
    For a proxy above a GpsStatementVerifier
    extract the upgrade events information and returns a list of the real verifier addresses
    and the block number in which they were set to be used.
    upgrade_info is a dict with address as keys (CallProxy address) and block:init_data values.
    From the init_data we extract the real verifier addresses.
    """
    upgrade_info = get_upgrade_detailed_event(proxy_contract=proxy_contract)
    flat_list = []
    for address, block_data_dict in upgrade_info.items():
        _items = list(block_data_dict.items())
        assert all(len(_item[1]) == 32 for _item in _items), 'Incompatible init data'
        flat_list += [
            (_item[0], Web3.toChecksumAddress(_item[1][-20:]))
            for _item in _items]
    flat_list = sorted(flat_list, key=lambda _item: _item[0])  # Sort by block number.
    return flat_list


class MemoryPagesFetcher:
    """
    Given a fact hash and using onchain data, retrieves the memory pages that the GPS statement
    verifier outputted for the relevant Cairo job.
    """

    def __init__(
            self, web3: Web3, memory_page_map: Dict[int, str], memory_page_facts_logs,
            memory_page_fact_registry_contract: Contract):
        self.web3 = web3
        self.memory_page_map = memory_page_map
        self.memory_page_facts_logs = memory_page_facts_logs
        self.memory_page_fact_registry_contract = memory_page_fact_registry_contract

    @classmethod
    def create(
            cls, web3: Web3, gps_statement_verifier_contract: Contract,
            memory_page_fact_registry_contract: Contract, is_verifier_proxied: bool,
            get_logs_max_chunk_size: int = DEFUALT_GET_LOGS_MAX_CHUNK_SIZE) -> 'MemoryPagesFetcher':
        """
        Creates an initialized instance by reading contract logs from the given web3 provider.
        If is_verifier_proxied is true, then gps_statement_verifier_contract is the proxy contract
        rather than the statement verifier implemantation.
        GetLogs call is splitted to chunks of size get_logs_max_chunk_size to avoid exceeding
        maximal query limitation.
        """
        last_block = web3.eth.block_number
        memory_page_map = _initialize_memory_page_map(
            memory_page_fact_registry_contract=memory_page_fact_registry_contract,
            from_block=MEMORY_PAGE_FACT_REGISTRY_START_BLOCK, to_block=last_block,
            get_logs_max_chunk_size=get_logs_max_chunk_size)
        if is_verifier_proxied:
            call_proxy = web3.eth.contract(
                address=gps_statement_verifier_contract.address, abi=CALL_PROXY_ABI)
            gps_statement_verifier_impl = extract_gps_verifiers(proxy_contract=call_proxy)
            gps_statement_verifier_impl_contracts = [
                web3.eth.contract(
                    abi=gps_statement_verifier_contract.abi,
                    address=gps_statement_verifier_impl_address)
                for _, gps_statement_verifier_impl_address in gps_statement_verifier_impl]
        else:
            gps_statement_verifier_impl_contracts = [gps_statement_verifier_contract]

        memory_page_facts_logs = _get_statement_verifier_memory_hashes_logs(
            statement_verifier_impl_contracts=gps_statement_verifier_impl_contracts,
            from_block=GPS_STATEMENT_VERIFIER_START_BLOCK, to_block=last_block,
            get_logs_max_chunk_size=get_logs_max_chunk_size)

        return cls(
            web3=web3, memory_page_map=memory_page_map,
            memory_page_facts_logs=memory_page_facts_logs,
            memory_page_fact_registry_contract=memory_page_fact_registry_contract)

    def _get_memory_pages_hashes_from_fact(self, fact_hash: bytes):
        """
        An auxiliary function for retrieveing the memory pages' hashes of a fact.
        """
        for log in self.memory_page_facts_logs:
            if log['args']['factHash'] == fact_hash:
                return log['args']['pagesHashes']
        raise Exception(f'Fact hash {fact_hash.hex()} was not registered in the verifier contract.')

    def get_memory_pages_from_fact(self, fact_hash: bytes) -> List[List[int]]:
        """
        Given a fact hash, retrieves the memory pages which are relevant for that fact.
        """
        memory_pages = []
        memory_pages_hashes = self._get_memory_pages_hashes_from_fact(fact_hash)

        assert memory_pages_hashes is not None
        for memory_page_hash in memory_pages_hashes:
            transaction_str = self.memory_page_map[int.from_bytes(memory_page_hash, 'big')]
            memory_pages_tx = self.web3.eth.getTransaction(HexStr(transaction_str))
            tx_decoded_values = self.memory_page_fact_registry_contract.decode_function_input(
                memory_pages_tx['input'])[1]['values']
            memory_pages.append(tx_decoded_values)
        return memory_pages


def get_asset_name_and_amount(asset_id: int, amount: int):
    asset_full_name = asset_id.to_bytes(15,'big').decode("utf-8")
    asset_name, resolution = re.split('-|\x00', asset_full_name)[:2]
    return asset_name, amount / (10 ** int(resolution))

def get_balance_and_asset_id(asset_serialized: int) -> Tuple[int, int]:
    """
    Given the on-chain-data serialization of a position_asset, returns its asset_id and balance.
    [ padding - 72 | asset_id - 120 | biased_balance - 64 bit ]
    """
    asset_id = \
        asset_serialized >> ONCHAIN_DATA_ASSET_ID_OFFSET % ASSET_ID_UPPER_BOUND
    biased_balance = asset_serialized & (2 ** ONCHAIN_DATA_ASSET_ID_OFFSET - 1)
    return asset_id, biased_balance + FUNDING_INDEX_LOWER_BOUND


def parse_funding_indices(data: List[int]) -> 'FundingIndicesState':
    """
    Given a list of ints that corresponds to a serialized funding indices, returns the deserealized
    object FundingIndicesState.
    [ asset_id_0, funding_index_0, ..., asset_id_n, funding_index_n, timestamp]
    """
    timestamp, data = data[-1], data[:-1]
    indices = {}
    assert len(data) % 2 == 0
    for asset_id, funding_index in zip(data[::2], data[1::2]):
        indices[asset_id] = funding_index + FUNDING_INDEX_LOWER_BOUND
    return FundingIndicesState(indices=indices, timestamp=timestamp)


def parse_position_state_update(data: List[int]) -> 'PositionStateUpdate':
    """
    Given a list of ints that corresponds to a serialized position state update, returns the
    deserealized object PositionStateUpdate.
    """
    position_id, public_key, biased_balance, funding_timestamp, *asset_serialization = data
    collateral_balance = biased_balance + FUNDING_INDEX_LOWER_BOUND
    asset_ids_and_balances = [get_balance_and_asset_id(asset) for asset in asset_serialization]
    return PositionStateUpdate(
        position_id=position_id,
        public_key=public_key,
        collateral_balance=collateral_balance,
        asset_balance_updates=dict(asset_ids_and_balances),
        funding_timestamp=funding_timestamp)


def parse_onchain_data(values: List[int]) -> 'PerpetualOnChainData':
    """
    Parses the on-chain data and returns the changes in an OnChainData instance with the parsed
    changes. The changes are encoded using 256bit words as follows:
    [- funding table size -----] <- 1 word size, index 0
    [- first_entry n_assets ---] <- 1 word size, index 1.
    [- first_entry assets -----] <- 2 * n_assets, [asset_id, funding_index]
           ...                   <-
    [- first_entry timestamp --] <- 1 word size, timestamp
    [- position update 0 size -] <- 1 word size, position_update_size (depends on number of assets)
    [- position update 0 ------] <- position update 0 size words, contains position_id, public_key,
           ...                   <- biased_balance, funding_timestamp, asset_serialization
    [- position update n ------] <- last update
    """
    position_state_updates = []
    funding_indices_mapping = {}
    values, funding_indices_table_size = values[1:], values[0]
    for _ in range(funding_indices_table_size):
        n_funding_indices = values[0]
        # Funding indices composed of n_funding_indices assets and timestamp.
        funding_indices = parse_funding_indices(values[1: 2 * n_funding_indices + 1 + 1])
        funding_indices_mapping[funding_indices.timestamp] = funding_indices
        values = values[2 * n_funding_indices + 2:]

    while len(values) != 0:
        curr_position_state_values = values[0]
        position_state_updates.append(parse_position_state_update(
            values[1: curr_position_state_values + 1]))
        values = values[curr_position_state_values + 1:]

    return PerpetualOnChainData(
        funding_indices_mapping=funding_indices_mapping, updates=position_state_updates)


def fetch_memory_pages_from_onchain_data(
        memory_pages_fetcher: MemoryPagesFetcher, verified_fact_event) -> List[List[int]]:
    """
    Given a StarkEx fact event, returns the memory pages that are relevant for this event.
    """
    fact_hash = verified_fact_event['args']['stateTransitionFact']
    return memory_pages_fetcher.get_memory_pages_from_fact(fact_hash)


def print_onchain_data(onchain_data: PerpetualOnChainData, block_num: int):
    print(f"Printing position's updates for block {block_num}")
    print('-------------------------------------------------------------')
    print(f'Number of updates in this batch: {len(onchain_data.updates)}')
    headers = [
        'position id', 'stark key', 'collateral balance', 'funding timestamp', 'position_assets']
    positions_updates = [
        [
            update.position_id, update.public_key, update.collateral_balance,
            update.funding_timestamp, update.asset_balance_updates.items()]
        for update in onchain_data.updates]
    print('headers:\n' + '\t'.join(headers) + '\n')
    for position in positions_updates:
        print(f'position: {position[0]}:\n{position}\n')

    if onchain_data.funding_indices_mapping:
        print("\nPrinting position's funding_indices_mapping:\n")
        print('timestamp\t funding indices')
        for timestamp, funding_indices in onchain_data.funding_indices_mapping.items():
            print(f'{timestamp}\t {funding_indices.indices}\n')


def parse_cmdline():
    parser = argparse.ArgumentParser(description='Blockchain Explorer')
    parser.add_argument(
        '--web3_http_provider', type=str, required=True,
        help='HTTP URL to the Ethereum node.')
    parser.add_argument(
        '--from_block', type=int, required=True,
        help='Ethereum block number')
    parser.add_argument(
        '--to_block', type=int, 
        help='Ethereum block number')
    parser.add_argument(
        '--output_file', type=str, default='output.json',
        help='Script output json file name, default is ./output.json')
    parser.add_argument(
        '--proxied_verifier', dest='is_verifier_proxied', action='store_true',
        help='Indicator that the statement verifier contract is proxy contract.')
    parser.add_argument(
        '--implementation_verifier', dest='is_verifier_proxied', action='store_false',
        help='Indicator that the statement verifier contract is implementation contract.')
    parser.set_defaults(is_verifier_proxied=True)
    return parser.parse_args()


def get_event_updates(verified_fact_event, memory_pages_fetcher):
    """
    Given verified_fact_event of a batch and memory pages_fetcher object, retrieve the memory pages
    that corresponds to the batch and parses it to an event update.
    """
    batch_memory_pages = fetch_memory_pages_from_onchain_data(
        memory_pages_fetcher=memory_pages_fetcher, verified_fact_event=verified_fact_event)
    # Asserts that at least one data availability page exists.
    assert len(batch_memory_pages) > 1, (
        "Data availability pages doesn't found for fact "
        f"{verified_fact_event['args']['stateTransitionFact']}")
    batch_onchain_data = [
        value for memory_page in batch_memory_pages[1:] for value in memory_page]
    return parse_onchain_data(values=batch_onchain_data), verified_fact_event['blockNumber']


def dump_json_file(position_update_list: list, output_file: str):
    with open(output_file, 'w') as file:
        json.dump(position_update_list, file, indent=4)


def dump_onchain_data(block, onchain_data: PerpetualOnChainData) -> dict:
    positions_updates_dict = {}
    for position_update in onchain_data.updates:
        positions_updates_dict.update(dump_position_update(block, position_update=position_update))
    return positions_updates_dict


def dump_position_update(block, position_update: PositionStateUpdate) -> dict:
    return {
        position_update.position_id: {
            'stark_key': hex(position_update.public_key),
            'assets': dump_position_update_assets(position_update=position_update),
            'block_num': block.number,
            'timestamp': block.timestamp
        }
    }


def dump_position_update_assets(position_update: PositionStateUpdate) -> dict:
    assets_dict = {
        "USDC": {
            'asset_type': 'ERC20',
            'amount': str(position_update.collateral_balance / USDC_RESOLUTION),
            'additional_attributes': {}
        }
    }
    synt_asset = [
        get_asset_name_and_amount(asset_id, amount)
        for asset_id, amount in position_update.asset_balance_updates.items()]
    synt_asset_dict = {
        asset_name: {
            'asset_type': 'SYNTH',
            'amount': str(real_amount),
            'additional_attributes': {
                'cached_funding_index': position_update.funding_timestamp
            }
        }
        for asset_name, real_amount in synt_asset}
    assets_dict.update(synt_asset_dict)
    return assets_dict


def main():
    """
    Given web3 http provider and block number, prints the batches updates that are registered on
    this block. If such updates do not exists, then prints "No position changes were found in block
    {fromblock} to {toBlock}".
    In order to run the script, you need to add the correct ethereum contracts addresses in the
    relevant places (marked with <Fill>) under CONTRACTS_JSON.
    """
    args = parse_cmdline()

    # Connect web3.
    web3 = Web3(Web3.HTTPProvider(args.web3_http_provider, request_kwargs={'timeout': 200}))
    assert web3.isConnected(), f'Cannot connect to http provider {args.web3_http_provider}.'
    logger.info('web3 connected.')

    # Get contracts.
    contract_names = ['StarkPerpetual', 'GpsStatementVerifier', 'MemoryPageFactRegistry']
    contracts_dict = load_contracts(web3=web3, contracts_names=contract_names)
    (main_contract, gps_statement_verifier_contract, memory_pages_contract) = \
        [contracts_dict[contract_name] for contract_name in contract_names]

    logger.info(f'Main Contract Events: {main_contract.events}')

    to_block = args.to_block or web3.eth.get_block_number()
    main_contract_events = main_contract.events.LogStateTransitionFact.createFilter(
        fromBlock=args.from_block,toBlock=to_block).get_all_entries()

    if not main_contract_events:
        logger.info(f'No position changes were found in block {args.from_block} to {to_block}.')
        exit(0)

    # Fetch the updates.
    memory_pages_fetcher = MemoryPagesFetcher.create(
        web3=web3, is_verifier_proxied=args.is_verifier_proxied,
        gps_statement_verifier_contract=gps_statement_verifier_contract,
        memory_page_fact_registry_contract=memory_pages_contract)

    position_update_list = []
    for main_contract_event in main_contract_events:
        onchain_data, block_num = get_event_updates(main_contract_event, memory_pages_fetcher)
        block = web3.eth.get_block(block_num)
        position_changes = list(dump_onchain_data(block, onchain_data=onchain_data).values())
        position_update_list.extend(position_changes)
    dump_json_file(position_update_list=position_update_list, output_file=args.output_file)


if __name__ == '__main__':
    sys.exit(main())
