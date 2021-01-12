package schemas

var GlobalIndexName = "global"
var GlobalSchema = `
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
            "HighestOrderlyBlock": {"type": "integer"}
        }
    }
}
`