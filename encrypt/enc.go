package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"strconv"
	"sync"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

var LoginRManager *LoginRTable

type LoginRInfo struct {
	ID    string
	R1    string
	R2    string
	RTime time.Time
}

type LoginRTable struct {
	List     map[string]LoginRInfo
	ListLock sync.Mutex
}

func generateSalt() []byte {
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic(err)
	}
	return salt
}

func generateRandomBytes(size int) ([]byte, error) {
	randomBytes := make([]byte, size)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

func aesEncrypt(key, plaintext []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(iv) == 0 {
		iv = make([]byte, aes.BlockSize)
	}
	//fmt.Printf("%v \n", iv)

	paddedPlaintext := addZeroPadding(plaintext, block.BlockSize())

	ciphertext := make([]byte, len(paddedPlaintext))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	return ciphertext, nil
}

func aesDecrypt(key, ciphertext []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(iv) == 0 {
		iv = make([]byte, aes.BlockSize)
	}
	//fmt.Printf("%v \n", iv)

	plaintext := make([]byte, len(ciphertext))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext = removeZeroPadding(plaintext)

	return plaintext, nil
}

func addZeroPadding(plaintext []byte, blockSize int) []byte {
	paddingSize := blockSize - (len(plaintext) % blockSize)
	padding := make([]byte, paddingSize)
	return append(plaintext, padding...)
}

func removeZeroPadding(plaintext []byte) []byte {
	paddingSize := 0
	for i := len(plaintext) - 1; i >= 0; i-- {
		if plaintext[i] != 0 {
			paddingSize = len(plaintext) - i - 1
			break
		}
	}

	return plaintext[:len(plaintext)-paddingSize]
}

// pbkdf2 is a helper function that performs PBKDF2 key derivation with the specified parameters.
func PBKDF2(data []byte, salt []byte, iterations int, keylen int, hash func() hash.Hash) []byte {
	dk := make([]byte, keylen)
	pbkdf2 := pbkdf2.Key(data, salt, iterations, keylen, hash)
	copy(dk, pbkdf2)
	return dk
}

func (t *LoginRTable) DumyCreateInfo(count int) {

	for i := 1; i <= count; i++ {
		fmt.Printf("add : %d \n", i)
		info := t.CreateInfo("id-"+strconv.Itoa(i), "test11")
		fmt.Printf("add : %v \n", *info)
		time.Sleep(100 * time.Millisecond)
	}
}

func (t *LoginRTable) CreateInfo(id, r1 string) (loginInfo *LoginRInfo) {
	var Info *LoginRInfo
	Info, err := t.GetInfo(id)

	if err == nil {
		fmt.Printf("Get Info : %v \n", Info)
		loginInfo = Info
	} else {
		fmt.Printf("Get Info : %v \n", Info)
		loginInfo = &LoginRInfo{}
	}
	r2 := hex.EncodeToString(generateSalt())
	loginInfo.ID = id
	loginInfo.R1 = r1
	loginInfo.R2 = r2
	loginInfo.RTime = time.Now()

	t.ListLock.Lock()
	defer t.ListLock.Unlock()
	t.List[id] = *loginInfo

	return
}

func (t *LoginRTable) RemoveInfo(id string) error {
	t.ListLock.Lock()
	defer t.ListLock.Unlock()

	delete(t.List, id)
	return nil
}

func (t *LoginRTable) GetInfo(id string) (*LoginRInfo, error) {
	//t.ListLock.Lock()
	//defer t.ListLock.Unlock()
	if info, ok := t.List[id]; ok {
		return &info, nil
	}
	return nil, fmt.Errorf("not found")
}

func (t *LoginRTable) PeekInfo(id string) (LoginRInfo, error) {
	t.ListLock.Lock()
	defer t.ListLock.Unlock()
	if info, ok := t.List[id]; ok {
		delete(t.List, id)
		return info, nil
	}
	return LoginRInfo{}, fmt.Errorf("not found")
}

func (t *LoginRTable) Clear() {
	t.ListLock.Lock()
	defer t.ListLock.Unlock()

	t.List = make(map[string]LoginRInfo)
}
func (t *LoginRTable) timeExpiredCheck() {

	go func() {
		for {
			if len(t.List) > 0 {
				for _, info := range t.List {
					if info.RTime.UnixMilli()+1000 < time.Now().UnixMilli() {
						t.RemoveInfo(info.ID)
						fmt.Printf("delete : %v \n", info)
					}
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

}

func TestEnc(pass string) {

	// r1, err := generateRandomBytes(16)
	// if err != nil {
	// 	fmt.Println("Failed to generate random bytes for r1:", err)
	// 	return
	// }

	//r1 := []byte("1234567890123456789012")
	//	fmt.Printf("r1 byte :  %v\n", r1)

	// Encrypt with AES-256
	aesKey := "12345678901234567890123456789012" //hex.EncodeToString(r1) //32byte string
	fmt.Printf("r1 str :  %s\n", aesKey)
	fmt.Printf("r1 byte :  %v\n", []byte(aesKey))
	fmt.Printf("r1 16 byte  :  %v\n", []byte(aesKey)[0:16])
	PBKDF2Key := PBKDF2([]byte(aesKey), []byte(aesKey)[0:16], 1000, 32, sha256.New)
	fmt.Printf("PBKDF2Key byte :  %v\n", PBKDF2Key)
	pbkdf2Str := hex.EncodeToString(PBKDF2Key)
	fmt.Printf("PBKDF2Key hexStr :  %s\n", pbkdf2Str)
	plaintext := []byte(pass)

	var iv []byte
	//ciphertext, err := aesEncrypt(PBKDF2Key, plaintext, iv)
	ciphertext, err := aesEncrypt([]byte(pbkdf2Str)[0:32], plaintext, iv)
	if err != nil {
		fmt.Println("AES encryption failed:", err)
		return
	}
	fmt.Printf("pass : %s -> Ciphertext byte : %v \n", string(plaintext), ciphertext)

	//암호화된 값을 hex encoding 값으로 받아서 복호화
	uEndHex := hex.EncodeToString(ciphertext)
	uDecHex, _ := hex.DecodeString(uEndHex)
	fmt.Println("HexEncoding : ", uEndHex)
	// Decrypt AES-256 ciphertext
	decryptedPlaintext, err := aesDecrypt([]byte(pbkdf2Str)[0:32], uDecHex, iv)
	if err != nil {
		fmt.Println("AES decryption failed:", err)
		return
	}
	fmt.Println("Decrypted plaintext:", string(decryptedPlaintext))

	//암호화된 값을 url encoding 값으로 받아서 복호화
	uEnc := base64.URLEncoding.EncodeToString(ciphertext)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Printf("URLEncoding : %v \n", uEnc)
	decryptedPlaintext, err = aesDecrypt([]byte(pbkdf2Str)[0:32], uDec, iv)
	if err != nil {
		fmt.Println("AES decryption failed:", err)
		return
	}
	fmt.Println("Decrypted plaintext:", string(decryptedPlaintext))

}

func init() {
	fmt.Println("Initialize memory login-key")
	LoginRManager = &LoginRTable{}
	LoginRManager.List = make(map[string]LoginRInfo)
	LoginRManager.timeExpiredCheck()
}
