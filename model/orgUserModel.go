package model

import "fast-bpm/utils"

type OrgUserModel struct {
	BaseModel
	UserWorkId      string `gorm:"column:userWorkId" form:"userWorkId" uri:"userWorkId" json:"userWorkId"`
	UserName        string `gorm:"column:userName" form:"userName" uri:"userName" json:"userName"`
	UserLoginName   string `gorm:"column:userLoginName" form:"userLoginName" uri:"userLoginName" json:"userLoginName"`
	UserPassword    string `gorm:"column:userPassword" form:"userPassword" uri:"userPassword" json:"userPassword"`
	UserAvatars     string `gorm:"column:userAvatars" form:"userAvatars" uri:"userAvatars" json:"userAvatars"`
	UserSort        int    `gorm:"column:userSort" form:"userSort" uri:"userSort" json:"userSort"`
	UserTitle       string `gorm:"column:userTitle" form:"userTitle" uri:"userTitle" json:"userTitle"`
	UserPost        string `gorm:"column:userPost" form:"userPost" uri:"userPost" json:"userPost"`
	UserSex         int    `gorm:"column:userSex" form:"userSex" uri:"userSex" json:"userSex"`
	UserState       string `gorm:"column:userState" form:"userState" uri:"userState" json:"userState"`
	UserIdCard      string `gorm:"column:userIdCard" form:"userIdCard" uri:"userIdCard" json:"userIdCard"`
	UserMobPhone    string `gorm:"column:userMobPhone" form:"userMobPhone" uri:"userMobPhone" json:"userMobPhone"`
	UserOfficePhone string `gorm:"column:userOfficePhone" form:"userOfficePhone" uri:"userOfficePhone" json:"userOfficePhone"`
	UserEmail       string `gorm:"column:userEmail" form:"userEmail" uri:"userEmail" json:"userEmail"`
	CmpyId          string `gorm:"column:cmpyId" form:"cmpyId" uri:"cmpyId" json:"cmpyId"`
}

func (u OrgUserModel) TableName() string {
	return utils.OrgUserInfo
}

func (u OrgUserModel) Clone() interface{} {
	return new(OrgUserModel)
}

func (u OrgUserModel) CloneList() interface{} {
	return new([]*OrgUserModel)
}
