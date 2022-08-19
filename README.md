## Architecture 🏞️

![https://i.postimg.cc/k4xPqtBF/Screen-Shot-2022-06-03-at-1-36-41-PM.pngapp](https://i.postimg.cc/k4xPqtBF/Screen-Shot-2022-06-03-at-1-36-41-PM.png)

Sample REST API for crowdfunding system build on top of [GIN](https://gin-gonic.com/) Framework, MariaDB combine with [GORM](https://gorm.io/index.html) and integrated with [Midtrans](https://midtrans.com/) payment gateway.

## Objective ⭐

Simple Crowdfunding platforms that enable interaction between fundraisers and the crowd. Funders can give funds to the project they want to help and fundraisers can explain what benefits the funder will get if the project is successful.

## Example Frontend (if any)🪴

![fe2](https://i.postimg.cc/hjBkK4Qv/fe2.jpg)
![fe](https://i.postimg.cc/ydywxq0P/fe.jpg)

## Main Feature 🧩

- auth jwt (register,login, email checker)
- upload user avatar
- get, create, update, detail campaign
- upload campaign images
- get campaign & user transaction
- create trasaction & give payment link (midtrans)
- admin web dashboard for monitoring user, campaign & transaction data

## Step to run the app 🎖️

- setup env for database & midtrans connection. You can see docs folder for see dummy database & Postman Documentation
- install [golang](https://go.dev/dl/),move to project folder & type: `go run main.go`
- you can see my blog post [Golang-Deployment](https://www.faeshal.com/posts/deploy-golang-on-linux-server-2ecedcde6752/) to see how to deploy go on a linux server

🗡 **August 2022 - [faeshal.com](https://faeshal.com)**
