# proch

`proxychange` makes it easy to change proxy setting for win10.

It's developed using [`systray`](https://github.com/getlantern/systray).

## Get Started

```bash
git clone https://github.com/Riki-Okunishi/proch.git
cd proch/cmd/proch
go install
```

## Usage
1. Create `setting.json`, the list of your wireless network and proxy setting.

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

2. Put `setting.json` into any folder.
3. Open this folder in terminal such as `cmd`.
4. execute the command, `proch` (if you want to execute in background, use `proch &`)