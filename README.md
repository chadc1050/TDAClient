# TD Ameritrade API Go Client

![TDA Client Release](https://img.shields.io/github/v/release/chadc1050/TDAClient)
[![Go Source Version](https://img.shields.io/github/go-mod/go-version/chadc1050/TDAClient/master)](https://go.dev/doc/devel/release#go1.19)

An unofficial TD Ameritrade Go Client for TD Ameritrade's Trading API.

### [TD Ameritrade Docs](https://developer.tdameritrade.com/)

## Authentication
TDA Ameritrade uses OAuth with a brokerage accounts credentials to generate the code needed to make requests on behalf
of that account to do this you will need to set up a callback url in
your [TD Ameritrade Developer Portal](https://developer.tdameritrade.com/user/me/apps)
so that TD Ameritrade can send you a refresh token when Authentication is completed using their web-based form.
Details on to do this can be found in their [Authentication FAQ Page](https://developer.tdameritrade.com/content/authentication-faq).


## Usage
To add TDAClient module to your project, add the following to your ```go.mod``` file:

```
require github.com/chadc1050/TDAClient latest
```

Then run ```github.com/chadc1050/TDAClient@latest``` command to pull the module into your project.