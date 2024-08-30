package padmin

import "fmt"

type LookupImpl struct {
	cli HttpClient
}

func (l *LookupImpl) GetOwner(topicDomain TopicDomain, tenant, namespace, topic string) (*LookupData, error) {
	url := fmt.Sprintf(UrlLookupBrokerFormat, topicDomain, tenant, namespace, topic)
	resp, err := l.cli.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body = new(LookupData)
	if err := EasyReader(resp, body); err != nil {
		return nil, err
	}
	return body, nil
}

func (l *LookupImpl) GetNamespaceBundle(topicDomain TopicDomain, tenant, namespace, topic string) (string, error) {
	url := fmt.Sprintf(UrlLookupGetNamespaceBundleFormat, topicDomain, tenant, namespace, topic)
	resp, err := l.cli.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return res, nil
}

func newLookup(cli HttpClient) Lookup {
	return &LookupImpl{cli: cli}
}
