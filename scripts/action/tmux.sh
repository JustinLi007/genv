#!/usr/bin/env bash

PROGNAME=$0
CWD="${1:-$PWD}"
echo "'$PROGNAME': cwd: $CWD"
SESSION="${CWD##*/}"
echo "'$PROGNAME': session: $SESSION"
if [[ "$SESSION" =~ ^\. ]]; then
  echo "'$PROGNAME': '$SESSION' is a hidden dir"
  echo "'$PROGNAME': new name = dot_${SESSION##.}"
  SESSION=dot_${SESSION##.}
fi
echo "'$PROGNAME': session: $SESSION"

if [[ -e "$CWD" ]]; then
  echo "'$PROGNAME': '$CWD' exists"
else
  echo "'$PROGNAME': '$CWD' does not exists" >&2
  exit 1
fi

if [[ -d "$CWD" ]]; then
  echo "'$PROGNAME': '$CWD' is a directory"
else
  echo "'$PROGNAME': '$CWD' is not a directory" >&2
  exit 1
fi

tty=false
if [[ -t 1 ]]; then
  tty=true
fi

in_tmux=false
if [[ -n "$TMUX" ]]; then
  in_tmux=true
fi

has_session=false
if tmux has-session -t "$SESSION" 2>/dev/null; then
  has_session=true
  echo "'$PROGNAME': '$SESSION'  exist"
fi

has_git=false
if stat $"$CWD/.git" 1>/dev/null 2>&1; then
  has_git=true
fi

attach_or_switch() {
  if [[ $tty == true ]]; then
    echo "'$PROGNAME': tty"
    if [[ $in_tmux == true ]]; then
      echo "'$PROGNAME': inside tmux, switch session '$SESSION'"
      tmux switch-client -t "$SESSION"
    else
      echo "'$PROGNAME': outside tmux, attach to session '$SESSION'"
      tmux a -t "$SESSION"
    fi
  else
    echo "'$PROGNAME': no tty"
  fi
  return
}

create_session() {
  if [[ $has_git == true ]]; then
    echo "'$PROGNAME': create new session '$SESSION', has git dir"
    tmux new -d -s "$SESSION" -c "$CWD" -nGit -- lazygit
  else
    echo "'$PROGNAME': create new session '$SESSION', no git dir"
    tmux new -d -s "$SESSION" -c "$CWD" -nGit
  fi
  return
}

create_windows() {
  echo "'$PROGNAME': create windows for session '$SESSION'"
  tmux neww -t "$SESSION" -c "$CWD" -nSrc -- nvim .
  # tmux neww -t "$SESSION" -nMisc
  tmux neww -t "$SESSION" -c "$CWD"
  tmux neww -t "$SESSION" -c "$CWD"
  tmux select-window -t "$SESSION:1"
  return
}

if [[ $has_session == true ]]; then
  echo "'$PROGNAME': session '$SESSION' exist"
  attach_or_switch
else
  echo "'$PROGNAME': session '$SESSION' does not exist"
  create_session
  create_windows
  attach_or_switch
fi
