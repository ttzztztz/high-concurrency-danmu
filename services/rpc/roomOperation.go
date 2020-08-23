//support crud for room, in GetRoom and UpdateRoom, there is a redundancy attribute: room_id
//when you need these two services, only supply creator_id
package rpc

import (
	"danmu/models"
	"danmu/protobuf"
	"danmu/services/db"
	"fmt"
	"time"
)

type RoomService struct{}

func (s *RoomService) GetRoom(req protobuf.RoomRequest, res *protobuf.RoomResponse) error {
	start := time.Now()

	roomId, creatorId := req.RoomId, req.CreatorId
	fmt.Println(roomId, creatorId)
	room := new(models.Room)
	if _, err := db.DB.Where("id = ? and creator_id = ?", roomId, creatorId).Get(room); err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("查询通过")

	end := time.Now()
	timeUsed := end.Sub(start).Milliseconds()

	res.Status = true
	res.RoomId = roomId
	res.CreatorId = creatorId
	res.Flow = room.Flow
	res.TimeUsed = timeUsed
	return nil
}

func (s *RoomService) AddRoom(req protobuf.RoomInfoRequest, res *protobuf.ChangeResponse) error {
	var (
		num int64
		err error
	)

	cId := req.CreatorId
	flow := req.Flow

	room := models.Room{
		Id:        0,
		CreatorId: cId,
		Flow:      flow,
	}

	if num, err = db.DB.Insert(&room); err != nil {
		fmt.Println(err.Error())
		return err
	}

	res.Status = true
	res.Num = num
	return nil
}

func (s *RoomService) UpdateFlow(req protobuf.RoomInfoRequest, res *protobuf.ChangeResponse) error {
	var (
		num int64
		err error
	)

	rId := req.RoomId
	cId := req.CreatorId
	flow := req.Flow

	room := models.Room{
		Id:        rId,
		CreatorId: cId,
		Flow:      flow,
	}
	if num, err = db.DB.Where("creator_id = ?", cId).Update(&room); err != nil {
		fmt.Println(err.Error())
		return err
	}

	res.Status = true
	res.Num = num
	return nil
}

func (s *RoomService) RemoveRoom(req protobuf.RoomRequest, res *protobuf.ChangeResponse) error {
	var (
		num int64
		err error
	)

	creatorId := req.CreatorId
	var user models.User
	if num, err = db.DB.Where("creator_id = ?", creatorId).Delete(&user); err != nil {
		fmt.Println(err.Error())
		return err
	}

	res.Status = true
	res.Num = num
	return nil
}
