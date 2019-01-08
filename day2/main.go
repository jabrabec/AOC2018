package main

import (
	"fmt"
	"strings"
)

func findChecksum() {
	boxIDsList := strings.Split(BOXIDS, "\n")

	masterCount := map[int]int{2: 0, 3: 0}

	for _, value := range boxIDsList {
		lettersSeen := make(map[string]int)
		for i, _ := range value {
			lettersSeen[string(value[i])] += 1
		}

		seenTwo := false
		seenThree := false
		for _, v := range lettersSeen {
			if seenTwo == true && seenThree == true {
				break
			}
			if v == 2 {
				if seenTwo == false {
					masterCount[2] += 1
					seenTwo = true
				}

			} else if v == 3 {
				if seenThree == false {
					masterCount[3] += 1
					seenThree = true
				}
			}
		}
	}

	checkSum := 1
	for _, v := range masterCount {
		checkSum = checkSum * v
	}
	fmt.Println(checkSum)
}

func main() {
	findChecksum()
}

const BOXIDS = `auxwcbzrmdvpsjfgkrthnkioqm
auxwcbzrmdvpsjfgbltonyijqe
auxwcbzrmdfpsefgklthnoioqe
auxwcbzrmdvpsjfgkluhnjisqe
auxwcbzrmdvesjfgdzthnyioqe
auxwcbzrmdvhsjfgklthnmijqe
auxwcbzridvpsjfgkltxeyioqe
ayxwcbzrgdvpsjfgklthiyioqe
ajxwcbzrmdvpsjfgklkhnyiode
auxwcbcrmdvpsjfqelthnyioqe
auxwcbzrmsvpsjsgklthnyiope
auxwcbzrmqvpsjzgklghnyioqe
auxwcbzrmdvpsjtqklthxyioqe
auxwcbzrmdopsjfdklthncioqe
auxwcbzrmdvpsjfgkltmhyfoqe
aixwcbzrmdvpsjfgllthdyeoqe
vuxicbzrmdepsjfgklthnyioqe
auxwcbbxmdkpsjfgklthnyioqe
auxwcbzrgdvpsofaklthnyioqe
auxycbzrmdvpsjfgklthnyuose
aujwcbzrmdvprjfgkltcnyioqe
auxwgbzrmdvpsjfgyzthnyioqe
auxwcbzrmavpsjfgkltsnyiome
auxwcbgrmdvpsjfgkdthnrioqe
kuxwcbzrmdvpsyfgklthnyioue
auxwcbzomdvpjdfgklthnyioqe
auxwcbzrmdppsjfgklthvyifqe
aunwubzrmdvpsjrgklthnyioqe
auxwcbzrmoipsjfgklbhnyioqe
auxwdbzrmdvpsjfgmlthnyioce
auxwcbzjmsvpsjfiklthnyioqe
auxwcbzrmwcpsjfcklthnyioqe
auxwcbzfmdvprjfhklthnyioqe
auxdcbzrgdvpsjfgklthnyxoqe
wuxwbbzrmdvpsjfgklthnyiote
auowcbjrmdvpsjfgklthnyfoqe
auxwsbzrmdvpsjfglltcnyioqe
quxwcbzrmdvpkjfgklthnyioqt
vuxwcbzrudvpsjfgklthnyioqi
puxwcbzrmdvgsjfgklthncioqe
luxdcbzrmdvpsjfgkothnyioqe
auxwcbzrmdvpsjfyklthfhioqe
auxwcbqrmdvpsjfgkldhnyiote
quxwcbzrmlvpsjfgklthnyioqi
auxwcbzgmdvpsjfoklthnyiuqe
auxwcbzrmdvpsbfgkltdjyioqe
auxwcbzsmdrpsjfgklthpyioqe
auxwcbzrmfvpsjfwklthnyiote
auxbkpzrmdvpsjfgklthnyioqe
auxwcbzrddvpsjfsklthnyroqe
abxwcbzrmdvpsjfgkltdnyivqe
awxwcbzrmvvpsjfgklthngioqe
auxwcbzrmkvgsjfgkltcnyioqe
auxwcbammdvpsjfgklthpyioqe
auxwcbhrmdvpsjfgtlthnuioqe
auxwcpzrmdvpbjogklthnyioqe
auxwcbzrmdvpslfgklbhkyioqe
auxwcbsrmdvpjjfgkldhnyioqe
auxwcbzrmdqpsjfgauthnyioqe
ydxwcbxrmdvpsjfgklthnyioqe
auxwcbzrmdvpejfgklthnyyofe
auxwchzrmxvpsjfgklthnyioqh
auxwcbzrtdvpsjfgklxhnzioqe
auxwcbyrmdvpsnfgklnhnyioqe
auxwcbzrcdvpsjugklihnyioqe
auxwcbzrddvpsjfgklhhnyiaqe
aumwtbzrmdvpsjfgklthnyitqe
auxucbzrmdvpsjfgklthwfioqe
auxwcbzrmdvpzmfgkllhnyioqe
auxwcbzrmdvpsjhgklthntiome
buxwzbzrmdvpszfgklthnyioqe
ouxwcbzsgdvpsjfgklthnyioqe
auxwcbzrmdvpsjfskltgnyioqz
auxwcbbrmdvpsjftklthnyioqu
quxocbzrmdvpsjfgklthfyioqe
acxwcbzrmdvpsjfgklfhnrioqe
auxwcbzrmdnpsjfrkjthnyioqe
wuxwybzrmdwpsjfgklthnyioqe
auxwgbxrmdvpsjfghlthnyioqe
atxwcbzrmdvnsjfgklthnyjoqe
acxwcbzmmdvpsjfbklthnyioqe
auxhcbzrmdvbsjbgklthnyioqe
auxwlbzrfdvpsjfgxlthnyioqe
auxwmbzrmdfpsjqgklthnyioqe
auxwcbzrmdvpsgfgklahnyigqe
auxwgbzrmdvpsjfgzldhnyioqe
auxwcbzrmdvpydfgklthnyiohe
auxwxbzrmdvpsjfsklchnyioqe
auxqcbzrmdvpsjfgqlthnyiwqe
auxwcozrmdvssbfgklthnyioqe
auxvcczrmdvpsufgklthnyioqe
auxwcbzrudvpsjfgklyhnyioxe
aulwcbzrmdvpsjqgknthnyioqe
auukcbzrmdvpsjfgklthtyioqe
auxwcszimdvpsjfgklthnyigqe
juxwcbzrbdvpsjfgklthnyboqe
auxwcbzrmdvpjofgklthnyioqj
auxwcbzrmdvpsjfgplfhnyione
auxwcbzrmdhpsjfgkltknyeoqe
luxwcqzrmdvpsjfgklthnbioqe
uuxwcbzrmdvpsjfgkithnyiiqe
auxwcbzrmdvpdjfgkrthnyeoqe
auuwcbnrmdvpsjfgklthnjioqe
auxwcnzrmdvpsjvgklthnyooqe
auxwcbzcmdvpsjfcklthnyiose
auxwcbzrldfpsjfgklthjyioqe
auxwcizrmdvpsjfjklthnymoqe
auxwcbtrmdvpsjfgtlphnyioqe
amxwcbzrmdvksjfgklthnyiove
auxwcbzrmdvpszfgkpthnyiuqe
auxwcbzrmdvxdjfgkltqnyioqe
auxwcbzrudvpsjfgklthnymiqe
auxwcbirmdvfsjfgklmhnyioqe
auwwcbzrndvprjfgklthnyioqe
auxwcbormdgpsjfgklbhnyioqe
auxwabzrmdupsjfgklthnyioqt
auxvcbzrmdvpsjfgkltrmyioqe
auxwcbzrmddpsjfsklthnyizqe
auxwcczrmuvpyjfgklthnyioqe
auxwcczrmdvpsnfgkpthnyioqe
auxkcbzrmdvpsjfhklihnyioqe
auxwcbzrmdvpsjfgklthnkijje
auxwcbzcmdvpsjpgkldhnyioqe
auxwcnzrudvpstfgklthnyioqe
xuxwcbzrgdvusjfgklthnyioqe
aaxwcbzrmdvpsjvgklthnyidqe
auxwcbztmdvpsjfgklthnyhqqe
auxwcbzrmfvpsjfgklthnyilfe
auxwcbzrmdvksjfgklthjyioqq
auxwcbzrmdzksjfgktthnyioqe
auxwcbzrmfvpszfgklohnyioqe
auxwckzamdvpsjfgklthnyioqs
auxwcmzrhdvpsjfaklthnyioqe
fuxwcbzrmdapsjfgklrhnyioqe
avxwxbzrmdvpsjfgklthniioqe
auxwubzrmevpsjfgkltpnyioqe
fuxwcbzrgdvpsjfgklhhnyioqe
duxwwbdrmdvpsjfgklthnyioqe
audwcbzrmdvpnjcgklthnyioqe
auxtcbzrmdvpsjmgklthnyyoqe
aucwcbwrmdepsjfgklthnyioqe
auxwcbzrudvpsjfpklthnyiose
auxwcbzridvpsjfsklthxyioqe
auxtcbzrmdvpscfgklyhnyioqe
auxwcbzrmdvppjfgklthnyivee
auxwdbzrmuvpskfgklthnyioqe
auxwubzrmdvosjfgklthnyiope
auxwcbzrmhnpsjfgklthnyimqe
auxwcbzrmdqpwjfgkltpnyioqe
auxwcbormdvpsjljklthnyioqe
auxwcbzrmdjpsjfgkltjpyioqe
auxwcbzrmdvpszfgklthkyizqe
auxwcbzighvpsjfgklthnyioqe
auxwcbzrmdlpsjfgcythnyioqe
auxwcbzumdvpsjflklthnyimqe
pdxwcbzrmdvpsjfgklthnyihqe
auxwcbzrsdvpsjfgklhhvyioqe
auxwcfzamdvpsjfgkmthnyioqe
aexwcdzrmdvpsjogklthnyioqe
auxxcbkrmavpsjfgklthnyioqe
auxwcbzredvssjfgklthryioqe
aupwqbzrmdvpsjfgklthnyioqc
auxwcbzrmdvpkcagklthnyioqe
auxwcbzrmdvwsbfgklthnlioqe
aunwcbzxmdvhsjfgklthnyioqe
auxwcbzrhddpsjfgklthnnioqe
ouxwcbzrmdvtsifgklthnyioqe
auxwcbzrmdqpsjfgklthnyfoqp
auxwrbzrhdvpsjfgolthnyioqe
auxwcbcqmdvpsjugklthnyioqe
auxwcbzrqdvpsjhgklthnjioqe
auxmcbzrmdvpsjfgmlthnyjoqe
auxwcbzrmdvpsjfgzlthnycoqv
auswcbzrmdvpsffgslthnyioqe
auxwcbzrfdvpsjfrmlthnyioqe
auxwcbzrmdvpsjngzlthnxioqe
auxwcbzrmdvpsjfuqlthnyiyqe
auxwzbzrrdvosjfgklthnyioqe
auxwcbzdmdvpsjfikxthnyioqe
guxwcbzrmdvpsjfgmlthnytoqe
auxwcbzrmdvpspfgkytenyioqe
auxvcbzrldvpsjfgklthnyhoqe
auxwcbzrmavpckfgklthnyioqe
autwcbzrmdvpsafgklthnyirqe
auxwcbzrxuvpsjfgklthmyioqe
auxwcbarmdppsjfgklthnywoqe
anxvcbzrmdvpsjfgklthnyijqe
auxwcbwrmdapsjngklthnyioqe
abxwcbzrmdvpsjugkltgnyioqe
auxwcbtrmdvpsjfgkltunyioue
aujwcbzrmovpsjfgklthryioqe
auxwcbzrydvpsjfgklthndikqe
auxwcbzrmdvpsjfgklmrnyioqo
auxwcbzrddvpsjfggithnyioqe
auxwcbzrmdvpfjfaklthlyioqe
fuxtcbzrmdvpsjfgklwhnyioqe
tuxwcbzrjdvpsjfgjlthnyioqe
auxwcbzrmdppsofgklthnyfoqe
auxvclzamdvpsjfgklthnyioqe
auxwcbzrmdvpsjfdklhhnzioqe
auxwcbzrmsvpsvdgklthnyioqe
arxfcbzrmdvpsvfgklthnyioqe
auxzcbzrmdvpsjfgklthnhioqj
auxwcbzrrdvpsjfgpltunyioqe
auxuibzrmdvpwjfgklthnyioqe
auxwcbzrwdqpsjfgklthnyooqe
aujwcbzrmdvpsjvgklthxyioqe
abxwcbzrmfvpsjfgklthnyxoqe
aurwcbzrmdvpshfgklthnyhoqe
auxwcbzjmdvpsjfgknthnycoqe
auxwcbzrmdvpsjfgklmhxwioqe
auxwcbzrmfvpsjfgklthnyiorq
auxwcbormdvpsjfgklwhnlioqe
auxwctzrmdvpsjfgklcknyioqe
awxwcbzrmdvpsjfgvlthnyiome
auxwcbzrmdvpsjfjklthnyixje
auxwcsxrmdvpsjfgkltsnyioqe
auxbmbzrmdvpsjfgklthnyioce
auxwcbzrmdvpsjfukzthnytoqe
aixwcbzrmdvpsjfgllthdyioqe
auxwcbzrmdypsjfgklthnlioqy
auxccbzrmdvpsjfgkltrnnioqe
auxwcznrmdvpsjfgklthnykoqe
auxwmqzrmdvpsjfgilthnyioqe
auxwcbzrmdvpdyfgolthnyioqe
auxwcbzrmdvpsjfgkmohnqioqe
auxwcfzrmzvpsjfoklthnyioqe
auxwjyzrmdvpsjfgulthnyioqe
auxwcgzredvpsjfgkxthnyioqe
wuxwcbtrmdvpsjfgklthnyiofe
auxwcbzrmdopsgfgklthncioqe
auxmcbzjmdvpsjfgklbhnyioqe
auxwlbzrmdvpsjffklthgyioqe
auxwcbzrmrvpsjfgqlthtyioqe
kuxwhbzrmdvpsjfgklthgyioqe
auxwcozrmdgpsjfgklthnydoqe
auxwdbzrmdvpdjfgklthgyioqe
auxwqbzrmdapsvfgklthnyioqe
auqwcbzridvjsjfgklthnyioqe
auxwckzrmdvpsjfoklthnyuoqe
auxwcbzvmdvpsjfgklghnyiome
auxtcbzrmdvpsjqgktthnyioqe
auxwcbzrmdvesjfgkljhnnioqe
auxwcbzrmpvpsqfgklthnqioqe
auxwcbzrmdcpsqfgklthnzioqe
yuxwcbzrmdvpsjggklthnlioqe
auxwcbzradvpsjftklthoyioqe
auxwcbzrmdvjujfgklmhnyioqe
auxwcbzrmdvpsrfgklpinyioqe
auxwobzrvqvpsjfgklthnyioqe`
