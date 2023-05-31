package main

import "fmt"

/**
--测试了gin的get和post，这个在redis.go方法里
--测试了mysql,这个是在sql.go里边
over
*/

func main() {
	db := GetInstance()
	/**
	写了恢复函数之后，最后不需要再单独关闭db，因为恢复函数会在正常退出或异常退出时都执行，
	所以可以保证db对象被关闭。如果你再单独关闭db，可能会导致重复关闭db对象，
	虽然不会影响程序的运行，但是不是一个好的编程习惯。
	*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
			// 关闭db对象
			db.Close()
		}
	}()

	TestMysql()
	// TestRedis()

}
