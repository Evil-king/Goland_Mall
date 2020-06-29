package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"log"
	"os"
)

//CreateUUID 生成UUID
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
//分布式ID
func IdWork() int64  {
	n, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	id := n.Generate()
	return int64(id)
}
