package insightcloudsec

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var _ Badges = (*badges)(nil)

type Badges interface {
	Create(target_org_resource_ids []string, badge_data map[string]string) error
	Update(org_resource_id string, badges map[string]string) error
	Delete(target_org_resource_ids []string, badges map[string]string) error
	ListCloudsWithBadges(badges map[string]string) ([]CloudBadges, error)
	ListResourceBadges(org_resource_id string) ([]Badge, error)
	ListResourceBadgeCount(resource_ids []string) (Resource_Count, error)
}

type badges struct {
	client *Client
}

type Badge struct {
	// The key and value of a given badge for use with filters, insights, etc.
	Key            string `json:"key"`
	Value          string `json:"value"`
	Auto_Generated bool   `json:"auto_generated,omitempty"`
}

type BadgesList struct {
	Badges []Badge `json:"badges"`
}

type BadgeRequest struct {
	Org_Resource_IDs []string `json:"target_resource_ids"`
	Badges           []Badge  `json:"badges"`
}

type Resource_Count struct {
	Resource_Count []BadgeResourceCount `json:"resource_count"`
}

type BadgeResourceCount struct {
	Resource_ID string `json:"resource_id"`
	Count       int32  `json:"count"`
}

type BadgeResourceCountRequest struct {
	Resource_IDs []string `json:"resource_ids"`
}

type CloudBadges struct {
	Resource_ID string `json:"resource_id"`
	Name        string `json:"name"`
}

func (s *badges) Create(target_org_resource_ids []string, badge_data map[string]string) error {
	// Creates a badge for target organization resource ids of key and value pairings provided in map

	_, err := s.client.makeRequest(http.MethodPost, "/v2/public/badges/create", BadgeRequest{Org_Resource_IDs: target_org_resource_ids, Badges: createBadgesFromMap(badge_data)})
	if err != nil {
		return err
	}

	return nil
}

func (s *badges) Update(org_resource_id string, badges map[string]string) error {
	// Updates cloud badges for given organization but overwrites any existing. USE WITH CAUTION

	_, err := s.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/badges/%s/update", org_resource_id), createBadgesFromMap(badges))
	if err != nil {
		return err
	}

	return nil
}

func (s *badges) Delete(target_org_resource_ids []string, badges map[string]string) error {
	// Deletes given list of badges defined as a map of key/values.

	_, err := s.client.makeRequest(http.MethodPost, "/v2/public/badges/delete", BadgeRequest{Org_Resource_IDs: target_org_resource_ids, Badges: createBadgesFromMap(badges)})
	if err != nil {
		return err
	}

	return nil
}

func (s *badges) ListCloudsWithBadges(badges map[string]string) ([]CloudBadges, error) {
	// Returns a list of cloud accounts what contain the given badges defined as a map of key / values.

	resp, err := s.client.makeRequest(http.MethodPost, "/v2/public/badge/clouds/list", badges)
	if err != nil {
		return []CloudBadges{}, err
	}

	var clouds []CloudBadges
	if err := json.NewDecoder(resp.Body).Decode(&clouds); err != nil {
		return []CloudBadges{}, err
	}

	return clouds, nil
}

func (s *badges) ListResourceBadges(org_resource_id string) ([]Badge, error) {
	// Returns a list of resource badges for a given organization

	resp, err := s.client.makeRequest(http.MethodPost, fmt.Sprintf("/v2/public/badges/%s/list", org_resource_id), nil)
	if err != nil {
		return []Badge{}, err
	}

	var ret []Badge
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return []Badge{}, err
	}

	return ret, nil
}

func (s *badges) ListResourceBadgeCount(resource_ids []string) (Resource_Count, error) {
	// Returns a list of badge counts for all resources.

	resp, err := s.client.makeRequest(http.MethodPost, "/v2/public/badges/count", BadgeResourceCountRequest{Resource_IDs: resource_ids})
	if err != nil {
		return Resource_Count{}, err
	}

	var ret Resource_Count
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return Resource_Count{}, err
	}

	return ret, nil
}

func createBadgesFromMap(m map[string]string) []Badge {
	badges := []Badge{}
	for badge, value := range m {
		item := Badge{
			Key:   badge,
			Value: value,
		}
		badges = append(badges, item)
	}

	return badges
}
