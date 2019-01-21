/**
 * @author 苗根 miaogen156@outlook.com
 * @date 2019/1/13 16:59
 */
package srv

import (
	"fmt"
	"gopkg.in/noxue/ormgo.v1"
	"noxue/model"
	"noxue/utils"
	"testing"
)

func TestAdService_AdExistsOfName(t *testing.T) {
	isExists, err := SrvAd.AdExistsOfName("hello")
	utils.CheckErr(err)
	if isExists != true {
		panic("判断存在出错")
	}
}

func TestAdService_AdExistsOfId(t *testing.T) {
	isExists, err := SrvAd.AdExistsOfId("5c3ddabf844ee3e1ff0658ec")
	utils.CheckErr(err)
	if isExists != true {
		panic("判断存在出错")
	}
}

func TestAdService_AdAdd(t *testing.T) {
	var ad model.Ad
	ad.Name = "hello"
	ad.Title = "world"
	ad.Content = "!!!"
	ad.Visible = true
	err := SrvAd.AdAdd(ad)
	if err != nil {
		t.Error(err)
	}
}

func TestAdService_AdFindById(t *testing.T) {
	id := "5c3ddabf844ee3e1ff0658ec"
	ad, err := SrvAd.AdFindById(id)
	utils.CheckErr(err)
	fmt.Println(ad.Id.Hex())
}

func TestAdService_AdFindByName(t *testing.T) {
	name := "hello"
	ad, err := SrvAd.AdFindByName(name)
	utils.CheckErr(err)
	fmt.Println(ad.Id.Hex())
}

//FIXME::这里选择有错误
func TestAdService_AdSelect(t *testing.T) {
	ads, err := SrvAd.AdSelect(ormgo.M{"name": "testtest123456"}, map[string]bool{"title": true}, nil)
	utils.CheckErr(err)
	for _, ad := range ads {
		fmt.Print(ad.Name)
	}
}

func TestAdService_AdUpdateById(t *testing.T) {
	var ad model.Ad
	ad.Name = "world"
	ad.Title = "hello"
	ad.Content = "@@@"
	ad.Visible = false

	err := SrvAd.AdUpdateById("5c3ddabf844ee3e1ff0658ec", &ad)
	utils.CheckErr(err)

}

func TestAdService_AdRemoveById(t *testing.T) {
	err := SrvAd.AdRemoveById("5c3ddabf844ee3e1ff0658ec")
	utils.CheckErr(err)
}

//测试多个存在相同name的存在--请在mongodb中插入两个相同名称的记录
func TestAdService_AdRemoveByName(t *testing.T) {
	name := "world"
	err := SrvAd.AdRemoveByName(name)
	utils.CheckErr(err)
}
