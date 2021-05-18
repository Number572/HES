package main

import (
	"crypto/rsa"
	"net"
	"sync"
)

// Basic structure describing the user.
// Stores the private key and list of friends.
type Client struct {
	handle      func(*Client, *Package)
	mutex       sync.Mutex
	privateKey  *rsa.PrivateKey
	mapping     map[string]bool
	connections map[string]net.Conn
	actions     map[string]chan []byte
	F2F         *friendToFriend
}

type friendToFriend struct {
	mutex   sync.Mutex
	enabled bool
	friends map[string]*rsa.PublicKey
}

type Package struct {
	Head HeadPackage `json:"head"`
	Body BodyPackage `json:"body"`
}

type HeadPackage struct {
	Title   string `json:"title"`
	Rand    []byte `json:"rand"`
	Sender  []byte `json:"sender"`
	Session []byte `json:"session"`
}

type BodyPackage struct {
	Data []byte `json:"data"`
	Hash []byte `json:"hash"`
	Sign []byte `json:"sign"`
	Npow uint64 `json:"npow"`
}
