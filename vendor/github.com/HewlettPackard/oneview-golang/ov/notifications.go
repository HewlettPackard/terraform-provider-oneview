package ov

import (
	"encoding/json"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type EmailNotificationList struct {
	AlertEmailDisabled bool            `json:"alertEmailDisabled,omitempty"`
	AlertEmailFilters  []utils.Nstring `json:"alertEmailFilters,omitempty"`
	Category           string          `json:"category,omitempty"`
	URI                utils.Nstring   `json:"uri,omitempty"`
	ETAG               string          `json:"eTag,omitempty"`
	Modified           string          `json:"modified,omitempty"`
	Created            string          `json:"created,omitempty"`
	Type               string          `json:"type,omitempty"`
	Password           utils.Nstring   `json:"password,omitempty"`
	SenderEmailAddress utils.Nstring   `json:"senderEmailAddress,omitempty"`
	SmtpPort           int             `json:"smtpPort,omitempty"`
	SmtpServer         utils.Nstring   `json:"smtpServer,omitempty"`
	SmtpProtocol       string          `json:"smtpProtocol,omitempty"`
}

type AlertEmailFilters struct {
	Disabled        bool            `json:"disabled,omitempty"`
	DisplayFilter   string          `json:"displayFilter,omitempty"`
	Emails          []utils.Nstring `json:"emails,omitempty"`
	Filter          string          `json:"filter,omitempty"`
	UserQueryFilter string          `json:"userQueryFilter,omitempty"`
	Limit           int             `json:"limit,omitempty"`
	LimitDuration   string          `json:"limitDuration,omitempty"`
	ScopeQuery      utils.Nstring   `json:"scopeQuery,omitempty"`
	FilterName      int             `json:"filterName,omitempty"`
}

type EmailFilterList struct {
	FilterName []utils.Nstring `json:"filterName,omitempty"`
}

type TestEmailResponse struct {
	Category               string        `json:"category,omitempty"`
	URI                    utils.Nstring `json:"uri,omitempty"`
	ETAG                   string        `json:"eTag,omitempty"`
	Modified               string        `json:"modified,omitempty"`
	Created                string        `json:"created,omitempty"`
	Type                   string        `json:"type,omitempty"`
	Password               utils.Nstring `json:"password,omitempty"`
	SenderEmailAddress     utils.Nstring `json:"senderEmailAddress,omitempty"`
	SmtpPort               int           `json:"smtpPort,omitempty"`
	UserProvidedSmtpServer utils.Nstring `json:"userProvidedSmtpServer,omitempty"`
	PreferredSmtpServer    utils.Nstring `json:"preferredSmtpServer,omitempty"`
}

type TestEmailRequest struct {
	HtmlMessageBody utils.Nstring   `json:"htmlMessageBody,omitempty"`
	Subject         utils.Nstring   `json:"subject,omitempty"`
	TextMessageBody utils.Nstring   `json:"textMessageBody,omitempty"`
	ToAddress       []utils.Nstring `json:"toAddress,omitempty"`
}

func (c *OVClient) GetEmailNotifications(filter string, sort string, start string, count string) (EmailNotificationList, error) {
	var (
		uri                = "/rest/appliance/notifications/email-config/"
		q                  = make(map[string]interface{})
		emailNotifications EmailNotificationList
	)

	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
	}

	if start != "" {
		q["start"] = start
	}

	if count != "" {
		q["count"] = count
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return emailNotifications, err
	}

	log.Debugf("GetEmailNotifications %s", data)
	if err := json.Unmarshal(data, &emailNotifications); err != nil {
		return emailNotifications, err
	}
	return emailNotifications, nil
}

func (c *OVClient) GetEmailNotificationsByFilter(filter string, sort string, start string, count string) (EmailFilterList, error) {
	var (
		uri          = "/rest/appliance/notifications/email-config/filters"
		q            = make(map[string]interface{})
		emailFilters EmailFilterList
	)

	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
	}

	if start != "" {
		q["start"] = start
	}

	if count != "" {
		q["count"] = count
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return emailFilters, err
	}

	log.Debugf("GetEmailNotificationsByFilter %s", data)
	if err := json.Unmarshal(data, &emailFilters); err != nil {
		return emailFilters, err
	}
	return emailFilters, nil
}

func (c *OVClient) GetEmailNotificationsConfiguration(filter string, sort string, start string, count string) (TestEmailResponse, error) {
	var (
		uri           = "/rest/appliance/notifications/test-email-config"
		q             = make(map[string]interface{})
		emailResponse TestEmailResponse
	)

	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
	}

	if start != "" {
		q["start"] = start
	}

	if count != "" {
		q["count"] = count
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return emailResponse, err
	}

	log.Debugf("GetEmailNotificationsConfiguration %s", data)
	if err := json.Unmarshal(data, &emailResponse); err != nil {
		return emailResponse, err
	}
	return emailResponse, nil
}

func (c *OVClient) SendTestEmail(email TestEmailRequest) error {
	log.Infof("Initialization of sending test email ")
	var (
		uri = "/rest/appliance/notifications/test-email"
		t   = (&Task{}).NewProfileTask(c)
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, email)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, email)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new test email request: %s", err)
		return err
	}

	log.Debugf("Response after sending test email %s", data)
	if err := json.Unmarshal(data, &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	err = t.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (c *OVClient) SendEmail(email TestEmailRequest) error {
	log.Infof("Initialization of sending email ")
	var (
		uri = "/rest/appliance/notifications/send-email"
		t   = (&Task{}).NewProfileTask(c)
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, email)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, email)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new email request: %s", err)
		return err
	}

	log.Debugf("Response after sending email %s", data)
	if err := json.Unmarshal(data, &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	err = t.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (c *OVClient) ConfigureAppliance(configuration EmailNotificationList) error {
	log.Infof("Initialization of sending email ")
	var (
		uri = "/rest/appliance/notifications/email-config"
		t   = (&Task{}).NewProfileTask(c)
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, configuration)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, configuration)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new email configuraion request: %s", err)
		return err
	}

	log.Debugf("Response after sending email configuration %s", data)
	if err := json.Unmarshal(data, &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	err = t.Wait()
	if err != nil {
		return err
	}

	return nil
}
