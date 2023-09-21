package api

import "time"

type AlertResponse struct {
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
	Alerts   []struct {
		Status string `json:"status"`
		Labels struct {
			DeviceHostname               string `json:"deviceHostname"`
			DeviceID                     string `json:"deviceId"`
			DeviceTagContainer           string `json:"deviceTag_Container"`
			DeviceTagBgp                 string `json:"deviceTag_bgp"`
			DeviceTagDevice              string `json:"deviceTag_device"`
			DeviceTagEos                 string `json:"deviceTag_eos"`
			DeviceTagEostrain            string `json:"deviceTag_eostrain"`
			DeviceTagHostname            string `json:"deviceTag_hostname"`
			DeviceTagModel               string `json:"deviceTag_model"`
			DeviceTagMpls                string `json:"deviceTag_mpls"`
			DeviceTagSerialnumber        string `json:"deviceTag_serialnumber"`
			DeviceTagSystype             string `json:"deviceTag_systype"`
			DeviceTagTapagg              string `json:"deviceTag_tapagg"`
			DeviceTagTerminattr          string `json:"deviceTag_terminattr"`
			DeviceTagTopologyNetworkType string `json:"deviceTag_topology_network_type"`
			DeviceTagType                string `json:"deviceTag_type"`
			DeviceTagZtp                 string `json:"deviceTag_ztp"`
			EventKey                     string `json:"eventKey"`
			EventType                    string `json:"eventType"`
			InterfaceID                  string `json:"interfaceId"`
			InterfaceTagDevice           string `json:"interfaceTag_device"`
			InterfaceTagInterface        string `json:"interfaceTag_interface"`
			InterfaceTagName             string `json:"interfaceTag_name"`
			InterfaceTagSpeed            string `json:"interfaceTag_speed"`
			InterfaceTagType             string `json:"interfaceTag_type"`
			Severity                     string `json:"severity"`
			TurbineID                    string `json:"turbineId"`
		} `json:"labels"`
		Annotations struct {
			Description string `json:"description"`
			Title       string `json:"title"`
		} `json:"annotations"`
		StartsAt     time.Time `json:"startsAt"`
		EndsAt       time.Time `json:"endsAt"`
		GeneratorURL string    `json:"generatorURL"`
		Fingerprint  string    `json:"fingerprint"`
	} `json:"alerts"`
	GroupLabels struct {
	} `json:"groupLabels"`
	CommonLabels struct {
		DeviceHostname               string `json:"deviceHostname"`
		DeviceID                     string `json:"deviceId"`
		DeviceTagContainer           string `json:"deviceTag_Container"`
		DeviceTagBgp                 string `json:"deviceTag_bgp"`
		DeviceTagDevice              string `json:"deviceTag_device"`
		DeviceTagEos                 string `json:"deviceTag_eos"`
		DeviceTagEostrain            string `json:"deviceTag_eostrain"`
		DeviceTagHostname            string `json:"deviceTag_hostname"`
		DeviceTagModel               string `json:"deviceTag_model"`
		DeviceTagMpls                string `json:"deviceTag_mpls"`
		DeviceTagSerialnumber        string `json:"deviceTag_serialnumber"`
		DeviceTagSystype             string `json:"deviceTag_systype"`
		DeviceTagTapagg              string `json:"deviceTag_tapagg"`
		DeviceTagTerminattr          string `json:"deviceTag_terminattr"`
		DeviceTagTopologyNetworkType string `json:"deviceTag_topology_network_type"`
		DeviceTagType                string `json:"deviceTag_type"`
		DeviceTagZtp                 string `json:"deviceTag_ztp"`
		EventType                    string `json:"eventType"`
		InterfaceTagDevice           string `json:"interfaceTag_device"`
		InterfaceTagSpeed            string `json:"interfaceTag_speed"`
		InterfaceTagType             string `json:"interfaceTag_type"`
		Severity                     string `json:"severity"`
		TurbineID                    string `json:"turbineId"`
	} `json:"commonLabels"`
	CommonAnnotations struct {
		Title string `json:"title"`
	} `json:"commonAnnotations"`
	ExternalURL     string `json:"externalURL"`
	Version         string `json:"version"`
	GroupKey        string `json:"groupKey"`
	TruncatedAlerts int    `json:"truncatedAlerts"`
}
