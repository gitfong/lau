#!/bin/sh

ps -ef | grep lau | grep -v grep | awk '{print $2}' | xargs kill