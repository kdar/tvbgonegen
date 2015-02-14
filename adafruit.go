package tvbgonegen

// var IRsignal = []int{
// 	240, 60,
// 	120, 60,
// 	60, 60,
// 	120, 60,
// 	60, 60,
// 	120, 60,
// 	60, 60,
// 	60, 60,
// 	120, 60,
// 	60, 60,
// 	60, 60,
// 	60, 60,
// 	60, 2700,
// 	240, 60,
// 	120, 60,
// 	60, 60,
// 	120, 60,
// 	60, 60,
// 	120, 60,
// 	60, 60,
// 	60, 60,
// 	120, 60,
// 	60, 60,
// 	60, 60,
// 	60, 60,
// 	60, 0,
// }

// // https://github.com/adafruit/Raw-IR-decoder-for-Arduino
// func compressAdafruit() {
// 	var IRsignal = []int{
// 		// ON, OFF (in 10's of microseconds)
// 		898, 434,
// 		64, 48,
// 		62, 48,
// 		64, 48,
// 		62, 158,
// 		62, 156,
// 		64, 156,
// 		64, 48,
// 		62, 48,
// 		62, 158,
// 		62, 158,
// 		62, 158,
// 		62, 48,
// 		64, 48,
// 		62, 48,
// 		64, 156,
// 		62, 158,
// 		62, 48,
// 		64, 156,
// 		64, 48,
// 		62, 48,
// 		64, 156,
// 		62, 48,
// 		64, 48,
// 		62, 48,
// 		64, 156,
// 		64, 48,
// 		62, 156,
// 		64, 156,
// 		64, 48,
// 		62, 158,
// 		62, 158,
// 		62, 158,
// 		62, 3914,
// 		892, 216,
// 		62, 2884,
// 		894, 214,
// 		62, 0}

// 	var times [][2]int //= make([][2]int, 4)
// 	var codeTimeIndex []int

// 	// times[0][0] = 60
// 	// times[0][1] = 60
// 	// times[1][0] = 60
// 	// times[1][1] = 2700
// 	// times[2][0] = 120
// 	// times[2][1] = 60
// 	// times[3][0] = 240
// 	// times[3][1] = 60

// 	for i := 0; i < len(IRsignal); i += 2 {
// 		pair := [2]int{IRsignal[i], IRsignal[i+1]}
// 		// if pair[1] == 0 {
// 		// 	continue
// 		// }

// 		found := false
// 		timeIndex := 0
// 		for x := 0; x < len(times); x++ {
// 			found = int(math.Abs(float64(times[x][0]-pair[0]))) < marginOfError
// 			found = found && (int(math.Abs(float64(times[x][1]-pair[1]))) < marginOfError || pair[1] == 0)
// 			if found {
// 				timeIndex = x
// 				break
// 			}
// 		}

// 		if !found {
// 			times = append(times, pair)
// 			timeIndex = len(times) - 1
// 		}

// 		codeTimeIndex = append(codeTimeIndex, timeIndex)

// 		// lastCode |= byte((uint(timeIndex) << codeShift)) & 0xFF
// 		// //fmt.Println(strconv.FormatInt(int64(lastCode), 2))
// 		// codeShift -= 2

// 		// if codeShift == 0 || i+2 >= len(IRsignal) {
// 		// 	codes = append(codes, lastCode)
// 		// 	codeShift = 8
// 		// 	lastCode = 0
// 		// }
// 	}

// 	str := ""
// 	for i, v := range times {
// 		str += fmt.Sprintf("%d,%d", v[0], v[1])
// 		if i < len(times)-1 {
// 			str += ", "
// 		}
// 	}

// 	fmt.Printf(timesTpl, str)

// 	// Determine how many bits we need to use
// 	bitLen := 0
// 	for i := 8; i >= 0; i-- {
// 		if (1<<uint(i-1))&(len(times)-1) > 0 {
// 			bitLen = i
// 			break
// 		}
// 	}

// 	var codesStr string
// 	for _, v := range codeTimeIndex {
// 		binary := strconv.FormatInt(int64(v), 2)
// 		zeros := bitLen - len(binary)
// 		for i := 0; i < zeros; i++ {
// 			binary = "0" + binary
// 		}

// 		codesStr += binary
// 	}

// 	codebig := big.NewInt(0)
// 	codebig.SetString(codesStr, 2)
// 	codesStr = fmt.Sprintf("%X", codebig)
// 	if len(codesStr)%2 != 0 {
// 		codesStr += "0"
// 	}

// 	str = ""
// 	for i := 0; i < len(codesStr); i += 2 {
// 		str += "0x" + string(codesStr[i]) + string(codesStr[i+1])
// 		if i < len(codesStr)-2 {
// 			str += ", "
// 		}
// 	}

// 	fmt.Printf(codesTpl, len(IRsignal)/2, bitLen, str)
// }
