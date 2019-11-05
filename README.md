# ipoib-core-dump

When try to create a link from InfiniBand net device, a core dump issue occur

## Reproducing Core Dump Issue

Run the following commands
```
# ip netns add testing
# go run coreDumpReproduce.go

```