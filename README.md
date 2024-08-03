# One-Env-CLI

One-Env-CLI is a command-line tool that streamlines environment creation using your password manager as the provider. It provides a convenient way to manage and create environments, such as Postman environments, quickly and securely.

![Beta](https://img.shields.io/badge/status-beta-yellow)

This entire project is currently in beta. All features and functionality are subject to change.

## Prerequisites

To use this project, you need to have the following dependencies installed (more password managers and plugins are able to be developed): 

Password managers: 

1. 1Password

- [1Password](https://1password.com/)
- [1Password CLI](https://developer.1password.com/docs/cli/)

*or* 

2. BitWarden

- [BitWarden](https://bitwarden.com/)
- [BitWarden CLI](https://bitwarden.com/help/cli/)


Plugins: 
- [Postman](https://www.postman.com/)
- [Postman API Key](https://web.postman.co/settings/me/api-keys)

### Limitations

- If using BitWarden, run bw sync to pull the latest vault data from the server.
- If using BitWarden, only Custom Fields data will be passed into the plugin. 

## Installation

### Using Homebrew

You can use [Homebrew](https://brew.sh/) (on macOS or Linux) to install one-env-cli.

Run the following command to install the tool:

```
brew tap ericflores108/tap
```
```
brew install one-env-cli
```
## Usage

To use One-Env-CLI, run the following command:

```
one-env-cli [command] [flags]
```

### Available Commands

- `add`: Add a password manager item to an integrated application.
  - `postman`: Add a password manager item to create a Postman environment. The workspace is set to default. 

### postman flags

- `-i, --item`: Specify the item name to add. Currently, the item must be a unique name in your password manager. 
- `-w, --workspace`: Specify the Postman workspace name to add the environment to. 
- `-v, --verbose`: Enable verbose output.

## Configuration

One-Env-CLI uses a configuration file to store settings. The default configuration file is located in the directory, `~/.one-env-cli`, and is in JSON format.

## Demo

![Demo GIF](./images/cli.gif)

### Default Configuration

The default configuration file looks like this:

```json
{
  "plugin": {
    "postman": {
      "keyName": "PostmanAPI",
      "type": "api-key"
    },
    "gcp": {
      "type": "cli"
    }
  },
  "provider": {
    "op": {
      "vault": "Developer",
      "enabled": true
    },
    "bw": {
      "enabled": false
    }
  },
  "cli": {
    "logging": {
      "level": "debug",
      "encoding": "json",
      "outputPaths": [
        "tmp/log/one-env-cli.json"
      ]
    }
  }
}

```

You can modify the configuration file to suit your needs. The available options are:

- `plugin.postman.keyName`: The name of the 1Password item that contains the Postman API key.
- `plugin.postman.keySecretName`: The name of the secret field within the 1Password item that holds the Postman API key.
- `provider.op.vault`: The name of the 1Password vault to use.
- `provider.op.enabled`: Enable the application to use 1Password.
- `provider.bw.enabled`: Enable the application to use BitWarden.
- `cli.logging.level`: The logging level (e.g., "debug", "info", "error").
- `cli.logging.encoding`: The encoding format for the log files (e.g., "json").
- `cli.logging.outputPaths`: An array of file paths where the logs will be written.

### Custom Configuration (Work in Progress)

The `config` command for setting up custom configuration is currently under development and will be available in a future release.

## Example

Here's an example of how to use One-Env-CLI to create a Postman environment:

```
one-env-cli add postman -i Strava
```

This command will retrieve the specified item, transform it into a Postman environment, and create the environment in Postman.

All password manager secrets are retrieved and authenticated through their respective CLI applications. No passwords or data is stored, instead it passes through from one environment to another. 

## License

One-Env-CLI is licensed under the [MIT License](LICENSE).

## Contact

If you have any questions or suggestions, feel free to reach out to the maintainer:

- Eric Flores
- Email: eflorty108@gmail.com