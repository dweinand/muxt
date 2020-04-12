#!/usr/bin/env sh

get_tmux_config() {
  echo `tmux show-option -g | grep $1 | cut -f2 -d ' ' | sed 's/"//'`
}

{{$session := .}}

tmux start-server\; has-session -t {{.Name}} 2>/dev/null

if [ "$?" -eq 1 ]; then
  {{if .Root}}cd "{{.Root}}"{{end}}
  {{if .Pre}}{{.Pre}}{{end}}
  echo "creating new session"

  TMUX= tmux new-session -d -s {{.Name}}{{if (gt (len .Window) 0)}}{{with (index .Window 0)}} -n {{.Name}}{{if .Root}} -c {{.Root}}{{end}}{{end}}{{end}}
  base_index=`get_tmux_config "base-index"`
  {{range $idx, $window := .Window}}
  echo "setting up window $(({{$idx}}+$base_index))"
  window_target={{$session.Name}}:$(({{$idx}}+$base_index))
  {{if (gt $idx 0)}}tmux new-window -n {{$window.Name}} -t {{$session.Name}}{{if $window.Root}} -c "{{.Root}}"{{end}}{{end}}
  {{if $session.PreWindow}}tmux send-keys -t $window_target "{{$session.PreWindow}}" C-m{{end}}
  {{range $cidx, $cmd := $window.Commands}}
  tmux send-keys -t $window_target "{{$cmd}}" C-m
  {{end}}
  {{range $pidx, $pane := $window.Pane}}
  echo "setting up pane $(({{$pidx}}+$base_index))"
  {{if (gt $pidx 0)}}tmux splitw{{if $window.Root}} -c {{$window.Root}}{{end}} -t $window_target{{end}}
  {{range $cidx, $cmd := $pane.Commands}}
  tmux send-keys -t $window_target.$(({{$pidx}}+$base_index)) "{{$cmd}}" C-m
  {{end}}
  {{end}}
  {{if $window.Layout}}tmux select-layout -t $window_target {{$window.Layout}}{{end}}

  {{end}}
  tmux select-window -t $base_index
fi

if [ -z "$TMUX" ]; then
  tmux -u attach-session -t {{.Name}}
else
  tmux -u switch-client -t {{.Name}}
fi
