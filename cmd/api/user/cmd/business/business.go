package user_business

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

type Business struct {
}

func NewBusiness() *Business {
	return &Business{}
}

func (b *Business) CreateCredentialConfigure(projectName, accessKey, secretAccessKey string, keyNumber uint) error {
	file, shouldReturn, err := b.getHomeDirectory()
	if shouldReturn {
		return err
	}
	profileContent := fmt.Sprintf("[%s]\naws_access_key_id = %s\naws_secret_access_key = %s\n",
		projectName+strconv.Itoa(int(keyNumber)), accessKey, secretAccessKey)

	_, err = file.WriteString(profileContent)
	if err != nil {
		return fmt.Errorf("failed to write credentials: %w", err)
	}
	defer file.Close()
	return nil
}

func (*Business) getHomeDirectory() (*os.File, bool, error) {
	targetUser, err := user.Lookup("projectManager")
	homeDir := targetUser.HomeDir
	if err != nil {
		return nil, true, fmt.Errorf("failed to get user home directory: %w", err)
	}

	// uid, _ := strconv.Atoi(targetUser.Uid)
	// gid, _ := strconv.Atoi(targetUser.Gid)

	credentialsPath := filepath.Join(homeDir, ".aws", "credentials")
	file, err := os.OpenFile(credentialsPath, os.O_WRONLY|os.O_APPEND|os.O_APPEND, 0777)
	if err != nil {
		return nil, true, fmt.Errorf("failed to open credentials file : %w", err)
	}
	return file, false, nil
}
