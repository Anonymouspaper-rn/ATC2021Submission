<!-- vscode-markdown-toc -->
* 1. [Schema](#Schema)
	* 1.1. [Block](#Block)
	* 1.2. [Code](#Code)
	* 1.3. [State](#State)
* 2. [Connection methods](#Connectionmethods)
* 3. [Search examples](#Searchexamples)
	* 3.1. [Search a transaction by transaction hash](#Searchatransactionbytransactionhash)
	* 3.2. [Search unique events and corresponding transactions' hash for BZx.](#SearchuniqueeventsandcorrespondingtransactionshashforBZx.)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

# The usage of data aggregator's database part

The database is based on Elasticsearch, and consists of three indices: block, code, and state. The search tasks of each index are independent. 

##  1. <a name='Schema'></a>Schema

The file (main.go) is used to create the schema of database.

###  1.1. <a name='Block'></a>Block

The block index contains most of the states in blockchain, which include
block information, normal transaction information, internal transaction information and demand-driven accounts' states. The details are shown in below.

*Block Information*
- Difficulty, long.
- ExtraData, text.
- GasLimit, long.
- GasUsed, long.
- Hash, keyword.
- Miner, keyword.
- Number, integer.
- TxnCount, integer.
*Normal transaction information*
- Transactions, nested:
    - CallFunction, keyword: the first 4 bytes of input data.
    - CallParameter, text: the remaining part of input data except for the first 4 bytes.
    - ConAddress, keyword: the address of created smart contract when the current normal transaction is a contract creation transaction.
    - CumGasUsed, long: the cumulative gas consumption of transactions in the same block before the current normal transaction.
    - FromAddress, keyword.
    - GasLimit, long.
    - GasPrice, keyword.
    - GasUsed, long.
    - Hash, keyword.
    - IntTxnCount, integer: the number of internal transactions triggered by the current normal transaction.
    - Nonce, integer.
    - Status, boolean.
    - ToAddress, keyword.
    - TxnIndex, integer: the serial number of the current normal transaction in the current block.
    - Value, keyword.
    *Internal transaction information*
    - InternalTxns, nested:
        - CallFunction, keyword.
        - CallParameter, text.
        - ConAddress, keyword.
        - EvmDepth, short: the depth of EVM stack when execution the current internal transaction.
        - FromAddress, keyword.
        - GasLimit, long.
        - OutPut, text: the return value of the current internal transaction.
        - ToAddress, keyword.
        - TxnIndex, integer: the serial number of the current internal transaction.
        - Type, short: the methods to trigger the current internal transaction, 240: CREATE, 241: CALL, 242: CALLCODE, 244: DELEGATECALL, 245: CREATE2, 250: STATICCALL.
    *Events*
    - Logs, nested:
        - Address, keyword: the address of the smart contract that contains the current event.
        - Topics, keyword array: the first topic is the event signature, the remaining topics are the indexed parameters of event.
        - Data, text: the normal parameters of event.
    *Demand-driven accounts' states*
    - ReadCommittedState, nested: accounts' states the transactions will load.
        - Address, keyword.
        - Balance, keyword.
        - CodeHash, keyword.
        - CodeSize, integer.
        - Nonce, integer.
        - Storage, nested:
            - Key, keyword.
            - Value, keyword.
    - GetCodeList, keyword array: addresses of smart contracts that the transactions will load. 
    - ChangedState, nested: accounts' states the transactions changed.
        - Address, keyword.
        - Balance, keyword.
        - Nonce, integer.
        - Storage, nested:
            - Key, keyword.
            - Value, keyword.
- Rewards, nested: mining information including uncle blocks
    - Balance, keyword.
    - Beneficiary, keyword: Miner address.

###  1.2. <a name='Code'></a>Code

The code index contains bytecodes of smart contracts in blockchain, which include deploying code and runtime code.

- Hash, keyword: block hash.
- Timestamp, date.
- Transactions, nested:
    - Hash, keyword: transaction hash.
    - Input, text: deploying code of the new smart contract created by the current normal transaction.
    - TxnIndex, integer: the serial number of the current normal transaction in the current block.
    - Contracts, nested:
        - Address, keyword.
        - Code, text.
        - Hash, keyword: code hash,

###  1.3. <a name='State'></a>State

The state index records accounts' creation and self-destruction.

- Hash, keyword: block hash.
- Timestamp, date.
- Transactions, nested:
    - Hash, keyword: transaction hash.
    - TxnIndex, integer.
    - Create, keyword array: touched accounts' address.
    - Suicide, keyword array: self-destructed smart contracts' address.

##  2. <a name='Connectionmethods'></a>Connection methods

We often connect to the database through two ways. First, access the front-end link [https://kibaelastic.icsr.wiki](https://kibaelastic.icsr.wiki) and use the Dev Tools to interact with the database. Note that, accessing this link requires your host to use the intranet of ZJU, otherwise, use [RVPN](rvpn.zju.edu.cn) of ZJU. Second, use the PyPI package `elasticsearch` in Python to interact with the database, which requires your host to use the intranet of the ICSR's computer room(192.168.1.*). We often use the first way for quick search or debug and test complicated search codes, and use the second way for running program.

Kibana Dev Tools

```
Get /index name/_search
{
    ...
}
```

Python 
``` shell
$ pip install elasticsearch
```

``` python
from elasticsearch import Elasticsearch
elastic = Elasticsearch("http://192.168.1.7:9200", timeout=4000)
response = elastic.search(index="index name", body={...})
```

##  3. <a name='Searchexamples'></a>Search examples

###  3.1. <a name='Searchatransactionbytransactionhash'></a>Search a transaction by transaction hash
``` json
{
    "_source": false,
    "query": {"nested": {
        "path": "Transactions",
        "query": {"bool": {"filter": {"term": {
            "Transactions.Hash": "0x3cf41ad4f703fe61368139b8482e75de53a335b9d76039ca071530bb5292b0c7"
        }}}},
        "inner_hits": {}
    }}
}
```

Query and filter contexts: [https://www.elastic.co/guide/en/elasticsearch/reference/7.5/query-filter-context.html](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/query-filter-context.html)

Nested query: [https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-nested-query.html](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-nested-query.html)

Source filtering: [https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-source-filtering](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-source-filtering)

Inner_hits: [https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-inner-hits](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-inner-hits)

###  3.2. <a name='SearchuniqueeventsandcorrespondingtransactionshashforBZx.'></a>Search unique events and corresponding transactions' hash for BZx.

The main smart contract address of BZx is 0x1cf226e9413addaf22412a2e182f9c0de44af002.

``` json
{
    "size": 0,
    "query": {"nested": {
        "path": "Transactions",
        "query": {"nested": {
            "path": "Transactions.Logs",
            "query": {"bool": {"filter": {"term": {"Transactions.Logs.Address": "0x1cf226e9413addaf22412a2e182f9c0de44af002"}}}}
        }}
    }},
    "aggs": {"result": {"scripted_metric": {
        "init_script": "state.eventToTxns = new HashMap();",
        "map_script": """
            for (t in params['_source']['Transactions']) {
                for (log in t['Logs']) {
                    if (log['Address'] == params['Contract']) {
                        if (!state.eventToTxns.containsKey(log['Topics'][0])) {
                            state.eventToTxns[log['Topics'][0]] = []
                        }
                        state.eventToTxns[log['Topics'][0]].add(t['Hash']);
                    }
                }
            }
        """,
        "combine_script": "return state.eventToTxns",
        "reduce_script": """
            Map res = new HashMap();
            for (s in states) {
                for (event in s.keySet()) {
                    if (res.containsKey(event)) {
                        res[event].addAll(s[event]);
                    } else {
                        res[event] = s[event];
                    }
                }
            }
            return res
        """,
        "params": {
            "Contract": "0x1cf226e9413addaf22412a2e182f9c0de44af002"
        }
    }}}
}
```

From/Size: [https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-from-size](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-from-size)

Scripted_metric: [https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-aggregations-metrics-scripted-metric-aggregation.html](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-aggregations-metrics-scripted-metric-aggregation.html)

Painless scripting language: [https://www.elastic.co/guide/en/elasticsearch/painless/current/index.html](https://www.elastic.co/guide/en/elasticsearch/painless/current/index.html)
