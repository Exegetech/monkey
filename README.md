# Monkey

TODO:
- Attach filenames and line numbers to tokens
- Support unicode
- Check code TODOs
- Type checking
- Parse scientific, binary, hexadecimal
- Power operator
- Support string, boolean, arrays, dict
- Debug flat for Pratt parsing

To Trace Pratt Parsing

Use `-1 * 2 + 3` then

```sh
go test -v -run TestOperatorPrecedenceParsing ./src/parser
```
