# cat-aws-ssm-param

Like `cat`, but for AWS Systems Manager Parameter Store!

Writes the decrypted value of a parameter to standard output.

- No dependencies
- Single binary to download
- Uses AWS credentials from the environment

## Installation

Download `cat-aws-ssm-param`. That's it!

```sh
curl -L https://github.com/railsware/cat-aws-ssm-param/releases/download/v1.0.0/cat-aws-ssm-param.linux.gz | gzip -cd >./cat-aws-ssm-param
```

## Build for yourself. Question authority

1. Clone this repo
2. Install Go
3. Install UPX: `brew install --build-from-source upx`
4. `make build-linux` (or `build-macos`)

## Usage

```bash
./cat-aws-ssm-param /myapp/nginx_ssl_certificate >/etc/nginx/ssl.crt
```

## Use cases

Package into Docker with programs that can't access the Parameter Store directly - like nginx.

---

(c) Railsware, Leonid Shevtsov 2020
