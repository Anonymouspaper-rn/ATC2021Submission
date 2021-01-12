package schemas

var StateIndexName = "state_realtime"
var StateSchema = `
{
    "settings": {
        "number_of_shards": 2,
        "number_of_replicas": 0,
        "refresh_interval": -1,
        "index.max_inner_result_window": 20000,
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
                    "Create": {"type": "keyword"},
                    "Hash": {"type": "keyword"},
                    "Reset": {"type": "keyword"},
                    "Suicide": {"type": "keyword"},
                    "TxnIndex": {"type": "integer"}
                }
            }
        }
    }
}`
