package data

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	authn "github.com/tx7do/kratos-authn/engine"
	authnEngine "github.com/tx7do/kratos-authn/engine"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gofrs/uuid"
	"github.com/redis/go-redis/v9"

	userV1 "kratos-monolithic-demo/gen/api/go/user/service/v1"
)

type UserTokenRepo struct {
	data          *Data
	log           *log.Helper
	authenticator authnEngine.Authenticator
}

func NewUserTokenRepo(data *Data, authenticator authnEngine.Authenticator, logger log.Logger) *UserTokenRepo {
	l := log.NewHelper(log.With(logger, "module", "user-token/repo/admin-service"))
	return &UserTokenRepo{
		data:          data,
		log:           l,
		authenticator: authenticator,
	}
}

// createAccessJwtToken 生成JWT访问令牌
func (r *UserTokenRepo) createAccessJwtToken(_ string, userId uint32) string {
	principal := authn.AuthClaims{
		Subject: strconv.FormatUint(uint64(userId), 10),
		Scopes:  make(authn.ScopeSet),
	}

	signedToken, err := r.authenticator.CreateIdentity(principal)
	if err != nil {
		return ""
	}

	return signedToken
}

// createRefreshToken 生成刷新令牌
func (r *UserTokenRepo) createRefreshToken() string {
	strUUID, _ := uuid.NewV4()
	return strUUID.String()
}

// GenerateToken 创建令牌
func (r *UserTokenRepo) GenerateToken(ctx context.Context, user *userV1.User) (accessToken string, refreshToken string, err error) {
	if accessToken = r.createAccessJwtToken(user.GetUserName(), user.GetId()); accessToken == "" {
		err = errors.New("create access token failed")
		return
	}

	if err = r.setAccessTokenToRedis(ctx, user.GetId(), accessToken, 0); err != nil {
		return
	}

	if refreshToken = r.createRefreshToken(); refreshToken == "" {
		err = errors.New("create refresh token failed")
		return
	}

	if err = r.setRefreshTokenToRedis(ctx, user.GetId(), refreshToken, 0); err != nil {
		return
	}

	return
}

// GenerateAccessToken 创建访问令牌
func (r *UserTokenRepo) GenerateAccessToken(ctx context.Context, userId uint32, userName string) (accessToken string, err error) {
	if accessToken = r.createAccessJwtToken(userName, userId); accessToken == "" {
		err = errors.New("create access token failed")
		return
	}

	if err = r.setAccessTokenToRedis(ctx, userId, accessToken, 0); err != nil {
		return
	}

	return
}

// GenerateRefreshToken 创建刷新令牌
func (r *UserTokenRepo) GenerateRefreshToken(ctx context.Context, user *userV1.User) (refreshToken string, err error) {
	if refreshToken = r.createRefreshToken(); refreshToken == "" {
		err = errors.New("create refresh token failed")
		return
	}

	if err = r.setRefreshTokenToRedis(ctx, user.GetId(), refreshToken, 0); err != nil {
		return
	}

	return
}

// RemoveToken 移除所有令牌
func (r *UserTokenRepo) RemoveToken(ctx context.Context, userId uint32) error {
	var err error
	if err = r.deleteAccessTokenFromRedis(ctx, userId); err != nil {
		r.log.Errorf("remove user access token failed: [%v]", err)
	}

	if err = r.deleteRefreshTokenFromRedis(ctx, userId); err != nil {
		r.log.Errorf("remove user refresh token failed: [%v]", err)
	}

	return err
}

// GetAccessToken 获取访问令牌
func (r *UserTokenRepo) GetAccessToken(ctx context.Context, userId uint32) string {
	return r.getAccessTokenFromRedis(ctx, userId)
}

// GetRefreshToken 获取刷新令牌
func (r *UserTokenRepo) GetRefreshToken(ctx context.Context, userId uint32) string {
	return r.getRefreshTokenFromRedis(ctx, userId)
}

const userAccessTokenKeyPrefix = "a_uat_"

func (r *UserTokenRepo) setAccessTokenToRedis(ctx context.Context, userId uint32, token string, expires int32) error {
	key := fmt.Sprintf("%s%d", userAccessTokenKeyPrefix, userId)
	return r.data.rdb.Set(ctx, key, token, time.Duration(expires)).Err()
}

func (r *UserTokenRepo) getAccessTokenFromRedis(ctx context.Context, userId uint32) string {
	key := fmt.Sprintf("%s%d", userAccessTokenKeyPrefix, userId)
	result, err := r.data.rdb.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("get redis user access token failed: %s", err.Error())
		}
		return ""
	}
	return result
}

func (r *UserTokenRepo) deleteAccessTokenFromRedis(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", userAccessTokenKeyPrefix, userId)
	return r.data.rdb.Del(ctx, key).Err()
}

const userRefreshTokenKeyPrefix = "a_urt_"

func (r *UserTokenRepo) setRefreshTokenToRedis(ctx context.Context, userId uint32, token string, expires int32) error {
	key := fmt.Sprintf("%s%d", userRefreshTokenKeyPrefix, userId)
	return r.data.rdb.Set(ctx, key, token, time.Duration(expires)).Err()
}

func (r *UserTokenRepo) getRefreshTokenFromRedis(ctx context.Context, userId uint32) string {
	key := fmt.Sprintf("%s%d", userRefreshTokenKeyPrefix, userId)
	result, err := r.data.rdb.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("get redis user refresh token failed: %s", err.Error())
		}
		return ""
	}
	return result
}

func (r *UserTokenRepo) deleteRefreshTokenFromRedis(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", userRefreshTokenKeyPrefix, userId)
	return r.data.rdb.Del(ctx, key).Err()
}
