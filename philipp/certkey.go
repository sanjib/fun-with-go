package main

import "path/filepath"

func localCertKey() (string, string) {
	cert := filepath.Join("C:/Users/sanji/lab/ssl/2020-sep/viaye.san/", "cert.pem")
	key := filepath.Join("C:/Users/sanji/lab/ssl/2020-sep/viaye.san/", "server.key")
	return cert, key
}

func remoteCertKey() (string, string) {
	cert := filepath.Join("/home/vrtapps/ssl/certs/", "virtual_apps_com_d2e98_7265f_1613001599_5bece736b9e35e4bf4b0a547892ad8c1.crt")
	key := filepath.Join("/home/vrtapps/ssl/keys/", "d2e98_7265f_e149b0045708bf0aaa1de70d10abd7f0.key")
	return cert, key
}
