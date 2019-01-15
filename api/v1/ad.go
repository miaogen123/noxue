/**
 * @author 苗根 miaogen156@outlook.com
 * @date 2019/1/12 17:59
 */
package v1

import (
	"github.com/gin-gonic/gin"
	"noxue/dao"
	"noxue/model"
	"noxue/srv"
	"noxue/utils"
)

var ApiAd AdApi

type AdApi struct {
}

type AdInfo struct {
	Name    string
	Title   string
	Content string
	Visible bool
}

//添加一个广告
func (AdApi) AdCreate(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			CheckError(c, e)
		}
	}()
	var adInfo AdInfo
	err := c.BindJSON(&adInfo)
	utils.CheckApiError(422, -1, err)

	var ad model.Ad
	ad.Name = adInfo.Name
	ad.Content = adInfo.Content
	ad.Title = adInfo.Title
	ad.Visible = adInfo.Visible
	err = srv.SrvAd.AdAdd(ad)
	utils.CheckErr(err)

	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "添加成功",
	})

}

//获取一个广告

//获取广告List

//更新广告
func (AdApi) AdUpdate(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			CheckError(c, e)
		}
	}()
	adId := c.Param("id")
	if adId == "" {
		panic("请求未携带ID")
	}

	exists, err := srv.SrvAd.AdExistsOfId(adId)
	utils.CheckApiError(422, -1, err)
	if !exists {
		panic("更新的广告不存在")
	}

	var adInfo AdInfo
	err = c.BindJSON(&adInfo)
	utils.CheckApiError(422, -1, err)

	var ad model.Ad
	ad.Name = adInfo.Name
	ad.Content = adInfo.Content
	ad.Title = adInfo.Title
	ad.Visible = adInfo.Visible
	err = srv.SrvAd.AdUpdateById(adId, &ad)
	utils.CheckErr(err)

	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "更新成功",
	})

}

//删除广告
func (AdApi) AdRemoveById(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			CheckError(c, e)
		}
	}()
	adId := c.Param("id")
	if adId == "" {
		panic("请求未携带ID")
	}

	err := dao.AdDao.AdRemoveById(adId, true)
	utils.CheckApiError(410, -1, err)
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "删除成功",
	})
}
