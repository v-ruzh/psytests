package main

//This program was created to calculate the results of a number of the assesment tests in order to evaluate the level of creativity and
//self-realization of the participants. It was created as a part of the study on the relation between the two features of a person.
//Данная программа подсчитывает результаты психологических опросников, направленных на оценку креативности и
//уровня самореализации испытуемого. Программа была создана как часть исследования о взаимодействии этих двух аспектов личности.
import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Johnson function calculates the results of the Johnson's Creativity Inventory.
// Функция Johnson считает результаты опросника креативности Джонсона.
func Johnson(a []string) string {
	Sum := 0
	for _, V := range a {
		Num, _ := strconv.Atoi(V)
		Sum += Num
	}
	JustSum := "Результат первого опросника:" + strconv.Itoa(Sum)
	return JustSum
}

// Williams function calculates the results of one of the tests from Frank Williams'Creativity Assessment Packet.
// Функция Williams считает результаты одного из тестов из батареи тестов Вильямса.
func Williams(a []string) string {
	var risk int
	var curiosity int
	var imagination int
	var complexity int
	var yes string = "Да, в основном верно."
	var no string = "Нет, в основном неверно."
	var maybe string = "Может быть, отчасти верно."
	var dontKnow string = "Не знаю, не могу решить."
	for i, v := range a {

		//1+, 21+, 25+, 35+, 36+, 43+, 44+,
		//5-, 8-, 22-, 29-, 32-, 34-,
		if ((((i + 1) == 1) || ((i + 1) == 21) || ((i + 1) == 25) || ((i + 1) == 35) || ((i + 1) == 36) || ((i + 1) == 43) || ((i + 1) == 44)) && (v == yes)) || ((((i + 1) == 5) || ((i + 1) == 8) || ((i + 1) == 22) || ((i + 1) == 29) || ((i + 1) == 32) || ((i + 1) == 34)) && (v == no)) {
			risk += 2
		} else if (((i+1) == 1) || ((i+1) == 21) || ((i+1) == 25) || ((i+1) == 35) || ((i+1) == 36) || ((i+1) == 43) || ((i+1) == 44)) && ((v == no) || (v == maybe)) || ((((i + 1) == 5) || ((i + 1) == 8) || ((i + 1) == 22) || ((i + 1) == 29) || ((i + 1) == 32) || ((i + 1) == 34)) && ((v == yes) || (v == maybe))) {
			risk += 1
		} else if ((i+1) == 1 || (i+1) == 21 || (i+1) == 25 || (i+1) == 35 || (i+1) == 36 || (i+1) == 43 || (i+1) == 44 || (i+1) == 5 || (i+1) == 8 || (i+1) == 22 || (i+1) == 29 || (i+1) == 32 || (i+1) == 34) && v == dontKnow {
			risk -= 1
		}
		//2+, 3+, 11+, 12+, 19+, 27+, 33+, 37+, 38+, 47+, 49+
		//28-,
		if (((i+1) == 2 || (i+1) == 3 || (i+1) == 11 || (i+1) == 12 || (i+1) == 19 || (i+1) == 27 || (i+1) == 33 || (i+1) == 37 || (i+1) == 38 || (i+1) == 47 || (i+1) == 49) && v == yes) || ((i+1) == 28 && v == no) {
			curiosity += 2
		} else if (((i+1) == 2 || (i+1) == 3 || (i+1) == 11 || (i+1) == 12 || (i+1) == 19 || (i+1) == 27 || (i+1) == 33 || (i+1) == 37 || (i+1) == 38 || (i+1) == 47 || (i+1) == 49) && (v == no || v == maybe)) || ((i+1) == 28 && (v == yes || v == maybe)) {
			curiosity += 1
		} else if ((i+1) == 2 || (i+1) == 3 || (i+1) == 11 || (i+1) == 12 || (i+1) == 19 || (i+1) == 27 || (i+1) == 33 || (i+1) == 37 || (i+1) == 38 || (i+1) == 47 || (i+1) == 49 || (i+1) == 28) && v == dontKnow {
			curiosity -= 1
		}
		//6+, 13+, 16+, 23+, 30+, 31+, 40+, 45+, 46+
		//14-, 20-, 39-,
		if (((i+1) == 6 || (i+1) == 13 || (i+1) == 16 || (i+1) == 23 || (i+1) == 30 || (i+1) == 31 || (i+1) == 40 || (i+1) == 45 || (i+1) == 46) && v == yes) || (((i+1) == 14 || (i+1) == 20 || (i+1) == 39) && v == no) {
			imagination += 2
		} else if (((i+1) == 6 || (i+1) == 13 || (i+1) == 16 || (i+1) == 23 || (i+1) == 30 || (i+1) == 31 || (i+1) == 40 || (i+1) == 45 || (i+1) == 46) && (v == no || v == maybe)) || (((i+1) == 14 || (i+1) == 20 || (i+1) == 39) && (v == yes || v == maybe)) {
			imagination += 1
		} else if ((i+1) == 6 || (i+1) == 13 || (i+1) == 16 || (i+1) == 23 || (i+1) == 30 || (i+1) == 31 || (i+1) == 40 || (i+1) == 45 || (i+1) == 46 || (i+1) == 14 || (i+1) == 20 || (i+1) == 39) && v == dontKnow {
			imagination -= 1
		}
		//7+, 10+, 15+, 18+, 26+, 42+, 50+
		//4-, 9-, 17-, 24-, 41-, 48-,
		if (((i+1) == 7 || (i+1) == 10 || (i+1) == 15 || (i+1) == 18 || (i+1) == 26 || (i+1) == 42 || (i+1) == 50) && v == yes) || (((i+1) == 4 || (i+1) == 9 || (i+1) == 17 || (i+1) == 24 || (i+1) == 41 || (i+1) == 48) && v == no) {
			complexity += 2
		} else if (((i+1) == 7 || (i+1) == 10 || (i+1) == 15 || (i+1) == 18 || (i+1) == 26 || (i+1) == 42 || (i+1) == 50) && (v == no || v == maybe)) || (((i+1) == 4 || (i+1) == 9 || (i+1) == 17 || (i+1) == 24 || (i+1) == 41 || (i+1) == 48) && (v == yes || v == maybe)) {
			complexity += 1
		} else if ((i+1) == 7 || (i+1) == 10 || (i+1) == 15 || (i+1) == 18 || (i+1) == 26 || (i+1) == 42 || (i+1) == 50 || (i+1) == 4 || (i+1) == 9 || (i+1) == 17 || (i+1) == 24 || (i+1) == 41 || (i+1) == 48) && v == dontKnow {
			complexity -= 1
		}
	}
	rawPoints := risk + curiosity + imagination + complexity
	result := "Результаты опросника Вильямса.\nСырые баллы:\n" + strconv.Itoa(rawPoints) + "\nРиск:\n" + strconv.Itoa(risk) + "\nЛюбопытство:\n" + strconv.Itoa(curiosity) + "\nВоображение:\n" + strconv.Itoa(imagination) + "\nСложность:\n" + strconv.Itoa(complexity)
	return result
}

// Leontiev function calculates the results of the Leontiev's life-meaning orientations test.
// Функция Leontiev считает результаты теста смысложизненных ориентаций Д. А. Леонтьева.
func Leontiev(a []string) string {
	var goals, lifeProcess, lifeResults, locCtrlSelf, locCtrlLife, meaningfulness int
	for i, v := range a {
		num, _ := strconv.Atoi(v)
		meaningfulness += num
		//Субшкала 1 (цели в жизни) - 3, 4, 10, 16, 17, 18.
		if ((i + 1) == 3) || ((i + 1) == 4) || ((i + 1) == 10) || ((i + 1) == 16) || ((i + 1) == 17) || ((i + 1) == 18) {
			goals += num
			//Субшкала 2 (процесс жизни) — 1, 2, 4, 5, 7, 9.
		} else if ((i + 1) == 1) || ((i + 1) == 2) || ((i + 1) == 4) || ((i + 1) == 5) || ((i + 1) == 7) || ((i + 1) == 9) {
			lifeProcess += num
			//Субшкала З (результат жизни) — 8, 9, 10, 12, 20.
		} else if ((i + 1) == 8) || ((i + 1) == 9) || ((i + 1) == 10) || ((i + 1) == 12) || ((i + 1) == 20) {
			lifeResults += num
			//Субшкала 4 (локус контроля — Я) — 1, 15, 16, 19.
		} else if ((i + 1) == 1) || ((i + 1) == 15) || ((i + 1) == 16) || ((i + 1) == 19) {
			locCtrlSelf += num
			//Субшкала 5 (локус контроля — жизнь) — 7, 10, 11, 14, 18, 19.
		} else if ((i + 1) == 7) || ((i + 1) == 10) || ((i + 1) == 11) || ((i + 1) == 14) || ((i + 1) == 18) || ((i + 1) == 19) {
			locCtrlLife += num
		}

	}
	leontiev := "Результат опросника Леонтьева.\nОбщий балл(осмысленность жизни):\n" + strconv.Itoa(meaningfulness) + "\nСубшкала 1 (цели в жизни):\n" + strconv.Itoa(goals) + "\nСубшкала 2 (процесс жизни):\n" + strconv.Itoa(lifeProcess) + "\nСубшкала 3 (результативность жизни):\n" + strconv.Itoa(lifeResults) + "\nСубшкала 4 (локус контроля - Я):\n" + strconv.Itoa(locCtrlSelf) + "\nСубшкала 5 (локус контроля - жизнь)\n" + strconv.Itoa(locCtrlLife)
	return leontiev
}

// Kudinov function calculates the results of the Kudinov's multidimensional self-fulfilment inventory.
// Функция Kudinov считает результаты многомерного опросника самореализации личности С. И. Кудинова.
func Kudinov(a []string) string {
	var oneSf, twoSf, threeSf, fourSf, fiveSf, sixSf, sevenSf, eightSf, nineSf, tenSf, elevenSf, twelveSf, thirteenSf, fourteenSf, fifteenSf, sixteenSf, _, oneS, twoS, threeS, fourS, fiveS, sixS, sevenS, eightS, nineS, tenS, elevenS, twelveS, thirteenS, fourteenS, fifteenS, sixteenS, _, oneP, twoP, threeP, fourP, fiveP, sixP, sevenP, eightP, nineP, tenP, elevenP, twelveP, thirteenP, fourteenP, fifteenP, sixteenP, seventeen int

	for i, v := range a {
		num, _ := strconv.Atoi(v)
		if num >= 4 {

			//ценностно-целевой компонент
			//1-Осм. Ц 1, 18, 35, 52, 69, 86
			//2-Осв. Ц 2, 19, 36, 53, 70, 87
			if ((i + 1) == 1) || ((i + 1) == 18) {
				oneSf += num
			} else if ((i + 1) == 35) || ((i + 1) == 52) {
				oneS += num
			} else if ((i + 1) == 69) || ((i + 1) == 86) {
				oneP += num
			} else if ((i + 1) == 2) || ((i + 1) == 19) {
				twoSf += num
			} else if ((i + 1) == 36) || ((i + 1) == 53) {
				twoS += num
			} else if ((i + 1) == 70) || ((i + 1) == 87) {
				twoP += num

				//динамический компонент
				//3-Э-P 3, 20, 37, 54, 71, 88
				//4-A-P 4, 21, 38, 55, 72, 89
			} else if ((i + 1) == 3) || ((i + 1) == 20) {
				threeSf += num
			} else if ((i + 1) == 37) || ((i + 1) == 54) {
				threeS += num
			} else if ((i + 1) == 71) || ((i + 1) == 88) {
				threeP += num

			} else if ((i + 1) == 4) || ((i + 1) == 21) {
				fourSf += num
			} else if ((i + 1) == 38) || ((i + 1) == 55) {
				fourS += num
			} else if ((i + 1) == 72) || ((i + 1) == 89) {
				fourP += num

				//эмоциональный ком-понент
				//5-C-T 5, 22, 39, 56, 73, 90
				//6-A-C 6, 23, 40, 57, 74, 91
			} else if ((i + 1) == 5) || ((i + 1) == 22) {
				fiveSf += num
			} else if ((i + 1) == 39) || ((i + 1) == 56) {
				fiveS += num
			} else if ((i + 1) == 73) || ((i + 1) == 90) {
				fiveP += num

			} else if ((i + 1) == 6) || ((i + 1) == 23) {
				sixSf += num
			} else if ((i + 1) == 40) || ((i + 1) == 57) {
				sixS += num
			} else if ((i + 1) == 74) || ((i + 1) == 91) {
				sixP += num

				//организационный компонент
				//7-И-H 7, 24, 41, 58, 75, 92
				//8-Э-H 8, 25, 42, 59, 76, 93
			} else if ((i + 1) == 7) || ((i + 1) == 24) {
				sevenSf += num
			} else if ((i + 1) == 41) || ((i + 1) == 58) {
				sevenS += num
			} else if ((i + 1) == 75) || ((i + 1) == 92) {
				sevenP += num

			} else if ((i + 1) == 8) || ((i + 1) == 25) {
				eightSf += num
			} else if ((i + 1) == 42) || ((i + 1) == 59) {
				eightS += num
			} else if ((i + 1) == 76) || ((i + 1) == 93) {
				eightP += num

				//мотивационный компонент
				//9-C-Ц 9, 26, 43, 60, 77, 94
				//10-Э-Ц 10, 27, 44, 61, 78, 95
			} else if ((i + 1) == 9) || ((i + 1) == 26) {
				nineSf += num
			} else if ((i + 1) == 43) || ((i + 1) == 60) {
				nineS += num
			} else if ((i + 1) == 77) || ((i + 1) == 94) {
				nineP += num

			} else if ((i + 1) == 10) || ((i + 1) == 27) {
				tenSf += num
			} else if ((i + 1) == 44) || ((i + 1) == 61) {
				tenS += num
			} else if ((i + 1) == 78) || ((i + 1) == 95) {
				tenP += num

				//когнитивный компонент
				//11-K-P 11, 28, 45, 62, 79, 96
				//12-K-H 12, 29, 46, 63, 80, 97
			} else if ((i + 1) == 11) || ((i + 1) == 28) {
				elevenSf += num
			} else if ((i + 1) == 45) || ((i + 1) == 62) {
				elevenS += num
			} else if ((i + 1) == 79) || ((i + 1) == 96) {
				elevenP += num

			} else if ((i + 1) == 12) || ((i + 1) == 29) {
				twelveSf += num
			} else if ((i + 1) == 46) || ((i + 1) == 63) {
				twelveS += num
			} else if ((i + 1) == 80) || ((i + 1) == 97) {
				twelveP += num

				//прогностический компонент
				//13-K-T 13, 30, 47, 64, 81, 98
				//14-Д-Т 14, 31, 48, 65, 82, 99
			} else if ((i + 1) == 13) || ((i + 1) == 30) {
				thirteenSf += num
			} else if ((i + 1) == 47) || ((i + 1) == 64) {
				thirteenS += num
			} else if ((i + 1) == 81) || ((i + 1) == 98) {
				thirteenP += num

			} else if ((i + 1) == 14) || ((i + 1) == 31) {
				fourteenSf += num
			} else if ((i + 1) == 48) || ((i + 1) == 65) {
				fourteenS += num
			} else if ((i + 1) == 82) || ((i + 1) == 99) {
				fourteenP += num

				//компетентно-личностный компонент
				//15-С-Б 15, 32, 49, 66, 83, 100
				//16-Л-Б 16, 33, 50, 67, 84, 101

			} else if ((i + 1) == 15) || ((i + 1) == 32) {
				fifteenSf += num
			} else if ((i + 1) == 49) || ((i + 1) == 66) {
				fifteenS += num
			} else if ((i + 1) == 83) || ((i + 1) == 100) {
				fifteenP += num

			} else if ((i + 1) == 16) || ((i + 1) == 33) {
				sixteenSf += num
			} else if ((i + 1) == 50) || ((i + 1) == 67) {
				sixteenS += num
			} else if ((i + 1) == 84) || ((i + 1) == 101) {
				sixteenP += num
				//Шкала искренности И-Т 17,34,51,68,85,102
			} else if ((i + 1) == 17) || ((i + 1) == 34) || ((i + 1) == 51) || ((i + 1) == 68) || ((i + 1) == 85) || ((i + 1) == 102) {
				seventeen += num
			}

		}
	}
	one := oneP + oneS + oneSf
	two := twoP + twoS + twoSf
	three := threeP + threeS + threeSf
	four := fourP + fourS + fourSf
	five := fiveP + fiveS + fiveSf
	six := sixP + sixS + sixSf
	seven := sevenP + sevenS + sevenSf
	eight := eightP + eightS + eightSf
	nine := nineP + nineS + nineSf
	ten := tenP + tenS + tenSf
	eleven := elevenP + elevenS + elevenSf
	twelve := twelveP + twelveS + twelveSf
	thirteen := thirteenP + thirteenS + thirteenSf
	fourteen := fourteenP + fourteenS + fourteenSf
	fifteen := fifteenP + fifteenS + fifteenSf
	sixteen := sixteenP + sixteenS + sixteenSf
	fulfilment := (one + three + five + seven + eleven + thirteen + ((nine + ten) / 2)) - (two + four + six + eight + twelve + fourteen + ((fifteen + sixteen) / 2))
	self := (oneSf + threeSf + fiveSf + sevenSf + elevenSf + thirteenSf + ((nineSf + tenSf) / 2)) - (twoSf + fourSf + sixSf + eightSf + twelveSf + fourteenSf + ((fifteenSf + sixteenSf) / 2))

	social := (oneS + threeS + fiveS + sevenS + elevenS + thirteenS + ((nineS + tenS) / 2)) - (twoS + fourS + sixS + eightS + twelveS + fourteenS + ((fifteenS + sixteenS) / 2))

	prof := (oneP + threeP + fiveP + sevenP + elevenP + thirteenP + ((nineP + tenP) / 2)) - (twoP + fourP + sixP + eightP + twelveP + fourteenP + ((fifteenP + sixteenP) / 2))
	Kudinov := "\nОпросник Кудинова.\nПервый и второй этапы.\nШкала искренности (17-И-Т):\n" + strconv.Itoa(seventeen) + "\nЦенностно-целевой компонент.\n1-Осм. Ц:\n" + strconv.Itoa(one) + "\n2-Осв. Ц:\n" + strconv.Itoa(two) + "\nДинамический компонент.\n3-Э-P:\n" + strconv.Itoa(three) + "\n4-A-P:\n" + strconv.Itoa(four) + "\nЭмоциональный компонент.\n5-C-T:\n" + strconv.Itoa(five) + "\n6-A-C:\n" + strconv.Itoa(six) + "\nОрганизационный компонент.\n7-И-H:\n" + strconv.Itoa(seven) + "\n8-Э-H:\n" + strconv.Itoa(eight) + "\nМотивационный компонент.\n9-C-Ц:\n" + strconv.Itoa(nine) + "\n10-Э-Ц\n" + strconv.Itoa(ten) + "\nКогнитивный компонент.\n11-K-P:\n" + strconv.Itoa(eleven) + "\n12-K-H:\n" + strconv.Itoa(twelve) + "\nПрогностический компонент.\n13-K-T:\n" + strconv.Itoa(thirteen) + "\n14-Д-Т:\n" + strconv.Itoa(fourteen) + "\nКомпетентно-личностный компонент.\n15-С-Б:\n" + strconv.Itoa(fifteen) + "\n16-Л-Б:\n" + strconv.Itoa(sixteen) + "\n\n\nЭтап три.\n\nБаллы самореализации:\n" + strconv.Itoa(fulfilment) + "\n\n\nЭтап 4.\nЛичностная самореализация:\n" + strconv.Itoa(self) + "\nСоциальная самореализация:\n" + strconv.Itoa(social) + "\nПрофессиональная самореализация:\n" + strconv.Itoa(prof)

	return Kudinov
}

// Cat function calculates the results of the self-actualization test (CAT is a cyrillic acronym)
// Функция Cat считает результаты самоатуализационного теста (САТ).
func Cat(s []string) string {
	var time01, supp02, val03, flex04, sense05, spon06, sfResp07, sfAccpt08, humNat09, syn10, aggr11, cont12, cur13, creat14 int
	var a string = "'А'"
	var b string = "'Б'"
	for f := range s {
		for j, l := range s[f] {
			if j == 0 {
				s[f] = strconv.QuoteRune(l)
			}
		}
	}
	for i := range s {
		//Компетентность во времени (базовая шкала, 17 вопросов). 11а, 16б, 18б, 21а, 28б, 38б, 40б, 41б, 45б, 60б, 64б, 71б, 76б, 82б, 91б, 106б, 126б.
		if (((i + 1) == 11) && (s[i] == a)) || ((i+1) == 16 && (s[i] == b)) || (((i + 1) == 18) && (s[i] == b)) || (((i + 1) == 21) && (s[i] == a)) || (((i + 1) == 28) && (s[i] == b)) || (((i + 1) == 38) && (s[i] == b)) || (((i + 1) == 40) && (s[i] == b)) || (((i + 1) == 41) && (s[i] == b)) || (((i + 1) == 45) && (s[i] == b)) || (((i + 1) == 60) && (s[i] == b)) || (((i + 1) == 64) && (s[i] == b)) || (((i + 1) == 71) && (s[i] == b)) || (((i + 1) == 76) && (s[i] == b)) || (((i + 1) == 82) && (s[i] == b)) || (((i + 1) == 91) && (s[i] == b)) || (((i + 1) == 106) && (s[i] == b)) || (((i + 1) == 126) && (s[i] == b)) {
			time01 += 1
		}
		//Шкала поддержки (базовая, 91 вопрос). 1б, 2б, 3а, 4а, 5б, 7б, 8а, 9а, 10а, 12б, 14б, 15б, 17а, 19а, 22б, 23а, 25б, 26б, 27б, 29а, 31б, 32а, 33б, 34а, 35б, 36б, 39б, 42а, 43а, 44б, 46а, 47б, 49б, 50б, 51б, 52а, 53а, 55а, 56а, 57а, 59а, 61б, 62б, 65б, 66а, 67б, 68а, 69б, 70а, 72б, 73а, 74б, 75б, 77а, 79б, 80а, 81а, 83а, 85б, 86а, 87б, 88б, 89б, 90а, 93а, 94а, 95б, 96а, 97а, 98а, 99б, 100а, 102а, 103б, 104а, 105б, 108б, 109а, 110а, 111б, 113а, 114а, 115а, 116б, 117б, 118а, 119б, 120а, 122а, 125б.
		if ((i+1) == 1) && (s[i] == b) || (((i + 1) == 2) && (s[i] == b)) || (((i + 1) == 3) && (s[i] == a)) || (((i + 1) == 4) && (s[i] == a)) || (((i + 1) == 5) && (s[i] == b)) || (((i + 1) == 7) && (s[i] == b)) || (((i + 1) == 8) && (s[i] == a)) || (((i + 1) == 9) && (s[i] == a)) || (((i + 1) == 10) && (s[i] == a)) || (((i + 1) == 12) && (s[i] == b)) || (((i + 1) == 14) && (s[i] == b)) || (((i + 1) == 15) && (s[i] == b)) || (((i + 1) == 17) && (s[i] == a)) || (((i + 1) == 19) && (s[i] == a)) || (((i + 1) == 22) && (s[i] == b)) || (((i + 1) == 23) && (s[i] == a)) || (((i + 1) == 25) && (s[i] == b)) || (((i + 1) == 26) && (s[i] == b)) || (((i + 1) == 27) && (s[i] == b)) || (((i + 1) == 29) && (s[i] == a)) || (((i + 1) == 31) && (s[i] == b)) || (((i + 1) == 32) && (s[i] == a)) || (((i + 1) == 33) && (s[i] == b)) || (((i + 1) == 34) && (s[i] == a)) || (((i + 1) == 35) && (s[i] == b)) || (((i + 1) == 36) && (s[i] == b)) || (((i + 1) == 39) && (s[i] == b)) || (((i + 1) == 42) && (s[i] == a)) || (((i + 1) == 43) && (s[i] == a)) || (((i + 1) == 44) && (s[i] == b)) || (((i + 1) == 46) && (s[i] == a)) || (((i + 1) == 47) && (s[i] == b)) || (((i + 1) == 49) && (s[49] == b)) || (((i + 1) == 50) && (s[i] == b)) || (((i + 1) == 51) && (s[i] == b)) || (((i + 1) == 52) && (s[i] == a)) || (((i + 1) == 53) && (s[i] == a)) || (((i + 1) == 55) && (s[i] == a)) || (((i + 1) == 56) && (s[i] == a)) || (((i + 1) == 57) && (s[i] == a)) || (((i + 1) == 59) && (s[i] == a)) || (((i + 1) == 61) && (s[i] == b)) || (((i + 1) == 62) && (s[i] == b)) || (((i + 1) == 65) && (s[i] == b)) || (((i + 1) == 66) && (s[i] == a)) || (((i + 1) == 67) && (s[i] == b)) || (((i + 1) == 68) && (s[i] == a)) || (((i + 1) == 69) && (s[i] == b)) || (((i + 1) == 70) && (s[i] == a)) || (((i + 1) == 72) && (s[i] == b)) || (((i + 1) == 73) && (s[i] == a)) || (((i + 1) == 74) && (s[i] == b)) || (((i + 1) == 75) && (s[i] == b)) || (((i + 1) == 77) && (s[i] == a)) || (((i + 1) == 79) && (s[i] == b)) || (((i + 1) == 80) && (s[i] == a)) || (((i + 1) == 81) && (s[i] == a)) || (((i + 1) == 83) && (s[i] == a)) || (((i + 1) == 85) && (s[i] == b)) || (((i + 1) == 86) && (s[i] == a)) || (((i + 1) == 87) && (s[i] == b)) || (((i + 1) == 88) && (s[i] == b)) || (((i + 1) == 89) && (s[i] == b)) || (((i + 1) == 90) && (s[i] == a)) || (((i + 1) == 93) && (s[i] == a)) || (((i + 1) == 94) && (s[i] == a)) || (((i + 1) == 95) && (s[i] == b)) || (((i + 1) == 96) && (s[i] == a)) || (((i + 1) == 97) && (s[i] == a)) || (((i + 1) == 98) && (s[i] == a)) || (((i + 1) == 99) && (s[i] == b)) || (((i + 1) == 100) && (s[i] == a)) || (((i + 1) == 102) && (s[i] == a)) || (((i + 1) == 103) && (s[i] == b)) || (((i + 1) == 104) && (s[i] == a)) || (((i + 1) == 105) && (s[i] == b)) || (((i + 1) == 108) && (s[i] == b)) || (((i + 1) == 109) && (s[i] == a)) || (((i + 1) == 110) && (s[i] == a)) || (((i + 1) == 111) && (s[i] == b)) || (((i + 1) == 113) && (s[i] == a)) || (((i + 1) == 114) && (s[i] == a)) || (((i + 1) == 115) && (s[i] == a)) || (((i + 1) == 116) && (s[i] == b)) || (((i + 1) == 117) && (s[i] == b)) || (((i + 1) == 118) && (s[i] == a)) || (((i + 1) == 119) && (s[i] == b)) || (((i + 1) == 120) && (s[i] == a)) || (((i + 1) == 122) && (s[i] == a)) || (((i + 1) == 125) && (s[i] == b)) {
			supp02 += 1

		}
		// Шкала ценностных ориентаций (дополнительная, 20 вопросов).17а, 29а, 42а, 49б, 50б, 53а, 56а, 59а, 67б, 68а, 69б, 80а, 81а, 90а, 93а, 97а, 99б, 113а, 114а, 122а.
		if ((i+1) == 17) && (s[i] == a) || (((i + 1) == 29) && (s[i] == a)) || (((i + 1) == 42) && (s[i] == a)) || (((i + 1) == 49) && (s[i] == b)) || (((i + 1) == 50) && (s[i] == b)) || (((i + 1) == 53) && (s[i] == a)) || (((i + 1) == 56) && (s[i] == a)) || (((i + 1) == 59) && (s[i] == a)) || (((i + 1) == 67) && (s[i] == b)) || (((i + 1) == 68) && (s[i] == a)) || (((i + 1) == 69) && (s[i] == b)) || (((i + 1) == 80) && (s[i] == a)) || (((i + 1) == 81) && (s[i] == a)) || (((i + 1) == 90) && (s[i] == a)) || (((i + 1) == 93) && (s[i] == a)) || (((i + 1) == 97) && (s[i] == a)) || (((i + 1) == 99) && (s[i] == b)) || (((i + 1) == 113) && (s[i] == a)) || (((i + 1) == 114) && (s[i] == a)) || (((i + 1) == 122) && (s[i] == a)) {
			val03 += 1

		}
		//Шкала гибкости поведения (дополнительная, 24 вопроса). 3а, 9а, 12б, 33б, 36б, 38б, 40б, 47б, 50б, 51б, 61б, 62б, 65б, 68а, 70а, 74б, 82б, 85б, 95б, 97а, 99б, 102а, 105б, 123б.
		if ((i+1) == 3) && (s[i] == a) || (((i + 1) == 9) && (s[i] == a)) || (((i + 1) == 12) && (s[i] == b)) || (((i + 1) == 33) && (s[i] == b)) || (((i + 1) == 36) && (s[i] == b)) || (((i + 1) == 38) && (s[i] == b)) || (((i + 1) == 40) && (s[i] == b)) || (((i + 1) == 47) && (s[i] == b)) || (((i + 1) == 50) && (s[i] == b)) || (((i + 1) == 51) && (s[i] == b)) || (((i + 1) == 61) && (s[i] == b)) || (((i + 1) == 62) && (s[i] == b)) || (((i + 1) == 65) && (s[i] == b)) || (((i + 1) == 68) && (s[i] == a)) || (((i + 1) == 70) && (s[i] == a)) || (((i + 1) == 74) && (s[i] == b)) || (((i + 1) == 82) && (s[i] == b)) || (((i + 1) == 85) && (s[i] == b)) || (((i + 1) == 95) && (s[i] == b)) || (((i + 1) == 97) && (s[i] == a)) || (((i + 1) == 99) && (s[i] == b)) || (((i + 1) == 102) && (s[i] == a)) || (((i + 1) == 105) && (s[i] == b)) || (((i + 1) == 123) && (s[i] == b)) {
			flex04 += 1

		}
		//Шкала сензитивности (дополнительная, 13 вопросов). 2б, 5б, 10а, 43а, 46а, 55а, 73а, 77а, 83а, 89б, 103б, 119б, 122а.
		if ((i+1) == 2) && (s[i] == b) || (((i + 1) == 5) && (s[i] == b)) || (((i + 1) == 10) && (s[i] == a)) || (((i + 1) == 43) && (s[i] == a)) || (((i + 1) == 46) && (s[i] == a)) || (((i + 1) == 55) && (s[i] == a)) || (((i + 1) == 73) && (s[i] == a)) || (((i + 1) == 77) && (s[i] == a)) || (((i + 1) == 83) && (s[i] == a)) || (((i + 1) == 89) && (s[i] == b)) || (((i + 1) == 103) && (s[i] == b)) || (((i + 1) == 119) && (s[i] == b)) || (((i + 1) == 122) && (s[i] == a)) {
			sense05 += 1

		}
		// Шкала спонтанности (дополнительная, 14 вопросов). 5б, 14б, 15б, 26б, 42а, 62а, 67б, 74б, 77а, 80а, 81а, 83а, 95б, 114а.
		if ((i+1) == 5) && (s[i] == b) || (((i + 1) == 14) && (s[i] == b)) || (((i + 1) == 15) && (s[i] == b)) || (((i + 1) == 26) && (s[i] == b)) || (((i + 1) == 42) && (s[i] == a)) || (((i + 1) == 62) && (s[i] == a)) || (((i + 1) == 67) && (s[i] == b)) || (((i + 1) == 74) && (s[i] == b)) || (((i + 1) == 77) && (s[i] == a)) || (((i + 1) == 80) && (s[i] == a)) || (((i + 1) == 81) && (s[i] == a)) || (((i + 1) == 83) && (s[i] == a)) || (((i + 1) == 95) && (s[i] == b)) || (((i + 1) == 114) && (s[i] == a)) {
			spon06 += 1

		}
		//Шкала самоуважения (дополнительная, 15 вопросов).2б, 3а, 7б, 23а, 29а, 44б, 53а, 66а, 69б, 98а, 100а, 102а, 106б, 114а, 122а.
		if ((i+1) == 2) && (s[i] == b) || (((i + 1) == 3) && (s[i] == a)) || (((i + 1) == 7) && (s[i] == b)) || (((i + 1) == 23) && (s[i] == a)) || (((i + 1) == 29) && (s[i] == a)) || (((i + 1) == 44) && (s[i] == b)) || (((i + 1) == 53) && (s[i] == a)) || (((i + 1) == 66) && (s[i] == a)) || (((i + 1) == 69) && (s[i] == b)) || (((i + 1) == 98) && (s[i] == a)) || (((i + 1) == 100) && (s[i] == a)) || (((i + 1) == 102) && (s[i] == a)) || (((i + 1) == 106) && (s[i] == b)) || (((i + 1) == 114) && (s[i] == a)) || (((i + 1) == 122) && (s[i] == a)) {
			sfResp07 += 1

		}
		// Шкала самопринятия (дополнительная, 21 вопрос).1б, 8а, 14б, 22б, 31б, 32а, 34а, 39б, 53а, 61б, 71б, 75б, 86а, 87б, 104а, 105б, 106б, 110б, 111б, 116б, 125б.
		if ((i+1) == 1) && (s[i] == b) || (((i + 1) == 8) && (s[i] == a)) || (((i + 1) == 14) && (s[i] == b)) || (((i + 1) == 22) && (s[i] == b)) || (((i + 1) == 31) && (s[i] == b)) || (((i + 1) == 32) && (s[i] == a)) || (((i + 1) == 34) && (s[i] == a)) || (((i + 1) == 39) && (s[i] == b)) || (((i + 1) == 53) && (s[i] == a)) || (((i + 1) == 61) && (s[i] == b)) || (((i + 1) == 71) && (s[i] == b)) || (((i + 1) == 75) && (s[i] == b)) || (((i + 1) == 86) && (s[i] == a)) || (((i + 1) == 87) && (s[i] == b)) || (((i + 1) == 104) && (s[i] == a)) || (((i + 1) == 105) && (s[i] == b)) || (((i + 1) == 106) && (s[i] == b)) || (((i + 1) == 110) && (s[i] == b)) || (((i + 1) == 111) && (s[i] == b)) || (((i + 1) == 116) && (s[i] == b)) || (((i + 1) == 125) && (s[i] == b)) {
			sfAccpt08 += 1

		}
		// Шкала представлений о природе человека (дополнительная, 10 вопросов). 23а, 25б, 27б, 50б, 66а, 90а, 94а, 97а, 99б, 113а.
		if ((i+1) == 23) && (s[i] == a) || (((i + 1) == 25) && (s[i] == b)) || (((i + 1) == 27) && (s[i] == b)) || (((i + 1) == 50) && (s[i] == b)) || (((i + 1) == 66) && (s[i] == a)) || (((i + 1) == 90) && (s[i] == a)) || (((i + 1) == 94) && (s[i] == a)) || (((i + 1) == 97) && (s[i] == a)) || (((i + 1) == 99) && (s[i] == b)) || (((i + 1) == 113) && (s[i] == a)) {
			humNat09 += 1

		}
		//  Шкала синергичности (дополнительная, 7 вопросов).50б, 68а, 91б, 93а, 97а, 99б, 113а.
		if ((i+1) == 50) && (s[i] == b) || (((i + 1) == 68) && (s[i] == a)) || (((i + 1) == 91) && (s[i] == b)) || (((i + 1) == 93) && (s[i] == a)) || (((i + 1) == 97) && (s[i] == a)) || (((i + 1) == 99) && (s[i] == b)) || (((i + 1) == 113) && (s[i] == a)) {
			syn10 += 1

		}
		//Шкала принятия агрессии (дополнительная, 16 вопросов). 5б, 8а, 10а, 15б, 19а, 29а, 39а, 43а, 46а, 56а, 57б, 67б, 85б, 93а, 94а, 115а.
		if ((i+1) == 5) && (s[i] == b) || (((i + 1) == 8) && (s[i] == a)) || (((i + 1) == 10) && (s[i] == a)) || (((i + 1) == 15) && (s[i] == b)) || (((i + 1) == 19) && (s[i] == a)) || (((i + 1) == 29) && (s[i] == a)) || (((i + 1) == 39) && (s[i] == a)) || (((i + 1) == 43) && (s[i] == a)) || (((i + 1) == 46) && (s[i] == a)) || (((i + 1) == 56) && (s[i] == a)) || (((i + 1) == 57) && (s[i] == b)) || (((i + 1) == 67) && (s[i] == b)) || (((i + 1) == 85) && (s[i] == b)) || (((i + 1) == 93) && (s[i] == a)) || (((i + 1) == 94) && (s[i] == a)) || (((i + 1) == 115) && (s[i] == a)) {
			aggr11 += 1

		}
		//Шкала контактности (дополнительная, 20 вопросов).17а, 5б, 7б, 23б, 26б, 36б, 46а, 65б, 70а, 73а, 74б, 75б, 79б, 96а,99б, 103б, 108б, 109а, 120а, 123б.
		if ((i+1) == 17) && (s[i] == a) || (((i + 1) == 5) && (s[i] == b)) || (((i + 1) == 7) && (s[i] == b)) || (((i + 1) == 23) && (s[i] == b)) || (((i + 1) == 26) && (s[i] == b)) || (((i + 1) == 36) && (s[i] == b)) || (((i + 1) == 46) && (s[i] == a)) || (((i + 1) == 65) && (s[i] == b)) || (((i + 1) == 70) && (s[i] == a)) || (((i + 1) == 73) && (s[i] == a)) || (((i + 1) == 74) && (s[i] == b)) || (((i + 1) == 75) && (s[i] == b)) || (((i + 1) == 79) && (s[i] == b)) || (((i + 1) == 96) && (s[i] == a)) || (((i + 1) == 99) && (s[i] == b)) || (((i + 1) == 103) && (s[i] == b)) || (((i + 1) == 108) && (s[i] == b)) || (((i + 1) == 109) && (s[i] == a)) || (((i + 1) == 120) && (s[i] == a)) || (((i + 1) == 123) && (s[i] == b)) {
			cont12 += 1

		}
		//Шкала познавательных потребностей (дополнительная, 11 вопросов).13а, 20б, 37а, 48а, 63б, 66а, 78б, 82б, 92а, 107б, 121б.
		if ((i+1) == 13) && (s[i] == a) || (((i + 1) == 20) && (s[i] == b)) || (((i + 1) == 37) && (s[i] == a)) || (((i + 1) == 48) && (s[i] == a)) || (((i + 1) == 63) && (s[i] == b)) || (((i + 1) == 66) && (s[i] == a)) || (((i + 1) == 78) && (s[i] == b)) || (((i + 1) == 82) && (s[i] == b)) || (((i + 1) == 92) && (s[i] == a)) || (((i + 1) == 107) && (s[i] == b)) || (((i + 1) == 121) && (s[i] == b)) {
			cur13 += 1

		}
		//Шкала креативности (дополнительная, 14 вопросов).6б, 24а, 30а, 42а, 54а, 58а, 59а, 68а, 84а, 101а, 105б, 112б, 123б, 124б.
		if ((i+1) == 6) && (s[i] == b) || (((i + 1) == 24) && (s[i] == a)) || (((i + 1) == 30) && (s[i] == a)) || (((i + 1) == 42) && (s[i] == a)) || (((i + 1) == 54) && (s[i] == a)) || (((i + 1) == 58) && (s[i] == a)) || (((i + 1) == 59) && (s[i] == a)) || (((i + 1) == 68) && (s[i] == a)) || (((i + 1) == 84) && (s[i] == a)) || (((i + 1) == 101) && (s[i] == a)) || (((i + 1) == 105) && (s[i] == b)) || (((i + 1) == 112) && (s[i] == b)) || (((i + 1) == 123) && (s[i] == b)) || (((i + 1) == 124) && (s[i] == b)) {
			creat14 += 1

		}
	}
	ttime01 := float64(time01) / 17 * 100
	tsupp02 := float64(supp02) / 91 * 100
	tval03 := float64(val03) / 20 * 100
	tflex04 := float64(flex04) / 24 * 100
	tsense05 := float64(sense05) / 13 * 100
	tspon06 := float64(spon06) / 14 * 100
	tsfResp07 := float64(sfResp07) / 15 * 100
	tsfAccpt08 := float64(sfAccpt08) / 21 * 100
	thumNat09 := float64(humNat09) / 10 * 100
	tsyn10 := float64(syn10) / 7 * 100
	taggr11 := float64(aggr11) / 16 * 100
	tcont12 := float64(cont12) / 20 * 100
	tcur13 := float64(cur13) / 11 * 100
	tcreat14 := float64(creat14) / 14 * 100

	results := "Баллы САТ. Сначала указаны сырые, затем Т-баллы.\n\n" + "Компетентность во времени:  " + strconv.Itoa(time01) + " " + strconv.FormatFloat(ttime01, 'f', 1, 64) + "\nШкала поддержки: " + strconv.Itoa(supp02) + " " + strconv.FormatFloat(tsupp02, 'f', 1, 64) + "\nШкала ценностных ориентаций: " + strconv.Itoa(val03) + " " + strconv.FormatFloat(tval03, 'f', 1, 64) + "\nШкала гибкости поведения: " + strconv.Itoa(flex04) + " " + strconv.FormatFloat(tflex04, 'f', 1, 64) + "\nШкала сензитивности: " + strconv.Itoa(sense05) + " " + strconv.FormatFloat(tsense05, 'f', 1, 64) + "\nШкала спонтанности: " + strconv.Itoa(spon06) + " " + strconv.FormatFloat(tspon06, 'f', 1, 64) + "\nШкала самоуважения: " + strconv.Itoa(sfResp07) + " " + strconv.FormatFloat(tsfResp07, 'f', 1, 64) + "\nШкала самопринятия: " + strconv.Itoa(sfAccpt08) + " " + strconv.FormatFloat(tsfAccpt08, 'f', 1, 64) + "\nШкала представлений о природе человека: " + strconv.Itoa(humNat09) + " " + strconv.FormatFloat(thumNat09, 'f', 1, 64) + "\nШкала синергичности: " + strconv.Itoa(syn10) + " " + strconv.FormatFloat(tsyn10, 'f', 1, 64) + "\nШкала принятия агрессии: " + strconv.Itoa(aggr11) + " " + strconv.FormatFloat(taggr11, 'f', 1, 64) + "\nШкала контактности: " + strconv.Itoa(cont12) + " " + strconv.FormatFloat(tcont12, 'f', 1, 64) + "\nШкала познавательных потребностей: " + strconv.Itoa(cur13) + " " + strconv.FormatFloat(tcur13, 'f', 1, 64) + "\nШкала креативности: " + strconv.Itoa(creat14) + " " + strconv.FormatFloat(tcreat14, 'f', 1, 64)
	return results
}

// check function checks if there's an error in program's execution and if there is one it generates a message for the user.
// Функция check проверяет выполнение программы на наличие ошибки и генерирует сообщение для пользователя в случае её присутствия.
func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

// createSlice function creates a slice of strings with all of the content of the .csv file used for further calculation, an example of such file is in the /main folder of this project which is called "test.csv". Current edition of the program requires 46 rows of content filled in accorrding to the example. The number of rows can be changed by reducing its number in the given code.
// Функция createSlice создает срез строк из .csv файла, используемого для рассчетов. Пример такого файла находится в папке /main под названием "test.csv". Настоящая версия программы требует наличия 46 строк в файле оформленных в соответствии с примером. При необходимости можно удалить обработку лишних строк в тексте программы.
func createSlice() []string {
	// rec := make([]string, 1300)
	//  pointR := &rec
	//var err error
	var rec []string
	f, err := os.Open("answers2903.csv")
	check(err)
	defer f.Close()
	csvReader := csv.NewReader(f)
	for {
		pointR, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		rec = append(rec, pointR...)
	}
	return rec
}

// splitAll function splits the initial raw slice of data raw by raw.
// Функция splitAll разделяет изначальный срез данных из таблицы на построчные подсрезы для дальнейшей работы.
func splitAll(a []string) ([]string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string) {
	headerRow, firstRow, secondRow, thirdRow, fourthRow, fifthRow, sixthRow, seventhRow, ninethRow, eightsRow, tenthRow, eleventhRow, twelvethRow, thirteenthRow, fourteenthRow, fifteenthRow, sixteenthRow, seventeenthRow, eighteenthRow, nineteenthRow, twenteenthRow, twentyfirstRow, twentysecondRow, twentythirdRow, twentyfourthRow, twentyfifthRow, twentysixthRow, twentyseventhRow, twentyeighthRow, twentyninthRow, thirtiethRow, thirtyfirstRow, thirtysecondRow, thirtythirdRow, thirtyfourthRow, thirtyfifthRow, thirtysixthRow, thirtyseventhRow, thirtyeightsRow, thirtyninethRow, fourtiethRow, fourtyfirstRow, fourtysecondRow, fourtythirdRow, fourtyfourthRow, fourtyfifthRow := make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000), make([]string, 2000)
	headerRow = a[:317]
	firstRow = a[318:636]
	secondRow = a[636:954]
	thirdRow = a[954:1272]
	fourthRow = a[1272:1590] //+318
	fifthRow = a[1590:1908]
	sixthRow = a[1908:2226]
	seventhRow = a[2226:2544]
	eightsRow = a[2544:2862]
	ninethRow = a[2862:3180]
	tenthRow = a[3180:3498]
	eleventhRow = a[3498:3816]
	twelvethRow = a[3816:4134]
	thirteenthRow = a[4134:4452]
	fourteenthRow = a[4452:4770]
	fifteenthRow = a[4770:5088]
	sixteenthRow = a[5088:5406]
	seventeenthRow = a[5406:5724]
	eighteenthRow = a[5724:6042]
	nineteenthRow = a[6042:6360]
	twenteenthRow = a[6360:6678]
	twentyfirstRow = a[6678:6996]
	twentysecondRow = a[6996:7314]
	twentythirdRow = a[7314:7632]
	twentyfourthRow = a[7632:7950]
	twentyfifthRow = a[7950:8268]
	twentysixthRow = a[8268:8586]
	twentyseventhRow = a[8586:8904]
	twentyeighthRow = a[8904:9222]
	twentyninthRow = a[9222:9540]
	thirtiethRow = a[9540:9858]
	thirtyfirstRow = a[9858:10176]
	thirtysecondRow = a[10176:10494]
	thirtythirdRow = a[10494:10812]
	thirtyfourthRow = a[10812:11130]
	thirtyfifthRow = a[11130:11448]
	thirtysixthRow = a[11448:11766]
	thirtyseventhRow = a[11766:12084]
	thirtyeightsRow = a[12084:12402]
	thirtyninethRow = a[12402:12720]
	fourtiethRow = a[12720:13038]
	fourtyfirstRow = a[13038:13356]
	fourtysecondRow = a[13356:13674]
	fourtythirdRow = a[13674:13992]
	fourtyfourthRow = a[13992:14310]
	fourtyfifthRow = a[14310:14628]

	return headerRow, firstRow, secondRow, thirdRow, fourthRow, fifthRow, sixthRow, seventhRow, ninethRow, eightsRow, tenthRow, eleventhRow, twelvethRow, thirteenthRow, fourteenthRow, fifteenthRow, sixteenthRow, seventeenthRow, eighteenthRow, nineteenthRow, twenteenthRow, twentyfirstRow, twentysecondRow, twentythirdRow, twentyfourthRow, twentyfifthRow, twentysixthRow, twentyseventhRow, twentyeighthRow, twentyninthRow, thirtiethRow, thirtyfirstRow, thirtysecondRow, thirtythirdRow, thirtyfourthRow, thirtyfifthRow, thirtysixthRow, thirtyseventhRow, thirtyeightsRow, thirtyninethRow, fourtiethRow, fourtyfirstRow, fourtysecondRow, fourtythirdRow, fourtyfourthRow, fourtyfifthRow
}

// splitSlice function splits a given row into a number of slices for each analytical method as well as separates the personal data part from the test answers.
// Функция splitSlice разделяет срез строки на подсрезы, содержащие отдельно личные данные и каждую из методик для дальнейших подсчетов.
func splitSlice(a []string) (personalData, johnsonP02, williamsP03, leontievP04, kudinovP05, catP06 []string) {
	personalData = a[0:11] //1
	johnsonP02 = a[11:19]  //2
	williamsP03 = a[19:69] //3
	leontievP04 = a[69:89] //4
	kudinovP05 = a[89:191] //5
	catP06 = a[191:317]    //6
	return
}

// contacts function creates a separate slice of data with participant's contact information
// Функция contacts создает отдельный срез с контактной информацией участника исследования.
func contacts(a []string) []string {
	contact := a[317:]
	return contact
}

// addDataToFile function adds a slice of strings to the results.txt file.
// Функция addDataToFile добавляет срез строк к файлу results.txt.
func addDataToFile(a []string) {
	filename := "results.txt"
	fi, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer fi.Close()
	check(err)
	fi.WriteString("\n")
	for i, _ := range a {
		_, err := fi.WriteString(a[i] + "\n")
		check(err)
	}
}

func main() {

	raw := createSlice()

	/*headerRow*/
	_, firstRow, secondRow, thirdRow, fourthRow, fifthRow, sixthRow, seventhRow, eightsRow, ninethRow, tenthRow, eleventhRow, twelvethRow, thirteenthRow, fourteenthRow, fifteenthRow, sixteenthRow, seventeenthRow, eighteenthRow, nineteenthRow, twenteenthRow, twentyfirstRow, twentysecondRow, twentythirdRow, twentyfourthRow, twentyfifthRow, twentysixthRow, twentyseventhRow, twentyeighthRow, twentyninthRow, thirtiethRow, thirtyfirstRow, thirtysecondRow, thirtythirdRow, thirtyfourthRow, thirtyfifthRow, thirtysixthRow, thirtyseventhRow, thirtyeightsRow, thirtyninethRow, fourtiethRow, fourtyfirstRow, fourtysecondRow, fourtythirdRow, fourtyfourthRow, fourtyfifthRow := splitAll(raw)
	//personalDataHead, johnsonHead, williamsHead, leontievHead, kudinovHead, catHead := splitSlicel(headerRow)
	personalData01, johnson01, williams01, leontiev01, kudinov01, cat01 := splitSlice(firstRow)
	personalData02, johnson02, williams02, leontiev02, kudinov02, cat02 := splitSlice(secondRow)
	personalData03, johnson03, williams03, leontiev03, kudinov03, cat03 := splitSlice(thirdRow)
	personalData04, johnson04, williams04, leontiev04, kudinov04, cat04 := splitSlice(fourthRow)
	personalData05, johnson05, williams05, leontiev05, kudinov05, cat05 := splitSlice(fifthRow)
	personalData06, johnson06, williams06, leontiev06, kudinov06, cat06 := splitSlice(sixthRow)
	personalData07, johnson07, williams07, leontiev07, kudinov07, cat07 := splitSlice(seventhRow)
	personalData08, johnson08, williams08, leontiev08, kudinov08, cat08 := splitSlice(eightsRow)
	personalData09, johnson09, williams09, leontiev09, kudinov09, cat09 := splitSlice(ninethRow)
	personalData10, johnson10, williams10, leontiev10, kudinov10, cat10 := splitSlice(tenthRow)
	personalData11, johnson11, williams11, leontiev11, kudinov11, cat11 := splitSlice(eleventhRow)
	personalData12, johnson12, williams12, leontiev12, kudinov12, cat12 := splitSlice(twelvethRow)
	personalData13, johnson13, williams13, leontiev13, kudinov13, cat13 := splitSlice(thirteenthRow)
	personalData14, johnson14, williams14, leontiev14, kudinov14, cat14 := splitSlice(fourteenthRow)
	personalData15, johnson15, williams15, leontiev15, kudinov15, cat15 := splitSlice(fifteenthRow)
	personalData16, johnson16, williams16, leontiev16, kudinov16, cat16 := splitSlice(sixteenthRow)
	personalData17, johnson17, williams17, leontiev17, kudinov17, cat17 := splitSlice(seventeenthRow)
	personalData18, johnson18, williams18, leontiev18, kudinov18, cat18 := splitSlice(eighteenthRow)
	personalData19, johnson19, williams19, leontiev19, kudinov19, cat19 := splitSlice(nineteenthRow)
	personalData20, johnson20, williams20, leontiev20, kudinov20, cat20 := splitSlice(twenteenthRow)
	personalData21, johnson21, williams21, leontiev21, kudinov21, cat21 := splitSlice(twentyfirstRow)
	personalData22, johnson22, williams22, leontiev22, kudinov22, cat22 := splitSlice(twentysecondRow)
	personalData23, johnson23, williams23, leontiev23, kudinov23, cat23 := splitSlice(twentythirdRow)
	personalData24, johnson24, williams24, leontiev24, kudinov24, cat24 := splitSlice(twentyfourthRow)
	personalData25, johnson25, williams25, leontiev25, kudinov25, cat25 := splitSlice(twentyfifthRow)
	personalData26, johnson26, williams26, leontiev26, kudinov26, cat26 := splitSlice(twentysixthRow)
	personalData27, johnson27, williams27, leontiev27, kudinov27, cat27 := splitSlice(twentyseventhRow)
	personalData28, johnson28, williams28, leontiev28, kudinov28, cat28 := splitSlice(twentyeighthRow)
	personalData29, johnson29, williams29, leontiev29, kudinov29, cat29 := splitSlice(twentyninthRow)
	personalData30, johnson30, williams30, leontiev30, kudinov30, cat30 := splitSlice(thirtiethRow)
	personalData31, johnson31, williams31, leontiev31, kudinov31, cat31 := splitSlice(thirtyfirstRow)
	personalData32, johnson32, williams32, leontiev32, kudinov32, cat32 := splitSlice(thirtysecondRow)
	personalData33, johnson33, williams33, leontiev33, kudinov33, cat33 := splitSlice(thirtythirdRow)
	personalData34, johnson34, williams34, leontiev34, kudinov34, cat34 := splitSlice(thirtyfourthRow)
	personalData35, johnson35, williams35, leontiev35, kudinov35, cat35 := splitSlice(thirtyfifthRow)
	personalData36, johnson36, williams36, leontiev36, kudinov36, cat36 := splitSlice(thirtysixthRow)
	personalData37, johnson37, williams37, leontiev37, kudinov37, cat37 := splitSlice(thirtyseventhRow)
	personalData38, johnson38, williams38, leontiev38, kudinov38, cat38 := splitSlice(thirtyeightsRow)
	personalData39, johnson39, williams39, leontiev39, kudinov39, cat39 := splitSlice(thirtyninethRow)
	personalData40, johnson40, williams40, leontiev40, kudinov40, cat40 := splitSlice(fourtiethRow)
	personalData41, johnson41, williams41, leontiev41, kudinov41, cat41 := splitSlice(fourtyfirstRow)
	personalData42, johnson42, williams42, leontiev42, kudinov42, cat42 := splitSlice(fourtysecondRow)
	personalData43, johnson43, williams43, leontiev43, kudinov43, cat43 := splitSlice(fourtythirdRow)
	personalData44, johnson44, williams44, leontiev44, kudinov44, cat44 := splitSlice(fourtyfourthRow)
	personalData45, johnson45, williams45, leontiev45, kudinov45, cat45 := splitSlice(fourtyfifthRow)

	separator := []string{"\n=============================\n"}
	fi, err := os.OpenFile("results.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer fi.Close()
	check(err)

	addDataToFile(personalData01)
	addDataToFile(contacts(firstRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson01))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams01))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev01))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov01))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat01))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData02)
	addDataToFile(contacts(secondRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson02))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams02))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev02))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov02))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat02))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData03)
	addDataToFile(contacts(thirdRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson03))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams03))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev03))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov03))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat03))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData04)
	addDataToFile(contacts(fourthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson04))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams04))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev04))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov04))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat04))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData05)
	addDataToFile(contacts(fifthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson05))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams05))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev05))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov05))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat05))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData06)
	addDataToFile(contacts(sixthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson06))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams06))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev06))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov06))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat06))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData07)
	addDataToFile(contacts(seventhRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson07))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams07))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev07))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov07))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat07))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData08)
	addDataToFile(contacts(eightsRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson08))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams08))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev08))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov08))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat08))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData09)
	addDataToFile(contacts(ninethRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson09))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams09))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev09))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov09))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat09))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData10)
	addDataToFile(contacts(tenthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson10))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams10))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev10))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov10))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat10))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData11)
	addDataToFile(contacts(eleventhRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson11))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams11))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev11))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov11))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat11))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData12)
	addDataToFile(contacts(twelvethRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson12))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams12))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev12))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov12))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat12))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData13)
	addDataToFile(contacts(thirteenthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson13))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams13))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev13))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov13))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat13))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData14)
	addDataToFile(contacts(fourteenthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson14))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams14))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev14))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov14))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat14))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData15)
	addDataToFile(contacts(fifteenthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson15))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams15))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev15))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov15))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat15))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData16)
	addDataToFile(contacts(sixteenthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson16))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams16))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev16))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov16))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat16))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData17)
	addDataToFile(contacts(seventeenthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson17))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams17))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev17))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov17))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat17))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData18)
	addDataToFile(contacts(eighteenthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson18))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams18))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev18))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov18))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat18))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData19)
	addDataToFile(contacts(nineteenthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson19))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams19))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev19))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov19))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat19))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData20)
	addDataToFile(contacts(twenteenthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson20))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams20))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev20))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov20))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat20))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData21)
	addDataToFile(contacts(twentyfirstRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson21))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams21))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev21))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov21))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat21))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData22)
	addDataToFile(contacts(twentysecondRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson22))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams22))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev22))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov22))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat22))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData23)
	addDataToFile(contacts(twentythirdRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson23))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams23))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev23))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov23))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat23))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData24)
	addDataToFile(contacts(twentyfourthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson24))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams24))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev24))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov24))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat24))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData25)
	addDataToFile(contacts(twentyfifthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson25))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams25))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev25))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov25))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat25))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData26)
	addDataToFile(contacts(twentysixthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson26))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams26))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev26))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov26))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat26))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData27)
	addDataToFile(contacts(twentyseventhRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson27))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams27))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev27))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov27))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat27))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData28)
	addDataToFile(contacts(twentyeighthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson28))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams28))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev28))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov28))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat28))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData29)
	addDataToFile(contacts(twentyninthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson29))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams29))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev29))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov29))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat29))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData30)
	addDataToFile(contacts(thirtiethRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson30))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams30))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev30))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov30))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat30))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData31)
	addDataToFile(contacts(thirtyfirstRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson31))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams31))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev31))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov31))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat31))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData32)
	addDataToFile(contacts(thirtysecondRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson32))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams32))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev32))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov32))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat32))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData33)
	addDataToFile(contacts(thirtythirdRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson33))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams33))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev33))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov33))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat33))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData34)
	addDataToFile(contacts(thirtyfourthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson34))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams34))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev34))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov34))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat34))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData35)
	addDataToFile(contacts(thirtyfifthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson35))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams35))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev35))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov35))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat35))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData36)
	addDataToFile(contacts(thirtysixthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson36))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams36))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev36))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov36))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat36))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData37)
	addDataToFile(contacts(thirtyseventhRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson37))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams37))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev37))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov37))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat37))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData38)
	addDataToFile(separator)
	addDataToFile(contacts(thirtyeightsRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson38))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams38))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev38))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov38))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat38))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData39)
	addDataToFile(contacts(thirtyninethRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson39))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams39))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev39))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov39))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat39))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData40)
	addDataToFile(contacts(fourtiethRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson40))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams40))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev40))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov40))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat40))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData41)
	addDataToFile(contacts(fourtyfirstRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson41))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams41))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev41))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov41))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat41))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData42)
	addDataToFile(contacts(fourtysecondRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson42))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams42))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev42))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov42))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat42))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData43)
	addDataToFile(contacts(fourtythirdRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson43))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams43))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev43))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov43))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat43))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData44)
	addDataToFile(contacts(fourtyfourthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson44))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams44))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev44))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov44))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat44))

	addDataToFile(separator)
	addDataToFile(separator)

	addDataToFile(personalData45)
	addDataToFile(contacts(fourtyfifthRow))
	addDataToFile(separator)
	fi.WriteString("\n" + Johnson(johnson45))
	addDataToFile(separator)
	fi.WriteString("\n" + Williams(williams45))
	addDataToFile(separator)
	fi.WriteString("\n" + Leontiev(leontiev45))
	addDataToFile(separator)
	fi.WriteString("\n" + Kudinov(kudinov45))
	addDataToFile(separator)
	fi.WriteString("\n" + Cat(cat45))

	fmt.Println("all done")
}
