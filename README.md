# Rotates pen input coordinates with laptop screen

This small app tries to fix tablet/laptop pen rotation issues when used with gnome desktop.

For now only works for HP Spectre 360 (kaby lake).

## Build and install

```bash
$ make

$ sudo make install # places built binary to /usr/local/bin/laptop-rotate-fix
```

Autostart this app with desktop environment.

```bash
$ cp laptop-rotation-fix.desktop $HOME/.config/autostart/
$ chmod a+x $HOME/.config/autostart/
```

## License

MIT