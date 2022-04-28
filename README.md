# duration-to-sec
Simple Go program that receives a string duration from stdin and outputs that in seconds, milliseconds, microseconds, nanoseconds

## Build
go build -o duration-to-sec main.go

## Run examples

Add `-precision int` to apply a multiplier to the default seconds output.
For milliseconds, add `-precision 1000`.
For microseconds, add `-precision 1000000`.
And so on.

```shell
# echo "1h" | duration-to-sec
3600
# echo "58m10s" | duration-to-sec
3490
# echo "1m" | duration-to-sec -precision 1000
60000
```
