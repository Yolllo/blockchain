package core

import (
	"log"
	"time"
)

const INACTIVITY_LIMIT = 3

func (c *Core) HeathcheckScheduler() {
	for {
		err := c.UpdateNodeStatus()
		if err != nil {
			// down more than 1 min (not shuffle)
			if c.NodeStatus.InactivityCounter == 1 {
				c.SendBotMessage("node is down")
				log.Println("node is down")
			}
			// down more than 3 min (needs restart)
			if c.NodeStatus.InactivityCounter == 3 {
				c.SendBotMessage("node is running...")
				log.Println("node is running...")
				c.ExecNode()
				time.Sleep(60 * time.Second)
			}
			c.NodeStatus.InactivityCounter++
		}

		if err == nil {
			// up after shuffling
			if c.NodeStatus.InactivityCounter > 0 && c.NodeStatus.InactivityCounter < INACTIVITY_LIMIT {
				log.Println("node is shuffled")
			}
			// up after restarting
			if c.NodeStatus.InactivityCounter >= INACTIVITY_LIMIT {
				c.SendBotMessage("node is up")
				log.Println("node is up")
			}
			c.NodeStatus.InactivityCounter = 0
		}

		time.Sleep(60 * time.Second)
	}
}
