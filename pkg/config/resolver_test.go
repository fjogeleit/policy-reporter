package config_test

import (
	"context"
	"testing"

	"github.com/fjogeleit/policy-reporter/pkg/config"
	"k8s.io/client-go/rest"
)

var testConfig = &config.Config{
	Loki: config.Loki{
		Host:            "http://localhost:3100",
		SkipExisting:    true,
		MinimumPriority: "debug",
	},
	Elasticsearch: config.Elasticsearch{
		Host:            "http://localhost:9200",
		Index:           "policy-reporter",
		Rotation:        "dayli",
		SkipExisting:    true,
		MinimumPriority: "debug",
	},
	Slack: config.Slack{
		Webhook:         "http://hook.slack:80",
		SkipExisting:    true,
		MinimumPriority: "debug",
	},
	Discord: config.Discord{
		Webhook:         "http://hook.discord:80",
		SkipExisting:    true,
		MinimumPriority: "debug",
	},
}

func Test_ResolveTarget(t *testing.T) {
	resolver := config.NewResolver(testConfig, nil)

	t.Run("Loki", func(t *testing.T) {
		client := resolver.LokiClient()
		if client == nil {
			t.Error("Expected Client, got nil")
		}

		client2 := resolver.LokiClient()
		if client != client2 {
			t.Error("Error: Should reuse first instance")
		}
	})
	t.Run("Elasticsearch", func(t *testing.T) {
		client := resolver.ElasticsearchClient()
		if client == nil {
			t.Error("Expected Client, got nil")
		}

		client2 := resolver.ElasticsearchClient()
		if client != client2 {
			t.Error("Error: Should reuse first instance")
		}
	})
	t.Run("Slack", func(t *testing.T) {
		client := resolver.SlackClient()
		if client == nil {
			t.Error("Expected Client, got nil")
		}

		client2 := resolver.SlackClient()
		if client != client2 {
			t.Error("Error: Should reuse first instance")
		}
	})
	t.Run("Discord", func(t *testing.T) {
		client := resolver.DiscordClient()
		if client == nil {
			t.Error("Expected Client, got nil")
		}

		client2 := resolver.DiscordClient()
		if client != client2 {
			t.Error("Error: Should reuse first instance")
		}
	})
}

func Test_ResolveTargets(t *testing.T) {
	resolver := config.NewResolver(testConfig, nil)

	clients := resolver.TargetClients()
	if count := len(clients); count != 4 {
		t.Errorf("Expected 4 Clients, got %d", count)
	}
}

func Test_ResolveSkipExistingOnStartup(t *testing.T) {
	var testConfig = &config.Config{
		Loki: config.Loki{
			Host:            "http://localhost:3100",
			SkipExisting:    true,
			MinimumPriority: "debug",
		},
		Elasticsearch: config.Elasticsearch{
			Host:            "http://localhost:9200",
			SkipExisting:    true,
			MinimumPriority: "debug",
		},
	}

	t.Run("Resolve false", func(t *testing.T) {
		testConfig.Elasticsearch.SkipExisting = false

		resolver := config.NewResolver(testConfig, nil)

		if resolver.SkipExistingOnStartup() == true {
			t.Error("Expected SkipExistingOnStartup to be false if one Client has SkipExistingOnStartup false configured")
		}
	})

	t.Run("Resolve true", func(t *testing.T) {
		testConfig.Elasticsearch.SkipExisting = true

		resolver := config.NewResolver(testConfig, nil)

		if resolver.SkipExistingOnStartup() == false {
			t.Error("Expected SkipExistingOnStartup to be true if all Client has SkipExistingOnStartup true configured")
		}
	})
}

func Test_ResolveTargetWithoutHost(t *testing.T) {
	config2 := &config.Config{
		Loki: config.Loki{
			Host:            "",
			SkipExisting:    true,
			MinimumPriority: "debug",
		},
		Elasticsearch: config.Elasticsearch{
			Host:            "",
			Index:           "policy-reporter",
			Rotation:        "dayli",
			SkipExisting:    true,
			MinimumPriority: "debug",
		},
		Slack: config.Slack{
			Webhook:         "",
			SkipExisting:    true,
			MinimumPriority: "debug",
		},
		Discord: config.Discord{
			Webhook:         "",
			SkipExisting:    true,
			MinimumPriority: "debug",
		},
	}

	t.Run("Loki", func(t *testing.T) {
		resolver := config.NewResolver(config2, nil)

		if resolver.LokiClient() != nil {
			t.Error("Expected Client to be nil if no host is configured")
		}
	})
	t.Run("Elasticsearch", func(t *testing.T) {
		resolver := config.NewResolver(config2, nil)

		if resolver.ElasticsearchClient() != nil {
			t.Error("Expected Client to be nil if no host is configured")
		}
	})
	t.Run("Slack", func(t *testing.T) {
		resolver := config.NewResolver(config2, nil)

		if resolver.SlackClient() != nil {
			t.Error("Expected Client to be nil if no host is configured")
		}
	})
	t.Run("Discord", func(t *testing.T) {
		resolver := config.NewResolver(config2, nil)

		if resolver.DiscordClient() != nil {
			t.Error("Expected Client to be nil if no host is configured")
		}
	})
}

func Test_ResolveResultClient(t *testing.T) {
	resolver := config.NewResolver(&config.Config{}, &rest.Config{})

	client1, err := resolver.PolicyResultClient(context.Background())
	if err != nil {
		t.Errorf("Unexpected Error: %s", err)
	}

	client2, err := resolver.PolicyResultClient(context.Background())
	if client1 != client2 {
		t.Error("A second call resolver.PolicyResultClient() should return the cached first client")
	}
}

func Test_ResolvePolicyClient(t *testing.T) {
	resolver := config.NewResolver(&config.Config{}, &rest.Config{})

	client1, err := resolver.PolicyReportClient(context.Background())
	if err != nil {
		t.Errorf("Unexpected Error: %s", err)
	}

	client2, err := resolver.PolicyReportClient(context.Background())
	if client1 != client2 {
		t.Error("A second call resolver.PolicyReportClient() should return the cached first client")
	}
}

func Test_ResolveClusterPolicyClient(t *testing.T) {
	resolver := config.NewResolver(&config.Config{}, &rest.Config{})

	client1, err := resolver.ClusterPolicyReportClient(context.Background())
	if err != nil {
		t.Errorf("Unexpected Error: %s", err)
	}

	client2, err := resolver.ClusterPolicyReportClient(context.Background())
	if client1 != client2 {
		t.Error("A second call resolver.ClusterPolicyReportClient() should return the cached first client")
	}
}

func Test_ResolveAPIServer(t *testing.T) {
	resolver := config.NewResolver(testConfig, &rest.Config{})

	server := resolver.APIServer()
	if server == nil {
		t.Error("Error: Should return API Server")
	}
}

func Test_ResolveClientWithInvalidK8sConfig(t *testing.T) {
	k8sConfig := &rest.Config{}
	k8sConfig.Host = "invalid/url"

	resolver := config.NewResolver(&config.Config{}, k8sConfig)

	_, err := resolver.PolicyReportClient(context.Background())
	if err == nil {
		t.Error("Error: 'host must be a URL or a host:port pair' was expected")
	}
}
