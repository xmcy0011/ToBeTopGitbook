package main

import (
	"fmt"
	"sync"
	"time"
)

type User interface {
	Add(deviceId int)

	Devices() int
}

type user struct {
	sync.Mutex
	devices []int
	userId  int
}

func NewUser(uid int) *user {
	return &user{userId: uid}
}

func (u *user) Add(deviceId int) {
	u.Lock()
	defer u.Unlock()
	u.devices = append(u.devices, deviceId)
}

func (u *user) Devices() int {
	u.Lock()
	defer u.Unlock()

	return len(u.devices)
}

type UserManager interface {
	Add(uid, deviceId int)
	DeviceCount() (count int)
}

type userManager struct {
	users map[int]User
	sync.Mutex
}

func NewUserManager() UserManager {
	return &userManager{users: make(map[int]User, 0)}
}

func (u *userManager) Add(uid, deviceId int) {
	u.Lock()
	defer u.Unlock()

	if curUser, ok := u.users[uid]; ok {
		curUser.Add(deviceId)
	} else {
		newUser := NewUser(uid)
		newUser.Add(deviceId)

		u.users[uid] = newUser
	}
}

func (u *userManager) DeviceCount() (count int) {
	u.Lock()
	defer u.Unlock()

	for _, u := range u.users {
		count += u.Devices()
	}
	return
}

type DeviceMetric interface {
	DeviceCount() (count int)
}

func main1() {

	um := NewUserManager()

	go func() {
		for i := 0; i < 100*1000; i++ {
			// 1个user 1个device
			um.Add(i, i)
			//time.Sleep(time.Microsecond * 1)
			time.Sleep(time.Nanosecond * 1)
		}
	}()

	// go func() {
	// 	for {
	// 		fmt.Println("device ", um.DeviceCount())
	// 		time.Sleep(time.Microsecond * 100)
	// 	}
	// }()

	// 模拟metric
	type MetricProxy struct {
		UserManager
	}
	in := &MetricProxy{
		UserManager: um,
	}

	go func(metric DeviceMetric) {
		fmt.Printf("%p %p %p \n", metric, um, in.UserManager)

		for {
			c := metric.DeviceCount()
			if c > 10*1000 {
				fmt.Println("metric goroutine exit")
				break
			}
			time.Sleep(time.Microsecond * 1)
		}
	}(in)

	time.Sleep(time.Second * 100)
}

func main() {
	w := make(chan string, 2)

	w <- "1"
	fmt.Println("write 1")

	w <- "2"
	fmt.Println("write 2")

	select {
	case w <- "3":
		fmt.Println("write 3")
	default:
		fmt.Println("msg flll")
	}
	// if len(w) == cap(w) {
	// 	fmt.Println("msg flll")
	// } else {
	// 	w <- "3"
	// 	fmt.Println("write 3")
	// }
}
