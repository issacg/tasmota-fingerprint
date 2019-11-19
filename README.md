# tasmota-fingerprint
## Synopsis
tasmota-fingerprint is a tool designed to assist in calculating the custom TLS fingerprints expected by the [Tasmota](https://github.com/arendst/Tasmota) firmware since version 6.5.0.15 as documented on the [Tasmota Wiki](https://github.com/arendst/Tasmota/wiki/TLS#fingerprint-validation)

## Usage
### Reading from a file
`tasmota-fingerprint /path/to/server.pem`
### Reading from a pipe
`/some/command/outputservercertificate | tasmota-fingerprint`

In either case, the input is expected to be a PEM-encoded x509 certificate with an RSA public key, as is required by the Tasmota project. 

## Installation
If you have [Go](https://golang.org) installed, simply `go get github.com/issacg/tasmota-fingerprint`

Otherwise, [binary releases](https://github.com/issacg/tasmota-fingerprint/releases/latest) may be available.

## Rationale
The current implementation of TLS in the Tasmota project, without significantly altering the source code, requires either the use of a certificate issued by [LetsEncrypt](https://letsencrypt.org), or a [TOFU](https://en.wikipedia.org/wiki/Trust_on_first_use) security model.  While these are both legitimate models for many use cases, they don't cover the use cases available in previous versions of the firmware where it was trivial to pre-calculate the expected fingerprint of any key.  

This software allows the expected fingerprint to be calculated in order to be configured on end-devices without needing a TOFU flow.

## License
Copyright 2019 Issac Goldstand <margol@beamartyr.net>

  This library is free software; you can redistribute it and/or
  modify it under the terms of the GNU Lesser General Public
  License as published by the Free Software Foundation; either
  version 2.1 of the License, or (at your option) any later version.
  This library is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
  Lesser General Public License for more details.
  You should have received a copy of the GNU Lesser General Public
  License along with this library; if not, write to the Free Software
  Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA

Having said that, the author of this software is interested in relicensing under the more permissive Apache 2.0 license.  If the original author of [this function](https://github.com/arendst/Tasmota/blob/6ec43737616d56db310d64a7bb7e856e1f62f053/tasmota/WiFiClientSecureLightBearSSL.cpp#L695) reads this notice and is willing to allow such relicensing, (s)he is encouraged to open an issue on GitHub or otherwise contact the author.  Thank you!
