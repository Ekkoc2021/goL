package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"math/rand"
)

// json和xml数据的格式化
// 有关json的数据格式化的库: "encoding/json"

// go基础接近尾声
// 在我看来json数据的格式化使用频率是比较大的,尽管api使用很简单
// 结构体<==>json数据
// 1,结构体格式化的属性必须全部可以被外部文件访问,可以使用tag标记属性
// 2,结构体==>json:Marsha()函数 返回[]byte
// 3,结构体<==json:MarshalIndent()函数
func jsondemo() {
	type student struct {
		Name string `json:"名字"`
		Id   int    `json:"编号"`
	}
	st1 := student{"ekko", 10001}
	marshal1, _ := json.Marshal(st1)
	fmt.Println(string(marshal1))

	st2 := student{}
	json.Unmarshal(marshal1, &st2)
	fmt.Println(st2)

	//json==>interface==>map使用接口解析,自动转为map
	var i interface{}
	json.Unmarshal(marshal1, &i)
	fmt.Println(i)
	fmt.Printf("%T", i)
	//使用需要通过类型断言取出才能用,v的类型也是interface,会根据具体字符串情况进行转换:数值都是float64
	m := i.(map[string]interface{})
	for k, v := range m {
		fmt.Println(k, v)
		fmt.Printf("%T", v)
	}

	//json==>map
	var m2 map[string]interface{}
	json.Unmarshal(marshal1, &m)
	//map==>json就不演示了
	for k, v := range m2 {
		fmt.Println(k, v)
		fmt.Printf("%T", v)
	}
	fmt.Println()
}

// xml
func xmldmeo() {
	// 抽取单个server对象
	type Server struct {
		ServerName string `xml:"serverName"`
		ServerIP   string `xml:"serverIP"`
	}

	type Servers struct {
		Name    xml.Name `xml:"servers"`
		Version int      `xml:"version"`
		Servers []Server `xml:"server"`
	}
	data := []byte("  <?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n  " +
		"  <servers version=\"1\">\n       " +
		" <server>\n           " +
		" <serverName>Shanghai_VPN</serverName>\n           " +
		" <serverIP>127.0.0.1</serverIP>\n      " +
		"  </server>\n       " +
		" <server>\n            " +
		"<serverName>Beijing_VPN</serverName>\n            " +
		"<serverIP>127.0.0.2</serverIP>\n       " +
		" </server>\n   " +
		" </servers>")

	var servers Servers
	err := xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("xml: %#v\n", servers)
	marshal, _ := xml.Marshal(servers)
	fmt.Println(string(marshal))
}

// 二进制读取和写入:MSGPack库
// go get -u github.com/vmihailenco/msgpack
type Person struct {
	Name string
	Age  int
	Sex  string
}

// 二进制写出
func writerJson(filename string) (err error) {
	var persons []*Person
	// 假数据
	for i := 0; i < 10; i++ {
		p := &Person{
			Name: fmt.Sprintf("name%d", i),
			Age:  rand.Intn(100),
			Sex:  "male",
		}
		persons = append(persons, p)
	}
	// 二进制json序列化
	data, err := msgpack.Marshal(persons)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// 二进制读取
func readJson(filename string) (err error) {
	var persons []*Person
	// 读文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 反序列化
	err = msgpack.Unmarshal(data, &persons)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range persons {
		fmt.Printf("%#v\n", v)
	}
	return
}

func MSGPackDemo() {
	err := writerJson("./ch11/ch11_8_json+xml/person.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	err2 := readJson("./ch11/ch11_8_json+xml/person.dat")
	if err2 != nil {
		fmt.Println(err)
		return
	}
}
func main() {
	jsondemo()
	xmldmeo()
	MSGPackDemo()
}
