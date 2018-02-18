package ovc

import (
	"github.com/openvcloud/ovc-controller-manager/client"
)

// Robot base stub
type Robot struct {
	zrobot *client.ZeroRobot
}

func NewRobot(url string) *Robot {
	robot := &Robot{
		zrobot: client.NewZeroRobot(),
	}

	robot.zrobot.BaseURI = url
	return robot
}

func (r *Robot) t() {

}
