package core

import (
	"fmt"
	"log"
	"time"
)

func (c *Core) ServiceScheduler() {
	for {

		secondMonitorIsAlive := c.CheckAlive("http://" + c.Config.Monitor.SecondAddr + "/service/alive")
		if !secondMonitorIsAlive {
			c.SendBotMessage("Second monitor (" + c.Config.Monitor.SecondAddr + ") not responding")
		}
		time.Sleep(1 * time.Second)

		if c.Config.Monitor.RedundancyLevel == 0 ||
			(c.Config.Monitor.RedundancyLevel == 1 && !secondMonitorIsAlive) {
			log.Println("checking instanceses")
			for _, v := range c.InstanceList {
				isAlive := c.CheckAlive(v.Endpoint)
				fmt.Println(isAlive)
				if !isAlive {
					log.Println(v.Name + " not responding")
					c.SendBotMessage(v.Name + " (" + v.Addr + ") server not responding")
				}
				time.Sleep(5 * time.Second)
			}
		}

		time.Sleep(60 * time.Second)
	}
}
