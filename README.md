# rd-client

Download torrents at high speeds using rd-client. It uses Real-Debrid and aria2c. It can monitor system clipboard to capture magnet links and automatically 
gets the real-debrid download link. It then passes them to aria2c for high speed download.

rd-client only works on Wayland as of now as it uses wl-clipboard.

## Usage

Run aria2c before running the app using: 
``` bash
aria2c --enable-rpc-input --rpc-secret="secret"
```

