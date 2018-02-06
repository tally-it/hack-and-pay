// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v6"
)

// Transaction is an object representing the database table.
type Transaction struct {
	PaymentID int         `boil:"payment_id" json:"payment_id" toml:"payment_id" yaml:"payment_id"`
	UserID    int         `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	SKUID     null.Int    `boil:"SKU_id" json:"SKU_id,omitempty" toml:"SKU_id" yaml:"SKU_id,omitempty"`
	Value     string      `boil:"value" json:"value" toml:"value" yaml:"value"`
	Tag       null.String `boil:"tag" json:"tag,omitempty" toml:"tag" yaml:"tag,omitempty"`
	AddedAt   time.Time   `boil:"added_at" json:"added_at" toml:"added_at" yaml:"added_at"`
	UpdatedAt null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *transactionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L transactionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TransactionColumns = struct {
	PaymentID string
	UserID    string
	SKUID     string
	Value     string
	Tag       string
	AddedAt   string
	UpdatedAt string
}{
	PaymentID: "payment_id",
	UserID:    "user_id",
	SKUID:     "SKU_id",
	Value:     "value",
	Tag:       "tag",
	AddedAt:   "added_at",
	UpdatedAt: "updated_at",
}

// transactionR is where relationships are stored.
type transactionR struct {
	User *User
}

// transactionL is where Load methods for each relationship are stored.
type transactionL struct{}

var (
	transactionColumns               = []string{"payment_id", "user_id", "SKU_id", "value", "tag", "added_at", "updated_at"}
	transactionColumnsWithoutDefault = []string{"user_id", "SKU_id", "value", "tag", "updated_at"}
	transactionColumnsWithDefault    = []string{"payment_id", "added_at"}
	transactionPrimaryKeyColumns     = []string{"payment_id"}
)

type (
	// TransactionSlice is an alias for a slice of pointers to Transaction.
	// This should generally be used opposed to []Transaction.
	TransactionSlice []*Transaction

	transactionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	transactionType                 = reflect.TypeOf(&Transaction{})
	transactionMapping              = queries.MakeStructMapping(transactionType)
	transactionPrimaryKeyMapping, _ = queries.BindMapping(transactionType, transactionMapping, transactionPrimaryKeyColumns)
	transactionInsertCacheMut       sync.RWMutex
	transactionInsertCache          = make(map[string]insertCache)
	transactionUpdateCacheMut       sync.RWMutex
	transactionUpdateCache          = make(map[string]updateCache)
	transactionUpsertCacheMut       sync.RWMutex
	transactionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single transaction record from the query, and panics on error.
func (q transactionQuery) OneP() *Transaction {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single transaction record from the query.
func (q transactionQuery) One() (*Transaction, error) {
	o := &Transaction{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for transactions")
	}

	return o, nil
}

// AllP returns all Transaction records from the query, and panics on error.
func (q transactionQuery) AllP() TransactionSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Transaction records from the query.
func (q transactionQuery) All() (TransactionSlice, error) {
	var o []*Transaction

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Transaction slice")
	}

	return o, nil
}

// CountP returns the count of all Transaction records in the query, and panics on error.
func (q transactionQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Transaction records in the query.
func (q transactionQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count transactions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q transactionQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q transactionQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if transactions exists")
	}

	return count > 0, nil
}

// UserG pointed to by the foreign key.
func (o *Transaction) UserG(mods ...qm.QueryMod) userQuery {
	return o.User(boil.GetDB(), mods...)
}

// User pointed to by the foreign key.
func (o *Transaction) User(exec boil.Executor, mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("user_id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(exec, queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
} // LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (transactionL) LoadUser(e boil.Executor, singular bool, maybeTransaction interface{}) error {
	var slice []*Transaction
	var object *Transaction

	count := 1
	if singular {
		object = maybeTransaction.(*Transaction)
	} else {
		slice = *maybeTransaction.(*[]*Transaction)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &transactionR{}
		}
		args[0] = object.UserID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &transactionR{}
			}
			args[i] = obj.UserID
		}
	}

	query := fmt.Sprintf(
		"select * from `users` where `user_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}
	defer results.Close()

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.User = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.User = foreign
				break
			}
		}
	}

	return nil
}

// SetUserG of the transaction to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Transactions.
// Uses the global database handle.
func (o *Transaction) SetUserG(insert bool, related *User) error {
	return o.SetUser(boil.GetDB(), insert, related)
}

// SetUserP of the transaction to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Transactions.
// Panics on error.
func (o *Transaction) SetUserP(exec boil.Executor, insert bool, related *User) {
	if err := o.SetUser(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetUserGP of the transaction to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Transactions.
// Uses the global database handle and panics on error.
func (o *Transaction) SetUserGP(insert bool, related *User) {
	if err := o.SetUser(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetUser of the transaction to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Transactions.
func (o *Transaction) SetUser(exec boil.Executor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `transactions` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, transactionPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.PaymentID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.UserID

	if o.R == nil {
		o.R = &transactionR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Transactions: TransactionSlice{o},
		}
	} else {
		related.R.Transactions = append(related.R.Transactions, o)
	}

	return nil
}

// TransactionsG retrieves all records.
func TransactionsG(mods ...qm.QueryMod) transactionQuery {
	return Transactions(boil.GetDB(), mods...)
}

// Transactions retrieves all the records using an executor.
func Transactions(exec boil.Executor, mods ...qm.QueryMod) transactionQuery {
	mods = append(mods, qm.From("`transactions`"))
	return transactionQuery{NewQuery(exec, mods...)}
}

// FindTransactionG retrieves a single record by ID.
func FindTransactionG(paymentID int, selectCols ...string) (*Transaction, error) {
	return FindTransaction(boil.GetDB(), paymentID, selectCols...)
}

// FindTransactionGP retrieves a single record by ID, and panics on error.
func FindTransactionGP(paymentID int, selectCols ...string) *Transaction {
	retobj, err := FindTransaction(boil.GetDB(), paymentID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindTransaction retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTransaction(exec boil.Executor, paymentID int, selectCols ...string) (*Transaction, error) {
	transactionObj := &Transaction{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `transactions` where `payment_id`=?", sel,
	)

	q := queries.Raw(exec, query, paymentID)

	err := q.Bind(transactionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from transactions")
	}

	return transactionObj, nil
}

// FindTransactionP retrieves a single record by ID with an executor, and panics on error.
func FindTransactionP(exec boil.Executor, paymentID int, selectCols ...string) *Transaction {
	retobj, err := FindTransaction(exec, paymentID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Transaction) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Transaction) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Transaction) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Transaction) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no transactions provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.UpdatedAt.Time.IsZero() {
		o.UpdatedAt.Time = currTime
		o.UpdatedAt.Valid = true
	}

	nzDefaults := queries.NonZeroDefaultSet(transactionColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	transactionInsertCacheMut.RLock()
	cache, cached := transactionInsertCache[key]
	transactionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			transactionColumns,
			transactionColumnsWithDefault,
			transactionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(transactionType, transactionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(transactionType, transactionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `transactions` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `transactions` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `transactions` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, transactionPrimaryKeyColumns))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into transactions")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.PaymentID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == transactionMapping["PaymentID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PaymentID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for transactions")
	}

CacheNoHooks:
	if !cached {
		transactionInsertCacheMut.Lock()
		transactionInsertCache[key] = cache
		transactionInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Transaction record. See Update for
// whitelist behavior description.
func (o *Transaction) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Transaction record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Transaction) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Transaction, and panics on error.
// See Update for whitelist behavior description.
func (o *Transaction) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Transaction.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Transaction) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	key := makeCacheKey(whitelist, nil)
	transactionUpdateCacheMut.RLock()
	cache, cached := transactionUpdateCache[key]
	transactionUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			transactionColumns,
			transactionPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update transactions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `transactions` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, transactionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(transactionType, transactionMapping, append(wl, transactionPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update transactions row")
	}

	if !cached {
		transactionUpdateCacheMut.Lock()
		transactionUpdateCache[key] = cache
		transactionUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q transactionQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q transactionQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for transactions")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TransactionSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o TransactionSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o TransactionSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TransactionSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `transactions` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, transactionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in transaction slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Transaction) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Transaction) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Transaction) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Transaction) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no transactions provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	nzDefaults := queries.NonZeroDefaultSet(transactionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	transactionUpsertCacheMut.RLock()
	cache, cached := transactionUpsertCache[key]
	transactionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			transactionColumns,
			transactionColumnsWithDefault,
			transactionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			transactionColumns,
			transactionPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert transactions, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "transactions", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `transactions` WHERE `payment_id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(transactionType, transactionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(transactionType, transactionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for transactions")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.PaymentID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == transactionMapping["PaymentID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PaymentID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for transactions")
	}

CacheNoHooks:
	if !cached {
		transactionUpsertCacheMut.Lock()
		transactionUpsertCache[key] = cache
		transactionUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single Transaction record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Transaction) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Transaction record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Transaction) DeleteG() error {
	if o == nil {
		return errors.New("models: no Transaction provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Transaction record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Transaction) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Transaction record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Transaction) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Transaction provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), transactionPrimaryKeyMapping)
	sql := "DELETE FROM `transactions` WHERE `payment_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from transactions")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q transactionQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q transactionQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no transactionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from transactions")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o TransactionSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o TransactionSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Transaction slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o TransactionSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TransactionSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Transaction slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `transactions` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, transactionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from transaction slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Transaction) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Transaction) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Transaction) ReloadG() error {
	if o == nil {
		return errors.New("models: no Transaction provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Transaction) Reload(exec boil.Executor) error {
	ret, err := FindTransaction(exec, o.PaymentID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *TransactionSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *TransactionSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TransactionSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty TransactionSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TransactionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	transactions := TransactionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `transactions`.* FROM `transactions` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, transactionPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&transactions)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TransactionSlice")
	}

	*o = transactions

	return nil
}

// TransactionExists checks if the Transaction row exists.
func TransactionExists(exec boil.Executor, paymentID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `transactions` where `payment_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, paymentID)
	}

	row := exec.QueryRow(sql, paymentID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if transactions exists")
	}

	return exists, nil
}

// TransactionExistsG checks if the Transaction row exists.
func TransactionExistsG(paymentID int) (bool, error) {
	return TransactionExists(boil.GetDB(), paymentID)
}

// TransactionExistsGP checks if the Transaction row exists. Panics on error.
func TransactionExistsGP(paymentID int) bool {
	e, err := TransactionExists(boil.GetDB(), paymentID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// TransactionExistsP checks if the Transaction row exists. Panics on error.
func TransactionExistsP(exec boil.Executor, paymentID int) bool {
	e, err := TransactionExists(exec, paymentID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
