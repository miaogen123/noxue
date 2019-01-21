/**
 * @author 苗根 miaogen156@outlook.com
 * @date 2019/1/12 16:03
 */
package srv

import (
	"errors"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/noxue/ormgo.v1"
	"noxue/dao"
	"noxue/model"
	"noxue/utils"
)

var SrvAd AdService

type AdService struct {
}

func init() {
	initAd()
}

//留空
func initAd() {

}

//判断指定名称广告是否存在
func (AdService) AdExistsOfName(name string) (isExists bool, err error) {
	n, err := dao.AdDao.AdCount(ormgo.M{"name": name})
	if err != nil {
		return
	}
	isExists = n > 0
	return
}

//判断指定ID广告是否存在
func (AdService) AdExistsOfId(id string) (isExists bool, err error) {
	n, err := dao.AdDao.AdCount(ormgo.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return
	}
	isExists = n > 0
	return
}

//添加广告
func (AdService) AdAdd(ad model.Ad) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()
	n, err := dao.AdDao.AdCount(ormgo.M{"name": ad.Name})
	utils.CheckErr(err)
	if n > 0 {
		utils.CheckErr(errors.New("该广告已存在"))
		return
	}
	err = dao.AdDao.AdInsert(ad.Name, ad.Title, ad.Content, ad.Visible)
	utils.CheckErr(err)

	return
}

//通过ID查找广告
func (AdService) AdFindById(id string) (ad model.Ad, err error) {
	ad, err = dao.AdDao.AdFindById(id)
	return
}

//通过name查找广告
func (AdService) AdFindByName(name string) (ad model.Ad, err error) {
	ad, err = dao.AdDao.AdFindByName(name)
	return
}

//选择广告
func (AdService) AdSelect(condition map[string]interface{}, fileds map[string]bool, sorts []string) (ads []model.Ad, err error) {
	ads, err = dao.AdDao.AdSelect(condition, fileds, sorts, 0, 0)
	return
}

//根据ID更新广告
func (AdService) AdUpdateById(id string, ad *model.Ad) (err error) {
	err = dao.AdDao.AdEditById(id, ad)
	return
}

//删除广告
func (AdService) AdRemoveById(id string) (err error) {
	err = dao.AdDao.AdRemoveById(id, true)
	return
}

//通过名字删除广告
func (AdService) AdRemoveByName(name string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()
	n, err := dao.AdDao.AdCount(ormgo.M{"name": name})
	utils.CheckErr(err)
	if n < 1 {
		err = errors.New("该广告不存在")
	} else if n > 1 {
		glog.Error("内部错误，数据库中存在同名的多个ad")
		err = errors.New("内部错误，数据库中存在同名的多个ad")
	}
	utils.CheckErr(err)
	ad, err := dao.AdDao.AdFindByName(name)
	utils.CheckErr(err)
	err = dao.AdDao.AdRemoveById(ad.Id.Hex(), true)
	utils.CheckErr(err)
	return
}
