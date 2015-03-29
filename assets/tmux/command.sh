#!/usr/bin/env sh

get_tmux_config() {
  echo `echo $tmux_config | grep $1 | cut -f2 -d ' ' | sed 's/"//'`
}

init_tmux_commands() {
  echo "starting tmux server"
  echo "getting tmux config"
  tmux_config=`tmux start-server\; show-option -g`;

  {{$session := .}}
  tmux start-server\; has-session -t {{.Name}} 2>/dev/null

  if [ "$?" -eq 1 ]; then
    base_index=`get_tmux_config "base-index"`

    TMUX=`tmux new-session -d -s {{.Name}}{{if (gt (len .Window) 0)}}{{with (index .Window 0)}} -n {{.Name}}{{if .Root}} -c {{.Root}}{{end}}{{end}}{{end}}`
  {{range $idx, $window := .Window}}
    echo "setting up window $(({{$idx}}+$base_index))"
    window_target={{$session.Name}}:$(({{$idx}}+$base_index))
    {{if (gt $idx 0)}}tmux new-window -n {{$window.Name}}{{if $window.Root}} -c {{.Root}}{{end}}{{end}}
    {{if $session.PreWindow}}tmux send-keys -t $window_target "{{$session.PreWindow}}" C-m{{end}}
    {{if $window.Command}}tmux send-keys -t $window_target "{{$window.Command}}" C-m{{end}}
    {{range $pidx, $pane := $window.Pane}}
    echo "setting up pane $(({{$pidx}}+$base_index))"
      {{if (gt $pidx 0)}}tmux splitw{{if $window.Root}} -c {{$window.Root}}{{end}} -t $window_target{{end}}
      {{if $pane.Command}}tmux send-keys -t $window_target.$(({{$pidx}}+$base_index)) "{{$pane.Command}}" C-m{{end}}
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
}

cd {{.Root}}
{{if .Pre}}{{.Pre}}{{end}}
init_tmux_commands > /dev/null
