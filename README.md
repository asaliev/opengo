# opengo: ChatGPT command line client

A minimalistic ChatGPT client written in Go.

## Configuration
  
There is only one required configuration parameter: `OPENAI_TOKEN`.

### Environment variables

You can set the token using an environment variable.

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
$ OPENGO_OPENAI_TOKEN=my_token ./opengo -q "Hello!"
```

### YAML Config

Alternatively you can create a config file `chatgpt.yaml` in the same directory as the executable with the following contents:

```yaml
OPENAI_TOKEN: "my_token"
``` 

## Running

todo

## License

This project is licensed under the MIT license. See LICENSE.md for more details.