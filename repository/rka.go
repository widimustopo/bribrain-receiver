package repository

import (
	"bribrain-receiver/model"
	"bribrain-receiver/rpc/rka"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Queries struct {
}

func (q *Queries) Adduser(in *rka.User) (*rka.ResponseUser, error) {
	userModel := &model.Users{
		Name: in.Name,
	}

	//var newIn model.Users
	tx := ConnectionDB.Table("users").Save(&userModel).Last(&userModel)

	if tx.Error != nil {
		fmt.Println("error Adduser:", tx.Error)
		return nil, tx.Error
	}

	created := timestamppb.New(*userModel.CreatedAt)
	return &rka.ResponseUser{
		Code:    http.StatusOK,
		Message: "success insert data user",
		Data: &rka.User{
			Id:        userModel.Id,
			Name:      userModel.Name,
			CreatedAt: created,
		},
	}, nil
}

func (q *Queries) AddRka(in *rka.RKA) (*rka.ResponseRKA, error) {
	rkaModel := &model.Rka{
		NameRka:   in.NameRka,
		TargetRka: in.TargetRka,
	}

	//var newIn *model.Rka
	tx := ConnectionDB.Table("rka").Save(&rkaModel).Last(&rkaModel)
	if tx.Error != nil {
		fmt.Println("error AddRka:", tx.Error)
		return nil, tx.Error
	}

	var newRkaModel []*model.Rka

	result := ConnectionDB.Table("rka").Find(&newRkaModel)
	if result.Error != nil {
		fmt.Println("error Find Rka:", result.Error)
		return nil, result.Error
	}

	var newRka []*rka.RKA
	for _, data := range newRkaModel {
		created := timestamppb.New(*data.CreatedAt)
		newRka = append(newRka, &rka.RKA{
			Id:        data.RkaId,
			NameRka:   data.NameRka,
			TargetRka: data.TargetRka,
			CreatedAt: created,
		})
	}

	return &rka.ResponseRKA{
		Code:    http.StatusOK,
		Message: "success insert data rka",
		Data:    newRka,
	}, nil
}

func (q *Queries) AddRkaUser(in *rka.RequestUserRKA) (*rka.ResponseUserRKA, error) {
	var rkaUserModel []*model.RKAUsers

	for _, inData := range in.UserRka {
		rkaUserModel = append(rkaUserModel, &model.RKAUsers{
			RkaId:   inData.RkaId,
			UserId:  inData.UserId,
			Visited: inData.Visited,
		})
	}

	tx := ConnectionDB.Table("rka_users").Save(&rkaUserModel).Last(&rkaUserModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := ConnectionDB.Table("rka_users").Find(&rkaUserModel).Where("user_id = ?", in.UserRka[0].UserId)
	if result.Error != nil {
		fmt.Println("error Find Rka:", result.Error)
		return nil, result.Error
	}

	var userRkaRpc []*rka.UserRKA
	for _, userRkaModel := range rkaUserModel {
		created := timestamppb.New(*userRkaModel.CreatedAt)
		userRkaRpc = append(userRkaRpc, &rka.UserRKA{
			Id:        userRkaModel.RkaUserId,
			RkaId:     userRkaModel.RkaId,
			UserId:    userRkaModel.UserId,
			Visited:   userRkaModel.Visited,
			CreatedAt: created,
		})
	}

	return &rka.ResponseUserRKA{
		Code:    http.StatusOK,
		Message: "success insert data user's rka",
		Data:    userRkaRpc,
	}, nil

}

func (q *Queries) AccumulationRKA(in *rka.RequestAccumulation) (*rka.ResponseAccumulationRKA, error) {
	var rkaUserModel []model.AccumulationRKA

	if in.Month > 0 && in.Yaer > 0 {
		result := ConnectionDB.Table("rka_users").
			Select("users.name as `name`, rka.name_rka as `rka_name`, sum(visited) as `visited`, rka.target_rka as `target`,MONTH(rka_users.created_at) AS `month`,YEAR(rka_users.created_at) AS `year`").
			Joins("JOIN users ON users.id = rka_users.user_id").
			Joins("JOIN rka ON rka.id = rka_users.rka_id").
			Where("rka_users.user_id = ? AND YEAR(rka_users.created_at) = ? AND MONTH(rka_users.created_at) = ?", in.UserId, in.Yaer, in.Month).
			Group("rka_id").
			Find(&rkaUserModel)
		if result.Error != nil {
			fmt.Println("error AccumulationRKA:", result.Error)
			return nil, result.Error
		}
	} else if in.Month > 0 && in.Yaer <= 0 {
		result := ConnectionDB.Table("rka_users").
			Select("users.name as `name`, rka.name_rka as `rka_name`, sum(visited) as `visited`, rka.target_rka as `target`,MONTH(rka_users.created_at) AS `month`,YEAR(rka_users.created_at) AS `year`").
			Joins("JOIN users ON users.id = rka_users.user_id").
			Joins("JOIN rka ON rka.id = rka_users.rka_id").
			Where("rka_users.user_id = ? AND MONTH(rka_users.created_at) = ?", in.UserId, in.Month).
			Group("rka_id").
			Find(&rkaUserModel)
		if result.Error != nil {
			fmt.Println("error AccumulationRKA:", result.Error)
			return nil, result.Error
		}
	} else if in.Month <= 0 && in.Yaer > 0 {
		result := ConnectionDB.Table("rka_users").
			Select("users.name as `name`, rka.name_rka as `rka_name`, sum(visited) as `visited`, rka.target_rka as `target`,MONTH(rka_users.created_at) AS `month`,YEAR(rka_users.created_at) AS `year`").
			Joins("JOIN users ON users.id = rka_users.user_id").
			Joins("JOIN rka ON rka.id = rka_users.rka_id").
			Where("rka_users.user_id = ? AND YEAR(rka_users.created_at) = ?", in.UserId, in.Yaer).
			Group("rka_id").
			Find(&rkaUserModel)
		if result.Error != nil {
			fmt.Println("error AccumulationRKA:", result.Error)
			return nil, result.Error
		}
	} else {
		result := ConnectionDB.Table("rka_users").
			Select("users.name as `name`, rka.name_rka as `rka_name`, sum(visited) as `visited`, rka.target_rka as `target`,MONTH(rka_users.created_at) AS `month`,YEAR(rka_users.created_at) AS `year`").
			Joins("JOIN users ON users.id = rka_users.user_id").
			Joins("JOIN rka ON rka.id = rka_users.rka_id").
			Where("rka_users.user_id = ?", in.UserId).
			Group("rka_id,MONTH(rka_users.created_at),YEAR(rka_users.created_at)").
			Find(&rkaUserModel)
		if result.Error != nil {
			fmt.Println("error AccumulationRKA:", result.Error)
			return nil, result.Error
		}
	}

	resbyte, err := json.Marshal(&rkaUserModel)
	if err != nil {
		fmt.Println("error marhsal : ", err)
	}

	return &rka.ResponseAccumulationRKA{
		Code:            http.StatusOK,
		Message:         "success accumulation rka data",
		ProgressVisited: resbyte,
	}, nil
}
