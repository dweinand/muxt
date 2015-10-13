# muxt

A commandline tool for setting up tmux environments.

## Usage

Muxt uses [TOML][toml]-based config files for defining your tmux sessions.

If your config file is placed in `~/.muxt`, you can start the session by simply
passing the base name of the file to the muxt command.

```bash
muxt [name]
```

Otherwise, you can specify an arbitrary path to a config file to use.

```bash
muxt path/to/config.toml
```

## Configuration

Muxt's configuration is based on [TOML v.0.4.0][toml]. See an [example][example]
for more infromation.

[toml]:https://github.com/toml-lang/toml/blob/master/versions/en/toml-v0.4.0.md
[example]:https://github.com/dweinand/muxt/blob/master/src/muxt/assets/config/example.toml
