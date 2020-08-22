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
	room := new(models.Room)
	if _, err := db.DB.Where("room_id = ? and creator_id = ?", roomId, creatorId).Get(room); err != nil {
		fmt.Println(err.Error())
		return err
	}

	end := time.Now()
	timeUsed := end.Sub(start).Milliseconds()

	res = &protobuf.RoomResponse{
		Status:    true,
		RoomId:    roomId,
		CreatorId: creatorId,
		Flow:      room.Flow,
		TimeUsed:  timeUsed,
	}
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

	res = &protobuf.ChangeResponse{
		Status: true,
		Num:    num,
	}
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
	if num, err = db.DB.Where("room_id = ? and creator_id = ?", rId, cId).Update(&room); err != nil {
		fmt.Println(err.Error())
		return err
	}

	res = &protobuf.ChangeResponse{
		Status: true,
		Num:    num,
	}
	return nil
}

func (s *RoomService) RemoveRoom(req protobuf.RoomRequest, res *protobuf.ChangeResponse) error {
	var (
		num int64
		err error
	)

	roomId, creatorId := req.RoomId, req.CreatorId
	var user models.User
	if num, err = db.DB.Where("room_id = ? and creator_id = ?", roomId, creatorId).Delete(&user); err != nil {
		fmt.Println(err.Error())
		return err
	}

	res = &protobuf.ChangeResponse{
		Status: true,
		Num:    num,
	}
	return nil
}
