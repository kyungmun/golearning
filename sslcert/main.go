package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	// 테스트용 인증서 및 개인 키 생성
	priv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Println("키 생성 실패:", err)
		return
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour) // 1년 유효 기간

	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{Organization: []string{"Test"}},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		fmt.Println("인증서 생성 실패:", err)
		return
	}

	// 개인 키를 PEM 형식으로 인코딩하여 파일로 저장
	privateKeyFile, err := os.Create("private_key.pem")
	if err != nil {
		fmt.Println("개인 키 파일 생성 실패:", err)
		return
	}
	defer privateKeyFile.Close()

	privateKeyBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		fmt.Println("개인 키 마샬링 실패:", err)
		return
	}

	privateKeyPEM := &pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKeyBytes}
	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		fmt.Println("개인 키 파일 저장 실패:", err)
		return
	}
	fmt.Println("개인 키 파일 저장 완료: private_key.pem")

	// 인증서를 PEM 형식으로 인코딩하여 파일로 저장
	certificateFile, err := os.Create("certificate.pem")
	if err != nil {
		fmt.Println("인증서 파일 생성 실패:", err)
		return
	}
	defer certificateFile.Close()

	certificatePEM := &pem.Block{Type: "CERTIFICATE", Bytes: certDER}
	err = pem.Encode(certificateFile, certificatePEM)
	if err != nil {
		fmt.Println("인증서 파일 저장 실패:", err)
		return
	}
	fmt.Println("인증서 파일 저장 완료: certificate.pem")

	// // PFX 파일로 저장
	// pfxData, err := createPFX(priv, certDER, "your_password")
	// if err != nil {
	// 	fmt.Println("PFX 파일 생성 실패:", err)
	// 	return
	// }

	// err = os.WriteFile("test.pfx", pfxData, 0600)
	// if err != nil {
	// 	fmt.Println("PFX 파일 저장 실패:", err)
	// 	return
	// }
	// fmt.Println("PFX 파일 저장 완료: test.pfx")
}

// func createPFX(privateKey *ecdsa.PrivateKey, certDER []byte, password string) ([]byte, error) {
// 	pfxData, err := pkcs12.Encode(rand.Reader, privateKey, &x509.Certificate{Raw: certDER}, nil, password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pfxData, nil
// }
