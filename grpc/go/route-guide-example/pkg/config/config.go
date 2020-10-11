package config

import (
	"flag"
	"log"
)

// Conf ...
type Conf struct {
	TLSCertPath        string
	TLSKeyPath         string
	HTTPOnly           bool
	Port               string
	Host               string
	TLSCAPath          string
	ServeAddr          string
	ServerHostOverride string
}

// JSONFeature ...
type JSONFeature struct {
	JSONfeatureFile string
}

// GetConfig ...
func GetConfig() *Conf {
	c := &Conf{}
	flag.StringVar(&c.TLSCertPath, "cert-path", "./certs/route-guide-server.fabulous.af/cert.pem", "The path to the PEM-encoded TLS certificate")
	flag.StringVar(&c.TLSKeyPath, "key-path", "./certs/route-guide-server.fabulous.af/key.pem", "The path to the unencrypted TLS key")
	//flag.BoolVar(&conf.HTTPOnly, "http-only", false, "Only listen on unencrypted HTTP (e.g. for proxied environments)")
	flag.StringVar(&c.Port, "port", ":10000", "The port to listen on (HTTPS).")
	flag.StringVar(&c.TLSCAPath, "ca-path", "./certs/minica.pem", "The path to the PEM-encoded TLS miniCA certificate")
	flag.StringVar(&c.ServeAddr, "server_addr", "localhost:10000", "The server address in the format of host:port")
	flag.StringVar(&c.ServerHostOverride, "server_host_override", "route-guide-server.fabulous.af", "The server name used to verify the hostname returned by the TLS handshake")
	flag.Parse()

	log.Printf("config: %+v", c)

	return c
}

// GetJSONFeature ...
func GetJSONFeature() *JSONFeature {
	jf := &JSONFeature{}
	flag.StringVar(&jf.JSONfeatureFile, "json_db_file", "", "A json file containing a list of features")
	flag.Parse()

	return jf
}
