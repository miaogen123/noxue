/**
 * @author 苗根 miaogen156@outlook.com
 * @date 2019/1/13 16:14
 */
package dao

import (
	"fmt"
	"gopkg.in/noxue/ormgo.v1"
	"noxue/model"
	"noxue/utils"
	"testing"
)

func init() {
	//ormgo.UseSoftDelete()
}

func TestAdDaoType_AdInsert(t *testing.T) {
	name := "testtest123456"
	title := "test"
	content := "test"
	n, err := AdDao.AdCount(ormgo.M{"name": name})
	if n > 0 {
		panic("记录已存在")
	}
	var visible bool = true
	err = AdDao.AdInsert(name, title, content, visible)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAdDaoType_AdSelect(t *testing.T) {
	res, err := AdDao.AdSelect(ormgo.M{
		"name": "testtest123456",
		//wrong selection
		//}, map[string]bool{"title": true}, nil, 0, 0)
	}, nil, nil, 0, 0)

	utils.CheckErr(err)
	if len(res) <= 0 {
		t.Error("DAO中的select出现错误")
		return
	}
	for _, ad := range res {
		fmt.Print(ad.Id.Hex())
	}

}

func TestAdDaoType_AdEditById(t *testing.T) {
	//TODO:修改下面的id
	id := "5c3ed86a844ee3e1ff0682ea"
	var ad model.Ad
	ad.Name = "123test"
	ad.Title = "123"
	ad.Content = "123"
	ad.Visible = false
	err := AdDao.AdEditById(id, &ad)
	utils.CheckErr(err)
	res, err := AdDao.AdSelect(ormgo.M{"name": "123test"}, nil, nil, 0, 0)
	utils.CheckErr(err)
	for _, ad := range res {
		fmt.Print(ad.Id.Hex())
	}
	return
}

func TestAdDaoType_AdFindById(t *testing.T) {
	id := ""
	ad, err := AdDao.AdFindById(id)
	utils.CheckErr(err)
	fid := ad.Id.Hex()
	if fid != id {
		panic("TestAdDaoType_AdFindById failed")
	}
}

func TestAdDaoType_AdRemoveById(t *testing.T) {
	id := "5c3ed86a844ee3e1ff0682ea"
	err := AdDao.AdRemoveById(id, false)
	utils.CheckErr(err)
	return
}
