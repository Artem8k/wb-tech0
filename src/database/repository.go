package database

import (
	"fmt"
	"log"
	"rest-api/src/database/models"
)

var Cache = []models.Order{}

type OrderRepository interface {
	// получение order по orderId из БД
	GetFromDb(uid string) models.Order

	// получение order по orderId из кеша
	Get(uid string) models.Order

	// получение order из кеша
	GetTest() models.Order

	// создание новых строк в БД в таблицах Payment, Delivery, Item, Order
	// добавление нового значения в кеш
	Create(models.Order)

	// добавление таблиц из бд в кеш при запуске приложения
	UpdateCacheOnStartUp()

	// добавление нового значения order в кеш при создании нового order
	UpdateCache(models.Order)
}

type Repository struct {
	db *Database
}

func NewOrderRepo(db *Database) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Get(uid string) models.Order {
	var o models.Order
	for _, elem := range Cache {
		if elem.OrderUid == uid {
			o = elem
		}
	}
	return o
}

func (r Repository) GetFromDb(uid string) models.Order {
	var o models.Order

	d, err := r.getDelivery(uid)

	if err != nil {
		log.Fatalln(err)
	}

	i, err := r.getItems(uid)

	if err != nil {
		log.Fatalln(err)
	}

	p, err := r.getPayment(uid)

	if err != nil {
		log.Fatalln(err)
	}

	err = r.db.Client.Get(&o,
		`SELECT * FROM "order" WHERE "uid" = $1 `, uid)

	if err != nil {
		log.Fatalln(err)
	}

	o.Delivery = d
	o.Items = i
	o.Payment = p

	return o
}

func (r Repository) Create(o models.Order) {

	tx, err := r.db.Client.Begin()
	if err != nil {
		fmt.Println(err)
	}

	ok := false
	defer func() {
		if !ok {
			tx.Rollback()
		}
	}()

	err = r.createOrder(o)

	if err != nil {
		fmt.Println(err)
	}

	err = r.createDelivery(o.Delivery, o.OrderUid)

	if err != nil {
		fmt.Println(err)
	}

	err = r.createItems(o.Items, o.OrderUid)

	if err != nil {
		fmt.Println(err)
	}

	err = r.createPayment(o.Payment, o.OrderUid)

	if err != nil {
		fmt.Println(err)
	}

	err = tx.Commit()

	if err != nil {
		fmt.Println(err)
	}

	order := r.GetFromDb(o.OrderUid)
	r.UpdateCache(order)

	ok = true
}

func (r Repository) UpdateCacheOnStartUp() {
	var o []models.Order

	err := r.db.Client.Select(&o,
		`SELECT * FROM "order"`)

	if err != nil {
		fmt.Println(err)
	}

	for _, el := range o {
		order := r.GetFromDb(el.OrderUid)
		r.UpdateCache(order)
	}
}

func (r Repository) GetTest() models.Order {
	return Cache[0]
}

func (r Repository) UpdateCache(o models.Order) {
	c := append(Cache, o)
	Cache = c
}

func (r Repository) getDelivery(uid string) (models.Delivery, error) {
	var d models.Delivery

	err := r.db.Client.Get(&d,
		`SELECT * FROM "delivery" WHERE "order_uid" = $1 `, uid)

	return d, err
}

func (r Repository) getItems(uid string) ([]models.Item, error) {
	var i []models.Item

	err := r.db.Client.Select(&i,
		`SELECT * FROM "item" WHERE "order_uid" = $1 `, uid)

	return i, err
}

func (r Repository) getPayment(uid string) (models.Payment, error) {
	var p models.Payment

	err := r.db.Client.Get(&p,
		`SELECT * FROM "payment" WHERE "order_uid" = $1 `, uid)

	return p, err
}

func (r Repository) createDelivery(d models.Delivery, order_uid string) error {

	_, err := r.db.Client.Exec(
		`INSERT INTO "delivery" (order_uid, name, phone, zip, city, address, region, email) VALUES($1,$2,$3,$4,$5,$6,$7,$8)`,
		order_uid, d.Name, d.Phone, d.Zip, d.City, d.Address, d.Region, d.Email)

	return err
}

func (r Repository) createItems(i []models.Item, order_uid string) error {
	var err error

	for _, elem := range i {
		_, err = r.db.Client.Exec(
			`INSERT INTO "item" (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
			VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`,
			order_uid, elem.Chrt_id, elem.Track_number, elem.Price, elem.Rid, elem.Name, elem.Sale, elem.Size, elem.Total_price, elem.Nm_id, elem.Brand, elem.Status)
	}

	return err
}

func (r Repository) createPayment(p models.Payment, order_uid string) error {

	_, err := r.db.Client.Exec(
		`INSERT INTO "payment" (order_uid, transaction, currency, provider, request_id, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		order_uid, p.Transaction, p.Currency, p.Provider, p.Request_id, p.Amount, p.Payment_dt, p.Bank, p.Delivery_cost, p.Goods_total, p.Custom_fee)

	return err
}

func (r Repository) createOrder(o models.Order) error {

	_, err := r.db.Client.Exec(
		`INSERT INTO "order" (uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		o.OrderUid, o.Track_number, o.Entry, o.Locale, o.Internal_signature, o.Customer_id, o.Delivery_service, o.Shardkey, o.Sm_id, o.Date_created, o.Oof_shard)

	return err
}
