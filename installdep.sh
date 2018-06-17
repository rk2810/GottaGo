curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
echo 'alias dep="$GOPATH/bin/dep"' >> ~/.bash_aliases
source ~/.bash_aliases
echo "Done."