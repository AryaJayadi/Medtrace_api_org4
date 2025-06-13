package main

import (
	"log"
	"os"

	"github.com/AryaJayadi/MedTrace_api_org4/cmd/fabric"
	"github.com/AryaJayadi/MedTrace_api_org4/internal/handlers"
	"github.com/AryaJayadi/MedTrace_api_org4/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	OrgName      = "Org1"
	mspID        = "Org1MSP"
	cryptoPath   = "../../../MedTrace_network/organizations/peerOrganizations/org1.medtrace.com"
	certPath     = cryptoPath + "/users/User1@org1.medtrace.com/msp/signcerts/User1@org1.medtrace.com-cert.pem"
	keyPath      = cryptoPath + "/users/User1@org1.medtrace.com/msp/keystore"
	tlsCertPath  = cryptoPath + "/peers/peer0.org1.medtrace.com/tls/ca.crt"
	peerEndpoint = "dns:///localhost:7051"
	gatewayPeer  = "peer0.org1.medtrace.com"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	orgConfig := fabric.OrgSetup{
		OrgName:      OrgName,
		MSPID:        mspID,
		CertPath:     certPath,
		KeyPath:      keyPath,
		TLSCertPath:  tlsCertPath,
		PeerEndpoint: peerEndpoint,
		GatewayPeer:  gatewayPeer,
	}
	orgSetup, err := fabric.Initialize(orgConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Fabric Org: %v", err)
	}

	chaincodeName := "medtrace_cc"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	channelName := "medtrace"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	network := orgSetup.Gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	drugService := services.NewDrugService(contract)
	organizationService := services.NewOrganizationService(contract)

	drugHandler := handlers.NewDrugHandler(drugService)
	organizationHandler := handlers.NewOrganizationHandler(organizationService)

	drugsGroup := e.Group("/drugs")
	drugsGroup.GET("/history/:drugID", drugHandler.GetHistoryDrug)

	organizationsGroup := e.Group("/organizations")
	organizationsGroup.GET("/", organizationHandler.GetOrganizations)

	e.Logger.Fatal(e.Start(":9090"))
}
