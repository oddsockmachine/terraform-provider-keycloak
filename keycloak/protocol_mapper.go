package keycloak

import "fmt"

// https://www.keycloak.org/docs-api/4.2/rest-api/index.html#_protocolmapperrepresentation
type protocolMapper struct {
	Id             string            `json:"id,omitempty"`
	Name           string            `json:"name"`
	Protocol       string            `json:"protocol"`
	ProtocolMapper string            `json:"protocolMapper"`
	Config         map[string]string `json:"config"`
}

var (
	addToAccessTokenField = "access.token.claim"
	addToIdTokenField     = "id.token.claim"
	addToUserInfoField    = "userinfo.token.claim"
	claimNameField        = "claim.name"
	claimValueField       = "claim.value"
	claimValueTypeField   = "jsonType.label"
	fullPathField         = "full.path"
	multivaluedField      = "multivalued"
	userAttributeField    = "user.attribute"
	userPropertyField     = "user.attribute"
)

func protocolMapperPath(realmId, clientId, clientScopeId string) string {
	parentResourceId := clientId
	parentResourcePath := "clients"

	if clientScopeId != "" {
		parentResourceId = clientScopeId
		parentResourcePath = "client-scopes"
	}

	return fmt.Sprintf("/realms/%s/%s/%s/protocol-mappers/models", realmId, parentResourcePath, parentResourceId)
}

func individualProtocolMapperPath(realmId, clientId, clientScopeId, mapperId string) string {
	return fmt.Sprintf("%s/%s", protocolMapperPath(realmId, clientId, clientScopeId), mapperId)
}

func (keycloakClient *KeycloakClient) listGenericProtocolMappers(realmId, clientId, clientScopeId string) ([]*protocolMapper, error) {
	var protocolMappers []*protocolMapper

	err := keycloakClient.get(protocolMapperPath(realmId, clientId, clientScopeId), &protocolMappers)
	if err != nil {
		return nil, err
	}

	return protocolMappers, nil
}