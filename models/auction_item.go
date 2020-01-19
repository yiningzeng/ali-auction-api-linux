package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type AuctionItem struct {
	Id                int64       `orm:"column(id);pk"`
	AuctionTargetId   int64     `orm:"column(auction_target_id)"`
	MnNotice          bool      `orm:"column(mn_notice)"`
	Credit            bool      `orm:"column(credit)"`
	ItemUrl           string    `orm:"column(item_url)"`
	Status            string    `orm:"column(status)"`
	Title             string    `orm:"column(title)"`
	PicUrl            string    `orm:"column(pic_url)"`
	InitialPrice      float64   `orm:"column(initial_price)"`
	CurrentPrice      float64   `orm:"column(current_price)"`
	ConsultPrice      float64   `orm:"column(consult_price)"`
	MarketPrice       float64   `orm:"column(market_price)"`
	SellOff           bool      `orm:"column(sell_off)"`
	Start             int64     `orm:"column(start)"`
	End               int64     `orm:"column(end)"`
	TimeToStart       int64     `orm:"column(time_to_start)"`
	TimeToEnd         int64     `orm:"column(time_to_end)"`
	ViewerCount       int64     `orm:"column(viewer_count)"`
	BidCount          int64     `orm:"column(bid_count)"`
	DelayCount        int64     `orm:"column(delay_count)"`
	ApplyCount        int64     `orm:"column(apply_count)"`
	CatNames          string    `orm:"column(cat_names)"`
	CollateralCatName string    `orm:"column(collateral_cat_name)"`
	XmppVersion       string    `orm:"column(xmpp_version)"`
	BuyRestrictions   int64     `orm:"column(buy_restrictions)"`
	SupportLoans      int64     `orm:"column(support_loans)"`
	SupportOrgLoan    int64     `orm:"column(support_org_loan)"`
	UpdateTime        time.Time `orm:"column(update_time);type(timestamp with time zone)"`
}

func (t *AuctionItem) TableName() string {
	return "auction_item"
}

func init() {
	orm.RegisterModel(new(AuctionItem))
}

// AddAuctionItem insert a new AuctionItem into database and returns
// last inserted Id on success.
func AddAuctionItem(m *AuctionItem) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAuctionItemById retrieves AuctionItem by Id. Returns error if
// Id doesn't exist
func GetAuctionItemById(id int64) (v *AuctionItem, err error) {
	o := orm.NewOrm()
	v = &AuctionItem{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAuctionItem retrieves all AuctionItem matches certain condition. Returns empty list if
// no records exist
func GetAllAuctionItem(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AuctionItem))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []AuctionItem
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateAuctionItem updates AuctionItem by Id and returns error if
// the record to be updated doesn't exist
func UpdateAuctionItemById(m *AuctionItem) (err error) {
	o := orm.NewOrm()
	v := AuctionItem{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAuctionItem deletes AuctionItem by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAuctionItem(id int64) (err error) {
	o := orm.NewOrm()
	v := AuctionItem{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AuctionItem{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
