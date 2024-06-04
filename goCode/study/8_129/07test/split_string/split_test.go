package split_string

import "testing"

// TestSplit 测试组，一次测试多个样例
// 测试组，如果样例太多，我们是没办法一眼看出来具体是哪个测试用例失败了
// 采用子测试
//func TestSplit(t *testing.T) {
//	type testCase struct {
//		str  string
//		sep  string
//		want []string
//	}
//	testGroup := []testCase{
//		{str: "bbbaab", sep: "b", want: []string{"", "", "", "aa", ""}},
//		{str: "abcdef", sep: "b", want: []string{"a", "cdef"}},
//		{str: "abcef", sep: "bc", want: []string{"a", "ef"}},
//		{str: "陕科大垃圾大傻逼", sep: "大", want: []string{"陕科", "垃圾", "傻逼"}},
//	}
//	for _, tc := range testGroup {
//		got := Split(tc.str, tc.sep)
//		if !reflect.DeepEqual(got, tc.want) {
//			t.Fatalf("want %#v, got %#v\n", tc.want, got)
//		}
//	}
//}

func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	//使用map，子测试，一条样例对应一个名字
	testGroup := map[string]testCase{
		"case1": {str: "bbbaab", sep: "b", want: []string{"", "", "", "aa", ""}},
		"case2": {str: "abcdef", sep: "b", want: []string{"a", "cdef"}},
		"case3": {str: "abcef", sep: "bc", want: []string{"a", "ef"}},
		"case4": {str: "陕科大垃圾大傻逼", sep: "大", want: []string{"陕科", "垃圾", "傻逼"}},
	}

	//遍历测试	如果需要单独跑某个测试样例，在终端中输入go test -v -run=TestSplit/"case1"
	//如果想知道测试覆盖率，看看能够跑多少语句，在终端中输入go test -cover
	//将覆盖率相关记录信息输出到文件中 go test -cover -coverprofile=c.out(文件名)
	//使用工具，查看覆盖率文件 go tool cover -html=c.out(文件名)
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if len(got) != len(tc.want) {
				t.Fatalf("want %#v, got %#v\n", tc.want, got)
			}
		})
	}
}
