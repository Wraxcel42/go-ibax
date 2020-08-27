/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
	Cycle        int64           `gorm:"not null" `           //
	Amount       decimal.Decimal `gorm:"not null default 0" ` //
	Expired      int64           `gorm:"null" `
	Status       int64           `gorm:"null"`            //
	Review       int64           `gorm:"null default 0" ` //
	Count        int64           `gorm:"null default 0" ` //
	Stakes       int64           `gorm:"null default 0" ` //
	Transfers    int64           `gorm:"null"  `          //
	Stime        int64           `gorm:"not null" `       //
	Etime        int64           `gorm:"not null" `       //
	DateUpdated  int64           `gorm:"not null" `
	DateCreated  int64           `gorm:"not null" `
}

// TableName returns name of table
func (MineStake) TableName() string {
	return `1_mine_stake`
}

// Get is retrieving model from database
func (m *MineStake) GetActiveMiner(time, availableStatus int64) (mp []MineStake, err error) {
	err = DBConn.Table(m.TableName()).
		Where("stime <= ? and etime >=? and status = ?", time, time, availableStatus).
		Order("devid asc").
		Scan(&mp).Error
	return mp, err
}

// Get is retrieving model from database
func (m *MineStake) GetExpiredMiner(time int64) (mp []MineStake, err error) {
	err = DBConn.Table(m.TableName()).
		Where("etime <=? and expired = 0", time).
		Order("etime asc").
		Limit(10).
		Scan(&mp).Error
	return mp, err
}

// Get is retrieving model from database
func (m *MineStake) UpdateExpired(t int64) error {
	m.Expired = 1
	m.DateUpdated = t
	return DBConn.Model(m).Updates(map[string]interface{}{"expired": m.Expired, "date_updated": m.DateUpdated}).Error
}

// Get is retrieving model from database
func (m *MineStake) Update() error {
	return DBConn.Save(m).Error
}