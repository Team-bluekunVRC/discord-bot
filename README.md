# discord-bot

Source code of the discord bot in the Just B Club discord server

## Developing

To hack on this repository, ensure you have the following:

* `yq` - `pip3 install yq` (one day we might use the go version, if it was better)
* `docker` -  [Ubuntu](https://docs.docker.com/engine/install/ubuntu/), [WSL2](https://hub.docker.com/editions/community/docker-ce-desktop-windows), [macOS](https://docs.docker.com/desktop/mac/install/)
* `asdf` - [all platforms](https://asdf-vm.com/guide/getting-started.html#_1-install-dependencies), setup plugins with: `for plugin in "golang" "jq"; do asdf plugin-add "$plugin"; done`

Ensure that you have the right host dependencies:

```bash
asdf install
```

You are ready to go! You can easily build a local version with `make`. Updated dependencies with `make dep` and to automatically reload changes, run `make watch`

## License

GPL-3.0
