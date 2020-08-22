package rpc

import (
	"danmu/models"
	"danmu/protobuf"
	"danmu/services/db"
	"errors"
	"fmt"
	"time"
)

type DanmuService struct{}

func (d *DanmuService) GetDanmuById(req protobuf.DanmuIdRequest, res *protobuf.DanmuResponse) error {
	var (
		err error
	)

	start := time.Now()
	var danmuList []models.Danmu
	if err = query(req, &danmuList); err != nil {
		fmt.Println(err)
		return err
	}
	retList := transform(danmuList)
	end := time.Now()

	timeUsed := end.Sub(start).Milliseconds()
	res = &protobuf.DanmuResponse{
		Status:    true,
		DanmuList: retList,
		TimeUsed:  timeUsed,
	}
	return nil
}

func (d *DanmuService) GetDanmuByUId(req protobuf.DanmuUIdRequest, res *protobuf.DanmuResponse) error {
	var (
		err error
	)

	start := time.Now()
	var danmuList []models.Danmu
	if err = query(req, &danmuList); err != nil {
		fmt.Println(err)
		return err
	}
	retList := transform(danmuList)
	end := time.Now()

	timeUsed := end.Sub(start).Milliseconds()
	res = &protobuf.DanmuResponse{
		Status:    true,
		DanmuList: retList,
		TimeUsed:  timeUsed,
	}
	return nil
}

func (d *DanmuService) GetDanmuByRoomId(req protobuf.DanmuRoomIdRequest, res *protobuf.DanmuResponse) error {
	var (
		err error
	)

	start := time.Now()
	var danmuList []models.Danmu
	if err = query(req, &danmuList); err != nil {
		fmt.Println(err)
		return err
	}
	retList := transform(danmuList)
	end := time.Now()

	timeUsed := end.Sub(start).Milliseconds()
	res = &protobuf.DanmuResponse{
		Status:    true,
		DanmuList: retList,
		TimeUsed:  timeUsed,
	}
	return nil
}

func (d *DanmuService) GetDanmuByUIdAndRoomId(req protobuf.DanmuRoomIdRequest, res *protobuf.DanmuResponse) error {
	var (
		err error
	)

	start := time.Now()
	var danmuList []models.Danmu
	if err = query(req, &danmuList); err != nil {
		fmt.Println(err)
		return err
	}
	retList := transform(danmuList)
	end := time.Now()

	timeUsed := end.Sub(start).Milliseconds()
	res = &protobuf.DanmuResponse{
		Status:    true,
		DanmuList: retList,
		TimeUsed:  timeUsed,
	}
	return nil
}

func (d *DanmuService) AddDanmu(req protobuf.DanmuRequest, res *protobuf.DanmuChangeResponse) error {
	var (
		err error
	)

	uid := req.Uid
	roomId := req.RoomId
	visible := req.Visible
	content := req.Content

	danmu := &models.Danmu{
		Id:         0,
		UId:        uid,
		RoomId:     roomId,
		Visible:    visible,
		Content:    content,
		CreateTime: time.Now(),
	}

	if _, err = db.DB.Insert(danmu); err != nil {
		fmt.Println(err.Error())
		return err
	}

	res = &protobuf.DanmuChangeResponse{
		Status: true,
	}
	return nil
}

func query(req interface{}, danmuList *[]models.Danmu) error {
	switch t := req.(type) {
	case protobuf.DanmuIdRequest:
		danmuId := t.Id
		if err := db.DB.ID(danmuId).Find(&danmuList); err != nil {
			fmt.Println(err)
			return err
		}
	case protobuf.DanmuUIdRequest:
		uid := t.Uid
		if err := db.DB.Where("uid = ?", uid).Find(&danmuList); err != nil {
			fmt.Println(err)
			return err
		}
	case protobuf.DanmuRoomIdRequest:
		roomId := t.RoomId
		if err := db.DB.Where("room_id = ?", roomId).Find(&danmuList); err != nil {
			fmt.Println(err)
			return err
		}
	case protobuf.DanmuUIdAndRoomIdRequest:
		uid := t.Uid
		roomId := t.RoomId
		if err := db.DB.Where("uid = ? and room_id = ?", uid, roomId).Find(&danmuList); err != nil {
			fmt.Println(err)
			return err
		}
	default:
		return errors.New("unknown request type")
	}

	return nil
}

func transform(danmuList []models.Danmu) []*protobuf.DanmuResponse_Result {
	var retList []*protobuf.DanmuResponse_Result
	for _, danmu := range danmuList {
		result := &protobuf.DanmuResponse_Result{
			Visible: danmu.Visible,
			Content: danmu.Content,
		}

		retList = append(retList, result)
	}
	return retList
}
