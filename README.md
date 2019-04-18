# structlog-usage-example

Usage example for Go package
[github.com/powerman/structlog](https://godoc.org/github.com/powerman/structlog).

## Output

```
main[24224] inf structlog-usage-example: `started` version 0.1.0
main[24224] inf   main: `started` version 0.1.0
main[24224] inf   main: 127.0.0.1:1234        /some/thing: `incoming request` someuser
main[24224] inf   main: 123.123.123.123:12345 /other/thing: `incoming request` someparam=value
main[24224] ERR    pkg: `failed to do something` arg=42 err: EOF
main[24224] WRN   main: `temporary error` err: EOF
main[24224] WRN   main: `temporary error` err: failed to do something: EOF
{"_a":"main","_l":"inf","_m":"started","_p":24224,"_t":"Apr 18 14:39:57.374319","_u":"main","version":"0.1.0"}
```
