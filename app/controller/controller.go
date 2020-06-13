package controller

import (
	"fmt"
	socketio "github.com/graarh/golang-socketio"
	"github.com/koushamad/goro/app/service"
)

func Connect (c *socketio.Channel) {
	fmt.Println("Connect")
	c.Join("Rashma")
}

func Disconnet (c *socketio.Channel) {
	fmt.Println("Disconnet")
	service.Disconnect(c.Id())
	c.Leave("Rashma")
}

func Error (c *socketio.Channel) {
	fmt.Println("Error")
	c.Leave("Rashma")
}

func Auth(c *socketio.Channel, auth service.Auth) string {
	fmt.Println("Auth")
	profileId := auth.Check(c.Id())
	c.Join(string(profileId))
	return "OK"
}

func Join(c *socketio.Channel, quiz service.Quiz ) string {
	fmt.Println("Join")
	if quiz.Join() {
		c.Join(quiz.QuizId)
	}
	return "OK"
}

func Leave(c *socketio.Channel, quiz service.Quiz ) string {
	fmt.Println("Leave")
	if quiz.Leave() {
		c.Leave(quiz.QuizId)
	}
	return "OK"
}

func Send(c *socketio.Channel, msg service.MessageSend) string {
	fmt.Println("Send")
	if service.Connect(c.Id()){
		profile := service.Profile(c.Id())
		err, message :=  msg.Send(profile.ID)
		if err == nil {
			fmt.Println("Receive")
			c.Emit("receive", message)
			c.BroadcastTo(msg.QuizId, "receive", message)
		}
	}
	return "OK"
}