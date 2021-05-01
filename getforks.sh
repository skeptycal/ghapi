#!/usr/bin/env zsh

gh repo list --fork --limit 100 >| forks
awk '{ print $1 }' forks >|forklist
