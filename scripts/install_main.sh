# !/bin/bash
printf "\e[1;34m==> Downloading shelby for $(uname -s)\e[0m\n"
wget https://github.com/athul/shelby/releases/download/0.1.0/shelby
printf "\e[1;34m==> Installing shelby\e[0m\n"
sudo mv -f shelby /usr/local/bin/shelby
printf "\e[1;32m==> shelby successfully installed 🎉🎉 Enjoy....."