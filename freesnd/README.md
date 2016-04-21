# freesnd

A command-line tool for interacting with [freesound](http://freesound.org).

# Install

```bash
go get github.com/briansorahan/freesound/freesnd
```

# Authentication

* Configure freesnd with your key and secret. As of right now, you can get them (or apply for them if you don't already have them) at [this link](http://freesound.org/apiv2/apply/). The key is "Client id" and the secret is "Client secret".

```bash
freesnd configure -key MY_FREESOUND_KEY -secret MY_FREESOUND_SECRET
```

* Get the URL for generating an auth code.

```bash
freesnd get-code
https://www.freesound.org/apiv2/oauth2/authorize?client_id=MY_FREESOUND_KEY&response_type=code
```

* Go to the URL to generate the auth code.

This will take you to a URL that looks [like this](https://raw.githubusercontent.com/briansorahan/freesound/master/freesnd/auth_code.png). Copy the generated auth code to your clipboard.

* Store the auth code.

```bash
freesnd authorize PASTE_AUTH_CODE
```
