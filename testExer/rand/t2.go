package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var theSameRandLen int
	var srcRandLen int
	var dstRandLen int
	theSameRandLen = rand.Intn(1000)
	for theSameRandLen < 500 {
		theSameRandLen = rand.Intn(1000)
	}
	for srcRandLen < 500 {
		srcRandLen = rand.Intn(1000)
	}
	for dstRandLen < 500 {
		dstRandLen = rand.Intn(1000)
	}
	theSame := []string{}
	SameStr := "The Same "
	for i := 0; i < theSameRandLen; i++ {
		theSame = append(theSame, SameStr+string(rand.Intn(95)+33)+"\n")
	}
	srcOnly := []string{}
	srcStr := "The Only Src "
	for i := 0; i < srcRandLen; i++ {
		srcOnly = append(srcOnly, srcStr+string(rand.Intn(95)+33)+"\n")
	}
	dstOnly := []string{}
	dstStr := "The Only Dst "
	for i := 0; i < dstRandLen; i++ {
		dstOnly = append(dstOnly, dstStr+string(rand.Intn(95)+33)+"\n")
	}

	fmt.Println(len(theSame))
	fmt.Println(len(srcOnly))
	fmt.Println(len(dstOnly))

	now := time.Now()
	fileNameStr := fmt.Sprintf("%d%d%d", now.Day(), now.Hour(), now.Second())

	src, srcErr := os.Create("../testFile/" + "src" + fileNameStr)
	if srcErr != nil {
		log.Fatalln(srcErr)
		// log.Fatal(srcErr)
	}
	theSameDst := []string{}
	for srcOnly != nil || theSame != nil {
		if rand.Intn(2) == 0 && srcOnly != nil {
			src.WriteString(srcOnly[0])
			if len(srcOnly) > 1 {
				srcOnly = srcOnly[1:]
			} else {
				srcOnly = nil
			}
			continue
		}
		if theSame != nil {
			src.WriteString(theSame[0])
			theSameDst = append(theSameDst, theSame[0])
			if len(theSame) > 1 {
				theSame = theSame[1:]
			} else {
				theSame = nil
			}
		}
	}
	src.Close()
	dst, dstErr := os.Create("../testFile/" + "dst" + fileNameStr)
	if dstErr != nil {
		log.Fatalln(dstErr)
		// t.Fatal(dstErr)
	}
	for dstOnly != nil || theSameDst != nil {
		if rand.Intn(2) == 0 && dstOnly != nil {
			dst.WriteString(dstOnly[0])
			if len(dstOnly) > 1 {
				dstOnly = dstOnly[1:]
			} else {
				dstOnly = nil
			}
			continue
		}
		if theSameDst != nil {
			dst.WriteString(theSameDst[0])
			if len(theSameDst) > 1 {
				theSameDst = theSameDst[1:]
			} else {
				theSameDst = nil
			}
		}
	}
	dst.Close()
}
