# proch

`proxychange` makes it easy to switch wireless LAN and its proxy settings for Windows 10.

It's developed using [`systray`](https://github.com/getlantern/systray).

<a href="./doc/README_ja.md">README(日本語)</a>

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

Note: `;<local>` is optional in `"proxyOverride"`. `;` is a separator of excluded IP addresses, `<local>` is the option to check "Don't use any proxy server for addresses in intra(local) network".

2. Put the `setting.json` file in the same directory with `proch.exe`. You can put them in any directory.
   `proch.exe` is installed in `%USERPROFILE%/go/bin` if you build with `go install -ldflags -H=windowsgui`.
3. Execute `proch.exe` by double-click or execute `proch` command in terminal if you installed with `go install`. If you have changed the location of `setting.json` as described below, you can call it from anywhere in the terminal.

## How to change `setting.json` location

You can change the path of `setting.json` with any path by adding/editing the registry key.

1. Open "Registry Editor".
2. Create the registry key `HKEY_LOCAL_MACHINE\SOFTWARE\Proch\SettingJson`.
3. Add the value to this key as "string value(REG_SZ)" and set the data to the directory path where you want to place the `setting.json` file.

If the registry key is not registered or the data of the key is invalid, proch will search the directory where `proch.exe` was execute.

## Add to startup

If you want proch to be up when your computer has started up, you can add proch to startup programs.

1. Open "File Explorer".
2. Input `shell:startup` to address bar.  This will open "Startup" folder.
3. Create a shortcut to `proch.exe` in the this folder.

##  License
MIT License