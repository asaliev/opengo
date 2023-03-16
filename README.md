# opengo: ChatGPT command line client

A minimalistic ChatGPT client written in Go for personal use.

## Configuration

There is only one required configuration parameter which is the OpenAI API token.

You can get your token via the [OpenAI account settings](https://platform.openai.com/account/api-keys) page.

### Environment variables

#### Linux/OSX

```bash
export OPENGO_OPENAI_TOKEN=my_token
```

#### Windows

```powershell
set OPENGO_OPENAI_TOKEN=my_token
```

### Runtime

With bash-like shells you can prefix the token when running the executable:

```bash
OPENGO_OPENAI_TOKEN=my_token ./opengo -q "Hello!"
```

### YAML Config

Alternatively you can create a config file `chatgpt.yaml` in the same directory as the executable with the following contents:

```yaml
OPENAI_TOKEN: "my_token"
``` 

Note, the environment variable takes precedence over the YAML config.

## Build

```bash
go build
```

## Run

### Standard input mode

```bash
./opengo -q "To be or not to be?"
```

### Interactive mode

```bash
$ ./opengo
Query (`exit` to quit):
to be or not to be?
...
```

## Test

```bash
go test . ./config
```

## License

This project is licensed under the MIT license. See LICENSE for more details.