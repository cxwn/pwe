package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

//  PublicKeyEncrypt RSA 公钥加密方法，第一个参数为需要加密的字符串，第二个参数为 RSA 公钥字符串。加密后返回一个 Base64 编码的字符串及错误。
func PublicKeyEncrypt(text string, publicKey string) (string, error) {
	var publicKeyBytes []byte = []byte(publicKey)

	block, _ := pem.Decode(publicKeyBytes)
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		panic(err)
	}

	// 类型断言
	pubKey := publicKeyInterface.(*rsa.PublicKey)

	// 对明文进行加密
	encryptText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(text))
	if err != nil {
		panic(err)
	}

	// 返回base64编码字符串
	return base64.StdEncoding.EncodeToString(encryptText), nil
}

// PrivateKeyDecrypt 私钥解密方法，第一个参数为base64编码的加密字符串，第二个参数为 RSA 私钥字符串。解密后返回原始字符串。
func PrivateKeyDecrypt(encryptText string, privateKey string) (string, error) {
	decryptTextBytes, err := base64.StdEncoding.DecodeString(encryptText)

	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode([]byte(privateKey))

	//X509解码
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	//对密文进行解密
	decryptText, _ := rsa.DecryptPKCS1v15(rand.Reader, priKey, decryptTextBytes)

	//返回明文
	return string(decryptText), nil
}


// GetPublicKeyFromPrivateKey 从私钥中推导公钥
func GetPublicKeyFromPrivateKey(privateKeyString string) string {
	privateKeyBytes := []byte(privateKeyString)
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		panic(errors.New("解码包含公钥的PEM块失败"))
	}
	publicKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}
	publicKeyDer, _ := x509.MarshalPKIXPublicKey(&publicKey.PublicKey)
	publicKeyBlock := pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyDer}
	return string(pem.EncodeToMemory(&publicKeyBlock))
}


func GenerateRSAKeys(bits int) {

	if createDir("keys") {
		fmt.Println("keys目录已成功新建，本次生成的公钥和私钥将存放于该目录，请谨慎保存！")
	} else {
		fmt.Println("keys目录已存在，本次生成的公钥和私钥将存放于该目录，请谨慎保存！")
	}

	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串，使用pem格式对x509输出的内容进行编码。
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	//创建文件保存私钥
	privateFile, err := os.Create("keys/private.key")
	if err != nil {
		panic(err)
	}

	defer privateFile.Close()

	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: X509PrivateKey}

	//将数据保存到文件
	pem.Encode(privateFile, &privateBlock)

	//获取公钥的数据
	publicKey := privateKey.PublicKey

	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}

	//创建用于保存公钥的文件
	publicFile, err := os.Create("keys/public.key")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()

	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "PUBLIC KEY", Bytes: X509PublicKey}

	pem.Encode(publicFile, &publicBlock)
}

// createDir 创建存放公钥和私钥的目录，如已经存在则不创建.
func createDir(name string) bool {
	_, err := os.Stat(name)
	if !os.IsExist(err) {
		e := os.Mkdir("keys", os.ModePerm)
		if e != nil {
			return false
		}
	} else {
		return false
	}
	return true
}