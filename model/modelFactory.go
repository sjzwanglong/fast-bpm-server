package model

import "strings"

func NewModelInstance(modelType string) interface{} {
	switch {
	case strings.HasSuffix(modelType, "cmpy"):
		return &OrgCmpyModel{}
	case strings.HasSuffix(modelType, "dept"):
		return &OrgDeptModel{}
	case strings.HasSuffix(modelType, "user"):
		return &OrgUserModel{}
	case strings.HasSuffix(modelType, "deptuser"):
		return &OrgDeptUserModel{}
	case strings.HasSuffix(modelType, "cmpys"):
		return &[]*OrgCmpyModel{}
	case strings.HasSuffix(modelType, "depts"):
		return &[]*OrgDeptModel{}
	case strings.HasSuffix(modelType, "users"):
		return &[]*OrgUserModel{}
	case strings.HasSuffix(modelType, "deptusers"):
		return &[]*OrgDeptUserModel{}
	default:
		return nil
	}
}
