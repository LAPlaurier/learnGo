package mypackage

// go强制一个目录下所有源文件应属于一个包，简化构建过程
// 所以同一目录下，不能有两个包名

var MyVar = 100     // 公开
var anotherVar = 40 // 私有

func MyFunc() int { //公开
	return anotherVar
}

func anotherFunc() int { // 私有
	return MyVar
}
