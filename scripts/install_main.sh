# !/bin/bash
OS=$(uname -s)
printf "\e[1;34m==> Downloading shelby for $(OS)\e[0m\n"
wget https://github.com/athul/shelby/releases/download/0.1.1/shelby-${OS}-amd64
printf "\e[1;34m==> Installing shelby\e[0m\n"
sudo mv -f shelby /usr/local/bin/shelby
sudo chmod 775 /usr/local/bin/shelby
printf "\e[1;32m==> shelby successfully installed 🎉🎉 Enjoy.....\n"
cat << "EOF"
 ____    __              ___    __
/\  _`\ /\ \            /\_ \  /\ \
\ \,\L\_\ \ \___      __\//\ \ \ \ \____  __  __
 \/_\__ \\ \  _ `\  /'__`\\ \ \ \ \ '__`\/\ \/\ \
   /\ \L\ \ \ \ \ \/\  __/ \_\ \_\ \ \L\ \ \ \_\ \
   \ `\____\ \_\ \_\ \____\/\____\\ \_,__/\/`____ \
    \/_____/\/_/\/_/\/____/\/____/ \/___/  `/___/> \
                                              /\___/
                                              \/__/
EOF
