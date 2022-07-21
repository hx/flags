# Flags

Remotely control devices connected to a Raspberry Pi, with some basic programmability.

An iOS client is available from the [iOS App Store](https://apps.apple.com/app/pi-flags/id1634687777?platform=iphone).

## Installation

Clone the repository:

```shell
git clone git@github.com:hx/flags
```

Build the executable:

```shell
./build.sh
```

Copy it to your Raspberry Pi:

```shell
scp out/flags-linux-arm64 pi@raspberrypi.local:flags
```

## Usage

To run the app, you need to specify:

- A state machine
- Any number of inputs
- Any number of outputs

For example:

```shell
./flags \
  dual-limits[1/2,4] \
  http[minimumsecurity@0.0.0.0:13579] \
  gpio[21,20,16,12] \
  'reset[0 3 * * *]'
```

> Note that if you're using ZSH, you'll need to quote arguments containing square brackets.

You most likely want to run this app as a service. I usually use
[this gist](https://gist.github.com/hx/e49e61d4337eed860313774935c3b68e) as a guide.

### State machines

#### `dual-limits`

A state machine that enforces a safe minimum and maximum number of flags being set, and, optionally, an unsafe minimum 
and maximum as well.

To set both safe and unsafe minimums to 1, and both safe and unsafe limits to 3:

```shell
dual-limits[1,3]
```

An unsafe minimum can be added as a prefix before a slash. An unsafe maximum can be added as a suffix after a slash.

In this example:

```shell
dual-limits[1/2,3/4]
```

- `1` is the unsafe minimum;
- `2` is the safe minimum;
- `3` is the safe maximum; and
- `4` is the unsafe maximum.

In a `dual-limits` state machine, by default, all flags up to the safe maximum are **on**.

### Inputs/outputs

#### `http`

Runs an HTTP server.

```shell
http[passphrase@bindaddr]
```

If `passphrase` is given, it should be included in all requests as the `Authorization` header.

Port 80 will be assumed if `bindaddr` does not include a port.

> HTTPS is not supported, and passphrases are transmitted in plain text. This app is designed to run on closed networks, 
> and is not secure for Internet exposure. If you need secure communication, please use a VPN, SSH tunnel, or other 
> security layer.

| Endpoint              | Description                                                                 | Example                                                          |
|-----------------------|-----------------------------------------------------------------------------|------------------------------------------------------------------|
| `GET  /flags`         | Retrieve the state of all known flags.                                      | `{"flags":[{"index":0,"state":false},{"index":0,"state":true}]}` |
| `POST /toggle/:index` | Toggle the flag at the given index. Returns the new state per `GET /flags`. | `POST /toggle/1`                                                 |

To perform an unsafe operation, include a header `X-Unsafe: 1`.

#### `gpio`

Mirrors flag state onto the given GPIO pins of a Raspberry Pi.

> Specifying this output on a device other than a Raspberry Pi will behave unpredictably, and likely panic (crash) the
> app.

This will bind the first flag to GPIO pin 21, and the second flag to GPIO pin 20:

```shell
gpio[21,20]
```

The pin will be **high** when the flag is **on**, and **low** when the flag is **off**. To reverse this behaviour, activate
inversion on the output with the letter `i`:

```shell
gpio[21,20,i]
```

Now, pins will be **low** when flags are **on** and vice-versa.

#### `reset`

Resets the state machine to its default state on the given [cron](https://crontab.guru/) schedule.

```shell
'reset[0 3 * * *]'
```

This will reset flag state at 3am every day, according to the OS's local time.

> Make sure to quote the entire argument so that your shell doesn't break it up at the spaces, or try to use the
> asterisks for globbing.
