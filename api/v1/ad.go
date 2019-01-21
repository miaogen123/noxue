/**
 * @author 苗根 miaogen156@outlook.com
 * @date 2019/1/12 17:59
 */
package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/noxue/ormgo.v1"
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
	utils.CheckApiError(500, -1, err)

	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "添加成功",
	})

}

//获取一个广告
func (AdApi) AdGetOne(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			CheckError(c, e)
		}
	}()
	adId := c.Param("id")
	if adId == "" {
		utils.CheckApiError(422, -1, errors.New("请求未包含ID"))
	}
	ad, err := srv.SrvAd.AdFindById(adId)
	utils.CheckApiError(500, -1, err)
	c.JSON(200, gin.H{
		"ad": ad,
	})

}

//获取广告List
func (AdApi) AdGetList(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			CheckError(c, e)
		}
	}()
	fieldsValue := c.Query("fieldsvalue")
	if fieldsValue != "" {
		ApiAd.AdGetByFields(c)
		return
	}
	sort, field, filter, _, _, _, err := utils.ParseSelectParam(c)
	utils.CheckApiError(400, -1, err)
	ads, err := srv.SrvAd.AdSelect(filter, field, sort)
	utils.CheckApiError(500, -1, err)

	c.JSON(200, gin.H{
		"ads": ads,
	})
}

//获取广告-限定字段内的多个值
func (AdApi) AdGetByFields(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			CheckError(c, e)
		}
	}()

	filter, err := utils.ParseFilterParams(c)
	utils.CheckApiError(422, -1, err)

	fmt.Print(filter)
	if len(filter) != 1 {
		err := errors.New("请求中未包含或者未只包含一个字段")
		panic(err)
	}

	var ads []model.Ad
	existFlag := make(map[string]bool)
	for field, key := range filter {
		if field == "id" {
			for _, val := range key {
				ad, err := srv.SrvAd.AdFindById(val)
				utils.CheckApiError(500, -1, err)
				ads = append(ads, ad)
			}
		} else {
			for _, val := range key {
				adsFromQuery, err := srv.SrvAd.AdSelect(ormgo.M{field: val}, nil, nil)
				utils.CheckApiError(500, -1, err)
				for _, ad := range adsFromQuery {
					if _, ok := existFlag[ad.Name]; !ok {
						ads = append(ads, ad)
					}
				}
			}
		}
	}

	c.JSON(200, gin.H{
		"ads": ads,
	})
}

//更新广告
func (AdApi) AdUpdate(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			CheckError(c, e)
		}
	}()
	adId := c.Param("id")
	if adId == "" {
		err := errors.New("请求更新的广告ID为空")
		utils.CheckApiError(422, -1, err)
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
	utils.CheckApiError(500, -1, err)

	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "更新成功",
	})

}

//删除广告 -- 不再使用
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

	err := srv.SrvAd.AdRemoveById(adId)
	utils.CheckApiError(410, -1, err)
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "删除成功",
	})
}

//删除多个广告
func (AdApi) AdRemoveByIds(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			CheckError(c, e)
		}
	}()

	filter, err := utils.ParseFilterParams(c)
	utils.CheckApiError(422, -1, err)
	fmt.Print(filter)
	if len(filter) != 1 {
		err := errors.New("请求中未包含或者未只包含一个字段")
		panic(err)
	}

	for _, key := range filter {
		for _, val := range key {
			err := srv.SrvAd.AdRemoveById(val)
			utils.CheckApiError(500, -1, err)
		}
	}

	c.JSON(200, gin.H{
		"msg": "删除成功",
	})

}
