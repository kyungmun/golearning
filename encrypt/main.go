package main

func main() {
	/*
		md5, err := crypt.Crypt("Admin12#$", "esFlh")
		if err != nil {
			fmt.Errorf("error:", err)
			return
		}

		result2, err := crypt.Crypt("Admin12#$", "$5$esFlh$")
		if err != nil {
			fmt.Errorf("error:", err)
			return
		}

		sha512, err := crypt.Crypt("Admin12#$", "$6$esFlh$")
		if err != nil {
			fmt.Errorf("error:", err)
			return
		}

		fmt.Println("MD5  pass : Admin12#$, salt : esFlh => ", md5)
		fmt.Println("SHA256 (5) pass : Admin12#$, salt : esFlh => ", result2)
		fmt.Println("SHA512 (6) pass : Admin12#$, salt : efFlh => ", sha512)
	*/
	//TestEnc("Admin12#$")
	//TestEnc("Admin13#$")
	TestEnc("Admin12#$")
}
