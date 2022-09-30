package insightcloudsec

var _ CloudOrgs = (*cloud_orgs)(nil)

type CloudOrgs interface {
	List()
}

type cloud_orgs struct {
	client *Client
}

func (o *cloud_orgs) List() {

}
