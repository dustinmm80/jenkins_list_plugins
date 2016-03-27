package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type Plugin struct {
	ShortName string `json:"shortName"`
	LongName  string `json:"longName"`
	Version   string `json:"version"`
	HasUpdate bool   `json:"hasUpdate"`
	Active    bool   `json:"active"`
}

type PluginList []Plugin

type ByName []Plugin

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].ShortName < n[j].ShortName }

func ListPlugins(url, username, password string, skipVerify bool) (PluginList, error) {
	body, err := doRequest(url, username, password, skipVerify)
	if err != nil {
		return PluginList{}, err
	}
	resp := pluginResponse{}
	json.Unmarshal(body, &resp)
	plugins := resp.Plugins
	sort.Sort(ByName(plugins))
	return plugins, nil
}

type pluginResponse struct {
	Plugins PluginList `json:"plugins"`
}

func doRequest(url, username, password string, skipVerify bool) ([]byte, error) {
	fullUrl := fmt.Sprintf("%s/pluginManager/api/json?depth=1", url)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", fullUrl, nil)
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
