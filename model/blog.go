/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */

package model


type Blog struct {
	Article `bson:",inline"`
}

func (this *Blog) GetCName() string {
	return "blog"
}