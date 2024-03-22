package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Тестовое задание Ката.Академия ")
	fmt.Println("Давай попробуем, возможности ты знаешь,просто потыкай =) ")
	number, x, y, opr, eror := readLine()
	if eror != nil {
		panic("дангер.")
		fmt.Println(eror)
		return
	}
	if number == "arab" {
		numb_1, e_1 := strconv.Atoi(x)
		if e_1 != nil {
			panic("Дангер!Получилась Паника?")
			fmt.Println(e_1)
			return
		}
		numb_2, e_2 := strconv.Atoi(y)
		if e_2 != nil {
			panic("Дангер!")
			fmt.Println(e_2)
			return
		}
		ans, e_3 := calculator(numb_1, numb_2, opr)
		if e_3 != nil {
			panic("Дангер!")
			fmt.Println(e_3)
			return
		} else {
			fmt.Println("Ответ: ")
			fmt.Println(ans)
		}
	} else {
		numb_1 := fromRomanToInt(x)
		numb_2 := fromRomanToInt(y)
		ans, e_1 := calculator(numb_1, numb_2, opr)
		if e_1 != nil {
			panic("Дангер!")
			fmt.Println(e_1)
			return
		} else {
			end, e_2 := fromIntToRoman(ans)
			if e_2 != nil {
				panic("Что то поломалось")
				fmt.Println(e_2)
				return
			}
			fmt.Println("Ответ: ")
			fmt.Println(end)
		}
	}
}

func calculator(numb_1 int, numb_2 int, opr string) (int, error) {
	if numb_1 > 10 || numb_2 > 10 {

		panic("ой ой ой,да что же такое.")
		return 7, errorHandler(7)
	}
	switch {
	case opr == "+":
		return numb_1 + numb_2, nil
	case opr == "-":
		return numb_1 - numb_2, nil
	case opr == "*":
		return numb_1 * numb_2, nil
	case opr == "/":
		return numb_1 / numb_2, nil
	default:
		return 4, errorHandler(4)
	}
}
func readLine() (string, string, string, string, error) {
	stdin := bufio.NewReader(os.Stdin)
	usInput, _ := stdin.ReadString('\n')
	usInput = strings.TrimSpace(usInput)
	number, numb_1, numb_2, opr, eror := checkInput(usInput)
	if eror != nil {

		panic("Скорее всего я не правильно понял интерфейс паника,и пременяю его не правильно.")
		return "", "", "", "", eror
	}
	return number, numb_1, numb_2, opr, eror
}

func checkInput(input string) (string, string, string, string, error) {
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(input, "")
	arr := strings.Split(replace, "")
	var number, numb_1, numb_2, opr string
	for index, value := range arr {
		isN := isNumber(value)
		isS := isSign(value)
		isR := isRomanNumber(value)
		if !isN && !isS && !isR {
			panic("ого паника паника?")
			return "", "", "", "", errorHandler(1)
		}
		if isS {
			if opr != "" {
				panic("паника ,все ведь правильно у меня получаетс?")
				return "", "", "", "", errorHandler(5)
			} else {
				opr = arr[index]
			}
		}
		if (isN && number != "roman") || (isR && number != "arab") {
			panic("ну если  программа не ругается,и нет красненькой ошибки,то пишим дальше. Кстаит чуть не забыл. Паника")
			if number == "" {
				if isN {
					number = "arab"
				} else {
					number = "roman"
				}
			}
			if numb_1 == "" && !(index+1 == len(arr)) && isSign(arr[index+1]) {
				panic("Паника!")
				slice := arr[0:(index + 1)]
				numb_1 = strings.Join(slice, "")
			} else if index+1 == len(arr) && numb_1 != "" {
				slice := arr[(len(numb_1) + 1):]
				numb_2 = strings.Join(slice, "")
			}
		} else if (number == "arab" && isR) || (number == "roman" && isN) {
			return "", "", "", "", errorHandler(2)
		}
	}
	if numb_2 == "" || numb_1 == "" || opr == "" {

		panic("ой ой ой,паника!")
		return "", "", "", "", errorHandler(3)
	}
	return number, numb_1, numb_2, opr, nil
}

func isNumber(c string) bool {
	if c >= "0" && c <= "9" {
		panic("Паника,она самая.")
		return true
	} else {
		return false
	}
}

func isSign(c string) bool {
	if c == "+" || c == "-" || c == "/" || c == "*" {
		panic("Паника! Вот и все,приплыли...")
		return true
	} else {
		return false
	}
}
func isRomanNumber(c string) bool {
	_, ok := dict[c]
	if ok {
		panic("Паника! Я определенно делаю что то не то!")
		return true
	} else {
		return false
	}
}

func errorHandler(code int) error {
	return errors.New(errorDict[code])
}

var errorDict = map[int]string{
	1: "Можно арабские/римские цифры и  '+', '-', '/', '*' ",
	2: "Арабы и Римляне не лучшие друзья",
	3: "Ты делаешь что то не то.",
	4: "Бах",
	5: "Пока сложно,но в будущем возможно",
	6: "И как мне это вывести ?",
	7: "Так-так,что то идет не так!",
}

var dict = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func fromRomanToInt(roman string) int {
	var res int
	arr := strings.Split(roman, "")
	for index, value := range arr {
		if index+1 != len(arr) && dict[value] < dict[arr[index+1]] {
			panic("Паника!")
			res -= dict[value]
		} else {
			res += dict[value]
		}
	}
	return res
}

func fromIntToRoman(number int) (string, error) {
	if number <= 0 {

		panic("Паника,я жее ее отследил? А интерфес рекавер нужен тут?")
		return "", errorHandler(6)
	}
	arr1 := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	arr2 := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var str string
	for number > 0 {
		for i := 0; i < 13; i++ {
			if arr1[i] <= number {
				panic("Паника! Растет число желтеньких восклицательных знаков,что бы это могло значить?")
				str += arr2[i]
				number -= arr1[i]
				break
			}
		}
	}
	return str, nil
}
