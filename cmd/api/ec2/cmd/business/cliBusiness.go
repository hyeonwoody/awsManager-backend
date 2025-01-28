package ec2_business

import (
	dto "awsManager/api/ec2/cmd/business/dto"
	domainDto "awsManager/api/ec2/cmd/domain/dto"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

type CliBusiness struct {
}

func NewCliBusiness() *CliBusiness {
	return &CliBusiness{}
}

func (b *CliBusiness) MountEbsVolume(command *domainDto.CliCommand) error {
	config, err := b.createSshClientConfig(&command.PrivateKeyName)
	if err != nil {
		return err
	}

	client, err := b.establishSshConnection(&command.PublicIp, config)
	if err != nil {
		return fmt.Errorf("failed to dial : %s", err)
	}

	mountPath := "/mnt/" + command.DeviceName
	b.createDirectory(client, &mountPath)

	if err != nil {
		return fmt.Errorf("failed to create session")
	}
	b.mountDisk(client, "/dev/"+command.DeviceName, mountPath)

	return nil
}

func (b *CliBusiness) mountDisk(client *ssh.Client, path string, mountPath string) error {

	session, _ := b.openSshSession(client)
	defer session.Close()
	formatErr := session.Run(fmt.Sprintf("sudo file -s %s | grep -q 'data' && sudo mkfs.ext4 %s ", path, path))
	if formatErr != nil {
		return fmt.Errorf("failed to format volume: %w", formatErr)
	}

	session, _ = b.openSshSession(client)
	defer session.Close()
	mountErr := session.Run(fmt.Sprintf("sudo mount %s %s", path, mountPath))
	if mountErr != nil {
		return fmt.Errorf("failed to mount volume: %w", mountErr)
	}

	session, _ = b.openSshSession(client)
	defer session.Close()
	fstabEntry := fmt.Sprintf("%s %s ext4 defaults,nofail 0 2", path, mountPath)
	if err := session.Run(fmt.Sprintf("echo '%s' | sudo tee -a /etc/fstab", fstabEntry)); err != nil {
		return fmt.Errorf("failed to add fstab entry: %w", err)
	}

	return nil
}

func (b *CliBusiness) MakeDir(command *domainDto.CliCommand) error {
	config, err := b.createSshClientConfig(&command.PrivateKeyName)
	if err != nil {
		return err
	}

	client, err := b.establishSshConnection(&command.PublicIp, config)
	if err != nil {
		return fmt.Errorf("failed to dial : %s", err)
	}
	defer client.Close()

	makeDirErr := b.createDirectory(client, &command.DeviceName)

	return makeDirErr
}

func (b *CliBusiness) createDirectory(client *ssh.Client, path *string) error {
	session, err := b.openSshSession(client)
	if err != nil {
		return nil
	}
	defer session.Close()
	cmd := "sudo mkdir -p " + *path
	runError := session.Run(cmd)
	if runError != nil {
		return fmt.Errorf("failed to create directory: %w", runError)
	}
	return nil
}

func (b *CliBusiness) Create(command *domainDto.CreateCommand) (*dto.Ec2Instance, error) {
	panic("Not Implemented")
}

func (b *CliBusiness) AddMemory(command *domainDto.AddMemoryCommand) error {
	config, err := b.createSshClientConfig(&command.PrivateKeyName)
	if err != nil {
		return err
	}

	client, err := b.establishSshConnection(&command.PublicIp, config)
	if err != nil {
		return fmt.Errorf("failed to dial : %s", err)
	}
	defer client.Close()

	swapfileErr := b.createSwapfile(client)

	return swapfileErr
}

func (b *CliBusiness) createSwapfile(client *ssh.Client) error {

	session, err := b.openSshSession(client)
	if err != nil {
		return fmt.Errorf("failed to create session: %s", err)
	}
	defer session.Close()

	output, err := session.CombinedOutput("[ -f /swapfile ] && echo 'exists' || echo 'not exists'")
	if err != nil {
		return fmt.Errorf("failed to check /swapfile: %s", err)
	}
	if string(output) == "not exists\n" {
		// Create 2GB swap file
		createSwapCmd := `
		sudo dd if=/dev/xvdm of=/mnt/xvdm/swapfile bs=128M count=64
        sudo chmod 600 /mnt/xvdm/swapfile
        sudo mkswap /mnt/xvdm/swapfile
        sudo swapon /mnt/xvdm/swapfile
        `
		session, err = b.openSshSession(client)
		if err != nil {
			return fmt.Errorf("failed to create session: %s", err)
		}
		defer session.Close()

		err = session.Run(createSwapCmd)
		if err != nil {
			return fmt.Errorf("failed to create swap file: %s", err)
		}

		// Update /etc/fstab
		updateFstabCmd := `echo '/mnt/xvdm/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab`
		session, err = b.openSshSession(client)
		if err != nil {
			return fmt.Errorf("failed to create session: %w", err)
		}
		defer session.Close()

		err = session.Run(updateFstabCmd)
		if err != nil {

			return fmt.Errorf("failed to update /etc/fstab: %s", err)
		}

		return fmt.Errorf("swap file created and /etc/fstab updated successfully")
	}
	fmt.Printf("Swap file already Exist")
	return nil
}

func (b *CliBusiness) openSshSession(client *ssh.Client) (*ssh.Session, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to open new ssh sesion %s", err)
	}
	return session, err
}

func (b *CliBusiness) createSshClientConfig(privateKeyName *string) (*ssh.ClientConfig, error) {
	targetUser, err := user.Lookup("projectManager")
	if err != nil {
		return nil, fmt.Errorf("failed to find user")
	}
	homeDir := targetUser.HomeDir
	pemBytes, err := os.ReadFile(homeDir + "/" + *privateKeyName + ".pem")
	if err != nil {
		return nil, fmt.Errorf("failed to read private key")
	}

	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		return nil, fmt.Errorf("unabele to parse private key")
	}

	config := &ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config, nil
}

func (b *CliBusiness) establishSshConnection(publicIp *string, config *ssh.ClientConfig) (*ssh.Client, error) {
	client, err := ssh.Dial("tcp", *publicIp+":22", config)
	if err != nil {
		//handle error
	}
	return client, nil
}

func (b *CliBusiness) InstallDocker(command *domainDto.InstallCommand) error {
	config, err := b.createSshClientConfig(&command.PrivateKeyName)
	if err != nil {
		return err
	}

	client, err := b.establishSshConnection(&command.PublicIp, config)
	if err != nil {
		return fmt.Errorf("failed to dial : %s", err)
	}
	defer client.Close()

	installErr := b.installDocker(client)
	return installErr
}

func (b *CliBusiness) installDocker(client *ssh.Client) error {

	session, err := b.openSshSession(client)
	if err != nil {
		return fmt.Errorf("failed to create session: %s", err)
	}
	defer session.Close()

	commands := []string{
		"sudo apt-get update",
		"sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common",
		"curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -",
		"sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"",
		"sudo apt-get update",
		"sudo apt-get install -y docker-ce docker-ce-cli containerd.io",
		"sudo systemctl start docker",
		"sudo systemctl enable docker",
		"sudo usermod -aG docker $USER",
	}

	for _, cmd := range commands {
		session, err := client.NewSession()
		if err != nil {
			return fmt.Errorf("unable to create session: %v", err)
		}
		defer session.Close()

		out, err := session.CombinedOutput(cmd)
		if err != nil {
			return fmt.Errorf("command execution failed: %v", err)
		}
		fmt.Printf("Command executed: %s\nOutput: %s\n", cmd, out)
	}

	fmt.Println("Docker installation completed. Please log out and log back in for group changes to take effect.")
	return nil
}

func (b *CliBusiness) InstallDockerNginx(command *domainDto.InstallDockerNginxCommand) error {
	privateKeyName := command.ProjectName + strconv.Itoa(int(command.KeyNumber))
	config, err := b.createSshClientConfig(&privateKeyName)
	if err != nil {
		return err
	}

	client, err := b.establishSshConnection(&command.PublicIp, config)
	if err != nil {
		return fmt.Errorf("failed to dial : %s", err)
	}
	defer client.Close()

	installErr := b.installDockerNginx(client)
	return installErr
}

func (b *CliBusiness) installDockerNginx(client *ssh.Client) error {
	path := "/mnt/xvdf/nginx-proxy"
	b.createDirectory(client, &path)
	b.createNginxConfig(client)
	b.createNginxDockerCompose(client)
	b.runDockerContainer(client, &path)
	return nil
}

func (b *CliBusiness) createNginxDockerCompose(client *ssh.Client) error {
	session, err := b.openSshSession(client)
	if err != nil {
		return err
	}
	defer session.Close()

	composeContent := `
services:
  nginx:
    image: nginx:latest
    ports:
      - "80:80"
      - "8153:8153"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf`

	cmd := fmt.Sprintf("mkdir -p /mnt/xvdf/nginx-proxy && cd /mnt/xvdf/nginx-proxy && echo '%s' | sudo tee docker-compose.yml", composeContent)
	output, runError := session.CombinedOutput(cmd)
	if runError != nil {
		return fmt.Errorf("failed to create docker compose: %w", runError)
	}
	fmt.Printf("Command output: %s\n", output)
	fmt.Println("docker-compose.yml file created successfully in nginx-proxy directory")
	return nil
}

func (b *CliBusiness) createNginxConfig(client *ssh.Client) error {
	session, err := b.openSshSession(client)
	if err != nil {
		return nil
	}
	defer session.Close()
	configContent := fmt.Sprintf(`events {
  worker_connections 1024;
}

http {
  upstream gocd_server {
    server %s:8153;
  }

  server {
    listen 8153;
    server_name proxy-nginx;

    # go-agent에서 go-server로의 요청 처리
    location / {
      proxy_pass http://gocd_server;
      proxy_set_header Host $host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
    }
  }
}`, b.getMyPublicIP())

	cmd := fmt.Sprintf("cd /mnt/xvdf/nginx-proxy && echo '%s' | sudo tee nginx.conf", configContent)
	output, runError := session.CombinedOutput(cmd)
	if runError != nil {
		return fmt.Errorf("failed to create docker compose: %w", runError)
	}
	fmt.Printf("Command output: %s\n", output)
	fmt.Println("nginx.conf file created successfully in nginx-proxy directory")
	return nil
}

func (b *CliBusiness) runDockerContainer(client *ssh.Client, path *string) error {
	session, err := b.openSshSession(client)
	if err != nil {
		return nil
	}
	defer session.Close()
	cmd := "cd " + *path + " && docker compose up -d"
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return fmt.Errorf("command execution failed: %v", err)
	}

	fmt.Printf("Command output: %s\n", output)
	fmt.Printf("%s docker container setup completed successfully\n", *path)
	return nil
}

func (b *CliBusiness) InstallDockerGoAgent(command *domainDto.InstallGoAgentCommand) error {
	privateKeyName := command.ProjectName + strconv.Itoa(int(command.KeyNumber))
	config, err := b.createSshClientConfig(&privateKeyName)
	if err != nil {
		return err
	}

	client, err := b.establishSshConnection(&command.PublicIp, config)
	if err != nil {
		return fmt.Errorf("failed to dial : %s", err)
	}
	defer client.Close()

	installErr := b.installDockerGoAgent(client, &command.GoServerIp, &privateKeyName)
	return installErr
}

func (b *CliBusiness) installDockerGoAgent(client *ssh.Client, goServerIp, privateKeyName *string) error {
	path := "/mnt/xvdf/go-agent"
	b.createDirectory(client, &path)
	b.createGoAgentDockerCompose(client, goServerIp, privateKeyName)
	//b.createNginxConfig(client)
	b.runDockerContainer(client, &path)
	return nil
}

func (b *CliBusiness) createGoAgentDockerCompose(client *ssh.Client, goServerIp, privateKeyName *string) error {
	session, err := b.openSshSession(client)
	if err != nil {
		return err
	}
	defer session.Close()

	composeContent := fmt.Sprintf(`services:
  gocd-agent:
    image: gocd/gocd-agent-docker-dind:v24.5.0
    privileged: true
    environment:
      - GO_SERVER_URL=http://%s:8153/go
      - AGENT_AUTO_REGISTER_KEY=e00bfc7c-1f53-4fbf-b931-cc69a32c2990
      - AGENT_AUTO_REGISTER_RESOURCES=your_resources
      - AGENT_AUTO_REGISTER_ENVIRONMENTS=your_environments
      - AGENT_AUTO_REGISTER_HOSTNAME=%s`, *goServerIp, *privateKeyName)

	cmd := fmt.Sprintf("mkdir -p /mnt/xvdf/go-agent && cd /mnt/xvdf/go-agent && echo '%s' | sudo tee docker-compose.yml", composeContent)
	output, runError := session.CombinedOutput(cmd)
	if runError != nil {
		return fmt.Errorf("failed to create docker compose: %w", runError)
	}
	fmt.Printf("Command output: %s\n", output)
	fmt.Println("docker-compose.yml file created successfully in go-agent directory")
	return nil
}

func (b *CliBusiness) InstallGoAgent(command *domainDto.InstallGoAgentCommand) error {
	privateKeyName := command.ProjectName + strconv.Itoa(int(command.KeyNumber))
	config, err := b.createSshClientConfig(&privateKeyName)
	if err != nil {
		return err
	}

	client, err := b.establishSshConnection(&command.PublicIp, config)
	if err != nil {
		return fmt.Errorf("failed to dial : %s", err)
	}
	defer client.Close()

	installErr := b.installGoAgent(client, &command.GoServerIp)
	return installErr
}

func (b *CliBusiness) installGoAgent(client *ssh.Client, goServerIp *string) error {
	path := "/mnt/xvdf/go-agent"
	b.createDirectory(client, &path)
	b.createGoAgent(client)
	b.createGoAgentConfig(client, goServerIp)
	b.restartGoAgent(client)
	//b.createNginxConfig(client)
	//b.runDockerContainer(client, &path)
	return nil
}

func (b *CliBusiness) createGoAgent(client *ssh.Client) error {
	session, err := b.openSshSession(client)
	if err != nil {
		return err
	}
	defer session.Close()

	cmd := `
    if [ -f /etc/debian_version ]; then
		sudo install -m 0755 -d /etc/apt/keyrings
		curl https://download.gocd.org/GOCD-GPG-KEY.asc | sudo gpg --dearmor -o /etc/apt/keyrings/gocd.gpg
        sudo chmod a+r /etc/apt/keyrings/gocd.gpg
echo "deb [signed-by=/etc/apt/keyrings/gocd.gpg] https://download.gocd.org /" | sudo tee /etc/apt/sources.list.d/gocd.list
		sudo apt-get update
        sudo apt-get install -y go-agent
    elif [ -f /etc/redhat-release ]; then
        sudo yum install -y go-agent
    else
        echo "Unsupported Linux distribution"
        exit 1
    fi
    `
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		fmt.Println("Error installing go-agent:", err)
		fmt.Println(string(output))
		return err
	}
	fmt.Println("GoCD agent installed successfully")
	fmt.Println(string(output))
	return nil
}

func (b *CliBusiness) createGoAgentConfig(client *ssh.Client, goServerIp *string) error {
	session, err := b.openSshSession(client)
	if err != nil {
		return err
	}
	defer session.Close()
	filePath := "/var/lib/go-agent/config/autoregister.properties"
	mkdirCmd := fmt.Sprintf("sudo mkdir -p %s", filepath.Dir(filePath))
	output, err := session.CombinedOutput(mkdirCmd)
	if err != nil {
		return fmt.Errorf("command failed: %v, output: %s", err, string(output))
	}
	fmt.Printf("Command output: %s\n", output)

	session, err = b.openSshSession(client)
	if err != nil {
		return err
	}
	defer session.Close()
	fileContent := `agent.auto.register.key=e00bfc7c-1f53-4fbf-b931-cc69a32c2990
agent.auto.register.resources=your_resources
agent.auto.register.environments=your_environments
agent.auto.register.hostname=your_agent_hostname
`
	writeFileCmd := fmt.Sprintf("echo '%s' | sudo tee %s", fileContent, filePath)
	output, err = session.CombinedOutput(writeFileCmd)
	if err != nil {
		return fmt.Errorf("failed to write autoregister.properties: %v", err)
	}
	fmt.Printf("Command output: %s\n", output)
	fmt.Println("autoregister.properties file configured successfully.")

	session, err = b.openSshSession(client)
	if err != nil {
		return err
	}
	defer session.Close()
	readFilePath := "/usr/share/go-agent/wrapper-config/wrapper-properties.conf"
	readFileCmd := fmt.Sprintf("cat %s", readFilePath)
	output, err = session.CombinedOutput(readFileCmd)
	if err != nil {
		return fmt.Errorf("command failed: %v, output: %s", err, string(output))
	}
	fmt.Printf("Command output: %s\n", output)

	content := string(output)
	updatedContent := strings.Replace(content,
		"wrapper.app.parameter.101=http://localhost",
		fmt.Sprintf("wrapper.app.parameter.101=http://%s", *goServerIp),
		-1)

	session, err = b.openSshSession(client)
	if err != nil {
		return err
	}
	defer session.Close()
	updateCmd := fmt.Sprintf("echo '%s' > %s", updatedContent, filePath)
	err = session.Run(updateCmd)
	if err != nil {
		return err
	}

	fmt.Printf("Command output: %s\n", output)
	fmt.Println("GoCD agent cofigured successfully.")
	return nil
}

func (b *CliBusiness) restartGoAgent(client *ssh.Client) error {
	session, err := b.openSshSession(client)
	if err != nil {
		return err
	}
	defer session.Close()
	restartCmd := "sudo systemctl restart go-agent"
	output, err := session.CombinedOutput(restartCmd)
	if err != nil {
		return nil
	}
	fmt.Printf("Command output: %s\n", output)
	fmt.Println("Gocd restarted successfully.")
	return nil
}

func (b *CliBusiness) getMyPublicIP() string {
	resp, err := http.Get("https://ifconfig.co")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(ip))
}
