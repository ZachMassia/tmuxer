tmuxer
======

Outputs a formated string to stdout for use in the tmux status bar.

Components are specified as command line arguments, the are printed out in the order passes in.

Usage
=======

Run `go get -u github.com/ZachMassia/tmuxer`

Add `set -g status-right "#(tmuxer bat)"` to your `~/.tmux.conf`
