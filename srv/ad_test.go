/**
 * @author 苗根 miaogen156@outlook.com
 * @date 2019/1/13 16:59
 */
package srv

import (
	"fmt"
	"noxue/dao"
	"noxue/utils"
	"testing"
)

func TestAdService_AdAdd(t *testing.T) {
	err := dao.AdDao.AdInsert("hello", "world", "!!!", true)
	if err != nil {
		t.Error(err)
	}
}

func TestAdService_AdExists(t *testing.T) {
	//TestAdService_AdAdd(t)
	isExists, err := SrvAd.AdExistsOfName("hello")
	utils.CheckErr(err)
	if isExists != true {
		panic("判断存在出错")
	}
}

func TestAdService_AdFindByName(t *testing.T) {
	name := "hello"
	ad, err := SrvAd.AdFindByName(name)
	utils.CheckErr(err)
	fmt.Println(ad.Id.Hex())
}

//测试多个存在相同name的存在--请在mongodb中插入两个相同名称的记录
func TestAdService_AdRemoveByName(t *testing.T) {
	name := "kkkk"
	SrvAd.AdRemoveByName(name)
}
