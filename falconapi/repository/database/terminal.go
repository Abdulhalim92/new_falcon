package database

import (
	"context"
	"falconapi/domain/entities"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log"
	"time"
)

func (db *database) GetAll(ctx context.Context) ([]entities.TerminalStatus, error) {
	var (
		terminalStatuses []entities.TerminalStatus
	)

	query := db.dbConn.WithContext(ctx).
		Select("t.endpoint_id,t.phone,t.region_id, t.endpoint_num,  t.address , max(p.created_at) as  last_created_payment, pd.lastping,pd.status,t.endpoint_disabled").
		Table("tendpoints t").
		Joins("LEFT OUTER  JOIN tpayments p ON p.agent_term_id=t.endpoint_id").
		Joins("LEFT OUTER JOIN (SELECT DISTINCT ON (endpoint_id) endpoint_id,created_at as lastping,status FROM tendpointpings WHERE created_at >= NOW() - INTERVAL '60 DAYS' ORDER BY endpoint_id, created_at desc,status) pd ON pd.endpoint_id = t.endpoint_id").
		Where("t.type = 100").
		Group("t.endpoint_num, t.endpoint_id,  t.address, pd.lastping,pd.status").
		Order("t.endpoint_id asc")

	if err := query.Scan(&terminalStatuses).Error; err != nil {
		db.log.Errorf("unable to get terminal statuses: %v", err)
		return nil, errors.Wrap(err, "unable to get terminal statuses")
	}

	return terminalStatuses, nil
}

func (db *database) GetAllWithRegionsNames(ctx context.Context) ([]entities.TerminalStatus, error) {
	var (
		terminalStatuses []entities.TerminalStatus
	)

	query := db.dbConn.WithContext(ctx).
		Select("t.endpoint_id,t.phone,t.region_id, r.name, t.endpoint_num,  t.address , max(p.created_at) as  last_created_payment, pd.lastping,pd.status,t.endpoint_disabled").
		Table("tendpoints t").
		Joins("LEFT OUTER JOIN tregions r ON r.id=t.region_id").
		Joins("LEFT OUTER  JOIN tpayments p ON p.agent_term_id=t.endpoint_id").
		Joins("LEFT OUTER JOIN (SELECT DISTINCT ON (endpoint_id) endpoint_id,created_at as lastping,status FROM tendpointpings WHERE created_at >= NOW() - INTERVAL '60 DAYS' ORDER BY endpoint_id, created_at desc,status) pd ON pd.endpoint_id = t.endpoint_id").
		Where("t.type = 100").
		Group("t.endpoint_num, t.endpoint_id,  t.address, pd.lastping,pd.status").
		Order("t.endpoint_id asc")

	if err := query.Scan(&terminalStatuses).Error; err != nil {
		db.log.Errorf("unable to get terminal statuses: %v", err)
		return nil, errors.Wrap(err, "unable to get terminal statuses")
	}

	return terminalStatuses, nil
}

func (db *database) GetRegions(ctx context.Context) ([]entities.TRegion, error) {
	var (
		region []entities.TRegion
	)

	if err := db.dbConn.WithContext(ctx).Find(&region).Error; err != nil {
		db.log.Errorf("unable to find regions: %v", err)
		return nil, errors.Wrap(err, "unable to find regions")
	}

	return region, nil
}

func (db *database) GetTerminalByID(ctx context.Context, id string) (*entities.TEndpoint, error) {
	var endpoint entities.TEndpoint

	if err := db.dbConn.WithContext(ctx).Where("endpoint_id = ?", id).Find(&endpoint).Error; err != nil {
		db.log.Errorf("unable to get terminal by ID: %v", err)
		return nil, errors.Wrap(err, "unable to get terminal by ID")
	}

	return &endpoint, nil
}

func (db *database) GenerateTermPassword(ctx context.Context, userID string, endpoint entities.TEndpoint) (*entities.TEndpoint, error) {
	var tendpoint entities.TEndpoint

	qr := db.dbConn.WithContext(ctx).Where("endpoint_id = ?", endpoint.ID).Find(&tendpoint).RowsAffected
	if qr == 0 {
		db.log.Errorf("terminal with ID %v not found", endpoint.ID)
		return nil, fmt.Errorf("terminal with ID %v not found", endpoint.ID)
	}

	err := db.dbConn.WithContext(ctx).Where("endpoint_id = ?", endpoint.ID).
		Model(&entities.TEndpoint{}).
		UpdateColumns(map[string]interface{}{
			"secretkey":  endpoint.SecretKey,
			"apiKey":     endpoint.APIKey,
			"updated_at": time.Now(),
			"updated_by": userID,
		}).Error
	if err != nil {
		db.log.Errorf("unable to update: %v", err)
		return nil, err
	}

	return &endpoint, err
}

func (db *database) CountTerminalByNum(ctx context.Context, incomingTerminalData entities.TerminalCreate) (int64, error) {
	var endpoint entities.TEndpoint

	endpointCount := db.dbConn.WithContext(ctx).Where("endpoint_num = ?", incomingTerminalData.Num).Find(&endpoint).RowsAffected

	return endpointCount, nil
}

func (db *database) NewTerminal(ctx context.Context, endpoint entities.TEndpoint) (*entities.TEndpoint, error) {
	if err := db.dbConn.WithContext(ctx).Create(&endpoint).Error; err != nil {
		db.log.Errorf("unable to create new terminal: %v", err)
		return nil, fmt.Errorf("unable to create new terminal: %v", err)
	}

	return &endpoint, nil
}

func (db *database) UpdateTerminal(ctx context.Context, endpoint entities.TEndpoint) (*entities.TEndpoint, error) {
	if err := db.dbConn.WithContext(ctx).Save(&endpoint).Error; err != nil {
		db.log.Errorf("unable to update terminal: %v", err)
		return nil, fmt.Errorf("unable to update terminal: %v", err)
	}

	return &endpoint, nil
}

func (db *database) ChangeApi(ctx context.Context, newApi *uuid.UUID, endpoint entities.TEndpoint) (*entities.TEndpoint, error) {
	if err := db.dbConn.WithContext(ctx).Model(&endpoint).UpdateColumn("SecretKey", newApi.String()).Error; err != nil {
		db.log.Errorf("couldn't update controller key: %v", err)
		return nil, fmt.Errorf("couldn't update controller key: %v", err)
	}

	return &endpoint, nil
}

func (db *database) GetXMLGate(ctx context.Context) ([]entities.TEndpoint, error) {
	var endpoints []entities.TEndpoint

	err := db.dbConn.WithContext(ctx).Order("endpoint_id asc").Where(&entities.TEndpoint{Type: 900}).Find(&endpoints).Error
	if err != nil {
		db.log.Errorf("unable to get XML Gate: %v", err)
		return nil, fmt.Errorf("unable to get XML Gate: %v", err)
	}

	return endpoints, nil
}

func (db *database) CreateXMLGate(ctx context.Context, endpoint entities.TEndpoint) (*entities.TEndpoint, error) {
	if err := db.dbConn.WithContext(ctx).Create(&endpoint).Error; err != nil {
		db.log.Errorf("unable to create XML Gate: %v", err)
		return nil, fmt.Errorf("unable to create XML Gate: %v", err)
	}

	return &endpoint, nil
}

func (db *database) UpdateTerminalXML(ctx context.Context, endpoint entities.TEndpoint) (*entities.TEndpoint, error) {
	if err := db.dbConn.WithContext(ctx).Save(&endpoint).Error; err != nil {
		db.log.Errorf("unable to update terminal XML: %v", err)
		return nil, fmt.Errorf("unable to update terminal XML: %v", err)
	}

	return &endpoint, nil
}

func (db *database) GetStatusPingByID(ctx context.Context, id string) (*entities.StatusTerminal, error) {
	var status entities.StatusTerminal

	tx := db.dbConn.WithContext(ctx).Select("t.endpoint_id, t.status, t.created_at, t.param1").
		Table("tendpointpings t,(SELECT endpoin_id, max(created_at) AS lastping FROM public.tendpoints GROUP BY endpoint_id s").
		Where("t.created_at = s.lastping and t.endpoin_id = s.endpoint_id and t.endpoint.id = ?", id)

	if err := tx.Scan(&status).Error; err != nil {
		db.log.Errorf("unable to find status by id %v with error: %v", id, err)
		return nil, fmt.Errorf("unable to find status by id %v with error: %v", id, err)
	}

	return &status, nil
}

func (db *database) GetAllCommands(ctx context.Context) ([]entities.EndpointCommand, error) {
	var cmds []entities.EndpointCommand

	if err := db.dbConn.WithContext(ctx).Order("cmd_id desc").Limit(30).Find(&cmds).Error; err != nil {
		db.log.Errorf("unable tp find commands: %v", err)
		return nil, fmt.Errorf("unable tp find commands: %v", err)
	}

	return cmds, nil
}

func (db *database) CreateCommand(ctx context.Context, cmd entities.EndpointCommand) (*entities.EndpointCommand, error) {
	if err := db.dbConn.WithContext(ctx).Create(&cmd).Error; err != nil {
		db.log.Errorf("unable to create command: %v", err)
		return nil, fmt.Errorf("unable to create command: %v", err)
	}

	return &cmd, nil
}

func (db *database) GetCommandsByDate(ctx context.Context, cmdSearch entities.CommandSearch) ([]entities.EndpointCommand, error) {
	var cmds []entities.EndpointCommand

	if err := db.dbConn.WithContext(ctx).Order("cmd_id desc").Where("created_at >= ? and created_at <= ?", cmdSearch.DateFrom, cmdSearch.DateTo).
		Find(&cmds); err != nil {
		db.log.Errorf("unable to find commands by date from %v to %v with error: %v", cmdSearch.DateFrom, cmdSearch.DateTo, err)
		return nil, fmt.Errorf("unable to find commands by date from %v to %v with error: %v", cmdSearch.DateFrom, cmdSearch.DateTo, err)
	}

	return cmds, nil
}

func (db *database) TerminalReportKiosk(ctx context.Context, req entities.IncomingTermReport) ([]entities.TermReportKiosk, error) {
	var reports []entities.TermReportKiosk

	tx := db.dbConn.WithContext(ctx)

	tx = tx.Where("submitted_at >= ? and submitted_at <= ?", req.Datefrom, req.Dateto)

	if req.Provider != nil {
		tx = tx.Where("p.service_id = ?", *req.Provider)
	}

	if req.Status != nil {
		tx = tx.Where("status = ?", *req.Status)
	}

	if req.Terminal != nil {
		tx = tx.Where("status = ?", *req.Status)
	}

	tx = tx.Select("s.service_desc as provider, t.endpoint_num as terminal, count(*), sum(p.amount)").
		Table("tpayments p, tservices s, tendpoints t").
		Where("p.service_id = s.service_id and p.agent_term_id = t.endpoint_id and t.type = 100").
		Order("terminal asc")

	if err := tx.Scan(&reports).Error; err != nil {
		db.log.Errorf("unable to get report kiosk: %v", err)
		return nil, fmt.Errorf("unable to get report kiosk: %v", err)
	}

	return reports, nil
}

func (db *database) TerminalReport(ctx context.Context, req entities.IncomingTermReport) ([]entities.TermReport, error) {
	var reports []entities.TermReport

	tx := db.dbConn.WithContext(ctx)

	tx = tx.Where("p.submitted_at >= ? and p.submitted_at <= ?", req.Datefrom, req.Dateto)

	if req.Status != nil {
		tx = tx.Where("p.status = ?", *req.Status)
	}

	if req.Terminal != nil {
		tx = tx.Where("t.endpoint_num = ?", *req.Terminal)
	}

	if req.Agent != nil {
		tx = tx.Where("p.agent_term_id = ?", *req.Agent)
	}

	tx = tx.Select("t.endpoint_num, t.address, sum(p.amount)").
		Table("tendpoints t, tpayments p").
		Where("t.endpoint_id = p.agent_term_id and t.type = 100").
		Group("t.endpoint_num, t.address").
		Order("t.endpoint_num asc")

	if err := tx.Scan(&reports).Error; err != nil {
		db.log.Errorf("unable to get terminal report: %v", err)
		return nil, fmt.Errorf("unable to get terminal report: %v", err)
	}

	return reports, nil
}

func (db *database) GetTerminalsByAgent(ctx context.Context, agentID int64) ([]entities.TEndpoint, error) {
	var terminals []entities.TEndpoint

	if err := db.dbConn.WithContext(ctx).Order("endpoint_id asc").Where("agent_id = ?", agentID).
		Find(&terminals).Error; err != nil {
		db.log.Errorf("unable to get terminal by agent ID (%v) with error: %v", agentID, err)
		return nil, fmt.Errorf("unable to get terminal by agent ID (%v) with error: %v", agentID, err)
	}

	return terminals, nil
}

func (db *database) CreateCommandAuto(ctx context.Context) error {
	var (
		endpoint []entities.EndpointId
		cmd      entities.EndpointCommand
	)

	if err := db.dbConn.WithContext(ctx).Table("tendpoints").Select("endpoint_id").
		Where("endpoint_disabled = false").Find(&endpoint).Error; err != nil {
		db.log.Errorf("unable to scan tendpoints: %v", err)
		return fmt.Errorf("unable to scan tendpoints: %v", err)
	}

	now := time.Now()
	prevDay := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local)
	prevDayFromat := prevDay.Format("2006-01-02")
	userID := int64(15)

	for _, value := range endpoint {
		cmd.CreatedAt = &now
		cmd.CreatedBy = &userID
		cmd.Proceeded = false
		cmd.AgentID = 3
		cmd.Command = 7
		cmd.Param1 = &prevDayFromat
		cmd.EndpointID = value.EndpointId
		if err := db.dbConn.WithContext(ctx).Create(&cmd).Error; err != nil {
			log.Println("err to create raw: ", err.Error())
			db.log.Errorf("unable to create raw: %v", err)
			return fmt.Errorf("unable to create raw: %v", err)
		}

	}

	return nil
}
