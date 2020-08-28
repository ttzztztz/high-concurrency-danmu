//support danmu query by {id, {uid}, {room_id}, {uid, room_id}}
//support danmu adding, need danmu uid, room_id, content, visible
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

func (d *DanmuService) GetDanmuById(req *protobuf.DanmuIdRequest, res *protobuf.DanmuResponse) error {
	start := time.Now()
	var danmuList []models.Danmu
	if err := query(req, &danmuList); err != nil {
		fmt.Println(err)
		return err
	}
	retList := transform(danmuList)
	end := time.Now()

	timeUsed := end.Sub(start).Milliseconds()
	res.Status = true
	res.DanmuList = retList
	res.TimeUsed = timeUsed
	return nil
}

func (d *DanmuService) GetDanmuByUId(req *protobuf.DanmuUIdRequest, res *protobuf.DanmuResponse) error {
	start := time.Now()
	var danmuList []models.Danmu
	if err := query(req, &danmuList); err != nil {
		fmt.Println(err)
		return err
	}
	retList := transform(danmuList)
	end := time.Now()

	timeUsed := end.Sub(start).Milliseconds()
	res.Status = true
	res.DanmuList = retList
	res.TimeUsed = timeUsed
	return nil
}

func (d *DanmuService) GetDanmuByRoomId(req *protobuf.DanmuRoomIdRequest, res *protobuf.DanmuResponse) error {
	start := time.Now()
	var danmuList []models.Danmu
	if err := query(req, &danmuList); err != nil {
		fmt.Println(err)
		return err
	}
	retList := transform(danmuList)
	end := time.Now()

	timeUsed := end.Sub(start).Milliseconds()
	res.Status = true
	res.DanmuList = retList
	res.TimeUsed = timeUsed
	return nil
}

func (d *DanmuService) GetDanmuByUIdAndRoomId(req *protobuf.DanmuUIdAndRoomIdRequest, res *protobuf.DanmuResponse) error {
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
	res.Status = true
	res.DanmuList = retList
	res.TimeUsed = timeUsed
	return nil
}

func (d *DanmuService) AddDanmu(req *protobuf.DanmuRequest, res *protobuf.DanmuChangeResponse) error {
	var (
		err error
	)

	uid := req.Uid
	roomId := req.RoomId
	visible := req.Visible
	content := req.Content

	danmu := &models.Danmu{
		Id:      0,
		Uid:     uid,
		Rid:     roomId,
		Visible: visible,
		Content: content,
		Created: time.Now(),
	}

	if _, err = db.DB.Table("danmu").Insert(danmu); err != nil {
		fmt.Println(err.Error())
		return err
	}

	res.Status = true
	return nil
}

func query(req interface{}, danmuList *[]models.Danmu) error {
	switch t := req.(type) {
	case protobuf.DanmuIdRequest:
		danmuId := t.Id
		if err := db.DB.Table("danmu").Where("id = ?", danmuId).Find(danmuList); err != nil {
			fmt.Println(err)
			return err
		}
	case protobuf.DanmuUIdRequest:
		uid := t.Uid
		if err := db.DB.Table("danmu").Where("uid = ?", uid).Find(danmuList); err != nil {
			fmt.Println(err)
			return err
		}
	case protobuf.DanmuRoomIdRequest:
		roomId := t.RoomId
		if err := db.DB.Table("danmu").Where("rid = ?", roomId).Find(danmuList); err != nil {
			fmt.Println(err)
			return err
		}
	case protobuf.DanmuUIdAndRoomIdRequest:
		uid := t.Uid
		roomId := t.RoomId
		if err := db.DB.Table("danmu").Where("uid = ? and rid = ?", uid, roomId).Find(danmuList); err != nil {
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
