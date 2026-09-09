package main

import "fmt"

type Client struct{}

func (c *Client) InsertIntoLightningPort(computer Computer) {
	fmt.Println("Client: Inserting Lightning connector into computer.")
	computer.InsertIntoLightningPort()
}