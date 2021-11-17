package hangman

/*

func AsciiArt () {
	tbAsciiArt := ArrayAscii()
	for i := 0; i < len(tbAsciiArt); i++ {
		fmt.Println(tbAsciiArt[i])
		fmt.Println("*")
	}
	tbLigne := [8] string {}
	doubleArray := [][] string {}

	for i := 0; i < len(tbAsciiArt); i++ {
		fmt.Println("88")
		fmt.Println(tbAsciiArt[i])
		fmt.Println("88")
		for j := 0; j < len(tbLigne); j ++ {
			fmt.Println(("88"))
			fmt.Println(tbAsciiArt[60])
			tbLigne[j] = SplitCara(tbAsciiArt[i], j)
			fmt.Println(doubleArray[i][j])
		}
	}
	fmt.Println(doubleArray)
}

func SplitCara (AsciiArt string, tbLigne[]string, i int ) string {

	splitC := strings.Split(AsciiArt, "\r\n" )
	for j := 0 ; j < len(splitC); j++ {
		fmt.Println(splitC[j])
		fmt.Println("*")
	}
	fmt.Println("++")
	return splitC[i]
}

func ArrayAscii () []string {
	f, e := os.Open("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\standard.txt")
	check(e) // check erreur : par de fichier
	defer f.Close()
	data, err := os.ReadFile("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\standard.txt")
	check(err)
	tbAsciiArt := strings.Split(string(data), "\r\n\r\n")
	return tbAsciiArt
}

*/
