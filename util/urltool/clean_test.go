package urltool

import "testing"

func Test_handleDoubleSlant(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "结尾有多个斜杠",
			args: args{
				url: "https://www.baidu.com//",
			},
			want: "https://www.baidu.com/",
		},
		{
			name: "中间有多个斜杠",
			args: args{
				url: "http://www.baidu.com//a//b//c//",
			},
			want: "http://www.baidu.com/a/b/c/",
		},
		{
			name: "maito协议",
			args: args{
				url: "mailto:zhangxian887//@gmail.com//",
			},
			want: "mailto:zhangxian887/@gmail.com/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleDoubleSlant(tt.args.url); got != tt.want {
				t.Errorf("handleDoubleSlant() = %v, want %v", got, tt.want)
			}
		})
	}
}
