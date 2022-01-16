appsocket support for macOS

Inspired by Jose E. Marchesi's [Wait, Hyperlinks!?](https://jemarch.net/pokology-20200202.html) article about hyperlinks in [GNU Poke](https://jemarch.net/poke.html).

macOS app code is based on [Handling macOS URL schemes with Go](https://blakewilliams.me/posts/handling-macos-url-schemes-with-go).

Egmont Koblinger's [Hyperlinks in terminal emulators proposal](https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda) that iTerm2 supports.

Bruno Haible's [appsocket proposal](https://www.mail-archive.com/xdg@lists.freedesktop.org/msg08765.html).

With a combination of these two proposals a terminal program can
implement clickable text for interaction.

[hyperlink-app-client](https://gitlab.com/darnir/hyperlink-app-client) is the same for Freedesktop.

### Install and test

```sh
go build -o AppClient.app/Contents/MacOS/AppClient
```

Copy `AppClient.app` to `/Applications`.

```sh
printf '\e]8;;app://0:1234/Hello\e\\Click me\e]8;;\e\\\n'
nc -l 1234
```

Click the link and `Hello` should be written to the terminal.

