# proch

`proxychange` makes it easy to change proxy setting for Windows 10.

It's developed using [`systray`](https://github.com/getlantern/systray).

## How to Build

```bash
git clone https://github.com/Riki-Okunishi/proch.git
cd proch/cmd/proch
go install -ldflags -H=windowsgui
```

## Usage
1. Create `setting.json` shown below. You list up your wireless network and proxy setting in this file.
   If the value of "proxyEnable" is true for a given network, then you should specify the values of "proxyServer" and "proxyOverride".

```json: setting.json
{
  "profiles": [
    {
      "ssid": "Proxy SSID",
      "proxyEnable": true,
      "proxyServer": "proxy.address:PORT",
      "proxyOverride": "exclude address[;<local>]"
    },
    {
      "ssid": "non-Proxy SSID",
      "proxyEnable": false
    }
  ]
}
```

2. Put the `setting.json` file in the same directory with `proch.exe`.
   `proch.exe` is installed in `%USERPROFILE%/go/bin` if you build with `go install`.
3. Execute `proch.exe` by double-click or execute `proch` command in terminal if you installed with `go install`.


## Add to startup

If you want proch to be up when your computer has started up, you can add proch to startup programs.

### Windows

1. Open "File Explorer".
2. Input `shell:startup` to address bar.  This will open "Startup" folder.
3. Create a shortcut to `proch.exe` in the "Startup" folder.

##  License
MIT License