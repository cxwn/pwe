package main

import (
	"fmt"
	"pwe/encrypt"
)

// 384 位的公钥长度，可加密包含数字、字母、符号的最长21位的密码，尝过该长度程序会报错，一般情况下256位加密即可满足要求，加密位数过高会导致效率低下。
var privateKey string = `-----BEGIN RSA PRIVATE KEY-----
MIHxAgEAAjEAw+3qUW6PYKDuF478HYzC7fV0at5NRSjHoLKEPqPNiBlg3QxV42bk
crVu+qZWJfORAgMBAAECMGuWMivjDQIffH4dOt2zFLr8JKAmT8HhQL8PW4Nw4dS6
b7rsv4OiRHDE2AjDP0Lj0QIZANZS6vwWFvySDLV8tf0WX/8yTlYcAxM7lQIZAOoH
UFbNq5QZRQy0F+P385hYRq6whlj5DQIYaRjGBBGs+fOAaeqar12ue0ym31DjLSY1
AhgWEYJ97P8VBB0CyajHEoaiAxEHQBYrJbECGDzN9H7nRPr+2naCeLXW5NKs49vo
LdRirA==
-----END RSA PRIVATE KEY-----`

var publicKey string = encrypt.GetPublicKeyFromPrivateKey(privateKey)

func main() {
	// encrypt.GenerateRSAKeys(384)
	fmt.Printf("\n----------------------------------------------------------------------------------------------------------------------------\n\n")
	passwords := []string{"thinkit_offline_asr@123$", "thinkit_offline@1234#!", "speech_offline_b@1234#!", "thinkit_offline_sd@123$", "thinkit_offline_hl@123$", "thinkit_offline_hb@123$", "thinkit_offline_sn@123$", "thinkit_offline_tj@123$", "thinkit_offline_jl@123$", "thinkit_offline_sxc@123$", "thinkit_offline_jx@123$", "I7N&q0DM=Ci-UK|!", "[%Ydh9]=vC3*{Lzl", "*oz{fXn,=KawHbjc", "/<-6(o?b%phKrSJB", "K0<%h=fs4*v48NT|", "bx@aSqE:0Y>li?5t", "]BSU6K}+C)DZFX^c", "Ao6P#$l!O3yEf0?h", "u(3MA-TbGYs<imFD", "G>!y$]7}SCzZrcQg", "qg#Hi-wyO^Y0Sh9D", "PwBzct-YEuG,iRkq", "*aQ^g0|2+$#U!DEj", "S$R9>lXDj$Jw]Gkz", "6c|_H(JK>0eCa{74", "{bPAMgE$}j0Z8QFz", "[LA1k4+=O~}&8@l>", "Lw_jy:TcrM/([Jfm", "SC1Zw9MHP3#yq!(a", "W%1}:Eg[S>&TSZMq", "4WX(65-E*)9Kachy", "<e$utN-&5Z|(l=gz", "AVshQ()P5K}ta$gF"}
	provinces := []string{"ASR", "TEST", "B", "SD", "HL", "HB", "SN", "TJ", "JL", "SXC", "JX", "LN", "GX", "HE", "HA", "YN", "NX", "GD", "ZJ", "GZ", "GS", "SC", "XJ", "BJ", "AH", "CQ", "FJ", "HI", "HN", "NM", "JS", "QH", "SH", "XZ"}
	for i, password := range passwords {
		for j, province := range provinces {
			if i == j {
				encryptText, _ := encrypt.PublicKeyEncrypt(password, publicKey)
				fmt.Println(province + "\t原始密码：" + password + "\t加密后密码：" + encryptText)

			}
	
		}
	}
	pause()
}

func pause() {
	fmt.Println("--------------------------------------------------------------------------------------------------------------------")
	var s string
	fmt.Println("退出请输入exit！")
	fmt.Scan(&s)

	if s == "exit" {
	} else {
		pause()
	}
}
