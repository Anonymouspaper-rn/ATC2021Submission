
## APIs to retrieve execution context
| | | | | | | |
|-|-|-|-|-|-|-|
| op.getN()      | stack.length() | memory.slice(start, end) | contract.getSelfAddress() | getBalance(addr) | getBlockNumber() | getPc()|
| op.toNumber()  | stack.peek(n)  | memory.getUint(offset)   | contract.getCodeAddress() | getNonce(addr)   | getTxnIndex()    | getGas()        |
| op.toString()  |                |                          | contract.getValue()       | getCode(addr)    | getTxnHash()     | getDepth()      |
|                |                |                          | contract.getInput()       | getStorage(addr) |                  | getReturnData() |

## APIs to assign, clear and check taint tags
| | | | | | 
|-|-|-|-|-|
| labelStack(n,tag) | labelMemory(offset,size,tag) | labelInput(o,s,t)   | labelReturnData(o,s,t)   | labelStorage(addr,slot,tag) |
| clearStack(n)     | clearMemory(offset,size)     | clearInput(o,s)     | clearReturnData(o,s)     | clearStorage(addr,slot)     |
| peekStack(n)      | peekMemory(offset)           | peekInput(o)        | peekReturnData(o)        | peekStorage(addr,slot)      |
|                   | peekMemorySlice(offset,size) | peekInputSlice(o,s) | peekReturnDataSlice(o,s) |                             |

## Other APIs

| | |
|-|-|
| cfg.hijack(isJump) | params.get(key) |

## Example

``` javascript
{
    sload: function(log){
        contextContract = toHex(log.contract.getSelfAddress())
        key = log.stack.peek(0).toString(16)
        tag = contract+"_"+key
        log.taint.labelStack(0, tag)
    },
    
    jumpi: function(log) {
        tags = log.taint.peekStack(1)
        for (tag in tags) {
            contextContract = tag.substring(0, tag.indexOf("_"))
            key = tag.substring(tag.indexOf("_"))
            console.log("Storage", key, "in contract", contextContract, "influenced the control flow.")
        }
    }
}
```