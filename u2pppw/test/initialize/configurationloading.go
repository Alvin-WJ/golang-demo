//package name must not be keywords, like "init"
package initialize



import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
)

func Configurationloading() {
	const filename = "../conf.properties"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func ConfigurationloadingPrint() {
	const filename = "/Users/wangjian/go/src/u2pppw/test/resources/conf.properties"
	file, err := os.Open(filename)
	if err != nil {
		//panic(err)
		fmt.Println("read success")
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func writeFile(){
	func main() {
		wireteString := "测试n"
		var filename = "./output.txt"
		var f *os.File

		/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
		if checkFileIsExist(filename) { //如果文件存在
			f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
			fmt.Println("文件存在")
		} else {
			f, err1 = os.Create(filename) //创建文件
			fmt.Println("文件不存在")
		}
		n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)

}

func deadloop() {
	for {
		fmt.Println("abc")
	}
}