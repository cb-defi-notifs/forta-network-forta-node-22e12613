package containers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/forta-network/forta-node/clients"
	"github.com/forta-network/forta-node/clients/docker"
	"github.com/forta-network/forta-node/config"
	log "github.com/sirupsen/logrus"
)

// Timeouts
const (
	BotPullTimeout  = time.Minute * 5
	BotStartTimeout = time.Minute * 5
)

// BotClient launches a bot.
type BotClient interface {
	EnsureBotImages(ctx context.Context, botConfigs []config.AgentConfig) []error
	LaunchBot(ctx context.Context, botConfig config.AgentConfig) error
	ShutDownBot(ctx context.Context, botConfig config.AgentConfig) error
	LoadBotContainers(ctx context.Context) ([]types.Container, error)
	StartWaitBotContainer(ctx context.Context, containerID string) error
}

type botClient struct {
	logConfig       config.LogConfig
	resourcesConfig config.ResourcesConfig
	client          clients.DockerClient
	botImageClient  clients.DockerClient
}

// NewBotClient creates a new bot client to manage bot containers.
func NewBotClient(
	logConfig config.LogConfig, resourcesConfig config.ResourcesConfig,
	client clients.DockerClient, botImageClient clients.DockerClient,
) BotClient {
	return &botClient{
		logConfig:       logConfig,
		resourcesConfig: resourcesConfig,
		client:          client,
		botImageClient:  botImageClient,
	}
}

// EnsureBotImages ensures that all of the bot images are locally available.
func (bc *botClient) EnsureBotImages(ctx context.Context, botConfigs []config.AgentConfig) []error {
	var imagePulls []docker.ImagePull
	for _, botConfig := range botConfigs {
		imagePulls = append(imagePulls, docker.ImagePull{
			Name: botConfig.ID,
			Ref:  botConfig.Image,
		})
	}
	return bc.client.EnsureLocalImages(ctx, BotPullTimeout, imagePulls)
}

// LaunchBot launches a bot by downloading docker image and starting the container.
func (bc *botClient) LaunchBot(ctx context.Context, botConfig config.AgentConfig) error {
	ctx, cancel := context.WithTimeout(ctx, BotStartTimeout)
	defer cancel()

	_, err := bc.client.GetContainerByName(ctx, botConfig.ContainerName())
	if !errors.Is(err, docker.ErrContainerNotFound) {
		log.WithField("container", botConfig.ContainerName()).Info("bot container exists - skipping launch")
		return nil
	}

	nwID, err := bc.client.CreatePublicNetwork(ctx, botConfig.ContainerName())
	if err != nil {
		return err
	}

	botContainerCfg := NewBotContainerConfig(nwID, botConfig, bc.logConfig, bc.resourcesConfig)
	botContainerCfg.LinkNetworkIDs, err = bc.getServiceContainerIDs(ctx)
	if err != nil {
		return err
	}

	_, err = bc.client.StartContainer(ctx, botContainerCfg)
	return err
}

func (bc *botClient) getServiceContainerIDs(ctx context.Context) (ids []string, err error) {
	for _, containerName := range []string{
		config.DockerScannerContainerName, config.DockerJSONRPCProxyContainerName,
		config.DockerJWTProviderContainerName, config.DockerPublicAPIProxyContainerName,
	} {
		container, err := bc.client.GetContainerByName(ctx, containerName)
		if errors.Is(err, docker.ErrContainerNotFound) {
			return nil, fmt.Errorf("failed to get service container ids while launching the bot: %v", err)
		}
		ids = append(ids, container.ID)
	}
	return ids, nil
}

// ShutDownBot shuts down a bot by shutting down the docker container.
func (bc *botClient) ShutDownBot(ctx context.Context, botConfig config.AgentConfig) error {
	container, err := bc.client.GetContainerByName(ctx, botConfig.ContainerName())
	if err != nil {
		return fmt.Errorf("failed to get the bot container to shut down: %v", err)
	}
	return bc.client.StopContainer(ctx, container.ID)
}

// LoadBotContainers loads the latest bot list for the running scanner.
func (bc *botClient) LoadBotContainers(ctx context.Context) ([]types.Container, error) {
	return bc.client.GetContainersByLabel(ctx, docker.LabelFortaIsBot, LabelValueFortaIsBot)
}

// StartWaitBotContainer starts the bot container and waits.
func (bc *botClient) StartWaitBotContainer(ctx context.Context, containerID string) error {
	if err := bc.client.StartContainerWithID(ctx, containerID); err != nil {
		return fmt.Errorf("failed to start container with id: %v", err)
	}
	return bc.client.WaitContainerStart(ctx, containerID)
}
