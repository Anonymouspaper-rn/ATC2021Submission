package schemas

var BlockIndexName = "block_realtime"
var BlockSchema = `
{
    "settings": {
        "number_of_shards": 88,
        "number_of_replicas": 0,
        "refresh_interval": -1,
        "index.max_result_window": 1000000,
        "index.max_inner_result_window": 20000,
        "index.mapping.nested_objects.limit": 100000
    },
    "mappings": {
        "properties": {
            "Difficulty": {"type": "long"},
            "ExtraData": {
                "type": "text",
                "index": false
            },
            "GasLimit": {"type": "long"},
            "GasUsed": {"type": "long"},
            "Hash": {"type": "keyword"},
            "Miner": {"type": "keyword"},
            "Number": {"type": "integer"},
            "Timestamp": {"type": "date"},
            "TxnCount": {"type": "integer"},
            "Transactions": {
                "type": "nested",
                "properties": {
                    "CallFunction": {"type": "keyword"},
                    "CallParameter": {
                        "type": "text",
                        "norms": false,
                        "index_options": "freqs"
                    },
                    "ConAddress": {"type": "keyword"},
                    "CumGasUsed": {
                        "type": "long",
                        "index": false
                    },
                    "FromAddress": {"type": "keyword"},
                    "GasLimit": {"type": "long"},
                    "GasPrice": {"type": "keyword"},
                    "GasUsed": {"type": "long"},
                    "Hash": {"type": "keyword"},
                    "IntTxnCount": {"type": "integer"},
                    "Nonce": {"type": "integer"},
                    "Status": {"type": "boolean"},
                    "ToAddress": {"type": "keyword"},
                    "TxnIndex": {"type": "integer"},
                    "Value": {"type": "keyword"},
                    "InternalTxns": {
                        "type": "nested",
                        "properties": {
                            "CallFunction": {"type": "keyword"},
                            "CallParameter": {
                                "type": "text",
                                "norms": false,
                                "index_options": "freqs"
                            },
                            "ConAddress": {"type": "keyword"},
                            "EvmDepth": {"type": "short"},
                            "FromAddress": {"type": "keyword"},
                            "GasLimit": {"type": "long"},
                            "Output": {
                                "type": "text",
                                "norms": false,
                                "index_options": "freqs"
                            },
                            "ToAddress": {"type": "keyword"},
                            "TxnIndex": {"type": "integer"},
                            "Type": {"type": "short"},
                            "Value": {"type": "keyword"}
                        }
                    },
                    "Logs": {
                        "type": "nested",
                        "properties": {
                            "Address": {"type": "keyword"},
                            "Data": {
                                "type": "text",
                                "norms": false,
                                "index_options": "freqs"
                            },
							"IntTxnIndex": {"type": "integer"},
                            "Topics": {
                                "type": "keyword",
                                "doc_values": false
                            }
                        }
                    },
                    "ReadCommittedState": {
                        "type": "nested",
                        "properties": {
                            "Address": {
                                "type": "keyword",
                                "doc_values": false
                            },
                            "Balance": {
                                "type": "keyword",
                                "doc_values": false
                            },
                            "CodeHash": {
                                "type": "keyword",
                                "doc_values": false
                            },
                            "CodeSize": {
                                "type": "integer",
                                "doc_values": false
                            },
                            "Nonce": {
                                "type": "integer",
                                "doc_values": false
                            },
                            "Storage": {
                                "type": "nested",
                                "properties": {
                                    "Key": {
                                        "type": "keyword",
                                        "doc_values": false
                                    },
                                    "Value": {
                                        "type": "keyword",
                                        "doc_values": false
                                    }
                                }
                            }
                        }
                    },
                    "ChangedState": {
                        "type": "nested",
                        "properties": {
                            "Address": {
                                "type": "keyword",
                                "doc_values": false
                            },
                            "Balance": {
                                "type": "keyword",
                                "doc_values": false
                            },
                            "Nonce": {
                                "type": "integer",
                                "doc_values": false
                            },
                            "Storage": {
                                "type": "nested",
                                "properties": {
                                    "Key": {
                                        "type": "keyword",
                                        "doc_values": false
                                    },
                                    "Value": {
                                        "type": "keyword",
                                        "doc_values": false
                                    }
                                }
                            }
                        }
                    },
                    "GetCodeList": {
                        "type": "keyword",
                        "doc_values": false
                    }
                }
            },
            "Rewards": {
                "type": "nested",
                "properties": {
                    "Balance": {"type": "keyword"},
                    "Beneficiary": {
                        "type": "keyword",
                        "doc_values": false
                    }
                }
            }
        }
    }
}`
