package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	frequency := 0

	changesList := strings.Split(INPUT, "\n")
	frequenciesSeen := make(map[int]bool)
	finished := false

	for !finished {
		for _, value := range changesList {
			modifier, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("Error converting str to int: %+v", err)
				break
			}
			frequency += modifier

			if _, ok := frequenciesSeen[frequency]; ok {
				finished = true
				break
			} else {
				frequenciesSeen[frequency] = true
			}

		}
	}

	fmt.Println(frequency)
}

const INPUT = `-19
+8
-10
+9
+2
+19
+5
+8
+13
+16
+4
-8
+16
-10
+11
+13
-12
-5
-15
-12
+9
-6
-18
+5
+3
+18
-3
+5
-19
+21
-9
-9
+20
-18
+11
+11
+17
+17
-20
-7
+5
+7
+7
+7
+7
+16
-12
+18
+10
+3
+15
-10
+8
+5
+7
-16
-16
-2
-15
+3
-6
+9
-15
+19
-12
-20
-1
-12
+11
-22
+10
+15
+2
-10
-19
+1
-18
-21
-4
+16
-15
+13
+20
+7
+23
+13
-3
+10
+8
-7
+4
-19
-7
-4
+19
+20
-18
-5
+11
+22
-14
+15
-3
+16
+20
-16
-11
-12
+7
-13
+7
+12
+20
+9
-15
+1
-14
+11
+16
+13
+9
-16
+18
-8
-6
+17
+7
+15
+7
-15
-15
-14
-16
+15
+17
+19
+6
-16
+13
-12
+20
-18
-11
-8
-14
+18
-11
+6
-16
+14
-10
-6
+1
+16
+4
-6
-8
+27
+18
+10
-3
-2
+11
-10
+11
+7
+19
-1
+11
+14
-16
-10
+15
-9
+2
-1
+13
-4
+6
-9
+10
-2
-10
+17
-14
-10
+15
+17
-12
-15
+18
+8
+13
+3
+17
+14
+1
+15
+18
+7
+15
-2
+9
+3
-1
-18
-15
-11
-10
-1
-14
+13
-18
-2
+5
+5
-7
-5
-12
+4
+19
+15
+9
-17
-10
-5
+13
+15
+7
-10
+12
+1
+9
-18
+14
+20
+2
+5
+7
+6
+13
+4
+12
-9
+17
+10
-5
-9
+18
+4
+5
-16
+1
+16
+1
-5
+19
-5
+14
-15
-3
+7
-15
+14
-7
+16
+13
+1
+17
-4
-9
+12
+15
+10
+5
+8
+31
+22
+12
+16
-17
+9
-21
-13
+22
-1
-2
+17
+14
+34
+18
-15
+10
+14
+11
+4
-9
+13
-12
+7
+3
-15
+6
+5
+21
-18
-20
+7
+18
-11
+8
-9
+11
-19
-2
-2
-20
-13
+16
-14
-1
+4
+20
-10
+6
-18
+15
+41
-1
+25
+21
+15
+1
-14
-1
-11
+17
-9
-5
-1
+21
-2
+9
+3
-13
-16
+1
+13
-4
-18
-8
-20
-1
+2
+1
+4
-8
-46
-29
-16
-5
-36
+35
-144
+3
+2
-7
-1
-11
-20
-22
-12
-18
-2
-18
+8
-18
+3
-1
+7
-18
+7
+19
-12
-19
-19
-14
+15
+15
+6
-17
+12
+11
+24
+16
+4
+12
+4
-18
+3
+19
-10
+12
-4
+19
+3
+2
+5
+19
+5
+22
-21
+1
-21
-19
+13
-12
+10
-12
-13
-2
-11
-16
-7
+10
-21
-3
+4
+22
+18
+11
+1
-10
+15
+5
-9
+2
+1
-24
-1
-2
+6
-26
-15
-16
+9
+25
+14
-7
-27
-24
+16
-30
-30
+26
+12
+3
-32
+3
-10
-4
+8
-19
-7
+15
+13
-3
+16
+7
-21
-23
-12
+14
-149
-77
+19
+6
-80
+11
-16
-19
-65
+6
-67327
-14
+16
-15
-10
+4
+16
-12
-3
-10
+7
+14
-2
+17
+8
-7
+15
-17
-13
-11
-18
+19
-22
-11
-1
-16
-2
+10
-12
-6
+11
+9
+1
+13
+9
-6
+19
-10
-4
-12
+14
+1
+19
+15
+8
-16
+1
+8
+18
-12
-5
+10
+15
-9
-15
+12
+8
-9
+3
+16
+5
+14
+18
-1
-5
+4
+17
+12
-9
-12
-15
+11
-19
+1
-13
+17
-11
-11
-14
+8
+10
-15
+18
-12
+4
-1
+10
+16
+11
+3
-11
-14
-4
+17
+13
+4
+12
-18
-16
-20
+14
-22
-20
-20
+6
+16
-15
-35
-7
-17
-20
+19
-17
-11
-9
-16
-3
+7
-10
-2
+7
+2
+2
+16
+14
+17
-8
-18
+2
-6
+3
-10
-8
+7
+5
+18
-20
-13
-17
-17
-7
+9
+13
-17
+16
+4
+12
-3
+1
-12
-7
+16
+13
+1
+14
-3
-4
+16
-21
-20
-17
+1
-14
-10
-5
-14
-1
-5
+19
+3
+16
-9
-13
+2
+3
-7
-15
+4
+6
+8
-10
+5
+12
+8
+18
-7
+8
-15
+5
-11
-15
+16
+16
+7
-2
+14
+18
+24
+9
+19
+17
+13
-16
-13
-14
-20
-5
+18
-3
+20
+13
+15
+7
+11
-15
-22
-15
-6
+18
-8
+18
-6
-24
-12
+21
+17
+18
+40
+22
+27
-4
-16
+17
-10
-18
+49
+15
-12
+7
+13
-16
+10
+15
-1
-3
+16
+6
-7
+4
-18
+13
-4
-16
+18
-5
+1
+9
-4
-2
+19
+17
-4
-14
-24
-21
+1
+6
-18
+8
+15
-18
-2
-31
-30
+4
+2
-28
-53
-33
+10
-11
-18
-19
-4
-6
-24
-6
-4
+11
-24
-16
+3
-1
-1
+13
-15
-10
-7
-3
+8
-11
+2
+5
+8
-3
-7
+15
+1
-11
-11
-2
-4
+16
+5
+14
-24
-14
-22
-11
-15
-6
+8
+15
+19
-14
-12
+1
+9
+5
+18
+6
-18
+16
-8
-10
-8
-9
+1
+11
-16
+8
+18
+12
+8
+2
+10
+11
-2
+14
+1
+11
-9
-18
-24
+2
-8
+15
-14
-18
+11
+2
-10
-11
-5
-13
-3
-12
-5
-13
+16
+5
-19
-16
-12
-4
+10
-15
+14
-6
+3
-18
+17
+12
-2
+6
+13
+16
-11
+3
+3
-20
-6
+8
-9
-19
-3
+16
-14
-14
+1
-12
-10
+19
-12
-18
+5
-3
+7
-20
-16
-18
+17
-23
+25
+65
+11
+11
+16
-1
+16
-9
+19
+13
-8
-10
+9
-18
-16
-16
-19
+1
-28
+4
+35
+13
+12
+15
-5
-3
+4
+16
+2
+10
+15
+17
+18
+27
+19
+15
+13
+6
-17
+12
+15
+9
+23
+78
-221
-28
+19
+18
-13
-22
-11
-5
-5
-9
-16
-2
+11
+30
+10
+28
-4
-72
-36
+10
-5
+77
+35
+102
+153
-29
-676
-67225
-5
+13
+19
+3
+15
+7
-11
-1
-7
+13
+17
+17
+6
-15
-9
+11
-8
-7
+10
+1
+13
+19
-11
-11
+2
+4
-11
-7
+10
+19
+17
+7
+12
+20
-2
+135638`
