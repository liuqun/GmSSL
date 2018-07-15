/* +build cgo */
package gmssl

/*
#include <openssl/evp.h>
*/
import "C"

import (
	"errors"
	"unsafe"
	"runtime"
)

/* generated by `gmssl list -cipher-algorithms | sort -f | uniq -i` */
func GetCipherNames() []string {
	return []string{
		"AES-128-CBC",
		"AES-128-CBC-HMAC-SHA1",
		"AES-128-CBC-HMAC-SHA256",
		"AES-128-CFB",
		"AES-128-CFB1",
		"AES-128-CFB8",
		"AES-128-CTR",
		"AES-128-ECB",
		"AES-128-OCB",
		"AES-128-OFB",
		"AES-128-XTS",
		"AES-192-CBC",
		"AES-192-CFB",
		"AES-192-CFB1",
		"AES-192-CFB8",
		"AES-192-CTR",
		"AES-192-ECB",
		"AES-192-OCB",
		"AES-192-OFB",
		"AES-256-CBC",
		"AES-256-CBC-HMAC-SHA1",
		"AES-256-CBC-HMAC-SHA256",
		"AES-256-CFB",
		"AES-256-CFB1",
		"AES-256-CFB8",
		"AES-256-CTR",
		"AES-256-ECB",
		"AES-256-OCB",
		"AES-256-OFB",
		"AES-256-XTS",
		"AES128",
		"aes128-wrap",
		"AES192",
		"aes192-wrap",
		"AES256",
		"aes256-wrap",
		"BF",
		"BF-CBC",
		"BF-CFB",
		"BF-ECB",
		"BF-OFB",
		"blowfish",
		"CAMELLIA-128-CBC",
		"CAMELLIA-128-CFB",
		"CAMELLIA-128-CFB1",
		"CAMELLIA-128-CFB8",
		"CAMELLIA-128-CTR",
		"CAMELLIA-128-ECB",
		"CAMELLIA-128-OFB",
		"CAMELLIA-192-CBC",
		"CAMELLIA-192-CFB",
		"CAMELLIA-192-CFB1",
		"CAMELLIA-192-CFB8",
		"CAMELLIA-192-CTR",
		"CAMELLIA-192-ECB",
		"CAMELLIA-192-OFB",
		"CAMELLIA-256-CBC",
		"CAMELLIA-256-CFB",
		"CAMELLIA-256-CFB1",
		"CAMELLIA-256-CFB8",
		"CAMELLIA-256-CTR",
		"CAMELLIA-256-ECB",
		"CAMELLIA-256-OFB",
		"CAMELLIA128",
		"CAMELLIA192",
		"CAMELLIA256",
		"CAST",
		"CAST-cbc",
		"CAST5-CBC",
		"CAST5-CFB",
		"CAST5-ECB",
		"CAST5-OFB",
		"ChaCha20",
		"ChaCha20-Poly1305",
		"DES",
		"DES-CBC",
		"DES-CFB",
		"DES-CFB1",
		"DES-CFB8",
		"DES-ECB",
		"DES-EDE",
		"DES-EDE-CBC",
		"DES-EDE-CFB",
		"DES-EDE-ECB",
		"DES-EDE-OFB",
		"DES-EDE3",
		"DES-EDE3-CBC",
		"DES-EDE3-CFB",
		"DES-EDE3-CFB1",
		"DES-EDE3-CFB8",
		"DES-EDE3-ECB",
		"DES-EDE3-OFB",
		"DES-OFB",
		"DES3",
		"des3-wrap",
		"DESX",
		"DESX-CBC",
		"id-aes128-CCM",
		"id-aes128-GCM",
		"id-aes128-wrap",
		"id-aes128-wrap-pad",
		"id-aes192-CCM",
		"id-aes192-GCM",
		"id-aes192-wrap",
		"id-aes192-wrap-pad",
		"id-aes256-CCM",
		"id-aes256-GCM",
		"id-aes256-wrap",
		"id-aes256-wrap-pad",
		"id-smime-alg-CMS3DESwrap",
		"IDEA",
		"IDEA-CBC",
		"IDEA-CFB",
		"IDEA-ECB",
		"IDEA-OFB",
		"RC2",
		"rc2-128",
		"rc2-40",
		"RC2-40-CBC",
		"rc2-64",
		"RC2-64-CBC",
		"RC2-CBC",
		"RC2-CFB",
		"RC2-ECB",
		"RC2-OFB",
		"RC4",
		"RC4-40",
		"RC4-HMAC-MD5",
		"SEED",
		"SEED-CBC",
		"SEED-CFB",
		"SEED-ECB",
		"SEED-OFB",
		"SMS4",
		"SMS4-CBC",
		"SMS4-CCM",
		"SMS4-CFB",
		"SMS4-CFB1",
		"SMS4-CFB8",
		"SMS4-CTR",
		"SMS4-ECB",
		"SMS4-GCM",
		"SMS4-OCB",
		"SMS4-OFB",
		"SMS4-WRAP",
		"SMS4-WRAP-PAD",
		"SMS4-XTS",
	}
}

func GetCipherKeyLength(name string) (int, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cipher := C.EVP_get_cipherbyname(cname)
	if cipher == nil {
		return 0, GetErrors()
	}
	return int(C.EVP_CIPHER_key_length(cipher)), nil
}

func GetCipherBlockLength(name string) (int, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cipher := C.EVP_get_cipherbyname(cname)
	if cipher == nil {
		return 0, GetErrors()
	}
	return int(C.EVP_CIPHER_block_size(cipher)), nil
}

func GetCipherIVLength(name string) (int, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cipher := C.EVP_get_cipherbyname(cname)
	if cipher == nil {
		return 0, GetErrors()
	}
	return int(C.EVP_CIPHER_iv_length(cipher)), nil
}

type CipherContext struct {
	ctx *C.EVP_CIPHER_CTX
}

func NewCipherContext(name string, eng *Engine, key, iv []byte, encrypt bool) (
	*CipherContext, error) {

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cipher := C.EVP_get_cipherbyname(cname)
	if cipher == nil {
		return nil, GetErrors()
	}

	if key == nil {
		return nil, errors.New("No key")
	}
	if len(key) != int(C.EVP_CIPHER_key_length(cipher)) {
		return nil, errors.New("Invalid key length")
	}

	if 0 != int(C.EVP_CIPHER_iv_length(cipher)) {
		if iv == nil {
			return nil, errors.New("No IV")
		}
		if len(iv) != int(C.EVP_CIPHER_iv_length(cipher)) {
			return nil, errors.New("Invalid IV")
		}
	}

	ctx := C.EVP_CIPHER_CTX_new()
	if ctx == nil {
		return nil, GetErrors()
	}

	ret := &CipherContext{ctx}
	runtime.SetFinalizer(ret, func(ret *CipherContext) {
		C.EVP_CIPHER_CTX_free(ret.ctx)
	})

	enc := 1
	if encrypt == false {
		enc = 0
	}

	if 1 != C.EVP_CipherInit(ctx, cipher, (*C.uchar)(&key[0]),
		(*C.uchar)(&iv[0]), C.int(enc)) {
		return nil, GetErrors()
	}

	return ret, nil
}

func (ctx *CipherContext) Update(in []byte) ([]byte, error) {
	outbuf := make([]byte, len(in)+int(C.EVP_CIPHER_CTX_block_size(ctx.ctx)))
	outlen := C.int(len(outbuf))
	if 1 != C.EVP_CipherUpdate(ctx.ctx, (*C.uchar)(&outbuf[0]), &outlen,
		(*C.uchar)(&in[0]), C.int(len(in))) {
		return nil, GetErrors()
	}
	return outbuf[:outlen], nil
}

func (ctx *CipherContext) Final() ([]byte, error) {
	outbuf := make([]byte, int(C.EVP_CIPHER_CTX_block_size(ctx.ctx)))
	outlen := C.int(len(outbuf))
	if 1 != C.EVP_CipherFinal(ctx.ctx, (*C.uchar)(&outbuf[0]), &outlen) {
		return nil, GetErrors()
	}
	return outbuf[:outlen], nil
}
