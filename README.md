### TOTP Generator
---
This a very niche CLI tool which generates TOTP based on URL or QR Code. 

```
$ totp-cli get secret -h
Get secret (totp)

Usage:
  totp-cli get [flags]

Flags:
  -d, --digits int        digits for the otp (default 6)
  -f, --filepath string   path to file
  -h, --help              help for get
      --uri string        TOTP URI
```

`--uri` and `-f` flag can't be provided together. 

`--digits` define how many digits you want in the OTP.

#### NOTE:

- URI should always have `issuer` field/param.
- If image is provided, it should be `png` or `jpeg`/`jpg` format only.