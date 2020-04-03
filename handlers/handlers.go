package handlers

// import (
// 	"log"
// 	"realtimechat/structs"
// 	"time"
// )

// // PingWS -
// func PingWS(ws *Conn) {

// 	for {
// 		<-time.After(5 * time.Second)
// 		msg := structs.Message{
// 			Type: structs.MTPing,
// 		}
// 		if err := Conn.Ws.WriteJSON(msg); err != nil {
// 			log.Printf("ws send ping err: %v", err)
// 		}
// 	}
// }
