# muxt

[![CircleCI](https://circleci.com/gh/dweinand/muxt.svg?style=svg&circle-token=d2dda66af3767333a8c0193974fd53b7df1ca619)]



A commandline tool for setting up tmux environments.

## Installation

If you use homebrew, installation is as simple as:

```shell
brew tap dweinand/utilities
brew install muxt
```
Otherwise you can download a binary for the [latest release][latest-release].

## Usage

Muxt uses [TOML][toml]-based config files for defining your tmux sessions.

If your config file is placed in `~/.muxt`, you can start the session by simply
passing the base name of the file to the muxt command.

```shell
muxt [name]
```

Otherwise, you can specify an arbitrary path to a config file to use.

```shell
muxt path/to/config.toml
```

## Configuration

Muxt's configuration is based on [TOML v.0.4.0][toml]. See an [example][example]
for more infromation.

## Development

Muxt uses [gb][gb] and you should install it before building or running tests.

To run tests:

```shell
make test
```

To build:

```shell
make build
```

## Contributing

Bug reports and pull requests are welcome on GitHub at
https://github.com/dweinand/muxt.

This project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [Contributor Covenant][coc] code of
conduct.

## Acknowledgments

Muxt is inspired by similar tools:
* [teamocil](https://github.com/remiprev/teamocil)
* [tmuxinator](https://github.com/tmuxinator/tmuxinator)

## License

Muxt is available as open source under the terms of the [MIT License][mit].

[latest-release]:https://github.com/dweinand/muxt/releases/latest
[toml]:https://github.com/toml-lang/toml/blob/master/versions/en/toml-v0.4.0.md
[example]:https://github.com/dweinand/muxt/blob/master/src/muxt/assets/config/example.toml
[gb]:http://getgb.io/
[coc]:http://contributor-covenant.org
[mit]:http://opensource.org/licenses/MIT
