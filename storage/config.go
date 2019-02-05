package storage

import (
	"context"
	"strconv"
	"time"

	"github.com/cybozu-go/neco"
)

// Default values
const (
	DefaultCheckUpdateInterval = 10 * time.Minute
	DefaultWorkerTimeout       = 60 * time.Minute
)

// PutEnvConfig stores proxy config to storage.
func (s Storage) PutEnvConfig(ctx context.Context, env string) error {
	return s.put(ctx, KeyEnv, env)
}

// GetEnvConfig returns proxy config from storage.
func (s Storage) GetEnvConfig(ctx context.Context) (string, error) {
	env, err := s.get(ctx, KeyEnv)
	if err == ErrNotFound {
		return neco.NoneEnv, nil
	}
	return env, err
}

// PutSlackNotification stores SlackNotification to storage
func (s Storage) PutSlackNotification(ctx context.Context, url string) error {
	return s.put(ctx, KeyNotificationSlack, url)
}

// GetSlackNotification returns SlackNotification from storage
// If not found, this returns ErrNotFound.
func (s Storage) GetSlackNotification(ctx context.Context) (string, error) {
	return s.get(ctx, KeyNotificationSlack)
}

// PutProxyConfig stores proxy config to storage.
func (s Storage) PutProxyConfig(ctx context.Context, proxy string) error {
	return s.put(ctx, KeyProxy, proxy)
}

// GetProxyConfig returns proxy config from storage.
func (s Storage) GetProxyConfig(ctx context.Context) (string, error) {
	return s.get(ctx, KeyProxy)
}

// PutQuayUsername stores proxy config to storage.
func (s Storage) PutQuayUsername(ctx context.Context, username string) error {
	return s.put(ctx, KeyQuayUsername, username)
}

// GetQuayUsername returns proxy config from storage.
func (s Storage) GetQuayUsername(ctx context.Context) (string, error) {
	return s.get(ctx, KeyQuayUsername)
}

// PutQuayPassword stores proxy config to storage.
func (s Storage) PutQuayPassword(ctx context.Context, passwd string) error {
	return s.put(ctx, KeyQuayPassword, passwd)
}

// GetQuayPassword returns proxy config from storage.
func (s Storage) GetQuayPassword(ctx context.Context) (string, error) {
	return s.get(ctx, KeyQuayPassword)
}

// PutCheckUpdateInterval stores check-update-interval config to storage.
func (s Storage) PutCheckUpdateInterval(ctx context.Context, d time.Duration) error {
	data := strconv.FormatInt(int64(d), 10)
	return s.put(ctx, KeyCheckUpdateInterval, data)
}

// GetCheckUpdateInterval returns check-update-interval config from storage. It
// returns default value if the key does not exist.
func (s Storage) GetCheckUpdateInterval(ctx context.Context) (time.Duration, error) {
	data, err := s.get(ctx, KeyCheckUpdateInterval)
	if err == ErrNotFound {
		return DefaultCheckUpdateInterval, nil
	}
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return time.Duration(i), nil
}

// PutWorkerTimeout stores worker-timeout config to storage.
func (s Storage) PutWorkerTimeout(ctx context.Context, d time.Duration) error {
	data := strconv.FormatInt(int64(d), 10)
	return s.put(ctx, KeyWorkerTimeout, data)
}

// GetWorkerTimeout returns worker-timeout config from storage. It returns
// default value if the key does not exist.
func (s Storage) GetWorkerTimeout(ctx context.Context) (time.Duration, error) {
	data, err := s.get(ctx, KeyWorkerTimeout)
	if err == ErrNotFound {
		return DefaultWorkerTimeout, nil
	}
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return time.Duration(i), nil
}

// PutGitHubToken stores github-token config to storage.
func (s Storage) PutGitHubToken(ctx context.Context, token string) error {
	return s.put(ctx, KeyGitHubToken, token)
}

// GetGitHubToken returns github-token from storage
// If not found, this returns empty string.
func (s Storage) GetGitHubToken(ctx context.Context) (string, error) {
	token, err := s.get(ctx, KeyGitHubToken)
	if err == ErrNotFound {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return token, nil
}
