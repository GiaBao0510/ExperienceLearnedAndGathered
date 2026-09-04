package main

import "fmt"

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (s *SecurityCode) checkCode(code int) error {
	if s.code != code {
		return fmt.Errorf("security code does not match")
	}

	fmt.Println("Security code matches")
	return nil
}