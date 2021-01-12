package schemas

var CodeIndexName = "code_realtime"
var CodeSchema = `
{
    "settings": {
        "number_of_shards": 2,
        "number_of_replicas": 0,
        "refresh_interval": -1,
        "index.max_inner_result_window": 10000,
        "index.mapping.nested_objects.limit": 100000
    },
    "mappings": {
        "properties": {
            "Hash": {"type": "keyword"},
            "Number": {"type": "integer"},
            "Timestamp": {"type": "date"},
            "Transactions": {
                "type": "nested",
                "properties": {
                    "Hash": {"type": "keyword"},
                    "Input": {
                        "type": "text",
                        "norms": false,
                        "index_options": "freqs"
                    },
                    "TxnIndex": {"type": "integer"},
                    "Contracts": {
                        "type": "nested",
                        "properties": {
                            "Address": {"type": "keyword"},
                            "Code": {
                                "type": "text",
                                "norms": false,
                                "index_options": "freqs"
                            },
                            "Hash": {"type": "keyword"}
                        }
                    }
                }
            }
        }
    }
}`
