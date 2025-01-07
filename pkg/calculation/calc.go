package calculation


import (
	"errors"
	
	"strconv"
	"strings"
)

// Определяем приоритет операций
var precedence = map[rune]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

// Функция для вычисления выражения
func Calc(expression string) (float64, error) {
	slise := modific(expression)
	if len(slise) == 0 {
		return 0, errors.New("пустое выражение")
	}

	postfix, err := Postfix(slise)
	if err != nil {
		return 0, err
	}

	return calculation(postfix)
}

// Токенизация входного выражения
func modific(expression string) []string {
	var result []string

	//буфер для работы со строкой
	var bild strings.Builder

	for _, run := range expression {
		if run == ' ' {
			continue
		}
		if isOperator(run) || run == '(' || run == ')' {
			if bild.Len() > 0 {
				result = append(result, bild.String())
				bild.Reset()
			}
			result = append(result, string(run))
		} else if isDigit(run) || run == '.' {
			bild.WriteRune(run)
		} else {
			return nil // Неверный символ
		}
	}
	if bild.Len() > 0 {
		result = append(result, bild.String())
	}

	return result
}

// Проверка оператора
func isOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

// Проверка  цифры
func isDigit(c rune) bool {
	return (c >= '0' && c <= '9') || c == '.'
}

// Преобразование инфиксного выражения в постфиксное
func Postfix(slise []string) ([]string, error) {
	var output []string
	var stack []rune

	for _, vol := range slise {

		if isDigit(rune(vol[0])) { // Проверка если число
			output = append(output, vol)
		} else if vol == "(" { // Если  открывающая скобка
			stack = append(stack, '(')
		} else if vol == ")" { // Если  закрывающая скобка

			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				output = append(output, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, errors.New("несоответствующая закрывающая скобка")
			}
			stack = stack[:len(stack)-1] // Удаляем открывающую скобку
		} else if isOperator(rune(vol[0])) { // Если оператор
			for len(stack) > 0 && precedence[rune(vol[0])] <= precedence[stack[len(stack)-1]] {
				output = append(output, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, rune(vol[0]))
		} else {
			return nil, errors.New("недопустимый токен: " + vol)
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == '(' {
			return nil, errors.New("несоответствующая открывающая скобка")
		}
		output = append(output, string(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

// Вычисление  выражения
func calculation(slise []string) (float64, error) {
	var stack []float64

	for _, vol := range slise {
		if isDigit(rune(vol[0])) { // Если  число
			num, err := strconv.ParseFloat(vol, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
	 } else if isOperator(rune(vol[0])) { // Если  оператор
			if len(stack) < 2 {
				return 0, errors.New("недостаточно операндов")
		 }
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			

			switch vol[0] {
			 case '+':
				stack = append(stack, a+b)
			 case '-':
				stack = append(stack, a-b)
			 case '*':
				stack = append(stack, a*b)
			 case '/':
				if b == 0 {
					return 0, errors.New("деление на ноль")
			 }
				stack = append(stack, a/b)
		  }
	  } else {
		  return 0, errors.New("недопустимый токен: " + vol)
	  }
   }

   if len(stack) != 1 {
	  return 0, errors.New("ошибка в вычислении")
   }

   return stack[0], nil
}
