package tasks

import "testing"

func TestSendMail1(t *testing.T) {
	type args struct {
		body    string
		email   string
		subject string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "t1", args: struct {
			body    string
			email   string
			subject string
		}{body: "body_1", email: "valid_email@gmail.com", subject: "sub_1"}, want: true},
		{name: "t1", args: struct {
			body    string
			email   string
			subject string
		}{body: "body_1", email: "not_a_valid_email.com", subject: "sub_1"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got, _ := SendMail(tt.args.email , tt.args.subject, tt.args.body); got != tt.want {
			//	t.Errorf("SendMail() = %v, want %v", got, tt.want)
			//}
		})
	}
}
