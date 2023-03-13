// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ITCPAuth interface {
		Start(ctx context.Context)
		Stop(ctx context.Context)
	}
)

var (
	localTCPAuth ITCPAuth
)

func TCPAuth() ITCPAuth {
	if localTCPAuth == nil {
		panic("implement not found for interface ITCPAuth, forgot register?")
	}
	return localTCPAuth
}

func RegisterTCPAuth(i ITCPAuth) {
	localTCPAuth = i
}
