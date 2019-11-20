package rules

import (
	"go/ast"

	"github.com/securego/gosec"
)

// NewModernTLSCheck creates a check for Modern TLS ciphers
// DO NOT EDIT - generated by tlsconfig tool
func NewModernTLSCheck(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	return &insecureConfigTLS{
		MetaData:     gosec.MetaData{ID: id},
		requiredType: "crypto/tls.Config",
		MinVersion:   0x0304,
		MaxVersion:   0x0304,
		goodCiphers: []string{
			"TLS_AES_128_GCM_SHA256",
			"TLS_AES_256_GCM_SHA384",
			"TLS_CHACHA20_POLY1305_SHA256",
		},
	}, []ast.Node{(*ast.CompositeLit)(nil)}
}

// NewIntermediateTLSCheck creates a check for Intermediate TLS ciphers
// DO NOT EDIT - generated by tlsconfig tool
func NewIntermediateTLSCheck(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	return &insecureConfigTLS{
		MetaData:     gosec.MetaData{ID: id},
		requiredType: "crypto/tls.Config",
		MinVersion:   0x0303,
		MaxVersion:   0x0304,
		goodCiphers: []string{
			"TLS_AES_128_GCM_SHA256",
			"TLS_AES_256_GCM_SHA384",
			"TLS_CHACHA20_POLY1305_SHA256",
			"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
			"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
			"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
			"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
			"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305",
			"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305",
			"TLS_DHE_RSA_WITH_AES_128_GCM_SHA256",
			"TLS_DHE_RSA_WITH_AES_256_GCM_SHA384",
		},
	}, []ast.Node{(*ast.CompositeLit)(nil)}
}

// NewOldTLSCheck creates a check for Old TLS ciphers
// DO NOT EDIT - generated by tlsconfig tool
func NewOldTLSCheck(id string, conf gosec.Config) (gosec.Rule, []ast.Node) {
	return &insecureConfigTLS{
		MetaData:     gosec.MetaData{ID: id},
		requiredType: "crypto/tls.Config",
		MinVersion:   0x0301,
		MaxVersion:   0x0304,
		goodCiphers: []string{
			"TLS_AES_128_GCM_SHA256",
			"TLS_AES_256_GCM_SHA384",
			"TLS_CHACHA20_POLY1305_SHA256",
			"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
			"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
			"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
			"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
			"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305",
			"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305",
			"TLS_DHE_RSA_WITH_AES_128_GCM_SHA256",
			"TLS_DHE_RSA_WITH_AES_256_GCM_SHA384",
			"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256",
			"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
			"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
			"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
			"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384",
			"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384",
			"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
			"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
			"TLS_DHE_RSA_WITH_AES_128_CBC_SHA256",
			"TLS_DHE_RSA_WITH_AES_256_CBC_SHA256",
			"TLS_RSA_WITH_AES_128_GCM_SHA256",
			"TLS_RSA_WITH_AES_256_GCM_SHA384",
			"TLS_RSA_WITH_AES_128_CBC_SHA256",
			"TLS_RSA_WITH_AES_256_CBC_SHA256",
			"TLS_RSA_WITH_AES_128_CBC_SHA",
			"TLS_RSA_WITH_AES_256_CBC_SHA",
			"TLS_RSA_WITH_3DES_EDE_CBC_SHA",
		},
	}, []ast.Node{(*ast.CompositeLit)(nil)}
}
