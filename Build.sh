#!/bin/sh
watcher -cmd="sh Update.sh" -recursive -pipe=true -list ./signaling &
canthefason_watcher -run pikabu-signaling/signaling/cmd -watch pikabu-signaling
