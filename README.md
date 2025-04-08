# awsManager

## 🧑‍💻: Intro
❓ Problem : AWS EC2 인스턴스 수동 관리 시, 휴먼 에러 발생할 가능성이 높고, 대규모 스케일링이나 테스트 시 번거로움 😮

❗ Idea : EC2 인스턴스 생성 / 종료를 자동화하여 효율성과 안정성을 확보하는 RestFul API 도구 개발 🤔

💯 Solution : AWS SDK와 연동된 EC2 생명주기를 코드로 관리 😁

</br>

## 🧱: Structure
```
.
├── backend
│   └── cmd
│       ├── api
│       │   ├── ami
│       │   ├── ec2
│       │   │   ├── cmd
│       │   │   │   ├── application
│       │   │   │   │   └── useCase
│       │   │   │   │       ├── dto
│       │   │   │   │       │   └── in
│       │   │   │   │       │       ├── createEc2Command.go
│       │   │   │   │       │       ├── findCommand.go
│       │   │   │   │       │       ├── initEc2Command.go
│       │   │   │   │       │       ├── installDockerCommand.go
│       │   │   │   │       │       └── installGoAgentCommand.go
│       │   │   │   │       ├── ec2ProjectFacade.go
│       │   │   │   │       ├── ec2ProjectFacadeInterface.go
│       │   │   │   │       ├── ec2UserProjectFacade.go
│       │   │   │   │       ├── ec2UserProjectFacadeInterface.go
│       │   │   │   │       └── ec2UserProjectFacade_test.go
│       │   │   │   ├── business
│       │   │   │   │   ├── cliBusiness.go
│       │   │   │   │   ├── cliBusiness_test.go
│       │   │   │   │   ├── dto
│       │   │   │   │   │   └── instance.go
│       │   │   │   │   ├── model
│       │   │   │   │   └── sdkBusiness.go
│       │   │   │   ├── domain
│       │   │   │   │   ├── dto
│       │   │   │   │   │   ├── addMemoryCommand.go
│       │   │   │   │   │   ├── attachEbsVolumeCommand.go
│       │   │   │   │   │   ├── configureHostCommand.go
│       │   │   │   │   │   ├── createCommand.go
│       │   │   │   │   │   ├── deleteCommand.go
│       │   │   │   │   │   ├── installBouncerCommand.go
│       │   │   │   │   │   ├── installCommand.go
│       │   │   │   │   │   ├── installDockerGoAgentCommand.go
│       │   │   │   │   │   ├── installDockerNginxCommand.go
│       │   │   │   │   │   ├── installGocdCommand.go
│       │   │   │   │   │   ├── installGoServerCommand.go
│       │   │   │   │   │   └── makeDirCommand.go
│       │   │   │   │   ├── service.go
│       │   │   │   │   ├── serviceInterface.go
│       │   │   │   │   └── service_test.go
│       │   │   │   ├── infrastructure
│       │   │   │   │   ├── repository.go
│       │   │   │   │   ├── repositoryInterface.go
│       │   │   │   │   └── repository_test.go
│       │   │   │   └── presentation
│       │   │   │       ├── handler.go
│       │   │   │       └── handlerInterface.go
│       │   │   └── main.go
│       │   ├── model
│       │   │   ├── boGocd.go
│       │   │   ├── ec2.go
│       │   │   ├── hostEntry.go
│       │   │   ├── project.go
│       │   │   ├── subProject.go
│       │   │   └── user.go
│       │   ├── project
│       │   │   ├── cmd
│       │   │   │   ├── domain
│       │   │   │   │   ├── secretValue.go
│       │   │   │   │   ├── service.go
│       │   │   │   │   └── serviceInterface.go
│       │   │   │   ├── infrastructure
│       │   │   │   │   ├── repository.go
│       │   │   │   │   └── repositoryInterface.go
│       │   │   │   ├── presentation
│       │   │   │   │   ├── handler.go
│       │   │   │   │   └── handlerInterface.go
│       │   │   │   └── subProject
│       │   │   │       ├── domain
│       │   │   │       │   ├── service.go
│       │   │   │       │   ├── serviceInterface.go
│       │   │   │       │   └── service_test.go
│       │   │   │       └── infrastructure
│       │   │   │           ├── repository.go
│       │   │   │           └── repositoryInterface.go
│       │   │   └── main.go
│       │   ├── untitled
│       │   │   └── cmd
│       │   │       ├── handler.go
│       │   │       ├── handlerInterface.go
│       │   │       ├── repository.go
│       │   │       ├── repositoryInterface.go
│       │   │       ├── service.go
│       │   │       └── serviceInterface.go
│       │   └── user
│       │       ├── cmd
│       │       │   ├── application
│       │       │   │   └── useCase
│       │       │   │       ├── dto
│       │       │   │       │   └── in
│       │       │   │       │       └── CreateUserCommand.go
│       │       │   │       ├── projectFacade.go
│       │       │   │       └── projectFacadeInterface.go
│       │       │   ├── business
│       │       │   │   ├── business.go
│       │       │   │   └── businessInterface.go
│       │       │   ├── domain
│       │       │   │   ├── dto
│       │       │   │   │   └── in
│       │       │   │   ├── service.go
│       │       │   │   ├── serviceInterface.go
│       │       │   │   └── service_test.go
│       │       │   ├── infrastructure
│       │       │   │   ├── repository.go
│       │       │   │   └── repositoryInterface.go
│       │       │   ├── model
│       │       │   └── presentation
│       │       │       ├── handler.go
│       │       │       └── handlerInterface.go
│       │       └── main.go
│       ├── awsManager
│       ├── database
│       │   └── database.go
│       ├── dependencyInjection
│       │   └── di.go
│       ├── go.mod
│       ├── go.sum
│       ├── main.go
│       └── nohup.out
├── cc.xml
├── database
│   └── Dockerfile
├── docker-compose.yml
```
</br>

## 🛢️: Entity Relationship Diagram
<p align="center">
  <img src="https://github.com/user-attachments/assets/1190c008-ac25-44b9-86ad-9c794508b38a" alt="Centered Image">
</p>

<br>

## ✅: Implementation
### 운영체제 환경 구축 자동화
### 프로젝트 CI / CD 자동화를 위한 GoCD 연계

</br>

## 📞: Contact
- 이메일: hyeonwoody@gmail.com
- 블로그: https://velog.io/@hyeonwoody
- 깃헙: https://github.com/hyeonwoody

</br>

## 🛠️: Technologies Used
> Golang

> MySQL

> SSH

</br>

## 📚: Libraries Used
> [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2)
