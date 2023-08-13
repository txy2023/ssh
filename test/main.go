package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/wanshantian/ssh/cmd"
)

func main() {
	user1 := &cmd.LoginInfo{
		User:     "tian",
		Ip:       "192.168.101.108",
		Port:     22,
		Password: "tian",
	}
	client, err := cmd.NewClient(user1)
	if err != nil {
		log.Panic(err)
	}

	s, err := client.NewStream()
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		s.Close()
		client.Close()
	}()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		ret := s.Run("echo hello")
		fmt.Println(ret)
		wg.Done()
	}()
	go func() {
		ret := s.Run("sleep 11 && pwd")
		fmt.Println(ret)
		wg.Done()
	}()
	wg.Wait()
}
