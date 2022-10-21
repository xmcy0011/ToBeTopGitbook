package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// DAO层查数据库
func queryByName(userName string) (int32, error) {
	if userName != "test" {
		return 0, errors.New("userName not found")
	}
	return 9527, nil
}

// Service层
func login(userName, pwd string) error {
	if _, err := queryByName(userName); err != nil {
		log.Println("queryByName error")
		return err
		// return errors.Wrap(err, "queryByName error")
	}

	// ... 余下登录 逻辑
	return nil
}

func loginWithErrorWrap(userName, pwd string) error {
	if _, err := queryByName(userName); err != nil {
		return errors.Wrap(err, "queryUserById error")
	}
	return nil
}

// 模拟 Biz 层调用 Service代码
func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	fmt.Println("example1: ")
	if err := login("admin", "123"); err != nil {
		log.Println(err)
	}

	fmt.Println("\nexample2: ")
	if err := loginWithErrorWrap("admin", "123"); err != nil {
		log.Printf("%+v", err)
	}
}
