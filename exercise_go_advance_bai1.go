package main

import (
	"errors"
	"fmt"
	"regexp"
	"time"
	"unicode"
)

const (
	minimumYear                     = 1900
	regexValidateEmail              = `\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	errorMessageInvalidName         = "invalid name"
	errorMessageInvalidBirthdayYear = "invalid birthday year"
	errorMessageInvalidEmail        = "invalid email"
	errorMessageInvalidPhone        = "invalid phone"
	lowerLenPhone                   = 11
	higherLenPhone                  = 12
	lowerInt                        = 100000000
	higherInt                       = 9999999999
)

type Person struct {
	name         string
	birthdayYear int
	age          int
	email        string
	phone        string
}

type PersonInterface interface {
	setName(string) error
	setBirthdayYear(int) error
	setEmail(string) error
	setPhone(interface{}) error
}

func (p *Person) setName(input string) error {
	isValid := p.validateName(input)
	if !isValid {
		return errors.New(errorMessageInvalidName)
	}
	p.name = input
	return nil
}

func (p *Person) setBirthdayYear(input int) error {
	isValid := p.validateBirthdayYear(input)
	if !isValid {
		return errors.New(errorMessageInvalidBirthdayYear)
	}
	p.birthdayYear = input
	age := p.calculateAge(input)
	p.age = age
	return nil
}

func (p *Person) setEmail(input string) error {
	isValid := p.validateEmail(input)
	if !isValid {
		return errors.New(errorMessageInvalidEmail)
	}
	p.email = input
	return nil
}

func (p *Person) setPhone(input interface{}) error {
	isValid, phone := p.validatePhone(input)
	if !isValid {
		return errors.New(errorMessageInvalidPhone)
	}
	p.phone = phone
	return nil
}

func (p *Person) validateName(input string) bool {
	first := rune(input[0])
	return unicode.IsUpper(first) && unicode.IsLetter(first)
}

func (p *Person) validateBirthdayYear(input int) bool {
	return input >= minimumYear
}

func (p *Person) calculateAge(input int) int {
	currentYear := time.Now().Year()
	return currentYear - input
}

func (p *Person) validateEmail(input string) bool {
	regexEmail := regexp.MustCompile(regexValidateEmail)
	return regexEmail.MatchString(input)
}

func (p *Person) validatePhone(input interface{}) (bool, string) {
	inputStr, isString := input.(string)
	inputInt, isInt := input.(int)
	if isString {
		first := inputStr[0]
		if first != '+' {
			return false, ""
		}
		if len(inputStr) < lowerLenPhone || len(inputStr) > higherLenPhone {
			return false, ""
		}
		return true, inputStr
	}
	if isInt {
		if inputInt < lowerInt || inputInt > higherInt {
			return false, ""
		}
		return true, fmt.Sprintf("0%d", inputInt)
	}
	return false, ""
}

func main() {
	p := &Person{}
	err := p.setName("TranBaoNgoc")
	if err != nil {
		fmt.Printf("setName fail with err = %v", err)
		return
	}
	err = p.setEmail("ngoc@ngoc.ngoc")
	if err != nil {
		fmt.Printf("setEmail fail with err = %v", err)
		return
	}
	err = p.setBirthdayYear(1900)
	if err != nil {
		fmt.Printf("setBirthdayYear fail with err = %v", err)
		return
	}
	err = p.setPhone("+44444444444")
	if err != nil {
		fmt.Printf("setEmail fail with err = %v", err)
		return
	}
	fmt.Println(p)
}
