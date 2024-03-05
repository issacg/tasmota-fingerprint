package main

/*
  Tasmota Fingerprint Generator - Used to calculate TLS fingerprints
  expected by Tasmota firmware

  Copyright (c) 2019 Issac Goldstand <margol@beamartyr.net>

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
*/

import (
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func reverse(ar []byte) []byte {
	for i := 0; i < len(ar)/2; i++ {
		j := len(ar) - i - 1
		ar[i], ar[j] = ar[j], ar[i]
	}
	return ar
}

func main() {
	var data []byte
	var err error
	if len(os.Args) > 1 && os.Args[1] != "-" {
		data, err = ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
	} else {
		info, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		if info.Mode()&os.ModeCharDevice != 0 {
			panic("Missing input on STDIN")
		}
		data, err = ioutil.ReadAll(os.Stdin)
	}

	block, _ := pem.Decode(data)
	if block == nil {
		panic("Error reading PEM data")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(err)
	}

	pub := cert.PublicKey.(*rsa.PublicKey)
	buf := make([]byte, binary.MaxVarintLen64)
	binary.LittleEndian.PutUint64(buf, uint64(pub.E))
	buf = reverse(buf)
	start := 0
	for start < len(buf) && buf[start] == 0 {
		start++
	}

	e_bytes := buf[start:]
	e_length := make([]byte, 4)
	e_length[0] = byte(len(e_bytes) >> 24 & 255)
	e_length[1] = byte(len(e_bytes) >> 16 & 255)
	e_length[2] = byte(len(e_bytes) >>  8 & 255)
	e_length[3] = byte(len(e_bytes) >>  0 & 255)

	n_bytes := pub.N.Bytes()
	n_length := make([]byte, 4)
	n_length[0] = byte(len(n_bytes) >> 24 & 255)
	n_length[1] = byte(len(n_bytes) >> 16 & 255)
	n_length[2] = byte(len(n_bytes) >>  8 & 255)
	n_length[3] = byte(len(n_bytes) >>  0 & 255)

	ctx := sha1.New()
	ctx.Write([]byte("\000\000\000\007")) // length of "ssh-rsa"
	ctx.Write([]byte("ssh-rsa"))
	ctx.Write(e_length)
	ctx.Write(e_bytes)
	ctx.Write(n_length)
	ctx.Write(n_bytes)
	fmt.Printf("% X (Tasmota v8.4.0+)\n", ctx.Sum(nil))
}
