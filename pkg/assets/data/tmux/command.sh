#!/usr/bin/env sh

tmux start-server
tmux has-session -t {{.Name}} 2>/dev/null

if [ "$?" -eq 1 ]; then
  cd {{.Root}}

  tmux new-session -d -s {{.Name}}{{with (index .Window 0)}} -n {{.Name}}{{if .Root}} -c {{.Root}}{{end}}{{end}}
fi

tmux -u attach-session -t {{.Name}}
