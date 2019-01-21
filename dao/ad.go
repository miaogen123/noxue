/**
 * @author 苗根 miaogen156@outlook.com
 * @date 2019/1/12 10:18
 */
package dao

import (
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"gopkg.in/noxue/ormgo.v1"
	"noxue/model"
	"noxue/utils"
	"time"
)

var AdDao *AdDaoType

func init() {
	AdDao = &AdDaoType{}
}

type AdDaoType struct {
}

//添加广告
func (AdDaoType) AdInsert(name, title, content string, visible bool) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()
	//判断是否已经存在同名广告
	n, err := AdDao.AdCount(ormgo.M{"Name": name})
	utils.CheckErr(err)
	if n > 0 {
		utils.CheckErr(errors.New("该广告名称已存在"))
		return
	}

	ad := model.NewAd()
	ad.Name = name
	ad.Title = title
	ad.Content = content
	ad.Visible = visible
	ad.Time.CreatedAt = time.Now().UTC()
	return ad.Save()
}

//广告查找
func (AdDaoType) AdSelect(condition ormgo.M, fields map[string]bool, sorts []string, page, size int) (Ads []model.Ad, err error) {
	query := ormgo.Query{
		Condition:  condition,
		SortFields: sorts,
		Selector:   fields,
		Skip:       (page - 1) * size,
	}
	err = ormgo.FindAll(query, &Ads)
	return
}

// 获取广告ById
func (AdDaoType) AdFindById(id string) (ad model.Ad, err error) {
	err = ormgo.FindById(id, nil, &ad)
	return
}

//获取广告ByName
func (AdDaoType) AdFindByName(name string) (ad model.Ad, err error) {
	err = ormgo.FindOne(ormgo.M{"name": name}, nil, &ad)
	return
}

//更新广告的内容ById
func (AdDaoType) AdEditById(id string, ad *model.Ad) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()

	if ad == nil {
		glog.Error("Ad不能为空")
		err = errors.New("Ad不能为空")
	}
	//判断是否有相同名字的ad
	n, err := AdDao.AdCount(ormgo.M{"name": ad.Name})
	utils.CheckErr(err)
	if n > 0 {
		utils.CheckErr(errors.New("该广告名称已存在"))
		return
	}

	ad.SetDoc(ad)
	err = ad.UpdateId(id, ormgo.M{
		"name":      ad.Name,
		"title":     ad.Title,
		"content":   ad.Content,
		"visible":   ad.Visible,
		"updatedat": time.Now().UTC(),
	})
	return
}

//统计广告个数
func (AdDaoType) AdCount(conditions ormgo.M) (n int, err error) {
	ad := &model.Ad{}
	ad.SetDoc(ad)
	n, err = ad.Count(ormgo.Query{
		Condition: conditions,
	})
	return
}

//根据ID删除广告
func (AdDaoType) AdRemoveById(id string, really bool) (err error) {
	ad := &model.Ad{}
	ad.SetDoc(ad)
	if really {
		err = ad.RemoveTrueById(id)
	} else {
		err = ad.RemoveById(id)
	}
	return
}
