export SHELBY_SHELL="bash"
prompt_shelby_load() {
if [ $? != 0 ]; then
    local prompt_symbol="\[\e[0;91m\]❯\[\e[0m\]"
  else
    local prompt_symbol="\[\e[0;92m\]❯\[\e[0m\]"
  fi

  PS1="$(~/.local/bin/shelby)\n${prompt_symbol} " 
}
PROMPT_COMMAND=prompt_shelby_load