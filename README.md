tmuxer
======

Outputs a formated string to stdout for use in the tmux status bar.

Currently components must be specified in main.go. In the future, command line flags will be used to specify desired components.

Usage
=======

Run `go get -u github.com/ZachMassia/tmuxer`

Add `set -g status-right "#(tmuxer) %H:%M"` to your `~/.tmux.conf`
