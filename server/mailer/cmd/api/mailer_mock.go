package main

type mockMailer struct {
}

func (m mockMailer) Send(recipient, sublejct, text string) error {
	return nil
}
